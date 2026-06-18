package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"upcycleconnect/bdd"
)

func GetProInvoicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	acheteurID, ok := r.Context().Value("userID").(int)
	if !ok || acheteurID == 0 {
		http.Error(w, `{"error": "Utilisateur non identifié"}`, http.StatusUnauthorized)
		return
	}

	factures, err := bdd.GetProInvoices(acheteurID)
	if err != nil {
		fmt.Println("Erreur GetProInvoices :", err)
		http.Error(w, `{"error": "Erreur serveur"}`, http.StatusInternalServerError)
		return
	}

	if factures == nil {
		factures = []map[string]interface{}{}
	}

	json.NewEncoder(w).Encode(factures)
}

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
	RevenuMois float64 `json:"revenuMois"`
	Evolution  int     `json:"evolution"`
}

func FinanceOverviewHandler(w http.ResponseWriter, r *http.Request) {

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

	volume, commission, err := bdd.GetFinanceOverviewMois()
	if err != nil {
		http.Error(w, `{"erreur": "Erreur lors du calcul des finances"}`, http.StatusInternalServerError)
		fmt.Printf("Erreur GetFinanceOverviewMois : %v\n", err)
		return
	}

	stats := FinanceOverview{
		VolumeMois: volume,
		RevenuMois: commission,
		Evolution:  5,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stats)
}

func AdminTransactionsHandler(w http.ResponseWriter, r *http.Request) {

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

	transactions, err := bdd.GetAdminTransactions()
	if err != nil {
		fmt.Println("CRASH TRANSACTIONS:", err.Error())
		http.Error(w, `{"erreur": "Erreur lors de la récupération des transactions"}`, http.StatusInternalServerError)
		return
	}

	if transactions == nil {
		transactions = []map[string]interface{}{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}
