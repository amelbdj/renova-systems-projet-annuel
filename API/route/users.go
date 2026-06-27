package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

func RoutesUsers() {

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
	http.HandleFunc("OPTIONS /api/user/profile", admin.UpdateProfileHandler)

	// --- Routes d'administration des comptes : réservées au personnel (Salarié / Administrateur) ---
	http.HandleFunc("GET /admin/users", auth.VerifyRoleMiddleware(admin.GetAllUsers, "Salarié", "Administrateur"))
	http.HandleFunc("POST /admin/users/add", auth.VerifyTokenMiddleware(admin.CreateUser))
	http.HandleFunc("DELETE /admin/users/delete/{id}", auth.VerifyRoleMiddleware(admin.DeletedUser, "Salarié", "Administrateur"))
	http.HandleFunc("PUT /admin/users/modify/{id}", auth.VerifyRoleMiddleware(admin.UpdateUser, "Salarié", "Administrateur"))
	http.HandleFunc("GET /admin/users/role/{role}", auth.VerifyRoleMiddleware(admin.GetUserByRole, "Salarié", "Administrateur"))
	http.HandleFunc("GET /admin/users/search", auth.VerifyRoleMiddleware(admin.GetUserByName, "Salarié", "Administrateur"))
	http.HandleFunc("PUT /admin/users/validate/{id}", auth.VerifyRoleMiddleware(admin.ValidateUser, "Salarié", "Administrateur"))
	http.HandleFunc("PUT /admin/users/refuse/{id}", auth.VerifyRoleMiddleware(admin.RefuseUser, "Salarié", "Administrateur"))
	http.HandleFunc("PUT /admin/users/ban/{id}", auth.VerifyRoleMiddleware(admin.BanUserHandler, "Salarié", "Administrateur"))

	// GET /admin/users/{id} : laissé accessible à tout utilisateur connecté
	// (l'app mobile s'en sert pour afficher SON propre profil)
	http.HandleFunc("GET /user/profile", admin.GetUserById)
	http.HandleFunc("GET /admin/users/{id}", auth.VerifyTokenMiddleware(admin.GetUserById))
	http.HandleFunc("GET /api/user/payment-history", admin.PaymentHistoryHandler)
	http.HandleFunc("POST /api/user/update-password", auth.VerifyTokenMiddleware(admin.UpdatePasswordHandler))
	http.HandleFunc("PUT /api/user/profile", auth.VerifyTokenMiddleware(admin.UpdateProfileHandler))

	http.HandleFunc("/api/user/ecostats", admin.GetEcoStatsHandler)
	http.HandleFunc("/api/user/stats", admin.GetEcoStatsHandler)
	http.HandleFunc("/api/user/achats", admin.GetMyPurchases)
}
