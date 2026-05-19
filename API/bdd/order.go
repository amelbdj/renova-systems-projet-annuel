package bdd

import (
	"fmt"
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
		res, err := Db.Exec("INSERT INTO pa2026.order (annonce_id, acheteur_id, montant, commission, date) VALUES (?, ?, ?, ?, NOW())", annonce.Id, acheteurId, montant, commission)
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

func PaymentHistory(userID int) ([]map[string]interface{}, error) {
	query := `
        SELECT 
            o.montant, 
            o.date, 
            a.titre 
        FROM pa2026.order o
        JOIN annonce a ON o.annonce_id = a.id
        WHERE a.id_user = ?
        ORDER BY o.date DESC`

	rows, err := Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []map[string]interface{}
	for rows.Next() {
		var montant float64
		var date, titre string
		rows.Scan(&montant, &date, &titre)

		item := map[string]interface{}{
			"titre":   titre,
			"montant": montant,
			"date":    date,
		}
		history = append(history, item)
	}
	return history, nil
}

// GetFinanceOverviewMois récupère le volume total et la commission du mois en cours
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
	// On récupère les infos de l'order ET le titre de l'annonce
	query := `
		SELECT 
			o.id_commande, 
			o.date_commande, 
			a.titre, 
			o.montant_total, 
			o.commission 
		FROM pa2026.order o
		JOIN pa2026.annonce a ON o.id_annonce = a.id
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

		err := rows.Scan(&id, &date, &titre, &montant, &commission)
		if err != nil {
			continue // S'il y a une erreur sur une ligne, on passe à la suivante
		}

		// On construit notre objet JSON
		item := map[string]interface{}{
			"id":         id,
			"date":       date,
			"titre":      titre,
			"montant":    montant,
			"commission": commission,
		}
		transactions = append(transactions, item)
	}
	
	return transactions, nil
}