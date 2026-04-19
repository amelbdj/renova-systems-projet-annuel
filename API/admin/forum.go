package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"upcycleconnect/bdd"
)

func GetForumMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	filtre := r.URL.Query().Get("filter")

	messages, err := bdd.GetForumMessages(filtre)
	if err != nil {
		http.Error(w, "Erreur serveur lors de la récupération des messages", http.StatusInternalServerError)
		fmt.Println("Erreur BDD GetForumMessages :", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func ModerateForumMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	idMessageStr := r.PathValue("id")
	idMessage, err := strconv.Atoi(idMessageStr)
	if err != nil {
		http.Error(w, "ID de message invalide", http.StatusBadRequest)
		return
	}

	var payload struct {
		Action string `json:"action"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Format de requête invalide", http.StatusBadRequest)
		return
	}

	err = bdd.ModerateForumMessage(idMessage, payload.Action)
	if err != nil {
		http.Error(w, "Erreur lors de la modération", http.StatusInternalServerError)
		fmt.Println("Erreur BDD ModerateForumMessage :", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Action de modération enregistrée avec succès"}`)
}

func GetForumStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	stats, err := bdd.GetForumStats()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des statistiques", http.StatusInternalServerError)
		fmt.Println("Erreur BDD GetForumStats :", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}