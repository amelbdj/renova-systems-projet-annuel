package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

// RoutesCategories : gestion des catégories
func RoutesCategories() {
	// --- OPTIONS (pré-vol CORS) ---
	http.HandleFunc("OPTIONS /admin/categories/add", admin.CreateCategorie)
	http.HandleFunc("OPTIONS /admin/categories/delete/{id}", admin.DeleteCategorie)
	http.HandleFunc("OPTIONS /admin/categories", admin.GetAllCategories)

	// --- Vraies routes ---
	http.HandleFunc("POST /admin/categories/add", auth.VerifyTokenMiddleware(admin.CreateCategorie))
	http.HandleFunc("DELETE /admin/categories/delete/{id}", auth.VerifyTokenMiddleware(admin.DeleteCategorie))
	http.HandleFunc("GET /admin/categories", auth.VerifyTokenMiddleware(admin.GetAllCategories))
}
