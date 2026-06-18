package models

type Evenement struct {
	Id               int     `json:"id"`
	Titre            string  `json:"titre"`
	Description      string  `json:"description"`
	DateDebut        string  `json:"date_debut"`
	DateFin          string  `json:"date_fin"`
	NbPlaces         int     `json:"nb_places"`
	StatutValidation string  `json:"statut_validation"`
	Format           string  `json:"format"`
	NomSalarie       string  `json:"nomSalarie"`
	PrenomSalarie    string  `json:"prenomSalarie"`
	Lieu             string  `json:"lieu"`
	Type             string  `json:"type"`
	IdSalarie        int     `json:"idSalarie"`
	Prix             float64 `json:"prix"`
	DejaInscrit      bool    `json:"deja_inscrit"`
	ImageUrl         string  `json:"image_url"`
	PdfUrl           string  `json:"pdf_url"`
}
