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
		SuccessURL: stripe.String("http://127.0.0.1:5500/Frontend/oneAnnonce.html?id=" + strconv.Itoa(annonceID) + "&payment=success"),
		CancelURL:  stripe.String("http://127.0.0.1:5500/Frontend/oneAnnonce.html?id=" + strconv.Itoa(annonceID)),
	}

	s, err := session.New(params)
	if err != nil {
		http.Error(w, "Erreur Stripe: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"url": s.URL})
}
