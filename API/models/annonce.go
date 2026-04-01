package models

type Annonce struct {
	Id               int     `json:"id"`
	Titre            string  `json:"titre"`
	Description      string  `json:"description"`
	Type             string  `json:"type"` // Don ou Vente
	Prix             float64 `json:"prix"`
	StatutValidation string  `json:"statut_validation"` // En attente, Validé, Refusé
	CodePostal       string  `json:"code_postal"`
	Ville            string  `json:"ville"`
	Prenom           string  `json:"prenom"`
	Nom              string  `json:"nom"`
	Categorie        string  `json:"categorie"`
}

type Categorie struct {
	Id      int    `json:"id"`
	Libelle string `json:"libelle"`
}