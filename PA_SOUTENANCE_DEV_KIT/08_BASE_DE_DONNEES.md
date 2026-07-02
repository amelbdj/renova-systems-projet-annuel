# 08 — Base de données (`pa2026`)

> SGBD : MySQL 8. Schéma + données : `db/init.sql` (import auto au 1er `docker compose up`). Dump de référence : `pa2026.sql`.
> Connexion locale : `docker exec -it uc_mysql mysql -uupcycle -pupcyclePass123 pa2026`.

## Tables principales (31 au total)

| Table | Rôle | Colonnes importantes |
|-------|------|----------------------|
| `utilisateur` | comptes | `id`, `nom`, `prenom`, `email`, `mot_de_passe` (bcrypt), `role` (enum), `validation`, `score`, `stripe_account_id`, `est_premium`, `plan_abo` |
| `annonce` | dons/ventes | `id`, `id_user`, `id_categorie`, `titre`, `description`, `type`, `prix`, `statut_validation`, `etat`, `poids_kg`, `ville`, `code_postal`, `image`, `is_sponsored`, `statut_vente`, `id_box`, `created_at` |
| `categorie` | matériaux | `id`, `libelle` (Textile, Bois, Plastique, Métal) |
| `conteneur` | points de collecte | `id`, `nom`, `adresse` |
| `box` / `box_conteneur` | casiers | `id`, `numero`, `statut`, `taille`, `id_conteneur` |
| `depot_box` / `historique_conteneurs` | dépôts/retraits | réservation, dépôt, retrait, dates, codes |
| `order` | commandes | `id`, acheteur, annonce, montant, commission, statut |
| `paiement` | paiements Stripe | montant, statut |
| `document` | factures/contrats PDF | `id_document`, `id_user`, `type_doc`, `url_pdf`, `id_commande` |
| `documents_legaux` | pièces justificatives users | `user_id`, `chemin_fichier` |
| `evenement` | events/formations | `id`, `id_salarie`, `titre`, `type`, `prix`, `date_debut`, `lieu`, `nb_places`, `image_url`, `statut_validation` |
| `inscription` | inscriptions events | user, event |
| `ressource_pedagogique` | PDF de formation | `url_fichier` |
| `article_news` | articles | `id_article`, `id_salarie`, `titre`, `contenu`, `image_url`, `statut` |
| `topic_forum` / `message_forum` | forum | `id_topic`, `titre` / `id_message`, `contenu`, `est_modere`, `est_signale` |
| `message` | messagerie chat | expéditeur, destinataire, `contenu`, `lu` |
| `notification` | notifs in-app | user, message, `est_lu` |
| `abonnement` / `plan_abo` | abonnements pro | plan, statut |
| `projet_pro` / `etapes_projet` | projets pro | avant/après, CO2 |
| `translations` / `languages` | multilingue | `lang_code`, `msg_key`, `msg_value` |
| `upcycling_score` | score éco | historique de points |
| `log_connexion` | logs auth | ip, date |
| `campagne_pub`, `partenaire`, `dictionnaire`, `plan_abo` | annexes | — |

## Relations / clés étrangères principales
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

## Requêtes SQL utiles

```sql
-- Voir tous les utilisateurs par rôle
SELECT id, prenom, nom, email, role, validation FROM utilisateur ORDER BY role;

-- Annonces en attente de validation
SELECT id, titre, prix, statut_validation FROM annonce WHERE statut_validation='En attente';

-- Catégories
SELECT * FROM categorie;

-- Conteneurs + nb de casiers
SELECT c.nom, c.adresse, COUNT(b.id) AS nb_box
FROM conteneur c LEFT JOIN box b ON b.id_conteneur=c.id GROUP BY c.id;

-- Score d'un utilisateur
SELECT id, prenom, score FROM utilisateur WHERE id=1;

-- Rendre visibles tous les messages du forum
UPDATE message_forum SET est_modere=0;

-- Donner un compte Stripe test à un salarié
UPDATE utilisateur SET stripe_account_id='acct_...' WHERE id=40;
```

## Import / export

```bash
# Export (sauvegarde)
docker exec uc_mysql mysqldump -uupcycle -pupcyclePass123 pa2026 > backup.sql

# Import (⚠️ utf8mb4 pour les accents)
docker exec -i uc_mysql mysql --default-character-set=utf8mb4 -uupcycle -pupcyclePass123 pa2026 < db/init.sql
```

## Générer une base vide vs remplie
- **Remplie** (démo) : `db/init.sql` contient schéma **+ données** → import automatique au 1er démarrage Docker, ou manuel (commande ci-dessus).
- **Vide** : importer uniquement les `CREATE TABLE` (retirer les `INSERT`), ou :
  ```sql
  -- vider une table sans supprimer la structure
  TRUNCATE TABLE annonce;
  ```

## Seed
- **Script de seed dédié : non trouvé** (`db/seed.*` inexistant). Le « seed » = les `INSERT` présents dans `db/init.sql`.
- **Plan proposé** (si demandé, à faire seulement sur validation) : créer `db/seed.sql` avec 1 compte par rôle (mot de passe bcrypt), 3–4 annonces, 1 conteneur + casiers, 1 événement, quelques catégories.
