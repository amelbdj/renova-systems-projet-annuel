package main

import (
	"fmt"
	"net/http"
	"upcycleconnect/admin"
	"upcycleconnect/auth"
	"upcycleconnect/bdd"
)

func Health(w http.ResponseWriter, r *http.Request) {
	err := bdd.Db.Ping()

	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "ping à la bdd")
}

func main() {
	bdd.Db = bdd.NewDB()
	http.HandleFunc("GET /{$}", Health)
	http.HandleFunc("POST /admin/login", admin.Login)

	// Eviter les problèmes de CORS pour les requêtes préliminaires (OPTIONS)
	// Note : definir un dossier route sera pertinant pour avoir des fichiers plus organisés et éviter d'avoir tout dans le main.go

	http.HandleFunc("OPTIONS /admin/users/delete/{id}", admin.DeletedUser)
	http.HandleFunc("OPTIONS /admin/users/add", admin.CreateUser)
	http.HandleFunc("OPTIONS /admin/users/modify/{id}", admin.UpdateUser)
	http.HandleFunc("OPTIONS /admin/users/validate/{id}", admin.ValidateUser)
	http.HandleFunc("OPTIONS /admin/users/refuse/{id}", admin.RefuseUser)

	http.HandleFunc("OPTIONS /admin/categories/add", admin.CreateCategorie)
	http.HandleFunc("OPTIONS /admin/categories/delete/{id}", admin.DeleteCategorie)

	http.HandleFunc("OPTIONS /admin/annonces/validate/{id}", admin.ValidateAnnonce)
	http.HandleFunc("OPTIONS /admin/annonces/refuse/{id}", admin.RefuseAnnonce)
	http.HandleFunc("OPTIONS /admin/annonces/add", admin.CreateAnnonce)
	http.HandleFunc("OPTIONS /admin/annonces/delete/{id}", admin.DeleteAnnonce)
	http.HandleFunc("OPTIONS /admin/annonces/modify/{id}", admin.UpdateAnnonce)

	http.HandleFunc("OPTIONS /admin/evenements/validate/{id}", admin.ValidateEvenement)
	http.HandleFunc("OPTIONS /admin/evenements/update/{id}", admin.UpdateEvenement)
	http.HandleFunc("OPTIONS /admin/evenements/delete/{id}", admin.DeleteEvenement)
	http.HandleFunc("OPTIONS /admin/evenements/refuse/{id}", admin.RefuseEvenement)
	http.HandleFunc("OPTIONS /admin/evenements/add", admin.CreateEvenement)
	http.HandleFunc("OPTIONS /admin/evenements", admin.GetAllEvenements)

	http.HandleFunc("OPTIONS /admin/box/confirm-deposit", admin.ConfirmDeposit)
	http.HandleFunc("OPTIONS /admin/box/collect-object", admin.CollectObject)
	http.HandleFunc("OPTIONS /admin/order/create", admin.CreateOrder)
	http.HandleFunc("OPTIONS /admin/box/create", admin.CreateBox)

	http.HandleFunc("OPTIONS /admin/articles/validate/{id}", admin.ValidateArticle)
	http.HandleFunc("OPTIONS /admin/articles/refuse/{id}", admin.RefuseArticle)
	http.HandleFunc("OPTIONS /admin/articles/delete/{id}", admin.DeleteArticle)
	http.HandleFunc("OPTIONS /admin/articles/add/{action}", admin.CreateArticle)
	http.HandleFunc("OPTIONS /admin/articles/modify/{id}/{action}", admin.ModifyArticle)
	http.HandleFunc("OPTIONS /admin/articles/salarie/{id}", admin.GetArticlesBySalarie)

	// --- USERS ---
	http.HandleFunc("GET /admin/users", auth.VerifyTokenMiddleware(admin.GetAllUsers))
	http.HandleFunc("POST /admin/users/add", auth.VerifyTokenMiddleware(admin.CreateUser))
	http.HandleFunc("DELETE /admin/users/delete/{id}", auth.VerifyTokenMiddleware(admin.DeletedUser))
	http.HandleFunc("PUT /admin/users/modify/{id}", auth.VerifyTokenMiddleware(admin.UpdateUser))
	http.HandleFunc("GET /admin/users/role/{role}", auth.VerifyTokenMiddleware(admin.GetUserByRole))
	http.HandleFunc("GET /admin/users/search", auth.VerifyTokenMiddleware(admin.GetUserByName))
	http.HandleFunc("PUT /admin/users/validate/{id}", auth.VerifyTokenMiddleware(admin.ValidateUser))
	http.HandleFunc("PUT /admin/users/refuse/{id}", auth.VerifyTokenMiddleware(admin.RefuseUser))
	http.HandleFunc("GET /user/profile", admin.GetUserById)

	// Auth, Tutorial & Upload
	http.HandleFunc("/api/upload-document", auth.VerifyTokenMiddleware(admin.UploadDocumentHandler))
	http.Handle("/view-uploads/", http.StripPrefix("/view-uploads/", http.FileServer(http.Dir("./uploads"))))
	http.HandleFunc("/auth/check-email", admin.VerifierEmail)
	http.HandleFunc("/auth/inscription", admin.Inscription)
	http.HandleFunc("/admin/login", admin.Login)
	http.HandleFunc("/update-tutorial", admin.UpdateTutorialStatus)

	// --- CATEGORIES ---
	http.HandleFunc("POST /admin/categories/add", auth.VerifyTokenMiddleware(admin.CreateCategorie))
	http.HandleFunc("DELETE /admin/categories/delete/{id}", auth.VerifyTokenMiddleware(admin.DeleteCategorie))
	http.HandleFunc("GET /admin/categories", auth.VerifyTokenMiddleware(admin.GetAllCategories))
	// --- ANNONCES ---
	http.HandleFunc("GET /admin/annonces", admin.GetAllAnnonces)
	http.HandleFunc("PUT /admin/annonces/validate/{id}", auth.VerifyTokenMiddleware(admin.ValidateAnnonce))
	http.HandleFunc("PUT /admin/annonces/refuse/{id}", auth.VerifyTokenMiddleware(admin.RefuseAnnonce))
	http.HandleFunc("POST /admin/annonces/add", auth.VerifyTokenMiddleware(admin.CreateAnnonce))
	http.HandleFunc("DELETE /admin/annonces/delete/{id}", auth.VerifyTokenMiddleware(admin.DeleteAnnonce))
	http.HandleFunc("PUT /admin/annonces/modify/{id}", auth.VerifyTokenMiddleware(admin.UpdateAnnonce))
	http.HandleFunc("GET /admin/annonces/search", auth.VerifyTokenMiddleware(admin.GetAnnonceByTitle))
	http.HandleFunc("GET /mes-annonces", admin.GetMyAnnonces)

	// --- EVENEMENTS ---
	http.HandleFunc("GET /admin/evenements", admin.GetAllEvenements)
	http.HandleFunc("POST /admin/evenements/add", auth.VerifyTokenMiddleware(admin.CreateEvenement))
	http.HandleFunc("PUT /admin/evenements/{id}", auth.VerifyTokenMiddleware(admin.UpdateEvenement))
	http.HandleFunc("DELETE /admin/evenements/{id}", auth.VerifyTokenMiddleware(admin.DeleteEvenement))
	http.HandleFunc("PUT /admin/evenements/validate/{id}", auth.VerifyTokenMiddleware(admin.ValidateEvenement))
	http.HandleFunc("PUT /admin/evenements/refuse/{id}", auth.VerifyTokenMiddleware(admin.RefuseEvenement))

	// --- LOGISTIQUE ---
	http.HandleFunc("POST /admin/orders/create", auth.VerifyTokenMiddleware(admin.CreateOrder))
	http.HandleFunc("POST /admin/box/confirm-deposit", auth.VerifyTokenMiddleware(admin.ConfirmDeposit))
	http.HandleFunc("POST /admin/box/collect-object", auth.VerifyTokenMiddleware(admin.CollectObject))
	http.HandleFunc("GET /admin/boxs", auth.VerifyTokenMiddleware(admin.GetAllBoxs))       // AJOUTÉ
	http.HandleFunc("POST /admin/box/create", auth.VerifyTokenMiddleware(admin.CreateBox)) // AJOUTÉ

	// --- ARTICLES / NEWS ---
	http.HandleFunc("GET /admin/articles", admin.GetAllArticles)
	http.HandleFunc("GET /admin/articles/salarie/{id}", auth.VerifyTokenMiddleware(admin.GetArticlesBySalarie))
	http.HandleFunc("GET /admin/articles/{id}", admin.GetArticleById)                                       // AJOUTÉ
	http.HandleFunc("PUT /admin/articles/validate/{id}", auth.VerifyTokenMiddleware(admin.ValidateArticle)) // AJOUTÉ
	http.HandleFunc("PUT /admin/articles/refuse/{id}", auth.VerifyTokenMiddleware(admin.RefuseArticle))
	http.HandleFunc("DELETE /admin/articles/delete/{id}", auth.VerifyTokenMiddleware(admin.DeleteArticle))       // AJOUTÉ
	http.HandleFunc("POST /admin/articles/add/{action}", auth.VerifyTokenMiddleware(admin.CreateArticle))        // AJOUTÉ
	http.HandleFunc("PUT /admin/articles/modify/{id}/{action}", auth.VerifyTokenMiddleware(admin.ModifyArticle)) // AJOUTÉ

	//Stripe payment
	http.HandleFunc("POST /admin/connect-stripe", auth.VerifyTokenMiddleware(admin.ConnectToStripe))
	http.HandleFunc("POST /api/stripe/webhook", admin.StripeWebhookHandler)

	fmt.Println("test de : http://localhost:8081")

	err := http.ListenAndServe(":8081", corsMiddleware(http.DefaultServeMux))
	if err != nil {
		panic(err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
