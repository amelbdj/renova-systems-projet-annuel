package models

type Etape struct {
	Id          int    `json:"id"`
	IdProjet    int    `json:"id_projet"`
	Titre       string `json:"titre"`
	Description string `json:"description"`
	Statut      string `json:"statut"`
	CreatedAt   string `json:"created_at"`
}
