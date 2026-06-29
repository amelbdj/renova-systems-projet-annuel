package main

import (
	"fmt"
	"net/http"
	"upcycleconnect/bdd"
	"upcycleconnect/route"
)

func main() {
	bdd.Db = bdd.NewDB()

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
	route.RoutesNotifications()
	route.RoutesDocuments()

	fmt.Println("test de : http://localhost:8081")
	fmt.Println(http.ListenAndServe(":8081", nil))
}
