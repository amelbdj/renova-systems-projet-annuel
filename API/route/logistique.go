package route

import (
	"net/http"
	"upcycleconnect/admin"
)

// RoutesLogistique : box, conteneurs et retraits (pickups)
func RoutesLogistique() {
	// --- OPTIONS (pré-vol CORS) ---
	http.HandleFunc("OPTIONS /api/user/boxes", admin.GetMyBoxes)
	http.HandleFunc("OPTIONS /api/admin/conteneur/{id}/boxes", admin.GetBoxesForConteneurHandler)
	http.HandleFunc("OPTIONS /api/admin/box/add", admin.AddSingleBoxHandler)
	http.HandleFunc("OPTIONS /api/admin/box/update", admin.UpdateBoxStatusHandler)
	http.HandleFunc("OPTIONS /api/box/reserve", admin.ReserveBox)
	http.HandleFunc("OPTIONS /api/box/deposit", admin.ConfirmDeposit)
	http.HandleFunc("OPTIONS /api/box/collect", admin.CollectObject)
	http.HandleFunc("OPTIONS /api/admin/conteneurs", admin.GetConteneursAdmin)
	http.HandleFunc("OPTIONS /api/admin/conteneur/create", admin.CreateConteneur)
	http.HandleFunc("OPTIONS /api/user/pickups/{id}", admin.GetUserPickupsHandler)

	// --- Vraies routes ---
	http.HandleFunc("POST /api/box/reserve", admin.ReserveBox)
	http.HandleFunc("POST /api/box/deposit", admin.ConfirmDeposit)
	http.HandleFunc("POST /api/box/collect", admin.CollectObject)
	http.HandleFunc("GET /api/admin/conteneurs", admin.GetConteneursAdmin)
	http.HandleFunc("POST /api/admin/conteneur/create", admin.CreateConteneur)
	http.HandleFunc("POST /api/admin/box/add", admin.AddSingleBoxHandler)
	http.HandleFunc("PUT /api/admin/box/update", admin.UpdateBoxStatusHandler)
	http.HandleFunc("GET /api/user/boxes", admin.GetMyBoxes)
	http.HandleFunc("GET /api/user/pickups/{id}", admin.GetUserPickupsHandler)
	// Le {id} entre accolades indique à Go que cette partie de l'URL est une variable dynamique !
	http.HandleFunc("GET /api/admin/conteneur/{id}/boxes", admin.GetBoxesForConteneurHandler)
}
