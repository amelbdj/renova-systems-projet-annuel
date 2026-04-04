package models

type DepositRequest struct {
	AnnonceId   int64 `json:"annonce_id"`
	ConteneurId int64 `json:"conteneur_id"`
}

type Box struct {
	Id           int    `json:"id"`
	Localisation string `json:"localisation"`
	Etat         string `json:"etat"`
}
