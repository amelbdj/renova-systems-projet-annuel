package bdd

import (
	"fmt"
	"upcycleconnect/models"
)

func GetArticles(searchWord string) ([]models.Article, error) {

	var Articles []models.Article

	if searchWord != "" {
		rows, err := Db.Query("SELECT id_article, id_salarie, titre, contenu, type, statut, DATE_FORMAT(created_at, '%d/%m/%Y') as created_at, utilisateur.nom, utilisateur.prenom FROM article_news INNER JOIN utilisateur ON article_news.id_salarie = utilisateur.id WHERE titre LIKE ? OR contenu LIKE ? ORDER BY id_article DESC", "%"+searchWord+"%", "%"+searchWord+"%")
		if err != nil {
			fmt.Println("Erreur lors de l'exécution de la requête : ", err)
			return nil, fmt.Errorf("get Articles : %v", err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var Article models.Article

			err := rows.Scan(&Article.Id, 
				&Article.IdSalarie, 
				&Article.Titre, 
				&Article.Contenu, 
				&Article.Type, 
				&Article.Statut,
				&Article.CreatedAt,
				&Article.NomAuteur,
				&Article.PrenomAuteur,
			)

			if err != nil {
				return nil, fmt.Errorf("get Articles : %v", err.Error())
			}

			
			Articles = append(Articles, Article)
		}
	} else {
		rows, err := Db.Query("SELECT id_article, id_salarie, titre, contenu, type, statut, DATE_FORMAT(created_at, '%d/%m/%Y') as created_at, utilisateur.nom, utilisateur.prenom FROM article_news INNER JOIN utilisateur ON article_news.id_salarie = utilisateur.id ORDER BY id_article DESC")
		if err != nil {
			fmt.Println("Erreur lors de l'exécution de la requête : ", err)
			return nil, fmt.Errorf("get Articles : %v", err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var Article models.Article

			err := rows.Scan(&Article.Id, 
				&Article.IdSalarie, 
				&Article.Titre, 
				&Article.Contenu, 
				&Article.Type, 
				&Article.Statut,
				&Article.CreatedAt,
				&Article.NomAuteur,
				&Article.PrenomAuteur,
			)

			if err != nil {
				return nil, fmt.Errorf("get Articles : %v", err.Error())
			}

			
			Articles = append(Articles, Article)
		}
	}

	return Articles, nil
}

func GetArticlesBySalarie(salarieID int) ([]models.Article, error) {

    rows, err := Db.Query("SELECT id_article, id_salarie, titre, contenu, type, statut FROM article_news WHERE id_salarie = ?", salarieID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var articles []models.Article
    for rows.Next() {
        var article models.Article
        err := rows.Scan(&article.Id, &article.IdSalarie, &article.Titre, &article.Contenu, &article.Type, &article.Statut)
        if err != nil {
            return nil, err
        }
        articles = append(articles, article)
    }
    return articles, nil
}

func ValidateArticle(id int) error {
    _, err := Db.Exec("UPDATE article_news SET statut = 'valide' WHERE id_article = ?", id)
    return err
}

func RefuseArticle(id int) error {
    _, err := Db.Exec("UPDATE article_news SET statut = 'refuse' WHERE id_article = ?", id)
    return err
}

func DeleteArticle(id int) error {
    _, err := Db.Exec("DELETE FROM article_news WHERE id_article = ?", id)
    return err
}

func CreateArticle(article models.Article, action string)error{
    var statutFinal string
    if action == "publier" {
        statutFinal = "en attente" 
    } else {
        statutFinal = "brouillon"
    }


              
    _, err := Db.Exec(`INSERT INTO article_news (id_salarie, titre, contenu, type, statut) 
              VALUES (?, ?, ?, ?, ?)`, article.IdSalarie, article.Titre, article.Contenu, article.Type, statutFinal)
    if err != nil {
        return  err
    }

   

    return  nil
}

func ModifyArticle(id int, titre, contenu, articleType, action string) error {
	var statutFinal string
	if action == "publier" {
		statutFinal = "en attente"
	} else {
		statutFinal = "brouillon"
	}
	_, err := Db.Exec(`UPDATE article_news SET titre = ?, contenu = ?, type = ?, statut = ? WHERE id_article = ?`, titre, contenu, articleType, statutFinal, id)
	if err != nil {
		return  err
	}
	return nil
}

func GetArticleById(id int) (models.Article, error) {

    var article models.Article 

    row := Db.QueryRow("SELECT id_article, id_salarie, titre, contenu, type, statut FROM article_news WHERE id_article = ?", id)
    err := row.Scan(&article.Id, &article.IdSalarie, &article.Titre, &article.Contenu, &article.Type, &article.Statut)
    if err != nil {
        return models.Article{}, err
    }
    return article, nil
}

