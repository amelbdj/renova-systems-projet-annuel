# Kit presentation DEV complet - UpcycleConnect

Document complet genere a partir des fichiers du kit.


---

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


---

# 01 - Script de presentation DEV 15 minutes

Objectif : tenir 14 min 30 pour garder 30 secondes de marge.

## 00:00 - 01:00

Personne : Faty  
Ecran : `landing.html` ou URL Docker `http://localhost:8088/`  
Action : ouvrir la page d'accueil, montrer menu, positionnement et parcours publics.  
Texte a dire :
"Bonjour, nous vous presentons UpcycleConnect, une plateforme metier dediee a l'upcycling et a l'economie circulaire. L'objectif n'est pas seulement de publier des annonces : la plateforme organise tout le cycle, depuis le depot d'un objet par un particulier, jusqu'a sa recuperation par un professionnel, avec validation, logistique, evenements et administration."  
Ce que ca prouve : vision produit, adaptation au sujet.  
Point de notation : presentation commerciale et fonctionnelle.  
Transition : "Je vais maintenant montrer comment un utilisateur entre dans la plateforme."

## 01:00 - 02:10

Personne : Faty  
Ecran : `login.html` puis choix d'espace.  
Action : montrer inscription rapidement, puis se connecter avec un compte de demo deja pret.  
Texte a dire :
"Nous avons gere plusieurs profils : particulier, professionnel, salarie et administrateur. Le parcours commence par une inscription ou une connexion. Le role est conserve dans la session et permet de rediriger l'utilisateur vers le bon espace."  
Ce que ca prouve : authentification, roles, UX.  
Point de notation : site utilisable et parcours clair.  
Transition : "On commence par le parcours particulier, car c'est le point d'entree principal de la plateforme."

## 02:10 - 03:40

Personne : Faty  
Ecran : `espClient.html`  
Action : montrer dashboard, mes annonces, bouton creation annonce, score ou statistiques si visibles.  
Texte a dire :
"Ici, le particulier peut suivre ses annonces, creer un depot, consulter ses reservations et voir ses indicateurs. Le formulaire d'annonce structure les informations utiles au traitement : titre, categorie, prix, et image si besoin. Une annonce creee passe ensuite dans un workflow de validation avant d'etre visible publiquement."  
Ce que ca prouve : dashboard client, creation annonce, workflow.  
Point de notation : fonctionnalite metier centrale.  
Transition : "Une fois l'annonce soumise, elle n'est pas publiee directement : elle passe par le back-office."

## 03:40 - 05:10

Personne : Amel  
Ecran : `admin_validations.html`  
Action : montrer les onglets annonces/evenements/contenus, une annonce en attente, boutons approuver/refuser.  
Texte a dire :
"Cote back-office, nous avons separe la saisie utilisateur de la publication. Les administrateurs et salaries peuvent verifier les annonces, les evenements et les contenus avant mise en ligne. Cela evite les contenus incorrects et permet de garder une qualite de catalogue."  
Ce que ca prouve : validation admin, moderation, logique metier.  
Point de notation : principaux cas de traitement.  
Transition : "Je vais montrer rapidement comment cette logique est reliee a l'API Go."

## 05:10 - 06:20

Personne : Amel  
Ecran : code `API/route/annonces.go`, `API/admin/annonce.go`, `API/bdd/annonceReq.go`  
Action : montrer une route de validation, handler, requete SQL.  
Texte a dire :
"La logique n'est pas seulement dans le front. Le front appelle une API Go. Les routes sont regroupees dans le dossier `route`, les handlers HTTP dans `admin`, et les requetes SQL dans `bdd`. Par exemple, valider une annonce appelle une route API qui met a jour son statut en base."  
Ce que ca prouve : separation front/back, API Go, SQL.  
Point de notation : comprehension technique.  
Transition : "Apres validation, l'annonce peut etre consultee dans le catalogue."

## 06:20 - 07:20

Personne : Faty  
Ecran : `annonceAll.html` puis `oneAnnonce.html`  
Action : ouvrir catalogue, rechercher/filtrer si possible, ouvrir detail annonce.  
Texte a dire :
"Le catalogue permet aux utilisateurs et aux professionnels de consulter les annonces validees. On retrouve une experience plus commerciale : affichage des objets, detail, prix et possibilite d'achat ou de recuperation selon le cas."  
Ce que ca prouve : catalogue public, donnees validees.  
Point de notation : demo visible et concrete.  
Transition : "La particularite du sujet est aussi la logistique autour des box et des codes."

## 07:20 - 08:40

Personne : Amel  
Ecran : `admin_conteneurs.html` puis eventuellement `simulateur.html`  
Action : montrer conteneurs, box, statuts, expliquer code PIN depot/retrait.  
Texte a dire :
"Pour gerer la recuperation physique, nous avons ajoute une partie logistique. Les conteneurs contiennent des box avec des statuts. Lors d'une reservation, le systeme genere des codes de depot ou de retrait stockes en base dans l'historique des conteneurs. Le simulateur permet de representer le comportement d'une borne ou d'un scanner."  
Ce que ca prouve : box, PIN, statut, traitement metier.  
Point de notation : adaptation au besoin client.  
Transition : "Le deuxieme espace important est celui du professionnel."

