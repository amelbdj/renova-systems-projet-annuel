package main

import (
	"fmt"
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
	"upcycleconnect/bdd"
)

func Health(w http.ResponseWriter, r *http.Request) {
	err := bdd.Db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "ping à la bdd")
}

func main() {
	bdd.Db = bdd.NewDB()
	http.HandleFunc("GET /{$}", Health)
	http.HandleFunc("POST /admin/login", admin.Login)

	// Eviter les problèmes de CORS pour les requêtes préliminaires (OPTIONS)

	http.HandleFunc("OPTIONS /admin/users/delete/{id}", admin.DeletedUser)
	http.HandleFunc("OPTIONS /admin/users/add", admin.CreateUser)
	http.HandleFunc("OPTIONS /admin/users/modify/{id}", admin.UpdateUser)
	
	http.HandleFunc("OPTIONS /admin/categories/add", admin.CreateCategorie)
	http.HandleFunc("OPTIONS /admin/categories/delete/{id}", admin.DeleteCategorie)
	
	http.HandleFunc("OPTIONS /admin/annonces/validate/{id}", admin.ValidateAnnonce)
	http.HandleFunc("OPTIONS /admin/annonces/refuse/{id}", admin.RefuseAnnonce)
	
	http.HandleFunc("OPTIONS /admin/evenements/validate/{id}", admin.ValidateEvenement)
	http.HandleFunc("OPTIONS /admin/evenements/refuse/{id}", admin.RefuseEvenement)

	// Routes protégées par le middleware d'authentification
	// Note : Le middleware doit être appliqué à chaque route qui nécessite une authentification


	// Users
    http.HandleFunc("GET /admin/users", auth.VerifyTokenMiddleware(admin.GetAllUsers))
    http.HandleFunc("POST /admin/users/add", auth.VerifyTokenMiddleware(admin.CreateUser))
    http.HandleFunc("DELETE /admin/users/delete/{id}", auth.VerifyTokenMiddleware(admin.DeletedUser))
    http.HandleFunc("PUT /admin/users/modify/{id}", auth.VerifyTokenMiddleware(admin.UpdateUser))
    http.HandleFunc("GET /admin/users/role/{role}", auth.VerifyTokenMiddleware(admin.GetUserByRole))
    http.HandleFunc("GET /admin/users/search", auth.VerifyTokenMiddleware(admin.GetUserByName))

    // Categories
    http.HandleFunc("POST /admin/categories/add", auth.VerifyTokenMiddleware(admin.CreateCategorie))
    http.HandleFunc("DELETE /admin/categories/delete/{id}", auth.VerifyTokenMiddleware(admin.DeleteCategorie))
    http.HandleFunc("GET /admin/categories", auth.VerifyTokenMiddleware(admin.GetAllCategories))

    // Annonces
    http.HandleFunc("GET /admin/annonces", auth.VerifyTokenMiddleware(admin.GetAllAnnonces))
    http.HandleFunc("PUT /admin/annonces/validate/{id}", auth.VerifyTokenMiddleware(admin.ValidateAnnonce))
    http.HandleFunc("PUT /admin/annonces/refuse/{id}", auth.VerifyTokenMiddleware(admin.RefuseAnnonce))

    // Evenements
    http.HandleFunc("GET /admin/evenements", auth.VerifyTokenMiddleware(admin.GetAllEvenements))
    http.HandleFunc("PUT /admin/evenements/validate/{id}", auth.VerifyTokenMiddleware(admin.ValidateEvenement))
    http.HandleFunc("PUT /admin/evenements/refuse/{id}", auth.VerifyTokenMiddleware(admin.RefuseEvenement))
	

	fmt.Println("test de : http://localhost:8081")
	http.ListenAndServe(":8081", nil)

	
}