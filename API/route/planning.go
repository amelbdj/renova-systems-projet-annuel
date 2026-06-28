package route

import (
	"net/http"
	"upcycleconnect/admin"
)

func RoutesPlanning() {
	http.HandleFunc("OPTIONS /user/planning", admin.GetUserPlanningHandler)
	http.HandleFunc("GET /user/planning", admin.GetUserPlanningHandler)
}
