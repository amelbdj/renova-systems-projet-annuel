# 07 — Guide utilisateur : Salarié (équipe interne)

Le **Salarié** anime la plateforme : il publie des événements/formations et des articles, soumis à validation de l'administrateur. Espace : **`/salarie`**.

## 1. Connexion
1. Ouvrir `/login`, saisir un compte **Salarié** (voir `03_COMPTES_DE_DEMONSTRATION.md`).
2. **Résultat** : redirection vers **`/salarie`** (Espace Salariés).

## 2. Espace salarié
- **Page d'accueil** : `/salarie` (tableau de bord).
- Sous-espaces :
  - **Événements** : `/salarie/evenements`
  - **Planning** : `/salarie/planning`
  - **Contenus / Articles** : `/salarie/content`
  - **Forum** : `/salarie/forum`

*(Capture : `dashboard_salarie.png`.)*

## 3. Créer un événement / une formation
1. Aller sur **`/salarie/evenements`**.
2. Cliquer sur **créer un événement**.
3. Remplir : **titre**, **type**, **description**, **date de début/fin**, **lieu**, **capacité**, **tarif**, **image**, éventuellement un **plan (PDF)** et des **ressources pédagogiques (PDF)**.
4. Valider.
5. **Résultat** : « Événement soumis avec succès. Il est en attente de validation. » → statut **En attente**.

> **Prérequis important** : pour déposer un événement **payant**, le salarié doit disposer d'un **compte Stripe** (identifiant `stripe_account_id`). Sans cela, un message invite à créer son compte Stripe.

*(Capture : `creation_evenement.png`.)*

## 4. Planning
- **Page** : `/salarie/planning` : vue calendrier des événements, gestion des inscrits, des ressources.

## 5. Articles / actualités
- **Page** : `/salarie/content` : rédaction d'articles (titre, contenu, image). Soumis à **validation admin** avant publication.

## 6. Forum
- **Page** : `/salarie/forum` : participation aux discussions d'entraide.

## 7. Actions nécessitant une validation admin
| Action du salarié | Validé par |
|---|---|
| Publication d'un **événement / formation** | Administrateur (`/admin/validations`) |
| Publication d'un **article** | Administrateur |

- Après décision de l'admin (validation ou refus), le salarié reçoit une **notification** (cloche 🔔) l'informant du résultat.
