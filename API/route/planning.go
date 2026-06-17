package route

import (
	"net/http"
	"upcycleconnect/admin"
)

// RoutesPlanning : planning de l'utilisateur
func RoutesPlanning() {
	http.HandleFunc("OPTIONS /user/planning", admin.GetUserPlanningHandler)
	http.HandleFunc("GET /user/planning", admin.GetUserPlanningHandler)
}
