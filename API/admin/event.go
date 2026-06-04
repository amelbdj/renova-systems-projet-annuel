package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"upcycleconnect/bdd"
	"upcycleconnect/models"
)

func GetAllEvenements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return 
	}

	searchWord := r.URL.Query().Get("search")
	fmt.Println("hello from GetAllEvenements")

	var idUser int
	if val := r.Context().Value("userID"); val != nil {
		idUser, _ = val.(int)
	} else if val := r.Context().Value("user_id"); val != nil {
		idUser, _ = val.(int)
	} else if val := r.Context().Value("id_user"); val != nil {
		idUser, _ = val.(int)
	}


	Evenements, err := bdd.GetEvenements(searchWord, idUser)

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
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
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
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
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
	err = bdd.RefuseEvenement(id)
	if err != nil {
		http.Error(w, "erreur de refus de l'Evenement", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Evenement refusée avec succès")
}

func CreateEvenement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return 
	}
	var Evenement models.Evenement
	fmt.Println("hello from CreateEvenement")
	err := json.NewDecoder(r.Body).Decode(&Evenement)
	if err != nil {
		http.Error(w, "données invalides", http.StatusBadRequest)
		fmt.Println("erreur", err)
		return
	}
	err = bdd.CreateEvenement(Evenement)
	if err != nil {
		http.Error(w, "erreur de création de l'Evenement", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Evenement créée avec succès")
}

func DeleteEvenement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
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
	err = bdd.DeleteEvenement(id)
	if err != nil {
		http.Error(w, "erreur de suppression de l'Evenement", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Evenement supprimée avec succès")
}

func UpdateEvenement(w http.ResponseWriter, r *http.Request) {
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
	var Evenement models.Evenement

	err = json.NewDecoder(r.Body).Decode(&Evenement)	
	if err != nil {
		http.Error(w, "données invalides", http.StatusBadRequest)
		return
	}
	err = bdd.UpdateEvenement(id, Evenement)
	if err != nil {
		http.Error(w, "erreur de mise à jour de l'Evenement", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Evenement mise à jour avec succès")
}

func InscrireClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return 
	}

	var insc models.InscriptionRequest
	err := json.NewDecoder(r.Body).Decode(&insc)
	if err != nil {
		http.Error(w, "données invalides", http.StatusBadRequest)
		return
	}

	err = bdd.InscrireClient(insc.IdUser, insc.IdEvent)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"erreur": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Client inscrit avec succès !"})
}

type DesinscriptionReq struct {
	IdUser  int `json:"id_user"`
	IdEvent int `json:"id_event"`
}

func DesinscriptionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, `{"erreur": "Méthode non autorisée"}`, http.StatusMethodNotAllowed)
		return
	}

	var req DesinscriptionReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"erreur": "Données JSON invalides"})
		return
	}

	err := bdd.SupprimerInscription(req.IdUser, req.IdEvent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"erreur": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Désinscription effectuée avec succès",
	})
}