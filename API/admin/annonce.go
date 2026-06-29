package admin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"upcycleconnect/bdd"
	"upcycleconnect/models"

	"strconv"
)

func GetAllAnnonces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	fmt.Println("hello from GetAllAnnonces")

	Annonces, err := bdd.GetAnnonces()

	if err != nil {
		http.Error(w, "erreur de récupération des annonces", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}

	response, err := json.Marshal(Annonces)

	if err != nil {
		http.Error(w, "erreur de conversion", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", response)
}

func ValidateAnnonce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}

	annonce, errGet := bdd.GetAnnonceById(id)

	err = bdd.ValidateAnnonce(id)
	if err != nil {
		http.Error(w, "erreur de validation de l'annonce", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}

	if errGet == nil && annonce.IdUser != 0 {
		msg := fmt.Sprintf("✅ Bonne nouvelle ! Ton annonce '%s' a été validée et est en ligne.", annonce.Titre)
		go SendPushNotification(strconv.Itoa(annonce.IdUser), msg)

		go NotifyAllPros(fmt.Sprintf("🆕 Nouveau matériau disponible : '%s'. Réservez-le dans le catalogue !", annonce.Titre))
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Annonce validée avec succès")
}

func RefuseAnnonce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}
	err = bdd.RefuseAnnonce(id)

	if err != nil {
		http.Error(w, "erreur de refus de l'annonce", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Annonce refusée avec succès")
}

func CreateAnnonce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if err := r.ParseMultipartForm(30 << 20); err != nil {
		http.Error(w, "Formulaire invalide", http.StatusBadRequest)
		return
	}

	var ann models.Annonce
	ann.Titre = r.FormValue("titre")
	ann.Description = r.FormValue("description")
	ann.Type = r.FormValue("type")
	ann.Prix, _ = strconv.ParseFloat(r.FormValue("prix"), 64)
	ann.IdCategorie, _ = strconv.Atoi(r.FormValue("id_categorie"))
	ann.IdUser, _ = strconv.Atoi(r.FormValue("id_user"))
	ann.Ville = r.FormValue("ville")
	ann.CodePostal = r.FormValue("code_postal")
	ann.Etat = r.FormValue("etat")
	ann.PoidsKg, _ = strconv.ParseFloat(r.FormValue("poids_kg"), 64)
	ann.Quantite, _ = strconv.Atoi(r.FormValue("quantite"))

	var stripeID string
	err := bdd.Db.QueryRow("SELECT stripe_account_id FROM utilisateur WHERE id = ?", ann.IdUser).Scan(&stripeID)

	if err != nil || stripeID == "" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "STRIPE_NOT_CONFIGURED")
		return
	}

	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()

		filePath := "./uploads/" + header.Filename

		dst, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Erreur stockage image", http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		io.Copy(dst, file)

		ann.Image = "/view-uploads/" + header.Filename
	}

	err = bdd.CreateAnnonce(ann)
	if err != nil {
		http.Error(w, "Erreur BDD", http.StatusInternalServerError)
		return
	}

	userIDStr := strconv.Itoa(ann.IdUser)
	message := fmt.Sprintf("Felicitations ! Votre annonce '%s' a bien ete cree.", ann.Titre)

	go SendPushNotification(userIDStr, message)

	NotifyAllAdmins(fmt.Sprintf("📢 Nouvelle annonce à valider : %s", ann.Titre))

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Annonce créée avec succès")
}

func DeleteAnnonce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")

	fmt.Println("hello from DeleteAnnonce")
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

	var userID int
	var titre string
	err = bdd.Db.QueryRow("SELECT id_user, titre FROM annonce WHERE id = ?", id).Scan(&userID, &titre)
	if err != nil {
		fmt.Println("Erreur lors de la recup des info:", err)
	}

	err = bdd.DeleteAnnonce(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if userID != 0 {
		userIDStr := strconv.Itoa(userID)
		message := fmt.Sprintf("Votre annonce '%s' a  ete supprimee.", titre)

		go SendPushNotification(userIDStr, message)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "annonce suppr")
}

