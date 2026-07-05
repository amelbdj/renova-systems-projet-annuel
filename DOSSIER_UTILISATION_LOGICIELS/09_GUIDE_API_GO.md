# 09 — Guide de l'API (Go)

## Présentation
L'API est écrite en **Go** avec la bibliothèque standard `net/http` (pas de framework). Elle expose une API REST/JSON, protégée par **JWT** et vérification de **rôle côté serveur**.

- **Emplacement du code** : dossier **`API/`** (point d'entrée `API/main.go`).
- **Port d'écoute** : **8081**.
- **Organisation** : `route/` (déclare les URLs) → `admin/` (handlers) → `bdd/` (SQL) → `models/` (structs).
- **Swagger/OpenAPI** : **non présent dans le code actuel** — la documentation ci-dessous est générée à partir des routes réelles (`API/route/*.go`).

## Lancer l'API
```bash
# Via Docker (recommandé) — depuis la racine du projet
docker compose up -d --build backend

# Sans Docker (Go + MySQL installés localement)
cd API
go run .        # écoute sur :8081
```

## Authentification
- Se connecter via `POST /admin/login` → renvoie un **token JWT**.
- Ajouter ce token sur chaque appel protégé : header `Authorization: Bearer <token>`.
- Le middleware `API/auth/jwt.go` vérifie le token (`VerifyTokenMiddleware`) et le rôle (`VerifyRoleMiddleware`).

```bash
# Obtenir un token
curl -X POST http://localhost:8081/admin/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test.admin@renova.test","mot_de_passe":"MOT_DE_PASSE"}'
```

## Codes d'erreur possibles
| Code | Signification |
|---|---|
| 200 / 201 | Succès |
| 400 | Requête invalide (données manquantes/malformées) |
| 401 | Non authentifié (token absent/expiré/invalide) |
| 403 | Rôle non autorisé |
| 404 | Ressource/route introuvable |
| 500 | Erreur serveur (souvent SQL) |

## Principales routes (générées depuis le code)

### Authentification — `API/route/auth.go`, `users.go`
| Méthode | Route | Accès | Description |
|---|---|---|---|
| POST | `/admin/login` | public | Connexion → token JWT |
| POST | `/auth/inscription` | public | Inscription (Particulier/Pro ; Admin/Salarié refusés) |
| POST | `/auth/check-email` | public | Vérifier disponibilité email |
| POST | `/auth/forgot-password` | public | Demande de réinitialisation |
| POST | `/auth/reset-password` | public | Réinitialisation via token |

### Utilisateurs — `API/route/users.go`
| Méthode | Route | Accès |
|---|---|---|
| GET | `/admin/users` | Salarié/Admin |
| GET | `/admin/users/{id}` | token |
| GET | `/admin/users/search?search=` | token |
| GET | `/admin/users/role/{role}` | token |
| POST | `/admin/users/add` | token |
| PUT | `/admin/users/modify/{id}` | token |
| DELETE | `/admin/users/delete/{id}` | token |
| PUT | `/admin/users/validate|refuse|ban/{id}` | token |

### Annonces — `API/route/annonces.go`
| Méthode | Route | Accès |
|---|---|---|
| GET | `/api/annonces/all?id=` | public |
| GET | `/mes-annonces?id=` | token |
| GET | `/admin/annonces` | token |
| POST | `/admin/annonces/add` | token (multipart : image) |
| PUT | `/admin/annonces/modify/{id}` | token |
| PUT | `/admin/annonces/validate|refuse/{id}` | token |
| DELETE | `/admin/annonces/delete/{id}` | token |
| POST | `/api/payment-annonce` | token |

### Catégories — `API/route/categories.go`
| Méthode | Route |
|---|---|
| GET | `/admin/categories` |
| POST | `/admin/categories/add` |
| DELETE | `/admin/categories/delete/{id}` |

### Conteneurs / box — `API/route/logistique.go`
| Méthode | Route |
|---|---|
| GET | `/api/admin/conteneurs` |
| GET | `/api/admin/conteneur/{id}/boxes` |
| POST | `/api/admin/conteneur/create` |
| POST | `/api/admin/box/add` |
| PUT | `/api/admin/box/update` |
| POST | `/api/box/reserve` `/deposit` `/collect` |
| POST | `/api/hardware/simulate-deposit` `/simulate-withdrawal` |
| GET | `/api/user/boxes?user_id=` , `/api/user/pickups/{id}` |

### Événements — `API/route/evenements.go`
| Méthode | Route |
|---|---|
| GET | `/admin/evenements` |
| POST | `/admin/evenements/add` (multipart) |
| PUT | `/admin/evenements/validate|refuse/{id}` |
| POST | `/admin/evenements/inscription|desinscription` |
| POST | `/api/web/checkout/evenement` (Stripe) |

### Finance — `API/route/finance.go`
| Méthode | Route |
|---|---|
| GET | `/admin/finance/overview` |
| GET | `/admin/finance/transactions` |

### Pro / abonnements — `API/route/pro.go`
`POST /api/pro/subscribe|upgrade|cancel|portal` · `GET /api/pro/sync|invoices|projets|etapes` · `POST /api/pro/projets/create` · `POST /api/pro/annonces/sponsor`

### Forum / Chat — `API/route/forum.go`, `chat.go`
`GET/POST /user/forums` · `GET/POST /user/forums/messages` · `GET /ws/chat` (WebSocket) · `GET /api/chat/conversations|history`

### Notifications — `API/route/notifications.go`
`POST /admin/notifications/send` · `GET /admin/notifications/user/{id}` · `POST /admin/notifications/user/{id}/read`

### Traductions — `API/route/traductions.go`
`GET /api/translations?lang=fr` · `GET /api/languages` · `POST /admin/translations/add` (Admin) · `GET /admin/translations/keys`

### Documents / uploads — `API/route/documents.go`, `auth.go`
`GET /admin/documents` · `GET /uploads/*` (images, PDF servis par le back-end)

## Exemples curl
```bash
# Lister les utilisateurs (admin)
curl http://localhost:8081/admin/users -H "Authorization: Bearer $TOKEN"

# Lister les annonces publiques
curl "http://localhost:8081/api/annonces/all?id=1"

# Statistiques éco d'un utilisateur
curl "http://localhost:8081/api/user/stats?user_id=1"

# Créer une catégorie
curl -X POST http://localhost:8081/admin/categories/add \
  -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" \
  -d '{"libelle":"Verre"}'
```

> **Postman** : créer une variable `token` (récupérée via `/admin/login`) et l'utiliser dans l'en-tête `Authorization: Bearer {{token}}` pour tester les routes protégées.
