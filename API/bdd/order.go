package bdd

import (
	"fmt"
	"sort"
	"upcycleconnect/models"
)

func CreateOrder(annonce models.Annonce, acheteurId int) (int, error) {
	montant := annonce.Prix
	commission := montant * 0.05
	var StatutVente string
	var orderId int64

	err := Db.QueryRow("SELECT statut_vente FROM pa2026.annonce WHERE id = ?", annonce.Id).Scan(&StatutVente)
	if err != nil {
		return 0, fmt.Errorf("CreateOrder (Select) : %s", err.Error())
	}

	if StatutVente != "VENDU" {

		res, err := Db.Exec("INSERT INTO pa2026.order (id_annonce, id_acheteur, montant_total, commission, date_commande, `type`) VALUES (?, ?, ?, ?, NOW(), 'annonce')", annonce.Id, acheteurId, montant, commission)
		if err != nil {
			return 0, fmt.Errorf("CreateOrder (Insert) : %s", err.Error())
		}

		orderId, _ = res.LastInsertId()

		_, err = Db.Exec("UPDATE pa2026.annonce SET statut_vente = 'VENDU' WHERE id = ?", annonce.Id)
		if err != nil {
			return 0, fmt.Errorf("CreateOrder (Update VENDU) : %s", err.Error())
		}
	} else {
		return 0, fmt.Errorf("Cet objet a déjà été vendu")
	}

	return int(orderId), nil
}

func planNomEtPrix(idPlan int) (string, float64) {
	switch idPlan {
	case 2:
		return "Plus", 45
	case 3:
		return "Pro", 99
	default:
		return "Premium", 25
	}
}

func PaymentHistory(userID int) ([]map[string]interface{}, error) {
	var history []map[string]interface{}

	achats, err := Db.Query("SELECT o.montant_total, o.date_commande, a.titre, COALESCE((SELECT url_pdf FROM document WHERE id_commande = o.id_commande AND type_doc = 'facture' ORDER BY id_document DESC LIMIT 1), '') FROM pa2026.`order` o JOIN annonce a ON o.id_annonce = a.id WHERE o.id_acheteur = ?", userID)
	if err != nil {
		return nil, err
	}
	for achats.Next() {
		var montant float64
		var date, titre, urlPdf string
		achats.Scan(&montant, &date, &titre, &urlPdf)
		history = append(history, map[string]interface{}{
			"type": "achat", "titre": titre, "montant": montant, "date": date, "url_pdf": urlPdf,
		})
	}
	achats.Close()

	ventes, err := Db.Query("SELECT o.montant_total, o.date_commande, a.titre, COALESCE((SELECT url_pdf FROM document WHERE id_commande = o.id_commande AND type_doc = 'facture' ORDER BY id_document DESC LIMIT 1), '') FROM pa2026.`order` o JOIN annonce a ON o.id_annonce = a.id WHERE a.id_user = ?", userID)
	if err != nil {
		return nil, err
	}
	for ventes.Next() {
		var montant float64
		var date, titre, urlPdf string
		ventes.Scan(&montant, &date, &titre, &urlPdf)
		history = append(history, map[string]interface{}{
			"type": "vente", "titre": titre, "montant": montant, "date": date, "url_pdf": urlPdf,
		})
	}
	ventes.Close()

	abos, err := Db.Query("SELECT a.id_plan, a.date_debut, COALESCE((SELECT url_pdf FROM document WHERE id_commande = a.id_abonnement AND type_doc = 'contrat' ORDER BY id_document DESC LIMIT 1), '') FROM abonnement a WHERE a.id_user = ?", userID)
	if err != nil {
		return nil, err
	}
	for abos.Next() {
		var idPlan int
		var date, urlPdf string
		abos.Scan(&idPlan, &date, &urlPdf)
		nom, montant := planNomEtPrix(idPlan)
		history = append(history, map[string]interface{}{
			"type": "abonnement", "titre": "Abonnement " + nom, "montant": montant, "date": date, "url_pdf": urlPdf,
		})
	}
	abos.Close()

	sort.Slice(history, func(i, j int) bool {
		return history[i]["date"].(string) > history[j]["date"].(string)
	})

	return history, nil
}

