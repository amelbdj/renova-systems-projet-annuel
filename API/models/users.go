package models

type User struct {
	Id            int     `json:"id"`
	Role          string  `json:"role"`
	Nom           string  `json:"nom"`
	Prenom        string  `json:"prenom"`
	Email         string  `json:"email"`
	Score         int     `json:"score"`
	MotDePasse    string  `json:"mot_de_passe,omitempty"`
	TutorielVu    bool    `json:"tutoriel_vu"`
	Validation    string  `json:"validation"`
	TypeStatut    *string `json:"type_statut,omitempty"`
	NomEntreprise *string `json:"nom_entreprise,omitempty"`
	Siret         *string `json:"siret,omitempty"`
}