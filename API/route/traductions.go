package route

import (
	"net/http"
	"upcycleconnect/admin"
)

func RoutesTraductions() {

	http.HandleFunc("OPTIONS /api/translations", admin.GetTranslations)
	http.HandleFunc("OPTIONS /api/languages", admin.GetLanguages)
	http.HandleFunc("OPTIONS /admin/translations/add", admin.AddLanguage)
	http.HandleFunc("OPTIONS /admin/translations/keys", admin.GetTranslationKeysHandler)

	http.HandleFunc("GET /api/translations", admin.GetTranslations)
	http.HandleFunc("GET /api/languages", admin.GetLanguages)
	http.HandleFunc("POST /admin/translations/add", admin.AddLanguage)
	http.HandleFunc("GET /admin/translations/keys", admin.GetTranslationKeysHandler)
}