func UpdateAnnonce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(30 << 20); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	var annonce models.Annonce
	annonce.Titre = r.FormValue("titre")
	annonce.Description = r.FormValue("description")
	annonce.Type = r.FormValue("type")
	annonce.Prix, _ = strconv.ParseFloat(r.FormValue("prix"), 64)
	annonce.IdCategorie, _ = strconv.Atoi(r.FormValue("id_categorie"))
	annonce.IdUser, _ = strconv.Atoi(r.FormValue("id_user"))
	annonce.CodePostal = r.FormValue("code_postal")
	annonce.Ville = r.FormValue("ville")
	annonce.Etat = r.FormValue("etat")
	annonce.PoidsKg, _ = strconv.ParseFloat(r.FormValue("poids_kg"), 64)
	annonce.Quantite, _ = strconv.Atoi(r.FormValue("quantite"))

	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		filePath := "./uploads/" + header.Filename
		dst, _ := os.Create(filePath)
		defer dst.Close()
		io.Copy(dst, file)
		annonce.Image = "/view-uploads/" + header.Filename
	} else {
		annonce.Image = r.FormValue("old_image_path")
	}

	err = bdd.UpdateAnnonce(id, annonce)
	if err != nil {
		http.Error(w, "erreur dans la mise à jour de l'annonce", http.StatusInternalServerError)
		return
	}

	if annonce.IdUser != 0 {
		userIDStr := strconv.Itoa(annonce.IdUser)
		message := fmt.Sprintf("Votre annonce '%s' a ete modifiee", annonce.Titre)

		go SendPushNotification(userIDStr, message)
	}

	w.WriteHeader(http.StatusOK)
}

func GetAnnonceByTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	fmt.Println("hello from GetAnnonceByTitle")
	query := r.URL.Query().Get("query")
	filtre := r.URL.Query().Get("filtre")
	Annonces, err := bdd.GetAnnonceByTitle(query, filtre)

	if err != nil {
		http.Error(w, "erreur de récupération des annonces", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}
	response, err := json.Marshal(Annonces)
	if err != nil {
		http.Error(w, "erreur de conversion", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", response)
}

func GetMyAnnonces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID nvalide", http.StatusBadRequest)
		return
	}

	annonces, err := bdd.GetAnnoncesByUser(userID)
	if err != nil {
		http.Error(w, "Erreur lors de la recup de vos annonces", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(annonces)
}

func ToggleSponsorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	annonceID, _ := strconv.Atoi(r.URL.Query().Get("annonce_id"))
	if userID == 0 || annonceID == 0 {
		http.Error(w, `{"error": "Paramètres manquants"}`, http.StatusBadRequest)
		return
	}

	var owner int
	var plan, sponsored string
	err := bdd.Db.QueryRow(`
		SELECT a.id_user, COALESCE(u.plan_abo, ''), COALESCE(a.is_sponsored, 0)
		FROM pa2026.annonce a
		JOIN pa2026.utilisateur u ON a.id_user = u.id
		WHERE a.id = ?`, annonceID).Scan(&owner, &plan, &sponsored)
	if err != nil {
		http.Error(w, `{"error": "Annonce introuvable"}`, http.StatusNotFound)
		return
	}

	if owner != userID {
		http.Error(w, `{"error": "Cette annonce ne vous appartient pas"}`, http.StatusForbidden)
		return
	}
	if plan != "pro" {
		http.Error(w, `{"error": "Le boost est réservé au plan Pro"}`, http.StatusForbidden)
		return
	}

	newState := 1
	if sponsored == "1" {
		newState = 0
	}
	if _, err := bdd.Db.Exec("UPDATE pa2026.annonce SET is_sponsored = ? WHERE id = ?", newState, annonceID); err != nil {
		http.Error(w, `{"error": "Impossible de mettre à jour l'annonce"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"is_sponsored": newState})
}

func GetValidatedAnnonces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	idStr := r.URL.Query().Get("id")
	currentUserID, _ := strconv.Atoi(idStr)

	annonces, err := bdd.GetValidatedAnnonces(currentUserID)

	if err != nil {
		fmt.Println("Erreur lors de la recup des annonces validées : ", err)
		http.Error(w, "Erreur recup des annonces", http.StatusInternalServerError)
		return
	}

	if annonces == nil {
		annonces = []models.Annonce{}
	}

	json.NewEncoder(w).Encode(annonces)
}

func GetOneAnnonce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		http.Error(w, "ID manquant", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	annonce, err := bdd.GetAnnonceById(id)
	if err != nil {
		fmt.Printf("LOG_DÉTAIL: Impossible de trouver l'annonce %d : %v\n", id, err)
		http.Error(w, "Annonce introuvable", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(annonce)
}

func ConfirmPaymentAndOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	annonceID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	buyerID, _ := strconv.Atoi(r.URL.Query().Get("buyer_id"))

	annonce, err := bdd.GetAnnonceById(annonceID)
	if err != nil {
		http.Error(w, "Annonce introuvable", http.StatusNotFound)
		return
	}

	orderID, err := bdd.CreateOrder(annonce, buyerID)
	if err != nil {
		http.Error(w, "Erreur création commande: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if _, pdfErr := GenerateInvoicePDF(orderID, buyerID, annonce.Titre, annonce.Prix); pdfErr != nil {
		fmt.Println("Erreur génération facture:", pdfErr)
	}

	var sellerID int
	err = bdd.Db.QueryRow("SELECT id_user FROM annonce WHERE id = ?", annonceID).Scan(&sellerID)
	if err != nil {
		http.Error(w, "Impossible de trouver le vendeur", http.StatusInternalServerError)
		return
	}

	var conteneurID int
	err = bdd.Db.QueryRow(`
        SELECT c.id 
        FROM conteneur c
        JOIN box b ON c.id = b.id_conteneur
        WHERE b.statut = 'libre' 
        LIMIT 1
    `).Scan(&conteneurID)

	if err != nil {
		http.Error(w, "Aucun conteneur avec des box libres n'est disponible", http.StatusInternalServerError)
		return
	}

	err = bdd.ReserveBox(annonceID, conteneurID, sellerID)
	if err != nil {
		http.Error(w, "Erreur logistique box: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if sellerID != 0 {
		sellerIDStr := strconv.Itoa(sellerID)
		msgSeller := fmt.Sprintf("Vendu ! Votre objet '%s' a ete achete. Une box a ete resercee pour votre depot.", annonce.Titre)
		go SendPushNotification(sellerIDStr, msgSeller)
	}

	if buyerID != 0 {
		buyerIDStr := strconv.Itoa(buyerID)
		msgBuyer := fmt.Sprintf("Paiement valide ! Le vendeur va bientot deposer '%s' dans la box.", annonce.Titre)
		go SendPushNotification(buyerIDStr, msgBuyer)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   "success",
		"order_id": orderID,
		"box_id":   conteneurID,
		"message":  "Paiement valide, annonce passee en EN ATTENTE DEPOT, et Box donnee au vendeur",
	})
}

func GetMyBoxes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))

	data, err := bdd.GetUserReservations(userID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if data == nil {
		data = []map[string]interface{}{}
	}

	json.NewEncoder(w).Encode(data)
}
func GetEcoStatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	userIDStr := r.URL.Query().Get("user_id")
	userID, _ := strconv.Atoi(userIDStr)

	stats, err := bdd.GetUserEcoStats(userID)
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stats)
}
func GetMyPurchases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))

	data, err := bdd.GetUserPurchases(userID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if data == nil {
		data = []map[string]interface{}{}
	}

	json.NewEncoder(w).Encode(data)
}
