package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

// RoutesStripe : connexion Stripe, webhook et paiement d'annonce
func RoutesStripe() {
	// --- OPTIONS (pré-vol CORS) ---
	http.HandleFunc("OPTIONS /admin/connect-stripe", admin.ConnectToStripe)
	http.HandleFunc("OPTIONS /api/payment-annonce", admin.PaymentAnnonce)

	// --- Vraies routes ---
	http.HandleFunc("POST /admin/connect-stripe", auth.VerifyTokenMiddleware(admin.ConnectToStripe))
	http.HandleFunc("POST /api/stripe/webhook", admin.StripeWebhookHandler)
	http.HandleFunc("POST /api/payment-annonce", admin.PaymentAnnonce)
}
