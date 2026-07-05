# 03 — Comptes de démonstration

> Ces comptes existent réellement dans la base fournie (`db/export_db_remplie.sql`, table `utilisateur`).
> **Les mots de passe sont hashés (bcrypt) en base** : ils ne peuvent pas être extraits du dump. → **à compléter** avec le mot de passe défini lors du seed.

## Comptes disponibles (validés)

| Rôle | Email | Mot de passe | Utilité du compte | Parcours à tester |
|---|---|---|---|---|
| **Administrateur** | `test.admin@renova.test` | *à compléter* | Back-office complet | Valider un compte / une annonce, voir finances, conteneurs, catégories |
| Administrateur | `test.admin@test.fr` | *à compléter* | (secours) | idem |
| **Salarié** | `test.salarie@renova.test` | *à compléter* | Espace interne | Créer un événement/formation, un article (soumis à validation) |
| Salarié | `test.salarie@test.fr` | *à compléter* | (secours) | idem |
| **Professionnel** | `test.pro@renova.test` | *à compléter* | Espace pro | Consulter les annonces, abonnement, projet upcycling, sponsoring |
| Professionnel | `test.pro@test.fr` | *à compléter* | (secours) | idem |
| **Particulier** | `test.client@renova.test` | *à compléter* | Espace client | Créer une annonce, suivre son statut, score éco, conteneurs |
| Particulier | `test.user@test.fr` | *à compléter* | (secours) | idem |

> **Vérifier la liste à jour** : `SELECT email, role, validation FROM utilisateur WHERE email LIKE 'test.%';`

## Visiteur (sans compte)
Aucun identifiant : accès aux pages publiques uniquement (accueil, connexion, inscription). Voir `04_GUIDE_UTILISATEUR_VISITEUR.md`.

## Si vous devez (re)créer des comptes de démonstration
1. **Par l'inscription** : page `/register` (rôles Particulier et Professionnel uniquement — l'inscription en Admin/Salarié est volontairement bloquée côté serveur pour la sécurité).
2. **Par l'administration** : `/admin` → carte/section Utilisateurs → « Ajouter » (permet de choisir n'importe quel rôle, y compris Salarié/Admin).
3. Un compte fraîchement inscrit peut être **« En attente »** → le valider depuis l'admin avant la démonstration.

### Liste de comptes recommandés à préparer pour une démo
| Rôle | Email suggéré | Mot de passe suggéré |
|---|---|---|
| Administrateur | `test.admin@renova.test` | *(à définir, ex : Admin2026!)* |
| Salarié | `test.salarie@renova.test` | *(à définir)* |
| Professionnel | `test.pro@renova.test` | *(à définir)* |
| Particulier | `test.client@renova.test` | *(à définir)* |

> Pour redéfinir un mot de passe connu sur un compte de test, le plus simple est de recréer le compte via `/register` (particulier/pro) ou via l'admin (tous rôles), puis de le valider.
