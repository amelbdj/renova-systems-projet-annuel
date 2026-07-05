package main

import (
	"flag"
	"fmt"
	"log"
	"upcycleconnect/bdd"
)

type demoAnnonce struct {
	UserID      int
	CategorieID int
	Titre       string
	Description string
	Type        string
	Prix        float64
	Validation  string
	CodePostal  string
	Ville       string
	Projet      string
	Etat        string
	Poids       float64
	Quantite    int
	Sponsored   int
	StatutVente string
}

func main() {
	deleteOnly := flag.Bool("delete", false, "supprime les annonces demo et retire les Stripe ID demo")
	flag.Parse()

	db := bdd.NewDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	if *deleteOnly {
		_, err = tx.Exec("DELETE FROM annonce WHERE titre LIKE '[DEMO] %'")
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}

		_, err = tx.Exec(`
			UPDATE utilisateur
			SET stripe_account_id = NULL,
			    stripe_verif_completed = 0,
			    stripe_customer_id = NULL,
			    est_premium = 0,
			    plan_abo = NULL
			WHERE email IN (
				'test.client@renova.test',
				'test.pro@renova.test',
				'test.salarie@renova.test'
			)
		`)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}

		if err = tx.Commit(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Donnees demo supprimees.")
		return
	}

	// Comptes demo principaux.
	_, err = tx.Exec(`
		UPDATE utilisateur
		SET stripe_customer_id = 'cus_demo_client_0001',
		    stripe_account_id = 'acct_demo_client_0001',
		    stripe_verif_completed = 1,
		    validation = 'Validé'
		WHERE email = 'test.client@renova.test'
	`)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.Exec(`
		UPDATE utilisateur
		SET stripe_customer_id = 'cus_demo_pro_0001',
		    stripe_account_id = 'acct_demo_pro_0001',
		    stripe_verif_completed = 1,
		    est_premium = 1,
		    plan_abo = 'pro',
		    nom_entreprise = 'Atelier Demo Upcycle',
		    siret = '12345678900011',
		    validation = 'Validé'
		WHERE email = 'test.pro@renova.test'
	`)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.Exec(`
		UPDATE utilisateur
		SET stripe_account_id = 'acct_demo_salarie_0001',
		    stripe_verif_completed = 1,
		    validation = 'Validé'
		WHERE email = 'test.salarie@renova.test'
	`)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	// On nettoie les anciennes annonces demo pour pouvoir relancer le script.
	_, err = tx.Exec("DELETE FROM annonce WHERE titre LIKE '[DEMO] %'")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	annonces := []demoAnnonce{
		{
			UserID:      38,
			CategorieID: 2,
			Titre:       "[DEMO] Lot de chutes de bois atelier",
			Description: "Chutes de bois propres, parfaites pour fabriquer des petites etageres ou objets decoratifs.",
			Type:        "materiau",
			Prix:        18,
			Validation:  "Validé",
			CodePostal:  "75011",
			Ville:       "Paris",
			Projet:      "Creation d'etageres murales ou supports de plantes.",
			Etat:        "Bon etat",
			Poids:       12.5,
			Quantite:    8,
			Sponsored:   0,
			StatutVente: "EN VENTE",
		},
		{
			UserID:      38,
			CategorieID: 4,
			Titre:       "[DEMO] Anciennes poignees en metal",
			Description: "Poignees et petites pieces metalliques issues d'un meuble demonte.",
			Type:        "materiau",
			Prix:        9,
			Validation:  "En attente",
			CodePostal:  "75010",
			Ville:       "Paris",
			Projet:      "Reemploi en mobilier ou decoration industrielle.",
			Etat:        "Usage",
			Poids:       3.2,
			Quantite:    20,
			Sponsored:   0,
			StatutVente: "EN VENTE",
		},
		{
			UserID:      38,
			CategorieID: 6,
			Titre:       "[DEMO] Chaise cannee a reparer",
			Description: "Chaise ancienne avec assise abimee, interessante pour atelier de restauration.",
			Type:        "objet",
			Prix:        14,
			Validation:  "Rejeté",
			CodePostal:  "93100",
			Ville:       "Montreuil",
			Projet:      "Restauration de mobilier.",
			Etat:        "Pour pieces",
			Poids:       5,
			Quantite:    1,
			Sponsored:   0,
			StatutVente: "EN VENTE",
		},
		{
			UserID:      41,
			CategorieID: 2,
			Titre:       "[DEMO] Presentoir bois recycle premium",
			Description: "Presentoir fabrique a partir de bois reutilise, propose par un professionnel premium.",
			Type:        "objet",
			Prix:        65,
			Validation:  "Validé",
			CodePostal:  "69002",
			Ville:       "Lyon",
			Projet:      "Mise en avant boutique responsable.",
			Etat:        "Bon etat",
			Poids:       9.8,
			Quantite:    2,
			Sponsored:   1,
			StatutVente: "EN VENTE",
		},
		{
			UserID:      41,
			CategorieID: 1,
			Titre:       "[DEMO] Lot textile pour creation artisanale",
			Description: "Lot de tissus recuperes, trie et prepare pour couture ou accessoires.",
			Type:        "materiau",
			Prix:        22,
			Validation:  "Validé",
			CodePostal:  "69007",
			Ville:       "Lyon",
			Projet:      "Sacs, pochettes et accessoires zero dechet.",
			Etat:        "Bon etat",
			Poids:       6.4,
			Quantite:    12,
			Sponsored:   0,
			StatutVente: "EN ATTENTE DEPOT",
		},
	}

	req, err := tx.Prepare(`
		INSERT INTO annonce
		(id_user, id_categorie, titre, description, type, prix, statut_validation,
		 code_postal, ville, projet_potentiel, etat, poids_kg, quantite, image,
		 is_sponsored, commission_prelevee, photo_url, statut_vente)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, '', ?, 0, '', ?)
	`)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	defer req.Close()

	for _, a := range annonces {
		_, err = req.Exec(
			a.UserID,
			a.CategorieID,
			a.Titre,
			a.Description,
			a.Type,
			a.Prix,
			a.Validation,
			a.CodePostal,
			a.Ville,
			a.Projet,
			a.Etat,
			a.Poids,
			a.Quantite,
			a.Sponsored,
			a.StatutVente,
		)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Donnees demo ajoutees.")
	fmt.Println("Comptes Stripe demo : test.client@renova.test, test.pro@renova.test, test.salarie@renova.test")
	fmt.Println("Annonces demo ajoutees :", len(annonces))
}
