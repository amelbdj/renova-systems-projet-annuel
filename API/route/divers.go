package route

import (
	"fmt"
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/bdd"
)

// Health : petit ping pour vérifier que la BDD répond
func Health(w http.ResponseWriter, r *http.Request) {
	err := bdd.Db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "ping à la bdd")
}

// RoutesDivers : health check, fichiers statiques et simulation matériel (hardware)
func RoutesDivers() {
	http.HandleFunc("GET /{$}", Health)

	// Simulation du matériel (bornes physiques)
	http.HandleFunc("/api/hardware/simulate-withdrawal", admin.SimulateWithdrawalHandler)
	http.HandleFunc("/api/hardware/simulate-deposit", admin.SimulateDepositHandler)

	// Si on demande une URL qui commence par /static/, on sert le fichier du dossier static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
