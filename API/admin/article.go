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

func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	searchWord := r.URL.Query().Get("search")
	articles, err := bdd.GetArticles(searchWord)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des articles", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(articles)
	if err != nil {
		http.Error(w, "Erreur JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetArticlesBySalarie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	salarieId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	articles, err := bdd.GetArticlesBySalarie(salarieId)
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	if articles == nil {
		articles = []models.Article{}
	}

	response, _ := json.Marshal(articles)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	articleId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	err = bdd.DeleteArticle(articleId)
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"message":"article suppr"}`)
}

func ValidateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	articleId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	article, errGet := bdd.GetArticleById(articleId)
	err = bdd.ValidateArticle(articleId)
	if err != nil {
		http.Error(w, "Erreur validation", http.StatusInternalServerError)
		return
	}

	if errGet == nil && article.IdSalarie != 0 {
		msg := fmt.Sprintf("✅ Super ! Ton article '%s' a été validé et publié.", article.Titre)
		go SendPushNotification(strconv.Itoa(article.IdSalarie), msg)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"message":"article validé"}`)
}

func RefuseArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	articleId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	err = bdd.RefuseArticle(articleId)
	if err != nil {
		http.Error(w, "Erreur refus", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"message":"article refusé"}`)
}

func ModifyArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	action := r.PathValue("action")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	errParse := r.ParseMultipartForm(10 << 20)
	if errParse != nil {
		http.Error(w, "Impossible de lire le formulaire", http.StatusBadRequest)
		return
	}

	titre := r.FormValue("titre")
	contenu := r.FormValue("contenu")
	articleType := r.FormValue("type")

	imageUrl := ""
	file, handler, errFile := r.FormFile("image")
	if errFile == nil {
		defer file.Close()
		os.MkdirAll("./static/uploads/articles", os.ModePerm)
		fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)
		path := "./static/uploads/articles/" + fileName
		f, errCreate := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		if errCreate == nil {
			defer f.Close()
			io.Copy(f, file)
			imageUrl = "static/uploads/articles/" + fileName
		}
	}

	err = bdd.ModifyArticle(id, titre, contenu, articleType, action, imageUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "article modifié avec succès"}`))
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	action := r.PathValue("action")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Impossible de lire le formulaire", http.StatusBadRequest)
		return
	}

	idSalarie, _ := strconv.Atoi(r.FormValue("id_salarie"))

	article := models.Article{
		IdSalarie: idSalarie,
		Titre:     r.FormValue("titre"),
		Contenu:   r.FormValue("contenu"),
		Type:      r.FormValue("type"),
	}

	file, handler, errFile := r.FormFile("image")
	if errFile == nil {
		defer file.Close()
		os.MkdirAll("./static/uploads/articles", os.ModePerm)
		fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), handler.Filename)
		path := "./static/uploads/articles/" + fileName
		f, errCreate := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		if errCreate == nil {
			defer f.Close()
			io.Copy(f, file)
			article.ImageUrl = "static/uploads/articles/" + fileName
		}
	}

	err = bdd.CreateArticle(article, action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	NotifyAllAdmins("Un nouvel article attend votre relecture.")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "article créé"}`))
}

func GetArticleById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	articleId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	article, err := bdd.GetArticleById(articleId)
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(article)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