## 08:40 - 10:00

Personne : Faty  
Ecran : `espPro.html`  
Action : montrer dashboard pro, annonces pro, projets d'upcycling, abonnement sans lancer Stripe si risque.  
Texte a dire :
"Le professionnel dispose d'un espace adapte a son usage : il peut consulter ses annonces, gerer des projets d'upcycling et suivre des etapes de transformation. Nous avons aussi prepare une logique premium, avec Stripe pour les abonnements, mais en demo nous privilegions l'ecran et les donnees existantes pour rester fiables."  
Ce que ca prouve : role pro, projets, logique premium.  
Point de notation : roles differencies, fonctionnalites metier.  
Transition : "La plateforme ne se limite pas aux annonces : elle propose aussi des contenus et formations."

## 10:00 - 11:10

Personne : Ndoya  
Ecran : `evenement.html` puis `salarie_events.html`  
Action : montrer evenements publics, puis cote salarie creation/validation requise.  
Texte a dire :
"Nous avons ajoute un module evenements et formations. Cote public, les utilisateurs consultent le catalogue. Cote salarie, un responsable peut creer un evenement ou une formation, ajouter des informations pratiques et des ressources PDF. Comme les annonces, ces elements peuvent passer par validation."  
Ce que ca prouve : module evenements/formations, ressources.  
Point de notation : richesse fonctionnelle.  
Transition : "Cote salarie, il existe aussi une mission de moderation et de contenus."

## 11:10 - 12:10

Personne : Ndoya  
Ecran : `salarie_forum.html`, `salarie_content.html`  
Action : montrer moderation forum, articles/conseils.  
Texte a dire :
"L'espace salarie n'est pas seulement administratif. Il permet de moderer le forum, suivre les contenus et participer a la vie de la plateforme. Cela repond au besoin d'une entreprise qui doit encadrer les echanges et maintenir la qualite des informations."  
Ce que ca prouve : moderation, contenu, role salarie.  
Point de notation : back-office operationnel.  
Transition : "Je vais maintenant montrer l'organisation technique de deploiement."

## 12:10 - 13:30

Personne : Ndoya  
Ecran : terminal + `docker-compose.yml` + `Frontend/nginx.conf`  
Action : montrer `docker compose ps`, services mysql/backend/frontend, port front Docker, proxy `/api`.  
Texte a dire :
"Le projet est prepare pour un deploiement Docker Compose. Nous avons trois services : MySQL pour la base, le backend Go expose en interne sur le port 8081, et le frontend Nginx. Nginx sert les pages statiques et redirige les appels `/api` vers le backend, ce qui permet d'avoir une architecture proche d'un deploiement reel."  
Ce que ca prouve : Docker, Nginx, separation, deploiement.  
Point de notation : stack attendue.  
Transition : "Pour terminer, Amel va montrer la base de donnees et les preuves techniques."

## 13:30 - 14:20

Personne : Amel  
Ecran : `db/init.sql` ou phpMyAdmin  
Action : montrer tables `utilisateur`, `annonce`, `evenement`, `conteneur`, `box`, `translations`, `document`.  
Texte a dire :
"La base SQL est pre-remplie pour permettre une demonstration stable. On retrouve les utilisateurs et roles, les annonces, les evenements, la logistique des box, les traductions, les documents et les paiements. Cela montre que l'application n'est pas une maquette : elle manipule de vraies donnees relationnelles."  
Ce que ca prouve : SQL, persistance, donnees demo.  
Point de notation : base de donnees et fonctionnement reel.  
Transition : "Nous concluons sur la valeur livree."

## 14:20 - 14:30

Personne : Faty  
Ecran : accueil ou admin dashboard.  
Action : conclusion.  
Texte a dire :
"Pour conclure, UpcycleConnect repond au sujet avec une plateforme multi-roles : particuliers, professionnels, salaries et administrateurs. Nous avons livre un parcours complet autour des annonces, de la validation, de la logistique, des evenements, de la base SQL et du deploiement Docker. Les integrations externes comme Stripe et SMTP sont preparees, mais notre demo met volontairement en avant les fonctionnalites les plus fiables."  
Ce que ca prouve : maitrise du perimetre, lucidite.  
Point de notation : conclusion claire et professionnelle.


---

# 02 - Deroule demo clic par clic

## Preparation avant de commencer

1. Lancer Docker : `docker compose up -d --build`.
2. Verifier : `docker compose ps`.
3. Ouvrir le front Docker : `http://localhost:8088/` si le conteneur expose ce port, sinon verifier la colonne PORTS.
4. Ouvrir aussi un terminal dans le dossier projet.
5. Ouvrir phpMyAdmin ou `db/init.sql` pour la preuve SQL.
6. Preparer des comptes de demo. Les emails existent dans `db/init.sql`, mais les mots de passe sont hashes. Si vous ne connaissez pas les mots de passe, resetter avant la soutenance.

