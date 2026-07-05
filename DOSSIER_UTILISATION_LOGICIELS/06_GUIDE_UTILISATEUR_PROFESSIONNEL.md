# 06 — Guide utilisateur : Professionnel / Artisan

Le **Professionnel** récupère de la matière première de réemploi et gère ses projets d'upcycling. Espace : **`/pro`** (interface en teal).

## 1. Connexion
1. Ouvrir `/login`, saisir un compte **Pro** (voir `03_COMPTES_DE_DEMONSTRATION.md`).
2. **Résultat** : redirection vers **`/pro`** (Espace Professionnels & Artisans).

> Un compte Pro peut être **« En attente »** de validation par un administrateur après inscription (une pièce justificative / SIRET est demandée).

## 2. Tableau de bord Pro
- **Page** : `/pro`.
- Indicateurs : **Projets actifs**, **CO₂ évité**.
- Encart **abonnement** (Freemium / Premium / Pro).

*(Capture : `dashboard_pro.png`.)*

## 3. Catalogue / consultation des annonces
- **Page** : `/annonces` (mêmes annonces que les particuliers, mais l'interface prend les **couleurs pro**).
- Filtres par catégorie/type, recherche, tri.
- **Réservation / achat** : le paiement passe par **Stripe** ; une **commission de 5 %** est prélevée par la plateforme.

## 4. Abonnements (Freemium / Premium)
- **Freemium** (gratuit) : dépôt de matériaux, accès aux annonces de base.
- **Premium / Plus / Pro** (payant) : tableaux de bord avancés, alertes prioritaires, visibilité accrue, factures PDF automatiques.
- **Actions** : « Passer au Premium », gestion via le **portail Stripe**, résiliation.
- Endpoints associés : `POST /api/pro/subscribe`, `/upgrade`, `/cancel`, `/portal`.

*(Capture : `abonnement_pro.png`.)*

## 5. Projets d'upcycling
- Depuis `/pro`, section **Projets** : créer un projet avec **photos avant / après** et **CO₂ évité estimé**.
- Suivi par **étapes** (à faire / en cours / terminé).
- Endpoints : `POST /api/pro/projets/create`, `GET /api/pro/projets`, `POST /api/pro/etapes/create`, `PUT /api/pro/etapes/statut`.

*(Capture : `projet_pro.png`.)*

## 6. Sponsoring d'annonces
- Un pro peut **sponsoriser** une de ses annonces pour la mettre en avant sur la marketplace.
- Endpoint : `POST /api/pro/annonces/sponsor`.

## 7. Factures / transactions
- Les paiements (abonnement, achats) génèrent des **factures/contrats PDF** côté serveur.
- Consultables via l'espace pro (`GET /api/pro/invoices`).

## 8. Limitations selon le rôle
- Un Pro **ne peut pas** accéder au back-office admin ni aux espaces salarié.
- Certaines fonctionnalités (tableaux avancés, alertes prioritaires) sont réservées aux **abonnés Premium**.
