package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

// RoutesPro : abonnements Premium / Pro (routes ajoutées par Faty)
func RoutesPro() {
	// --- OPTIONS (pré-vol CORS) ---
	http.HandleFunc("OPTIONS /api/pro/upgrade", admin.UpgradeToPremiumHandler)
	http.HandleFunc("OPTIONS /api/pro/portal", admin.CustomerPortalHandler)
	http.HandleFunc("OPTIONS /api/pro/sync", admin.SyncPremiumStatusHandler)
	http.HandleFunc("OPTIONS /api/pro/subscribe", admin.CreateProSubscriptionHandler)

	// --- Vraies routes ---
	http.HandleFunc("POST /api/pro/upgrade", admin.UpgradeToPremiumHandler)
	http.HandleFunc("POST /api/pro/portal", admin.CustomerPortalHandler)
	http.HandleFunc("GET /api/pro/sync", admin.SyncPremiumStatusHandler)
	http.HandleFunc("POST /api/pro/subscribe", auth.VerifyTokenMiddleware(admin.CreateProSubscriptionHandler))
}
