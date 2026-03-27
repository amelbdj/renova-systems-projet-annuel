package bdd

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
