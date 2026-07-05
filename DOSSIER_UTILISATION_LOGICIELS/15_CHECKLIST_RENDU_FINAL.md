# 15 — Checklist de rendu final (MyGES)

## Livrables obligatoires
- [ ] **Code source complet** zippé (dossiers `API/`, `Frontend/`, `db/`, fichiers Docker).
- [ ] **Export SQL base vide** → `db/export_db_vide.sql` ✅ (généré).
- [ ] **Export SQL base remplie** → `db/export_db_remplie.sql` ✅ (généré).
- [ ] **Dossier d'utilisation** (ce dossier) en **PDF** → `DOSSIER_COMPLET_UTILISATION_LOGICIELS.pdf`.
- [ ] **URL publique** fonctionnelle → **https://upcycleconnect.pro/**.
- [ ] **Comptes de test** documentés (avec mots de passe **complétés**) → `03_COMPTES_DE_DEMONSTRATION.md`.
- [ ] **README** à la racine du projet.
- [ ] **docker-compose.yml** (+ `docker-compose-prod.yml`).
- [ ] **`.env.example`** présent.
- [ ] **Documentation API** → `09_GUIDE_API_GO.md`.
- [ ] **Preuve de lancement local** (capture `docker ps` / site sur `localhost:8088`).
- [ ] **Preuve de déploiement externe** (capture `https://upcycleconnect.pro` + cadenas HTTPS).

## Vérifications techniques avant dépôt
- [ ] `docker compose up -d --build` fonctionne **sur une machine propre** (testé chez un camarade).
- [ ] La base s'importe automatiquement (ou via `export_db_remplie.sql`).
- [ ] Les 4 rôles se connectent et accèdent à leur espace.
- [ ] L'API répond (`/admin/login` → token ; route protégée → 401 sans token).
- [ ] Les uploads (photos d'annonces) fonctionnent en ligne (limite Nginx hôte à 30 Mo).
- [ ] Le site public est **accessible depuis l'extérieur** (réseau mobile, pas seulement le Wi-Fi local).

## Points à finaliser (rappel du diagnostic)
- [ ] **Compléter les mots de passe** des comptes de test.
- [ ] **Ajouter les captures d'écran** (voir `14_CAPTURE_ECRANS_A_AJOUTER.md`).
- [ ] Mentionner honnêtement les éléments **absents/partiels** : Swagger (absent), tests automatisés (absents), vérification INSEE du SIRET (absente — seule la clé de Luhn est vérifiée), stockage `date_naissance` (absent), calcul du score par matériau (partiel).
- [ ] Externaliser les **secrets** (Stripe/OneSignal/JWT) via `.env` pour la version rendue.

## Sécurité / cohérence
- [ ] Vérifier qu'aucun secret sensible réel ne traîne en clair dans le dépôt public.
- [ ] Vérifier que le fichier `.env` **réel** n'est pas commité (seul `.env.example`).
