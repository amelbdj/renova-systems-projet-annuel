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
	http.HandleFunc("OPTIONS /admin/users", admin.GetAllUsers)
	http.HandleFunc("OPTIONS /admin/users/role/{role}", admin.GetUserByRole)
	http.HandleFunc("OPTIONS /admin/users/search", admin.GetUserByName)
	http.HandleFunc("OPTIONS /admin/users/{id}", admin.GetUserById)
	http.HandleFunc("OPTIONS /admin/users/ban/{id}", admin.BanUserHandler)
	http.HandleFunc("OPTIONS /user/profile", admin.GetUserById)

	http.HandleFunc("OPTIONS /admin/categories/add", admin.CreateCategorie)
	http.HandleFunc("OPTIONS /admin/categories/delete/{id}", admin.DeleteCategorie)
	http.HandleFunc("OPTIONS /admin/categories", admin.GetAllCategories)

	http.HandleFunc("OPTIONS /admin/annonces/validate/{id}", admin.ValidateAnnonce)
	http.HandleFunc("OPTIONS /admin/annonces", admin.GetAllAnnonces)
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
	http.HandleFunc("OPTIONS /admin/evenements/inscription", admin.InscrireClient)

	http.HandleFunc("OPTIONS /admin/articles/validate/{id}", admin.ValidateArticle)
	http.HandleFunc("OPTIONS /admin/articles/refuse/{id}", admin.RefuseArticle)
	http.HandleFunc("OPTIONS /admin/articles/delete/{id}", admin.DeleteArticle)
	http.HandleFunc("OPTIONS /admin/articles/add/{action}", admin.CreateArticle)
	http.HandleFunc("OPTIONS /admin/articles/modify/{id}/{action}", admin.ModifyArticle)
	http.HandleFunc("OPTIONS /admin/articles/salarie/{id}", admin.GetArticlesBySalarie)
	http.HandleFunc("OPTIONS /admin/articles/{id}", admin.GetArticleById)
	http.HandleFunc("OPTIONS /admin/articles", admin.GetAllArticles)
	http.HandleFunc("OPTIONS /api/user/boxes", admin.GetMyBoxes)
	http.HandleFunc("OPTIONS /api/admin/conteneur/{id}/boxes", admin.GetBoxesForConteneurHandler)
	http.HandleFunc("OPTIONS /api/admin/box/add", admin.AddSingleBoxHandler)
	http.HandleFunc("OPTIONS /api/admin/box/update", admin.UpdateBoxStatusHandler)

	// --- TRANSLATIONS OPTIONS ---
	http.HandleFunc("OPTIONS /api/translations", admin.GetTranslations)
	http.HandleFunc("OPTIONS /api/languages", admin.GetLanguages)
	http.HandleFunc("OPTIONS /admin/translations/add", admin.AddLanguage)
	http.HandleFunc("OPTIONS /admin/translations/keys", admin.GetTranslationKeysHandler)

	http.HandleFunc("OPTIONS /admin/forum/messages", admin.GetForumMessages)
	http.HandleFunc("OPTIONS /admin/forum/messages/moderate/{id}", admin.ModerateForumMessage)
	http.HandleFunc("OPTIONS /admin/forum/stats", admin.GetForumStats)

	http.HandleFunc("OPTIONS /api/box/reserve", admin.ReserveBox)
	http.HandleFunc("OPTIONS /api/box/deposit", admin.ConfirmDeposit)
	http.HandleFunc("OPTIONS /api/box/collect", admin.CollectObject)
	http.HandleFunc("OPTIONS /api/admin/conteneurs", admin.GetConteneursAdmin)
	http.HandleFunc("OPTIONS /api/admin/conteneur/create", admin.CreateConteneur)

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
	http.HandleFunc("GET /admin/users/{id}", auth.VerifyTokenMiddleware(admin.GetUserById))
	http.HandleFunc("PUT /admin/users/ban/{id}", auth.VerifyTokenMiddleware(admin.BanUserHandler))
	http.HandleFunc("GET /api/user/payment-history", admin.PaymentHistoryHandler)
	http.HandleFunc("OPTIONS /api/user/payment-history", admin.PaymentHistoryHandler)

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
	http.HandleFunc("GET /api/annonces/all", admin.GetValidatedAnnonces)
	http.HandleFunc("GET /api/annonces", admin.GetOneAnnonce)
	// http.HandleFunc("OPTIONS /api/annonces/vendre", admin.AnnVendu)
	// http.HandleFunc("PUT /api/annonces/vendre", admin.AnnVendu)
	http.HandleFunc("POST /api/annonces/vendre", admin.ConfirmPaymentAndOrder)
	http.HandleFunc("OPTIONS /api/annonces/vendre", admin.ConfirmPaymentAndOrder)
	http.HandleFunc("/api/user/stats", admin.GetEcoStatsHandler)

	// --- EVENEMENTS ---
	http.HandleFunc("GET /admin/evenements", admin.GetAllEvenements)
	http.HandleFunc("POST /admin/evenements/add", auth.VerifyTokenMiddleware(admin.CreateEvenement))
	http.HandleFunc("PUT /admin/evenements/{id}", auth.VerifyTokenMiddleware(admin.UpdateEvenement))
	http.HandleFunc("DELETE /admin/evenements/delete/{id}", auth.VerifyTokenMiddleware(admin.DeleteEvenement))
	http.HandleFunc("PUT /admin/evenements/validate/{id}", auth.VerifyTokenMiddleware(admin.ValidateEvenement))
	http.HandleFunc("PUT /admin/evenements/refuse/{id}", auth.VerifyTokenMiddleware(admin.RefuseEvenement))
	http.HandleFunc("POST /admin/evenements/inscription", admin.InscrireClient)
	// --- LOGISTIQUE ---
	http.HandleFunc("POST /api/box/reserve", admin.ReserveBox)
	http.HandleFunc("POST /api/box/deposit", admin.ConfirmDeposit)
	http.HandleFunc("POST /api/box/collect", admin.CollectObject)
	http.HandleFunc("GET /api/admin/conteneurs", admin.GetConteneursAdmin)
	http.HandleFunc("POST /api/admin/conteneur/create", admin.CreateConteneur)
	http.HandleFunc("POST /api/admin/box/add", admin.AddSingleBoxHandler)
	http.HandleFunc("PUT /api/admin/box/update", admin.UpdateBoxStatusHandler)

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
	http.HandleFunc("OPTIONS /admin/connect-stripe", admin.ConnectToStripe)
	http.HandleFunc("OPTIONS /api/payment-annonce", admin.PaymentAnnonce)
	http.HandleFunc("POST /api/payment-annonce", admin.PaymentAnnonce)

	// boxes
	http.HandleFunc("GET /api/user/boxes", admin.GetMyBoxes)
	// Le {id} entre accolades indique à Go que cette partie de l'URL est une variable dynamique !
	http.HandleFunc("GET /api/admin/conteneur/{id}/boxes", admin.GetBoxesForConteneurHandler)

	// --- TRADUCTIONS ---
	http.HandleFunc("GET /api/translations", admin.GetTranslations)
	http.HandleFunc("GET /api/languages", admin.GetLanguages)
	http.HandleFunc("POST /admin/translations/add", admin.AddLanguage)
	http.HandleFunc("GET /admin/translations/keys", admin.GetTranslationKeysHandler)

	http.HandleFunc("GET /admin/forum/messages", auth.VerifyTokenMiddleware(admin.GetForumMessages))
	http.HandleFunc("PUT /admin/forum/messages/moderate/{id}", auth.VerifyTokenMiddleware(admin.ModerateForumMessage))
	http.HandleFunc("GET /admin/forum/stats", auth.VerifyTokenMiddleware(admin.GetForumStats))

	http.HandleFunc("OPTIONS /user/planning", admin.GetUserPlanningHandler)
	http.HandleFunc("GET /user/planning", admin.GetUserPlanningHandler)

	http.HandleFunc("OPTIONS /user/forums", admin.GetForumsHandler)
	http.HandleFunc("GET /user/forums", admin.GetForumsHandler)
	http.HandleFunc("POST /user/forums", admin.GetForumsHandler)

	http.HandleFunc("OPTIONS /user/forums/messages", admin.ForumClientMessagesHandler)
	http.HandleFunc("GET /user/forums/messages", admin.ForumClientMessagesHandler)
	http.HandleFunc("POST /user/forums/messages", admin.ForumClientMessagesHandler)

	// --- MESSAGERIE (WEBSOCKET & HISTORIQUE) ---

	// 1. L'historique des messages (Besoin du Token pour la sécurité)
	// --- MESSAGERIE (WEBSOCKET & HISTORIQUE) ---

	// 1. L'historique des messages (Besoin du Token pour la sécurité)
	http.HandleFunc("OPTIONS /api/chat/history", admin.GetChatHistoryHandler)
	http.HandleFunc("GET /api/chat/history", auth.VerifyTokenMiddleware(admin.GetChatHistoryHandler))

	// 2. Le WebSocket (ATTENTION : Un WebSocket s'initie TOUJOURS avec un GET !)
	http.HandleFunc("GET /ws/chat", admin.ChatHandler)

