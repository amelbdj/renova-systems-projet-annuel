package bdd

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"upcycleconnect/models"
)

func ReserveBox(annonceId int, conteneurId int, vendeurId int) error {
	var boxId int

	err := Db.QueryRow("SELECT id FROM box WHERE id_conteneur = ? AND statut = 'libre' LIMIT 1", conteneurId).Scan(&boxId)
	if err != nil {
		return errors.New("aucun casier n'est disponible dans ce conteneur")
	}

	var acheteurId int
	err = Db.QueryRow("SELECT id_acheteur FROM `order` WHERE id_annonce = ? ORDER BY date_commande DESC LIMIT 1", annonceId).Scan(&acheteurId)
	if err != nil {
		return errors.New("impossible de trouver l'acheteur pour cette annonce")
	}

	pinCode := generateRandomPIN()
	barcode := fmt.Sprintf("UC-%d-%d", annonceId, boxId)

	_, err = Db.Exec(`
		INSERT INTO historique_conteneurs 
		(conteneur_id, annonce_id, vendeur_id, acheteur_id, code_ouverture, code_barre_recuperation, date_reservation) 
		VALUES (?, ?, ?, ?, ?, ?, NOW())`,
		boxId, annonceId, vendeurId, acheteurId, pinCode, barcode)

	if err != nil {
		return err
	}

	_, err = Db.Exec("UPDATE annonce SET statut_vente = 'EN ATTENTE DEPOT' WHERE id = ?", annonceId)
	if err != nil {
		return err
	}

	_, err = Db.Exec("UPDATE box SET statut = 'reservee', code_secret = ? WHERE id = ?", pinCode, boxId)
	if err != nil {
		return err
	}

	return nil
}

