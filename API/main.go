package main

import (
	"fmt"
	"net/http"
	"upcycleconnect/bdd"
	"upcycleconnect/route"
)

func main() {
	bdd.Db = bdd.NewDB()

	// On enregistre toutes les routes, classées par thématique dans le dossier route/
	route.RoutesAuth()
	route.RoutesUsers()
	route.RoutesCategories()
	route.RoutesAnnonces()
	route.RoutesEvenements()
	route.RoutesArticles()
	route.RoutesLogistique()
	route.RoutesStripe()
	route.RoutesTraductions()
	route.RoutesForum()
	route.RoutesChat()
	route.RoutesPlanning()
	route.RoutesFinance()
	route.RoutesDivers()
	route.RoutesPro()

	fmt.Println("test de : http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
