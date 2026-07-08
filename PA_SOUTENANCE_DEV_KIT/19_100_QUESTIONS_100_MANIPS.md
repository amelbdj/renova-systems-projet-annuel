# 19 — 100 questions théoriques + 100 manips de code

> Banque de révision massive pour la soutenance. Réponses courtes volontairement.
> Colonne de gauche = ce que le jury peut demander ; à droite = la réponse / la solution.

---

# PARTIE 1 — 100 QUESTIONS THÉORIQUES

## Architecture & général (1-12)

1. **C'est quoi l'architecture globale du projet ?** → Front statique (HTML/CSS/JS) servi par Nginx + une API Go + une base MySQL, le tout dans Docker.
2. **Pourquoi séparer frontend et backend ?** → Chacun évolue indépendamment ; le back expose une API, le front la consomme.
3. **Nginx sert à quoi chez vous ?** → Servir les fichiers statiques ET faire proxy `/api/` vers le backend Go.
4. **Pourquoi Docker ?** → Même environnement partout (dev, VM), déploiement reproductible.
5. **C'est quoi une API REST ?** → Une interface HTTP où chaque URL = une ressource, et les verbes (GET/POST/PUT/DELETE) = les actions.
6. **Le front et le back sont sur le même serveur ?** → Oui en prod (même origine), Nginx route `/api/` vers Go.
7. **Combien de rôles utilisateurs ?** → 4 : Utilisateur, Pro, Salarié, Administrateur.
8. **Où est la logique métier ?** → Dans le backend Go (les handlers + les requêtes bdd).
9. **Le front peut-il accéder directement à MySQL ?** → Non, jamais. Il passe toujours par l'API.
10. **Qu'est-ce qui tourne sur quel port ?** → Backend Go : 8081 ; front (Nginx) : 80 ; MySQL : 3306.
11. **C'est quoi le langage du backend ?** → Go (Golang).
12. **Pourquoi Go plutôt que PHP/Node ?** → Compilé, rapide, typé, un seul binaire à déployer.

## Go / backend (13-30)

13. **C'est quoi un handler en Go ?** → Une fonction `func(w http.ResponseWriter, r *http.Request)` qui traite une requête.
14. **Comment on déclare une route ?** → `http.HandleFunc("GET /chemin", handler)`.
15. **C'est quoi un middleware ?** → Une fonction qui enveloppe un handler pour faire un contrôle avant (ex. vérifier le token).
16. **Comment Go renvoie du JSON ?** → `json.NewEncoder(w).Encode(donnees)`.
17. **Comment lire le corps JSON d'une requête ?** → `json.NewDecoder(r.Body).Decode(&struct)`.
18. **C'est quoi une struct ?** → Un type qui regroupe des champs (comme un objet).
19. **À quoi servent les tags `json:"..."` ?** → Dire quel nom de clé JSON correspond à quel champ Go.
20. **Comment récupérer un paramètre d'URL ?** → `r.URL.Query().Get("id")` (query) ou `r.PathValue("id")` (chemin).
21. **Comment gérer les erreurs en Go ?** → On renvoie une `error` et on la teste : `if err != nil { ... }`.
22. **C'est quoi `defer` ?** → Reporte l'exécution d'une ligne à la fin de la fonction (ex. `defer rows.Close()`).
23. **Comment on fait une requête SQL ?** → `Db.Query(...)` (lecture), `Db.Exec(...)` (écriture), `Db.QueryRow(...)` (une ligne).
24. **C'est quoi le `?` dans les requêtes ?** → Un paramètre préparé (évite les injections SQL).
25. **Pourquoi utiliser des requêtes préparées ?** → Sécurité : empêche l'injection SQL.
26. **C'est quoi `rows.Scan()` ?** → Copie les colonnes d'une ligne SQL dans des variables Go.
27. **C'est quoi `COALESCE` dans vos SELECT ?** → Renvoie une valeur par défaut si la colonne est NULL.
28. **Comment compiler le backend ?** → `cd API && go build ./...`.
29. **C'est quoi un paramètre variadique ?** → `...string` : accepte 0, 1 ou plusieurs valeurs (ex. `rolesAutorises`).
30. **Comment on lit une variable d'environnement ?** → `os.Getenv("NOM")`.

## JavaScript / frontend (31-42)

