package models

type Evenement struct {
	Id               int    `json:"id"`
	Titre            string `json:"titre"`
	Description      string `json:"description"`
	DateDebut        string `json:"date_debut"`
	DateFin          string `json:"date_fin"`
	NbPlaces         int    `json:"nb_places"`
	StatutValidation string `json:"statut_validation"`
	Format           string `json:"format"` // En ligne ou Présentiel
	NomSalarie       string `json:"nomSalarie"`
	PrenomSalarie    string `json:"prenomSalarie"`
	Lieu             string `json:"lieu"`
	Type             string `json:"type"`
	IdSalarie        int    `json:"idSalarie"` // ID du salarié qui a créé l'événement

}