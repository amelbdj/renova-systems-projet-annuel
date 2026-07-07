package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"upcycleconnect/bdd"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/account"
	"github.com/stripe/stripe-go/v81/accountlink"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"github.com/stripe/stripe-go/v81/paymentintent"
	"github.com/stripe/stripe-go/v81/price"
	"github.com/stripe/stripe-go/v81/product"
)

const StripeSecretKey = "sk_test_51TNFHBHbaxF1KOTtH89RRHNJQSQXSPVtOHMJDHicr1LW4XYeY4ZC6nYWwzVbDvFUUI58YA7KlJs9BiUyP5zD4XU300gaAUPVpI"

func getStripeSecretKey() string {
	key := os.Getenv("STRIPE_SECRET_KEY")
	if key != "" {
		return key
	}
	return StripeSecretKey
}

func getFrontendBaseURL() string {
	url := os.Getenv("FRONTEND_BASE_URL")
	if url != "" {
		return strings.TrimRight(url, "/")
	}
	return "http://localhost/renova-systems-projet-annuel/Frontend"
}

func frontURL(path string) string {
	return getFrontendBaseURL() + "/" + strings.TrimLeft(path, "/")
}

func getPlanInfo(plan string) (name string, amountCents int64, ok bool) {
	switch planKey(plan) {
	case "premium":
		return "Abonnement Premium Pro", 2500, true
	case "plus":
		return "Abonnement Plus Pro", 4500, true
	case "pro":
		return "Abonnement Pro", 9900, true
	default:
		return "", 0, false
	}
}

func planKey(plan string) string {
	switch plan {
	case "plus":
		return "plus"
	case "pro":
		return "pro"
	default:
		return "premium"
	}
}

func ConnectToStripe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	stripe.Key = getStripeSecretKey()

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
		RefreshURL: stripe.String(frontURL("profil.html")),
		ReturnURL:  stripe.String(frontURL("profil.html?stripe=success")),
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
	stripe.Key = getStripeSecretKey()
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
		SuccessURL: stripe.String(frontURL("oneAnnonce.html?id=" + strconv.Itoa(annonceID) + "&buyer_id=" + strconv.Itoa(buyerID) + "&payment=success")),
		CancelURL:  stripe.String(frontURL("oneAnnonce.html?id=" + strconv.Itoa(annonceID))),
	}

	s, err := session.New(params)
	if err != nil {
		http.Error(w, "Erreur Stripe: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"url": s.URL})
}

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

	stripe.Key = getStripeSecretKey()
	unitAmount := int64(prix * 100)
	commission := (unitAmount * 5) / 100

	params := &stripe.PaymentIntentParams{
		Amount:               stripe.Int64(unitAmount),
		Currency:             stripe.String(string(stripe.CurrencyEUR)),
		ApplicationFeeAmount: stripe.Int64(commission),
		TransferData: &stripe.PaymentIntentTransferDataParams{
			Destination: stripe.String(stripeAccountIDSeller),
		},
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		http.Error(w, "Erreur Stripe", http.StatusInternalServerError)
		return
	}

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

	stripe.Key = getStripeSecretKey()

	unitAmount := int64(prix * 100)
	commissionCentimes := int64(float64(unitAmount) * 0.05)
	commissionEuros := prix * 0.05

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
		SuccessURL: stripe.String(frontURL("evenement.html?paiement=success")),
		CancelURL:  stripe.String(frontURL("evenement.html?paiement=cancel")),
	}
	params.AddMetadata("id_event", strconv.Itoa(req.IdEvent))
	params.AddMetadata("id_user", strconv.Itoa(req.IdUser))

	s, err := session.New(params)
	if err != nil {
		http.Error(w, "Erreur Stripe: "+err.Error(), http.StatusInternalServerError)
		return
	}

	queryOrder := "INSERT INTO `order` (id_acheteur, id_annonce, montant_total, commission, date_commande) VALUES (?, ?, ?, ?, NOW())"

	result, errOrder := bdd.Db.Exec(queryOrder, req.IdUser, req.IdEvent, prix, commissionEuros)

	if errOrder != nil {
		fmt.Printf("ERREUR INSERTION ORDER : %v\n", errOrder)
	} else {

		idCommandeCreee, _ := result.LastInsertId()
		fmt.Printf("Commande %d créée avec %.2f€ de commission !\n", idCommandeCreee, commissionEuros)

		queryPaiement := "INSERT INTO paiement (id_commande, stripe_id, statut) VALUES (?, ?, ?)"
		_, errPaiement := bdd.Db.Exec(queryPaiement, idCommandeCreee, s.ID, "pending")
		if errPaiement != nil {
			fmt.Printf("ERREUR INSERTION PAIMENT : %v\n", errPaiement)
		} else {
			fmt.Println("Paiement mis en attente avec succès dans la BDD.")
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"checkout_url": s.URL})
}

func CreateProSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, `{"error": "ID utilisateur manquant"}`, http.StatusBadRequest)
		return
	}

	plan := r.URL.Query().Get("plan")
	planName, planAmount, ok := getPlanInfo(plan)
	if !ok {
		http.Error(w, `{"error": "Plan inconnu"}`, http.StatusBadRequest)
		return
	}
	plan = planKey(plan)

	stripe.Key = getStripeSecretKey()

	prodParams := &stripe.ProductParams{
		Name: stripe.String(planName),
	}
	prod, errProd := product.New(prodParams)
	if errProd != nil {
		fmt.Println("Erreur création produit :", errProd)
		http.Error(w, `{"error": "Impossible de créer le produit"}`, http.StatusInternalServerError)
		return
	}

	priceParams := &stripe.PriceParams{
		Product:    stripe.String(prod.ID),
		UnitAmount: stripe.Int64(planAmount),
		Currency:   stripe.String(string(stripe.CurrencyEUR)),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
	}
	newPrice, errPrice := price.New(priceParams)
	if errPrice != nil {
		fmt.Println("Erreur création prix :", errPrice)
		http.Error(w, `{"error": "Impossible de créer le prix"}`, http.StatusInternalServerError)
		return
	}

	fmt.Println("price id works:", newPrice.ID)

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		Mode:               stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(newPrice.ID),
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL:        stripe.String(frontURL("espPro.html?abo=success&session_id={CHECKOUT_SESSION_ID}")),
		CancelURL:         stripe.String(frontURL("espPro.html?abo=cancel")),
		ClientReferenceID: stripe.String(userID),
	}
	params.AddMetadata("plan", plan)

	s, err := session.New(params)
	if err != nil {
		fmt.Println("Stripe Error:", err)
		http.Error(w, `{"error": "Impossible de contacter Stripe"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"url": s.URL})
}