## Comptes a preparer

Comptes visibles dans `db/init.sql` :

- `test.client@renova.test` role Utilisateur
- `test.admin@renova.test` role Administrateur
- `test.salarie@renova.test` role Salarie
- `test.pro@renova.test` role Pro
- `test.user@test.fr` role Utilisateur
- `test.salarie@test.fr` role Salarie
- `test.admin@test.fr` role Administrateur
- `test.pro@test.fr` role Pro

Mot de passe : non lisible dans le dump car hash bcrypt. A preparer avant demo, par exemple `Test123!`.

## Demo principale

1. Ouvrir `http://localhost:8088/`.
   - Montrer : accueil, menu, positionnement UpcycleConnect.
   - Dire : "plateforme metier d'economie circulaire".
   - Eviter : rester trop longtemps sur le design.

2. Aller sur `login.html`.
   - Montrer : choix des espaces.
   - Se connecter avec un compte utilisateur.
   - Page attendue : dashboard particulier.
   - Eviter : creer un compte en live si le temps est serre.

3. Dashboard particulier.
   - Cliquer sur "Mes annonces" ou section annonce.
   - Montrer : annonces existantes, creation annonce.
   - Dire : "l'annonce passe ensuite en validation".
   - Eviter : upload image lourd.

4. Aller sur admin.
   - Se connecter avec compte admin.
   - Ouvrir `admin_validations.html`.
   - Montrer : onglets annonces, evenements, contenus.
   - Cliquer seulement si une donnee de test est prete.
   - Eviter : valider toutes les donnees et vider la demo.

5. Montrer code API.
   - Ouvrir `API/route/annonces.go`.
   - Ouvrir `API/admin/annonce.go`.
   - Ouvrir `API/bdd/annonceReq.go`.
   - Montrer : route -> handler -> requete SQL.

6. Catalogue.
   - Ouvrir `annonceAll.html`.
   - Ouvrir une annonce.
   - Montrer : seuls les contenus valides apparaissent.
   - Eviter : lancer paiement Stripe si non prepare.

7. Logistique.
   - Ouvrir `admin_conteneurs.html`.
   - Montrer : conteneurs, box, statuts.
   - Ouvrir `simulateur.html` si un code de test existe.
   - Dire : "le code simule l'interaction avec une borne".

8. Professionnel.
   - Se connecter compte pro.
   - Ouvrir `espPro.html`.
   - Montrer : annonces, projets, etapes, premium.
   - Eviter : checkout Stripe live sauf deja teste.

9. Evenements et formations.
   - Ouvrir `evenement.html`.
   - Puis salarie : `salarie/salarie_events.html`.
   - Montrer : creation evenement/formation, PDF ressources.
   - Eviter : upload gros PDF.

10. Salarie moderation.
    - Ouvrir `salarie/salarie_forum.html`.
    - Montrer : moderation messages et bannissement possible.
    - Eviter : bannir un compte utile.

11. Docker / Nginx.
    - Terminal : `docker compose ps`.
    - Ouvrir `docker-compose.yml`.
    - Ouvrir `Frontend/nginx.conf`.
    - Dire : "Nginx sert le front et proxy les appels API".

12. Base SQL.
    - Ouvrir `db/init.sql`.
    - Montrer tables : `utilisateur`, `annonce`, `evenement`, `conteneur`, `box`, `translations`, `document`.
    - Dire : "donnees prechargees pour demo fiable".

13. Conclusion.
    - Revenir accueil ou admin dashboard.
    - Dire la phrase de conclusion du script.


---

# 03 - Captures a preparer

