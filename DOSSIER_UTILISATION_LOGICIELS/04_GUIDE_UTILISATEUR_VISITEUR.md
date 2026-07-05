# 04 — Guide utilisateur : Visiteur (non connecté)

Le **visiteur** est un internaute non authentifié. Il a accès aux pages publiques mais pas aux espaces personnels.

## Accéder à la page d'accueil
1. Ouvrir **https://upcycleconnect.pro/** (URL `/`).
2. La page d'accueil (landing) présente le concept UpcycleConnect.
3. Le **sélecteur de langue** (FR/EN) est disponible en bas à droite de l'écran.

*(Capture à ajouter : `accueil.png`.)*

## Pages accessibles sans compte
| Page | URL propre | Contenu |
|---|---|---|
| Accueil | `/` | Présentation du service |
| Connexion | `/login` | Formulaire d'authentification |
| Inscription | `/register` | Création de compte (Particulier ou Professionnel) |
| Mot de passe oublié | `/reset-password` | Réinitialisation |
| À propos | `/a-propos` | Informations sur le projet |

> Les pages « métier » (annonces, espaces client/pro/salarié/admin) exigent une **connexion**. Un garde de session (`Frontend/script/auth_guard.js`) redirige vers `/login` si aucun jeton n'est présent.

## Consulter / créer un compte
### S'inscrire
1. Cliquer sur **S'inscrire** (`/register`).
2. Choisir un **profil** : **Particulier** ou **Professionnel** (les rôles Salarié/Administrateur ne sont pas proposés à l'inscription).
3. Remplir le formulaire :
   - **Particulier** : nom, prénom, email, mot de passe, **date de naissance** (l'inscription est refusée si moins de **18 ans**).
   - **Professionnel** : + **nom d'entreprise** et **SIRET** (14 chiffres, validé par clé de contrôle).
4. Valider → le compte est créé. Un **Particulier** est actif immédiatement ; un **Professionnel** peut être **« En attente »** de validation par un administrateur.

### Se connecter
1. Cliquer sur **Connexion** (`/login`).
2. Saisir email + mot de passe → redirection automatique vers l'espace correspondant au rôle.

*(Captures à ajouter : `connexion.png`, `inscription.png`.)*

## Limitations d'un visiteur
- Ne peut **pas** créer d'annonce, ni déposer/récupérer en conteneur.
- Ne voit **pas** les espaces client/pro/salarié/admin.
- Doit **s'inscrire ou se connecter** pour toute action personnalisée.
