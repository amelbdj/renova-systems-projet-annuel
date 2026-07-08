package admin

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"upcycleconnect/bdd"
	"upcycleconnect/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	clients   = make(map[int]*websocket.Conn)
	clientsMu sync.Mutex
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	userIdStr := r.URL.Query().Get("userId")
	userId, _ := strconv.Atoi(userIdStr)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erreur Upgrade:", err)
		return
	}
	defer conn.Close()

	clientsMu.Lock()
	clients[userId] = conn
	clientsMu.Unlock()

	log.Printf("Utilisateur %d est connecté au chat", userId)

	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {

			clientsMu.Lock()
			delete(clients, userId)
			clientsMu.Unlock()
			break
		}

		err = bdd.SaveMessage(msg)
		if err != nil {
			log.Println("Erreur BDD Save:", err)
		}

		clientsMu.Lock()
		destConn, online := clients[msg.DestinataireID]
		clientsMu.Unlock()

		if online {
			destConn.WriteJSON(msg)
		}
	}
}

func GetChatHistoryHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	annonceId, _ := strconv.Atoi(r.URL.Query().Get("annonce_id"))
	user1, _ := strconv.Atoi(r.URL.Query().Get("user1"))
	user2, _ := strconv.Atoi(r.URL.Query().Get("user2"))

	if annonceId == 0 || user1 == 0 || user2 == 0 {
		http.Error(w, "Paramètres manquants", http.StatusBadRequest)
		return
	}

	messages, err := bdd.GetConversation(annonceId, user1, user2)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de l'historique", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var msg models.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Message invalide", http.StatusBadRequest)
		return
	}

	if msg.AnnonceID == 0 || msg.ExpediteurID == 0 || msg.DestinataireID == 0 || msg.Contenu == "" {
		http.Error(w, "Informations manquantes", http.StatusBadRequest)
		return
	}

	err = bdd.SaveMessage(msg)
	if err != nil {
		log.Println("Erreur SaveMessage:", err)
		http.Error(w, "Erreur sauvegarde message", http.StatusInternalServerError)
		return
	}

	clientsMu.Lock()
	destConn, online := clients[msg.DestinataireID]
	clientsMu.Unlock()

	if online {
		destConn.WriteJSON(msg)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Message envoye"})
}

func GetConversationsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userIdStr := r.URL.Query().Get("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil || userId == 0 {
		http.Error(w, "ID utilisateur invalide", http.StatusBadRequest)
		return
	}

	conversations, err := bdd.GetUserConversations(userId)
	if err != nil {
		log.Println("Erreur GetUserConversations:", err)
		http.Error(w, "Erreur lors de la récupération des conversations", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(conversations)
}
