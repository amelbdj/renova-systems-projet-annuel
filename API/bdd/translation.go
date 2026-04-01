package bdd

import (
	"log"
	"upcycleconnect/models"
) 

type TranslationPayload struct {
	LangCode     string `json:"lang_code"`
	Translations []struct {
		Key   string `json:"msg_key"`
		Value string `json:"msg_value"`
	} `json:"translations"`
}

func GetTranslationsByLang(lang string) (map[string]string, error) {

	locales := make(map[string]string) // map vide pour stock les trad

	rows, err := Db.Query("SELECT msg_key, msg_value FROM translations WHERE lang_code = ?", lang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return nil, err
		}
		// On remplit la map : "login.title" -> "Connexion"
		locales[key] = value
	}

	return locales, nil
}

func GetLanguages() ([]models.Language, error) {
	rows, err := Db.Query("SELECT code, name FROM languages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var languages []models.Language
	for rows.Next() {
		var lang models.Language
		if err := rows.Scan(&lang.Code, &lang.Name); err != nil {
			return nil, err
		}
		languages = append(languages, lang)
	}
	return languages, nil
}



func AddNewLanguage(payload TranslationPayload) error {

	_, err := Db.Exec("INSERT IGNORE INTO languages (code) VALUES (?)", payload.LangCode)
	if err != nil {
		log.Println("Erreur lors de la création de la langue dans la table mère :", err)
		return err
	}
	// On prépare la requête d'insertion
	stmt, err := Db.Prepare("INSERT INTO translations (lang_code, msg_key, msg_value) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, t := range payload.Translations {
		_, err := stmt.Exec(payload.LangCode, t.Key, t.Value)
		if err != nil {
			log.Println("Erreur lors de l'insertion de la clé:", t.Key, err)
		}
	}
	return nil
}

// Fonction pour récupérer toutes les clés uniques de traduction
func GetAllTranslationKeys() ([]string, error) {
	// On demande à MySQL de nous lister toutes les clés sans doublons (DISTINCT)
	rows, err := Db.Query("SELECT DISTINCT msg_key FROM translations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var keys []string
	for rows.Next() {
		var key string
		if err := rows.Scan(&key); err != nil {
			return nil, err
		}
		keys = append(keys, key)
	}
	return keys, nil
}