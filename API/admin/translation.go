package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"upcycleconnect/bdd"
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

func GetTranslations(w http.ResponseWriter, r *http.Request) {
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
func GetLanguages(w http.ResponseWriter, r *http.Request) {
	users, err := bdd.GetLanguages()

	if err != nil {
		http.Error(w, "erreur de récupération des langues", http.StatusInternalServerError)

		return
	}

	response, err := json.Marshal(users)

	if err != nil {
		http.Error(w, "erreur de conversion", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", response)
}

func AddLanguage(w http.ResponseWriter, r *http.Request) {
 w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return 
    }

	var payload bdd.TranslationPayload

	// On lit le JSON envoyé par le Front
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil || payload.LangCode == "" {
		http.Error(w, `{"erreur": "Données invalides"}`, http.StatusBadRequest)
		return
	}

	// On envoie à la BDD
	err = bdd.AddNewLanguage(payload)
	if err != nil {
		http.Error(w, `{"erreur": "Erreur SQL"}`, http.StatusInternalServerError)
		return
	}

	// On répond que tout s'est bien passé
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Nouvelle langue ajoutée avec succès !"})
}

func GetTranslationKeysHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	keys, err := bdd.GetAllTranslationKeys()
	if err != nil {
		http.Error(w, `{"erreur": "Erreur BDD"}`, http.StatusInternalServerError)
		return
	}

	// On renvoie le tableau de clés en JSON
	json.NewEncoder(w).Encode(keys)
}