package bdd

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
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

	// 5. Mise à jour de l'état de l'annonce (pour qu'elle ne soit plus modifiable) 
	_, err = Db.Exec("UPDATE annonce SET statut_vente = 'EN ATTENTE DEPOT' WHERE id = ?", annonceId)
	if err != nil {
		return err
	}

return nil
}

// generateRandomPIN crée un code à 6 chiffres 
func generateRandomPIN() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return fmt.Sprintf("%06d", n)
}

func ConfirmDeposit(pinCode string) error {
    var histId int
    var conteneurId int
    var annonceId int

    // 1. On cherche si ce code PIN existe pour un dépôt qui n'a pas encore eu lieu
    // On récupère l'ID de l'historique, du conteneur et de l'annonce liés
    
    err := Db.QueryRow("SELECT id, conteneur_id, annonce_id FROM historique_conteneurs WHERE code_ouverture = ? AND date_depot_effective IS NULL", pinCode).Scan(&histId, &conteneurId, &annonceId)
    if err != nil {
        return errors.New("code PIN invalide, expiré ou déjà utilisé")
    }

    // 2. On marque la date de dépôt réelle dans l'historique
    _, err = Db.Exec("UPDATE historique_conteneurs SET date_depot_effective = NOW() WHERE id = ?", histId)
    if err != nil {
        return err
    }

    // 3. On fait passer le box à l'état 'OCCUPE' (Inventaire intelligent)
    _, err = Db.Exec("UPDATE box_conteneur SET etat = 'OCCUPE' WHERE id = ?", conteneurId)
    if err != nil {
        return err
    }

    // 4. On met à jour l'annonce : elle est maintenant physiquement en box
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

    // 1. On vérifie si le code-barres existe et si l'objet est bien en box
	 err := Db.QueryRow("SELECT id, conteneur_id, annonce_id FROM historique_conteneurs WHERE code_barre_recuperation = ? AND date_retrait_effective IS NULL", barcode).Scan(&histID, &conteneurId, &annonceId)
    if err != nil {
        return errors.New("code-barres invalide ou objet déjà récupéré")
    }

    // 2. Mise à jour de l'historique : On enregistre qui a récupéré et quand
    _, err = Db.Exec("UPDATE historique_conteneurs SET date_retrait_effective = NOW(), professionnel_id = ? WHERE id = ?", professionnelId, histID)
    if err != nil { return err }

    // 3. Libération du conteneur (Inventaire intelligent)
    _, err = Db.Exec("UPDATE box_conteneur SET etat = 'LIBRE' WHERE id = ?", conteneurId)
    if err != nil { return err }

    // 4. Mise à jour de l'annonce : l'objet est officiellement "Upcyclé"
    _, err = Db.Exec("UPDATE annonce SET statut = 'RECUPERE' WHERE id = ?", annonceId)
    if err != nil { return err }

    // 5. Calcul de l'Upcycling Score (C'est ici que ton algo intervient !)
    err = CalculateAndAddScore(annonceId, professionnelId)
    if err != nil {
        fmt.Println("Erreur lors du calcul du score :", err)
        // On ne bloque pas le processus si le score échoue, on log juste l'erreur
    }

    return nil
}

func CalculateAndAddScore(annonceId int, professionnelId int) error {
    var poids float64
    var materiau string
    var particulierId int

    // 1. Récupérer les infos de l'objet récupéré
    err := Db.QueryRow("SELECT poids, type_materiau, particulier_id FROM annonce WHERE id = ?", annonceId).Scan(&poids, &materiau, &particulierId)
    if err != nil {
        return err
    }

    // 2. Définir les coefficients (Logique métier) ft vinted
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

    // 3. Calcul du score final
    gainScore := poids * coef

    // 4. Mise à jour du score de l'utilisateur en base
    _, err = Db.Exec("UPDATE users SET score = score + ? WHERE id = ?", gainScore, particulierId)
    
    return err
}