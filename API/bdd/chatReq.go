package bdd

import (
	"upcycleconnect/models"
)

func SaveMessage(msg models.Message) error {
	query := `INSERT INTO pa2026.message (annonce_id, expediteur_id, destinataire_id, contenu) 
              VALUES (?, ?, ?, ?)`
	_, err := Db.Exec(query, msg.AnnonceID, msg.ExpediteurID, msg.DestinataireID, msg.Contenu)
	return err
}

func GetConversation(annonceID, user1, user2 int) ([]models.Message, error) {
	var messages []models.Message
	query := `SELECT id, annonce_id, expediteur_id, destinataire_id, contenu, lu, date_envoi 
              FROM pa2026.message 
              WHERE annonce_id = ? 
              AND ((expediteur_id = ? AND destinataire_id = ?) OR (expediteur_id = ? AND destinataire_id = ?))
              ORDER BY date_envoi ASC`

	rows, err := Db.Query(query, annonceID, user1, user2, user2, user1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m models.Message
		err := rows.Scan(&m.ID, &m.AnnonceID, &m.ExpediteurID, &m.DestinataireID, &m.Contenu, &m.Lu, &m.DateEnvoi)
		if err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}
	return messages, nil
}

func GetUserConversations(userID int) ([]map[string]interface{}, error) {
	query := `
        SELECT DISTINCT 
            CASE WHEN expediteur_id = ? THEN destinataire_id ELSE expediteur_id END as contact_id,
            u.nom, u.prenom, a.titre as annonce_titre, a.id as annonce_id
        FROM pa2026.message m
        JOIN pa2026.utilisateur u ON u.id = (CASE WHEN m.expediteur_id = ? THEN m.destinataire_id ELSE m.expediteur_id END)
        JOIN pa2026.annonce a ON a.id = m.annonce_id
        WHERE m.expediteur_id = ? OR m.destinataire_id = ?
    `
	rows, err := Db.Query(query, userID, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var contactID, annonceID int
		var nom, prenom, titre string
		rows.Scan(&contactID, &nom, &prenom, &titre, &annonceID)

		results = append(results, map[string]interface{}{
			"contact_id":    contactID,
			"nom":           nom,
			"prenom":        prenom,
			"annonce_id":    annonceID,
			"annonce_titre": titre,
		})
	}
	return results, nil
}
