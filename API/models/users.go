package models

type User struct {
	Id                   int     `json:"id"`
	Role                 string  `json:"role"`
	Nom                  string  `json:"nom"`
	Prenom               string  `json:"prenom"`
	Email                string  `json:"email"`
	Score                int     `json:"score"`
	MotDePasse           string  `json:"mot_de_passe,omitempty"`
	TutorielVu           bool    `json:"tutoriel_vu"`
	Validation           string  `json:"validation"`
	TypeStatut           *string `json:"type_statut,omitempty"`
	NomEntreprise        *string `json:"nom_entreprise,omitempty"`
	Siret                *string `json:"siret,omitempty"`
	CheminFichier        string  `json:"chemin_fichier"`
	StripeAccountId      *string `json:"stripe_account_id"`
	StripeVerifCompleted bool    `json:"stripe_verif_completed"`
}

type UpdatePasswordInput struct {
	ID          int    `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type RefuseRequest struct {
	Motif string `json:"motif"`
}
type InscriptionRequest struct {
	IdUser  int `json:"id_user"`
	IdEvent int `json:"id_event"`
}
