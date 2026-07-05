# 10 — Installation et lancement en local

> Méthode recommandée : **Docker Compose** (aucune installation de Go/MySQL requise). Une méthode manuelle est fournie en fin de fiche.

## A. Lancement avec Docker Compose (recommandé)

### 1. Récupérer le projet
```bash
git clone <URL_DU_DEPOT> upcycleconnect
cd upcycleconnect
```
*(ou copier le dossier du projet — celui qui contient `docker-compose.yml`, `API/`, `Frontend/`, `db/`).*

### 2. Configurer les variables d'environnement (optionnel)
Les valeurs par défaut fonctionnent sans `.env`. Pour personnaliser :
```bash
cp .env.example .env
# éditer .env : mots de passe, WEB_PORT, clés Stripe/OneSignal…
```
Contenu de `.env.example` :
```
WEB_PORT=80
MYSQL_ROOT_PASSWORD=change_me_root
DB_NAME=pa2026
DB_USER=upcycle
DB_PASS=change_me_app
STRIPE_SECRET_KEY=
STRIPE_WEBHOOK_SECRET=
ONESIGNAL_APP_ID=
ONESIGNAL_API_KEY=
```

### 3. Lancer toute la stack
```bash
WEB_PORT=8088 docker compose up -d --build
```
- Première exécution : build ~2–5 min (téléchargement + compilation Go).
- `WEB_PORT=8088` évite un conflit avec le port 80.
- La **base est importée automatiquement** au premier démarrage via `db/init.sql` — rien à importer manuellement.

### 4. Vérifier que tout fonctionne
```bash
docker ps      # 3 conteneurs : uc_mysql (healthy), uc_backend, uc_frontend, (+ uc_phpmyadmin en local)
```
- **Front** : ouvrir `http://localhost:8088/` → la page d'accueil s'affiche.
- **API** : `curl -s -o /dev/null -w "%{http_code}" http://localhost:8081/admin/users` → `401` (API vivante, protégée).
- Se connecter avec un compte de démonstration.

### 5. Arrêter / relancer
```bash
docker compose stop     # arrêt (données conservées)
docker compose start    # relance
docker compose down      # supprime les conteneurs (volumes/données conservés)
docker compose down -v   # ⚠️ supprime AUSSI les données (réimport de init.sql au prochain up)
```

## B. Importer un export SQL manuellement (si besoin)
```bash
# Base remplie (schéma + données)
docker exec -i uc_mysql mysql --default-character-set=utf8mb4 -uupcycle -pupcyclePass123 pa2026 < db/export_db_remplie.sql

# Base vide (structure seule)
docker exec -i uc_mysql mysql --default-character-set=utf8mb4 -uupcycle -pupcyclePass123 pa2026 < db/export_db_vide.sql
```

## C. Lancement manuel (sans Docker) — avancé
Prérequis : **Go 1.25** + **MySQL 8** installés.
```bash
# 1) Base : créer la base pa2026 et importer un export
mysql -uroot -p -e "CREATE DATABASE pa2026 CHARACTER SET utf8mb4;"
mysql -uroot -p pa2026 < db/export_db_remplie.sql

# 2) API Go (variables d'env par défaut : localhost/root/root/pa2026)
cd API
go build ./...     # vérifie la compilation
go run .           # écoute sur :8081

# 3) Front : servir les fichiers statiques du dossier Frontend
#    (via WAMP, ou un serveur statique). Ouvrir la page login.html.
```
> En mode WAMP, `Frontend/script/config.js` détecte l'URL `/Frontend/` et pointe l'API sur `http://localhost:8081`.

## Récapitulatif express
```
1. Docker Desktop lancé
2. Dossier du projet ouvert dans un terminal
3. WEB_PORT=8088 docker compose up -d --build
4. Attendre ~3 min → http://localhost:8088
5. Se connecter avec un compte de démonstration
```
