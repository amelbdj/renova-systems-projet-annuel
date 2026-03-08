package bdd

import (
	"fmt"
	"upcycleconnect/models"
)

func GetUsers() ([]models.User, error) {

	var Users []models.User

	rows, err := Db.Query("SELECT id, nom, prenom, email, mot_de_passe, role, type_statut, nom_entreprise, siret FROM pa2026.utilisateur")

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

func DeletedUser(id int) error {

	// Vérifie si ID EXISTE

	_, err := Db.Query("SELECT id FROM pa2026.utilisateur WHERE id = ?", id)
	if err != nil {

		return fmt.Errorf("l'utilisateur n'existe pas : %d", id)

	}

	// sUPPRESSION
	_, err = Db.Exec(
		"DELETE FROM pa2026.utilisateur WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("mise à jour échouée : %v", err)
	}

	return nil
}

func UpdateUserById(user models.User) error {
	
	result, err := Db.Exec(
		"UPDATE pa2026.utilisateur SET nom = ?, prenom = ?, email = ?, mot_de_passe = ?, role = ?, type_statut = ?, nom_entreprise = ?, siret = ? WHERE id = ?",
		user.Nom,
		user.Prenom,
		user.Email,
		user.MotDePasse,
		user.Role,
		user.TypeStatut,
		user.NomEntreprise,
		user.Siret,
		user.Id,
	)
	if err != nil {
		return fmt.Errorf("mise à jour échouée : %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("aucun utilisateur trouvé avec l'id %d", user.Id)
	}

	return nil
}

func GetUserById(id int) ([]models.User, error) {
	var Users []models.User

	rows, err := Db.Query("SELECT id, nom, prenom, email, mot_de_passe, role, type_statut, nom_entreprise, siret FROM partiel.User WHERE id = ?", id)

	if err != nil {
		return nil, fmt.Errorf("get User by id : %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var User models.User

		err := rows.Scan(&User.Id, &User.Nom, &User.Prenom, &User.Email, &User.MotDePasse, &User.Role, &User.TypeStatut, &User.NomEntreprise, &User.Siret)

		if err != nil {
			return nil, fmt.Errorf("get User by name : %v", err.Error())
		}

		Users = append(Users, User)
	}

	err = rows.Err()

	if err != nil {
		return nil, fmt.Errorf("get User by id : %v", err.Error())
	}
	return Users, nil
}

