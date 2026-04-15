package admin

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"upcycleconnect/auth"
	"upcycleconnect/bdd"

	"net/http"
	"strconv"
	"upcycleconnect/models"
)

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

	userBdd, err := bdd.LoginUser(user.Email, user.MotDePasse)
	if err != nil {
		fmt.Println("Erreur login :", err)
		http.Error(w, "Email ou mot de passe incorrect", http.StatusUnauthorized)
		return
	}

	var monStatut string

	if userBdd.TypeStatut == nil {
		monStatut = "En attente"
	} else {
		monStatut = *userBdd.TypeStatut
	}

	token, err := auth.GenerateJWT(userBdd.Id, userBdd.Role)
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	reponse := map[string]string{
		"token":  token,
		"role":   userBdd.Role,
		"statut": monStatut,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reponse)
}

func Inscription(w http.ResponseWriter, r *http.Request) {
	// 1. The "Security Bouncers" (CORS)
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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}

	user, err := bdd.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUserByRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
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

	// 5. Réponse Pro en JSON
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"message": "Utilisateur refusé avec motif enregistré"}`)
}

func UploadDocumentHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Gestion des CORS (comme tu as fait pour CreateUser)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Println("hello from uploadDocument")

	// 2. On limite la taille du fichier (ici 10 MB maximum)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println("Erreur ParseMultipartForm :", err)
		http.Error(w, "Fichier trop lourd ou formulaire invalide", http.StatusBadRequest)
		return
	}

	// 3. Récupération des champs textes (envoyés par le JS)
	userID := r.FormValue("user_id")
	typeDocument := r.FormValue("type_document") // ex: "KBIS" ou "DIPLOME"

	// 4. Récupération du fichier physique
	file, handler, err := r.FormFile("document")
	if err != nil {
		fmt.Println("Erreur récupération fichier :", err)
		http.Error(w, "Impossible de lire le fichier joint", http.StatusBadRequest)
		return
	}
	defer file.Close() // Très important pour ne pas bloquer la mémoire !

	// 5. Création du dossier "uploads" s'il n'existe pas
	cheminDossier := "./uploads/documents/"
	err = os.MkdirAll(cheminDossier, os.ModePerm)
	if err != nil {
		http.Error(w, "Erreur serveur : impossible de créer le dossier", http.StatusInternalServerError)
		return
	}

	// 6. Sauvegarde du fichier sur le serveur
	// On met l'ID de l'utilisateur dans le nom du fichier pour éviter les doublons
	nomFichierFinal := fmt.Sprintf("user_%s_%s", userID, handler.Filename)
	cheminComplet := filepath.Join(cheminDossier, nomFichierFinal)

	dst, err := os.Create(cheminComplet)
	if err != nil {
		fmt.Println("Erreur création fichier serveur :", err)
		http.Error(w, "Erreur lors de l'enregistrement du fichier", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	io.Copy(dst, file) // On copie le contenu du fichier téléchargé dans notre nouveau fichier

	// 7. Enregistrement en Base de Données
	err = bdd.InsertDocument(userID, typeDocument, cheminComplet)
	if err != nil {
		fmt.Println("Erreur BDD :", err)
		http.Error(w, "Erreur base de données", http.StatusInternalServerError)
		return
	}

	// 8. Réponse finale (JSON)
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
