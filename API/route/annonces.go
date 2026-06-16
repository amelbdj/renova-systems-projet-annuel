package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

// RoutesAnnonces : gestion des annonces (validation, vente, paiement mobile...)
func RoutesAnnonces() {
	// --- OPTIONS (pré-vol CORS) ---
	http.HandleFunc("OPTIONS /admin/annonces/validate/{id}", admin.ValidateAnnonce)
	http.HandleFunc("OPTIONS /admin/annonces", admin.GetAllAnnonces)
	http.HandleFunc("OPTIONS /admin/annonces/refuse/{id}", admin.RefuseAnnonce)
	http.HandleFunc("OPTIONS /admin/annonces/add", admin.CreateAnnonce)
	http.HandleFunc("OPTIONS /admin/annonces/delete/{id}", admin.DeleteAnnonce)
	http.HandleFunc("OPTIONS /admin/annonces/modify/{id}", admin.UpdateAnnonce)
	http.HandleFunc("OPTIONS /api/annonces/vendre", admin.ConfirmPaymentAndOrder)
	http.HandleFunc("OPTIONS /api/mobile/payment-intent", admin.PaymentIntentMobile)

	// --- Vraies routes ---
	http.HandleFunc("GET /admin/annonces", admin.GetAllAnnonces)
	http.HandleFunc("PUT /admin/annonces/validate/{id}", auth.VerifyTokenMiddleware(admin.ValidateAnnonce))
	http.HandleFunc("PUT /admin/annonces/refuse/{id}", auth.VerifyTokenMiddleware(admin.RefuseAnnonce))
	http.HandleFunc("POST /admin/annonces/add", auth.VerifyTokenMiddleware(admin.CreateAnnonce))
	http.HandleFunc("DELETE /admin/annonces/delete/{id}", auth.VerifyTokenMiddleware(admin.DeleteAnnonce))
	http.HandleFunc("PUT /admin/annonces/modify/{id}", auth.VerifyTokenMiddleware(admin.UpdateAnnonce))
	http.HandleFunc("GET /admin/annonces/search", auth.VerifyTokenMiddleware(admin.GetAnnonceByTitle))
	http.HandleFunc("GET /mes-annonces", admin.GetMyAnnonces)
	http.HandleFunc("GET /api/annonces/all", admin.GetValidatedAnnonces)
	http.HandleFunc("GET /api/annonces", admin.GetOneAnnonce)
	http.HandleFunc("POST /api/annonces/vendre", admin.ConfirmPaymentAndOrder)
	http.HandleFunc("POST /api/mobile/payment-intent", admin.PaymentIntentMobile) // android payment intent
}
