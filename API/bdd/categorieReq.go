package bdd

import (
	"fmt"
	"upcycleconnect/models"
)

func GetCategories() ([]models.Categorie, error) {

	var Categories []models.Categorie

	rows, err := Db.Query("SELECT id, libelle FROM pa2026.categorie")

	if err != nil {
		return nil, fmt.Errorf("get Categories : %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		var Categorie models.Categorie

		err := rows.Scan(&Categorie.Id, &Categorie.Libelle)

		if err != nil {
			return nil, fmt.Errorf("get Categories : %v", err.Error())
		}
		Categories = append(Categories, Categorie)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("get Categories : %v", err.Error())
	}

	return Categories, nil
}

func CreateCategorie(Categorie models.Categorie) error {

var count int
    checkQuery := "SELECT COUNT(*) FROM categorie WHERE libelle = ?"
    err := Db.QueryRow(checkQuery, Categorie.Libelle).Scan(&count)
    
    if err != nil {
        return fmt.Errorf("Erreur vérification libellé : %s", err.Error())
    }

    if count > 0 {
		
        return fmt.Errorf("Le libellé %s est déjà utilisé", Categorie.Libelle)
    }



	_, err = Db.Exec("INSERT INTO pa2026.categorie (libelle) VALUES (?)", Categorie.Libelle)

	if err != nil {
		return fmt.Errorf("CreateCategorie : %s", err.Error())
	}
	return nil
}

func DeleteCategorie(id int) error {

	_, err := Db.Query("SELECT id FROM pa2026.categorie WHERE id = ?", id)
	if err != nil {
		fmt.Println("erreur",err)
		return fmt.Errorf("la catégorie n'existe pas : %d", id)

	}

	// sUPPRESSION
	_, err = Db.Exec(
		"DELETE FROM pa2026.categorie WHERE id = ?", id)
	if err != nil {
				fmt.Println("erreur",err)

		return fmt.Errorf("mise à jour échouée : %v", err)
	}

	return nil
}