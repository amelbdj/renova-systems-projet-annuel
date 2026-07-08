# 09 — Comptes de test

> Les mots de passe sont **hashés (bcrypt)** en base : impossibles à extraire. Les comptes ci-dessous existent réellement dans `pa2026` (table `utilisateur`).
> **Mot de passe : à confirmer** (celui que tu as défini au seed / que tu connais). Note-le dans la colonne prévue avant la soutenance.

## Comptes existants (validés, prêts pour la démo)

| Rôle | Email | Statut | Mot de passe (à remplir) |
|------|-------|--------|--------------------------|
| **Administrateur** | `test.admin@renova.test` | Validé | ____________ |
| Administrateur | `test.admin@test.fr` | Validé | ____________ |
| Administrateur (réel) | `amelbdj213@gmail.com` | Validé | ____________ |
| **Salarié** | `test.salarie@renova.test` | Validé | ____________ |
| Salarié | `test.salarie@test.fr` | Validé | ____________ |
| **Pro** | `test.pro@renova.test` | Validé | ____________ |
| Pro | `test.pro@test.fr` | Validé | ____________ |
| **Utilisateur** (particulier) | `test.client@renova.test` | Validé | ____________ |
| Utilisateur | `test.user@test.fr` | Validé | ____________ |

> Vérifier la liste à jour : `SELECT email, role, validation FROM utilisateur WHERE validation='Validé';`

## Usage en démo (parcours par compte)

### Administrateur (`test.admin@renova.test`) → redirige vers `/admin`
- Vue d'ensemble (KPI, activité récente dynamique, badge validations).
- **Utilisateurs** (`/admin/users`) : pagination 10/page, valider/refuser/bannir.
- **Validations** (`/admin/validations`) : approuver une annonce/un event → il disparaît en direct.
- **Finances** (`/admin/finances`), **Documents** (PDF), **Conteneurs** (`/admin/conteneurs` + bouton **Rapport logistique PDF**).

### Salarié (`test.salarie@renova.test`) → `/salarie`
- Créer un **événement/formation** (⚠️ nécessite `stripe_account_id`), planning, forum, articles.

### Pro (`test.pro@renova.test`) → `/pro`
- Espace pro **en teal**, abonnements, projets avant/après, sponsoriser une annonce.
- Page annonces/profil aux **couleurs pro**.

### Utilisateur (`test.client@renova.test`) → `/client`
- Créer une **annonce** (don/vente), voir le tableau de bord (Annonces, Dépôt actif, Score éco — tous dynamiques), déposer/récupérer en conteneur (simulateur).

## Si tu veux (re)créer des comptes propres
- Via l'inscription : page `/register` (rôle Utilisateur/Pro).
- Via l'admin : `/admin/users` → « Ajouter » (permet de choisir le rôle).
- Un compte créé peut être **En attente** → le valider depuis l'admin avant la démo.
