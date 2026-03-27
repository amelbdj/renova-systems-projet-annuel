package admin

import (
	"encoding/json"
	"net/http"
	"strings"
	"upcycleconnect/bdd"
	"upcycleconnect/models"
)



var CacheTraductions = make(map[string]map[string]string)

// RefreshCache charge la BDD en RAM (à appeler dans le main.go)
func RefreshCache() {
	// On récupère toutes les langues actives
	rows, err := bdd.Db.Query("SELECT lang_code, msg_key, msg_value FROM translations")
	if err != nil {
		return
	}
	defer rows.Close()

	// On vide le cache actuel pour le mettre à jour
	newCache := make(map[string]map[string]string)

	for rows.Next() {
		var lang, key, value string
		if err := rows.Scan(&lang, &key, &value); err == nil {
			if newCache[lang] == nil {
				newCache[lang] = make(map[string]string)
			}
			newCache[lang][key] = value
		}
	}
	CacheTraductions = newCache
}

func GetTranslationsHandler(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		lang = "fr"
	}

	translations := CacheTraductions[lang]
	
	if translations == nil {
		var err error
		translations, err = bdd.GetTranslationsByLang(lang)
		if err != nil {
			http.Error(w, "Erreur BDD", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MapToNestedJSON(translations))
}

// MapToNestedJSON : "nav.home" -> {"nav": {"home": "..."}}
func MapToNestedJSON(flatmap map[string]string) map[string]interface{} {
	nested := make(map[string]interface{})
	for key, value := range flatmap {
		parts := strings.Split(key, ".")
		current := nested
		for i, part := range parts {
			if i == len(parts)-1 {
				current[part] = value
			} else {
				if _, ok := current[part]; !ok {
					current[part] = make(map[string]interface{})
				}
				current = current[part].(map[string]interface{})
			}
		}
	}
	return nested
}

// Renvoie la liste des langues pour le menu de Faty
func GetLanguagesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := bdd.Db.Query("SELECT code, name FROM languages")
	if err != nil {
		http.Error(w, "Erreur BDD", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var languages []models.Language
	for rows.Next() {
		var l models.Language
		if err := rows.Scan(&l.Code, &l.Name); err == nil {
			languages = append(languages, l)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(languages)
}