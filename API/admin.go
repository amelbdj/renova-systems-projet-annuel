package main

import (
	"fmt"
	"net/http"
	"upcycleconnect/admin"
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

	http.HandleFunc("GET /admin/users", admin.GetAllUsers)
	http.HandleFunc("POST /admin/users/add", admin.CreateUser)
	http.HandleFunc("OPTIONS /admin/users/delete/{id}", admin.DeletedUser)
	http.HandleFunc("DELETE /admin/users/delete/{id}", admin.DeletedUser)
	http.HandleFunc("OPTIONS /admin/users/add", admin.CreateUser)
	http.HandleFunc("PUT /admin/users/modify/{id}", admin.UpdateUser)
	http.HandleFunc("OPTIONS /admin/users/modify/{id}", admin.UpdateUser)
	http.HandleFunc("GET /admin/users/role/{role}", admin.GetUserByRole)
	http.HandleFunc("GET /admin/users/search", admin.GetUserByName)



	http.HandleFunc("POST /admin/categories/add", admin.CreateCategorie)
	http.HandleFunc("OPTIONS /admin/categories/add", admin.CreateCategorie)
	http.HandleFunc("GET /admin/categories", admin.GetAllCategories)

	http.HandleFunc("GET /admin/annonces", admin.GetAllAnnonces)
	http.HandleFunc("PUT /admin/annonces/validate/{id}", admin.ValidateAnnonce)
	http.HandleFunc("OPTIONS /admin/annonces/validate/{id}", admin.ValidateAnnonce)
	http.HandleFunc("OPTIONS /admin/annonces/refuse/{id}", admin.RefuseAnnonce)
	http.HandleFunc("PUT /admin/annonces/refuse/{id}", admin.RefuseAnnonce)

	http.HandleFunc("GET /admin/evenements", admin.GetAllEvenements)
	http.HandleFunc("PUT /admin/evenements/validate/{id}", admin.ValidateEvenement)
	http.HandleFunc("OPTIONS /admin/evenements/validate/{id}", admin.ValidateEvenement)
	http.HandleFunc("OPTIONS /admin/evenements/refuse/{id}", admin.RefuseEvenement)
	http.HandleFunc("PUT /admin/evenements/refuse/{id}", admin.RefuseEvenement)




	fmt.Println("test de : http://localhost:8081")
	http.ListenAndServe(":8081", nil)

	
}