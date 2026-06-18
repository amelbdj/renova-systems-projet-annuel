package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

func RoutesPro() {

	http.HandleFunc("OPTIONS /api/pro/upgrade", admin.UpgradeToPremiumHandler)
	http.HandleFunc("OPTIONS /api/pro/portal", admin.CustomerPortalHandler)
	http.HandleFunc("OPTIONS /api/pro/sync", admin.SyncPremiumStatusHandler)
	http.HandleFunc("OPTIONS /api/pro/subscribe", admin.CreateProSubscriptionHandler)

	http.HandleFunc("POST /api/pro/upgrade", admin.UpgradeToPremiumHandler)
	http.HandleFunc("POST /api/pro/portal", admin.CustomerPortalHandler)
	http.HandleFunc("GET /api/pro/sync", admin.SyncPremiumStatusHandler)
	http.HandleFunc("POST /api/pro/subscribe", auth.VerifyTokenMiddleware(admin.CreateProSubscriptionHandler))
}
