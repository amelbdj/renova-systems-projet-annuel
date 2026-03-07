package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"upcycleconnect/bdd"
	// "upcycleconnect/models"
	// "strconv"
)

func GetAllAnnonces(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from GetAllAnnonces")

	Annonces, err := bdd.GetAnnonces()

	if err != nil {
		http.Error(w, "erreur de récupération des annonces", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}

	response, err := json.Marshal(Annonces)

	if err != nil {
		http.Error(w, "erreur de conversion", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", response)
}

