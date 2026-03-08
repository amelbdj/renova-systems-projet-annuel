package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"upcycleconnect/bdd"

	// "upcycleconnect/models"
	"strconv"
)

func GetAllEvenements(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("hello from GetAllEvenements")

	Evenements, err := bdd.GetEvenements()

	if err != nil {
		http.Error(w, "erreur de récupération des Evenements", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}

	response, err := json.Marshal(Evenements)

	if err != nil {
		http.Error(w, "erreur de conversion", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", response)
}

func ValidateEvenement(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")


	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}	
	err = bdd.ValidateEvenement(id)

	if err != nil {
		http.Error(w, "erreur de validation de l'Evenement", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Evenement validée avec succès")
}

func RefuseEvenement(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	

	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}	
	err = bdd.RefuseEvenement(id)

	if err != nil {
		http.Error(w, "erreur de refus de l'Evenement", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Evenement refusée avec succès")
}
