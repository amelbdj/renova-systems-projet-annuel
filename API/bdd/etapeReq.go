package bdd

import "upcycleconnect/models"

func CreateEtape(e models.Etape) error {
	_, err := Db.Exec(
		`INSERT INTO etapes_projet (id_projet, titre, description, statut) VALUES (?, ?, ?, ?)`,
		e.IdProjet, e.Titre, e.Description, e.Statut,
	)
	return err
}

func GetEtapesByProjet(idProjet int) ([]models.Etape, error) {
	rows, err := Db.Query(
		`SELECT id_etape, id_projet, titre, description, statut, created_at FROM etapes_projet WHERE id_projet = ? ORDER BY created_at ASC`,
		idProjet,
	)
	if err != nil { return nil, err }
	defer rows.Close()

	var etapes []models.Etape
	for rows.Next() {
		var e models.Etape
		if err := rows.Scan(&e.Id, &e.IdProjet, &e.Titre, &e.Description, &e.Statut, &e.CreatedAt); err != nil {
			return nil, err
		}
		etapes = append(etapes, e)
	}
	return etapes, nil
}

func DeleteEtape(id int) error {
	_, err := Db.Exec(`DELETE FROM etapes_projet WHERE id_etape = ?`, id)
	return err
}

func UpdateEtapeStatut(id int, statut string) error {
	_, err := Db.Exec(`UPDATE etapes_projet SET statut = ? WHERE id_etape = ?`, statut, id)
	return err
}