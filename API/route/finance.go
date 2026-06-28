package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

func RoutesFinance() {

	http.HandleFunc("OPTIONS /admin/finance/overview", admin.FinanceOverviewHandler)
	http.HandleFunc("OPTIONS /admin/finance/transactions", admin.AdminTransactionsHandler)
	http.HandleFunc("OPTIONS /api/pro/invoices", admin.GetProInvoicesHandler)

	http.HandleFunc("GET /admin/finance/overview", admin.FinanceOverviewHandler)
	http.HandleFunc("GET /admin/finance/transactions", admin.AdminTransactionsHandler)

	http.HandleFunc("GET /api/pro/invoices", auth.VerifyTokenMiddleware(admin.GetProInvoicesHandler))
}
