package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"upcycleconnect/bdd"
	"upcycleconnect/models"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	fmt.Println("hello from GetAllCategories")

	Categories, err := bdd.GetCategories()

	if err != nil {
		http.Error(w, "erreur de récupération des catégories", http.StatusInternalServerError)

		return
	}

	response, err := json.Marshal(Categories)

	if err != nil {
		http.Error(w, "erreur de conversion", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", response)
}

func CreateCategorie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	fmt.Println("hello from CreateCategorie")
	
	if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return 
    }
	var CategorieDto models.Categorie

	err := json.NewDecoder(r.Body).Decode(&CategorieDto)

	if err != nil {
		http.Error(w,
			"Impossible de décoder un modèle Categorie au format json",
			http.StatusBadRequest)
		return

	}

	err = bdd.CreateCategorie(CategorieDto)
	if err != nil {
		http.Error(w,
			"erreur dans la création d'une catégorie",
			http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteCategorie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	fmt.Println("hello from DeleteCategorie")	
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}	
idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}

	err = bdd.DeleteCategorie(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "catégorie suppr")
}