func ConfirmDeposit(pinCode string) error {
	var histId int
	var boxId int
	var annonceId int

	err := Db.QueryRow("SELECT id, conteneur_id, annonce_id FROM historique_conteneurs WHERE code_ouverture = ? AND date_depot_effective IS NULL", pinCode).Scan(&histId, &boxId, &annonceId)
	if err != nil {
		return errors.New("code PIN invalide, expiré ou déjà utilisé")
	}

	_, err = Db.Exec("UPDATE historique_conteneurs SET date_depot_effective = NOW() WHERE id = ?", histId)
	if err != nil {
		return err
	}

	_, err = Db.Exec("UPDATE box SET statut = 'occupee' WHERE id = ?", boxId)
	if err != nil {
		return err
	}

	_, err = Db.Exec("UPDATE annonce SET statut_vente = 'EN BOX' WHERE id = ?", annonceId)
	if err != nil {
		return err
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
			b.numero,
			c.nom,
			c.adresse,
			b.statut
		FROM historique_conteneurs h
		JOIN annonce a ON h.annonce_id = a.id
		JOIN box b ON h.conteneur_id = b.id
		JOIN conteneur c ON b.id_conteneur = c.id
		WHERE h.vendeur_id = ? AND h.date_depot_effective IS NULL`

	rows, err := Db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []map[string]interface{}
	for rows.Next() {
		var code, barcode, date, titre, nomConteneur, adresse, etat string
		var numBox int

		err := rows.Scan(&code, &barcode, &date, &titre, &numBox, &nomConteneur, &adresse, &etat)
		if err != nil {
			continue
		}

		res := map[string]interface{}{
			"lieu":       nomConteneur + " - " + adresse,
			"numero_box": fmt.Sprintf("Casier n°%d", numBox),
			"code_pin":   code,
			"barcode":    barcode,
			"objet":      titre,
			"date":       date,
			"etat":       etat,
		}
		reservations = append(reservations, res)
	}
	return reservations, nil
}

func GetUserPickups(acheteurID int) ([]map[string]interface{}, error) {

	query := `
		SELECT 
			h.code_ouverture, 
			h.code_barre_recuperation, 
			h.date_depot_effective,
			a.titre,
			b.numero,
			c.nom,
			c.adresse,
			a.statut_vente
		FROM historique_conteneurs h
		JOIN annonce a ON h.annonce_id = a.id
		JOIN box b ON h.conteneur_id = b.id
		JOIN conteneur c ON b.id_conteneur = c.id
		WHERE h.acheteur_id = ? 
		  AND h.date_depot_effective IS NOT NULL 
		  AND h.date_retrait_effective IS NULL`

	rows, err := Db.Query(query, acheteurID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pickups []map[string]interface{}
	for rows.Next() {
		var codePIN, barcode, dateDepot, titre, nomConteneur, adresse, statutVente string
		var numBox int

		err := rows.Scan(&codePIN, &barcode, &dateDepot, &titre, &numBox, &nomConteneur, &adresse, &statutVente)
		if err != nil {
			continue
		}

		pickup := map[string]interface{}{
			"lieu":         nomConteneur + " - " + adresse,
			"numero_box":   fmt.Sprintf("Casier n°%d", numBox),
			"code_pin":     codePIN,
			"barcode":      barcode,
			"objet":        titre,
			"date_depot":   dateDepot,
			"statut_vente": statutVente,
			"date":         dateDepot,
			"etat":         statutVente,
		}
		pickups = append(pickups, pickup)
	}
	return pickups, nil
}

func GetConteneursAdmin() ([]models.ConteneurAvecStats, error) {
	var conteneurs []models.ConteneurAvecStats

	query := `
        SELECT c.id, c.nom, c.adresse, COUNT(b.id) as total_boxes 
        FROM conteneur c 
        LEFT JOIN box b ON c.id = b.id_conteneur 
        GROUP BY c.id`

	rows, err := Db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erreur GetConteneursAdmin : %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var c models.ConteneurAvecStats
		err := rows.Scan(&c.Id, &c.Nom, &c.Adresse, &c.TotalBoxes)
		if err != nil {
			return nil, fmt.Errorf("erreur scan : %v", err)
		}
		conteneurs = append(conteneurs, c)
	}
	return conteneurs, nil
}

func generateRandomPIN() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return fmt.Sprintf("%06d", n)
}

func CalculateAndAddScore(annonceId int, professionnelId int) error {
	var poids float64
	var materiau string
	var particulierId int

	err := Db.QueryRow("SELECT poids_kg, id_user FROM pa2026.annonce WHERE id = ?", annonceId).Scan(&poids, &particulierId)
	if err != nil {
		return err
	}

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

	gainScore := poids * coef
	_, err = Db.Exec("UPDATE pa2026.utilisateur SET score = score + ? WHERE id = ?", gainScore, particulierId)
	return err
}

func GarbageCollectBox() error {

	rows, err := Db.Query("SELECT annonce_id, conteneur_id FROM historique_conteneurs WHERE date_reservation < NOW() - INTERVAL 2 DAY AND date_depot_effective IS NULL")
	if err != nil {
		return fmt.Errorf("Cleanup (select): %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var annonceId int
		var boxId int
		err := rows.Scan(&annonceId, &boxId)
		if err != nil {
			continue
		}

		_, err = Db.Exec("UPDATE box SET statut = 'libre', code_secret = NULL WHERE id = ?", boxId)
		if err != nil {
			fmt.Printf("Erreur libération box %d: %v", boxId, err)
		}

		_, err = Db.Exec("UPDATE annonce SET statut_vente = 'EN VENTE' WHERE id = ?", annonceId)
		if err != nil {
			fmt.Printf("Erreur remise en vente annonce %d: %v", annonceId, err)
		}

		_, err = Db.Exec("UPDATE historique_conteneurs SET date_reservation = NULL WHERE annonce_id = ? AND date_depot_effective IS NULL", annonceId)

		fmt.Printf("Nettoyage réussi pour l'annonce %d (Box %d libéré)\n", annonceId, boxId)
	}

	return nil
}

func CreateConteneurAvecBox(nom string, adresse string, nombreDeBoxs int) error {

	res, err := Db.Exec("INSERT INTO conteneur (nom, adresse) VALUES (?, ?)", nom, adresse)
	if err != nil {
		return fmt.Errorf("erreur insert conteneur: %v", err)
	}

	conteneurID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("erreur récupération ID: %v", err)
	}

	for i := 1; i <= nombreDeBoxs; i++ {
		_, err := Db.Exec(
			"INSERT INTO box (id_conteneur, numero, taille, statut) VALUES (?, ?, ?, ?)",
			conteneurID, i, "M", "libre",
		)

		if err != nil {
			return fmt.Errorf("erreur insert box n°%d: %v", i, err)
		}
	}

	return nil
}

func GetBoxesByConteneurID(conteneurID string) ([]models.Box, error) {
	var boxes []models.Box

	rows, err := Db.Query(`SELECT id, id_conteneur, numero, taille, statut, code_secret 
              FROM box 
              WHERE id_conteneur = ? 
              ORDER BY numero ASC`, conteneurID)
	if err != nil {
		return nil, fmt.Errorf("erreur requête GetBoxes: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Box

		err := rows.Scan(
			&b.Id,
			&b.IdConteneur,
			&b.Numero,
			&b.Taille,
			&b.Statut,
			&b.CodeSecret,
		)

		if err != nil {
			return nil, fmt.Errorf("erreur scan box: %v", err)
		}

		boxes = append(boxes, b)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erreur itération boxes: %v", err)
	}

	if boxes == nil {
		boxes = []models.Box{}
	}

	return boxes, nil
}

func AddSingleBoxToConteneur(conteneurID int, taille string) error {
	var maxNumero *int
	err := Db.QueryRow("SELECT MAX(numero) FROM box WHERE id_conteneur = ?", conteneurID).Scan(&maxNumero)

	if err != nil {
		return fmt.Errorf("erreur recherche max numero: %v", err)
	}

	nouveauNumero := 1
	if maxNumero != nil {
		nouveauNumero = *maxNumero + 1
	}

	_, err = Db.Exec(
		"INSERT INTO box (id_conteneur, numero, taille, statut) VALUES (?, ?, ?, 'libre')",
		conteneurID, nouveauNumero, taille,
	)

	return err
}

func UpdateBoxStatus(boxID int, statut string) error {
	// On normalise le statut en minuscules pour rester cohérent (libre, occupee, reservee, maintenance)
	statut = strings.ToLower(strings.TrimSpace(statut))
	_, err := Db.Exec("UPDATE box SET statut = ? WHERE id = ?", statut, boxID)
	return err
}

func CollectObject(pinCode string, professionnelId int) error {
	var histID int
	var boxId int
	var annonceId int

	err := Db.QueryRow("SELECT id, conteneur_id, annonce_id FROM historique_conteneurs WHERE code_ouverture = ? AND date_retrait_effective IS NULL", pinCode).Scan(&histID, &boxId, &annonceId)
	if err != nil {
		return errors.New("code PIN de retrait invalide ou objet déjà récupéré")
	}

	_, err = Db.Exec("UPDATE historique_conteneurs SET date_retrait_effective = NOW(), professionnel_id = ? WHERE id = ?", professionnelId, histID)
	if err != nil {
		return err
	}

	_, err = Db.Exec("UPDATE box SET statut = 'libre', code_secret = NULL WHERE id = ?", boxId)
	if err != nil {
		return err
	}

	_, err = Db.Exec("UPDATE annonce SET statut_vente = 'RECUPERE' WHERE id = ?", annonceId)
	if err != nil {
		return err
	}

	err = CalculateAndAddScore(annonceId, professionnelId)
	if err != nil {
		fmt.Println("Erreur lors du calcul du score :", err)
	}

	return nil
}

func SimulateHardwareDeposit(pinCode string) error {
	var histID int
	var boxId int
	var annonceId int

	err := Db.QueryRow("SELECT id, conteneur_id, annonce_id FROM historique_conteneurs WHERE code_ouverture = ? AND date_depot_effective IS NULL", pinCode).Scan(&histID, &boxId, &annonceId)
	if err != nil {
		return errors.New("code PIN invalide ou objet déjà déposé")
	}

	_, err = Db.Exec("UPDATE historique_conteneurs SET date_depot_effective = NOW() WHERE id = ?", histID)
	if err != nil {
		return err
	}

	_, err = Db.Exec("UPDATE box SET statut = 'occupee' WHERE id = ?", boxId)
	if err != nil {
		return err
	}

	_, err = Db.Exec("UPDATE annonce SET statut_vente = 'EN ATTENTE DE RECUPERATION' WHERE id = ?", annonceId)
	if err != nil {
		return err
	}

	return nil
}

func SimulateHardwareWithdrawal(codeBarre string) error {
	var histID int
	var boxId int
	var annonceId int

	err := Db.QueryRow("SELECT id, conteneur_id, annonce_id FROM historique_conteneurs WHERE code_barre_recuperation = ? AND date_depot_effective IS NOT NULL AND date_retrait_effective IS NULL", codeBarre).Scan(&histID, &boxId, &annonceId)
	if err != nil {
		return errors.New("code-barre de retrait invalide, objet pas encore déposé ou déjà récupéré")
	}

	_, err = Db.Exec("UPDATE historique_conteneurs SET date_retrait_effective = NOW() WHERE id = ?", histID)
	if err != nil {
		return err
	}

	_, err = Db.Exec("UPDATE box SET statut = 'libre', code_secret = NULL WHERE id = ?", boxId)
	if err != nil {
		return err
	}

	_, err = Db.Exec("UPDATE annonce SET statut_vente = 'VENDU' WHERE id = ?", annonceId)
	if err != nil {
		return err
	}

	return nil
}