31. **Comment on appelle l'API depuis le front ?** → `fetch(url, options)`.
32. **C'est quoi `async/await` ?** → Une écriture lisible du code asynchrone (attendre une réponse sans bloquer).
33. **C'est quoi une Promise ?** → Un objet représentant un résultat futur (réussi ou échoué).
34. **Comment on stocke le token côté front ?** → Dans `localStorage`.
35. **Différence localStorage / sessionStorage ?** → localStorage persiste après fermeture ; sessionStorage non.
36. **Comment modifier le DOM ?** → `document.getElementById(...)`, `.innerHTML`, `.textContent`, `createElement`.
37. **C'est quoi `data-i18n` ?** → Un attribut pour la traduction : le texte est remplacé selon la langue.
38. **Comment on gère plusieurs langues ?** → `translate.js` + un dictionnaire ; `appliquerTraductions()` remplace les textes.
39. **C'est quoi `addEventListener` ?** → Attache une fonction à un événement (clic, chargement…).
40. **`DOMContentLoaded` sert à quoi ?** → Lancer du code quand le HTML est prêt.
41. **Pourquoi `?v=2` sur un script ?** → Forcer le navigateur à recharger la nouvelle version (cache-busting).
42. **C'est quoi le CORS ?** → Une sécurité navigateur ; le backend autorise l'origine via des en-têtes `Access-Control-*`.

## Base de données / SQL (43-54)

43. **Quel SGBD ?** → MySQL 8.
44. **C'est quoi une clé primaire ?** → Identifiant unique d'une ligne (`id`).
45. **C'est quoi une clé étrangère ?** → Une colonne qui référence la clé primaire d'une autre table.
46. **C'est quoi `ON DELETE CASCADE` ?** → Supprime automatiquement les lignes enfants quand le parent est supprimé.
47. **Différence INNER JOIN / LEFT JOIN ?** → INNER = seulement les correspondances ; LEFT = tout à gauche même sans correspondance.
48. **C'est quoi un index ?** → Une structure qui accélère les recherches sur une colonne.
49. **C'est quoi une transaction ?** → Un groupe d'opérations tout-ou-rien (COMMIT / ROLLBACK).
50. **Comment compter des lignes ?** → `SELECT COUNT(*) FROM table`.
51. **C'est quoi `GROUP BY` ?** → Regrouper des lignes pour agréger (COUNT, SUM…).
52. **C'est quoi un ENUM ?** → Une colonne à valeurs limitées (ex. `role`, `statut_vente`).
53. **Comment éviter les injections SQL ?** → Requêtes préparées avec `?`, jamais de concaténation.
54. **Comment sauvegarder la base ?** → `mysqldump` ; restaurer avec `mysql < fichier.sql`.

## Sécurité / JWT / bcrypt (55-70)

55. **Comment stockez-vous les mots de passe ?** → Hachés avec bcrypt (coût 10), jamais en clair.
56. **Peut-on déchiffrer un hash bcrypt ?** → Non ; on compare seulement (`CompareHashAndPassword`).
57. **C'est quoi un JWT ?** → Un jeton signé contenant des infos (id, rôle, expiration).
58. **Le JWT est-il chiffré ?** → Non, il est **signé** (lisible mais infalsifiable sans la clé).
59. **Comment le serveur vérifie un token ?** → Il revérifie la signature avec sa clé secrète (`jwtKey`).
60. **Différence authentification / autorisation ?** → Auth = qui es-tu (token) ; autorisation = as-tu le droit (rôle).
61. **Différence 401 / 403 ?** → 401 = pas/mauvais token ; 403 = connecté mais pas les droits.
62. **Où est stocké le rôle ?** → Dans le token (signé), pas envoyé à la main par le front.
63. **Peut-on tricher en changeant son rôle ?** → Non : la signature ne correspond plus → token invalide → 401.
64. **Combien de temps vit un token ?** → 24 h (`ExpiresAt`).
65. **Que se passe-t-il à l'expiration ?** → `token.Valid` = faux → 401 → redirection login.
66. **C'est quoi le middleware `VerifyTokenMiddleware` ?** → Vérifie que l'utilisateur est connecté (token valide).
67. **C'est quoi `VerifyRoleMiddleware` ?** → Vérifie en plus que le rôle est autorisé.
68. **La sécurité front (auth_guard) suffit-elle ?** → Non, c'est du confort ; la vraie sécurité est backend.
69. **C'est quoi HS256 ?** → Algorithme de signature symétrique du JWT (une seule clé secrète).
70. **Comment protéger la clé JWT ?** → Variable d'environnement, jamais commitée.

