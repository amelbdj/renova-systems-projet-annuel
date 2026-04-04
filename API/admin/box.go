package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"upcycleconnect/bdd"
)

func ReserveBox(w http.ResponseWriter, r *http.Request) {
    // Structure pour lire ce que le Front envoie
    var req struct {
        AnnonceId     int `json:"annonce_id"`
        ConteneurId   int `json:"conteneur_id"`
        ParticulierId int `json:"particulier_id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Données invalides", http.StatusBadRequest)
        return
    }

    err := bdd.ReserveBox(req.AnnonceId, req.ConteneurId, req.ParticulierId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Box réservé avec succès")
}

func ConfirmDeposit(w http.ResponseWriter, r *http.Request) {
    var req struct {
        PinCode string `json:"pin_code"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Code PIN manquant", http.StatusBadRequest)
        return
    }

    err := bdd.ConfirmDeposit(req.PinCode)
    if err != nil {
        http.Error(w, "Code PIN incorrect ou expiré", http.StatusUnauthorized)
        return
    }

    fmt.Fprint(w, "Dépôt validé, le box est verrouillé")
}

func CollectObject(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Barcode        string `json:"barcode"`
        ProfessionnelId int  `json:"professionnel_id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Données de collecte invalides", http.StatusBadRequest)
        return
    }

    err := bdd.CollectObject(req.Barcode, req.ProfessionnelId)
    if err != nil {
        http.Error(w, "Erreur lors de la collecte : "+err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprint(w, "Objet récupéré, box libéré et score mis à jour !")
}

func GetAllBoxs(w http.ResponseWriter, r *http.Request) {
    Boxs, err := bdd.GetBox()
    if err != nil {
        http.Error(w, "Erreur lors de la récupération des boxs : "+err.Error(), http.StatusInternalServerError)
        return
    }
   response, err := json.Marshal(Boxs)

	if err != nil {
		http.Error(w, "erreur de conversion", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", response)
}