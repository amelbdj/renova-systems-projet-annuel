package admin

import (
	"encoding/json"
	"fmt"
	"upcycleconnect/bdd"

	"net/http"
	"upcycleconnect/models"
	// "strconv"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from GetAllUsers")

	users, err := bdd.GetUsers()

	if err != nil {
		http.Error(w, "erreur de récupération des utilisateurs", http.StatusInternalServerError)

		return
	}

	response, err := json.Marshal(users)

	if err != nil {
		http.Error(w, "erreur de conversion", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("METHOD :", r.Method)
fmt.Println("CONTENT TYPE :", r.Header.Get("Content-Type"))
fmt.Println("URL :", r.URL.Path)
	var UserDto models.User

	err := json.NewDecoder(r.Body).Decode(&UserDto)

	if err != nil {
		http.Error(w,
			"Impossible de décoder un modèle User au format json",
			http.StatusBadRequest)
		return

	}

	err = bdd.CreateUser(UserDto)
	if err != nil {
		http.Error(w,
			"erreur dans la création d'un utilisateur",
			http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
