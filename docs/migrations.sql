-- ============================================================
-- Migrations a appliquer sur une base pa2026 EXISTANTE
-- (pour une base neuve, reimporter pa2026.sql qui est deja a jour)
-- Importer avec : mysql --default-character-set=utf8mb4 -uroot -proot pa2026 < docs/migrations.sql
-- A jouer une seule fois. Si une colonne existe deja, ignorer l'erreur correspondante.
-- ============================================================

-- --- Evenements / formations ---

-- Plan du cours (texte libre)
ALTER TABLE evenement ADD COLUMN plan_cours TEXT;

-- Types d'evenements acceptes (le formulaire propose ces 4)
ALTER TABLE evenement MODIFY type ENUM('evenement','formation','atelier','reunion') DEFAULT NULL;

-- Le PDF de formation est range dans ressource_pedagogique (lie a l'evenement)
ALTER TABLE ressource_pedagogique ADD COLUMN id_event INT DEFAULT NULL;
-- L'ancienne colonne pdf_url n'est plus utilisee :
-- ALTER TABLE evenement DROP COLUMN pdf_url;   -- si elle existe encore

-- --- Commandes / finances ---

-- Distinguer une commande d'annonce d'une commande d'evenement
ALTER TABLE `order` ADD COLUMN `type` VARCHAR(20) DEFAULT 'annonce';

-- --- Flux box (depot / retrait) ---

-- Statuts de vente utilises par le cycle achat -> depot -> retrait
ALTER TABLE annonce MODIFY statut_vente
  ENUM('EN VENTE','EN ATTENTE DEPOT','RESERVEE','EN BOX','EN ATTENTE DE RECUPERATION','RECUPERE','VENDU')
  DEFAULT 'EN VENTE';

-- --- Traductions (espace salarie : champs formation + stat forum) ---

DELETE FROM translations WHERE msg_key IN (
  'salarie.events.plan_label',
  'salarie.events.plan_ph',
  'salarie.events.plan_pdf_label',
  'salarie.events.resources_label',
  'salarie.events.resources_hint',
  'salarie.stat.reports_pending'
);

INSERT INTO translations (lang_code, msg_key, msg_value) VALUES
('fr','salarie.events.plan_label','Plan du cours'),
('en','salarie.events.plan_label','Course outline'),
('fr','salarie.events.plan_ph','Décrivez le déroulé / programme du cours…'),
('en','salarie.events.plan_ph','Describe the course outline / programme…'),
('fr','salarie.events.plan_pdf_label','Plan du cours (PDF, optionnel)'),
('en','salarie.events.plan_pdf_label','Course outline (PDF, optional)'),
('fr','salarie.events.resources_label','Ressources PDF (plusieurs possibles)'),
('en','salarie.events.resources_label','PDF resources (multiple allowed)'),
('fr','salarie.events.resources_hint','Réservé aux formations · fichiers PDF uniquement.'),
('en','salarie.events.resources_hint','Trainings only · PDF files only.'),
('fr','salarie.stat.reports_pending','Signalements à traiter'),
('en','salarie.stat.reports_pending','Reports to handle');
