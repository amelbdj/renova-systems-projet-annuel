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
	err := Db.QueryRow("SELECT etat FROM box_conteneur WHERE id_box = ?", conteneurId).Scan(&etat)
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

	_, err = Db.Exec("UPDATE box_conteneur SET etat = 'RESERVEE' WHERE id_box = ?", conteneurId)
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
	if err != nil {
		return err
	}

	// 3. Libération du conteneur (Inventaire intelligent)
	_, err = Db.Exec("UPDATE box_conteneur SET etat = 'LIBRE' WHERE id = ?", conteneurId)
	if err != nil {
		return err
	}

	// 4. Mise à jour de l'annonce : l'objet est officiellement "Upcyclé"
	_, err = Db.Exec("UPDATE annonce SET statut = 'RECUPERE' WHERE id = ?", annonceId)
	if err != nil {
		return err
	}

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
		"plastique": 8.0,
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

	// On utilise INTERVAL 2 DAY (48h) pour la rotation des box

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

		// A. Libérer le BOX physique
		_, err = Db.Exec("UPDATE box_conteneur SET etat = 'LIBRE' WHERE id = ?", conteneurId)
		if err != nil {
			fmt.Printf("Erreur libération box %d: %v", conteneurId, err)
		}

		// B. Remettre l'ANNONCE en ligne (pour qu'un autre pro puisse l'acheter)
		_, err = Db.Exec("UPDATE annonce SET statut = 'En vente' WHERE id = ?", annonceId)
		if err != nil {
			fmt.Printf("Erreur remise en vente annonce %d: %v", annonceId, err)
		}

		// C. Marquer l'HISTORIQUE comme expiré (au lieu de DELETE pour garder la trace)
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

func GetUserReservations(userID int) ([]map[string]interface{}, error) {
	query := `
        SELECT 
            h.code_ouverture, 
            h.code_barre_recuperation, 
            h.date_reservation,
            a.titre,
            b.id_box,
            b.localisation,
            b.etat
        FROM historique_conteneurs h
        JOIN annonce a ON h.annonce_id = a.id
        JOIN box_conteneur b ON h.conteneur_id = b.id_box
        WHERE h.particulier_id = ? AND h.date_retrait_effective IS NULL`

	rows, err := Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []map[string]interface{}
	for rows.Next() {
		var code, barcode, date, titre, loc, etat string
		var idBox int
		rows.Scan(&code, &barcode, &date, &titre, &idBox, &loc, &etat)

		res := map[string]interface{}{
			"id_box":       fmt.Sprintf("BOX-%03d", idBox),
			"code_pin":     code,
			"barcode":      barcode,
			"objet":        titre,
			"date":         date,
			"localisation": loc,
			"etat":         etat,
		}
		reservations = append(reservations, res)
	}
	return reservations, nil
}
