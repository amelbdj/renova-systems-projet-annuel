package route

import (
	"net/http"
	"upcycleconnect/admin"
)

// RoutesFinance : tableau de bord financier de l'admin
func RoutesFinance() {
	// --- OPTIONS (pré-vol CORS) ---
	http.HandleFunc("OPTIONS /admin/finance/overview", admin.FinanceOverviewHandler)
	http.HandleFunc("OPTIONS /admin/finance/transactions", admin.AdminTransactionsHandler)

	// --- Vraies routes ---
	http.HandleFunc("GET /admin/finance/overview", admin.FinanceOverviewHandler)
	http.HandleFunc("GET /admin/finance/transactions", admin.AdminTransactionsHandler)
}
