package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

func RoutesNotifications() {
	http.HandleFunc("OPTIONS /admin/notifications/send", admin.SendNotificationToAudience)
	http.HandleFunc("POST /admin/notifications/send", auth.VerifyTokenMiddleware(admin.SendNotificationToAudience))

	http.HandleFunc("OPTIONS /admin/notifications/user/{id}", admin.GetUserNotifications)
	http.HandleFunc("GET /admin/notifications/user/{id}", auth.VerifyTokenMiddleware(admin.GetUserNotifications))

	http.HandleFunc("OPTIONS /admin/notifications/user/{id}/read", admin.MarkNotificationsRead)
	http.HandleFunc("POST /admin/notifications/user/{id}/read", auth.VerifyTokenMiddleware(admin.MarkNotificationsRead))
}
