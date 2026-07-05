package admin

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"upcycleconnect/auth"
	"upcycleconnect/bdd"

	"net/http"
	"strconv"
	"upcycleconnect/models"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/account"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"github.com/stripe/stripe-go/v81/subscription"
	"golang.org/x/crypto/bcrypt"

	portalsession "github.com/stripe/stripe-go/v81/billingportal/session"
)

type UpdatePasswordInput struct {
	ID          int    `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return strings.Split(forwarded, ",")[0]
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Println("hello from login")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Erreur décodage :", err)
		http.Error(w, "Impossible de décoder le JSON", http.StatusBadRequest)
		return
	}

	clientIP := GetIP(r)

	userBdd, err := bdd.LoginUser(user.Email, user.MotDePasse, clientIP)
	if err != nil {
		fmt.Println("Erreur login :", err)
		http.Error(w, "Email ou mot de passe incorrect", http.StatusUnauthorized)
		return
	}

	var validation string

	if userBdd.Validation == "" {
		validation = "En attente"
	} else {
		validation = userBdd.Validation
	}

	token, err := auth.GenerateJWT(userBdd.Id, userBdd.Role)
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	reponse := map[string]interface{}{
		"token":      token,
		"id":         userBdd.Id,
		"role":       userBdd.Role,
		"prenom":     userBdd.Prenom,
		"score":      userBdd.Score,
		"tutorielVu": userBdd.TutorielVu,
		"validation": validation,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reponse)
}

// siretValide verifie qu'un SIRET est correct : 14 chiffres + cle de controle
// de Luhn (on double un chiffre sur deux en partant de la droite, et la somme
// totale doit etre un multiple de 10). Detecte les SIRET inventes.
func siretValide(siret string) bool {
	if len(siret) != 14 {
		return false
	}
	somme := 0
	for i := 0; i < 14; i++ {
		c := siret[i]
		if c < '0' || c > '9' {
			return false // caractere non numerique
		}
		n := int(c - '0')
		// Position depuis la droite = 14 - i ; on double les positions paires.
		if (14-i)%2 == 0 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		somme += n
	}
	return somme%10 == 0
}

func Inscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	// Securite : on interdit l'auto-inscription en tant qu'Administrateur ou Salarie.
	// Ces comptes internes sont crees uniquement par un admin (via /admin/users/add).
	// HasPrefix "Salari" gere aussi "Salarie" et les variantes d'encodage du "é".
	if newUser.Role == "Administrateur" || strings.HasPrefix(newUser.Role, "Salari") {
		http.Error(w, "Création de ce type de compte non autorisée", http.StatusForbidden)
		return
	}

	// Un professionnel doit fournir un SIRET valide (14 chiffres + cle de Luhn).
	if newUser.Role == "Pro" || newUser.Role == "Professionnel" {
		siret := ""
		if newUser.Siret != nil {
			siret = *newUser.Siret
		}
		if !siretValide(siret) {
			http.Error(w, "SIRET invalide : 14 chiffres avec une clé de contrôle correcte", http.StatusBadRequest)
			return
		}
	}

	hashedPwd, _ := auth.HashPassword(newUser.MotDePasse)
	newUser.MotDePasse = hashedPwd

	_, err = bdd.CreateUser(newUser)
	if err != nil {
		http.Error(w, "Database error or Email already exists", http.StatusInternalServerError)
		return
	}
	fmt.Println("✅ Inscription réussie, j'alerte les admins !")
	NotifyAllAdmins("👤 Un nouvel utilisateur s'est  sur UpcycleConnect !")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from GetAllUsers")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	users, err := bdd.GetUsers()

	if err != nil {
		fmt.Println(err)
		http.Error(w, "erreur de récupération des utilisateurs", http.StatusInternalServerError)

		return
	}

	response, err := json.Marshal(users)

	if err != nil {
		http.Error(w, "erreur de conversion", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Println("hello from createUser")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Erreur décodage :", err)
		http.Error(w, "Impossible de décoder le JSON", http.StatusBadRequest)
		return
	}

	user.MotDePasse, err = auth.HashPassword(user.MotDePasse)
	if err != nil {
		fmt.Println("Erreur hashage mot de passe :", err)
		http.Error(w, "Erreur lors du hashage du mot de passe", http.StatusInternalServerError)
		return
	}

	nouvelID, err := bdd.CreateUser(user)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion en BDD :", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	NotifyAllAdmins("👤 Un nouvel utilisateur s'est  sur UpcycleConnect !")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Utilisateur créé avec succès",
		"id":      nouvelID,
	})
}

func DeletedUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	fmt.Println("hello from DeletedUser")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}

	err = bdd.DeletedUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "utilisateur suppr")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Println("hello from updateUser")

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"ID invalide"}`, http.StatusBadRequest)
		return
	}

	var userDto models.User
	if err := json.NewDecoder(r.Body).Decode(&userDto); err != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		fmt.Println("Erreur décodage :", err)
		return

	}

	userDto.Id = id

	if err := bdd.UpdateUserById(userDto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, `{"error":"Donnees invalides"}`, http.StatusBadRequest)
		return
	}

	if user.Id == 0 {
		idContext, ok := r.Context().Value("userID").(int)
		if ok {
			user.Id = idContext
		}
	}

	if user.Id == 0 || user.Nom == "" || user.Prenom == "" || user.Email == "" {
		http.Error(w, `{"error":"Champs obligatoires manquants"}`, http.StatusBadRequest)
		return
	}

	err = bdd.UpdateUserProfile(user.Id, user.Nom, user.Prenom, user.Email)
	if err != nil {
		http.Error(w, `{"error":"Erreur lors de la mise a jour"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Profil mis a jour",
	})
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		idStr = r.PathValue("id")
	}
	id, _ := strconv.Atoi(idStr)

	user, err := bdd.GetUserById(id)
	if err != nil {
		http.Error(w, "Utilisateur introuvable", http.StatusNotFound)
		return
	}

	if user.StripeAccountId != nil && *user.StripeAccountId != "0" && *user.StripeAccountId != "" && !user.StripeVerifCompleted {
		stripe.Key = getStripeSecretKey()

		acc, err := account.GetByID(*user.StripeAccountId, nil)

		if err == nil && acc.PayoutsEnabled {
			_, execErr := bdd.Db.Exec("UPDATE utilisateur SET stripe_verif_completed = 1 WHERE id = ?", id)
			if execErr == nil {
				user.StripeVerifCompleted = true
			}
		} else if err != nil {
			fmt.Println("Erreur Stripe :", err)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUserByRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	role := r.PathValue("role")

	users, err := bdd.GetUserByRole(role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	nameQuery := r.URL.Query().Get("name")
	roleQuery := r.URL.Query().Get("role")

	users, err := bdd.GetUserByName(nameQuery, roleQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func ValidateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}

	prenom, email, err := bdd.ValidateUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go bdd.EnvoyerEmailValidation(email, prenom)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"message": "Utilisateur validé avec succès !"}`)
}

