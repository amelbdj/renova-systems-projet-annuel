package bdd

import (
	"fmt"
	"upcycleconnect/models"
)

func GetUsers() ([]models.User, error) {

	var Users []models.User

	rows, err := Db.Query("SELECT id_user, nom, prenom, email, mot_de_passe, role, type_statut, nom_entreprise, siret FROM pa2026.utilisateur")

	if err != nil {
		return nil, fmt.Errorf("get Users : %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		var User models.User

		err := rows.Scan(&User.Id, &User.Nom, &User.Prenom, &User.Email, &User.MotDePasse, &User.Role, &User.TypeStatut, &User.NomEntreprise, &User.Siret)

		if err != nil {
			return nil, fmt.Errorf("get Users : %v", err.Error())
		}
		Users = append(Users, User)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("get Users : %v", err.Error())
	}

	return Users, nil
}

func CreateUser(User models.User) error {

var count int
    checkQuery := "SELECT COUNT(*) FROM utilisateur WHERE email = ?"
    err := Db.QueryRow(checkQuery, User.Email).Scan(&count)
    
    if err != nil {
        return fmt.Errorf("Erreur vérification email : %s", err.Error())
    }

    if count > 0 {
		
        return fmt.Errorf("L'email %s est déjà utilisé", User.Email)
    }



	_, err = Db.Exec("INSERT INTO pa2026.utilisateur (nom, prenom, email, mot_de_passe, role, type_statut, nom_entreprise, siret) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", User.Nom, User.Prenom, User.Email, User.MotDePasse, User.Role, User.TypeStatut, User.NomEntreprise, User.Siret)

	if err != nil {
		return fmt.Errorf("CreateUser : %s", err.Error())
	}
	return nil
}