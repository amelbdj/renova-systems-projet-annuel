package bdd

import (
	"database/sql"
	"fmt"
	"upcycleconnect/models"
)


func GetUserPlanning(userID int) ([]models.PlanningItem, error) {
    planning := []models.PlanningItem{}

 // events
    
    rowsEv, err := Db.Query("SELECT evenement.id, evenement.titre, DATE_FORMAT(evenement.date_debut, '%d-%m-%Y'), evenement.lieu FROM evenement JOIN inscription ON evenement.id = inscription.id_event WHERE inscription.id_user = ?", userID)
    if err == nil {
        defer rowsEv.Close()
        for rowsEv.Next() {
            var item models.PlanningItem
            rowsEv.Scan(&item.ID, &item.Titre, &item.Date, &item.Lieu)
            
            item.Type = "evenement"
            item.Theme = "bl" 
            item.Meta = "Inscrit" 
            
            planning = append(planning, item)
        }
    } else {
        fmt.Println("🚨 ERREUR SQL Événements :", err)
    }

// depot
	lignesDepots, errDepot := Db.Query("SELECT d.id_depot, a.titre, d.date_depot, b.localisation, d.code_ouverture FROM depot_box d JOIN box_conteneur b ON d.id_box = b.id JOIN annonce a ON d.id_annonce = a.id WHERE a.id_user = ? AND d.date_depot IS NOT NULL", userID)
    if errDepot == nil { 
        defer lignesDepots.Close() 
        for lignesDepots.Next() {
            var item models.PlanningItem
            var code sql.NullString

            lignesDepots.Scan(&item.ID, &item.Titre, &item.Date, &item.Lieu, &code)

            item.Type = "box_depot" 
            if code.Valid && code.String != "" {
                item.Meta = code.String 
            }
            
            planning = append(planning, item)
        }
    } else {
        fmt.Println("erreur dépôts :", errDepot)
    }

	// retrait

    lignesRetraits, errRetrait := Db.Query(" SELECT d.id_depot, a.titre, DATE_FORMAT(d.date_retrait, '%d-%m-%Y'), b.localisation, d.code_ouverture FROM depot_box d JOIN box_conteneur b ON d.id_box = b.id JOIN annonce a ON d.id_annonce = a.id WHERE a.id_user = ? AND d.date_retrait IS NOT NULL", userID)
    if errRetrait == nil { 
        defer lignesRetraits.Close() 
        for lignesRetraits.Next() {
            var item models.PlanningItem
            var code sql.NullString

            lignesRetraits.Scan(&item.ID, &item.Titre, &item.Date, &item.Lieu, &code)

            item.Type = "box_retrait" 
            if code.Valid && code.String != "" {
                item.Meta = code.String 
            }
            
            planning = append(planning, item)
        }
    } else {
        fmt.Println("erreur retraits :", errRetrait)
    }

    return planning, nil
}