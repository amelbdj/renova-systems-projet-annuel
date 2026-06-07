package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"upcycleconnect/bdd"
)

const (
	OneSignalAppID  = "79a53223-420a-46c8-83d9-1ca162fcb64f"
	OneSignalAPIKey = "os_v2_app_pgstei2cbjdmra6zdsqwf7fwj7lecom3g6lu4pumdxlt4rvgw66selidc5gwe5r2gpo7pr7cdhfecc55xgdksr5rplozsejaaghkzsa"
)


func SendPushNotification(userID string, message string) {
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
    
	// LA FAMEUSE LIGNE QUI VA NOUS DONNER LA RÉPONSE DE ONESIGNAL :
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