package admin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"upcycleconnect/bdd"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/webhook"
)

// Utilise bien le secret généré par ta commande "stripe listen"
const WebhookSecret = "whsec_ca8df90af4b7eb3045da3e1ab058edeb4ea8597a2d49b88adec258b0037f4fbc"

func StripeWebhookHandler(w http.ResponseWriter, r *http.Request) {


	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("❌ ERREUR : Impossible de lire le corps de la requête", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	event, err := webhook.ConstructEventWithOptions(payload, r.Header.Get("Stripe-Signature"), WebhookSecret, webhook.ConstructEventOptions{
		IgnoreAPIVersionMismatch: true,
	})
	if err != nil {
		fmt.Println("ERREUR DE SIGNATURE : Vérifie ton WebhookSecret", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("SIGNATURE VALIDE Événement :", event.Type)

	// Cas 1 : Mise à jour d'un compte Stripe Connect
	if event.Type == "account.updated" {
		var account stripe.Account
		err := json.Unmarshal(event.Data.Raw, &account)
		if err == nil && account.PayoutsEnabled {
			// Mise à jour du statut de l'utilisateur dans la table 'utilisateur'
			_, err := bdd.Db.Exec("UPDATE utilisateur SET stripe_verif_completed = 1 WHERE stripe_account_id = ?", account.ID)
			if err != nil {
				fmt.Println("ERREUR SQL (account.updated) :", err)
			}
		}
	}

	// Cas 2 : Succès d'un paiement (Checkout Session)
	if event.Type == "checkout.session.completed" {
		var session stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &session)
		if err != nil {
			fmt.Println("ERREUR JSON :", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Récupération des données envoyées lors de la création de la session
		idEventStr := session.Metadata["id_event"]
		idUserStr := session.Metadata["id_user"]
		idEvent, _ := strconv.Atoi(idEventStr)
		idUser, _ := strconv.Atoi(idUserStr)

		// Données financières
		stripeID := session.ID
		montantTotal := float64(session.AmountTotal) / 100.0

		fmt.Printf("Paiement reçu ! User: %d, Event: %d, Montant: %.2f€\n", idUser, idEvent, montantTotal)

		// Insertion dans la table 'inscription' 
		_, err = bdd.Db.Exec("INSERT INTO inscription (id_user, id_event) VALUES (?, ?)", idUser, idEvent)
		if err != nil {
			fmt.Println("ERREUR SQL (Table inscription) :", err)
		} else {
			fmt.Println("Inscription enregistrée avec succès !")
		}

		// On commence par créer une commande dans la table 'order'
		res, err := bdd.Db.Exec("INSERT INTO `order` (id_acheteur, id_annonce, montant_total, commission) VALUES (?, ?, ?, ?)",
			idUser, idEvent, montantTotal, 0.0) // On utilise id_event comme id_annonce ici pour simplifier

		if err != nil {
            fmt.Println("❌ ERREUR SQL (Table order) :", err)
		} else {
			// On récupère l'ID de la commande pour l'associer au paiement
			lastID, _ := res.LastInsertId()

			_, err = bdd.Db.Exec("INSERT INTO paiement(id_commande, stripe_id, statut) VALUES (?, ?, ?)",
				lastID, stripeID, "succeeded")

			if err != nil {
				fmt.Println("ERREUR SQL (Table paiement) :", err)
			} else {
				fmt.Println("Paiement stocké en base de données !")
			}
		}
	}

	w.WriteHeader(http.StatusOK)
}