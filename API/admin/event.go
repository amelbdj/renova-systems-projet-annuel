package admin

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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

	var idSalarie int
	var titre string
	errInfo := bdd.Db.QueryRow("SELECT id_salarie, titre FROM evenement WHERE id = ?", id).Scan(&idSalarie, &titre)
	if errInfo == nil && idSalarie != 0 {
		msg := fmt.Sprintf("✅ Votre événement '%s' a été validé et est maintenant en ligne !", titre)
		go SendPushNotification(strconv.Itoa(idSalarie), msg)
		bdd.CreateNotification(idSalarie, msg)
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

	var idSalarie int
	var titre string
	errInfo := bdd.Db.QueryRow("SELECT id_salarie, titre FROM evenement WHERE id = ?", id).Scan(&idSalarie, &titre)
	if errInfo == nil && idSalarie != 0 {
		msg := fmt.Sprintf("❌ Votre événement '%s' a été refusé par un responsable.", titre)
		go SendPushNotification(strconv.Itoa(idSalarie), msg)
		bdd.CreateNotification(idSalarie, msg)
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

	fmt.Println("hello from CreateEvenement (Multipart Mode)")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du formulaire", http.StatusBadRequest)
		fmt.Println("Erreur ParseMultipartForm :", err)
		return
	}

	var Evenement models.Evenement
	Evenement.IdSalarie, _ = strconv.Atoi(r.FormValue("idSalarie"))
	Evenement.Titre = r.FormValue("titre")
	Evenement.Type = r.FormValue("type")
	Evenement.Description = r.FormValue("description")
	Evenement.DateDebut = r.FormValue("date_debut")
	Evenement.DateFin = r.FormValue("date_fin")
	Evenement.Lieu = r.FormValue("lieu")
	Evenement.NbPlaces, _ = strconv.Atoi(r.FormValue("capacite"))
	Evenement.Prix, _ = strconv.ParseFloat(r.FormValue("tarif"), 64)
	Evenement.PlanCours = r.FormValue("plan_cours")

	dateDebut, errDate := time.ParseInLocation("2006-01-02 15:04:05", Evenement.DateDebut, time.Local)
	if errDate != nil {
		http.Error(w, "Date de début invalide", http.StatusBadRequest)
		return
	}
	if dateDebut.Before(time.Now()) {
		http.Error(w, "Impossible de créer un événement à une date ou une heure déjà passée", http.StatusBadRequest)
		return
	}

	file, handler, errFile := r.FormFile("image")
	if errFile == nil {
		defer file.Close()
		os.MkdirAll("./static/uploads/events", os.ModePerm)
		nomFichier := fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)
		cheminComplet := "./static/uploads/events/" + nomFichier

		f, err := os.OpenFile(cheminComplet, os.O_WRONLY|os.O_CREATE, 0666)
		if err == nil {
			defer f.Close()
			io.Copy(f, file)
			Evenement.ImageUrl = "static/uploads/events/" + nomFichier
		} else {
			fmt.Println("Erreur création fichier :", err)
		}
	}

	newId, err := bdd.CreateEvenement(Evenement)
	if err != nil {
		http.Error(w, "erreur de création de l'Evenement", http.StatusInternalServerError)
		fmt.Println("erreur bdd.CreateEvenement :", err)
		return
	}

	planFile, planHandler, errPlan := r.FormFile("plan_pdf")
	if errPlan == nil {
		defer planFile.Close()
		urlPlan, errSave := enregistrerPdf(planFile, planHandler)
		if errSave == nil {
			bdd.CreateRessource(Evenement.IdSalarie, int(newId), "Plan du cours", urlPlan)
		} else {
			fmt.Println("Erreur enregistrement plan PDF :", errSave)
		}
	}

	if r.MultipartForm != nil {
		fichiers := r.MultipartForm.File["ressources"]
		for i := 0; i < len(fichiers); i++ {
			f, errOpen := fichiers[i].Open()
			if errOpen != nil {
				continue
			}
			urlRes, errSave := enregistrerPdf(f, fichiers[i])
			f.Close()
			if errSave == nil {
				bdd.CreateRessource(Evenement.IdSalarie, int(newId), fichiers[i].Filename, urlRes)
			} else {
				fmt.Println("Erreur enregistrement ressource PDF :", errSave)
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Evenement créée avec succès")
}

func enregistrerPdf(file multipart.File, handler *multipart.FileHeader) (string, error) {
	if strings.ToLower(filepath.Ext(handler.Filename)) != ".pdf" {
		return "", fmt.Errorf("le fichier doit être au format PDF")
	}
	os.MkdirAll("./static/uploads/formations", os.ModePerm)
	nomFichier := fmt.Sprintf("%d_%s", time.Now().UnixNano(), handler.Filename)
	cheminComplet := "./static/uploads/formations/" + nomFichier

	f, err := os.OpenFile(cheminComplet, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()
	io.Copy(f, file)
	return "static/uploads/formations/" + nomFichier, nil
}

func GetInscritsEvenement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
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

	inscrits, err := bdd.GetInscrits(id)
	if err != nil {
		http.Error(w, "erreur de récupération des inscrits", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}

	response, err := json.Marshal(inscrits)
	if err != nil {
		http.Error(w, "erreur de conversion", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", response)
}

func GetRessourcesEvenement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
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

	ressources, err := bdd.GetRessources(id)
	if err != nil {
		http.Error(w, "erreur de récupération des ressources", http.StatusInternalServerError)
		fmt.Println("erreur", err)
		return
	}

	response, err := json.Marshal(ressources)
	if err != nil {
		http.Error(w, "erreur de conversion", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", response)
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
