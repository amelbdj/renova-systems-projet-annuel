package bdd

import (
	"fmt"
	"upcycleconnect/models"
)

func GetForumMessages(filtre string) ([]models.MessageForum, error) {
	var messages []models.MessageForum

	// Utilisation de LEFT JOIN et COALESCE pour garantir que le message s'affiche TOUJOURS,
	// même si l'utilisateur ou le topic a été supprimé de la base de données.
	requete := `
		SELECT m.id_message, m.id_topic, m.id_user, m.contenu, m.est_modere, m.est_signale, m.date_creation, 
			   COALESCE(u.nom, 'Anonyme'), COALESCE(u.prenom, 'Utilisateur'), COALESCE(t.titre, 'Topic inconnu')
		FROM pa2026.message_forum m
		LEFT JOIN pa2026.utilisateur u ON m.id_user = u.id
		LEFT JOIN pa2026.topic_forum t ON m.id_topic = t.id_topic
		WHERE m.est_modere = 0
	`

	// Si le salarié clique sur le bouton "Signalés", on filtre pour ne garder que ceux-là
	if filtre == "signales" {
		requete += " AND m.est_signale = 1"
	}

	// On trie toujours du plus récent au plus ancien
	requete += " ORDER BY m.date_creation DESC"

	rows, err := Db.Query(requete)
	if err != nil {
		return nil, fmt.Errorf("erreur récupération messages forum : %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var msg models.MessageForum
		// Le Scan lira 'Utilisateur Anonyme' si la personne n'existe plus en base
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

	switch action {
	case "approuver":
		requete = "UPDATE pa2026.message_forum SET est_modere = 1, est_signale = 0 WHERE id_message = ?"
	case "masquer":
		requete = "UPDATE pa2026.message_forum SET est_modere = 1 WHERE id_message = ?"
	default:
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