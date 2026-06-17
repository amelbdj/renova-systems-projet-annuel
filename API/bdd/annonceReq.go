package bdd

import (
	"fmt"
	"strings"
	"upcycleconnect/models"
)

func GetAnnonces() ([]models.Annonce, error) {

	var Annonces []models.Annonce

	rows, err := Db.Query("SELECT DATE_FORMAT(a.created_at, '%d-%m-%Y') as created_at, a.id, a.titre, a.description, a.type, a.prix, a.statut_validation, a.code_postal, a.ville, a.etat, a.poids_kg, a.quantite, a.image, u.nom, u.prenom, c.libelle FROM pa2026.annonce a INNER JOIN pa2026.utilisateur u ON u.id = a.id_user INNER JOIN pa2026.categorie c ON c.id = a.id_categorie")

	if err != nil {
		return nil, fmt.Errorf("get Annonces : %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		var Annonce models.Annonce

		err := rows.Scan(&Annonce.CreatedAt, &Annonce.Id, &Annonce.Titre, &Annonce.Description, &Annonce.Type, &Annonce.Prix, &Annonce.StatutValidation,
			&Annonce.CodePostal, &Annonce.Ville, &Annonce.Etat, &Annonce.PoidsKg, &Annonce.Quantite, &Annonce.Image,
			&Annonce.Nom, &Annonce.Prenom,
			&Annonce.Categorie)

		if err != nil {
			return nil, fmt.Errorf("get Annonces : %v", err.Error())
		}
		Annonces = append(Annonces, Annonce)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("get Annonces : %v", err.Error())
	}

	return Annonces, nil
}

func ValidateAnnonce(annonceId int) error {

	result, err := Db.Exec("UPDATE pa2026.annonce SET statut_validation = 'Validé' WHERE id = ?", annonceId)

	if err != nil {
		return fmt.Errorf("mise à jour échouée : %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {

		return err
	}

	if rows == 0 {
		return fmt.Errorf("aucune annonce trouvée avec l'id %d", annonceId)
	}

	return nil
}

func RefuseAnnonce(annonceId int) error {
	result, err := Db.Exec("UPDATE pa2026.annonce SET statut_validation = 'Rejeté' WHERE id = ?", annonceId)

	if err != nil {
		return fmt.Errorf("mise à jour échouée : %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("aucune annonce trouvée avec l'id %d", annonceId)
	}

	return nil
}

func CreateAnnonce(annonce models.Annonce) error {
	query := `INSERT INTO pa2026.annonce 
              (titre, description, type, prix, code_postal, ville, etat, poids_kg, quantite, id_user, id_categorie, image)
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := Db.Exec(query,
		annonce.Titre,
		annonce.Description,
		annonce.Type,
		annonce.Prix,
		annonce.CodePostal,
		annonce.Ville,
		annonce.Etat,
		annonce.PoidsKg,
		annonce.Quantite,
		annonce.IdUser,
		annonce.IdCategorie,
		annonce.Image,
	)

	if err != nil {
		fmt.Println("CreateAnnonce Error:", err)
		return fmt.Errorf("CreateAnnonce : %s", err.Error())
	}

	return nil
}

func DeleteAnnonce(id int) error {

	_, err := Db.Query("SELECT id FROM pa2026.annonce WHERE id = ?", id)
	if err != nil {
		fmt.Println("erreur", err)
		return fmt.Errorf("l'annonce n'existe pas : %d", id)

	}

	// sUPPRESSION
	_, err = Db.Exec(
		"DELETE FROM pa2026.annonce WHERE id = ?", id)
	if err != nil {
		fmt.Println("erreur", err)

		return fmt.Errorf("mise à jour échouée : %v", err)
	}

	return nil
}

func UpdateAnnonce(annonceId int, annonce models.Annonce) error {
	StatutVente := ""

	var id int
	err := Db.QueryRow("SELECT id FROM pa2026.annonce WHERE id = ?", annonceId).Scan(&id)
	if err != nil {
		return fmt.Errorf("l'annonce n'existe pas : %d", annonceId)
	}

	err = Db.QueryRow("SELECT statut_vente FROM pa2026.annonce WHERE id = ?", annonceId).Scan(&StatutVente)
	if err != nil {
		return fmt.Errorf("UpdateAnnonce Scan Error: %s", err.Error())
	}

	if StatutVente != "EN ATTENTE DEPOT" {
		_, err = Db.Exec(
			"UPDATE pa2026.annonce SET titre = ?, description = ?, type = ?, prix = ?, code_postal = ?, ville = ?, etat = ?, poids_kg = ?, quantite = ?, id_user = ?, id_categorie = ?, image = ? WHERE id = ?",
			annonce.Titre,
			annonce.Description,
			annonce.Type,
			annonce.Prix,
			annonce.CodePostal,
			annonce.Ville,
			annonce.Etat,
			annonce.PoidsKg,
			annonce.Quantite,
			annonce.IdUser,
			annonce.IdCategorie,
			annonce.Image,
			annonceId,
		)

		if err != nil {
			return fmt.Errorf("UpdateAnnonce Exec Error: %s", err.Error())
		}
	}

	return nil
}

func GetAnnonceByTitle(query string, filtre string) ([]models.Annonce, error) {
	query = strings.ToUpper(query)
	search := "%" + query + "%"
	if filtre != "Tout" && filtre != "" {
		var Annonces []models.Annonce
		rows, err := Db.Query("SELECT id, titre, description, type, prix, statut_validation, code_postal, ville, etat, poids_kg, quantite, nom, prenom, categorie FROM pa2026.annonce WHERE (UPPER(titre) LIKE ?) AND statut_validation = ?", search, filtre)

		if err != nil {
			fmt.Println("Erreur lors de l'exécution de la requête : ", err)
			return nil, fmt.Errorf("get Annonce by title : %v", err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var Annonce models.Annonce

			err := rows.Scan(&Annonce.Id, &Annonce.Titre, &Annonce.Description, &Annonce.Type, &Annonce.Prix, &Annonce.StatutValidation, &Annonce.CodePostal, &Annonce.Ville, &Annonce.Etat, &Annonce.PoidsKg, &Annonce.Quantite, &Annonce.Nom, &Annonce.Prenom, &Annonce.Categorie)

			if err != nil {
				fmt.Println("Erreur lors de l'exécution de la requête : ", err)
				return nil, fmt.Errorf("get Annonce by title : %v", err.Error())
			}

			Annonces = append(Annonces, Annonce)
		}

		err = rows.Err()

		if err != nil {
			return nil, fmt.Errorf("get Annonce by title : %v", err.Error())
		}
		return Annonces, nil
	} else {
		var Annonces []models.Annonce
		search := "%" + query + "%"
		rows, err := Db.Query("SELECT id, titre, description, type, prix, statut_validation, code_postal, ville, etat, poids_kg, quantite, nom, prenom, categorie FROM pa2026.annonce WHERE UPPER(titre) LIKE ?", search)

		if err != nil {
			return nil, fmt.Errorf("get Annonce by title : %v", err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var Annonce models.Annonce

			err := rows.Scan(&Annonce.Id, &Annonce.Titre, &Annonce.Description, &Annonce.Type, &Annonce.Prix, &Annonce.StatutValidation, &Annonce.CodePostal, &Annonce.Ville, &Annonce.Etat, &Annonce.PoidsKg, &Annonce.Quantite, &Annonce.Nom, &Annonce.Prenom, &Annonce.Categorie)

			if err != nil {
				return nil, fmt.Errorf("get Annonce by title : %v", err.Error())
			}

			Annonces = append(Annonces, Annonce)
		}

		err = rows.Err()

		if err != nil {
			return nil, fmt.Errorf("get Annonce by title : %v", err.Error())
		}
		return Annonces, nil
	}
}

func GetAnnonceById(id int) (models.Annonce, error) {
	var a models.Annonce

	query := `
        SELECT 
         a.id, 
         a.id_user, -- <-- AJOUTÉ ICI
         a.titre, a.description, a.type, a.prix, a.statut_vente, a.statut_validation, 
         a.code_postal, a.ville, a.etat, a.poids_kg, a.quantite,
         COALESCE(pa2026.utilisateur.nom, ''), 
         COALESCE(pa2026.utilisateur.prenom, ''), 
         COALESCE(pa2026.categorie.libelle, ''), 
         COALESCE(a.image, '')
    FROM pa2026.annonce a
    LEFT JOIN pa2026.utilisateur ON a.id_user = pa2026.utilisateur.id
    LEFT JOIN pa2026.categorie ON a.id_categorie = pa2026.categorie.id
    WHERE a.id = ?`

	err := Db.QueryRow(query, id).Scan(
		&a.Id,
		&a.IdUser,
		&a.Titre, &a.Description, &a.Type, &a.Prix, &a.StatutVente, &a.StatutValidation,
		&a.CodePostal, &a.Ville, &a.Etat, &a.PoidsKg, &a.Quantite,
		&a.Nom, &a.Prenom, &a.Categorie, &a.Image,
	)

	if err != nil {
		return a, fmt.Errorf("get Annonce by id : %v", err)
	}

	return a, nil
}

func GetAnnoncesByUser(userID int) ([]models.Annonce, error) {
	var list []models.Annonce

	// 🟢 CORRECTION : On protège TOUTES les colonnes contre les valeurs NULL
	query := `
		SELECT id, titre, COALESCE(prix, 0), COALESCE(id_categorie, 0), 
		       COALESCE(statut_vente, ''), COALESCE(statut_validation, ''), COALESCE(image, '') 
		FROM pa2026.annonce 
		WHERE id_user = ?`

	rows, err := Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Annonce models.Annonce
		if err := rows.Scan(&Annonce.Id, &Annonce.Titre, &Annonce.Prix, &Annonce.IdCategorie, &Annonce.StatutVente, &Annonce.StatutValidation, &Annonce.Image); err != nil {
			return nil, err
		}
		list = append(list, Annonce)
	}
	return list, nil
}

func GetValidatedAnnonces(currentUserID int) ([]models.Annonce, error) {
    var Annonces []models.Annonce

    query := `
        SELECT 
        a.id, a.titre, a.description, a.type, a.prix, a.statut_validation, 
        a.code_postal, a.ville, a.etat, a.poids_kg, a.quantite,
        COALESCE(u.nom, ''), 
        COALESCE(u.prenom, ''), 
        COALESCE(c.libelle, ''), 
        COALESCE(a.image, ''),
        a.statut_vente
    FROM pa2026.annonce a
    LEFT JOIN pa2026.utilisateur u ON a.id_user = u.id
    LEFT JOIN pa2026.categorie c ON a.id_categorie = c.id
    WHERE a.statut_validation = 'Validé' 
    AND a.id_user != ?
    -- LA CORRECTION EST ICI : on exige explicitement que l'annonce soit "En vente"
    AND a.statut_vente = 'En vente'`

    rows, err := Db.Query(query, currentUserID)
    if err != nil {
        return nil, fmt.Errorf("Erreur Query: %v", err)
    }
    defer rows.Close()

    for rows.Next() {
        var a models.Annonce
        err := rows.Scan(
            &a.Id, &a.Titre, &a.Description, &a.Type, &a.Prix, &a.StatutValidation,
            &a.CodePostal, &a.Ville, &a.Etat, &a.PoidsKg, &a.Quantite,
            &a.Nom, &a.Prenom, &a.Categorie, &a.Image, &a.StatutVente,
        )
        if err != nil {
            return nil, fmt.Errorf("Erreur Scan: %v", err)
        }
        Annonces = append(Annonces, a)
    }
    return Annonces, nil
}

func GetUserEcoStats(userID int) (map[string]interface{}, error) {
	var score float64
	var objetsDonnes int
	var dechetsEvites float64

	err := Db.QueryRow("SELECT COALESCE(score, 0) FROM utilisateur WHERE id = ?", userID).Scan(&score)
	if err != nil {
		score = 0
	}

	query := `
        SELECT 
            COUNT(id), 
            COALESCE(SUM(poids_kg), 0)
        FROM annonce 
        WHERE id_user = ? AND statut_vente = 'RECUPERE'`

	err = Db.QueryRow(query, userID).Scan(&objetsDonnes, &dechetsEvites)
	if err != nil {
		fmt.Println("Erreur lors du calcul des stats éco :", err)
	}

	return map[string]interface{}{
		"score":          score,
		"objets_donnes":  objetsDonnes,
		"dechets_evites": dechetsEvites,
	}, nil
}
func GetUserPurchases(buyerID int) ([]map[string]interface{}, error) {
    // On utilise des guillemets normaux (" ") pour pouvoir intégrer les backticks (`) autour du mot 'order'
    query := "SELECT h.code_barre_recuperation, h.date_reservation, a.titre, b.numero, c.nom, c.adresse, b.statut " +
             "FROM pa2026.`order` o " +
             "JOIN pa2026.annonce a ON o.id_annonce = a.id " +
             "JOIN pa2026.historique_conteneurs h ON h.annonce_id = a.id " +
             "JOIN pa2026.box b ON h.conteneur_id = b.id " +
             "JOIN pa2026.conteneur c ON b.id_conteneur = c.id " +
             "WHERE o.id_acheteur = ? AND h.date_retrait_effective IS NULL"

    rows, err := Db.Query(query, buyerID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var achats []map[string]interface{}
    for rows.Next() {
        var barcode, date, titre, nomConteneur, adresse, etat string
        var numBox int
        
        err := rows.Scan(&barcode, &date, &titre, &numBox, &nomConteneur, &adresse, &etat)
        if err != nil {
            continue
        }

        res := map[string]interface{}{
            "lieu":         nomConteneur + " - " + adresse,
            "numero_box":   fmt.Sprintf("Casier n°%d", numBox),
            "barcode":      barcode,
            "objet":        titre,
            "date":         date,
            "etat":         etat,
        }
        achats = append(achats, res)
    }
    return achats, nil
}