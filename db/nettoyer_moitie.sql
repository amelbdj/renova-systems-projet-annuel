SET FOREIGN_KEY_CHECKS=0;
-- On supprime la moitie haute des faux comptes (ids 107500..114999) et leurs annonces
DELETE FROM annonce     WHERE id_user >= 107500 AND id_user < 120000;
DELETE FROM utilisateur WHERE id      >= 107500 AND id      < 120000;
-- On rend le reste interessant : ~8% d'annonces sponsorisees
UPDATE annonce SET is_sponsored = 1        WHERE id_user >= 100000 AND id_user < 107500 AND (id % 12) = 0;
-- Un peu de variete dans les statuts de vente
UPDATE annonce SET statut_vente = 'VENDU'  WHERE id_user >= 100000 AND (id % 25) = 0;
UPDATE annonce SET statut_vente = 'EN BOX' WHERE id_user >= 100000 AND (id % 37) = 0;
SET FOREIGN_KEY_CHECKS=1;
