package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

func RoutesNotifications() {
	http.HandleFunc("OPTIONS /admin/notifications/send", admin.SendNotificationToAudience)
	http.HandleFunc("POST /admin/notifications/send", auth.VerifyTokenMiddleware(admin.SendNotificationToAudience))
}
