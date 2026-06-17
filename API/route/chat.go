package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

// RoutesChat : messagerie (historique, conversations, websocket)
func RoutesChat() {
	// --- OPTIONS (pré-vol CORS) ---
	http.HandleFunc("OPTIONS /api/chat/history", admin.GetChatHistoryHandler)
	http.HandleFunc("OPTIONS /api/chat/conversations", admin.GetConversationsHandler)

	// --- Vraies routes ---
	http.HandleFunc("GET /api/chat/history", auth.VerifyTokenMiddleware(admin.GetChatHistoryHandler))
	http.HandleFunc("GET /api/chat/conversations", auth.VerifyTokenMiddleware(admin.GetConversationsHandler))
	http.HandleFunc("GET /ws/chat", admin.ChatHandler)
}
