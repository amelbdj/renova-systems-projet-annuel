package bdd

import (
	"fmt"
	"upcycleconnect/models"
)

func GetEvenements(searchWord string, idUser int) ([]models.Evenement, error) {

	var Evenements []models.Evenement

	if searchWord != "" {
		query := `SELECT evenement.id, evenement.titre, evenement.description, 
		          DATE_FORMAT(evenement.date_debut, '%d/%m/%Y a %H:%i') as date_debut, 
		          DATE_FORMAT(evenement.date_fin, '%d/%m/%Y a %H:%i') as date_fin, 
		          evenement.nb_places, evenement.statut_validation, evenement.format, 
		          evenement.lieu, evenement.type, evenement.id_salarie, utilisateur.nom,
		          utilisateur.prenom, evenement.prix, COALESCE(evenement.image_url, '') as image_url,
		          COALESCE(evenement.plan_cours, '') as plan_cours,
		          COALESCE((SELECT url_fichier FROM ressource_pedagogique WHERE id_event = evenement.id ORDER BY id_ressource DESC LIMIT 1), '') as pdf_url,
		          (SELECT COUNT(*) FROM inscription WHERE id_user = ? AND id_event = evenement.id) > 0 AS deja_inscrit
		          FROM pa2026.Evenement
		          INNER JOIN utilisateur ON utilisateur.id = Evenement.id_salarie
		          WHERE evenement.titre LIKE ?`

		rows, err := Db.Query(query, idUser, "%"+searchWord+"%")
		if err != nil {
			return nil, fmt.Errorf("get Evenements : %v", err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var Evenement models.Evenement
			err := rows.Scan(&Evenement.Id, &Evenement.Titre, &Evenement.Description, &Evenement.DateDebut, &Evenement.DateFin, &Evenement.NbPlaces, &Evenement.StatutValidation, &Evenement.Format, &Evenement.Lieu, &Evenement.Type, &Evenement.IdSalarie, &Evenement.NomSalarie, &Evenement.PrenomSalarie, &Evenement.Prix, &Evenement.ImageUrl, &Evenement.PlanCours, &Evenement.PdfUrl, &Evenement.DejaInscrit)
			if err != nil {
				return nil, fmt.Errorf("get Evenements scan : %v", err.Error())
			}
			Evenements = append(Evenements, Evenement)
		}
		return Evenements, nil

	} else {
		query := `SELECT evenement.id, evenement.titre, evenement.description, 
		          DATE_FORMAT(evenement.date_debut, '%d/%m/%Y a %H:%i') as date_debut, 
		          DATE_FORMAT(evenement.date_fin, '%d/%m/%Y a %H:%i') as date_fin, 
		          evenement.nb_places, evenement.statut_validation, evenement.format, 
		          evenement.lieu, evenement.type, evenement.id_salarie, utilisateur.nom,
		          utilisateur.prenom, evenement.prix, COALESCE(evenement.image_url, '') as image_url,
		          COALESCE(evenement.plan_cours, '') as plan_cours,
		          COALESCE((SELECT url_fichier FROM ressource_pedagogique WHERE id_event = evenement.id ORDER BY id_ressource DESC LIMIT 1), '') as pdf_url,
		          (SELECT COUNT(*) FROM inscription WHERE id_user = ? AND id_event = evenement.id) > 0 AS deja_inscrit
		          FROM pa2026.Evenement
		          INNER JOIN utilisateur ON utilisateur.id = Evenement.id_salarie`

		rows, err := Db.Query(query, idUser)
		if err != nil {
			return nil, fmt.Errorf("get Evenements : %v", err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var Evenement models.Evenement
			err := rows.Scan(&Evenement.Id, &Evenement.Titre, &Evenement.Description, &Evenement.DateDebut, &Evenement.DateFin, &Evenement.NbPlaces, &Evenement.StatutValidation, &Evenement.Format, &Evenement.Lieu, &Evenement.Type, &Evenement.IdSalarie, &Evenement.NomSalarie, &Evenement.PrenomSalarie, &Evenement.Prix, &Evenement.ImageUrl, &Evenement.PlanCours, &Evenement.PdfUrl, &Evenement.DejaInscrit)
			if err != nil {
				return nil, fmt.Errorf("get Evenements scan : %v", err.Error())
			}
			Evenements = append(Evenements, Evenement)
		}
		return Evenements, nil
	}
}

func ValidateEvenement(EvenementId int) error {
	result, err := Db.Exec("UPDATE pa2026.Evenement SET statut_validation = 'valide' WHERE id = ?", EvenementId)
	if err != nil {
		return fmt.Errorf("mise à jour échouée : %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("aucune Evenement trouvée avec l'id %d", EvenementId)
	}
	return nil
}

func RefuseEvenement(EvenementId int) error {
	result, err := Db.Exec("UPDATE pa2026.Evenement SET statut_validation = 'refuse' WHERE id = ?", EvenementId)
	if err != nil {
		return fmt.Errorf("mise à jour échouée : %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("aucune Evenement trouvée avec l'id %d", EvenementId)
	}
	return nil
}

func CreateEvenement(Evenement models.Evenement) (int64, error) {
	result, err := Db.Exec("INSERT INTO pa2026.Evenement (titre, description, date_debut, date_fin, nb_places, statut_validation, format, lieu, type, id_salarie, prix, image_url, plan_cours) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		Evenement.Titre, Evenement.Description, Evenement.DateDebut, Evenement.DateFin, Evenement.NbPlaces, "en attente", Evenement.Format, Evenement.Lieu, Evenement.Type, Evenement.IdSalarie, Evenement.Prix, Evenement.ImageUrl, Evenement.PlanCours)
	if err != nil {
		return 0, fmt.Errorf("création de l'événement échouée : %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func CreateRessource(idSalarie int, idEvent int, titre string, urlFichier string) error {
	_, err := Db.Exec("INSERT INTO pa2026.ressource_pedagogique (id_salarie, id_event, titre, url_fichier) VALUES (?, ?, ?, ?)",
		idSalarie, idEvent, titre, urlFichier)
	if err != nil {
		return fmt.Errorf("création de la ressource pédagogique échouée : %v", err)
	}
	return nil
}

func GetRessources(idEvent int) ([]models.Ressource, error) {
	var ressources []models.Ressource

	query := "SELECT titre, url_fichier FROM ressource_pedagogique WHERE id_event = ? ORDER BY id_ressource ASC"
	rows, err := Db.Query(query, idEvent)
	if err != nil {
		return nil, fmt.Errorf("get ressources : %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var ressource models.Ressource
		err := rows.Scan(&ressource.Titre, &ressource.UrlFichier)
		if err != nil {
			return nil, fmt.Errorf("get ressources scan : %v", err.Error())
		}
		ressources = append(ressources, ressource)
	}
	return ressources, nil
}

func DeleteEvenement(EvenementId int) error {
	result, err := Db.Exec("DELETE FROM pa2026.Evenement WHERE id = ?", EvenementId)
	if err != nil {
		return fmt.Errorf("suppression échouée : %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("aucune Evenement trouvée avec l'id %d", EvenementId)
	}
	return nil
}

func UpdateEvenement(EvenementId int, Evenement models.Evenement) error {
	result, err := Db.Exec("UPDATE pa2026.Evenement SET titre = ?, description = ?, date_debut = ?, date_fin = ?, nb_places = ?, format = ?, lieu = ?, type = ?, prix = ? WHERE id = ?",
		Evenement.Titre, Evenement.Description, Evenement.DateDebut, Evenement.DateFin, Evenement.NbPlaces, Evenement.Format, Evenement.Lieu, Evenement.Type, Evenement.Prix, EvenementId)
	if err != nil {
		return fmt.Errorf("mise à jour échouée : %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("aucune Evenement trouvée avec l'id %d", EvenementId)
	}
	return nil
}

func InscrireClient(idUser int, idEvent int) error {
	var nbPlaces int
	var inscrits int
	var dateDebut string

	err := Db.QueryRow(`SELECT nb_places, date_debut, 
              (SELECT COUNT(*) FROM inscription WHERE id_event = ?) 
              FROM evenement WHERE id = ?`, idEvent, idEvent).Scan(&nbPlaces, &dateDebut, &inscrits)
	if err != nil {
		fmt.Println("Erreur SQL (Select):", err)
		return err
	}

	var estDansLeFutur bool
	err = Db.QueryRow("SELECT (date_debut > NOW()) FROM evenement WHERE id = ?", idEvent).Scan(&estDansLeFutur)
	if err != nil {
		fmt.Println("Erreur lors de la vérification date (MySQL):", err)
	} else if !estDansLeFutur {
		return fmt.Errorf("les inscriptions sont fermées, cet événement est déjà terminé")
	}

	if inscrits >= nbPlaces {
		return fmt.Errorf("plus de place (max: %d)", nbPlaces)
	}

	var check int
	err = Db.QueryRow("SELECT COUNT(*) FROM inscription WHERE id_user = ? AND id_event = ?", idUser, idEvent).Scan(&check)
	if err != nil {
		return err
	}
	if check > 0 {
		return fmt.Errorf("vous êtes déjà inscrit à cet événement")
	}

	_, err = Db.Exec("INSERT INTO inscription (id_user, id_event) VALUES (?, ?)", idUser, idEvent)
	if err != nil {
		fmt.Println("Erreur SQL (Insert):", err)
		return err
	}

	return nil
}

func GetInscrits(idEvent int) ([]models.Inscrit, error) {
	var inscrits []models.Inscrit

	query := `SELECT utilisateur.nom, utilisateur.prenom, utilisateur.email,
	          DATE_FORMAT(inscription.date_inscrip, '%d/%m/%Y a %H:%i') as date_inscrip
	          FROM inscription
	          INNER JOIN utilisateur ON utilisateur.id = inscription.id_user
	          WHERE inscription.id_event = ?
	          ORDER BY inscription.date_inscrip ASC`

	rows, err := Db.Query(query, idEvent)
	if err != nil {
		return nil, fmt.Errorf("get inscrits : %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var inscrit models.Inscrit
		err := rows.Scan(&inscrit.Nom, &inscrit.Prenom, &inscrit.Email, &inscrit.DateInscription)
		if err != nil {
			return nil, fmt.Errorf("get inscrits scan : %v", err.Error())
		}
		inscrits = append(inscrits, inscrit)
	}
	return inscrits, nil
}

func SupprimerInscription(idUser int, idEvent int) error {
	query := "DELETE FROM inscription WHERE id_user = ? AND id_event = ?"
	result, err := Db.Exec(query, idUser, idEvent)
	if err != nil {
		return fmt.Errorf("erreur lors de la suppression SQL : %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("aucune inscription trouvée pour cet utilisateur et cet événement")
	}

	return nil
}
