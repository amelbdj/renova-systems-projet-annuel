package models

type DepositRequest struct {
	AnnonceId   int64 `json:"annonce_id"`
	ConteneurId int64 `json:"conteneur_id"`
}

type Conteneur struct {
	Id           int    `json:"id"`
	Nom          string `json:"nom"`
	Adresse      string `json:"adresse"`
	DateCreation string `json:"date_creation"` // string est plus simple à gérer en JS que time.Time
}

// Structure pour la petite porte (le casier)
type Box struct {
	Id          int    `json:"id"`
	IdConteneur int    `json:"id_conteneur"`
	Numero      int    `json:"numero"`
	Taille      string `json:"taille"`
	Statut      string `json:"statut"`
	// Le pointeur *string est une astuce magique en Go pour gérer le fait
	// que le code secret peut être NULL (vide) dans ta base de données !
	CodeSecret *string `json:"code_secret"`
}

type ConteneurAvecStats struct {
	Id         int    `json:"id"`
	Nom        string `json:"nom"`
	Adresse    string `json:"adresse"`
	TotalBoxes int    `json:"total_boxes"` // Ce champ n'existe pas en table, il vient du COUNT(b.id) !
}