| Nom fichier recommande | Page | Pourquoi | Moment du script | Phrase si capture utilisee |
|---|---|---|---|---|
| `01_accueil.png` | Accueil | Prouver le site public | 00:00 | "Pour rester dans le temps, voici l'ecran d'accueil de la plateforme." |
| `02_login_roles.png` | Login | Montrer choix des espaces | 01:00 | "On retrouve ici les differents espaces relies aux roles." |
| `03_dashboard_client.png` | `espClient.html` | Prouver espace particulier | 02:10 | "Ce dashboard montre le suivi des annonces et des actions utilisateur." |
| `04_creation_annonce.png` | Formulaire annonce | Fonction centrale | 02:30 | "Le formulaire structure les informations avant validation." |
| `05_annonce_attente.png` | Admin validations | Workflow attente | 03:40 | "L'annonce arrive ici avant publication." |
| `06_annonce_validee_catalogue.png` | Catalogue | Prouver publication | 06:20 | "Apres validation, l'annonce devient visible dans le catalogue." |
| `07_admin_utilisateurs.png` | Admin users | Roles et gestion | Bonus admin | "Le back-office permet de suivre les roles et statuts." |
| `08_admin_conteneurs.png` | Admin conteneurs | Logistique box | 07:20 | "La logistique repose sur des conteneurs et box suivis en base." |
| `09_simulateur_pin.png` | Simulateur | PIN depot/retrait | 07:50 | "Le simulateur represente l'interaction avec une borne." |
| `10_dashboard_pro.png` | `espPro.html` | Role professionnel | 08:40 | "L'espace pro regroupe annonces, projets et options premium." |
| `11_projets_upcycling.png` | Projets pro | Upcycling concret | 09:15 | "Les professionnels peuvent suivre leurs projets de transformation." |
| `12_evenements_publics.png` | `evenement.html` | Catalogue evenements | 10:00 | "Le module evenements ajoute une dimension formation." |
| `13_salarie_events.png` | Salarie events | Creation formation | 10:30 | "Les salaries peuvent creer des evenements soumis a validation." |
| `14_moderation_forum.png` | Salarie forum | Moderation | 11:10 | "La plateforme prevoit aussi la moderation des echanges." |
| `15_multilingue.png` | Select langue | FR/EN | A glisser entre pages | "La traduction est geree par la table translations et l'API." |
| `16_finances.png` | Admin finances | Commissions | Bonus | "Le back-office finance exploite les commandes et paiements." |
| `17_pdf_facture.png` | Facture PDF | Preuve document | Bonus | "Les documents PDF sont generes cote backend." |
| `18_docker_compose_ps.png` | Terminal | Deploiement | 12:10 | "Les trois services Docker sont lances : front, API et base." |
| `19_nginx_proxy.png` | `Frontend/nginx.conf` | Reverse proxy | 12:40 | "Nginx route `/api` vers le backend Go." |
| `20_db_tables.png` | phpMyAdmin ou SQL | Base pre-remplie | 13:30 | "La demonstration repose sur une base SQL relationnelle." |

Priorite absolue : captures 01 a 10, 18, 19, 20.


---

# 04 - Plan B demo

Regle orale : ne jamais dire "normalement ca marche". Dire plutot : "Pour securiser la demonstration et respecter le temps, nous allons utiliser le jeu de donnees prepare."

| Fonctionnalite | Ce qui peut buguer | Reaction | Capture | Phrase professionnelle |
|---|---|---|---|---|
| Docker | Port deja utilise | Montrer `docker compose ps` et changer de port si besoin | `18_docker_compose_ps.png` | "L'environnement Docker est pret ; ici le port local est adapte a la machine de demonstration." |
| Connexion | Mot de passe oublie / role incorrect | Utiliser un autre compte ou reset avant demo | `02_login_roles.png` | "Nous utilisons un compte de demonstration preconfigure pour eviter de perdre du temps sur la saisie." |
| Creation annonce | Upload ou validation formulaire | Montrer annonce deja creee | `04_creation_annonce.png`, `05_annonce_attente.png` | "Le formulaire existe ; pour garder le rythme, nous montrons l'annonce deja inseree dans le workflow." |
| Validation admin | Aucune donnee en attente | Creer avant ou montrer capture | `05_annonce_attente.png` | "Le workflow est visible ici avec une donnee de demonstration deja preparee." |
| Catalogue | Catalogue vide | Verifier `db/init.sql` / utiliser capture | `06_annonce_validee_catalogue.png` | "La publication depend du statut valide ; voici une annonce validee visible cote catalogue." |
| Stripe annonce | Redirection externe lente | Ne pas lancer paiement | `17_pdf_facture.png` | "L'integration Stripe est presente cote API ; en presentation, nous montrons le flux sans declencher de paiement externe." |
| Stripe abonnement | Session checkout impossible | Montrer ecran abonnement + code API | `10_dashboard_pro.png` | "La logique premium est integree, mais nous evitons une dependance externe pendant la demo chronometree." |
| PDF | Fichier absent | Montrer code `API/admin/pdf.go` et capture PDF | `17_pdf_facture.png` | "La generation PDF est realisee cote backend ; voici le resultat attendu avec une donnee preparee." |
| SMTP | Mail non recu | Montrer code et ecran | capture login/reset | "L'envoi mail depend de la configuration SMTP ; la route et le parcours sont implementes." |
| OneSignal | Push non recu | Montrer notifications internes | capture notifications | "Les notifications internes sont visibles ; le push externe depend de la configuration OneSignal." |
| Box / PIN | Code inconnu | Montrer admin conteneurs + simulateur sans valider | `08_admin_conteneurs.png`, `09_simulateur_pin.png` | "Les codes sont generes et stockes dans l'historique ; nous montrons la logique sans consommer une donnee utile." |
| Multilingue | Traductions chargees lentement | Montrer table translations/API | `15_multilingue.png` | "Les textes sont servis par l'API de traduction et la table SQL." |
| Salarie events | Upload PDF bloque | Montrer formulaire et donnees existantes | `13_salarie_events.png` | "Le formulaire accepte des ressources ; nous evitons un upload lourd en live." |
| Moderation | Bannir mauvais compte | Ne pas cliquer ban | `14_moderation_forum.png` | "L'action existe mais nous ne l'executons pas sur un compte de demonstration utile." |
| Swagger | Prof demande documentation | Dire absent | aucune | "Nous n'avons pas integre Swagger ; les routes sont structurees dans `API/route` et peuvent etre documentees en evolution." |


