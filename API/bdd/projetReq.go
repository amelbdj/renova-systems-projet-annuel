package bdd

import "upcycleconnect/models"

func CreateProjet(p models.Projet) error {
	_, err := Db.Exec(
		`INSERT INTO projet_pro (id_user, titre, desc_etapes, url_photo_avant, url_photo_apres, statut, co2_evite) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		p.IdUser, p.Titre, p.Description, p.PhotoAvant, p.PhotoApres, p.Statut, p.Co2Evite,
	)
	return err
}

func GetProjetsByUser(idUser int) ([]models.Projet, error) {
	rows, err := Db.Query(
		`SELECT id_projet, id_user, titre, desc_etapes, COALESCE(url_photo_avant, ''), COALESCE(url_photo_apres, ''), COALESCE(statut, 'en_cours'), COALESCE(co2_evite, 0) FROM projet_pro WHERE id_user = ?`,
		idUser,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projets []models.Projet
	for rows.Next() {
		var p models.Projet
		if err := rows.Scan(&p.Id, &p.IdUser, &p.Titre, &p.Description, &p.PhotoAvant, &p.PhotoApres, &p.Statut, &p.Co2Evite); err != nil {
			return nil, err
		}
		projets = append(projets, p)
	}
	return projets, nil
}

func DeleteProjet(id int) error {
	Db.Exec(`DELETE FROM etapes_projet WHERE id_projet = ?`, id)
	_, err := Db.Exec(`DELETE FROM projet_pro WHERE id_projet = ?`, id)
	return err
}

func UpdateProjet(p models.Projet) error {
	_, err := Db.Exec(
		`UPDATE projet_pro SET titre = ?, desc_etapes = ?, url_photo_avant = ?, url_photo_apres = ?, statut = ?, co2_evite = ? WHERE id_projet = ?`,
		p.Titre, p.Description, p.PhotoAvant, p.PhotoApres, p.Statut, p.Co2Evite, p.Id,
	)
	return err
}
