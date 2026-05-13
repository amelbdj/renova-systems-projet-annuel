package models

import "time"

type Message struct {
	ID             int       `json:"id"`
	AnnonceID      int       `json:"annonce_id"`
	ExpediteurID   int       `json:"expediteur_id"`
	DestinataireID int       `json:"destinataire_id"`
	Contenu        string    `json:"contenu"`
	Lu             bool      `json:"lu"`
	DateEnvoi      time.Time `json:"date_envoi"`
}