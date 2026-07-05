package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	"upcycleconnect/bdd"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	nbUsers := flag.Int("count", 15000, "nombre d'utilisateurs a creer")
	password := flag.String("password", "Test123!", "mot de passe des utilisateurs de test")
	deleteOnly := flag.Bool("delete", false, "supprime les utilisateurs de test")
	flag.Parse()

	db := bdd.NewDB()
	defer db.Close()

	if *deleteOnly {
		_, err := db.Exec("DELETE FROM utilisateur WHERE email LIKE 'user_seed_%@renova.test'")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Utilisateurs de test supprimes.")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(*password), 10)
	if err != nil {
		log.Fatal(err)
	}

	debut := time.Now()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	req, err := tx.Prepare(`
		INSERT INTO utilisateur
		(role, nom, prenom, email, mot_de_passe, tutoriel_vu, score, validation, nom_entreprise, siret)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		nom = VALUES(nom),
		prenom = VALUES(prenom),
		mot_de_passe = VALUES(mot_de_passe)
	`)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	defer req.Close()

	roles := []string{"Utilisateur", "Utilisateur", "Utilisateur", "Pro", "Salarié"}

	for i := 1; i <= *nbUsers; i++ {
		role := roles[i%len(roles)]
		nom := fmt.Sprintf("SeedNom%d", i)
		prenom := fmt.Sprintf("SeedPrenom%d", i)
		email := fmt.Sprintf("user_seed_%05d@renova.test", i)
		score := i % 500
		nomEntreprise := ""
		siret := ""

		if role == "Pro" {
			nomEntreprise = fmt.Sprintf("Entreprise Seed %d", i)
			siret = fmt.Sprintf("123456789%05d", i%100000)
		}

		_, err = req.Exec(role, nom, prenom, email, string(hash), 1, score, "Validé", nomEntreprise, siret)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}

		if i%1000 == 0 {
			fmt.Println(i, "utilisateurs crees...")
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Seeding termine :", *nbUsers, "utilisateurs en", time.Since(debut))
	fmt.Println("Mot de passe :", *password)
}
