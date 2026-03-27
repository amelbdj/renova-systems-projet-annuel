package models

type Translation struct {
	Id       int    `json:"id"`
	LangCode string `json:"lang_code"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}