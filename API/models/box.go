package models

type DepositRequest struct {
	AnnonceId   int64 `json:"annonce_id"`
	ConteneurId int64 `json:"conteneur_id"`
}

type Conteneur struct {
	Id           int    `json:"id"`
	Nom          string `json:"nom"`
	Adresse      string `json:"adresse"`
	DateCreation string `json:"date_creation"`
}

type Box struct {
	Id          int    `json:"id"`
	IdConteneur int    `json:"id_conteneur"`
	Numero      int    `json:"numero"`
	Taille      string `json:"taille"`
	Statut      string `json:"statut"`

	CodeSecret *string `json:"code_secret"`
}

type ConteneurAvecStats struct {
	Id         int    `json:"id"`
	Nom        string `json:"nom"`
	Adresse    string `json:"adresse"`
	TotalBoxes int    `json:"total_boxes"`
}
