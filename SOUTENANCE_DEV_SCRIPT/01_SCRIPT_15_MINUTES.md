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
