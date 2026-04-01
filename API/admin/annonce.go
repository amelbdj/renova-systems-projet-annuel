package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"upcycleconnect/bdd"

	// "upcycleconnect/models"
	"strconv"
)

func GetAllAnnonces(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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

func ValidateAnnonce(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return 
    }

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}	
	err = bdd.ValidateAnnonce(id)

	if err != nil {
		http.Error(w, "erreur de validation de l'annonce", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Annonce validée avec succès")
}

func RefuseAnnonce(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return 
    }

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}	
	err = bdd.RefuseAnnonce(id)

	if err != nil {
		http.Error(w, "erreur de refus de l'annonce", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Annonce refusée avec succès")
}
