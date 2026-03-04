package models

type Annonce struct {
	ID               int     `json:"id"`
	Titre            string  `json:"titre"`
	Description      string  `json:"description"`
	Type             string  `json:"type"` // Don ou Vente
	Prix             float64 `json:"prix"`
	StatutValidation string  `json:"statut_validation"` // En attente, Validé, Refusé
	CodePostal       string  `json:"code_postal"`
	Ville            string  `json:"ville"`
	IdUser           int     `json:"id_user"`
	IdCategorie      int     `json:"id_categorie"`
}

type Categorie struct {
	ID      int    `json:"id"`
	Libelle string `json:"libelle"`
}