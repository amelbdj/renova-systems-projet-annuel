package bdd

import "strings"

func GetAllDocuments() ([]map[string]interface{}, error) {
	query := `
		SELECT d.id_document, d.type_doc, d.url_pdf,
		       COALESCE(d.date_creation, NOW()),
		       COALESCE(u.prenom, ''), COALESCE(u.nom, '')
		FROM document d
		LEFT JOIN utilisateur u ON d.id_user = u.id
		ORDER BY d.id_document DESC`

	rows, err := Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []map[string]interface{}
	for rows.Next() {
		var id int
		var typeDoc, urlPdf, date, prenom, nom string
		rows.Scan(&id, &typeDoc, &urlPdf, &date, &prenom, &nom)
		documents = append(documents, map[string]interface{}{
			"id":          id,
			"type":        typeDoc,
			"url_pdf":     urlPdf,
			"date":        date,
			"utilisateur": strings.TrimSpace(prenom + " " + nom),
		})
	}
	return documents, nil
}
