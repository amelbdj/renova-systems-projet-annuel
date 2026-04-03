package bdd

import (
	"fmt"
	"upcycleconnect/models"
)

func CreateOrder(annonce models.Annonce, acheteurId int) (int,error) {

	montant := annonce.Prix
	commission := montant * 0.05
	var StatutVente string
	var orderId int64


	err := Db.QueryRow("SELECT statut FROM pa2026.annonce WHERE id = ?", annonce.Id).Scan(&StatutVente)

	if err != nil {
	
		return 0, fmt.Errorf("CreateOrder : %s", err.Error())
		
	}

	if StatutVente == "En vente" {

		res, err := Db.Exec("INSERT INTO pa2026.order (annonce_id, acheteur_id, montant, commission, date) VALUES (?, ?, ?, ?, NOW())", annonce.Id, acheteurId, montant, commission)

		if err != nil {
		return 0, fmt.Errorf("CreateOrder : %s", err.Error())
		}

		orderId, _ = res.LastInsertId()
		_, err = Db.Exec("UPDATE pa2026.annonce SET statut_vente = 'EN ATTENTE DEPOT' WHERE id = ?", annonce.Id)

		if err != nil {
			return 0, fmt.Errorf("CreateOrder : %s", err.Error())
		}
	}
	return int (orderId), nil

}