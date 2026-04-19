package admin

import (
	"encoding/json"
	"io"
	"net/http"
	"upcycleconnect/bdd"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/webhook"
)

const WebhookSecret = "whsec_ca8df90af4b7eb3045da3e1ab058edeb4ea8597a2d49b88adec258b0037f4fbc"

func StripeWebhookHandler(w http.ResponseWriter, r *http.Request) {
	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"), WebhookSecret)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if event.Type == "account.updated" {
		var account stripe.Account
		err := json.Unmarshal(event.Data.Raw, &account)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if account.PayoutsEnabled {
			_, err := bdd.Db.Exec("UPDATE utilisateur SET stripe_verif_completed = 1 WHERE stripe_account_id = ?", account.ID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}

	w.WriteHeader(http.StatusOK)
}
