package route

import (
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

func RoutesAuth() {
	http.HandleFunc("POST /admin/login", admin.Login)
	http.HandleFunc("/admin/login", admin.Login)
	http.HandleFunc("/auth/check-email", admin.VerifierEmail)
	http.HandleFunc("/auth/inscription", admin.Inscription)
	http.HandleFunc("/update-tutorial", admin.UpdateTutorialStatus)

	http.HandleFunc("/api/upload-document", auth.VerifyTokenMiddleware(admin.UploadDocumentHandler))
	http.Handle("/view-uploads/", http.StripPrefix("/view-uploads/", http.FileServer(http.Dir("./uploads"))))
}
