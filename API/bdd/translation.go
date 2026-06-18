package bdd

import (
	"log"
	"upcycleconnect/models"
)

type TranslationPayload struct {
	LangCode string                 `json:"lang_code"`
	LangName string                 `json:"lang_name"`
	Data     map[string]interface{} `json:"data"`
}

func Aplatir(prefixe string, data map[string]interface{}, resultat map[string]string) {
	for cle, valeur := range data {
		nouvelleCle := cle
		if prefixe != "" {
			nouvelleCle = prefixe + "." + cle
		}

		if sousObjet, ok := valeur.(map[string]interface{}); ok {
			Aplatir(nouvelleCle, sousObjet, resultat)
		} else if texte, ok := valeur.(string); ok {

			resultat[nouvelleCle] = texte
		}
	}
}

func GetTranslationsByLang(lang string) (map[string]string, error) {

	locales := make(map[string]string)

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

		locales[key] = value
	}

	return locales, nil
}

func GetLanguages() ([]models.Language, error) {
	rows, err := Db.Query("SELECT code, name FROM languages WHERE is_active = 1")
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

	_, err := Db.Exec(
		"INSERT INTO languages (code, name, is_active) VALUES (?, ?, 1) ON DUPLICATE KEY UPDATE name = VALUES(name)",
		payload.LangCode, payload.LangName,
	)
	if err != nil {
		log.Println("Erreur lors de la création de la langue :", err)
		return err
	}

	traductions := make(map[string]string)
	Aplatir("", payload.Data, traductions)

	stmt, err := Db.Prepare("INSERT INTO translations (lang_code, msg_key, msg_value) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE msg_value = VALUES(msg_value)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for cle, valeur := range traductions {
		_, err := stmt.Exec(payload.LangCode, cle, valeur)
		if err != nil {
			log.Println("Erreur lors de l'insertion de la clé:", cle, err)
		}
	}
	return nil
}

func GetAllTranslationKeys() ([]string, error) {

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