## Stripe / paiement (71-80)

71. **Comment gérez-vous les paiements ?** → Via Stripe ; on ne manipule aucune donnée bancaire.
72. **C'est quoi Stripe Connect ?** → Chaque vendeur a un compte Stripe relié ; il reçoit l'argent directement.
73. **Comment prenez-vous une commission ?** → `ApplicationFeeAmount` (5 %) + `TransferData` vers le vendeur.
74. **C'est quoi une Checkout Session ?** → La page de paiement hébergée par Stripe.
75. **Pourquoi les montants sont ×100 ?** → Stripe travaille en centimes (25 € = 2500).
76. **C'est quoi un webhook ?** → Stripe appelle notre serveur pour confirmer un paiement.
77. **Pourquoi un webhook et pas la page de succès ?** → Le client peut fermer son navigateur ; le webhook arrive toujours.
78. **Comment vérifie-t-on que le webhook vient de Stripe ?** → Vérification de la signature avec `WebhookSecret`.
79. **Êtes-vous en vrai argent ?** → Non, mode test (`sk_test_`), carte `4242 4242 4242 4242`.
80. **Quels types de paiement gérez-vous ?** → Achat d'annonce, inscription événement, abonnement Pro récurrent.

## Docker / déploiement (81-90)

81. **C'est quoi une image Docker ?** → Un modèle figé de l'application (code + dépendances).
82. **C'est quoi un conteneur ?** → Une instance qui tourne à partir d'une image.
83. **C'est quoi docker-compose ?** → Un fichier qui décrit/lance plusieurs conteneurs ensemble.
84. **C'est quoi un Dockerfile ?** → La recette pour construire une image.
85. **Comment reconstruire un service ?** → `docker compose up -d --build <service>`.
86. **Comment voir les logs ?** → `docker compose logs -f <service>`.
87. **C'est quoi un volume ?** → Un stockage persistant (ex. les uploads, la base) hors du conteneur.
88. **Pourquoi un volume pour MySQL ?** → Garder les données même si le conteneur est recréé.
89. **Comment passer une config au conteneur ?** → Variables d'environnement (`environment:` dans compose).
90. **Comment déployez-vous une nouvelle version ?** → Build image → push → `pull` + `up -d` sur la VM.

## HTTP / API / divers (91-100)

