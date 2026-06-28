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
	http.HandleFunc("OPTIONS /api/pro/cancel", admin.CancelSubscriptionHandler)
	http.HandleFunc("OPTIONS /api/pro/annonces/sponsor", admin.ToggleSponsorHandler)
	http.HandleFunc("OPTIONS /api/pro/projets", admin.GetProjetsHandler)
	http.HandleFunc("OPTIONS /api/pro/projets/create", admin.CreateProjetHandler)
	http.HandleFunc("OPTIONS /api/pro/projets/delete", admin.DeleteProjetHandler)
	http.HandleFunc("OPTIONS /api/pro/projets/update", admin.UpdateProjetHandler)
	http.HandleFunc("OPTIONS /api/pro/etapes/create", admin.CreateEtapeHandler)
	http.HandleFunc("OPTIONS /api/pro/etapes", admin.GetEtapesHandler)
	http.HandleFunc("OPTIONS /api/pro/etapes/delete", admin.DeleteEtapeHandler)
	http.HandleFunc("OPTIONS /api/pro/etapes/statut", admin.UpdateEtapeStatutHandler)

	http.HandleFunc("POST /api/pro/upgrade", admin.UpgradeToPremiumHandler)
	http.HandleFunc("POST /api/pro/portal", admin.CustomerPortalHandler)
	http.HandleFunc("GET /api/pro/sync", admin.SyncPremiumStatusHandler)
	http.HandleFunc("POST /api/pro/subscribe", auth.VerifyTokenMiddleware(admin.CreateProSubscriptionHandler))
	http.HandleFunc("POST /api/pro/cancel", admin.CancelSubscriptionHandler)
	http.HandleFunc("POST /api/pro/annonces/sponsor", auth.VerifyTokenMiddleware(admin.ToggleSponsorHandler))
	http.HandleFunc("GET /api/pro/projets", auth.VerifyTokenMiddleware(admin.GetProjetsHandler))
	http.HandleFunc("POST /api/pro/projets/create", auth.VerifyTokenMiddleware(admin.CreateProjetHandler))
	http.HandleFunc("DELETE /api/pro/projets/delete", auth.VerifyTokenMiddleware(admin.DeleteProjetHandler))
	http.HandleFunc("PUT /api/pro/projets/update", auth.VerifyTokenMiddleware(admin.UpdateProjetHandler))
	http.HandleFunc("POST /api/pro/etapes/create", auth.VerifyTokenMiddleware(admin.CreateEtapeHandler))
	http.HandleFunc("GET /api/pro/etapes", auth.VerifyTokenMiddleware(admin.GetEtapesHandler))
	http.HandleFunc("DELETE /api/pro/etapes/delete", auth.VerifyTokenMiddleware(admin.DeleteEtapeHandler))
	http.HandleFunc("PUT /api/pro/etapes/statut", auth.VerifyTokenMiddleware(admin.UpdateEtapeStatutHandler))
}
