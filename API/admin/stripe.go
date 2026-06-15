package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"upcycleconnect/bdd"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/account"
	"github.com/stripe/stripe-go/v81/accountlink"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"github.com/stripe/stripe-go/v81/paymentintent"
)

const StripeSecretKey = "sk_test_51TNFHBHbaxF1KOTtH89RRHNJQSQXSPVtOHMJDHicr1LW4XYeY4ZC6nYWwzVbDvFUUI58YA7KlJs9BiUyP5zD4XU300gaAUPVpI"

func ConnectToStripe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	stripe.Key = StripeSecretKey

	userIDStr := r.URL.Query().Get("id")
	userID, _ := strconv.Atoi(userIDStr)

	var stripeID string
	bdd.Db.QueryRow("SELECT stripe_account_id FROM utilisateur WHERE id = ?", userID).Scan(&stripeID)

	if stripeID == "" {
		params := &stripe.AccountParams{
			Type: stripe.String(string(stripe.AccountTypeExpress)),
			Capabilities: &stripe.AccountCapabilitiesParams{
				Transfers: &stripe.AccountCapabilitiesTransfersParams{
					Requested: stripe.Bool(true),
				},
			},
		}

		acc, err := account.New(params)
		if err != nil {
			http.Error(w, "Erreur création compte Stripe", http.StatusInternalServerError)
			return
		}
		stripeID = acc.ID

		bdd.Db.Exec("UPDATE utilisateur SET stripe_account_id = ? WHERE id = ?", stripeID, userID)
	}

	linkParams := &stripe.AccountLinkParams{
		Account:    stripe.String(stripeID),
		RefreshURL: stripe.String("http://127.0.0.1:5500/Frontend/profil.html"),
		ReturnURL:  stripe.String("http://127.0.0.1:5500/Frontend/profil.html?stripe=success"),
		Type:       stripe.String("account_onboarding"),
	}

	link, err := accountlink.New(linkParams)
	if err != nil {
		http.Error(w, "Erreur création lien Stripe", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"stripe_id": stripeID,
		"url":       link.URL,
	})
}

func PaymentAnnonce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	annonceID, _ := strconv.Atoi(r.URL.Query().Get("annonce_id"))
	buyerID, _ := strconv.Atoi(r.URL.Query().Get("buyer_id"))

	var buyerStripeID string
	err := bdd.Db.QueryRow("SELECT stripe_account_id FROM utilisateur WHERE id = ?", buyerID).Scan(&buyerStripeID)

	if err != nil || buyerStripeID == "" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "BUYER_STRIPE_NOT_CONFIGURED")
		return
	}

	var prix float64
	var titre, stripeAccountIDSeller string
	query := `
        SELECT a.prix, a.titre, u.stripe_account_id 
        FROM pa2026.annonce a 
        JOIN pa2026.utilisateur u ON a.id_user = u.id 
        WHERE a.id = ?`

	err = bdd.Db.QueryRow(query, annonceID).Scan(&prix, &titre, &stripeAccountIDSeller)
	if err != nil {
		http.Error(w, "Annonce introuvable", http.StatusNotFound)
		return
	}

	unitAmount := int64(prix * 100)
	commission := (unitAmount * 5) / 100

	commissionAppli := commission
	stripe.Key = StripeSecretKey
	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("eur"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(titre),
					},
					UnitAmount: stripe.Int64(unitAmount),
				},
				Quantity: stripe.Int64(1),
			},
		},
		PaymentIntentData: &stripe.CheckoutSessionPaymentIntentDataParams{
			ApplicationFeeAmount: stripe.Int64(commissionAppli),
			TransferData: &stripe.CheckoutSessionPaymentIntentDataTransferDataParams{
				Destination: stripe.String(stripeAccountIDSeller),
			},
		},
		SuccessURL: stripe.String("http://127.0.0.1:5500/Frontend/oneAnnonce.html?id=" + strconv.Itoa(annonceID) + "&buyer_id=" + strconv.Itoa(buyerID) + "&payment=success"),
		CancelURL:  stripe.String("http://127.0.0.1:5500/Frontend/oneAnnonce.html?id=" + strconv.Itoa(annonceID)),
	}

	s, err := session.New(params)
	if err != nil {
		http.Error(w, "Erreur Stripe: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"url": s.URL})
}

// Nouvelle route pour l'application Android
func PaymentIntentMobile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	annonceID, _ := strconv.Atoi(r.URL.Query().Get("annonce_id"))
	var prix float64
	var stripeAccountIDSeller string
	
	query := `
        SELECT a.prix, u.stripe_account_id 
        FROM pa2026.annonce a 
        JOIN pa2026.utilisateur u ON a.id_user = u.id 
        WHERE a.id = ?`

	err := bdd.Db.QueryRow(query, annonceID).Scan(&prix, &stripeAccountIDSeller)
	if err != nil || stripeAccountIDSeller == "" {
		http.Error(w, "Vendeur non configuré pour Stripe", http.StatusBadRequest)
		return
	}

	stripe.Key = StripeSecretKey
	unitAmount := int64(prix * 100)
	commission := (unitAmount * 5) / 100

	// Au lieu d'une session Web, on crée une intention de paiement silencieuse
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(unitAmount),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		ApplicationFeeAmount: stripe.Int64(commission),
		TransferData: &stripe.PaymentIntentTransferDataParams{
			Destination: stripe.String(stripeAccountIDSeller),
		},
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true), // Nécessaire pour le SDK Android
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		http.Error(w, "Erreur Stripe", http.StatusInternalServerError)
		return
	}

	// On renvoie le secret au téléphone Android !
	json.NewEncoder(w).Encode(map[string]string{
		"client_secret": pi.ClientSecret,
	})
}

func CreateEventCheckoutSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req struct {
		IdUser  int `json:"id_user"`
		IdEvent int `json:"id_event"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}

	// 1. Récupération des informations de l'événement et du créateur
	var titre string
	var prix float64
	var stripeAccountId string
	
	err := bdd.Db.QueryRow(`
		SELECT e.titre, e.prix, u.stripe_account_id 
		FROM evenement e 
		JOIN utilisateur u ON e.id_salarie = u.id 
		WHERE e.id = ?`, req.IdEvent).Scan(&titre, &prix, &stripeAccountId)
	
	if err != nil {
		http.Error(w, "Événement introuvable", http.StatusNotFound)
		return
	}

	stripe.Key = StripeSecretKey
	
	// 2. Calculs (Stripe en centimes, Base de données en euros)
	unitAmount := int64(prix * 100)
	commissionCentimes := int64(float64(unitAmount) * 0.05) // 5% pour Stripe
	commissionEuros := prix * 0.05                          // 5% pour la BDD

	// 3. Création de la session Stripe
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("eur"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Inscription : " + titre),
					},
					UnitAmount: stripe.Int64(unitAmount),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		PaymentIntentData: &stripe.CheckoutSessionPaymentIntentDataParams{
			ApplicationFeeAmount: stripe.Int64(commissionCentimes),
			TransferData: &stripe.CheckoutSessionPaymentIntentDataTransferDataParams{
				Destination: stripe.String(stripeAccountId),
			},
		},
		SuccessURL: stripe.String("http://localhost:8081/evenement.html?paiement=success"),
		CancelURL:  stripe.String("http://localhost:8081/evenement.html?paiement=cancel"),
	}
	
	params.AddMetadata("id_event", strconv.Itoa(req.IdEvent))
	params.AddMetadata("id_user", strconv.Itoa(req.IdUser))

	s, err := session.New(params)
	if err != nil {
		http.Error(w, "Erreur Stripe: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. ÉTAPE BDD 1 : Création de la commande dans la table `order`
	// Utilisation des backticks pour `order` car c'est un mot réservé en SQL
	queryOrder := "INSERT INTO `order` (id_acheteur, id_annonce, montant_total, commission, date_commande) VALUES (?, ?, ?, ?, NOW())"
	
	result, errOrder := bdd.Db.Exec(queryOrder, req.IdUser, req.IdEvent, prix, commissionEuros)

	if errOrder != nil {
		fmt.Printf("ERREUR INSERTION ORDER : %v\n", errOrder)
	} else {
		// On récupère l'ID généré pour cette nouvelle commande
		idCommandeCreee, _ := result.LastInsertId()
		fmt.Printf("Commande %d créée avec %.2f€ de commission !\n", idCommandeCreee, commissionEuros)

		// 5. ÉTAPE BDD 2 : Liaison avec Stripe dans la table `paiement` (sans le "e")
		queryPaiement := "INSERT INTO paiement (id_commande, stripe_id, statut) VALUES (?, ?, ?)"
		
		_, errPaiement := bdd.Db.Exec(queryPaiement, idCommandeCreee, s.ID, "pending")
		if errPaiement != nil {
			fmt.Printf("ERREUR INSERTION PAIMENT : %v\n", errPaiement)
		} else {
			fmt.Println("Paiement mis en attente avec succès dans la BDD.")
		}
	}

	// 6. Réponse envoyée au front-end
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"checkout_url": s.URL})
}