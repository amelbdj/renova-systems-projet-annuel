package admin

import (
	"encoding/json"
	"net/http"
	"strconv"
	"upcycleconnect/bdd"
	"upcycleconnect/models"
)

func CreateEtapeHandler(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w, "POST")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var e models.Etape
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}
	if e.Statut == "" {
		e.Statut = "a_faire"
	}

	if err := bdd.CreateEtape(e); err != nil {
		http.Error(w, "Erreur création étape", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Étape créée"})
}

func GetEtapesHandler(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w, "GET")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	idProjet, err := strconv.Atoi(r.URL.Query().Get("id_projet"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	etapes, err := bdd.GetEtapesByProjet(idProjet)
	if err != nil {
		http.Error(w, "Erreur récupération", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(etapes)
}

func DeleteEtapeHandler(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w, "DELETE")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	if err := bdd.DeleteEtape(id); err != nil {
		http.Error(w, "Erreur suppression", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Étape supprimée"})
}

func UpdateEtapeStatutHandler(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w, "PUT")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var body struct {
		Id     int    `json:"id"`
		Statut string `json:"statut"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}
	if err := bdd.UpdateEtapeStatut(body.Id, body.Statut); err != nil {
		http.Error(w, "Erreur mise à jour", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Statut mis à jour"})
}
