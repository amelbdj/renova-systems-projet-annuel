package models

type Article struct {
	Id           int    `json:"id"`
	IdSalarie    int    `json:"id_salarie"`
	NomAuteur    string `json:"nom_auteur"`
	PrenomAuteur string `json:"prenom_auteur"`
	Titre        string `json:"titre"`
	Slug         string `json:"slug"` // Si tu as ajouté la colonne slug
	Contenu      string `json:"contenu"`
	ImageUrl     string `json:"image_url"`
	Type         string `json:"type"`   // news, conseil, pedagogique
	Statut       string `json:"statut"` // brouillon, valide, en attente, refuse
	CreatedAt    string `json:"created_at"`
}