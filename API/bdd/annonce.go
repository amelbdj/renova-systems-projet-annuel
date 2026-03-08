package bdd

import (
	"fmt"
	"upcycleconnect/models"
)

func GetAnnonces() ([]models.Annonce, error) {

	var Annonces []models.Annonce

	rows, err := Db.Query("SELECT annonce.id, annonce.titre, annonce.description, utilisateur.nom, utilisateur.prenom, categorie.libelle FROM pa2026.annonce INNER JOIN utilisateur ON utilisateur.id = annonce.id_user INNER JOIN categorie ON categorie.id = annonce.id_categorie")

	if err != nil {
		return nil, fmt.Errorf("get Annonces : %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		var Annonce models.Annonce
	
		err := rows.Scan(&Annonce.Id, &Annonce.Titre, &Annonce.Description, &Annonce.Nom, &Annonce.Prenom, &Annonce.Categorie)

		if err != nil {
			return nil, fmt.Errorf("get Annonces : %v", err.Error())
		}
		Annonces = append(Annonces, Annonce)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("get Annonces : %v", err.Error())
	}

	return Annonces, nil
}

func ValidateAnnonce(annonceId int) error {

	result, err := Db.Exec("UPDATE pa2026.annonce SET statut_validation = 'valide' WHERE id = ?", annonceId)

if err != nil {
		return fmt.Errorf("mise à jour échouée : %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("aucune annonce trouvée avec l'id %d", annonceId)
	}

	return nil
}

func RefuseAnnonce(annonceId int) error {
		result, err := Db.Exec("UPDATE pa2026.annonce SET statut_validation = 'refuse' WHERE id = ?", annonceId)

	if err != nil {
		return fmt.Errorf("mise à jour échouée : %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("aucune annonce trouvée avec l'id %d", annonceId)
	}

	return nil
}