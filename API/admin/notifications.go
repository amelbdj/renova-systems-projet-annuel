package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"upcycleconnect/bdd"
)

func SendNotificationToAudience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req struct {
		Cible   string `json:"cible"`
		Message string `json:"message"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, `{"erreur":"données invalides"}`, http.StatusBadRequest)
		return
	}

	if req.Message == "" {
		http.Error(w, `{"erreur":"le message est vide"}`, http.StatusBadRequest)
		return
	}

	var ids []string
	if req.Cible == "particuliers" {
		ids, _ = bdd.GetParticulierIDs()
	} else if req.Cible == "pros" {
		ids, _ = bdd.GetProfessionalIDs()
	} else {
		particuliers, _ := bdd.GetParticulierIDs()
		pros, _ := bdd.GetProfessionalIDs()
		ids = append(particuliers, pros...)
	}

	for _, id := range ids {
		// SendPushNotification cree aussi la notif "cloche" en base.
		go SendPushNotification(id, req.Message)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Notification envoyée",
		"nombre":  len(ids),
	})
}

func GetUserNotifications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
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

	notifs, err := bdd.GetNotifications(id)
	if err != nil {
		http.Error(w, "erreur de récupération des notifications", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}
	if notifs == nil {
		notifs = []map[string]interface{}{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifs)
}

func MarkNotificationsRead(w http.ResponseWriter, r *http.Request) {
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

	err = bdd.MarkNotificationsRead(id)
	if err != nil {
		http.Error(w, "erreur de mise à jour", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

var (
	OneSignalAppID  = envOr("ONESIGNAL_APP_ID", "79a53223-420a-46c8-83d9-1ca162fcb64f")
	OneSignalAPIKey = envOr("ONESIGNAL_API_KEY", "os_v2_app_pgstei2cbjdmra6zdsqwf7fwj7lecom3g6lu4pumdxlt4rvgw66selidc5gwe5r2gpo7pr7cdhfecc55xgdksr5rplozsejaaghkzsa")
)

func SendPushNotification(userID string, message string) {
	// Notification "cloche" en base (en plus du push OneSignal), pour que
	// l'utilisateur la retrouve dans l'app meme sans push / hors HTTPS.
	if idInt, err := strconv.Atoi(userID); err == nil {
		bdd.CreateNotification(idInt, message)
	}

	payload := map[string]interface{}{
		"app_id": OneSignalAppID,
		"include_aliases": map[string][]string{
			"external_id": {userID},
		},
		"target_channel": "push",
		"contents": map[string]string{
			"en": message,
			"fr": message,
		},
		"headings": map[string]string{"fr": "Upcycle Connect"},
	}

	jsonData, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erreur création requête OneSignal:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Key "+OneSignalAPIKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erreur envoi OneSignal:", err)
		return
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	fmt.Println("📡 RÉPONSE ONESIGNAL :", buf.String())
}

func NotifyAllAdmins(message string) {
	adminIDs, err := bdd.GetAllAdminIDs()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des admins:", err)
		return
	}
	fmt.Printf("🔍 Admins trouvés: %v. Envoi de la notification...\n", adminIDs)

	for _, id := range adminIDs {
		go SendPushNotification(id, message)
	}
}

func NotifyAllPros(message string) {
	proIDs, err := bdd.GetProfessionalIDs()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des professionnels:", err)
		return
	}
	fmt.Printf("🔍 Professionnels trouvés: %v. Envoi de la notification...\n", proIDs)

	for _, id := range proIDs {
		go SendPushNotification(id, message)
	}
}
