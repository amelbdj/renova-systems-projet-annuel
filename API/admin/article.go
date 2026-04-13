package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"upcycleconnect/bdd"
	"upcycleconnect/models"
)

func GetAllArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("hello from get articles")

	
	articles, err := bdd.GetArticles()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des articles", http.StatusInternalServerError)
		return
	}

	if err != nil {
		fmt.Println(err)
		http.Error(w, "erreur de récupération des utilisateurs", http.StatusInternalServerError)

		return
	}

		response, err := json.Marshal(articles)
		if err != nil {
			http.Error(w, "Erreur lors de la conversion des articles en JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)

}

func GetArticlesBySalarie(w http.ResponseWriter, r *http.Request) {

	salarieId, err := strconv.Atoi(r.URL.Query().Get("salarieId"))
	if err != nil {
		http.Error(w, "ID de salarié invalide", http.StatusBadRequest)
		return
	}

	articles, err := bdd.GetArticlesBySalarie(salarieId)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des articles", http.StatusInternalServerError)
		return
	}	

	response, err := json.Marshal(articles)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion des articles en JSON", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

   if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return 
    }
	articleId, err := strconv.Atoi(r.URL.Query().Get("articleId"))
	if err != nil {
		http.Error(w, "ID d'article invalide", http.StatusBadRequest)
		return
	}

	err = bdd.DeleteArticle(articleId)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression de l'article", http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "utilisateur suppr")
}
func ValidateArticle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

   if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return 
    }
	articleId, err := strconv.Atoi(r.URL.Query().Get("articleId"))
	if err != nil {
		http.Error(w, "ID d'article invalide", http.StatusBadRequest)
		return
	}

	err = bdd.ValidateArticle(articleId)
	if err != nil {
		http.Error(w, "Erreur lors de la validation de l'article", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "article validé")
}

func RefuseArticle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	   if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return 
    }

	articleId, err := strconv.Atoi(r.URL.Query().Get("articleId"))
	if err != nil {
		http.Error(w, "ID d'article invalide", http.StatusBadRequest)
		return
	}
	
	err = bdd.RefuseArticle(articleId)
	if err != nil {
		http.Error(w, "Erreur lors du refus de l'article", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "article refusé")
}

func ModifyArticle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	   if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return 
    }

	action := r.PathValue("action")
	if action == "" {
		http.Error(w, "Action invalide", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID d'article invalide", http.StatusBadRequest)
		return
	}



    fmt.Println("hello from modify article")

	var article models.Article
	err = json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		fmt.Println("Erreur décodage :", err)
		http.Error(w, "Impossible de décoder le JSON", http.StatusBadRequest)
		return
	}

	err = bdd.ModifyArticle(id, article.Titre, article.Contenu, article.Type, action)
	if err != nil {
		fmt.Println("Erreur lors de la modification en BDD :", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}



}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	action := r.PathValue("action")
	if action == "" {
		http.Error(w, "Action invalide", http.StatusBadRequest)
		return
	}

    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return 
    }

    fmt.Println("hello from create article")

    var article models.Article
    err := json.NewDecoder(r.Body).Decode(&article)
    if err != nil {
        fmt.Println("Erreur décodage :", err)
        http.Error(w, "Impossible de décoder le JSON", http.StatusBadRequest)
        return
    }

   
    
    err = bdd.CreateArticle(article, action)
	if err != nil {
    fmt.Println("Erreur lors de l'insertion en BDD :", err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
    
    w.Header().Set("Content-Type", "application/json") 
    w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "article créé")                  
    

}