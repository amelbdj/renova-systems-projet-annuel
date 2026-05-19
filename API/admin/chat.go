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

// Configuration du passage de HTTP à WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // À affiner plus tard pour la sécurité de l'infra
	},
}

// Gestion des clients connectés
var (
	clients   = make(map[int]*websocket.Conn) // Map UserId -> Connexion
	clientsMu sync.Mutex                     // Protection pour les accès concurrents
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Accepte les requêtes de n'importe où
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Autorise le token !
		userIdStr := r.URL.Query().Get("userId")
	userId, _ := strconv.Atoi(userIdStr)

	// 2. "Upgrade" la connexion en WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erreur Upgrade:", err)
		return
	}
	defer conn.Close()

	// 3. On enregistre le client comme étant en ligne
	clientsMu.Lock()
	clients[userId] = conn
	clientsMu.Unlock()

	log.Printf("Utilisateur %d est connecté au chat", userId)

	// 4. Boucle de lecture des messages
	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			// Si erreur (déconnexion), on nettoie
			clientsMu.Lock()
			delete(clients, userId)
			clientsMu.Unlock()
			break
		}

		// 5. Sauvegarde en BDD (via notre fichier bdd/chat.go)
		err = bdd.SaveMessage(msg)
		if err != nil {
			log.Println("Erreur BDD Save:", err)
		}

		// 6. Envoi au destinataire s'il est en ligne
		clientsMu.Lock()
		destConn, online := clients[msg.DestinataireID]
		clientsMu.Unlock()

		if online {
			destConn.WriteJSON(msg)
		}
	}
}

func GetChatHistoryHandler(w http.ResponseWriter, r *http.Request) {
	// --- 1. LES HEADERS CORS QUI RÉSOLVENT TON ERREUR ---
	w.Header().Set("Access-Control-Allow-Origin", "*") // Accepte les requêtes de n'importe où
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Autorise le token !

	// Si c'est juste la question de sécurité du navigateur (OPTIONS), on s'arrête là et on dit OK
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// --- 2. LE RESTE DU CODE (On ne change rien ici) ---
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
// GetConversationsHandler gère la requête pour lister les contacts d'un utilisateur
func GetConversationsHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Headers CORS pour autoriser le passage du token et des requêtes
	w.Header().Set("Access-Control-Allow-Origin", "*")
w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// 2. Récupération de l'ID utilisateur depuis les paramètres
	userIdStr := r.URL.Query().Get("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil || userId == 0 {
		http.Error(w, "ID utilisateur invalide", http.StatusBadRequest)
		return
	}

	// 3. Appel à la fonction BDD (celle que je t'ai donnée juste avant dans bdd/chat.go)
	conversations, err := bdd.GetUserConversations(userId)
	if err != nil {
		log.Println("Erreur GetUserConversations:", err)
		http.Error(w, "Erreur lors de la récupération des conversations", http.StatusInternalServerError)
		return
	}

	// 4. Envoi de la réponse en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(conversations)
}