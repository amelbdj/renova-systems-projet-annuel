# 08 — Guide Administrateur (back-office)

L'**Administrateur** dispose d'un back-office complet. Espace : **`/admin`**. Toutes les routes `/admin/*` sont protégées côté serveur (JWT + rôle).

## Connexion
1. Ouvrir `/login`, saisir un compte **Administrateur** (voir `03_COMPTES_DE_DEMONSTRATION.md`).
2. **Résultat** : redirection vers **`/admin`** (Vue d'ensemble). Le menu latéral donne accès aux modules.

*(Capture : `backoffice_admin.png`.)*

---

## Module 1 — Vue d'ensemble (`/admin`)
- **Objectif** : synthèse en temps réel.
- **Contenu** : KPI (utilisateurs actifs, revenus du mois, conteneurs/casiers, déchets sauvés), **Centre d'actions requises** (annonces/événements/articles en attente, casiers en panne), **Activité récente** (dernières annonces), **badge de validations** dans le menu.
- **Modules intégrés à cette page** : **Envoyer une notification**, **Catégories** (liste/ajout/suppression), **Langues / Traductions** (import d'une langue).
- **Précaution** : les compteurs sont calculés à chaque chargement ; rafraîchir la page après une action.

## Module 2 — Utilisateurs (`/admin/users`)
- **Objectif** : gérer les comptes.
- **Actions** :
  - Lister les utilisateurs (**pagination 10/page**), rechercher, filtrer par rôle.
  - **Valider** / **Refuser** un compte en attente (bouton ✓ / ✕).
  - **Bannir** un utilisateur.
  - **Ajouter** un utilisateur (choix du rôle, y compris Salarié/Admin).
  - **Modifier** / **Supprimer** un utilisateur.
- **Résultat attendu** : le statut se met à jour immédiatement dans le tableau.
- **Précaution** : un refus/bannissement peut déclencher un e-mail/une notification ; vérifier l'identité avant d'agir.

*(Captures : `validation_utilisateur.png`, `liste_utilisateurs.png`.)*

## Module 3 — Validations (`/admin/validations`)
- **Objectif** : modérer les contenus soumis.
- **Onglets** : Tout voir / Annonces / Événements / Contenus (articles).
- **Actions** : **Approuver** ou **Refuser** une annonce, un événement, un article.
- **Résultat attendu** : l'élément validé/refusé **disparaît de la liste en direct** ; le badge de validations se met à jour ; l'auteur reçoit une notification.

*(Capture : `validation_annonce.png`.)*

## Module 4 — Logistique & Box (`/admin/conteneurs`)
- **Objectif** : gérer le réseau physique.
- **Actions** :
  - Voir les conteneurs (nom, adresse, nombre de casiers) et l'état des casiers (libre / occupé / maintenance).
  - **Créer un conteneur**, **ajouter un casier**, **changer le statut** d'un casier.
  - **Générer un rapport logistique PDF** (bouton « 📄 Rapport logistique »).
- **Précaution** : ne pas supprimer un conteneur contenant des objets déposés.

*(Capture : `conteneurs.png`.)*

## Module 5 — Finances (`/admin/finances`)
- **Objectif** : suivre l'activité financière.
- **Contenu** : volume d'affaires du mois, **revenus (commission 5 %)**, tableau des transactions (facture/annonce, montant, commission, statut).
- Endpoints : `GET /admin/finance/overview`, `GET /admin/finance/transactions`.

*(Capture : `finances.png`.)*

## Module 6 — Documents (`/admin/documents`)
- **Objectif** : consulter les factures/contrats PDF.
- **Actions** : lister les documents (type, utilisateur, commande) et **ouvrir le PDF** (« Ouvrir le PDF »).
- Les PDF sont générés côté serveur lors d'un paiement (facture) ou d'une souscription (contrat) et stockés dans le volume `uploads`.

## Module 7 — Envoyer une notification (sur `/admin`)
- **Objectif** : communiquer avec une audience.
- **Actions** : choisir la cible (Particuliers / Professionnels / Tous), saisir un message, **Envoyer**.
- **Résultat** : chaque destinataire reçoit la notification (cloche + push OneSignal en prod).

## Module 8 — Catégories (sur `/admin`)
- **Objectif** : gérer les catégories de matériaux (Textile, Bois, Plastique, Métal…).
- **Actions** : voir la liste, **＋ Ajouter** (libellé), **supprimer**.
- **Résultat** : la nouvelle catégorie apparaît immédiatement dans le filtre de la page annonces.

## Module 9 — Langues / Traductions (sur `/admin`)
- **Objectif** : ajouter une langue à l'interface (multilingue).
- **Actions** :
  1. **⬇ Télécharger le modèle (FR)** → fichier JSON des textes en français.
  2. Traduire les **valeurs** (garder les clés) dans le fichier.
  3. **＋ Ajouter une nouvelle langue** → saisir le **code** (ex : `es`), le **nom** (ex : `Español`), sélectionner le fichier, **Importer**.
- **Résultat** : la langue apparaît dans le sélecteur ; l'import est réservé aux administrateurs (route protégée).

## Statistiques
- Présentes sous forme de **KPI** sur la Vue d'ensemble et la page Finances (pas de module « statistiques » séparé).

## Multilingue
- Géré via `data-i18n` + table `translations` ; import de langue décrit ci-dessus (Module 9).