http.HandleFunc("POST /admin/evenements/desinscription", admin.DesinscriptionHandler)
http.HandleFunc("OPTIONS /admin/evenements/desinscription", admin.DesinscriptionHandler)
	
// 1. On autorise la vraie requête GET
// Routes pour les revenus (Overview)
http.HandleFunc("GET /admin/finance/overview", admin.FinanceOverviewHandler)
http.HandleFunc("OPTIONS /admin/finance/overview", admin.FinanceOverviewHandler)

// Routes pour le tableau des transactions
http.HandleFunc("GET /admin/finance/transactions", admin.AdminTransactionsHandler)
http.HandleFunc("OPTIONS /admin/finance/transactions", admin.AdminTransactionsHandler)
	// 3. La liste des conversations pour le Dashboard
	// On ajoute OPTIONS pour le CORS et GET avec le Middleware de sécurité
	http.HandleFunc("OPTIONS /api/chat/conversations", admin.GetConversationsHandler)
	http.HandleFunc("GET /api/chat/conversations", auth.VerifyTokenMiddleware(admin.GetConversationsHandler))

	http.HandleFunc("POST /admin/evenements/desinscription", admin.DesinscriptionHandler)
	http.HandleFunc("OPTIONS /admin/evenements/desinscription", admin.DesinscriptionHandler)

	fmt.Println("test de : http://localhost:8081")

	http.ListenAndServe(":8081", nil)
}