---

# 05 - Checklist points a montrer

## A. Obligatoire a montrer

- [ ] Site public : `landing.html`
- [ ] Connexion : `login.html`
- [ ] Au moins deux roles : Utilisateur + Admin, idealement aussi Salarie ou Pro
- [ ] Dashboard particulier : `espClient.html`
- [ ] Creation ou formulaire annonce
- [ ] Catalogue annonce : `annonceAll.html`
- [ ] Workflow / statut : annonce en attente puis validation admin
- [ ] Back-office admin : `admin_dashboard.html` ou `admin_validations.html`
- [ ] Utilisateurs / roles : `admin_users.html`
- [ ] Validation / moderation : admin validations ou salarie forum
- [ ] API Go : `API/route/*.go`, `API/admin/*.go`, endpoint `/api/languages`
- [ ] Base SQL : `db/init.sql`, tables principales
- [ ] Docker : `docker compose ps`
- [ ] Nginx : `Frontend/nginx.conf`
- [ ] Site Docker : `http://localhost:8088/` ou port affiche par Docker

## B. Bonus a montrer si stable

- [ ] Upcycling Score : dashboard client / `API/bdd/boxReq.go`
- [ ] Commission : admin finances
- [ ] Facture PDF : document deja genere
- [ ] Code PIN : simulateur depot/retrait
- [ ] Conteneur : admin conteneurs
- [ ] Notifications : cloche ou panneau notifications
- [ ] Multilingue : FR/EN
- [ ] Paiement : ecran Stripe ou capture, pas forcement live
- [ ] Statistiques : admin dashboard / finances / salarie forum
- [ ] Evenements : catalogue public
- [ ] Formations : creation salarie + ressources PDF
- [ ] Projets d'upcycling : dashboard pro

## A ne pas promettre

- [ ] Swagger/OpenAPI complet : non trouve
- [ ] QR code complet : non trouve clairement dans le code
- [ ] Paiement reel sans preparation Stripe
- [ ] Envoi mail SMTP sans test juste avant


---

# 06 - Script version urgence 10 minutes

## 00:00 - 01:00 - Faty

Ecran : accueil.  
Dire : "UpcycleConnect est une plateforme metier pour connecter particuliers, professionnels et equipe interne autour de l'upcycling."  
Montrer : menu, promesse, acces.

## 01:00 - 02:00 - Faty

Ecran : login.  
Dire : "Le site gere plusieurs roles et redirige chacun vers son espace."  
Montrer : connexion utilisateur.

## 02:00 - 03:30 - Faty

Ecran : dashboard particulier.  
Dire : "Le particulier peut creer et suivre ses annonces. Une annonce n'est pas publiee directement : elle passe par validation."  
Montrer : formulaire annonce ou annonce existante.

## 03:30 - 05:00 - Amel

Ecran : admin validations.  
Dire : "Le back-office controle les annonces, evenements et contenus avant publication."  
Montrer : annonce en attente, boutons valider/refuser.

## 05:00 - 06:00 - Amel

Ecran : code API.  
Dire : "La logique est portee par une API Go. Les routes sont separees des handlers et des requetes SQL."  
Montrer : `route/annonces.go`, `admin/annonce.go`, `bdd/annonceReq.go`.

## 06:00 - 07:00 - Faty

Ecran : catalogue annonces.  
Dire : "Une fois validee, l'annonce est visible dans le catalogue."  
Montrer : liste + detail.

## 07:00 - 08:00 - Ndoya

Ecran : dashboard pro.  
Dire : "Le professionnel dispose d'un espace dedie pour ses annonces, ses projets d'upcycling et les options premium."  
Montrer : `espPro.html`.

## 08:00 - 09:00 - Ndoya

Ecran : Docker/Nginx.  
Dire : "Le projet se deploie avec Docker Compose : MySQL, backend Go et frontend Nginx avec proxy API."  
Montrer : `docker compose ps`, `nginx.conf`.

## 09:00 - 10:00 - Amel

Ecran : SQL.  
Dire : "La base est pre-remplie avec utilisateurs, annonces, evenements, box et traductions pour garantir une demo stable."  
Conclusion : "Nous avons livre une plateforme multi-roles fonctionnelle, avec workflows metier, API Go, base SQL et deploiement Docker."


---

# 07 - Questions post presentation

1. Pourquoi avoir choisi Go pour l'API ?  
Reponse : Go est simple a compiler, rapide, et adapte pour exposer des routes HTTP claires. Dans notre projet, les routes sont dans `API/route`, les handlers dans `API/admin`, et les requetes SQL dans `API/bdd`.

