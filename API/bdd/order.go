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
