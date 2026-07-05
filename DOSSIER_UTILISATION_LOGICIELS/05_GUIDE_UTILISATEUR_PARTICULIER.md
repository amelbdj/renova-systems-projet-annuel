# 05 — Guide utilisateur : Particulier

Le **Particulier** donne ou vend des objets et suit son impact écologique. Espace : **`/client`**.

## 1. Connexion
1. Ouvrir `/login`.
2. Saisir l'email et le mot de passe du compte particulier (voir `03_COMPTES_DE_DEMONSTRATION.md`).
3. **Résultat** : redirection automatique vers **`/client`** (Espace Particulier).

*(Capture : `connexion.png`.)*

## 2. Tableau de bord
- **Page** : `/client`.
- Affiche en haut trois indicateurs **dynamiques** : **Annonces** (nombre d'annonces actives), **Dépôt actif** (objets en conteneur), **Score éco**.
- Plus bas : la liste **Mes Annonces** et le **Système de Conteneurs**.

*(Capture : `dashboard_particulier.png`.)*

## 3. Créer une annonce
1. Sur `/client`, cliquer sur **＋ (Ajouter une annonce)** dans la section « Mes Annonces ».
2. Remplir le formulaire :
   - **Titre**, **Catégorie** (liste issue de la base : Textile, Bois, Plastique, Métal), **Type** (Vente / Don gratuit), **Prix** (si vente), **Description**.
   - **Photo de l'objet** : bouton **Choisir un fichier** (formats image, taille max 5 Mo).
3. Cliquer sur **Soumettre l'annonce**.
4. **Résultat attendu** : message « Votre annonce sera soumise à validation avant d'être publiée (délai 24 h max) ». L'annonce apparaît avec le statut **En attente**.

*(Captures : `creation_annonce.png`, `upload_photo.png`.)*

## 4. Suivre le statut d'une annonce
- Dans « Mes Annonces », chaque carte affiche un **badge de statut** :
  - **En attente** : en cours de validation par l'admin.
  - **Validée** : publiée et visible sur la marketplace.
  - **Rejetée** : refusée par l'admin.
- Boutons disponibles : **Modifier** et **Supprimer**.

## 5. Consulter / rechercher les annonces
- **Page** : `/annonces`.
- **Filtres** : par **catégorie** (dynamique depuis la base) et par **type** (Vente / Don).
- **Recherche** par mot-clé, **tri** (plus récentes, prix, A→Z), vue grille/liste.

*(Capture : `marketplace.png`.)*

## 6. Conteneurs : dépôt et retrait
Quand un objet est vendu/réservé, un **casier** est attribué :
1. Le particulier obtient un **code PIN** (dépôt) et un **code-barres** (retrait), visibles dans la section conteneurs de `/client`.
2. **Dépôt** : au conteneur, saisir le **PIN** → l'objet est marqué déposé.
3. **Retrait** : l'acheteur scanne le **code-barres** → l'objet est récupéré.
- Un **simulateur** (`/simulateur`) permet de tester dépôt et retrait sans matériel physique.

*(Capture : `conteneurs_particulier.png`.)*

## 7. Upcycling Score
- **Page** : `/client`, encart « Mon Upcycling Score ».
- Le score augmente lorsqu'un objet déposé est **récupéré** (calcul : poids × coefficient).
- Affiche aussi les statistiques d'impact (objets donnés, déchets évités).

## 8. Notifications
- Une **cloche** 🔔 affiche une pastille rouge en cas de notification non lue.
- **Clic** sur la cloche → panneau listant les notifications (ex : annonce validée). Les notifications passent en « lues » à l'ouverture.

## 9. Profil
- **Page** : `/profil` (accessible via l'avatar en haut à droite). Permet de consulter/mettre à jour ses informations. La couleur de la page s'adapte au rôle.
