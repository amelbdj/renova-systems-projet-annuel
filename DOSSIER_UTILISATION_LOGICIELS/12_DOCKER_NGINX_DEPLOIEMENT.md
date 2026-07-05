# 12 — Docker, Nginx & déploiement

## Rôle de Docker
Docker **empaquette chaque brique** du projet (base, API, front) dans un conteneur isolé et reproductible. Avantage : le projet tourne **à l'identique** sur n'importe quelle machine disposant de Docker, sans installer Go, MySQL ou un serveur web.

## Rôle de Docker Compose
`docker-compose.yml` **orchestre** tous les conteneurs d'une seule commande : il définit les services, le réseau, les volumes (persistance) et l'ordre de démarrage.

- **`docker-compose.yml`** : stack **locale** (build des images depuis `API/` et `Frontend/`). Inclut aussi **phpMyAdmin** (interface web de la base, http://localhost:8082) pour le développement.
- **`docker-compose-prod.yml`** : stack **prod/VM** utilisant les images publiées sur Docker Hub (`amelbdj/upcycle-backend`, `amelbdj/upcycle-frontend`).

### Services (local)
| Service | Conteneur | Image / build | Port | Rôle |
|---|---|---|---|---|
| `mysql` | `uc_mysql` | `mysql:8.0` | interne | Base de données `pa2026` |
| `backend` | `uc_backend` | build `./API` | 8081 | API Go |
| `frontend` | `uc_frontend` | build `./Frontend` | `WEB_PORT`→80 | Nginx (front + proxy) |
| `phpmyadmin` | `uc_phpmyadmin` | `phpmyadmin:latest` | 8082 | Admin base (local) |

### Volumes (persistance)
- `mysql_data` → données MySQL.
- `uploads_data` → fichiers uploadés (images, PDF).
- `documents_data` → factures/contrats PDF.

## Rôle de Nginx
Nginx (conteneur `frontend`, config `Frontend/nginx.conf`) :
- **sert les fichiers statiques** du front (HTML/CSS/JS) ;
- **fait office de reverse-proxy** : `/api/*` → API Go (`backend:8081`), `/uploads/*` → fichiers du back-end ;
- **réécrit les URLs propres** (`/login`, `/admin`, `/annonces`…) vers les fichiers `.html` correspondants.

En **production**, un second Nginx (sur la VM Ubuntu) ajoute le **HTTPS** et le domaine `upcycleconnect.pro` devant la stack Docker.

## Lancer les conteneurs
```bash
# Local
WEB_PORT=8088 docker compose up -d --build

# Prod (sur la VM)
docker compose -f docker-compose-prod.yml pull
docker compose -f docker-compose-prod.yml up -d
```

## Voir les logs
```bash
docker compose logs -f backend      # API en direct
docker compose logs -f frontend     # Nginx
docker compose logs --tail=100 mysql
docker ps                           # état + ports
```

## Redémarrer un service
```bash
docker compose up -d --build backend    # rebuild + redémarre l'API seule
docker compose up -d --build frontend   # rebuild + redémarre le front seul
docker compose restart backend          # redémarre sans rebuild
```

## Prouver que le site est accessible depuis l'extérieur
- Ouvrir **https://upcycleconnect.pro/** depuis n'importe quel appareil (hors réseau local).
- Le certificat **HTTPS** est valide (cadenas dans le navigateur).
- L'en-tête serveur indique `nginx (Ubuntu)` sur une **IP publique** (VM), preuve que ce n'est pas du localhost.
- Vérification rapide : `curl -I https://upcycleconnect.pro/` → réponse `200`/`HTTP 2`.

## Schéma du flux
```
[Utilisateur Internet]
        │  HTTPS (https://upcycleconnect.pro)
        ▼
[Nginx hôte (VM Ubuntu)]  ── TLS / domaine
        │
        ▼
[Conteneur frontend : Nginx]
        ├── fichiers statiques (Front)
        ├── /api/*     → [Conteneur backend : API Go :8081]
        └── /uploads/* → fichiers (images, PDF)
                                   │
                                   ▼
                        [Conteneur mysql : base pa2026]
```

## Mettre à jour la production
```bash
# 1) Construire et publier les images
docker build -t amelbdj/upcycle-backend:vNN ./API   && docker push amelbdj/upcycle-backend:vNN
docker build -t amelbdj/upcycle-frontend:vNN ./Frontend && docker push amelbdj/upcycle-frontend:vNN
# 2) Sur la VM : mettre les tags à jour dans docker-compose-prod.yml, puis
docker compose pull && docker compose up -d
```
