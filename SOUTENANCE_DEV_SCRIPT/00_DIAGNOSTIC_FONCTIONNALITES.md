# 00 - Diagnostic des fonctionnalites

Diagnostic base sur les fichiers presents dans le depot au 04/07/2026. Les fonctionnalites ci-dessous ne doivent pas etre presentees comme "finies" si le statut indique partielle, instable ou absente.

| Fonctionnalite | Statut | Fichiers concernes | Role concerne | Ecran a montrer | Risque demo | Recommandation |
|---|---|---|---|---|---|---|
| Page d'accueil publique | Fonctionnelle | `Frontend/landing.html`, `Frontend/a-propos.html`, `Frontend/script/config.js` | Visiteur | Accueil | Faible | Montrer en live |
| Inscription | Fonctionnelle | `Frontend/register.html`, `Frontend/script/register.js`, `API/route/auth.go`, `API/admin/users.go` | Visiteur | Inscription | Moyen si DB deja remplie/email existant | Montrer le formulaire, eviter de creer un compte en live si le temps est court |
| Connexion | Fonctionnelle | `Frontend/login.html`, `Frontend/script/login.js`, `API/admin/users.go`, `API/auth/jwt.go` | Tous | Login | Faible si compte de demo pret | Montrer en live |
| Mot de passe oublie | Partielle | `Frontend/reset-password.html`, `Frontend/script/resetPassword.js`, `API/admin/users.go`, `API/bdd/userReq.go` | Tous | Login/reset password | SMTP necessite config correcte | Mentionner, montrer ecran, eviter en live sauf SMTP verifie |
| Roles et protections | Fonctionnelle | `Frontend/script/auth_guard.js`, `API/auth/jwt.go`, `API/route/users.go` | Client, Pro, Salarie, Admin | Redirection dashboard | Moyen si localStorage corrompu | Montrer connexion avec 2 roles minimum |
| Dashboard particulier | Fonctionnelle | `Frontend/espClient.html`, `Frontend/script/dashClient.js` | Utilisateur | Espace client | Moyen selon donnees | Montrer en live avec donnees prechargees |
| Creation annonce | Fonctionnelle | `Frontend/espClient.html`, `Frontend/script/dashClient.js`, `API/admin/annonce.go`, `API/bdd/annonceReq.go` | Utilisateur / Pro | Formulaire annonce | Upload image peut dependre Docker/uploads | Montrer formulaire et/ou creer une annonce simple |
| Liste annonces utilisateur | Fonctionnelle | `Frontend/espClient.html`, `Frontend/script/dashClient.js`, `/mes-annonces` | Utilisateur / Pro | Mes annonces | Faible | Montrer en live |
| Catalogue annonces | Fonctionnelle | `Frontend/annonceAll.html`, `Frontend/script/annonces/annonceAll.js` | Utilisateur, Pro | Catalogue | Faible | Montrer en live |
| Detail annonce | Fonctionnelle | `Frontend/oneAnnonce.html`, `Frontend/script/annonces/oneAnnonce.js` | Utilisateur, Pro | Detail annonce | Moyen si paiement/achat | Montrer detail, eviter paiement reel si non prepare |
| Workflow annonce | Fonctionnelle | `API/bdd/annonceReq.go`, `API/admin/annonce.go`, `Frontend/admin_validations.html`, `Frontend/script/admin/annonce.js` | Admin / Salarie | Validations admin | Faible avec donnees en attente | Montrer avec annonce preexistante |
| Validation / refus admin | Fonctionnelle | `Frontend/admin_validations.html`, `API/route/annonces.go`, `API/route/evenements.go`, `API/route/articles.go` | Admin / Salarie | Centre validation | Moyen si aucune donnee en attente | Preparer une annonce en attente avant la soutenance |
| Back-office utilisateurs | Fonctionnelle | `Frontend/admin_users.html`, `Frontend/script/admin/user.js`, `API/admin/users.go` | Admin / Salarie | Admin utilisateurs | Faible | Montrer recherche, roles, statut |
| Bannissement utilisateur | Fonctionnelle | `API/bdd/userReq.go`, `API/admin/users.go`, `Frontend/script/salarie/forum.js` | Salarie / Admin | Moderation forum / utilisateurs | Moyen, action sensible | Montrer le bouton et expliquer, eviter de bannir un vrai compte |
| Emails validation/refus/ban/reset | Partielle | `API/bdd/userReq.go`, `API/admin/users.go` | Systeme | Pas d'ecran direct | Depend de SMTP | Mentionner comme integration technique, ne pas baser la demo dessus |
| Documents legaux / KYC | Partielle | `Frontend/register.html`, `API/admin/users.go`, `db/init.sql` table `documents_legaux` | Pro | Inscription pro / admin users | Stockage present mais parcours complet a verifier | Montrer comme element de validation, ne pas promettre KYC complet |
| Dashboard professionnel | Fonctionnelle | `Frontend/espPro.html`, `Frontend/script/dashPro.js`, `API/route/pro.go` | Pro | Espace Pro | Beaucoup de modules, risque de lenteur | Montrer seulement annonces + projets + abonnement si pret |
| Projets d'upcycling | Fonctionnelle | `Frontend/script/dashPro.js`, `API/admin/projet.go`, `API/admin/etape.go`, `API/bdd/projetReq.go`, `API/bdd/etapeReq.go` | Pro | Dashboard pro | Faible avec donnees | Montrer donnees existantes |
| Etapes projet | Fonctionnelle | `API/admin/etape.go`, `API/bdd/etapeReq.go` | Pro | Detail projet | Moyen | Montrer via donnee existante |
| Abonnement premium | Partielle / instable | `API/admin/stripe.go`, `API/admin/users.go`, `Frontend/script/dashPro.js` | Pro | Dashboard pro abonnement | Stripe necessite cles et retour checkout | Montrer ecran + expliquer, eviter paiement live sauf deja teste |
| Paiement annonce Stripe | Partielle / instable | `API/admin/stripe.go`, `Frontend/script/annonces/oneAnnonce.js` | Utilisateur / Pro | Detail annonce | Stripe externe + redirection | Preparer capture ou simulation |
| Paiement formation Stripe | Partielle / instable | `API/admin/stripe.go`, `Frontend/script/affichageEvt.js` | Utilisateur | Page evenements | Stripe externe | Eviter sauf compte Stripe test pret |
| Factures PDF | Fonctionnelle mais sensible | `API/admin/pdf.go`, `API/bdd/order.go`, `Frontend/profil.html`, `Frontend/admin_documents.html` | Client / Pro / Admin | Profil historique / documents | Generation depend chemin documents | Montrer PDF deja genere si present |
| Commissions / finances | Fonctionnelle | `Frontend/admin_finances.html`, `Frontend/script/admin/finance.js`, `API/bdd/order.go`, `API/admin/orderReq.go` | Admin | Finances | Donnees parfois faibles | Montrer via donnees prechargees |
| Conteneurs / box | Fonctionnelle | `Frontend/admin_conteneurs.html`, `Frontend/script/admin/box.js`, `API/admin/box.go`, `API/bdd/boxReq.go` | Admin | Logistique & Box | Moyen si aucune box libre | Montrer en live avec conteneurs precharges |
| Reservation box | Fonctionnelle | `API/admin/box.go`, `API/bdd/boxReq.go`, `Frontend/script/dashClient.js` | Utilisateur | Espace client | Moyen | Montrer statut et code via donnees existantes |
| Code PIN depot | Fonctionnelle | `API/bdd/boxReq.go`, `API/admin/box.go`, `Frontend/simulateur.html` | Utilisateur / systeme box | Simulateur | Moyen si code inconnu | Utiliser une donnee preparee |
| Code retrait / code-barres | Fonctionnelle | `API/bdd/boxReq.go`, `API/admin/box.go`, `Frontend/simulateur.html` | Acheteur / Pro | Simulateur | Moyen si code inconnu | Montrer la logique, eviter improvisation |
| QR code | Non trouve / a ne pas promettre | Recherche code: pas de module QR clair | Tous | Aucun | Eleve | Ne pas promettre QR code |
| Forum client | Fonctionnelle | `Frontend/forum.html`, `Frontend/script/forum.js`, `API/route/forum.go` | Utilisateur | Forum | Faible | Montrer rapide si temps |
| Moderation forum | Fonctionnelle | `Frontend/salarie/salarie_forum.html`, `Frontend/script/salarie/forum.js`, `API/admin/forum.go` | Salarie | Moderation | Faible | Montrer en live |
| Articles / conseils | Fonctionnelle | `Frontend/article.html`, `Frontend/salarie/salarie_content.html`, `Frontend/script/salarie/article.js`, `API/admin/article.go` | Public / Salarie / Admin | Conseils & News | Faible | Montrer liste et creation brouillon si besoin |
| Evenements / formations | Fonctionnelle | `Frontend/evenement.html`, `Frontend/salarie/salarie_events.html`, `API/admin/event.go`, `API/bdd/eventReq.go` | Public / Salarie / Admin | Evenements | Moyen si paiement formation | Montrer catalogue + creation en attente |
| Ressources PDF formation | Fonctionnelle mais a verifier | `API/admin/event.go`, `Frontend/script/salarie/event.js` | Salarie | Creation formation | Upload peut dependre Docker | Montrer champ, eviter gros fichier |
| Planning | Partielle | `Frontend/salarie/salarie_planning.html`, `Frontend/script/planning.js`, `API/route/planning.go` | Utilisateur / Salarie | Planning | Depend inscriptions | Montrer si donnees visibles |
| Notifications | Fonctionnelle partielle | `API/admin/notifications.go`, `Frontend/script/notifCloche.js`, `Frontend/script/salarie/userInfo.js` | Tous | Cloche notifications | OneSignal externe partiel | Montrer notifications internes, eviter promesse push mobile complete |
| Multilingue | Fonctionnelle | `API/admin/translation.go`, `Frontend/script/admin/translate.js`, `db/init.sql` tables `languages`, `translations` | Tous / Admin | Select langue | Faible | Montrer changement FR/EN |
| Upcycling Score | Fonctionnelle partielle | `API/bdd/boxReq.go`, `API/bdd/annonceReq.go`, `db/init.sql` table `upcycling_score` | Utilisateur | Dashboard client | Calcul lie aux retraits/depot | Montrer score existant, ne pas recalculer en live |
| Messagerie / WebSocket | Partielle | `Frontend/test_chat.html`, `Frontend/script/messagerie/messagerie.js`, `API/admin/chat.go` | Utilisateurs | Chat | WebSocket + donnees | Bonus seulement si deja ouvert |
| API Go | Fonctionnelle | `API/main.go`, `API/route/*.go`, `API/admin/*.go`, `API/bdd/*.go` | Technique | Code / endpoint health | Faible | Montrer code + endpoint `/api/languages` |
| Swagger / documentation API | Absente | Recherche `swagger/openapi` sans resultat | Technique | Aucun | Eleve | Non trouve / a ne pas promettre |
| Base SQL pre-remplie | Fonctionnelle | `db/init.sql`, `docker-compose.yml` | Technique | SQL / phpMyAdmin | Faible | Montrer tables et donnees |
| Docker Compose | Fonctionnelle | `docker-compose.yml`, `API/Dockerfile`, `Frontend/Dockerfile` | Technique | Terminal | Faible si conteneurs up | Montrer `docker compose ps` |
| Nginx reverse proxy | Fonctionnelle | `Frontend/nginx.conf` | Technique | Code config | Faible | Montrer `/api/` proxy vers backend |
| Site externe | Partielle | Compose prevu, URL depend infra | Public | Navigateur | Depend reseau/VM | Si disponible, montrer; sinon montrer Docker local |

## Conclusion diagnostic

Fonctionnalites les plus rentables en demo live : accueil, login, dashboard client, annonces, validation admin, utilisateurs/roles, dashboard salarie, evenements, multilingue, Docker/Nginx, API Go.

Fonctionnalites a montrer avec prudence : Stripe, emails SMTP, PDF, codes PIN/retrait, OneSignal, mot de passe oublie.

Fonctionnalites a ne pas promettre : Swagger/OpenAPI, QR code complet.