func GetProInvoices(userID int) ([]map[string]interface{}, error) {
	query := `
        SELECT
            o.id_commande,
            o.date_commande,
            COALESCE(a.titre, e.titre, 'Transaction') AS titre,
            o.montant_total,
            o.commission,
            COALESCE(o.type, 'annonce') AS type
        FROM pa2026.order o
        LEFT JOIN pa2026.annonce a ON o.type = 'annonce' AND o.id_annonce = a.id
        LEFT JOIN pa2026.evenement e ON o.type = 'evenement' AND o.id_annonce = e.id
        WHERE o.id_acheteur = ?
        ORDER BY o.date_commande DESC`

	rows, err := Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var factures []map[string]interface{}

	for rows.Next() {
		var id int
		var date string
		var titre string
		var montant float64
		var commission float64
		var typeCommande string

		err := rows.Scan(&id, &date, &titre, &montant, &commission, &typeCommande)
		if err != nil {
			continue
		}

		factures = append(factures, map[string]interface{}{
			"id":         id,
			"date":       date,
			"titre":      titre,
			"montant":    montant,
			"commission": commission,
			"type":       typeCommande,
		})
	}

	return factures, nil
}

func GetFinanceOverviewMois() (float64, float64, error) {
	query := `
        SELECT 
            COALESCE(SUM(montant_total), 0) AS total_volume,
            COALESCE(SUM(commission), 0) AS total_commission
        FROM pa2026.order 
        WHERE MONTH(date_commande) = MONTH(CURRENT_DATE()) 
        AND YEAR(date_commande) = YEAR(CURRENT_DATE())`

	var volume float64
	var commission float64

	err := Db.QueryRow(query).Scan(&volume, &commission)
	if err != nil {
		return 0, 0, fmt.Errorf("erreur SQL GetFinanceOverviewMois : %v", err)
	}

	return volume, commission, nil
}

func GetAdminTransactions() ([]map[string]interface{}, error) {

	query := `
        SELECT
            o.id_commande,
            o.date_commande,
            COALESCE(a.titre, e.titre, 'Transaction') AS titre,
            o.montant_total,
            o.commission,
            COALESCE(o.type, 'annonce') AS type,
            COALESCE(p.statut, 'payé') AS statut
        FROM pa2026.order o
        LEFT JOIN pa2026.annonce a ON o.type = 'annonce' AND o.id_annonce = a.id
        LEFT JOIN pa2026.evenement e ON o.type = 'evenement' AND o.id_annonce = e.id
        LEFT JOIN pa2026.paiement p ON p.id_commande = o.id_commande
        ORDER BY o.date_commande DESC
        LIMIT 50`

	rows, err := Db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Erreur SQL Transactions: %v", err)
	}
	defer rows.Close()

	var transactions []map[string]interface{}

	for rows.Next() {
		var id int
		var date string
		var titre string
		var montant float64
		var commission float64
		var typeCommande string
		var statut string

		err := rows.Scan(&id, &date, &titre, &montant, &commission, &typeCommande, &statut)
		if err != nil {
			continue
		}

		item := map[string]interface{}{
			"id":         id,
			"date":       date,
			"titre":      titre,
			"montant":    montant,
			"commission": commission,
			"type":       typeCommande,
			"statut":     statut,
		}
		transactions = append(transactions, item)
	}

	abos, err := Db.Query(`
        SELECT
            a.id_abonnement,
            a.date_debut,
            a.id_plan,
            COALESCE(a.statut, 'actif') AS statut,
            COALESCE(u.email, 'Utilisateur') AS email
        FROM pa2026.abonnement a
        LEFT JOIN pa2026.utilisateur u ON a.id_user = u.id
        ORDER BY a.date_debut DESC
        LIMIT 50`)
	if err != nil {
		return nil, fmt.Errorf("Erreur SQL Abonnements: %v", err)
	}
	defer abos.Close()

	for abos.Next() {
		var id int
		var date string
		var idPlan int
		var statut string
		var email string

		err := abos.Scan(&id, &date, &idPlan, &statut, &email)
		if err != nil {
			continue
		}

		nom, montant := planNomEtPrix(idPlan)

		item := map[string]interface{}{
			"id":         id,
			"date":       date,
			"titre":      "Abonnement " + nom + " - " + email,
			"montant":    montant,
			"commission": montant,
			"type":       "abonnement",
			"statut":     statut,
		}
		transactions = append(transactions, item)
	}

	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i]["date"].(string) > transactions[j]["date"].(string)
	})

	if len(transactions) > 50 {
		transactions = transactions[:50]
	}

	return transactions, nil
}
