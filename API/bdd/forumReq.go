package bdd

import (
	"fmt"
	"upcycleconnect/models"
)

func GetForumMessages(filtre string) ([]models.MessageForum, error) {
	var messages []models.MessageForum

	requete := `
		SELECT m.id_message, m.id_topic, m.id_user, m.contenu, m.est_modere, m.est_signale, m.date_creation, 
		       u.nom, u.prenom, t.titre 
		FROM pa2026.message_forum m
		INNER JOIN pa2026.utilisateur u ON m.id_user = u.id
		INNER JOIN pa2026.topic_forum t ON m.id_topic = t.id_topic
	`

	if filtre == "signales" {
		requete += " WHERE m.est_signale = 1 AND m.est_modere = 0"
	} else {
		requete += " ORDER BY m.date_creation DESC"
	}

	rows, err := Db.Query(requete)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération messages forum : %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var msg models.MessageForum
		err := rows.Scan(
			&msg.IdMessage, &msg.IdTopic, &msg.IdUser, &msg.Contenu, 
			&msg.EstModere, &msg.EstSignale, &msg.DateCreation, 
			&msg.NomAuteur, &msg.PrenomAuteur, &msg.TitreTopic,
		)
		if err != nil {
			return nil, fmt.Errorf("erreur scan message forum : %v", err)
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

func ModerateForumMessage(idMessage int, action string) error {
	var requete string

	if action == "approuver" {
		requete = "UPDATE pa2026.message_forum SET est_modere = 1, est_signale = 0 WHERE id_message = ?"
	} else if action == "masquer" {
		requete = "UPDATE pa2026.message_forum SET est_modere = 1 WHERE id_message = ?"
	} else {
		return fmt.Errorf("action de modération inconnue")
	}

	_, err := Db.Exec(requete, idMessage)
	return err
}

func GetForumStats() (models.StatistiqueForum, error) {
	var stats models.StatistiqueForum

	err := Db.QueryRow("SELECT COUNT(*) FROM pa2026.message_forum WHERE date_creation >= NOW() - INTERVAL 7 DAY").Scan(&stats.MessagesSemaine)
	if err != nil {
		return stats, err
	}

	err = Db.QueryRow("SELECT COUNT(*) FROM pa2026.message_forum WHERE est_signale = 1 AND est_modere = 0").Scan(&stats.Signalements)
	if err != nil {
		return stats, err
	}

	err = Db.QueryRow("SELECT COUNT(DISTINCT id_user) FROM pa2026.message_forum WHERE date_creation >= NOW() - INTERVAL 7 DAY").Scan(&stats.MembresActifs) // distinc pour EVITER LES DOUBLONS
	if err != nil {
		return stats, err
	}

	return stats, nil
}