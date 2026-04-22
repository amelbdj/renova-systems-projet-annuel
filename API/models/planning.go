package models

type PlanningItem struct {
	ID    int    `json:"id"`
	Type  string `json:"type"` // "evenement", "formation", "box"
	Titre string `json:"titre"`
	Date  string `json:"date"`
	Lieu  string `json:"lieu"`
	Meta  string `json:"meta"` // "Payé", "À récupérer", etc.
	Theme string `json:"theme"`
}