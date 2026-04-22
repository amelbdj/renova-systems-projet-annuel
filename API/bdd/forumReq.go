package bdd

import (
	"fmt"
	"upcycleconnect/models"
)

func GetForumMessages(filtre string) ([]models.MessageForum, error) {
	var messages []models.MessageForum

	
	requete := `
		SELECT m.id_message, m.id_topic, m.id_user, m.contenu, m.est_modere, m.est_signale, m.date_creation, 
			   COALESCE(u.nom, 'Anonyme'), COALESCE(u.prenom, 'Utilisateur'), COALESCE(t.titre, 'Topic inconnu')
		FROM pa2026.message_forum m
		LEFT JOIN pa2026.utilisateur u ON m.id_user = u.id
		LEFT JOIN pa2026.topic_forum t ON m.id_topic = t.id_topic
		WHERE m.est_modere = 0
	`

	if filtre == "signales" {
		requete += " AND m.est_signale = 1"
	}

	requete += " ORDER BY m.date_creation DESC"

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

func GetAllTopics() ([]models.ForumTopic, error) {
	topics := []models.ForumTopic{}

	lignes, err := Db.Query("SELECT t.id_topic, t.titre, DATE_FORMAT(t.date_creation, '%d-%m-%Y à %H:%i') as date_creation, u.prenom as auteur, (SELECT COUNT(*) FROM message_forum m WHERE m.id_topic = t.id_topic) as nb_reponses FROM topic_forum t LEFT JOIN utilisateur u ON t.id_user = u.id ORDER BY t.date_creation DESC")
	if err != nil {
		fmt.Println("ERREUR SQL Forum :", err)
		return topics, err
	}
	defer lignes.Close()

	for lignes.Next() {
		var t models.ForumTopic
		lignes.Scan(&t.Id, &t.Titre, &t.Date, &t.Auteur, &t.NbReponses)
		topics = append(topics, t)
	}

	return topics, nil
}

func GetMessagesByTopicClient(topicId int) ([]models.MessageForum, error) {
    var messages []models.MessageForum

    rows, err := Db.Query(`
        SELECT m.id_message, m.id_topic, m.id_user, m.contenu, m.est_modere, m.est_signale, 
               DATE_FORMAT(m.date_creation, '%d-%m-%Y à %H:%i'), 
               COALESCE(u.nom, 'Anonyme'), COALESCE(u.prenom, 'Utilisateur'), COALESCE(t.titre, 'Topic inconnu')
        FROM pa2026.message_forum m
        LEFT JOIN pa2026.utilisateur u ON m.id_user = u.id
        LEFT JOIN pa2026.topic_forum t ON m.id_topic = t.id_topic
        WHERE m.id_topic = ? AND m.est_modere = 0
        ORDER BY m.date_creation ASC
    `, topicId)

	
    if err != nil {
        return nil, fmt.Errorf("erreur récupération messages du topic : %v", err)
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
            return nil, fmt.Errorf("erreur scan message topic : %v", err)
        }
        messages = append(messages, msg)
    }

    return messages, nil
}

func AjouterMessageForum(topicId int, userId int, contenu string) error {

	_, err := Db.Exec("INSERT INTO pa2026.message_forum (id_topic, id_user, contenu, est_modere, est_signale) VALUES (?, ?, ?, 0, 0)", topicId, userId, contenu)
	
	if err != nil {
		return fmt.Errorf("erreur lors de l'insertion du message : %v", err)
	}

	return nil
}

func CreerNouveauSujet(userId int, titre string, premierMessage string) error {
	resultat, err := Db.Exec("INSERT INTO pa2026.topic_forum (id_user, titre) VALUES (?, ?)", userId, titre)
	if err != nil {
		return fmt.Errorf("erreur création topic : %v", err)
	}

	topicId, err := resultat.LastInsertId()
	if err != nil {
		return fmt.Errorf("erreur récupération ID topic : %v", err)
	}

	_, err = Db.Exec("INSERT INTO pa2026.message_forum (id_topic, id_user, contenu, est_modere, est_signale) VALUES (?, ?, ?, 0, 0)", topicId, userId, premierMessage)
	if err != nil {
		return fmt.Errorf("erreur insertion premier message : %v", err)
	}

	return nil
}