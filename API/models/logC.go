package models

type LogConnexion struct {
	Id            int    `json:"id"`
	IdUser        int    `json:"id_user"`
	Ip            string `json:"ip"`
	DateConnexion string `json:"date_connexion"` // Ou time.Time si tu gères bien les dates en Go
}