type RefuseRequest struct {
	Motif string `json:"motif"`
}

func RefuseUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}

	var req RefuseRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		req.Motif = "Non spécifié"
	}

	prenom, email, err := bdd.RefuseUser(id, req.Motif)
	if err != nil {
		http.Error(w, "Erreur BDD : "+err.Error(), http.StatusInternalServerError)
		return
	}

	go bdd.EnvoyerEmailRefus(email, prenom, req.Motif)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"message": "Utilisateur refusé avec motif enregistré"}`)
}

func UploadDocumentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Println("hello from uploadDocument")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println("Erreur ParseMultipartForm :", err)
		http.Error(w, "Fichier trop lourd ou formulaire invalide", http.StatusBadRequest)
		return
	}

	userID := r.FormValue("user_id")
	typeDocument := r.FormValue("type_document")

	file, handler, err := r.FormFile("document")
	if err != nil {
		fmt.Println("Erreur récupération fichier :", err)
		http.Error(w, "Impossible de lire le fichier joint", http.StatusBadRequest)
		return
	}
	defer file.Close()

	cheminDossier := "./uploads/documents/"
	err = os.MkdirAll(cheminDossier, os.ModePerm)
	if err != nil {
		http.Error(w, "Erreur serveur : impossible de créer le dossier", http.StatusInternalServerError)
		return
	}

	nomFichierFinal := fmt.Sprintf("user_%s_%s", userID, handler.Filename)
	cheminComplet := filepath.Join(cheminDossier, nomFichierFinal)

	dst, err := os.Create(cheminComplet)
	if err != nil {
		fmt.Println("Erreur création fichier serveur :", err)
		http.Error(w, "Erreur lors de l'enregistrement du fichier", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	io.Copy(dst, file)

	err = bdd.InsertDocument(userID, typeDocument, cheminComplet)
	if err != nil {
		fmt.Println("Erreur BDD :", err)
		http.Error(w, "Erreur base de données", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, `{"message": "Document sauvegardé avec succès"}`)
}

func VerifierEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	emailTape := r.URL.Query().Get("email")

	emailExiste, err := bdd.CheckEmailExists(emailTape)

	if err != nil {
		http.Error(w, "Erreur du serveur ou de la base de données", http.StatusInternalServerError)
		return
	}

	if emailExiste == true {
		fmt.Fprintf(w, "true")
	} else {
		fmt.Fprintf(w, "false")
	}
}

func UpdateTutorialStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var data struct {
		UserID int `json:"id"`
	}
	json.NewDecoder(r.Body).Decode(&data)

	_, err := bdd.Db.Exec("UPDATE pa2026.utilisateur SET tutoriel_vu = 1 WHERE id = ?", data.UserID)

	if err != nil {
		http.Error(w, "Erreur BDD", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func BanUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	idStr := r.PathValue("id")
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID utilisateur invalide", http.StatusBadRequest)
		return
	}

	// On récupère l'utilisateur AVANT le bannissement pour avoir son email et son prénom
	user, errUser := bdd.GetUserById(userId)

	err = bdd.BanUser(userId)
	if err != nil {
		http.Error(w, "Erreur serveur lors du bannissement", http.StatusInternalServerError)
		return
	}

	// On envoie l'email d'information (en arrière-plan, pour ne pas bloquer la réponse)
	if errUser == nil && user.Email != "" {
		go bdd.EnvoyerEmailBannissement(user.Email, user.Prenom)
	} else {
		fmt.Println("Bannissement : impossible d'envoyer l'email (utilisateur introuvable)")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Utilisateur banni avec succès"}`)
}

func UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Méthode non autorisée"}`, http.StatusMethodNotAllowed)
		return
	}

	var input UpdatePasswordInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error": "Données invalides"}`, http.StatusBadRequest)
		return
	}

	var hashedDBPassword string
	err := bdd.Db.QueryRow("SELECT mot_de_passe FROM pa2026.utilisateur WHERE id = ?", input.ID).Scan(&hashedDBPassword)
	if err != nil {
		http.Error(w, `{"error": "Utilisateur non trouvé"}`, http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedDBPassword), []byte(input.OldPassword))
	if err != nil {
		http.Error(w, `{"error": "L'ancien mot de passe est incorrect"}`, http.StatusUnauthorized)
		return
	}
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), 10)
	if err != nil {
		http.Error(w, `{"error": "Erreur lors du hachage"}`, http.StatusInternalServerError)
		return
	}
	_, err = bdd.Db.Exec("UPDATE pa2026.utilisateur SET mot_de_passe = ? WHERE id = ?", string(newHashedPassword), input.ID)
	if err != nil {
		http.Error(w, `{"error": "Erreur de mise à jour en BDD"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Mot de passe mis à jour avec succès !"})
}

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok || userID == 0 {
		http.Error(w, `{"error": "Vous devez être connecté"}`, http.StatusUnauthorized)
		return
	}

	var input struct {
		Email       string `json:"email"`
		NewPassword string `json:"new_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error": "Données invalides"}`, http.StatusBadRequest)
		return
	}
	if len(input.NewPassword) < 6 {
		http.Error(w, `{"error": "Le mot de passe doit faire au moins 6 caractères"}`, http.StatusBadRequest)
		return
	}

	var foundID int
	if err := bdd.Db.QueryRow("SELECT id FROM pa2026.utilisateur WHERE id = ? AND email = ?", userID, input.Email).Scan(&foundID); err != nil {
		http.Error(w, `{"error": "L'adresse e-mail ne correspond pas à votre compte"}`, http.StatusBadRequest)
		return
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), 10)
	if err != nil {
		http.Error(w, `{"error": "Erreur lors du hachage"}`, http.StatusInternalServerError)
		return
	}
	_, err = bdd.Db.Exec("UPDATE pa2026.utilisateur SET mot_de_passe = ? WHERE id = ?", string(newHashedPassword), userID)
	if err != nil {
		http.Error(w, `{"error": "Erreur de mise à jour en BDD"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Mot de passe réinitialisé avec succès !"})
}

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var input struct {
		Email string `json:"email"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, `{"error": "Donnees invalides"}`, http.StatusBadRequest)
		return
	}

	email := strings.TrimSpace(input.Email)
	if email == "" {
		http.Error(w, `{"error": "Email obligatoire"}`, http.StatusBadRequest)
		return
	}

	var userID int
	var prenom string
	err = bdd.Db.QueryRow("SELECT id, prenom FROM pa2026.utilisateur WHERE email = ?", email).Scan(&userID, &prenom)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Si le compte existe, un email a ete envoye."})
		return
	}

	creerTableResetPassword()

	token, err := genererTokenReset()
	if err != nil {
		http.Error(w, `{"error": "Erreur serveur"}`, http.StatusInternalServerError)
		return
	}

	bdd.Db.Exec("UPDATE pa2026.reset_password_token SET used = 1 WHERE id_user = ? AND used = 0", userID)
	_, err = bdd.Db.Exec(`
		INSERT INTO pa2026.reset_password_token (id_user, token, expires_at, used)
		VALUES (?, ?, DATE_ADD(NOW(), INTERVAL 1 HOUR), 0)
	`, userID, token)
	if err != nil {
		http.Error(w, `{"error": "Erreur BDD"}`, http.StatusInternalServerError)
		return
	}

	lien := frontURL("reset-password.html?token=" + token + "&email=" + email)
	go bdd.EnvoyerEmailResetPassword(email, prenom, lien)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Si le compte existe, un email a ete envoye."})
}

func ResetPasswordTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var input struct {
		Email       string `json:"email"`
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, `{"error": "Donnees invalides"}`, http.StatusBadRequest)
		return
	}

	input.Email = strings.TrimSpace(input.Email)
	input.Token = strings.TrimSpace(input.Token)
	if input.Email == "" || input.Token == "" {
		http.Error(w, `{"error": "Lien de reinitialisation invalide"}`, http.StatusBadRequest)
		return
	}

	if len(input.NewPassword) < 6 {
		http.Error(w, `{"error": "Le mot de passe doit faire au moins 6 caracteres"}`, http.StatusBadRequest)
		return
	}

	creerTableResetPassword()

	var userID int
	err = bdd.Db.QueryRow(`
		SELECT u.id
		FROM pa2026.reset_password_token r
		JOIN pa2026.utilisateur u ON u.id = r.id_user
		WHERE u.email = ? AND r.token = ? AND r.used = 0 AND r.expires_at > NOW()
	`, input.Email, input.Token).Scan(&userID)
	if err != nil {
		http.Error(w, `{"error": "Lien invalide ou expire"}`, http.StatusBadRequest)
		return
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), 10)
	if err != nil {
		http.Error(w, `{"error": "Erreur lors du hachage"}`, http.StatusInternalServerError)
		return
	}

	_, err = bdd.Db.Exec("UPDATE pa2026.utilisateur SET mot_de_passe = ? WHERE id = ?", string(newHashedPassword), userID)
	if err != nil {
		http.Error(w, `{"error": "Erreur de mise a jour en BDD"}`, http.StatusInternalServerError)
		return
	}

	bdd.Db.Exec("UPDATE pa2026.reset_password_token SET used = 1 WHERE token = ?", input.Token)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Mot de passe reinitialise avec succes !"})
}

func creerTableResetPassword() {
	bdd.Db.Exec(`
		CREATE TABLE IF NOT EXISTS pa2026.reset_password_token (
			id INT AUTO_INCREMENT PRIMARY KEY,
			id_user INT NOT NULL,
			token VARCHAR(100) NOT NULL,
			expires_at DATETIME NOT NULL,
			used TINYINT(1) DEFAULT 0,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
}

func genererTokenReset() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes) + fmt.Sprint(time.Now().Unix()), nil
}

func UpgradeToPremiumHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userID := r.URL.Query().Get("id")
	sessionID := r.URL.Query().Get("session_id") // 🟢 Get the session ID from the URL

	if userID == "" || sessionID == "" {
		http.Error(w, `{"error": "Missing parameters"}`, http.StatusBadRequest)
		return
	}

	stripe.Key = getStripeSecretKey()

	// Ask Stripe for the session details to get the Customer ID
	s, err := session.Get(sessionID, nil)
	if err != nil {
		fmt.Println("Error fetching Stripe session:", err)
		http.Error(w, `{"error": "Invalid session"}`, http.StatusInternalServerError)
		return
	}

	customerID := s.Customer.ID

	plan := s.Metadata["plan"]
	if plan != "plus" && plan != "pro" {
		plan = "premium"
	}

	query := "UPDATE utilisateur SET est_premium = 1, plan_abo = ?, stripe_customer_id = ? WHERE id = ?"
	_, dbErr := bdd.Db.Exec(query, plan, customerID, userID)
	if dbErr != nil {
		fmt.Println("Database error:", dbErr)
		http.Error(w, `{"error": "Could not upgrade user"}`, http.StatusInternalServerError)
		return
	}

	planId := 1
	if plan == "plus" {
		planId = 2
	} else if plan == "pro" {
		planId = 3
	}
	res, aboErr := bdd.Db.Exec("INSERT INTO abonnement (id_user, id_plan, date_debut, statut) VALUES (?, ?, NOW(), 'actif')", userID, planId)
	if aboErr != nil {
		fmt.Println("Erreur enregistrement abonnement (historique):", aboErr)
	} else {
		abonnementID, _ := res.LastInsertId()
		uid, _ := strconv.Atoi(userID)
		if _, pdfErr := GenerateContractPDF(uid, plan, int(abonnementID)); pdfErr != nil {
			fmt.Println("Erreur génération contrat:", pdfErr)
		}
	}

	newSubID := ""
	if s.Subscription != nil {
		newSubID = s.Subscription.ID
	}
	listParams := &stripe.SubscriptionListParams{
		Customer: stripe.String(customerID),
		Status:   stripe.String("active"),
	}
	iter := subscription.List(listParams)
	for iter.Next() {
		sub := iter.Subscription()
		if sub.ID != newSubID {
			if _, cancelErr := subscription.Cancel(sub.ID, nil); cancelErr != nil {
				fmt.Println("Impossible d'annuler l'ancien abonnement:", cancelErr)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Account upgraded!"})
}

func CustomerPortalHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userID := r.URL.Query().Get("id")

	// Fetch their Customer ID from your database
	var customerID string
	err := bdd.Db.QueryRow("SELECT stripe_customer_id FROM utilisateur WHERE id = ?", userID).Scan(&customerID)

	if err != nil || customerID == "" {
		http.Error(w, `{"error": "Customer not found"}`, http.StatusNotFound)
		return
	}

	stripe.Key = getStripeSecretKey()

	// Create the portal session
	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String(customerID),
		ReturnURL: stripe.String(frontURL("espPro.html")),
	}

	ps, err := portalsession.New(params)
	if err != nil {
		fmt.Println("Error creating portal:", err)
		http.Error(w, `{"error": "Stripe error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"url": ps.URL})
}

func CancelSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userID := r.URL.Query().Get("id")
	var customerID string
	err := bdd.Db.QueryRow("SELECT stripe_customer_id FROM utilisateur WHERE id = ?", userID).Scan(&customerID)
	if err != nil || customerID == "" {
		http.Error(w, `{"error": "Aucun compte Stripe trouvé"}`, http.StatusNotFound)
		return
	}

	stripe.Key = getStripeSecretKey()

	listParams := &stripe.SubscriptionListParams{
		Customer: stripe.String(customerID),
		Status:   stripe.String("active"),
	}
	iter := subscription.List(listParams)
	cancelled := 0
	for iter.Next() {
		sub := iter.Subscription()
		if _, cancelErr := subscription.Cancel(sub.ID, nil); cancelErr != nil {
			fmt.Println("Erreur annulation abonnement:", cancelErr)
		} else {
			cancelled++
		}
	}

	bdd.Db.Exec("UPDATE utilisateur SET est_premium = 0, plan_abo = NULL WHERE id = ?", userID)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "Abonnement résilié",
		"cancelled": cancelled,
	})
}

func SyncPremiumStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userID := r.URL.Query().Get("id")
	var customerID string

	// 1. Récupérer le Stripe Customer ID
	err := bdd.Db.QueryRow("SELECT stripe_customer_id FROM utilisateur WHERE id = ?", userID).Scan(&customerID)
	if err != nil || customerID == "" {
		http.Error(w, `{"error": "Aucun compte Stripe trouvé"}`, http.StatusNotFound)
		return
	}

	stripe.Key = getStripeSecretKey()

	// 2. Demander à Stripe si un abonnement "actif" existe pour ce client
	params := &stripe.SubscriptionListParams{
		Customer: stripe.String(customerID),
		Status:   stripe.String("active"),
	}
	iter := subscription.List(params)

	estPremium := 0
	if iter.Next() {
		estPremium = 1 // On a trouvé un abonnement actif !
	}

	// 3. Mettre à jour la base de données avec la vraie réponse de Stripe
	if estPremium == 0 {
		bdd.Db.Exec("UPDATE utilisateur SET est_premium = 0, plan_abo = NULL WHERE id = ?", userID)
	} else {
		bdd.Db.Exec("UPDATE utilisateur SET est_premium = 1 WHERE id = ?", userID)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"est_premium": estPremium})
}
