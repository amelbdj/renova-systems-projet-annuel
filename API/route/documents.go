package route

import (
	"net/http"
	"upcycleconnect/admin"
)

func RoutesDocuments() {
	http.HandleFunc("OPTIONS /admin/documents", admin.GetAllDocumentsHandler)
	http.HandleFunc("GET /admin/documents", admin.GetAllDocumentsHandler)
}