2. Pourquoi separer front et back ?  
Reponse : Le front gere l'affichage et l'experience utilisateur, tandis que l'API centralise la logique metier, la securite, les roles et la base de donnees.

3. Ou est la logique metier ?  
Reponse : Principalement dans `API/admin` pour les handlers et `API/bdd` pour les operations SQL, par exemple validation annonce, box, score, paiement.

4. Comment sont geres les roles ?  
Reponse : Le role est stocke dans le JWT et dans le localStorage cote front. L'API verifie les roles avec `VerifyRoleMiddleware`.

5. Pourquoi verifier les roles cote API ?  
Reponse : Le front peut etre modifie par l'utilisateur. La vraie protection doit donc etre cote serveur.

6. Comment fonctionne le workflow annonce ?  
Reponse : Une annonce est creee avec un statut, puis l'admin ou salarie peut la valider ou la refuser via les routes admin.

7. Comment eviter qu'une annonce non validee soit visible ?  
Reponse : Le catalogue utilise les annonces validees, notamment via les requetes de `GetValidatedAnnonces`.

8. Comment fonctionne l'Upcycling Score ?  
Reponse : Le score est calcule lors de certaines actions logistiques dans `boxReq.go`, puis ajoute a l'utilisateur selon l'objet et la recuperation.

9. Comment sont calculees les commissions ?  
Reponse : Les commandes et paiements sont stockes en base. Les vues finances utilisent `order.go` pour calculer volume et revenu.

10. Comment fonctionne la base de donnees ?  
Reponse : MySQL contient les tables utilisateurs, annonces, evenements, box, paiements, documents et traductions. Le dump est dans `db/init.sql`.

11. Comment importer la base ?  
Reponse : En Docker, `db/init.sql` est monte dans `/docker-entrypoint-initdb.d/init.sql` au premier demarrage du conteneur MySQL.

12. Comment exporter la base ?  
Reponse : Avec phpMyAdmin ou `mysqldump`, puis on versionne un dump propre si necessaire.

13. Comment lancer le projet en local Docker ?  
Reponse : `docker compose up -d --build`, puis ouvrir le port du conteneur frontend visible dans `docker compose ps`.

14. Comment fonctionne Docker Compose ?  
Reponse : Il lance trois services : MySQL, backend Go et frontend Nginx, connectes sur le reseau `uc_net`.

15. Comment fonctionne Nginx ?  
Reponse : Nginx sert les fichiers front et redirige `/api` vers le backend avec `proxy_pass http://backend:8081`.

16. Comment prouver que ce n'est pas seulement du localhost WAMP ?  
Reponse : Montrer `docker compose ps`, ouvrir l'URL Docker, puis montrer que les appels passent par `/api`.

17. Comment ajouter une route API ?  
Reponse : Ajouter le handler dans `API/admin`, la route dans `API/route`, et la fonction SQL dans `API/bdd` si besoin.

18. Comment proteger une route ?  
Reponse : Entourer le handler avec `auth.VerifyTokenMiddleware` ou `auth.VerifyRoleMiddleware`.

19. Comment ajouter un champ a une annonce ?  
Reponse : Modifier la table SQL, le modele Go, la requete `CreateAnnonce/UpdateAnnonce`, puis le formulaire front.

20. Que faire si la base ne repond pas ?  
Reponse : Verifier `docker compose ps`, les logs MySQL, les variables `DB_HOST`, `DB_USER`, `DB_PASS`.

21. Que faire si l'API renvoie 500 ?  
Reponse : Lire les logs backend, identifier la route appelee, verifier la requete SQL et les donnees envoyees.

22. Que se passe-t-il si un non-admin appelle une route admin ?  
Reponse : Les routes sensibles utilisent `VerifyRoleMiddleware`, donc l'API doit retourner une erreur d'acces.

23. Pourquoi certains endpoints ne sont pas proteges ?  
Reponse : Certains endpoints publics servent le catalogue ou les traductions. Les actions sensibles doivent etre protegees.

24. Comment fonctionne Stripe ?  
Reponse : Le backend cree des sessions checkout ou payment intent avec les cles Stripe. Les webhooks peuvent mettre a jour les statuts.

25. Pourquoi ne pas faire le paiement en live ?  
Reponse : Stripe depend d'un service externe. En soutenance courte, on montre le code et les ecrans, sauf si l'environnement test est valide juste avant.

26. Comment sont generees les factures PDF ?  
Reponse : `API/admin/pdf.go` genere les factures ou contrats et enregistre le chemin dans la table `document`.

27. Comment fonctionne le multilingue ?  
Reponse : Les textes sont stockes dans `translations`, servis par `/api/translations`, puis appliques dans le front.

28. Comment ajouter une langue ?  
Reponse : Ajouter une ligne dans `languages`, ajouter les traductions correspondantes et utiliser l'interface admin de traduction.

