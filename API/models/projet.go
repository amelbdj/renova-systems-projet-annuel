package models

type Projet struct {
	Id          int     `json:"id"`
	IdUser      int     `json:"id_user"`
	Titre       string  `json:"titre"`
	Description string  `json:"description"`
	Categorie   string  `json:"categorie"`
	Statut      string  `json:"statut"`
	AvantDesc   string  `json:"avant_desc"`
	ApresDesc   string  `json:"apres_desc"`
	Co2Evite    float64 `json:"co2_evite"`
	CreatedAt   string  `json:"created_at"`
	Photo       string  `json:"photo"`
}