91. **Différence GET / POST ?** → GET lit (paramètres dans l'URL) ; POST envoie des données (dans le corps).
92. **À quoi sert PUT ?** → Modifier une ressource existante.
93. **À quoi sert DELETE ?** → Supprimer une ressource.
94. **C'est quoi un code 200 / 404 / 500 ?** → 200 OK, 404 introuvable, 500 erreur serveur.
95. **C'est quoi une requête preflight OPTIONS ?** → Le navigateur demande l'autorisation avant un PUT/DELETE (CORS).
96. **Comment les notifications marchent ?** → Une ligne en base (cloche) + un push OneSignal, via `SendPushNotification`.
97. **Comment le forum s'actualise en temps réel ?** → Polling : `setInterval` recharge les messages toutes les 4 s.
98. **C'est quoi le Swagger ?** → Une doc interactive de l'API (`/swagger`, basée sur `openapi.yaml`).
99. **Comment testez-vous une route protégée ?** → `curl` avec l'en-tête `Authorization: Bearer <token>`.
100. **Quelle est la partie la plus complexe ?** → Le paiement Stripe (Connect + commission + webhook).

---

# PARTIE 2 — 100 MANIPS DE CODE

> Format : **objectif** → **fichier(s)** → **solution courte**. Fais-les sur ton environnement local.

## Frontend — HTML / CSS (1-20)

1. Changer la couleur d'accent des annonces → `style/annonceAll.css` → `:root{--ac:255,120,0;}`.
2. Ajouter un lien nav « À propos » → une page HTML → `<a class="nav-link" href="/a-propos">À propos</a>`.
3. Changer le titre d'un onglet → `<title>` de la page → nouveau texte.
4. Mettre le logo plus grand → CSS du `.nav-logo img` → `height:48px;`.
5. Changer le texte d'un bouton → le HTML du bouton → nouveau libellé.
6. Ajouter un footer → fin du `<body>` → `<footer>...</footer>`.
7. Arrondir les cartes annonces → `.listing-card` → `border-radius:16px;`.
8. Changer la police → `<link>` Google Fonts + `font-family` dans le CSS.
9. Masquer un élément → CSS → `display:none;`.
10. Ajouter une ombre aux cartes → `.card` → `box-shadow:0 4px 12px rgba(0,0,0,.3);`.
11. Rendre la nav collante → `.topnav` → `position:sticky; top:0;`.
12. Changer le nombre de colonnes de la grille → `.listings-grid` → `grid-template-columns:repeat(4,1fr);`.
13. Ajouter un `placeholder` à un champ → l'input → `placeholder="Votre texte"`.
14. Rendre un champ obligatoire (HTML) → l'input → `required`.
15. Changer la couleur du thème Pro → `style/client.css` `.theme-pro` → modifier `--blue`.
16. Ajouter une icône Font Awesome → `<i class="fas fa-star"></i>`.
17. Centrer un bloc → CSS → `margin:0 auto; max-width:...`.
18. Changer le favicon → `<link rel="icon" href="...">`.
19. Ajouter un espace entre sections → CSS → `padding`/`margin`.
20. Passer un texte en gras → CSS → `font-weight:700;`.

## Frontend — JavaScript (21-45)

21. Changer le nombre d'annonces par page → `annonceAll.js` → `ANN_PAR_PAGE = 12`.
22. Changer le nombre d'users par page → `admin/user.js` → `USERS_PAR_PAGE = 20`.
23. Logguer le nombre de pages → `renderPageAnnonces` → `console.log(nbPages)`.
24. Afficher une alerte au chargement → `document.addEventListener("DOMContentLoaded",()=>alert("ok"))`.
25. Trier les annonces par prix croissant → appeler `sortListings('price-asc')`.
26. Filtrer les annonces gratuites → filtrer `allAnnonces` sur `prix<=0`.
27. Ajouter un bouton « retour en haut » → `window.scrollTo({top:0})`.
28. Afficher le nombre de résultats → `resultCount.textContent = items.length`.
29. Changer l'intervalle de polling forum → `forum.js` → `setInterval(...,2000)`.
30. Rediriger si non connecté → `if(!localStorage.getItem("token")) location.href="/login"`.
31. Récupérer l'id utilisateur → `localStorage.getItem("userId")`.
32. Ajouter un header Authorization à un fetch → `headers:{Authorization:"Bearer "+token}`.
33. Afficher une erreur si le fetch échoue → `.catch(err=>alert("Erreur"))`.
34. Formater un prix avec « € » → `` `${prix} €` ``.
35. Vider un conteneur → `element.innerHTML = ""`.
36. Créer une carte dynamiquement → `document.createElement("div")` + `appendChild`.
37. Ajouter un écouteur sur un bouton → `btn.addEventListener("click",fn)`.
38. Changer la langue par défaut → `translate.js` (langue initiale).
39. Désactiver un bouton → `btn.disabled = true`.
40. Faire défiler vers un élément → `el.scrollIntoView({behavior:"smooth"})`.
41. Empêcher le rechargement d'un form → `e.preventDefault()`.
42. Lire la valeur d'un input → `document.getElementById("x").value`.
43. Ajouter une classe conditionnellement → `el.classList.toggle("on", condition)`.
44. Recharger la page après une action → `window.location.reload()`.
45. Construire un FormData → `const fd=new FormData(); fd.append("titre",...)`.

## Backend Go (46-70)

46. Créer un endpoint `GET /admin/users/count` → route + handler + requête `COUNT(*)`.
47. Passer la commission de 5 % à 10 % → `admin/stripe.go` → `*10)/100` et `0.10`.
48. Réserver une route aux Admins → `route/...` → `VerifyRoleMiddleware(handler,"Administrateur")`.
49. Autoriser Admin OU Salarié → `VerifyRoleMiddleware(handler,"Administrateur","Salarié")`.
50. Ajouter un champ `telephone` renvoyé → `models/users.go` + SELECT + `Scan`.
51. Refuser une annonce sans titre → `if r.FormValue("titre")==""{http.Error(...,400);return}`.
52. Changer la durée du token (12 h) → `jwt.go` → `Add(12 * time.Hour)`.
53. Ajouter un log dans un handler → `fmt.Println("appel handler X")`.
54. Renvoyer un JSON custom → `json.NewEncoder(w).Encode(map[string]string{"ok":"1"})`.
55. Lire un paramètre d'URL → `r.URL.Query().Get("id")`.
56. Lire un param de chemin → `r.PathValue("id")`.
57. Ajouter un endpoint `GET /admin/stats` → route + handler renvoyant `{users,annonces}`.
58. Créer une nouvelle table via une requête → `Db.Exec("CREATE TABLE ...")`.
59. Ajouter une validation de rôle inconnue → renvoyer 400 si rôle non prévu.
60. Trier une requête SQL → ajouter `ORDER BY date_creation DESC`.
61. Limiter une requête → ajouter `LIMIT 10`.
62. Filtrer par statut → `WHERE statut_vente = ?`.
63. Ajouter un `COALESCE` sur un champ nullable → `COALESCE(ville,'')`.
64. Gérer la méthode OPTIONS → `if r.Method=="OPTIONS"{w.WriteHeader(200);return}`.
65. Renvoyer 404 si introuvable → `http.Error(w,"introuvable",http.StatusNotFound)`.
66. Hacher un mot de passe → `auth.HashPassword(mdp)`.
67. Comparer un mot de passe → `bcrypt.CompareHashAndPassword(hash,saisi)`.
68. Ajouter un champ à une struct → nouveau champ + tag `json`.
69. Lire une variable d'env avec défaut → helper `envOr("NOM","defaut")`.
70. Ajouter une route DELETE → `http.HandleFunc("DELETE /admin/x/{id}",handler)`.

## SQL / Base de données (71-85)

71. Ajouter une colonne → `ALTER TABLE utilisateur ADD COLUMN telephone VARCHAR(20);`.
72. Supprimer une colonne → `ALTER TABLE x DROP COLUMN y;`.
73. Compter les annonces par catégorie → `GROUP BY c.id` avec `COUNT(a.id)`.
74. Valider un utilisateur → `UPDATE utilisateur SET validation='Validé' WHERE id=42;`.
75. Compter les annonces sponsorisées → `SELECT COUNT(*) FROM annonce WHERE is_sponsored=1;`.
76. Trouver les 5 dernières annonces → `ORDER BY id DESC LIMIT 5;`.
77. Supprimer les annonces d'un user → `DELETE FROM annonce WHERE id_user=?;`.
78. Masquer un message forum → `UPDATE message_forum SET est_modere=1 WHERE id_message=?;`.
79. Lister les users par rôle → `SELECT * FROM utilisateur WHERE role='Pro';`.
80. Créer un index → `CREATE INDEX idx_ville ON annonce(ville);`.
81. Compter les users → `SELECT COUNT(*) FROM utilisateur;`.
82. Mettre à jour un prix → `UPDATE annonce SET prix=10 WHERE id=?;`.
83. Sauvegarder la base → `mysqldump -u.. -p.. pa2026 > b.sql`.
84. Restaurer la base → `mysql -u.. -p.. pa2026 < b.sql`.
85. Voir la structure d'une table → `DESCRIBE annonce;`.

## Docker / debug / outils (86-100)

86. Reconstruire le backend → `docker compose up -d --build backend`.
87. Voir les logs backend → `docker compose logs -f backend`.
88. Entrer dans le conteneur MySQL → `docker exec -it uc_mysql mysql -u.. -p..`.
89. Lister les conteneurs → `docker ps`.
90. Redémarrer un service → `docker compose restart backend`.
91. Voir les variables d'env d'un conteneur → `docker exec uc_backend env`.
92. Vérifier qu'un port écoute → `netstat -ano | grep 8081`.
93. Arrêter tous les conteneurs → `docker compose down`.
94. Reconstruire le front → `docker build -t upcycle-frontend ./Frontend`.
95. Compiler le back sans lancer → `cd API && go build ./...`.
96. Tester une route au curl → `curl http://localhost:8081/... -H "Authorization: Bearer $TOKEN"`.
97. Provoquer un 500 et lire l'erreur → mauvais param + `docker compose logs backend`.
98. Vider le cache navigateur → Ctrl+Shift+R.
99. Ouvrir la console front → F12 → onglet Console.
100. Voir les requêtes réseau → F12 → onglet Réseau (Network).

---

## Comment s'entraîner
- **J-3** : lis les 100 questions, coche celles que tu ne sais pas.
- **J-2** : refais les manips 21-24, 46-52, 71-74 (les plus demandées).
- **J-1** : relis seulement tes cases non cochées + les fiches 15/16/18.
- **Jour J** : garde ce document ouvert, il sert d'aide-mémoire.
