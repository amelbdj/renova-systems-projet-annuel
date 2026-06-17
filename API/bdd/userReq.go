package bdd

import (
	"database/sql"
	"fmt"
	"net/smtp"
	"strconv"
	"strings"
	"upcycleconnect/models"

	"golang.org/x/crypto/bcrypt" //gestion hash mdp
)

func LoginUser(email string, motDePasse string, ip string) (models.User, error) {
	var user models.User

	err := Db.QueryRow("SELECT id, mot_de_passe, role, type_statut, prenom, score, tutoriel_vu, validation FROM pa2026.utilisateur WHERE email = ?", email).Scan(
		&user.Id,
		&user.MotDePasse,
		&user.Role,
		&user.TypeStatut,
		&user.Prenom,
		&user.Score,
		&user.TutorielVu,
		&user.Validation,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("email ou mot de passe incorrect")
		}
		return user, fmt.Errorf("erreur BDD : %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.MotDePasse), []byte(motDePasse))
	if err != nil {
		return user, fmt.Errorf("email ou mot de passe incorrect")
	}

	user.Email = email

	errLog := LogConnexion(user.Id, ip)
	if errLog != nil {
		fmt.Println("Erreur lors de l'enregistrement du log de connexion :", errLog)
	}

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

func GetUserById(id int) (models.User, error) {
	var user models.User

	err := Db.QueryRow("SELECT id, nom, prenom, email, mot_de_passe, role, type_statut, nom_entreprise, siret, score, validation, stripe_account_id, stripe_verif_completed, est_premium, stripe_customer_id FROM pa2026.utilisateur WHERE id = ?", id).Scan(
		&user.Id, &user.Nom, &user.Prenom, &user.Email, &user.MotDePasse,
		&user.Role, &user.TypeStatut, &user.NomEntreprise, &user.Siret,
		&user.Score, &user.Validation,
		&user.StripeAccountId,
		&user.StripeVerifCompleted,
		&user.EstPremium,
		&user.StripeCustomerId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("utilisateur non trouvé")
		}
		return models.User{}, fmt.Errorf("get User by id error: %v", err)
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

func ValidateUser(id int) (string, string, error) { // Retourne prenom, email, erreur
	_, err := Db.Exec(
		"UPDATE pa2026.utilisateur SET validation = 'Validé' WHERE id = ?",
		id,
	)
	if err != nil {
		return "", "", fmt.Errorf("validate user : %v", err.Error())
	}

	var prenom, email string
	err = Db.QueryRow("SELECT prenom, email FROM pa2026.utilisateur WHERE id = ?", id).Scan(&prenom, &email)
	if err != nil {
		return "", "", fmt.Errorf("erreur récupération infos pour email : %v", err.Error())
	}

	return prenom, email, nil
}

func RefuseUser(id int, motif string) (string, string, error) { // On ajoute les retours string, string
	_, err := Db.Exec(
		"UPDATE pa2026.utilisateur SET validation = 'Rejeté', motif_refus = ? WHERE id = ?",
		motif,
		id,
	)
	if err != nil {
		fmt.Println("Erreur:", err)
		return "", "", fmt.Errorf("refuse user : %v", err.Error())
	}

	// 2. On récupère ses infos pour lui envoyer l'e-mail
	var prenom, email string
	err = Db.QueryRow("SELECT prenom, email FROM pa2026.utilisateur WHERE id = ?", id).Scan(&prenom, &email)
	if err != nil {
		return "", "", fmt.Errorf("erreur récupération infos pour email refus : %v", err.Error())
	}

	return prenom, email, nil
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

func LogConnexion(idUser int, ip string) error {

	_, err := Db.Exec("INSERT INTO pa2026.log_connexion (id_user, ip, date_connexion) VALUES (?, ?, NOW())", idUser, ip)

	if err != nil {
		return err
	}

	return nil
}

func BanUser(userId int) error {
	_, err := Db.Exec("UPDATE pa2026.utilisateur SET validation = 'Banni' WHERE id = ?", userId)
	if err != nil {
		return fmt.Errorf("erreur lors du bannissement de l'utilisateur : %v", err)
	}

	_, err = Db.Exec("DELETE FROM pa2026.message_forum WHERE id_user = ?", userId)
	if err != nil {
		return fmt.Errorf("erreur lors de la suppression des messages : %v", err)
	}

	return nil
}

func GetAllAdminIDs() ([]string, error) {
	rows, err := Db.Query("SELECT id FROM pa2026.utilisateur WHERE role = 'Administrateur'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var adminIDs []string
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err == nil {
			// On convertit direct en string pour OneSignal
			adminIDs = append(adminIDs, strconv.Itoa(id))
		}
	}
	return adminIDs, nil
}

func EnvoyerEmailValidation(emailDestinataire string, prenom string) error {
	expediteur := "noreply@upcycleconnect.fr"
	motDePasse := "voir avec ndoya"
	serveurSMTP := "192.168.80.10"
	port := "25" // verif avec ndoya

	auth := smtp.PlainAuth("", expediteur, motDePasse, serveurSMTP)

	sujet := "Subject: UpcycleConnect - Votre compte est validé ! 🎉\n"
	typeMIME := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	corpsMessage := fmt.Sprintf(`
		<html>
			<body style="font-family: Arial, sans-serif; color: #333;">
				<h2>Bonjour %s,</h2>
				<p>Nous avons le plaisir de vous informer que votre compte a été validé par l'administration d'<strong>UpcycleConnect</strong> !</p>
				<p>Vous pouvez dès à présent vous connecter à votre espace et accéder à vos services.</p>
				<br>
				<p>À très vite,</p>
				<p><em>L'équipe UpcycleConnect</em></p>
			</body>
		</html>
	`, prenom)

	messageComplet := []byte(sujet + typeMIME + corpsMessage)

	adresseServeur := serveurSMTP + ":" + port
	err := smtp.SendMail(adresseServeur, auth, expediteur, []string{emailDestinataire}, messageComplet)
	if err != nil {
		return fmt.Errorf("erreur de connexion à hMailServer (192.168.80.10) : %v", err)
	}

	fmt.Println(" E-mail envoyé avec succès via la DMZ à", emailDestinataire)
	return nil
}

// Dans userReq.go (à la suite de ta fonction EnvoyerEmailValidation)

func EnvoyerEmailRefus(emailDestinataire string, prenom string, motif string) {
	expediteur := "noreply@upcycleconnect.fr"
	motDePasse := "voir avec ndoya"
	serveurSMTP := "192.168.80.10"
	port := "25" // verif avec ndoya

	auth := smtp.PlainAuth("", expediteur, motDePasse, serveurSMTP)

	sujet := "Subject: UpcycleConnect - Information concernant votre inscription\n"
	typeMIME := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	corpsMessage := fmt.Sprintf(`
		<html>
			<body style="font-family: Arial, sans-serif; color: #333;">
				<h2>Bonjour %s,</h2>
				<p>Nous faisons suite à votre demande d'inscription sur la plateforme <strong>UpcycleConnect</strong>.</p>
				<p>Après examen, nous sommes au regret de vous informer que nous ne pouvons pas valider votre compte pour la raison suivante :</p>
				<blockquote style="border-left: 4px solid #e74c3c; padding-left: 15px; color: #555; font-style: italic; background-color: #f9f9f9; padding: 10px;">
					%s
				</blockquote>
				<p>Si vous pensez qu'il s'agit d'une erreur ou si vous avez des éléments complémentaires à nous fournir, n'hésitez pas à nous contacter.</p>
				<br>
				<p>Cordialement,</p>
				<p><em>L'équipe UpcycleConnect</em></p>
			</body>
		</html>
	`, prenom, motif)

	messageComplet := []byte(sujet + typeMIME + corpsMessage)
	adresseServeur := serveurSMTP + ":" + port

	err := smtp.SendMail(adresseServeur, auth, expediteur, []string{emailDestinataire}, messageComplet)
	if err != nil {
		fmt.Printf("Erreur d'envoi d'e-mail de refus à %s : %v\n", emailDestinataire, err)
		return
	}
	fmt.Printf(" E-mail de refus envoyé avec succès à %s\n", emailDestinataire)
}
