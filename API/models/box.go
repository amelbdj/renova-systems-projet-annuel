package models

type DepositRequest struct {
	AnnonceId   int64 `json:"annonce_id"`
	ConteneurId int64 `json:"conteneur_id"`
}

type Box struct {
	Id           int    `json:"id"`
	Localisation string `json:"localisation"`
	Type         string `json:"type"`
	Etat         string `json:"etat"`
	Capacite     int    `json:"capacite"`
}
