package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"upcycleconnect/bdd"
	"upcycleconnect/models"
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

func GetForumsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "GET" {
		topics, err := bdd.GetAllTopics()
		if err != nil {
			http.Error(w, "Erreur serveur", http.StatusInternalServerError)
			return
		}

		if topics == nil {
			topics = []models.ForumTopic{}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(topics)
		return
	}

	if r.Method == "POST" {
		var req struct {
			IdUser  int    `json:"id_user"`
			Titre   string `json:"titre"`
			Message string `json:"message"`
		}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || req.Titre == "" || req.Message == "" {
			http.Error(w, "Données invalides", http.StatusBadRequest)
			return
		}

		err = bdd.CreerNouveauSujet(req.IdUser, req.Titre, req.Message)
		if err != nil {
			http.Error(w, "Erreur lors de la création", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}
}

func ForumClientMessagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "GET" {
		topicStr := r.URL.Query().Get("topic_id")
		topicId, err := strconv.Atoi(topicStr)
		if err != nil {
			http.Error(w, "Id de topic invalide ou manquant", http.StatusBadRequest)
			return
		}

		messages, err := bdd.GetMessagesByTopicClient(topicId)
		if err != nil {
			fmt.Println("Erreur GET messages forum client:", err)
			http.Error(w, "Erreur serveur lors de la récupération", http.StatusInternalServerError)
			return
		}

		if messages == nil {
			messages = []models.MessageForum{}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages)
		return
	}

	if r.Method == "POST" {
		var req models.MessageForum
		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil || req.Contenu == "" || req.IdTopic == 0 || req.IdUser == 0 {
			http.Error(w, "Données invalides ou message vide", http.StatusBadRequest)
			return
		}

		err = bdd.AjouterMessageForum(req.IdTopic, req.IdUser, req.Contenu)
		if err != nil {
			fmt.Println("Erreur POST nouveau message forum:", err)
			http.Error(w, "Erreur lors de l'enregistrement", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
}
