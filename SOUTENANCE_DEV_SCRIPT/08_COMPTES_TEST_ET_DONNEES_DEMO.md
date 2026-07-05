# 08 - Comptes test et donnees demo

## Comptes trouves dans le projet

Source : `db/init.sql`, table `utilisateur`.

| Email | Role | Statut | Usage conseille |
|---|---|---|---|
| `test.client@renova.test` | Utilisateur | Valide | Demo dashboard particulier |
| `test.admin@renova.test` | Administrateur | Valide | Demo back-office |
| `test.salarie@renova.test` | Salarie | Valide | Demo moderation / evenements |
| `test.pro@renova.test` | Pro | Valide | Demo espace pro |
| `test.user@test.fr` | Utilisateur | Valide | Compte secours client |
| `test.salarie@test.fr` | Salarie | Valide | Compte secours salarie |
| `test.admin@test.fr` | Administrateur | Valide | Compte secours admin |
| `test.pro@test.fr` | Pro | Valide | Compte secours pro |

## Mot de passe

Les mots de passe sont stockes en bcrypt dans la base. Ils ne sont donc pas lisibles depuis le dump SQL.

Action recommandee avant soutenance :

1. Choisir un mot de passe commun de demo, par exemple `Test123!`.
2. Le definir pour les 4 comptes principaux.
3. Verifier une connexion pour chaque role.
4. Noter les identifiants sur une fiche papier pour eviter la panique.

## Donnees utiles deja presentes

### Annonces

Source : table `annonce`.

Utilisation :
- montrer catalogue ;
- montrer mes annonces ;
- montrer statut annonce ;
- preparer une annonce en attente pour validation.

Action avant demo :
- garder au moins une annonce validee ;
- garder au moins une annonce en attente ;
- ne pas tout valider avant la soutenance.

### Evenements

Source : table `evenement`.

Donnees presentes :
- formations ;
- ateliers ;
- evenements valides ;
- au moins un evenement en attente selon dump.

Utilisation :
- montrer catalogue evenements ;
- montrer creation salarie ;
- montrer validation admin.

### Conteneurs et box

Sources : tables `conteneur`, `box`, `historique_conteneurs`.

Utilisation :
- montrer administration logistique ;
- expliquer depot/retrait ;
- montrer code PIN ou code de recuperation si donnee prete.

Action avant demo :
- reperer un code de depot/retrait utilisable ;
- ne pas consommer ce code avant la soutenance.

### Traductions

Sources : tables `languages`, `translations`.

Utilisation :
- montrer select langue FR/EN ;
- montrer route API `/api/translations?lang=fr`.

### Documents / PDF

Sources : table `document`, `API/admin/pdf.go`.

Utilisation :
- montrer factures/contrats si fichiers presents ;
- sinon montrer code et capture.

### Projets pro

Sources : tables `projet_pro`, `etapes_projet`.

Utilisation :
- montrer espace pro ;
- montrer suivi de projets d'upcycling.

## Jeu de demo conseille

- Client : creer ou afficher une annonce.
- Admin : valider/refuser une annonce.
- Salarie : creer un evenement ou moderer le forum.
- Pro : consulter catalogue/projets/premium.

## Donnees a ne pas modifier pendant la demo

- Le compte admin principal.
- Le compte salarie principal.
- Une annonce en attente.
- Une annonce validee.
- Un conteneur avec box disponible.
- Un code PIN prepare.

## Verification 30 minutes avant

- [ ] `docker compose ps` OK.
- [ ] Front Docker accessible.
- [ ] Login client OK.
- [ ] Login admin OK.
- [ ] Login salarie OK.
- [ ] Login pro OK.
- [ ] Admin validations contient au moins une donnee.
- [ ] Catalogue contient au moins une annonce.
- [ ] Conteneurs visibles.
- [ ] Traductions chargees.
- [ ] Captures ouvertes dans un dossier accessible.
