package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

func RoutesForum() {

	http.HandleFunc("OPTIONS /admin/forum/messages", admin.GetForumMessages)
	http.HandleFunc("OPTIONS /admin/forum/messages/moderate/{id}", admin.ModerateForumMessage)
	http.HandleFunc("OPTIONS /admin/forum/stats", admin.GetForumStats)
	http.HandleFunc("OPTIONS /user/forums", admin.GetForumsHandler)
	http.HandleFunc("OPTIONS /user/forums/messages", admin.ForumClientMessagesHandler)

	http.HandleFunc("GET /admin/forum/messages", auth.VerifyTokenMiddleware(admin.GetForumMessages))
	http.HandleFunc("PUT /admin/forum/messages/moderate/{id}", auth.VerifyTokenMiddleware(admin.ModerateForumMessage))
	http.HandleFunc("GET /admin/forum/stats", auth.VerifyTokenMiddleware(admin.GetForumStats))

	http.HandleFunc("GET /user/forums", admin.GetForumsHandler)
	http.HandleFunc("POST /user/forums", admin.GetForumsHandler)
	http.HandleFunc("GET /user/forums/messages", admin.ForumClientMessagesHandler)
	http.HandleFunc("POST /user/forums/messages", admin.ForumClientMessagesHandler)
}
