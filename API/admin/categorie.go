package admin

import (
	"encoding/json"
	"fmt"
	"upcycleconnect/bdd"

	"net/http"
	"upcycleconnect/models"
	// "strconv"
)

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("METHOD :", r.Method)
fmt.Println("CONTENT TYPE :", r.Header.Get("Content-Type"))
fmt.Println("URL :", r.URL.Path)
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