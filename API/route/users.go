package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

// RoutesUsers : toutes les routes liées aux utilisateurs (CRUD, validation, profil, mot de passe...)
func RoutesUsers() {
	// --- OPTIONS (pré-vol CORS) ---
	http.HandleFunc("OPTIONS /admin/users/delete/{id}", admin.DeletedUser)
	http.HandleFunc("OPTIONS /admin/users/add", admin.CreateUser)
	http.HandleFunc("OPTIONS /admin/users/modify/{id}", admin.UpdateUser)
	http.HandleFunc("OPTIONS /admin/users/validate/{id}", admin.ValidateUser)
	http.HandleFunc("OPTIONS /admin/users/refuse/{id}", admin.RefuseUser)
	http.HandleFunc("OPTIONS /admin/users", admin.GetAllUsers)
	http.HandleFunc("OPTIONS /admin/users/role/{role}", admin.GetUserByRole)
	http.HandleFunc("OPTIONS /admin/users/search", admin.GetUserByName)
	http.HandleFunc("OPTIONS /admin/users/{id}", admin.GetUserById)
	http.HandleFunc("OPTIONS /admin/users/ban/{id}", admin.BanUserHandler)
	http.HandleFunc("OPTIONS /user/profile", admin.GetUserById)
	http.HandleFunc("OPTIONS /api/user/payment-history", admin.PaymentHistoryHandler)
	http.HandleFunc("OPTIONS /api/user/update-password", admin.UpdatePasswordHandler)

	// --- Vraies routes ---
	http.HandleFunc("GET /admin/users", auth.VerifyTokenMiddleware(admin.GetAllUsers))
	http.HandleFunc("POST /admin/users/add", auth.VerifyTokenMiddleware(admin.CreateUser))
	http.HandleFunc("DELETE /admin/users/delete/{id}", auth.VerifyTokenMiddleware(admin.DeletedUser))
	http.HandleFunc("PUT /admin/users/modify/{id}", auth.VerifyTokenMiddleware(admin.UpdateUser))
	http.HandleFunc("GET /admin/users/role/{role}", auth.VerifyTokenMiddleware(admin.GetUserByRole))
	http.HandleFunc("GET /admin/users/search", auth.VerifyTokenMiddleware(admin.GetUserByName))
	http.HandleFunc("PUT /admin/users/validate/{id}", auth.VerifyTokenMiddleware(admin.ValidateUser))
	http.HandleFunc("PUT /admin/users/refuse/{id}", auth.VerifyTokenMiddleware(admin.RefuseUser))
	http.HandleFunc("GET /user/profile", admin.GetUserById)
	http.HandleFunc("GET /admin/users/{id}", auth.VerifyTokenMiddleware(admin.GetUserById))
	http.HandleFunc("PUT /admin/users/ban/{id}", auth.VerifyTokenMiddleware(admin.BanUserHandler))
	http.HandleFunc("GET /api/user/payment-history", admin.PaymentHistoryHandler)
	http.HandleFunc("POST /api/user/update-password", auth.VerifyTokenMiddleware(admin.UpdatePasswordHandler))

	// Stats éco & achats de l'utilisateur
	http.HandleFunc("/api/user/ecostats", admin.GetEcoStatsHandler)
	http.HandleFunc("/api/user/stats", admin.GetEcoStatsHandler)
	http.HandleFunc("/api/user/achats", admin.GetMyPurchases)
}
