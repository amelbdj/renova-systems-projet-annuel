package models

type Projet struct {
	Id          int     `json:"id"`
	IdUser      int     `json:"id_user"`
	Titre       string  `json:"titre"`
	Description string  `json:"description"`
	Categorie   string  `json:"categorie"`
	Statut      string  `json:"statut"`
	PhotoAvant  string  `json:"photo_avant"`
	PhotoApres  string  `json:"photo_apres"`
	Co2Evite    float64 `json:"co2_evite"`
	CreatedAt   string  `json:"created_at"`
}