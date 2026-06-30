package route

import (
	"net/http"
	"os"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
)

func RoutesAuth() {
	http.HandleFunc("POST /admin/login", admin.Login)
	http.HandleFunc("/admin/login", admin.Login)
	http.HandleFunc("/auth/check-email", admin.VerifierEmail)
	http.HandleFunc("/auth/inscription", admin.Inscription)
	http.HandleFunc("POST /auth/forgot-password", admin.ForgotPasswordHandler)
	http.HandleFunc("OPTIONS /auth/forgot-password", admin.ForgotPasswordHandler)
	http.HandleFunc("POST /auth/reset-password", admin.ResetPasswordTokenHandler)
	http.HandleFunc("OPTIONS /auth/reset-password", admin.ResetPasswordTokenHandler)
	http.HandleFunc("/update-tutorial", admin.UpdateTutorialStatus)

	http.HandleFunc("/api/upload-document", auth.VerifyTokenMiddleware(admin.UploadDocumentHandler))

	// Dossier d'uploads configurable (UPLOAD_DIR), defaut ./uploads (== /app/uploads en Docker)
	uploadDir := os.Getenv("UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = "./uploads"
	}
	// Nouvelle route propre : /uploads/annonces/..., /uploads/articles/..., /uploads/documents/...
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir(uploadDir))))

	// Routes legacy (compat avec les anciennes donnees deja en base)
	http.Handle("/view-uploads/", http.StripPrefix("/view-uploads/", http.FileServer(http.Dir(uploadDir))))
	http.Handle("/view-documents/", http.StripPrefix("/view-documents/", http.FileServer(http.Dir("./documents"))))
}
