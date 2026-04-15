package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"upcycleconnect/bdd"
)


func CreateOrder(w http.ResponseWriter, r *http.Request) {
    var req struct {
        AnnonceId int `json:"annonce_id"`
        AcheteurId int `json:"acheteur_id"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Données invalides", http.StatusBadRequest)
        return
    }

    annonce, err := bdd.GetAnnonceById(req.AnnonceId)
    if err != nil {
        http.Error(w, "Annonce introuvable", http.StatusNotFound)
        return
    }

    orderId, err := bdd.CreateOrder(annonce, req.AcheteurId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "ID de commande : %d", orderId)
}