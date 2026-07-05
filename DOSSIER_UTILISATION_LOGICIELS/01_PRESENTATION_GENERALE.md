# 01 — Présentation générale

## UpcycleConnect en une phrase
**UpcycleConnect** est une plateforme web d'**économie circulaire** qui met en relation des particuliers, des professionnels (artisans) et l'équipe interne pour **donner, vendre, récupérer et transformer** des objets et matériaux, via un réseau de **conteneurs connectés**.

- **Produit** : UpcycleConnect
- **Prestataire** : Renova Systems
- **URL publique** : **https://upcycleconnect.pro/**

## Objectif métier
Réduire le gaspillage en donnant une seconde vie aux objets et matériaux :
- un **particulier** dépose un objet (don ou vente) dans un casier d'un conteneur ;
- un **professionnel/artisan** récupère la matière pour la transformer (projets d'upcycling) ;
- la plateforme **mesure l'impact** (Upcycling Score, CO₂ évité) et **sécurise les transactions** (paiement Stripe, commission).

## Public cible
- **Particuliers** soucieux de donner/vendre plutôt que jeter.
- **Professionnels / artisans** en quête de matière première de réemploi.
- **Collectivités / structures** exploitant le réseau de conteneurs (via l'équipe interne).

## Rôles utilisateurs
| Rôle | Description | Espace |
|---|---|---|
| **Visiteur** | Non connecté : accès aux pages publiques (accueil, connexion, inscription) | — |
| **Utilisateur (Particulier)** | Crée des annonces, dépose/récupère en conteneur, suit son score éco | `/client` |
| **Professionnel (Pro)** | Consulte/achète les annonces, gère abonnements et projets d'upcycling | `/pro` |
| **Salarié** | Publie événements/formations et articles (soumis à validation) | `/salarie` |
| **Administrateur** | Back-office complet : utilisateurs, validations, conteneurs, finances, documents | `/admin` |

## Fonctionnalités principales
- Comptes et **authentification sécurisée** (JWT) avec 4 rôles.
- **Annonces** de dons/ventes avec photos, catégories et **validation** par l'admin.
- **Réseau de conteneurs** : réservation d'un casier, dépôt par code PIN, retrait par code-barres.
- **Paiement en ligne** (Stripe) avec **commission de 5 %** reversée à la plateforme.
- **Upcycling Score** : mesure de l'impact écologique de chaque utilisateur.
- **Événements & formations** organisés par les salariés.
- **Espace professionnel** : abonnements (freemium/premium), projets avant/après, sponsoring.
- **Notifications** (cloche in-app + push), **multilingue** (FR/EN), **factures PDF**.
- **Back-office** d'administration complet.

## Architecture logicielle (résumé)
```
[Navigateur] → HTTPS → [Nginx reverse-proxy]
        ├── fichiers statiques (Front HTML/CSS/JS)
        ├── /api/*   → API Go (port 8081)
        └── /uploads/* → fichiers (images, PDF)
                              │
                        [API Go] ─ JWT + rôles ─ requêtes SQL
                              │
                        [MySQL 8]  base « pa2026 »
```
- **Front** et **API** sont **séparés** et déployés dans des conteneurs Docker distincts.
- Toute action sensible est **vérifiée côté serveur** (le front n'est jamais une source de confiance).
- Les fichiers uploadés sont stockés dans un **volume Docker persistant** (jamais en base).
