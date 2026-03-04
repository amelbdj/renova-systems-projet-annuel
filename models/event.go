package models

import "time"

type Evenement struct {
	ID               int       `json:"id"`
	Titre            string    `json:"titre"`
	Description      string    `json:"description"`
	DateDebut        time.Time `json:"date_debut"`
	DateFin          time.Time `json:"date_fin"`
	NbPlaces         int       `json:"nb_places"`
	StatutValidation string    `json:"statut_validation"`
	Format           string    `json:"format"` // En ligne ou Présentiel
	IdSalarie        int       `json:"id_salarie"`
}