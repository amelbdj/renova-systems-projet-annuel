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
	http.HandleFunc("POST /admin/categories/add", admin.CreateCategorie)
	http.HandleFunc("GET /admin/categories", admin.GetAllCategories)

	fmt.Println("test de : http://localhost:8081")
	http.ListenAndServe(":8081", nil)

	
}