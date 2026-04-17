package bdd

import (
	"database/sql"
	"fmt"
	"strings"
	"upcycleconnect/models"

	"golang.org/x/crypto/bcrypt" //gestion hash mdp
)

func LoginUser(email string, motDePasse string) (models.User, error) {
	var user models.User

	err := Db.QueryRow("SELECT id, mot_de_passe, role, validation FROM pa2026.utilisateur WHERE email = ?", email).Scan(&user.Id, &user.MotDePasse, &user.Role, &user.Validation)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("email ou mot de passe incorrect")
		}
		return user, fmt.Errorf("erreur BDD : %v", err)
	}

	fmt.Printf("Mot de passe reçu du JSON : '%s'\n", motDePasse)
	fmt.Printf("Hash BDD trouvé         : '%s'\n", user.MotDePasse)



	err = bcrypt.CompareHashAndPassword([]byte(user.MotDePasse), []byte(motDePasse))
	if err != nil {
		fmt.Println("Erreur Bcrypt :", err) 
		return user, fmt.Errorf("email ou mot de passe incorrect")
	}

	user.Email = email
	return user, nil
}
func GetUsers() ([]models.User, error) {

	var Users []models.User
	var cheminDoc sql.NullString

	rows, err := Db.Query("SELECT u.id, u.nom, u.prenom, u.email, u.role, u.score, u.validation, d.chemin_fichier FROM pa2026.utilisateur u LEFT JOIN pa2026.documents_legaux d ON u.id = d.user_id")

	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête : ", err)

		return nil, fmt.Errorf("get Users : %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		var User models.User

		err := rows.Scan(&User.Id,
			&User.Nom,
			&User.Prenom,
			&User.Email,
			&User.Role,
			&User.Score,
			&User.Validation,
			&cheminDoc)

		if err != nil {
			return nil, fmt.Errorf("get Users : %v", err.Error())
		}

		if cheminDoc.Valid {
			User.CheminFichier = cheminDoc.String
		} else {
			User.CheminFichier = ""
		}
		Users = append(Users, User)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("get Users : %v", err.Error())
	}

	return Users, nil
}

