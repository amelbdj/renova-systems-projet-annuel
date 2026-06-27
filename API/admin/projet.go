package admin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
	"upcycleconnect/bdd"
	"upcycleconnect/models"
)

func corsHeaders(w http.ResponseWriter, methods string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", methods+", OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func saveUpload(r *http.Request, field string) (string, bool) {
	file, header, err := r.FormFile(field)
	if err != nil {
		return "", false
	}
	defer file.Close()

	name := fmt.Sprintf("%d_%s", time.Now().UnixNano(), header.Filename)
	out, err := os.Create("./uploads/" + name)
	if err != nil {
		return "", false
	}
	defer out.Close()
	io.Copy(out, file)
	return "/view-uploads/" + name, true
}

func CreateProjetHandler(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w, "POST")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	r.ParseMultipartForm(10 << 20)
	idUser, _ := strconv.Atoi(r.FormValue("id_user"))

	co2, _ := strconv.ParseFloat(r.FormValue("co2_evite"), 64)
	statut := r.FormValue("statut")
	if statut == "" {
		statut = "en_cours"
	}

	p := models.Projet{
		IdUser:      idUser,
		Titre:       r.FormValue("titre"),
		Description: r.FormValue("description"),
		Statut:      statut,
		Co2Evite:    co2,
	}
	if url, ok := saveUpload(r, "photo_avant"); ok {
		p.PhotoAvant = url
	}
	if url, ok := saveUpload(r, "photo_apres"); ok {
		p.PhotoApres = url
	}
	if err := bdd.CreateProjet(p); err != nil {
		http.Error(w, "Erreur création", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Projet créé"})
}

func GetProjetsHandler(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w, "GET")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	idUser, err := strconv.Atoi(r.URL.Query().Get("id_user"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	projets, err := bdd.GetProjetsByUser(idUser)
	if err != nil {
		http.Error(w, "Erreur récupération", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(projets)
}

func DeleteProjetHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := bdd.DeleteProjet(id); err != nil {
		http.Error(w, "Erreur suppression", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Projet supprimé"})
}

func UpdateProjetHandler(w http.ResponseWriter, r *http.Request) {
	corsHeaders(w, "PUT")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	r.ParseMultipartForm(10 << 20)
	id, _ := strconv.Atoi(r.FormValue("id"))

	co2, _ := strconv.ParseFloat(r.FormValue("co2_evite"), 64)

	statut := r.FormValue("statut")
	if statut == "" {
		statut = "en_cours"
	}

	p := models.Projet{
		Id:          id,
		Titre:       r.FormValue("titre"),
		Description: r.FormValue("description"),
		Statut:      statut,
		Co2Evite:    co2,
	}

	if url, ok := saveUpload(r, "photo_avant"); ok {
		p.PhotoAvant = url
	} else {
		p.PhotoAvant = r.FormValue("old_photo_avant")
	}
	if url, ok := saveUpload(r, "photo_apres"); ok {
		p.PhotoApres = url
	} else {
		p.PhotoApres = r.FormValue("old_photo_apres")
	}

	if err := bdd.UpdateProjet(p); err != nil {
		http.Error(w, "Erreur mise à jour", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Projet mis à jour"})
}
