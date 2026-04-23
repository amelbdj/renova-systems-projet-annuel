package models

import "time"

type Annonce struct {
	Id                 int       `json:"id"`
	Titre              string    `json:"titre"`
	Description        string    `json:"description"`
	Type               string    `json:"type"` // Don ou Vente
	Prix               float64   `json:"prix"`
	StatutValidation   string    `json:"statut_validation"` // En attente, Validé, Refusé
	CodePostal         string    `json:"code_postal"`
	Ville              string    `json:"ville"`
	Prenom             string    `json:"prenom"`
	Nom                string    `json:"nom"`
	Categorie          string    `json:"categorie"`
	CreatedAt          time.Time `json:"created_at"`
	Etat               string    `json:"etat"`     // 'Neuf', 'Bon état', etc.
	PoidsKg            float64   `json:"poids_kg"` // Pour l'impact écologique
	Quantite           int       `json:"quantite"`
	CommissionPrelevee float64   `json:"commission_prelevee"` // Commission 5-10%
	PhotoUrl           string    `json:"photo_url"`
	IdUser             int       `json:"id_user"`
	IdCategorie        int       `json:"id_categorie"`
	Image              string    `json:"image"`
	StatutVente        string    `json:"statut_vente"`
}

type Categorie struct {
	Id      int    `json:"id"`
	Libelle string `json:"libelle"`
}
