package models

type DepositRequest struct {
	AnnonceId   int64 `json:"annonce_id"`
	ConteneurId int64 `json:"conteneur_id"`
}
