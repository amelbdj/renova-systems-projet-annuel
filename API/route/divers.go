package route

import (
	"fmt"
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/bdd"
)

func Health(w http.ResponseWriter, r *http.Request) {
	err := bdd.Db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "ping à la bdd")
}

func RoutesDivers() {
	http.HandleFunc("GET /{$}", Health)

	http.HandleFunc("/api/hardware/simulate-withdrawal", admin.SimulateWithdrawalHandler)
	http.HandleFunc("/api/hardware/simulate-deposit", admin.SimulateDepositHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
