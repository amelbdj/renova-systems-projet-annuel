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

	
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	
	if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

	fmt.Println("hello from get articles")

	searchWord := r.URL.Query().Get("search")

	
	articles, err := bdd.GetArticles(searchWord)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des articles", http.StatusInternalServerError)
		fmt.Println(err)
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	
	if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

salarieId, err := strconv.Atoi(r.PathValue("id"))
if err != nil {
    fmt.Println("Erreur conversion ID :", err)
    http.Error(w, "ID de salarié invalide ou manquant", http.StatusBadRequest)
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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
   
	if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return 
    }
	articleId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID d'article invalide", http.StatusBadRequest)
		fmt.Println("Erreur conversion ID :", err)
		return
	}

	err = bdd.DeleteArticle(articleId)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression de l'article", http.StatusInternalServerError)
				fmt.Println( err)

		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "article suppr")
}
func ValidateArticle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

   if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return 
    }
	articleId, err := strconv.Atoi(r.PathValue("id"))
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
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	   if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return 
    }

	articleId, err := strconv.Atoi(r.PathValue("id"))
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
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

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
	if action == "" {
		http.Error(w, "Action invalide", http.StatusBadRequest)
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
    w.Write([]byte(`{"message": "article créé"}`))               
    

}

func GetArticleById(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	articleId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID d'article invalide", http.StatusBadRequest)
		return
	}
	article, err := bdd.GetArticleById(articleId)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de l'article", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion de l'article en JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

