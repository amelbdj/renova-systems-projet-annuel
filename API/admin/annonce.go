package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"upcycleconnect/bdd"
	"upcycleconnect/models"

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

func CreateAnnonce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var annonce models.Annonce
	err := json.NewDecoder(r.Body).Decode(&annonce)
	if err != nil {
		http.Error(w, "données invalides", http.StatusBadRequest)
		return
	}
	err = bdd.CreateAnnonce(annonce)
	if err != nil {
		http.Error(w,
			"erreur dans la création d'une annonce",
			http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteAnnonce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println("hello from DeleteAnnonce")
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

	err = bdd.DeleteAnnonce(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "annonce suppr")
}

func UpdateAnnonce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "id invalide", http.StatusBadRequest)
		return
	}
	var annonce models.Annonce
	err = json.NewDecoder(r.Body).Decode(&annonce)
	if err != nil {
		http.Error(w, "données invalides", http.StatusBadRequest)
		return
	}
	err = bdd.UpdateAnnonce(id, annonce)
	if err != nil {
		http.Error(w,
			"erreur dans la mise à jour de l'annonce",
			http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Annonce mise à jour avec succès")
}

func GetAnnonceByTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println("hello from GetAnnonceByTitle")
	query := r.URL.Query().Get("query")
	filtre := r.URL.Query().Get("filtre")
	Annonces, err := bdd.GetAnnonceByTitle(query, filtre)

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

func GetMyAnnonces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID nvalide", http.StatusBadRequest)
		return
	}

	annonces, err := bdd.GetAnnoncesByUser(userID)
	if err != nil {
		http.Error(w, "Erreur lors de la recup de vos annonces", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(annonces)
}
