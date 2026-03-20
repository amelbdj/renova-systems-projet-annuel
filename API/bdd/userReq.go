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
    
    err := Db.QueryRow("SELECT id, mot_de_passe, role FROM pa2026.utilisateur WHERE email = ?", email).Scan(&user.Id, &user.MotDePasse, &user.Role)

    if err != nil {
        if err == sql.ErrNoRows {
            return user, fmt.Errorf("email ou mot de passe incorrect")
        }
        return user, fmt.Errorf("erreur BDD : %v", err)
    }

    // 🕵️ LES 3 LIGNES D'ESPIONNAGE :
    fmt.Println("--- ESPIONNAGE LOGIN ---")
    fmt.Printf("Mot de passe reçu du JSON : '%s'\n", motDePasse)
    fmt.Printf("Hash BDD trouvé         : '%s'\n", user.MotDePasse)

    err = bcrypt.CompareHashAndPassword([]byte(user.MotDePasse), []byte(motDePasse))
    if err != nil {
        fmt.Println("Erreur Bcrypt :", err) // Ça nous dira exactement pourquoi Bcrypt bloque !
        return user, fmt.Errorf("email ou mot de passe incorrect")
    }
    
    user.Email = email
    return user, nil
}
func GetUsers() ([]models.User, error) {

	var Users []models.User

	rows, err := Db.Query("SELECT id, nom, prenom, email, mot_de_passe, role, type_statut, nom_entreprise, siret, score FROM pa2026.utilisateur")

	if err != nil {
				fmt.Println("Erreur lors de l'exécution de la requête : ", err)

		return nil, fmt.Errorf("get Users : %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		var User models.User

		err := rows.Scan(&User.Id, &User.Nom, &User.Prenom, &User.Email, &User.MotDePasse, &User.Role, &User.TypeStatut, &User.NomEntreprise, &User.Siret, &User.Score)

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
    
    err := Db.QueryRow( "SELECT COUNT(*) FROM utilisateur WHERE email = ?", User.Email).Scan(&count)
    
    if err != nil {
        return fmt.Errorf("Erreur vérification email : %s", err.Error())
    }

    if count > 0 {
		
        return fmt.Errorf("L'email %s est déjà utilisé", User.Email)
    }



	_, err = Db.Exec("INSERT INTO pa2026.utilisateur (nom, prenom, email, mot_de_passe, role, type_statut, nom_entreprise, siret) VALUES (UPPER(?), UPPER(?), ?, ?, ?, ?, ?, ?)", User.Nom, User.Prenom, User.Email, User.MotDePasse, User.Role, User.TypeStatut, User.NomEntreprise, User.Siret)

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

	rows, err := Db.Query("SELECT id, nom, prenom, email, mot_de_passe, role, type_statut, nom_entreprise, siret, score FROM pa2026.utilisateur WHERE id = ?", id)

	if err != nil {
		return nil, fmt.Errorf("get User by id : %v", err.Error())
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
		return nil, fmt.Errorf("get User by id : %v", err.Error())
	}
	return Users, nil
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
	search := "%"+query+"%"
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
	search := "%"+query+"%"
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