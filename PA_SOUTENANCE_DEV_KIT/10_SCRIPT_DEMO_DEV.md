# 10 — Script de démonstration (15 minutes)

> Objectif : montrer une plateforme **complète, déployée en ligne, multi-rôles**, en valorisant les **4 tableaux de bord** (Particulier, Professionnel, Salarié, Administrateur) et leurs **fonctionnalités clés**.
> Support : **https://upcycleconnect.pro/** (prod, preuve de mise en ligne). Garder un onglet local `http://localhost:8088` prêt pour une éventuelle modif en direct.
> Avoir les 4 comptes de test déjà notés (voir `09_COMPTES_TEST.md`).

## Minutage global (15 min)
| Temps | Séquence |
|---|---|
| 0:00 – 1:30 | Intro + accueil + connexion |
| 1:30 – 4:00 | **Dashboard Particulier** |
| 4:00 – 6:30 | **Dashboard Professionnel** |
| 6:30 – 9:00 | **Dashboard Salarié** |
| 9:00 – 13:00 | **Dashboard Administrateur** (le point fort) |
| 13:00 – 15:00 | Preuve technique (Docker / API / HTTPS) + conclusion |

---

## 0:00 – 1:30 · Introduction & connexion
**À montrer** : la page d'accueil sur `https://upcycleconnect.pro/`.
**À dire :**
> « UpcycleConnect est une plateforme d'économie circulaire : les particuliers donnent ou vendent des objets, les professionnels récupèrent la matière, via un réseau de **conteneurs connectés**. Elle est **déployée en ligne**, en HTTPS, sur une infrastructure Docker. Elle gère **4 rôles**, chacun avec son tableau de bord dédié. »

- Montrer rapidement : le **sélecteur de langue** (FR/EN) et les **URLs propres** (`/login`, `/annonces`).
- Cliquer **Connexion** → se connecter en **Particulier**.

---

## 1:30 – 4:00 · Dashboard PARTICULIER (`/client`)
> Fil conducteur : « le parcours d'un objet, du dépôt à la récupération ».

**Fonctionnalités clés à mettre en valeur :**
1. **Tableau de bord dynamique** : montrer les 3 indicateurs en haut — **Annonces**, **Dépôt actif**, **Score éco** (« ces chiffres sont calculés en temps réel, pas codés en dur »).
2. **Créer une annonce** (action forte) : cliquer **＋ Ajouter une annonce** → remplir titre, **catégorie** (issue de la base), type Vente/Don, prix, **photo** → **Soumettre**.
   > « L'annonce part en **validation** : rien n'est publié sans contrôle admin. »
3. **Marketplace** (`/annonces`) : montrer les **filtres par catégorie** (dynamiques) et le **thème violet** du particulier.
4. **Conteneurs & Upcycling Score** : montrer la section conteneurs (code PIN / code-barres) et l'encart **score écologique**.
   > « Quand l'objet déposé est récupéré, l'utilisateur gagne des points d'impact. »

**À dire (transition) :** « Voyons maintenant le même écosystème côté professionnel. »

---

## 4:00 – 6:30 · Dashboard PROFESSIONNEL (`/pro`)
> Se déconnecter, se reconnecter en **Pro**.

**Fonctionnalités clés :**
1. **Interface adaptée au rôle** : montrer que l'espace pro est en **teal**, et que la page annonces reprend **les couleurs pro** (thème dynamique selon le rôle).
2. **Modèle Freemium / Premium** : montrer l'encart abonnement (« Passer au Premium »).
   > « Le paiement passe par **Stripe**, avec une **commission de 5 %** prélevée par la plateforme via Stripe Connect. »
3. **Projets d'upcycling** : montrer un projet avec **photos avant / après** et **CO₂ évité**.
4. **Sponsoring** : mentionner qu'un pro peut mettre une annonce en avant.

**À dire (transition) :** « Ces contenus et événements sont animés par l'équipe interne — les salariés. »

