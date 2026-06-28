package models

type Annonce struct {
	Id                 int     `json:"id"`
	Titre              string  `json:"titre"`
	Description        string  `json:"description"`
	Type               string  `json:"type"`
	Prix               float64 `json:"prix"`
	StatutValidation   string  `json:"statut_validation"`
	CodePostal         string  `json:"code_postal"`
	Ville              string  `json:"ville"`
	Prenom             string  `json:"prenom"`
	Nom                string  `json:"nom"`
	Categorie          string  `json:"categorie"`
	CreatedAt          string  `json:"created_at"`
	Etat               string  `json:"etat"`
	PoidsKg            float64 `json:"poids_kg"`
	Quantite           int     `json:"quantite"`
	CommissionPrelevee float64 `json:"commission_prelevee"`
	PhotoUrl           string  `json:"photo_url"`
	IdUser             int     `json:"id_user"`
	IdCategorie        int     `json:"id_categorie"`
	Image              string  `json:"image"`
	StatutVente        string  `json:"statut_vente"`
	IsSponsored        bool    `json:"is_sponsored"`
	PlanAbo            string  `json:"plan_abo"`
}

type Categorie struct {
	Id      int    `json:"id"`
	Libelle string `json:"libelle"`
}