29. Comment fonctionnent les box ?  
Reponse : Les box appartiennent a des conteneurs. Les reservations et retraits sont traces dans `historique_conteneurs`.

30. Ou sont stockes les codes PIN ?  
Reponse : Dans l'historique des conteneurs, notamment les champs de code lies au depot ou a la recuperation.

31. Pourquoi un simulateur hardware ?  
Reponse : Pour representer le comportement d'une borne sans avoir de materiel physique en soutenance.

32. Comment fonctionnent les evenements ?  
Reponse : Les salaries peuvent creer evenements/formations, les utilisateurs peuvent les consulter et s'inscrire.

33. Comment sont gerees les ressources PDF de formation ?  
Reponse : Le front salarie envoie les fichiers, le backend les stocke et les relie a l'evenement.

34. Comment fonctionne le forum ?  
Reponse : Les utilisateurs creent des sujets/messages, et les salaries peuvent moderer les messages.

35. Comment fonctionne le bannissement ?  
Reponse : L'API met a jour l'utilisateur en base, cree une notification et peut envoyer un mail de bannissement.

36. Comment fonctionne le mot de passe oublie ?  
Reponse : L'utilisateur demande un reset, l'API cree un token, envoie un lien par mail, puis `reset-password.html` renvoie le nouveau mot de passe.

37. Est-ce que le SMTP est obligatoire ?  
Reponse : Oui pour recevoir le mail en vrai. Sans SMTP, le parcours peut etre montre mais l'envoi reel ne doit pas etre promis.

38. Quelles fonctionnalites sont les plus stables ?  
Reponse : Accueil, login, roles, annonces, validation admin, catalogue, SQL, Docker/Nginx, multilingue.

39. Quelles fonctionnalites sont plus risquees ?  
Reponse : Stripe, SMTP, OneSignal, PDF si les dossiers ne sont pas montes, et les codes PIN si la donnee n'est pas preparee.

40. Qu'est-ce qui manque ?  
Reponse : Swagger/OpenAPI n'est pas trouve. Le QR code complet n'est pas clairement implemente. Certaines integrations externes demandent une configuration.

41. Comment ameliorer le projet apres soutenance ?  
Reponse : Ajouter documentation OpenAPI, tests automatises, seeds de comptes demo, journal d'audit admin, et stabiliser les integrations externes.

42. Pourquoi utiliser des donnees pre-remplies ?  
Reponse : Pour garantir une demonstration fiable dans un temps limite et montrer le workflow sans attendre la creation de chaque etape.


---

# 08 - Comptes test et donnees demo

## Comptes trouves dans le projet

Source : `db/init.sql`, table `utilisateur`.

| Email | Role | Statut | Usage conseille |
|---|---|---|---|
| `test.client@renova.test` | Utilisateur | Valide | Demo dashboard particulier |
| `test.admin@renova.test` | Administrateur | Valide | Demo back-office |
| `test.salarie@renova.test` | Salarie | Valide | Demo moderation / evenements |
| `test.pro@renova.test` | Pro | Valide | Demo espace pro |
| `test.user@test.fr` | Utilisateur | Valide | Compte secours client |
| `test.salarie@test.fr` | Salarie | Valide | Compte secours salarie |
| `test.admin@test.fr` | Administrateur | Valide | Compte secours admin |
| `test.pro@test.fr` | Pro | Valide | Compte secours pro |

## Mot de passe

Les mots de passe sont stockes en bcrypt dans la base. Ils ne sont donc pas lisibles depuis le dump SQL.

Action recommandee avant soutenance :

1. Choisir un mot de passe commun de demo, par exemple `Test123!`.
2. Le definir pour les 4 comptes principaux.
3. Verifier une connexion pour chaque role.
4. Noter les identifiants sur une fiche papier pour eviter la panique.

## Donnees utiles deja presentes

### Annonces

Source : table `annonce`.

Utilisation :
- montrer catalogue ;
- montrer mes annonces ;
- montrer statut annonce ;
- preparer une annonce en attente pour validation.

Action avant demo :
- garder au moins une annonce validee ;
- garder au moins une annonce en attente ;
- ne pas tout valider avant la soutenance.

### Evenements

Source : table `evenement`.

Donnees presentes :
- formations ;
- ateliers ;
- evenements valides ;
- au moins un evenement en attente selon dump.

Utilisation :
- montrer catalogue evenements ;
- montrer creation salarie ;
- montrer validation admin.

### Conteneurs et box

Sources : tables `conteneur`, `box`, `historique_conteneurs`.

Utilisation :
- montrer administration logistique ;
- expliquer depot/retrait ;
- montrer code PIN ou code de recuperation si donnee prete.

Action avant demo :
- reperer un code de depot/retrait utilisable ;
- ne pas consommer ce code avant la soutenance.

### Traductions

Sources : tables `languages`, `translations`.

Utilisation :
- montrer select langue FR/EN ;
- montrer route API `/api/translations?lang=fr`.

### Documents / PDF

