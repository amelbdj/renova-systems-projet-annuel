package admin

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"

	"upcycleconnect/auth"
	"upcycleconnect/bdd"

	"net/http"
	"strconv"
	"upcycleconnect/models"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/account"
)

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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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

func Inscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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

	hashedPwd, _ := auth.HashPassword(newUser.MotDePasse)
	newUser.MotDePasse = hashedPwd

	_, err = bdd.CreateUser(newUser)
	if err != nil {
		http.Error(w, "Database error or Email already exists", http.StatusInternalServerError)
		return
	}

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

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

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

	if user.StripeAccountId != "" && !user.StripeVerifCompleted {
		stripe.Key = StripeSecretKey
		acc, err := account.GetByID(user.StripeAccountId, nil)

		if err == nil && acc.PayoutsEnabled {
			_, execErr := bdd.Db.Exec("UPDATE utilisateur SET stripe_verif_completed = 1 WHERE id = ?", id)
			if execErr == nil {
				user.StripeVerifCompleted = true
			}
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
	nameQuery := r.URL.Query().Get("name") // recup les valeurs de la query string(diff de path variable)
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
	err = bdd.ValidateUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "utilisateur validé")
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

	err = bdd.RefuseUser(id, req.Motif)
	if err != nil {
		http.Error(w, "Erreur BDD : "+err.Error(), http.StatusInternalServerError)
		return
	}

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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
