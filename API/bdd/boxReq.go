package bdd

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"upcycleconnect/models"
)


func ReserveBox(annonceId int, conteneurId int, particulierId int) error {
	var boxId int

	err := Db.QueryRow("SELECT id FROM box WHERE id_conteneur = ? AND statut = 'libre' LIMIT 1", conteneurId).Scan(&boxId)
	if err != nil {
		return errors.New("aucun casier n'est disponible dans ce conteneur")
	}

	pinCode := generateRandomPIN()
	barcode := fmt.Sprintf("UC-%d-%d", annonceId, boxId)

	_, err = Db.Exec("INSERT INTO historique_conteneurs (box_id, annonce_id, particulier_id, code_ouverture, code_barre_recuperation, date_reservation) VALUES (?, ?, ?, ?, ?, NOW())", boxId, annonceId, particulierId, pinCode, barcode)
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

	err := Db.QueryRow("SELECT id, box_id, annonce_id FROM historique_conteneurs WHERE code_ouverture = ? AND date_depot_effective IS NULL", pinCode).Scan(&histId, &boxId, &annonceId)
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

	_, err = Db.Exec("UPDATE annonce SET statut = 'EN BOX' WHERE id = ?", annonceId)
	if err != nil {
		return err
	}

	return nil
}

func CollectObject(barcode string, professionnelId int) error {
	var histID int
	var boxId int
	var annonceId int

	err := Db.QueryRow("SELECT id, box_id, annonce_id FROM historique_conteneurs WHERE code_barre_recuperation = ? AND date_retrait_effective IS NULL", barcode).Scan(&histID, &boxId, &annonceId)
	if err != nil {
		return errors.New("code-barres invalide ou objet déjà récupéré")
	}

	_, err = Db.Exec("UPDATE historique_conteneurs SET date_retrait_effective = NOW(), professionnel_id = ? WHERE id = ?", professionnelId, histID)
	if err != nil {
		return err
	}

	_, err = Db.Exec("UPDATE box SET statut = 'libre', code_secret = NULL WHERE id = ?", boxId)
	if err != nil {
		return err
	}

	_, err = Db.Exec("UPDATE annonce SET statut = 'RECUPERE' WHERE id = ?", annonceId)
	if err != nil {
		return err
	}

	err = CalculateAndAddScore(annonceId, professionnelId)
	if err != nil {
		fmt.Println("Erreur lors du calcul du score :", err)
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
		JOIN box b ON h.box_id = b.id
		JOIN conteneur c ON b.id_conteneur = c.id
		WHERE h.particulier_id = ? AND h.date_retrait_effective IS NULL`

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
			"lieu":         nomConteneur + " - " + adresse,
			"numero_box":   fmt.Sprintf("Casier n°%d", numBox),
			"code_pin":     code,
			"barcode":      barcode,
			"objet":        titre,
			"date":         date,
			"etat":         etat,
		}
		reservations = append(reservations, res)
	}
	return reservations, nil
}



// Remplace ton ancien GetBox(). Ça renvoie la liste des meubles pour l'Admin.
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
		err := rows.Scan(&c.ID, &c.Nom, &c.Adresse, &c.TotalBoxes)
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

	err := Db.QueryRow("SELECT poids, type_materiau, particulier_id FROM annonce WHERE id = ?", annonceId).Scan(&poids, &materiau, &particulierId)
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
	_, err = Db.Exec("UPDATE users SET score = score + ? WHERE id = ?", gainScore, particulierId)
	return err
}

func GarbageCollectBox() error { 
	rows, err := Db.Query("SELECT annonce_id, box_id FROM historique_conteneurs WHERE date_reservation < NOW() - INTERVAL 2 DAY AND date_depot_effective IS NULL")
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

		_, err = Db.Exec("UPDATE annonce SET statut = 'En vente' WHERE id = ?", annonceId)
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