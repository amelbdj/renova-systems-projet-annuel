package admin

import (
	"encoding/json"
	"fmt"

	"upcycleconnect/bdd"

	"net/http"
	"strconv"
	"upcycleconnect/models"
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

	var user models.User
err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w,
			"Impossible de décoder un modèle User au format json",
			http.StatusBadRequest)
		return

	}
}

func DeletedUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}

	err = bdd.DeletedUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "utilisateur suppr")
}
