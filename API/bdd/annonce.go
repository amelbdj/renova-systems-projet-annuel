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

// func CreateAnnonce(Annonce models.Annonce) error {

// var count int
//     checkQuery := "SELECT COUNT(*) FROM Annonce WHERE libelle = ?"
//     err := Db.QueryRow(checkQuery, Annonce.Libelle).Scan(&count)
    
//     if err != nil {
//         return fmt.Errorf("Erreur vérification libellé : %s", err.Error())
//     }

//     if count > 0 {
		
//         return fmt.Errorf("Le libellé %s est déjà utilisé", Annonce.Libelle)
//     }



// 	_, err = Db.Exec("INSERT INTO pa2026.Annonce (libelle) VALUES (?)", Annonce.Libelle)

// 	if err != nil {
// 		return fmt.Errorf("CreateAnnonce : %s", err.Error())
// 	}
// 	return nil
// }