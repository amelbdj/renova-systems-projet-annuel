package bdd

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"upcycleconnect/models"
)

func ReserveBox(annonceId int, conteneurId int, particulierId int) error {

    

	var etat string
	err := Db.QueryRow("SELECT etat FROM box_conteneur WHERE id = ?", conteneurId).Scan(&etat)
	if err != nil {
		return err
	}
	if etat != "LIBRE" {
		return errors.New("le conteneur n'est pas disponible")
	}

	pinCode := generateRandomPIN() 
	barcode := fmt.Sprintf("UC-%d-%d", annonceId, conteneurId) // Identifiant unique pour le professionnel

	
	_, err = Db.Exec("INSERT INTO historique_conteneurs (conteneur_id, annonce_id, particulier_id, code_ouverture, code_barre_recuperation, date_reservation) VALUES (?, ?, ?, ?, ?, NOW())", conteneurId, annonceId, particulierId, pinCode, barcode)
	
	if err != nil {
		return err
	}

	_, err = Db.Exec("UPDATE annonce SET statut_vente = 'EN ATTENTE DEPOT' WHERE id = ?", annonceId)
	if err != nil {
		return err
	}

return nil
}

func generateRandomPIN() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return fmt.Sprintf("%06d", n)
}

func ConfirmDeposit(pinCode string) error {
    var histId int
    var conteneurId int
    var annonceId int


    
    err := Db.QueryRow("SELECT id, conteneur_id, annonce_id FROM historique_conteneurs WHERE code_ouverture = ? AND date_depot_effective IS NULL", pinCode).Scan(&histId, &conteneurId, &annonceId)
    if err != nil {
        return errors.New("code PIN invalide, expiré ou déjà utilisé")
    }

    _, err = Db.Exec("UPDATE historique_conteneurs SET date_depot_effective = NOW() WHERE id = ?", histId)
    if err != nil {
        return err
    }

    _, err = Db.Exec("UPDATE box_conteneur SET etat = 'OCCUPE' WHERE id = ?", conteneurId)
    if err != nil {
        return err
    }

    _, err = Db.Exec("UPDATE annonce SET statut = 'EN BOX' WHERE id = ?", annonceId)
    if err != nil {
        return err
    }

    return nil
}

func CollectObject(barcode string, professionnelId int) error {
    var histID int
    var conteneurId int
    var annonceId int

	 err := Db.QueryRow("SELECT id, conteneur_id, annonce_id FROM historique_conteneurs WHERE code_barre_recuperation = ? AND date_retrait_effective IS NULL", barcode).Scan(&histID, &conteneurId, &annonceId)
    if err != nil {
        return errors.New("code-barres invalide ou objet déjà récupéré")
    }

    _, err = Db.Exec("UPDATE historique_conteneurs SET date_retrait_effective = NOW(), professionnel_id = ? WHERE id = ?", professionnelId, histID)
    if err != nil { return err }

    _, err = Db.Exec("UPDATE box_conteneur SET etat = 'LIBRE' WHERE id = ?", conteneurId)
    if err != nil { return err }

    _, err = Db.Exec("UPDATE annonce SET statut = 'RECUPERE' WHERE id = ?", annonceId)
    if err != nil { return err }

    err = CalculateAndAddScore(annonceId, professionnelId)
    if err != nil {
        fmt.Println("Erreur lors du calcul du score :", err)
    }

    return nil
}

func CalculateAndAddScore(annonceId int, professionnelId int) error {
    var poids float64
    var materiau string
    var particulierId int

    err := Db.QueryRow("SELECT poids, type_materiau, particulier_id FROM annonce WHERE id = ?", annonceId).Scan(&poids, &materiau, &particulierId)
    if err != nil {
        return err
    }

    coefficients := map[string]float64{
        "textile":   15.0,
        "metal":     10.0,
        "bois":      5.0,
        "plastique":  8.0,
        "autre":     3.0,
    }

    coef, exists := coefficients[materiau]
    if !exists {
        coef = coefficients["autre"]
    }

    gainScore := poids * coef

    _, err = Db.Exec("UPDATE users SET score = score + ? WHERE id = ?", gainScore, particulierId)
    
    return err
}

func GetBox() ([]models.Box, error) {

	var Boxs []models.Box

	rows, err := Db.Query("SELECT id, localisation, type_materiau_accepte, etat, capacite FROM box_conteneur")

	if err != nil {
				fmt.Println("Erreur lors de l'exécution de la requête : ", err)

		return nil, fmt.Errorf("get Boxs : %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		var Box models.Box

		err := rows.Scan(&Box.Id, 
            &Box.Localisation, 
            &Box.Type,
            &Box.Etat,
            &Box.Capacite)

		if err != nil {
			return nil, fmt.Errorf("get Boxs : %v", err.Error())
		}

		Boxs = append(Boxs, Box)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("get Boxs : %v", err.Error())
	}

	return Boxs, nil
}

func GarbageCollectBox() error { // faire la suppre via cron



	rows, err := Db.Query("SELECT annonce_id, conteneur_id FROM historique_conteneurs WHERE date_reservation < NOW() - INTERVAL 2 DAY AND date_depot_effective IS NULL")
	if err != nil {
		return fmt.Errorf("Cleanup (select): %v", err)
	}
	defer rows.Close()

	for rows.Next() {
	var annonceId int
    var conteneurId int
    err := rows.Scan(&annonceId, &conteneurId)

    if err != nil {
    continue 
    }       

		_, err = Db.Exec("UPDATE box_conteneur SET etat = 'LIBRE' WHERE id = ?", conteneurId)
		if err != nil {
			fmt.Printf("Erreur libération box %d: %v", conteneurId, err)
		}

		_, err = Db.Exec("UPDATE annonce SET statut = 'En vente' WHERE id = ?", annonceId)
		if err != nil {
			fmt.Printf("Erreur remise en vente annonce %d: %v", annonceId, err)
		}

		_, err = Db.Exec("UPDATE historique_conteneurs SET date_reservation = NULL WHERE annonce_id = ? AND date_depot_effective IS NULL", annonceId)
            
		fmt.Printf("Nettoyage réussi pour l'annonce %d (Box %d libéré)", annonceId, conteneurId)
	}

	return nil
}

func CreateBox(localisation string, boxType string, capacite int) error {
    _, err := Db.Exec("INSERT INTO box_conteneur (localisation, type_materiau_accepte, etat, capacite) VALUES (?, ?, 'LIBRE', ?)", localisation, boxType, capacite)
    
    if err != nil {
        return fmt.Errorf("CreateBox: %v", err)
    }
    return nil
}
