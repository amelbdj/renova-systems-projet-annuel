package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"upcycleconnect/bdd"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AnnonceId  int `json:"annonce_id"`
		AcheteurId int `json:"acheteur_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	annonce, err := bdd.GetAnnonceById(req.AnnonceId)
	if err != nil {
		http.Error(w, "Annonce introuvable", http.StatusNotFound)
		return
	}

	orderId, err := bdd.CreateOrder(annonce, req.AcheteurId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "ID de commande : %d", orderId)
}

func PaymentHistoryHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userID, _ := strconv.Atoi(r.URL.Query().Get("id"))

	data, err := bdd.PaymentHistory(userID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(data)
}
type FinanceOverview struct {
	VolumeMois float64 `json:"volumeMois"`
	RevenuMois float64 `json:"revenuMois"` // Tes 5% de commission
	Evolution  int     `json:"evolution"`
}

func FinanceOverviewHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Gestion du CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, `{"erreur": "Méthode non autorisée"}`, http.StatusMethodNotAllowed)
		return
	}

	// 2. Appel à la BDD (Logique métier)
	volume, commission, err := bdd.GetFinanceOverviewMois()
	if err != nil {
		http.Error(w, `{"erreur": "Erreur lors du calcul des finances"}`, http.StatusInternalServerError)
		fmt.Printf("Erreur GetFinanceOverviewMois : %v\n", err)
		return
	}

	// 3. Préparation des données pour le JS
	stats := FinanceOverview{
		VolumeMois: volume,
		RevenuMois: commission,
		Evolution:  5, // On laisse à 5% en dur pour le moment pour le design
	}

	// 4. Envoi de la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stats)
}




func AdminTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Gestion du CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, `{"erreur": "Méthode non autorisée"}`, http.StatusMethodNotAllowed)
		return
	}

	// 2. Appel à la BDD
	transactions, err := bdd.GetAdminTransactions()
	if err != nil {
		fmt.Println("CRASH TRANSACTIONS:", err.Error())
		http.Error(w, `{"erreur": "Erreur lors de la récupération des transactions"}`, http.StatusInternalServerError)
		return
	}

	// 3. Si aucune transaction, on renvoie un tableau vide pour ne pas faire planter le JS
	if transactions == nil {
		transactions = []map[string]interface{}{}
	}

	// 4. Envoi de la réponse JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}