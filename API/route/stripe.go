package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

func RoutesStripe() {

	http.HandleFunc("OPTIONS /admin/connect-stripe", admin.ConnectToStripe)
	http.HandleFunc("OPTIONS /api/payment-annonce", admin.PaymentAnnonce)

	http.HandleFunc("POST /admin/connect-stripe", auth.VerifyTokenMiddleware(admin.ConnectToStripe))
	http.HandleFunc("POST /api/stripe/webhook", admin.StripeWebhookHandler)
	http.HandleFunc("POST /api/payment-annonce", admin.PaymentAnnonce)
}
