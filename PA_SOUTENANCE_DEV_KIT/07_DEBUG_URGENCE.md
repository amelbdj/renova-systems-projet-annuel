# 07 — Debug d'urgence (soutenance)

> Réflexe n°1 : `docker ps` (tout tourne ?) puis `docker compose logs -f backend`.

| # | Problème | Symptôme | Cause probable | Fichier à vérifier | Commande | Correction rapide |
|---|----------|----------|----------------|--------------------|----------|-------------------|
| 1 | **Front ne démarre pas** | Conteneur `uc_frontend` exit | Erreur Nginx / `backend` introuvable en upstream | `Frontend/nginx.conf` | `docker compose logs frontend` | Lancer la stack **complète** (`docker compose up -d`) — le front seul échoue car `proxy_pass backend` |
| 2 | **API ne démarre pas** | `uc_backend` redémarre en boucle | Erreur de compilation / port pris / DB absente | logs backend | `docker compose logs -f backend` | Corriger l'erreur Go ; vérifier `mysql` healthy avant |
| 3 | **DB ne répond pas** | 500 sur toutes les routes | MySQL pas prêt / mauvais identifiants | `API/bdd/db.go`, `docker-compose.yml` (DB_*) | `docker compose logs mysql` ; `docker exec uc_mysql mysqladmin ping -uroot -p...` | Attendre `healthy` ; vérifier `DB_HOST=mysql` |
| 4 | **Erreur CORS** | Console: *blocked by CORS policy* | Handler sans `OPTIONS` ou sans header `Authorization` autorisé | `API/admin/<domaine>.go` | — | Ajouter `Access-Control-Allow-Headers: Content-Type, Authorization` + route `OPTIONS`. **N'arrive pas en prod** (même origine via `/api`) |
| 5 | **401 Unauthorized** | Appels rejetés | Token absent/expiré | `API/auth/jwt.go`, `localStorage` | Console : `localStorage.getItem('token')` | Se reconnecter (le token est régénéré au login) |
| 6 | **403 Forbidden** | *Accès refusé…* | Mauvais rôle | `API/auth/jwt.go` (`VerifyRoleMiddleware`) | `localStorage.getItem('userRole')` | Utiliser un compte du bon rôle |
| 7 | **500 Internal** | Réponse 500 | Erreur SQL / scan / nil | logs backend + `API/bdd/*Req.go` | `docker compose logs -f backend` | Lire la ligne d'erreur Go (nom de colonne, table…) |
| 8 | **Docker ne build pas** | `docker build` échoue | Erreur Go / fichier manquant | `API/Dockerfile`, `go.mod` | `cd API && go build ./...` | Corriger l'erreur Go affichée avant de rebuild |
| 9 | **Variable d'env manquante** | Stripe/OneSignal KO, DB par défaut | `.env` absent | `.env`, `docker-compose.yml` | `docker exec uc_backend env \| grep DB_` | Copier `.env.example` → `.env` et remplir |
| 10 | **Endpoint introuvable (404)** | 404 sur une route | Route non déclarée / faute de frappe | `API/route/*.go` + `API/main.go` | `grep -r "chemin" API/route/` | Vérifier la déclaration `HandleFunc` et que `RoutesX()` est appelé dans `main.go` |
| 11 | **Modif faite mais invisible** | Rien ne change à l'écran | **Cache navigateur** ou conteneur pas rebuild | — | — | **Ctrl+Shift+R** ; sinon `docker compose up -d --build frontend` |
| 12 | **Accents cassés (« ? »)** | Textes avec `?` | Import DB en latin1 | `db/init.sql`, table `translations` | — | Réimporter avec `--default-character-set=utf8mb4` |
| 13 | **413 Request Too Large** | Upload image échoue | Limite Nginx (hôte VM) | `Frontend/nginx.conf` (conteneur OK 30M) / Nginx hôte | — | Ajouter `client_max_body_size 30M;` dans le Nginx **hôte** de la VM |
| 14 | **Paiement « impossible de contacter le serveur »** | Erreur au checkout event | Salarié sans `stripe_account_id` | `API/admin/stripe.go`, table `utilisateur` | `SELECT stripe_account_id FROM utilisateur WHERE id=..` | Renseigner un `stripe_account_id` de test |

## Le réflexe « ça marche pas » en 4 étapes
1. `docker ps` → les 3 conteneurs sont-ils **Up** (et mysql **healthy**) ?
2. `docker compose logs -f backend` → lire la vraie erreur.
3. Console navigateur (F12) → onglet **Réseau** : quel code HTTP ? (401/403/500/CORS/404)
4. Reproduire l'appel en `curl` pour isoler front vs back.
