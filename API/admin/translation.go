package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"upcycleconnect/bdd"
)



func GetTranslations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		lang = "fr"
	}

	// On récupère les traductions directement depuis la BDD
	translations, err := bdd.GetTranslationsByLang(lang)
	if err != nil {
		http.Error(w, "Erreur BDD", http.StatusInternalServerError)
		return
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var payload bdd.TranslationPayload

	// On lit le JSON envoyé par le Front
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil || payload.LangCode == "" || payload.LangName == "" {
		http.Error(w, `{"erreur": "Données invalides (code et nom de langue obligatoires)"}`, http.StatusBadRequest)
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
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}


	keys, err := bdd.GetAllTranslationKeys()
	if err != nil {
		http.Error(w, `{"erreur": "Erreur BDD"}`, http.StatusInternalServerError)
		return
	}

	// On renvoie le tableau de clés en JSON
	json.NewEncoder(w).Encode(keys)
}