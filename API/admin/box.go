package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"upcycleconnect/bdd"
)



func ReserveBox(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS") // ⚠️ Ne pas oublier les méthodes
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {
		var req struct {
			AnnonceId     int `json:"annonce_id"`
			ConteneurId   int `json:"conteneur_id"` // L'ID du meuble !
			ParticulierId int `json:"particulier_id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}

		err := bdd.ReserveBox(req.AnnonceId, req.ConteneurId, req.ParticulierId)
		if err != nil {
			fmt.Println("Erreur ReserveBox :", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"message": "Box réservée avec succès"}`)
		return
	}
}

func ConfirmDeposit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {
		var req struct {
			PinCode string `json:"pin_code"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Code PIN manquant", http.StatusBadRequest)
			return
		}

		err := bdd.ConfirmDeposit(req.PinCode)
		if err != nil {
			fmt.Println("Erreur ConfirmDeposit :", err)
			http.Error(w, "Code PIN incorrect ou expiré", http.StatusUnauthorized)
			return
		}

		// --- NOTIFICATION ACHETEUR ---
		// On cherche à qui appartient cet objet et quel est le code pour l'ouvrir
		var acheteurID int
		var titre string
		var numBox string
		var codeRetrait string

		// ⚠️ Adapte le nom de tes tables/colonnes si elles sont un peu différentes
		query := `
			SELECT o.acheteur_id, a.titre, b.id, b.pin_code 
			FROM box b
			JOIN annonce a ON b.id_annonce = a.id
			JOIN orders o ON o.annonce_id = a.id
			WHERE b.pin_code = ? LIMIT 1
		`
		errInfo := bdd.Db.QueryRow(query, req.PinCode).Scan(&acheteurID, &titre, &numBox, &codeRetrait)

		if errInfo == nil && acheteurID != 0 {
			msg := fmt.Sprintf("🔓 Ton objet '%s' t'attend ! Tu peux le récupérer au Casier n°%s.", titre, numBox)
			go SendPushNotification(strconv.Itoa(acheteurID), msg)
		} else {
			fmt.Println("Impossible de trouver l'acheteur pour lui envoyer la notif :", errInfo)
		}
		// -----------------------------

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"message": "Dépôt validé, la box est verrouillée"}`)
		return
	}
}

func CollectObject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS") // Souvent en POST pour envoyer des données JSON
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {
		var req struct {
			Barcode         string `json:"barcode"`
			ProfessionnelId int    `json:"professionnel_id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Données de collecte invalides", http.StatusBadRequest)
			return
		}

		err := bdd.CollectObject(req.Barcode, req.ProfessionnelId)
		if err != nil {
			fmt.Println("Erreur CollectObject :", err)
			http.Error(w, "Erreur lors de la collecte : "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"message": "Objet récupéré, box libérée et score mis à jour !"}`)
		return
	}
}




// Remplace  "GetAllBoxs"
func GetConteneursAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from GetConteneursAdmin")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "GET" {
		conteneurs, err := bdd.GetConteneursAdmin()
		if err != nil {
			fmt.Println("Erreur BDD GetConteneursAdmin :", err)
			http.Error(w, "Erreur lors de la récupération des conteneurs", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(conteneurs)
		return
	}
}

// Remplace  "CreateBox"
func CreateConteneur(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {
		var req struct {
			Nom          string `json:"nom"`
			Adresse      string `json:"adresse"`
			NombreDeBoxs int    `json:"nombre_de_boxs"` // L'admin choisit combien de portes il y a dans ce meuble
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}

		err := bdd.CreateConteneurAvecBox(req.Nom, req.Adresse, req.NombreDeBoxs)
		if err != nil {
			fmt.Println("Erreur création conteneur :", err)
			http.Error(w, "Erreur serveur lors de la création", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"message": "Conteneur et boxes créés avec succès"}`)
		return
	}
}

func GetBoxesForConteneurHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "GET" {
		conteneurID := r.PathValue("id")
		
		if conteneurID == "" {
			http.Error(w, "ID du conteneur manquant", http.StatusBadRequest)
			return
		}

		// On appelle la fonction BDD qu'on vient de créer
		boxes, err := bdd.GetBoxesByConteneurID(conteneurID)
		if err != nil {
			fmt.Println("Erreur BDD GetBoxesForConteneurHandler :", err)
			http.Error(w, "Erreur serveur", http.StatusInternalServerError)
			return
		}

		// On envoie le tableau au JavaScript
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(boxes)
		return
	}
}

func AddSingleBoxHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {
		var req struct {
			IDConteneur int    `json:"id_conteneur"`
			Taille      string `json:"taille"` // "S", "M", ou "L"
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}

		if req.Taille == "" {
			req.Taille = "M"
		}

		err := bdd.AddSingleBoxToConteneur(req.IDConteneur, req.Taille)
		if err != nil {
			fmt.Println("Erreur AddSingleBox :", err)
			http.Error(w, "Erreur serveur", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"message": "Nouvelle porte ajoutée avec succès"}`)
		return
	}
}

func UpdateBoxStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "PUT" {
		var req struct {
			BoxID  int    `json:"box_id"`
			Statut string `json:"statut"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			fmt.Println("Erreur décodage UpdateBoxStatusHandler :", err)
			return
		}
		err := bdd.UpdateBoxStatus(req.BoxID, req.Statut)
		if err != nil {
			fmt.Println("Erreur UpdateBoxStatus :", err)
			http.Error(w, "Erreur serveur", http.StatusInternalServerError)
			fmt.Println("Erreur UpdateBoxStatusHandler :", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"message": "Statut mis à jour"}`)
		return
	}
}