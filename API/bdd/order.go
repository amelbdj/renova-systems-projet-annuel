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

    achats, err := Db.Query("SELECT o.montant, o.date, a.titre FROM pa2026.`order` o JOIN annonce a ON o.annonce_id = a.id WHERE o.acheteur_id = ?", userID)
    if err != nil {
        return nil, err
    }
    for achats.Next() {
        var montant float64
        var date, titre string
        achats.Scan(&montant, &date, &titre)
        history = append(history, map[string]interface{}{
            "type": "achat", "titre": titre, "montant": montant, "date": date,
        })
    }
    achats.Close()

    ventes, err := Db.Query("SELECT o.montant, o.date, a.titre FROM pa2026.`order` o JOIN annonce a ON o.annonce_id = a.id WHERE a.id_user = ?", userID)
    if err != nil {
        return nil, err
    }
    for ventes.Next() {
        var montant float64
        var date, titre string
        ventes.Scan(&montant, &date, &titre)
        history = append(history, map[string]interface{}{
            "type": "vente", "titre": titre, "montant": montant, "date": date,
        })
    }
    ventes.Close()

    abos, err := Db.Query("SELECT id_plan, date_debut FROM abonnement WHERE id_user = ?", userID)
    if err != nil {
        return nil, err
    }
    for abos.Next() {
        var idPlan int
        var date string
        abos.Scan(&idPlan, &date)
        nom, montant := planNomEtPrix(idPlan)
        history = append(history, map[string]interface{}{
            "type": "abonnement", "titre": "Abonnement " + nom, "montant": montant, "date": date,
        })
    }
    abos.Close()

    sort.Slice(history, func(i, j int) bool {
        return history[i]["date"].(string) > history[j]["date"].(string)
    })

    return history, nil
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

	return transactions, nil
}
