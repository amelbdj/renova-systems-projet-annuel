package models

type PlanningItem struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Titre string `json:"titre"`
	Date  string `json:"date"`
	Lieu  string `json:"lieu"`
	Meta  string `json:"meta"`
	Theme string `json:"theme"`
}
