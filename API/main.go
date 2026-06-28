package main

import (
	"fmt"
	"net/http"
	"upcycleconnect/bdd"
	"upcycleconnect/route"
)

func main() {
	bdd.Db = bdd.NewDB()

	route.RoutesAuth()
	route.RoutesUsers()
	route.RoutesCategories()
	route.RoutesAnnonces()
	route.RoutesEvenements()
	route.RoutesArticles()
	route.RoutesLogistique()
	route.RoutesStripe()
	route.RoutesTraductions()
	route.RoutesForum()
	route.RoutesChat()
	route.RoutesPlanning()
	route.RoutesFinance()
	route.RoutesDivers()
	route.RoutesPro()
	route.RoutesNotifications()

<<<<<<< HEAD
=======
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
	http.HandleFunc("POST /api/user/update-password", auth.VerifyTokenMiddleware(admin.UpdatePasswordHandler))
	http.HandleFunc("OPTIONS /api/user/update-password", admin.UpdatePasswordHandler)
	http.HandleFunc("POST /api/pro/upgrade", admin.UpgradeToPremiumHandler)
	http.HandleFunc("OPTIONS /api/pro/upgrade", admin.UpgradeToPremiumHandler)
	http.HandleFunc("POST /api/pro/portal", admin.CustomerPortalHandler)
	http.HandleFunc("OPTIONS /api/pro/portal", admin.CustomerPortalHandler)
	http.HandleFunc("POST /api/pro/cancel", admin.CancelSubscriptionHandler)
	http.HandleFunc("OPTIONS /api/pro/cancel", admin.CancelSubscriptionHandler)
	http.HandleFunc("GET /api/pro/sync", admin.SyncPremiumStatusHandler)
	http.HandleFunc("OPTIONS /api/pro/sync", admin.SyncPremiumStatusHandler)

	// Auth, Tutorial & Upload
	http.HandleFunc("/api/upload-document", auth.VerifyTokenMiddleware(admin.UploadDocumentHandler))
	http.Handle("/view-uploads/", http.StripPrefix("/view-uploads/", http.FileServer(http.Dir("./uploads"))))
	http.HandleFunc("/auth/check-email", admin.VerifierEmail)
	http.HandleFunc("/auth/inscription", admin.Inscription)
	http.HandleFunc("POST /auth/reset-password", auth.VerifyTokenMiddleware(admin.ResetPasswordHandler))
	http.HandleFunc("OPTIONS /auth/reset-password", admin.ResetPasswordHandler)
	http.HandleFunc("/admin/login", admin.Login)
	http.HandleFunc("/update-tutorial", admin.UpdateTutorialStatus)
	http.HandleFunc("/api/user/ecostats", admin.GetEcoStatsHandler)

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
	http.HandleFunc("/api/user/achats", admin.GetMyPurchases)
	http.HandleFunc("POST /api/mobile/payment-intent", admin.PaymentIntentMobile) // android payment intent
	http.HandleFunc("OPTIONS /api/mobile/payment-intent", admin.PaymentIntentMobile)

	// --- EVENEMENTS ---
	// --- EVENEMENTS ---
	http.HandleFunc("GET /admin/evenements", auth.VerifyTokenMiddleware(admin.GetAllEvenements)) // 🟢 CORRECTION ICI
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
	http.HandleFunc("POST /api/pro/subscribe", auth.VerifyTokenMiddleware(admin.CreateProSubscriptionHandler))
	http.HandleFunc("OPTIONS /api/pro/subscribe", admin.CreateProSubscriptionHandler)
	http.HandleFunc("POST /api/pro/annonces/sponsor", auth.VerifyTokenMiddleware(admin.ToggleSponsorHandler))
	http.HandleFunc("OPTIONS /api/pro/annonces/sponsor", admin.ToggleSponsorHandler)

	// boxes
	http.HandleFunc("GET /api/user/boxes", admin.GetMyBoxes)
	http.HandleFunc("GET /api/user/pickups/{id}", admin.GetUserPickupsHandler)
	http.HandleFunc("OPTIONS /api/user/pickups/{id}", admin.GetUserPickupsHandler)
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

	http.HandleFunc("OPTIONS /api/chat/history", admin.GetChatHistoryHandler)
	http.HandleFunc("GET /api/chat/history", auth.VerifyTokenMiddleware(admin.GetChatHistoryHandler))

	http.HandleFunc("GET /ws/chat", admin.ChatHandler)

	http.HandleFunc("POST /admin/evenements/desinscription", admin.DesinscriptionHandler)
	http.HandleFunc("OPTIONS /admin/evenements/desinscription", admin.DesinscriptionHandler)

	http.HandleFunc("GET /admin/finance/overview", admin.FinanceOverviewHandler)
	http.HandleFunc("OPTIONS /admin/finance/overview", admin.FinanceOverviewHandler)

	http.HandleFunc("GET /admin/finance/transactions", admin.AdminTransactionsHandler)
	http.HandleFunc("OPTIONS /admin/finance/transactions", admin.AdminTransactionsHandler)

	http.HandleFunc("OPTIONS /api/chat/conversations", admin.GetConversationsHandler)
	http.HandleFunc("GET /api/chat/conversations", auth.VerifyTokenMiddleware(admin.GetConversationsHandler))

	http.HandleFunc("/api/hardware/simulate-withdrawal", admin.SimulateWithdrawalHandler)
	http.HandleFunc("/api/hardware/simulate-deposit", admin.SimulateDepositHandler)
	http.HandleFunc("POST /api/web/checkout/evenement", admin.CreateEventCheckoutSession)
	http.HandleFunc("OPTIONS /api/web/checkout/evenement", admin.CreateEventCheckoutSession)

	
http.HandleFunc("GET /api/pro/projets", auth.VerifyTokenMiddleware(admin.GetProjetsHandler))
http.HandleFunc("POST /api/pro/projets/create", auth.VerifyTokenMiddleware(admin.CreateProjetHandler))
http.HandleFunc("OPTIONS /api/pro/projets", admin.GetProjetsHandler)
http.HandleFunc("OPTIONS /api/pro/projets/create", admin.CreateProjetHandler)
http.HandleFunc("DELETE /api/pro/projets/delete", auth.VerifyTokenMiddleware(admin.DeleteProjetHandler))
http.HandleFunc("PUT /api/pro/projets/update", auth.VerifyTokenMiddleware(admin.UpdateProjetHandler))
http.HandleFunc("OPTIONS /api/pro/projets/delete", admin.DeleteProjetHandler)
http.HandleFunc("OPTIONS /api/pro/projets/update", admin.UpdateProjetHandler)

http.HandleFunc("POST /api/pro/etapes/create", auth.VerifyTokenMiddleware(admin.CreateEtapeHandler))
http.HandleFunc("GET /api/pro/etapes", auth.VerifyTokenMiddleware(admin.GetEtapesHandler))
http.HandleFunc("DELETE /api/pro/etapes/delete", auth.VerifyTokenMiddleware(admin.DeleteEtapeHandler))
http.HandleFunc("PUT /api/pro/etapes/statut", auth.VerifyTokenMiddleware(admin.UpdateEtapeStatutHandler))
http.HandleFunc("OPTIONS /api/pro/etapes/create", admin.CreateEtapeHandler)
http.HandleFunc("OPTIONS /api/pro/etapes", admin.GetEtapesHandler)
http.HandleFunc("OPTIONS /api/pro/etapes/delete", admin.DeleteEtapeHandler)
http.HandleFunc("OPTIONS /api/pro/etapes/statut", admin.UpdateEtapeStatutHandler)

	// Cette ligne dit à Go : "Si on te demande une URL qui commence par /static/, va chercher le fichier dans le dossier static de mon PC"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
>>>>>>> origin/faty
	fmt.Println("test de : http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
