# 13 — Dépannage rapide (FAQ)

> Réflexe n°1 : `docker ps` (tout tourne ?) puis `docker compose logs -f backend`.

| Problème | Symptôme | Cause probable | Commande / fichier | Correction |
|---|---|---|---|---|
| **Front ne démarre pas** | `uc_frontend` s'arrête | Nginx ne trouve pas `backend` en upstream | `docker compose logs frontend` ; `Frontend/nginx.conf` | Lancer la stack **complète** (`docker compose up -d`), pas le front seul |
| **API ne démarre pas** | `uc_backend` redémarre en boucle | Erreur de compilation / MySQL pas prêt | `docker compose logs -f backend` ; `cd API && go build ./...` | Corriger l'erreur Go ; attendre `uc_mysql` **healthy** |
| **Base ne répond pas** | 500 sur toutes les routes | MySQL non prêt / mauvais identifiants | `docker compose logs mysql` ; `API/bdd/db.go` | Attendre `healthy` ; vérifier `DB_HOST=mysql` |
| **Erreur 401** | Appels rejetés | Token absent/expiré | Console : `localStorage.getItem('token')` ; `API/auth/jwt.go` | Se reconnecter |
| **Erreur 403** | « Accès refusé » | Mauvais rôle | `localStorage.getItem('userRole')` ; `VerifyRoleMiddleware` | Utiliser un compte du bon rôle |
| **Erreur 500** | Réponse 500 | Erreur SQL / scan / nil | `docker compose logs -f backend` ; `API/bdd/*Req.go` | Lire la ligne d'erreur Go (colonne/table) |
| **Erreur CORS** | *blocked by CORS policy* | Handler sans `OPTIONS`/header Authorization | `API/admin/<domaine>.go` | Ajouter `Access-Control-Allow-Headers: Content-Type, Authorization` + route `OPTIONS`. **N'arrive pas en prod** (même origine via `/api`) |
| **Docker ne build pas** | `docker build` échoue | Erreur Go / fichier manquant | `cd API && go build ./...` ; `API/Dockerfile` | Corriger l'erreur affichée avant de rebuild |
| **Variable d'env manquante** | Stripe/DB par défaut | `.env` absent | `docker exec uc_backend env \| grep DB_` ; `.env` | Copier `.env.example` → `.env` et remplir |
| **Import SQL échoue / accents cassés** | erreurs SQL ou « ? » à la place des accents | Import sans utf8mb4 | — | Réimporter avec `--default-character-set=utf8mb4` |
| **Endpoint introuvable (404)** | 404 sur une route | Route non déclarée | `grep -r "chemin" API/route/` ; `API/main.go` | Vérifier `HandleFunc` + appel `RoutesX()` dans `main.go` |
| **Site externe ne répond pas** | `upcycleconnect.pro` inaccessible | VM/conteneur arrêté, Nginx hôte KO | Sur la VM : `docker ps` ; `sudo systemctl status nginx` | Relancer la stack (`docker compose up -d`) ; `sudo systemctl restart nginx` |
| **Nginx ne redirige pas** | 404 sur `/api/...` ou URLs propres | Config Nginx non à jour | `Frontend/nginx.conf` ; `sudo nginx -t` | Vérifier les blocs `location /api/`, `/uploads/`, URLs propres ; recharger Nginx |
| **413 Request Too Large (upload)** | upload photo refusé | Limite Nginx **hôte** (1 Mo) | Nginx hôte de la VM | Ajouter `client_max_body_size 30M;` dans la conf Nginx **hôte** |
| **Modif invisible** | Rien ne change à l'écran | Cache navigateur / conteneur pas rebuild | — | **Ctrl+Shift+R** ; sinon `docker compose up -d --build frontend` |

## Méthode générale « ça ne marche pas » (4 étapes)
1. `docker ps` → les conteneurs sont-ils **Up** (mysql **healthy**) ?
2. `docker compose logs -f backend` → lire la vraie erreur.
3. Console navigateur (F12) → onglet **Réseau** : quel code HTTP ?
4. Reproduire l'appel en `curl` pour isoler front vs back.
