package admin

import (
	"encoding/json"
	"net/http"
	"upcycleconnect/bdd"
)

func GetAllDocumentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	docs, err := bdd.GetAllDocuments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if docs == nil {
		docs = []map[string]interface{}{}
	}
	json.NewEncoder(w).Encode(docs)
}
