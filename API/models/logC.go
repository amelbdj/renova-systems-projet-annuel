package models

type LogConnexion struct {
	Id            int    `json:"id"`
	IdUser        int    `json:"id_user"`
	Ip            string `json:"ip"`
	DateConnexion string `json:"date_connexion"`
}
