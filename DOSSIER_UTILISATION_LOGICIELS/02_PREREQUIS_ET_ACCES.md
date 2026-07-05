# 02 — Prérequis et accès

## Accès simple (utilisateur / correcteur qui veut juste tester en ligne)
- **Aucune installation.** Un navigateur récent suffit.
- **Navigateur recommandé** : Google Chrome ou Microsoft Edge à jour (Firefox fonctionne également).
- **URL publique** : **https://upcycleconnect.pro/**
- Se connecter avec un **compte de démonstration** (voir `03_COMPTES_DE_DEMONSTRATION.md`).

### Accéder au site
1. Ouvrir le navigateur.
2. Saisir l'adresse **https://upcycleconnect.pro/** → la page d'accueil s'affiche.
3. Cliquer sur **Connexion** (URL `/login`) et saisir un compte de démonstration.
4. Selon le rôle, l'utilisateur est redirigé vers son espace (`/client`, `/pro`, `/salarie`, `/admin`).

## Prérequis développeur / correcteur (pour lancer en local)
Le projet est **entièrement conteneurisé** : le seul prérequis obligatoire est **Docker**.

| Logiciel | Nécessaire ? | Usage |
|---|---|---|
| **Docker Desktop** + **Docker Compose** | ✅ Obligatoire | Lance toute la stack (base, API, front) |
| **Git** | Recommandé | Récupérer le dépôt (`git clone`) |
| **Go** (1.25) | Optionnel | Uniquement pour lancer l'API **sans** Docker |
| **MySQL 8** | Optionnel | Uniquement en mode « sans Docker » (sinon fourni par le conteneur) |
| **Node / npm** | ❌ Non requis | Le front est du HTML/JS statique, **pas de build** |
| **PostgreSQL** | ❌ Non | Le projet utilise **MySQL**, pas PostgreSQL |

> Résumé : avec **Docker Desktop** seul, tout fonctionne. Go/MySQL ne servent que pour un lancement manuel avancé.

## URL du projet
| Élément | URL |
|---|---|
| Site public (prod) | **https://upcycleconnect.pro/** |
| API en prod | `https://upcycleconnect.pro/api/...` (proxy Nginx → back-end) |
| Front en local (Docker) | `http://localhost:8088/` (port configurable via `WEB_PORT`) |
| API en local (Docker) | `http://localhost:8081/...` (exposée pour debug) |

## Détail de la communication front ↔ API
Le fichier `Frontend/script/config.js` choisit automatiquement l'adresse de l'API :
- **Développement WAMP local** (URL contenant `/Frontend/`) → `http://localhost:8081`.
- **Docker / prod** → `origin + "/api"` (ex : `https://upcycleconnect.pro/api`), le proxy Nginx retirant le préfixe `/api`.
