# 11 — Base de données

## Informations générales
- **Nom de la base** : `pa2026`
- **Type** : **MySQL 8** (moteur InnoDB, charset `utf8mb4`)
- **Utilisateur applicatif** : `upcycle` (mot de passe local par défaut : `upcyclePass123`) — `root` disponible pour l'admin.
- **Connexion depuis l'API** : `API/bdd/db.go` (`NewDB()`), identifiants via variables d'environnement.

## Emplacement des exports SQL
| Fichier | Contenu | Chemin |
|---|---|---|
| **`export_db_vide.sql`** | **Structure seule** (CREATE TABLE, sans données) | `db/export_db_vide.sql` |
| **`export_db_remplie.sql`** | **Structure + données** (31 tables + jeux de données de démonstration) | `db/export_db_remplie.sql` |
| `init.sql` | Schéma + données importé automatiquement au 1er démarrage Docker | `db/init.sql` |

> Les deux fichiers `export_db_vide.sql` et `export_db_remplie.sql` sont **fournis** dans `db/`. Ils ont été générés avec `mysqldump` (voir commandes ci-dessous).

## Base vide vs base remplie
- **Base vide** (`export_db_vide.sql`) : uniquement la **structure** (tables, colonnes, clés). Utile pour repartir d'une base neuve et y injecter ses propres données.
- **Base remplie** (`export_db_remplie.sql`) : la structure **+** des **données de démonstration** (comptes de test, catégories, conteneurs, annonces, événements…). C'est celle à utiliser pour une **démonstration immédiate**.

## Commandes d'import
```bash
# Vers le conteneur MySQL (recommandé) — toujours en utf8mb4 pour les accents
docker exec -i uc_mysql mysql --default-character-set=utf8mb4 -uupcycle -pupcyclePass123 pa2026 < db/export_db_remplie.sql

# Vers un MySQL local (sans Docker)
mysql --default-character-set=utf8mb4 -uroot -p pa2026 < db/export_db_remplie.sql
```

## Commandes d'export (régénérer les fichiers)
```bash
# Base remplie (structure + données)
docker exec uc_mysql mysqldump --default-character-set=utf8mb4 -uupcycle -pupcyclePass123 --databases pa2026 > db/export_db_remplie.sql

# Base vide (structure seule)
docker exec uc_mysql mysqldump --default-character-set=utf8mb4 -uupcycle -pupcyclePass123 --no-data --databases pa2026 > db/export_db_vide.sql
```

## Tables principales (31 au total)
| Table | Rôle |
|---|---|
| `utilisateur` | comptes (rôle, validation, score, siret, stripe_account_id…) |
| `annonce` | dons/ventes (titre, prix, type, statut_validation, statut_vente, image, id_categorie, id_user…) |
| `categorie` | catégories de matériaux (Textile, Bois, Plastique, Métal) |
| `conteneur`, `box`, `box_conteneur` | réseau de casiers |
| `depot_box`, `historique_conteneurs` | cycle de vie dépôt/retrait (PIN, code-barres, dates) |
| `order`, `paiement` | commandes et paiements (commission) |
| `document`, `documents_legaux` | factures/contrats PDF, pièces justificatives |
| `evenement`, `inscription`, `ressource_pedagogique` | événements/formations |
| `article_news` | articles/actualités |
| `topic_forum`, `message_forum`, `message` | forum et messagerie |
| `notification` | notifications in-app |
| `abonnement`, `plan_abo`, `projet_pro`, `etapes_projet` | espace pro |
| `translations`, `languages` | multilingue |
| `upcycling_score`, `log_connexion` | impact éco, logs |

## Relations importantes (clés étrangères)
```
utilisateur (1) ──< annonce (id_user)
categorie   (1) ──< annonce (id_categorie)
utilisateur (1) ──< evenement (id_salarie)
evenement   (1) ──< inscription >── (1) utilisateur
conteneur   (1) ──< box (id_conteneur)
annonce     (1) ──< order >── (1) utilisateur (acheteur)
utilisateur (1) ──< document (id_user)
topic_forum (1) ──< message_forum (id_topic)
utilisateur (1) ──< notification
```

## Comptes de test présents dans la base remplie
La base remplie contient les comptes de démonstration : `test.admin@renova.test`, `test.salarie@renova.test`, `test.pro@renova.test`, `test.client@renova.test` (+ variantes `@test.fr`).
Mots de passe **hashés (bcrypt)** → non lisibles dans le dump. Voir `03_COMPTES_DE_DEMONSTRATION.md`.
