package models

import "time"

type Order struct {
	Id         int       `json:"id"`
	AnnonceId  int       `json:"annonce_id"`
	AcheteurId int       `json:"acheteur_id"`
	Montant    float64   `json:"montant"`
	Commission float64   `json:"commission"`
	Date       time.Time `json:"date"`
}
