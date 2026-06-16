package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

// RoutesArticles : gestion des articles / news
func RoutesArticles() {
	// --- OPTIONS (pré-vol CORS) ---
	http.HandleFunc("OPTIONS /admin/articles/validate/{id}", admin.ValidateArticle)
	http.HandleFunc("OPTIONS /admin/articles/refuse/{id}", admin.RefuseArticle)
	http.HandleFunc("OPTIONS /admin/articles/delete/{id}", admin.DeleteArticle)
	http.HandleFunc("OPTIONS /admin/articles/add/{action}", admin.CreateArticle)
	http.HandleFunc("OPTIONS /admin/articles/modify/{id}/{action}", admin.ModifyArticle)
	http.HandleFunc("OPTIONS /admin/articles/salarie/{id}", admin.GetArticlesBySalarie)
	http.HandleFunc("OPTIONS /admin/articles/{id}", admin.GetArticleById)
	http.HandleFunc("OPTIONS /admin/articles", admin.GetAllArticles)

	// --- Vraies routes ---
	http.HandleFunc("GET /admin/articles", admin.GetAllArticles)
	http.HandleFunc("GET /admin/articles/salarie/{id}", auth.VerifyTokenMiddleware(admin.GetArticlesBySalarie))
	http.HandleFunc("GET /admin/articles/{id}", admin.GetArticleById)
	http.HandleFunc("PUT /admin/articles/validate/{id}", auth.VerifyTokenMiddleware(admin.ValidateArticle))
	http.HandleFunc("PUT /admin/articles/refuse/{id}", auth.VerifyTokenMiddleware(admin.RefuseArticle))
	http.HandleFunc("DELETE /admin/articles/delete/{id}", auth.VerifyTokenMiddleware(admin.DeleteArticle))
	http.HandleFunc("POST /admin/articles/add/{action}", auth.VerifyTokenMiddleware(admin.CreateArticle))
	http.HandleFunc("PUT /admin/articles/modify/{id}/{action}", auth.VerifyTokenMiddleware(admin.ModifyArticle))
}
