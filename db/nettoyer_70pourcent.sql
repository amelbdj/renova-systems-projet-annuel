-- ============================================================
--  Nettoyage : supprime ~70% des annonces seedees (trop nombreuses)
--
--  On GARDE :
--    - les vraies annonces d'origine (id_user < 100000)
--    - toutes les annonces sponsorisees (is_sponsored = 1)
--    - les annonces vendues ou en box (interessantes pour la demo)
--
--  On SUPPRIME 70% du reste (les fausses annonces seedees),
--  reparti uniformement grace a (id % 10) < 7.
--
--  Les enfants (messages, depots, commandes, historique) lies a
--  ces annonces sont supprimes avant, pour ne pas casser les cles etrangeres.
-- ============================================================

USE pa2026;

SET SQL_SAFE_UPDATES = 0;

-- 1) On liste les annonces a supprimer dans une table temporaire
DROP TEMPORARY TABLE IF EXISTS tmp_annonces_a_supprimer;
CREATE TEMPORARY TABLE tmp_annonces_a_supprimer AS
SELECT id
FROM annonce
WHERE id_user >= 100000
  AND is_sponsored = 0
  AND statut_vente NOT IN ('VENDU', 'EN BOX')
  AND (id % 10) < 7;

-- 2) On supprime d'abord les lignes enfants qui pointent vers ces annonces
DELETE FROM historique_conteneurs
WHERE annonce_id IN (SELECT id FROM tmp_annonces_a_supprimer);

DELETE FROM depot_box
WHERE id_annonce IN (SELECT id FROM tmp_annonces_a_supprimer);

DELETE FROM `order`
WHERE id_annonce IN (SELECT id FROM tmp_annonces_a_supprimer);

-- (la table `message` a ON DELETE CASCADE : ses lignes partent automatiquement)

-- 3) On supprime enfin les annonces elles-memes
DELETE FROM annonce
WHERE id IN (SELECT id FROM tmp_annonces_a_supprimer);

DROP TEMPORARY TABLE IF EXISTS tmp_annonces_a_supprimer;

SET SQL_SAFE_UPDATES = 1;

-- 4) Verification : combien reste-t-il d'annonces ?
SELECT COUNT(*) AS annonces_restantes FROM annonce;
