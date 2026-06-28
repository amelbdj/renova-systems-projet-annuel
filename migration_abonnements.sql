-- =====================================================================
-- Migration : abonnements Pro à 3 niveaux (Premium / Plus / Pro)
-- À exécuter dans phpMyAdmin sur la base pa2026 AVANT de tester.
-- =====================================================================

-- 1) Plan souscrit par l'utilisateur : 'premium' | 'plus' | 'pro' | NULL
--    NULL = aucun abonnement payant. (est_premium reste le flag "a un abo actif")
ALTER TABLE `utilisateur`
  ADD COLUMN `plan_abo` VARCHAR(20) DEFAULT NULL;

-- (Optionnel) Aligner les utilisateurs déjà premium sur le plan 'premium'
UPDATE `utilisateur` SET `plan_abo` = 'premium' WHERE `est_premium` = 1 AND `plan_abo` IS NULL;

-- 2) Drapeau "annonce sponsorisée / boostée" (perk du plan Pro).
--    /!\ Si vous obtenez l'erreur "Duplicate column name 'is_sponsored'",
--        c'est que la colonne existe déjà : ignorez simplement cette ligne.
ALTER TABLE `annonce`
  ADD COLUMN `is_sponsored` TINYINT(1) NOT NULL DEFAULT 0;
