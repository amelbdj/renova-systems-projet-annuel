package admin

import (
	"encoding/json"
	"net/http"
	"strconv"
	"upcycleconnect/bdd"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/account"
	"github.com/stripe/stripe-go/v81/accountlink"
)

const StripeSecretKey = "sk_test_51TNFHBHbaxF1KOTtH89RRHNJQSQXSPVtOHMJDHicr1LW4XYeY4ZC6nYWwzVbDvFUUI58YA7KlJs9BiUyP5zD4XU300gaAUPVpI"

func ConnectToStripe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

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
		RefreshURL: stripe.String("http://localhost:5500/Frontend/profil.html"),
		ReturnURL:  stripe.String("http://localhost:5500/Frontend/profil.html?stripe=success"),
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
