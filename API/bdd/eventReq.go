package bdd

import (
	"fmt"
	"upcycleconnect/models"
)

func GetEvenements() ([]models.Evenement, error) {

	var Evenements []models.Evenement

	rows, err := Db.Query("SELECT evenement.id, evenement.titre, evenement.description, evenement.date_debut, evenement.date_fin, evenement.nb_places, evenement.statut_validation, evenement.format, evenement.lieu, evenement.type, evenement.id_salarie, utilisateur.nom, utilisateur.prenom FROM pa2026.Evenement INNER JOIN utilisateur ON utilisateur.id = Evenement.id_salarie")

	if err != nil {
		return nil, fmt.Errorf("get Evenements : %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		var Evenement models.Evenement
	
		err := rows.Scan(&Evenement.Id, &Evenement.Titre, &Evenement.Description,&Evenement.DateDebut, &Evenement.DateFin, &Evenement.NbPlaces, &Evenement.StatutValidation, &Evenement.Format, &Evenement.Lieu, &Evenement.Type, &Evenement.IdSalarie, &Evenement.NomSalarie, &Evenement.PrenomSalarie)

		if err != nil {
			return nil, fmt.Errorf("get Evenements : %v", err.Error())
		}
		Evenements = append(Evenements, Evenement)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("get Evenements : %v", err.Error())
	}

	return Evenements, nil
}

func ValidateEvenement(EvenementId int) error {

	result, err := Db.Exec("UPDATE pa2026.Evenement SET statut_validation = 'valide' WHERE id = ?", EvenementId)

if err != nil {
		return fmt.Errorf("mise à jour échouée : %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("aucune Evenement trouvée avec l'id %d", EvenementId)
	}

	return nil
}

func RefuseEvenement(EvenementId int) error {
		result, err := Db.Exec("UPDATE pa2026.Evenement SET statut_validation = 'refuse' WHERE id = ?", EvenementId)

	if err != nil {
		return fmt.Errorf("mise à jour échouée : %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("aucune Evenement trouvée avec l'id %d", EvenementId)
	}

	return nil
}

func CreateEvenement(Evenement models.Evenement) error {

	_, err := Db.Exec("INSERT INTO pa2026.Evenement (titre, description, date_debut, date_fin, nb_places, statut_validation, format, lieu, type, id_salarie) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		Evenement.Titre, Evenement.Description, Evenement.DateDebut, Evenement.DateFin, Evenement.NbPlaces, "en attente", Evenement.Format, Evenement.Lieu, Evenement.Type, Evenement.IdSalarie)
	if err != nil {
		return fmt.Errorf("création de l'événement échouée : %v", err)
	}
	return nil
}

func DeleteEvenement(EvenementId int) error {
	result, err := Db.Exec("DELETE FROM pa2026.Evenement WHERE id = ?", EvenementId)

	if err != nil {
		return fmt.Errorf("suppression échouée : %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("aucune Evenement trouvée avec l'id %d", EvenementId)
	}
	return nil
}


func UpdateEvenement(EvenementId int, Evenement models.Evenement) error {

	result, err := Db.Exec("UPDATE pa2026.Evenement SET titre = ?, description = ?, date_debut = ?, date_fin = ?, nb_places = ?, format = ?, lieu = ?, type = ? WHERE id = ?",
		Evenement.Titre, Evenement.Description, Evenement.DateDebut, Evenement.DateFin, Evenement.NbPlaces, Evenement.Format, Evenement.Lieu, Evenement.Type, EvenementId)
	if err != nil {
		return fmt.Errorf("mise à jour échouée : %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("aucune Evenement trouvée avec l'id %d", EvenementId)
	}
	return nil
}