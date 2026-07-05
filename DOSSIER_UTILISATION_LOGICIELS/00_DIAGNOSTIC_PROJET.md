# 00 — Diagnostic du projet

> Analyse factuelle du dépôt `renova-systems-projet-annuel` (produit **UpcycleConnect**, prestataire **Renova Systems**).

## Structure du projet
```
renova-systems-projet-annuel/
├── API/                    ← Back-end Go (API REST, port 8081)
│   ├── main.go             ← point d'entrée (enregistre les routes + ListenAndServe :8081)
│   ├── route/              ← déclaration des routes HTTP (par domaine)
│   ├── admin/              ← handlers HTTP (logique des endpoints) + upload, pdf, stripe, webhook
│   ├── bdd/                ← accès base (une fonction = une requête SQL) + db.go (connexion)
│   ├── models/             ← structs Go (User, Annonce, Evenement…)
│   ├── auth/jwt.go         ← JWT + middlewares de rôle
│   └── Dockerfile          ← build multi-stage du back-end
├── Frontend/               ← Front-end statique (HTML/CSS/JS) servi par Nginx
│   ├── *.html              ← pages (login, espClient, espPro, admin_*, salarie/*, annonceAll…)
│   ├── script/             ← JS (config.js, login.js, dashClient.js, admin/*, salarie/*, notifCloche.js)
│   ├── style/              ← CSS
│   ├── assets/             ← logo + favicon (logoUpcycle.png / .ico)
│   ├── nginx.conf          ← config Nginx (proxy /api, /uploads, URLs propres)
│   └── Dockerfile          ← build du conteneur front (nginx)
├── db/
│   ├── init.sql            ← schéma + données (import auto au 1er démarrage Docker)
│   ├── export_db_vide.sql      ← EXPORT structure seule (généré)
│   ├── export_db_remplie.sql   ← EXPORT structure + données (généré)
│   └── fix_translations_pro.sql← correctif d'accents (utilitaire)
├── docker-compose.yml      ← stack LOCALE (build local)
├── docker-compose-prod.yml ← stack VM/PROD (images Docker Hub amelbdj/upcycle-*)
├── .env.example            ← modèle de variables d'environnement
└── pa2026.sql              ← dump SQL de référence (historique)
```

## Technologies détectées
| Élément | Technologie | Preuve |
|---|---|---|
| Back-end | **Go** (`net/http`, sans framework) | `API/main.go`, `API/go.mod` |
| Base de données | **MySQL 8** | `docker-compose.yml` (`mysql:8.0`), `API/bdd/db.go` |
| Front-end | **HTML / CSS / JS vanilla** + **Nginx** | `Frontend/`, `Frontend/nginx.conf` |
| Conteneurs | **Docker + Docker Compose** | `docker-compose.yml`, `docker-compose-prod.yml`, Dockerfiles |
| Authentification | **JWT (HS256)** | `API/auth/jwt.go` |
| Paiement | **Stripe** (Checkout + Connect) | `API/admin/stripe.go`, `webhook.go` |
| Notifications push | **OneSignal** | `API/admin/notifications.go` |
| PDF | **go-pdf/fpdf** (serveur) + **jsPDF** (client) | `API/admin/pdf.go`, `Frontend/script/admin/box.js` |

## Emplacements clés
| Élément | Chemin |
|---|---|
| Front-end | `Frontend/` |
| API Go | `API/` (entrée `API/main.go`) |
| Base de données (schéma + seed) | `db/init.sql` |
| Exports SQL | `db/export_db_vide.sql`, `db/export_db_remplie.sql` |
| Docker (local) | `docker-compose.yml` |
| Docker (prod/VM) | `docker-compose-prod.yml` |
| Nginx | `Frontend/nginx.conf` |
| Variables d'environnement | `.env.example` (à copier en `.env`) |
| Connexion DB (code) | `API/bdd/db.go` |
| Config API côté front | `Frontend/script/config.js` |

## Fonctionnalités réellement présentes
- Authentification (inscription, connexion, mot de passe oublié) — **JWT**.
- Gestion des utilisateurs et **rôles/permissions** (Utilisateur, Pro, Salarié, Administrateur).
- **Annonces** (dons/ventes) : création, upload photo, validation admin, filtres par catégorie.
- **Catégories** : liste + ajout + suppression (back-office).
- **Conteneurs / box** : réservation → dépôt (PIN) → retrait (code-barres), simulateur matériel.
- **Commandes / commission** (5 %) via Stripe ; suivi finances admin.
- **Upcycling Score** (impact écologique).
- **Événements / formations** (création salarié, inscription, paiement Stripe).
- **Articles / actualités**, **Forum**, **Messagerie (WebSocket)**.
- **Espace Pro** : abonnements (freemium/premium), projets avant/après, sponsoring d'annonces.
- **Back-office admin** : validations, utilisateurs (pagination), conteneurs, finances, documents.
- **Notifications** : en base (cloche) + push OneSignal.
- **Génération PDF** : factures/contrats (serveur), rapport logistique (client).
- **Multilingue** (FR/EN) : import de langue depuis l'admin.
- **Déploiement Docker** + Nginx reverse-proxy + HTTPS (URL publique).

## Fonctionnalités absentes ou à ne pas promettre
| Élément | État |
|---|---|
| Documentation **Swagger/OpenAPI** | **Non présent dans le code actuel** (doc API générée à partir des routes) |
| **Tests automatisés** (`*_test.go`) | **Non présent dans le code actuel** (tests manuels curl/navigateur) |
| Stockage de la **date de naissance** | **Non présent** (colonne absente ; la vérif d'âge 18+ est côté front) |
| Vérification **réelle** du SIRET (API INSEE) | **Non présent** — seule la **clé de Luhn** est validée (format) |
| Calcul du score par **matériau** | **Présent partiellement** — le coefficient retombe sur « autre » (le matériau n'est pas relu dans `CalculateAndAddScore`) |

## Points à compléter manuellement
- **Mots de passe des comptes de test** : hashés (bcrypt) en base → non extractibles, à renseigner (voir `03_COMPTES_DE_DEMONSTRATION.md`).
- **Captures d'écran** du dossier (voir `14_CAPTURE_ECRANS_A_AJOUTER.md`).
- Secrets (Stripe/OneSignal/JWT) : valeurs de repli présentes dans le code → à externaliser en `.env` pour la prod.
