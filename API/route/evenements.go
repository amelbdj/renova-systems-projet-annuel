package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

// RoutesEvenements : gestion des événements (validation, inscription, checkout...)
func RoutesEvenements() {
	// --- OPTIONS (pré-vol CORS) ---
	http.HandleFunc("OPTIONS /admin/evenements/validate/{id}", admin.ValidateEvenement)
	http.HandleFunc("OPTIONS /admin/evenements/update/{id}", admin.UpdateEvenement)
	http.HandleFunc("OPTIONS /admin/evenements/delete/{id}", admin.DeleteEvenement)
	http.HandleFunc("OPTIONS /admin/evenements/refuse/{id}", admin.RefuseEvenement)
	http.HandleFunc("OPTIONS /admin/evenements/add", admin.CreateEvenement)
	http.HandleFunc("OPTIONS /admin/evenements", admin.GetAllEvenements)
	http.HandleFunc("OPTIONS /admin/evenements/inscription", admin.InscrireClient)
	http.HandleFunc("OPTIONS /admin/evenements/desinscription", admin.DesinscriptionHandler)
	http.HandleFunc("OPTIONS /api/web/checkout/evenement", admin.CreateEventCheckoutSession)

	// --- Vraies routes ---
	http.HandleFunc("GET /admin/evenements", auth.VerifyTokenMiddleware(admin.GetAllEvenements))
	http.HandleFunc("POST /admin/evenements/add", auth.VerifyTokenMiddleware(admin.CreateEvenement))
	http.HandleFunc("PUT /admin/evenements/{id}", auth.VerifyTokenMiddleware(admin.UpdateEvenement))
	http.HandleFunc("DELETE /admin/evenements/delete/{id}", auth.VerifyTokenMiddleware(admin.DeleteEvenement))
	http.HandleFunc("PUT /admin/evenements/validate/{id}", auth.VerifyTokenMiddleware(admin.ValidateEvenement))
	http.HandleFunc("PUT /admin/evenements/refuse/{id}", auth.VerifyTokenMiddleware(admin.RefuseEvenement))
	http.HandleFunc("POST /admin/evenements/inscription", admin.InscrireClient)
	http.HandleFunc("POST /admin/evenements/desinscription", admin.DesinscriptionHandler)
	http.HandleFunc("POST /api/web/checkout/evenement", admin.CreateEventCheckoutSession)
}