func CreateUser(User models.User) (int64, error) {

	var count int

	err := Db.QueryRow("SELECT COUNT(*) FROM pa2026.utilisateur WHERE email = ?", User.Email).Scan(&count)

	if err != nil {
		return 0, fmt.Errorf("Erreur vérification email : %s", err.Error())
	}

	if count > 0 {
		return 0, fmt.Errorf("L'email %s est déjà utilisé", User.Email)
	}

	result, err := Db.Exec("INSERT INTO pa2026.utilisateur (nom, prenom, email, mot_de_passe, role, validation, nom_entreprise, siret) VALUES (UPPER(?), UPPER(?), ?, ?, ?, 'En attente', ?, ?)", User.Nom, User.Prenom, User.Email, User.MotDePasse, User.Role, User.NomEntreprise, User.Siret)

	if err != nil {
		return 0, fmt.Errorf("CreateUser : %s", err.Error())
	}

	nouvelID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Erreur lors de la récupération de l'ID : %s", err.Error())
	}

	return nouvelID, nil
}
func DeletedUser(id int) error {


	_, err := Db.Query("SELECT id FROM pa2026.utilisateur WHERE id = ?", id)
	if err != nil {

		return fmt.Errorf("l'utilisateur n'existe pas : %d", id)

	}

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

func GetUserById(id int) (models.User, error) { // On retire les []
    var user models.User // Un seul user, pas un slice

    err := Db.QueryRow("SELECT id, nom, prenom, email, mot_de_passe, role, type_statut, nom_entreprise, siret, score, validation FROM pa2026.utilisateur WHERE id = ?", id).Scan(
        &user.Id, &user.Nom, &user.Prenom, &user.Email, &user.MotDePasse, 
        &user.Role, &user.TypeStatut, &user.NomEntreprise, &user.Siret, 
        &user.Score, &user.Validation,
    )

    if err != nil {
        return models.User{}, fmt.Errorf("get User by id : %v", err)
    }

    return user, nil
}

func GetUserByRole(role string) ([]models.User, error) {
	var Users []models.User

	rows, err := Db.Query("SELECT id, nom, prenom, email, mot_de_passe, role, type_statut, nom_entreprise, siret, score FROM pa2026.utilisateur WHERE role = ?", role)

	if err != nil {
		return nil, fmt.Errorf("get User by role : %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var User models.User

		err := rows.Scan(&User.Id, &User.Nom, &User.Prenom, &User.Email, &User.MotDePasse, &User.Role, &User.TypeStatut, &User.NomEntreprise, &User.Siret, &User.Score)

		if err != nil {
			return nil, fmt.Errorf("get User by role : %v", err.Error())
		}

		Users = append(Users, User)
	}

	err = rows.Err()

	if err != nil {
		return nil, fmt.Errorf("get User by role : %v", err.Error())
	}
	return Users, nil
}

func GetUserByName(query string, role string) ([]models.User, error) {
	query = strings.ToUpper(query)
	search := "%" + query + "%"
	if role != "Tous les rôles" && role != "" {
		var Users []models.User
		rows, err := Db.Query("SELECT id, nom, prenom, email, mot_de_passe, role, type_statut, nom_entreprise, siret, score FROM pa2026.utilisateur WHERE (UPPER(nom) LIKE ? OR UPPER(prenom) LIKE ?) AND role = ?", search, search, role)

		if err != nil {
			fmt.Println("Erreur lors de l'exécution de la requête : ", err)
			return nil, fmt.Errorf("get User by name : %v", err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var User models.User

			err := rows.Scan(&User.Id, &User.Nom, &User.Prenom, &User.Email, &User.MotDePasse, &User.Role, &User.TypeStatut, &User.NomEntreprise, &User.Siret, &User.Score)

			if err != nil {
				fmt.Println("Erreur lors de l'exécution de la requête : ", err)
				return nil, fmt.Errorf("get User by name : %v", err.Error())
			}

			Users = append(Users, User)
		}

		err = rows.Err()

		if err != nil {
			return nil, fmt.Errorf("get User by name : %v", err.Error())
		}
		return Users, nil
	} else {
		var Users []models.User
		search := "%" + query + "%"
		rows, err := Db.Query("SELECT id, nom, prenom, email, mot_de_passe, role, type_statut, nom_entreprise, siret, score FROM pa2026.utilisateur WHERE (UPPER(nom) LIKE ? OR UPPER(prenom) LIKE ?)", search, search)

		if err != nil {
			return nil, fmt.Errorf("get User by name : %v", err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var User models.User

			err := rows.Scan(&User.Id, &User.Nom, &User.Prenom, &User.Email, &User.MotDePasse, &User.Role, &User.TypeStatut, &User.NomEntreprise, &User.Siret, &User.Score)

			if err != nil {
				return nil, fmt.Errorf("get User by name : %v", err.Error())
			}

			Users = append(Users, User)
		}

		err = rows.Err()

		if err != nil {
			return nil, fmt.Errorf("get User by name : %v", err.Error())
		}
		return Users, nil
	}
}

func ValidateUser(id int) error {
	_, err := Db.Exec(
		"UPDATE pa2026.utilisateur SET validation = 'Validé' WHERE id = ?",
		id,
	)
	if err != nil {
		return fmt.Errorf("validate user : %v", err.Error())
	}
	return nil
}

func RefuseUser(id int, motif string) error {
	_, err := Db.Exec(
		"UPDATE pa2026.utilisateur SET validation = 'Rejeté', motif_refus = ? WHERE id = ?",
		motif,
		id,
	)
	if err != nil {
		fmt.Println("Erreur:", err)

		return fmt.Errorf("refuse user : %v", err.Error())
	}
	return nil
}

// Dans ton fichier bdd/documents.go (ou là où tu gères la BDD)
func InsertDocument(userID string, typeDocument string, cheminFichier string) error {
	// On insère le document avec le statut "En attente" par défaut
	requeteSQL := `
		INSERT INTO pa2026.documents_legaux (user_id, type_document, chemin_fichier, statut_document) 
		VALUES (?, ?, ?, 'En attente')
	`
	_, err := Db.Exec(requeteSQL, userID, typeDocument, cheminFichier)
	if err != nil {
		return fmt.Errorf("erreur lors de l'insertion du document : %s", err.Error())
	}
	return nil
}

func CheckEmailExists(email string) (bool, error) {
	var count int

	requete := "SELECT COUNT(*) FROM utilisateur WHERE email = ?"

	err := Db.QueryRow(requete, email).Scan(&count)

	if err != nil {
		return false, fmt.Errorf("erreur lors de la vérification de l'email : %s", err.Error())
	}

	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
