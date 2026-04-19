package models

type MessageForum struct {
	IdMessage    int    `json:"id_message"`
	IdTopic      int    `json:"id_topic"`
	IdUser       int    `json:"id_user"`
	Contenu      string `json:"contenu"`
	EstModere    bool   `json:"est_modere"`
	EstSignale   bool   `json:"est_signale"`
	DateCreation string `json:"date_creation"` // Formaté en string pour le JSON

	TitreTopic   string `json:"titre_topic"`
	NomAuteur    string `json:"nom_auteur"`
	PrenomAuteur string `json:"prenom_auteur"`
}

type StatistiqueForum struct {
	MessagesSemaine int `json:"messages_semaine"`
	MembresActifs   int `json:"membres_actifs"`
	Signalements    int `json:"signalements_en_attente"`
	ModerationsMois int `json:"moderations_mois"`
}