Sources : table `document`, `API/admin/pdf.go`.

Utilisation :
- montrer factures/contrats si fichiers presents ;
- sinon montrer code et capture.

### Projets pro

Sources : tables `projet_pro`, `etapes_projet`.

Utilisation :
- montrer espace pro ;
- montrer suivi de projets d'upcycling.

## Jeu de demo conseille

- Client : creer ou afficher une annonce.
- Admin : valider/refuser une annonce.
- Salarie : creer un evenement ou moderer le forum.
- Pro : consulter catalogue/projets/premium.

## Donnees a ne pas modifier pendant la demo

- Le compte admin principal.
- Le compte salarie principal.
- Une annonce en attente.
- Une annonce validee.
- Un conteneur avec box disponible.
- Un code PIN prepare.

## Verification 30 minutes avant

- [ ] `docker compose ps` OK.
- [ ] Front Docker accessible.
- [ ] Login client OK.
- [ ] Login admin OK.
- [ ] Login salarie OK.
- [ ] Login pro OK.
- [ ] Admin validations contient au moins une donnee.
- [ ] Catalogue contient au moins une annonce.
- [ ] Conteneurs visibles.
- [ ] Traductions chargees.
- [ ] Captures ouvertes dans un dossier accessible.


---

# 09 - Timing repetition

Objectif final : 14 min 30, pas 15 min.

## Repartition cible

| Personne | Temps total cible | Parties |
|---|---:|---|
| Faty | 5 min 00 | Accueil, login, client, catalogue, conclusion |
| Amel | 5 min 00 | Validation, API Go, logistique, base SQL |
| Ndoya | 4 min 30 | Evenements, salarie, Docker/Nginx |

## Repetition 1 - Lecture lente

But : comprendre le script.

- Lire sans cliquer.
- Chronometrer chaque personne.
- Couper les phrases trop longues.
- Objectif : moins de 17 min.

## Repetition 2 - Clics reels

But : caler les onglets.

- Ouvrir les pages dans l'ordre du fichier `02_DEROULE_DEMO_CLIC_PAR_CLIC.md`.
- Verifier que chaque page charge sans erreur rouge bloquante.
- Noter les temps morts.
- Objectif : moins de 15 min 30.

## Repetition 3 - Version soutenance

But : tenir le timing.

- Une seule personne controle la souris ou bien chaque personne controle sa partie, mais il faut choisir avant.
- Interdiction d'improviser une fonctionnalite externe.
- Si une action prend plus de 15 secondes, passer au plan B.
- Objectif : 14 min 30.

## Signaux de transition

- Faty vers Amel : "Une fois l'annonce soumise, elle passe par le back-office."
- Amel vers Faty : "Apres validation, l'annonce est visible cote catalogue."
- Faty vers Ndoya : "La plateforme ne s'arrete pas au particulier : elle sert aussi les professionnels et l'equipe interne."
- Ndoya vers Amel : "Je termine sur le deploiement, puis Amel montre la base SQL."
- Amel vers conclusion : "On revient sur la valeur globale livree."

## Regles de parole

- Ne pas dire "mon code" : dire "notre API", "notre front", "notre workflow".
- Ne pas dire "normalement" : dire "pour securiser la demo".
- Ne pas dire "c'est juste une page" : dire "cet ecran permet a l'utilisateur de...".
- Ne pas s'excuser si une integration externe n'est pas lancee : expliquer que le risque est maitrise.

## Checklist juste avant passage

- [ ] Navigateur ouvert sur accueil Docker.
- [ ] Zoom navigateur a 90 ou 100 %, pas devtools visible.
- [ ] Comptes de test notes.
- [ ] Captures de secours ouvertes dans un dossier.
- [ ] Terminal pret avec `docker compose ps`.
- [ ] IDE ouvert sur `API/route/annonces.go`.
- [ ] phpMyAdmin ou `db/init.sql` pret.
- [ ] Chacune connait sa premiere phrase.
- [ ] Chacune connait sa transition.

## Chrono de secours

Si a 07:30 vous n'avez pas fini admin validation, couper :
- forum ;
- paiement Stripe live ;
- PDF live ;
- creation evenement live.

Si a 12:00 vous n'avez pas commence Docker, passer directement a :
- `docker compose ps`;
- `nginx.conf`;
- conclusion.

---

# Generation PDF

PDF non genere automatiquement dans cet environnement : `pandoc` n'est pas installe et Python ne dispose pas du module `reportlab`.

Commande recommandee si Pandoc est installe :

```powershell
pandoc SOUTENANCE_DEV_SCRIPT/KIT_PRESENTATION_DEV_COMPLET.md -o SOUTENANCE_DEV_SCRIPT/KIT_PRESENTATION_DEV_COMPLET.pdf
```

Alternative Python apres installation de reportlab :

```powershell
python -m pip install reportlab
```

Puis convertir le Markdown via un outil Markdown PDF de VS Code, Typora, Obsidian ou Pandoc.