---

## 6:30 – 9:00 · Dashboard SALARIÉ (`/salarie`)
> Se reconnecter en **Salarié**.

**Fonctionnalités clés :**
1. **Espace interne** : présenter les sous-espaces (Événements, Planning, Contenus, Forum).
2. **Créer un événement / une formation** (`/salarie/evenements`) : ouvrir le formulaire, montrer les champs (titre, date, lieu, tarif, **image**, **PDF de plan/ressources**) → soumettre.
   > « L'événement part **en attente de validation**. Et un salarié doit avoir un **compte Stripe** pour proposer un événement payant — contrôle fait côté serveur. »
3. **Notifications** : montrer la **cloche** 🔔 — « le salarié est notifié quand l'admin valide ou refuse son contenu ».

**À dire (transition) :** « Toutes ces validations, c'est le rôle de l'administrateur — le cœur du back-office. »

---

## 9:00 – 13:00 · Dashboard ADMINISTRATEUR (`/admin`) — POINT FORT
> Se reconnecter en **Admin**. Prendre le temps ici, c'est le module le plus riche.

**Fonctionnalités clés (dans l'ordre) :**
1. **Vue d'ensemble** : KPI (utilisateurs, revenus, conteneurs), **activité récente dynamique**, **badge de validations** dans le menu.
2. **Validations** (`/admin/validations`) — *effet « waouh »* : approuver **l'annonce créée à l'étape Particulier** → elle **disparaît en direct** de la liste, et le badge se met à jour.
3. **Utilisateurs** (`/admin/users`) : montrer la **pagination (10/page)**, la **recherche**, et **valider / refuser / bannir** un compte.
4. **Logistique & Box** (`/admin/conteneurs`) : ouvrir un conteneur → casiers, puis cliquer **« 📄 Rapport logistique »** → **un PDF se télécharge** (génération côté client).
5. **Finances** (`/admin/finances`) : volume, **commission 5 %**, transactions.
6. **Modules avancés** (sur la Vue d'ensemble) : montrer rapidement **Catégories** (ajout en direct → apparaît dans le filtre annonces), **Langues/Traductions** (import d'une langue), **Envoyer une notification** (cloche + push).

**À dire :**
> « Chaque action sensible est **vérifiée côté serveur** avec un contrôle de rôle : le back-office est totalement protégé, un particulier ne peut jamais y accéder même en manipulant le front. »

---

## 13:00 – 15:00 · Preuve technique & conclusion
**À montrer (terminal + navigateur) :**
- `docker ps` → **3 conteneurs Up** (base, API Go, front Nginx) : « front et back séparés, orchestrés par Docker Compose ».
- Un appel **API** en direct (Postman/curl) : `POST /admin/login` → token, puis une route protégée : « API en Go, sécurisée par **JWT** ».
- La barre d'adresse **https://upcycleconnect.pro** avec le **cadenas HTTPS** : « déployé en ligne, derrière Nginx, sur IP publique — ce n'est pas du localhost ».

**Phrase de clôture :**
> « En résumé : 4 rôles, 4 tableaux de bord, un parcours complet du dépôt à la récupération, un paiement sécurisé, le tout déployé en production avec Docker, Nginx et HTTPS. »

---

## À ÉVITER pendant la démo
- Dérouler un **paiement Stripe complet** en live (dépend des comptes Connect + webhook — risque de blocage). Le **décrire** suffit.
- Tester le **push OneSignal en local** (désactivé sur localhost, HTTPS requis) — le montrer en prod ou l'expliquer.
- Ouvrir des données de test brutes/incohérentes (comptes « Rejeté », annonces `azerty…`).
- Rester bloqué sur un bug live : avoir des **captures de secours** de chaque écran clé.

## Plan B (si le réseau/la prod tombe)
- Basculer sur `http://localhost:8088` (stack Docker locale) — même parcours.
- Si tout tombe : dérouler avec les **captures d'écran** de secours.
