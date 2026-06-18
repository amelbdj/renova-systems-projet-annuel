-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : mer. 15 avr. 2026 à 21:29
-- Version du serveur : 8.0.42
-- Version de PHP : 8.3.14

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `pa2026`
--

-- --------------------------------------------------------

--
-- Structure de la table `abonnement`
--

DROP TABLE IF EXISTS `abonnement`;
CREATE TABLE IF NOT EXISTS `abonnement` (
  `id_abonnement` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `id_plan` int DEFAULT NULL,
  `date_debut` date DEFAULT NULL,
  `date_fin` date DEFAULT NULL,
  `statut` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id_abonnement`),
  UNIQUE KEY `id_abonnement` (`id_abonnement`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `annonce`
--

DROP TABLE IF EXISTS `annonce`;
CREATE TABLE IF NOT EXISTS `annonce` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `id_categorie` int DEFAULT NULL,
  `titre` varchar(150) NOT NULL,
  `description` text,
  `type` varchar(20) DEFAULT NULL,
  `prix` decimal(10,2) DEFAULT '0.00',
  `statut_validation` enum('En attente','Validé','Rejeté') NOT NULL DEFAULT 'En attente',
  `code_postal` varchar(10) DEFAULT NULL,
  `ville` varchar(100) DEFAULT NULL,
  `projet_potentiel` text,
  `etat` enum('Neuf','Bon etat','Usage','Pour pieces') NOT NULL,
  `poids_kg` decimal(10,2) DEFAULT '0.00',
  `quantite` int DEFAULT '1',
  `is_sponsored` tinyint(1) DEFAULT '0',
  `commission_prelevee` decimal(10,2) DEFAULT '0.00',
  `id_box` int DEFAULT NULL,
  `photo_url` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `statut_vente` enum('EN VENTE','EN ATTENTE DEPOT','RESERVEE') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'EN VENTE',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_annonce` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `annonce`
--

INSERT INTO `annonce` (`id`, `id_user`, `id_categorie`, `titre`, `description`, `type`, `prix`, `statut_validation`, `code_postal`, `ville`, `projet_potentiel`, `etat`, `poids_kg`, `quantite`, `is_sponsored`, `commission_prelevee`, `id_box`, `photo_url`, `created_at`, `statut_vente`) VALUES
(1, 1, 3, 'Tondeuse à gazon électrique', 'Je vends ma tondeuse en très bon état, servie 3 fois.', 'Vente', 45.50, 'Validé', '93420', 'Villepinte', NULL, 'Neuf', 0.00, 1, 0, 0.00, 1, 'sdf', '2026-04-03 12:29:38', 'EN VENTE');

-- --------------------------------------------------------

--
-- Structure de la table `article_news`
--

DROP TABLE IF EXISTS `article_news`;
CREATE TABLE IF NOT EXISTS `article_news` (
  `id_article` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_salarie` int DEFAULT NULL,
  `titre` varchar(150) DEFAULT NULL,
  `slug` varchar(150) DEFAULT NULL,
  `contenu` text,
  `image_url` varchar(255) DEFAULT NULL,
  `type` varchar(50) DEFAULT NULL,
  `statut` enum('brouillon','valide','en attente','refuse') DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_article`),
  UNIQUE KEY `id_article` (`id_article`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `article_news`
--

INSERT INTO `article_news` (`id_article`, `id_salarie`, `titre`, `slug`, `contenu`, `image_url`, `type`, `statut`, `created_at`) VALUES
(1, 1, 'Bienvenue sur UpcycleConnect', NULL, 'tewgzdtfhzhffffffffff', NULL, 'Tutoriel', 'brouillon', '2026-04-13 17:19:04'),
(3, 1, '1 2 3 test micro', NULL, 'hehehehehe', NULL, 'Conseil pratique', 'en attente', '2026-04-13 20:15:59'),
(4, 1, 'why not', NULL, 'penser a bien dormir ', NULL, 'Conseil pratique', 'brouillon', '2026-04-13 20:16:18');

-- --------------------------------------------------------

--
-- Structure de la table `box_conteneur`
--

DROP TABLE IF EXISTS `box_conteneur`;
CREATE TABLE IF NOT EXISTS `box_conteneur` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `localisation` varchar(255) DEFAULT NULL,
  `etat` enum('MAINTENANCE','LIBRE','OCCUPE','') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'LIBRE',
  `type_materiau_accepte` varchar(50) DEFAULT NULL,
  `capacite` int DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_box` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `box_conteneur`
--

INSERT INTO `box_conteneur` (`id`, `localisation`, `etat`, `type_materiau_accepte`, `capacite`) VALUES
(1, 'Paris - Entrepôt Nord', 'LIBRE', 'Bois', 1),
(2, 'Paris - Entrepôt Nord', 'OCCUPE', 'Métal', 1),
(3, 'Lyon - Centre', 'MAINTENANCE', 'Plastique', 1),
(4, 'Lyon - Centre', 'LIBRE', 'Verre', 1),
(5, 'Marseille - Port', 'OCCUPE', 'Bois', 1),
(6, 'Paris - Entrepôt Nord', 'LIBRE', 'Bois', 1),
(7, 'Paris - Entrepôt Nord', 'OCCUPE', 'Métal', 1),
(8, 'Paris - Entrepôt Nord', 'LIBRE', 'Plastique', 1),
(9, 'Paris - Entrepôt Nord', 'MAINTENANCE', 'Électronique', 1),
(10, 'Paris - Entrepôt Nord', 'LIBRE', 'Bois', 1),
(11, 'Lyon - Centre de Tri', 'OCCUPE', 'Carton', 1),
(12, 'Lyon - Centre de Tri', 'LIBRE', 'Verre', 1),
(13, 'Lyon - Centre de Tri', 'LIBRE', 'Métal', 1),
(14, 'Lyon - Centre de Tri', 'MAINTENANCE', 'Bois', 1),
(15, 'Lyon - Centre de Tri', 'OCCUPE', 'Textile', 1),
(16, 'Bordeaux - Zone Eco', 'LIBRE', 'Plastique', 1),
(17, 'Bordeaux - Zone Eco', 'OCCUPE', 'Électronique', 1),
(18, 'Bordeaux - Zone Eco', 'LIBRE', 'Métal', 1),
(19, 'Bordeaux - Zone Eco', 'LIBRE', 'Bois', 1),
(20, 'Bordeaux - Zone Eco', 'MAINTENANCE', 'Papier', 1),
(21, 'eee', 'LIBRE', 'plastique', 22);

-- --------------------------------------------------------

--
-- Structure de la table `campagne_pub`
--

DROP TABLE IF EXISTS `campagne_pub`;
CREATE TABLE IF NOT EXISTS `campagne_pub` (
  `id_campagne` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_partenaire` int DEFAULT NULL,
  `titre` varchar(150) DEFAULT NULL,
  `budget` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`id_campagne`),
  UNIQUE KEY `id_campagne` (`id_campagne`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `categorie`
--

DROP TABLE IF EXISTS `categorie`;
CREATE TABLE IF NOT EXISTS `categorie` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `libelle` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_categorie` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `categorie`
--

INSERT INTO `categorie` (`id`, `libelle`) VALUES
(1, 'Textile'),
(2, 'Bois'),
(3, 'Plastique'),
(4, 'Métal');

-- --------------------------------------------------------

--
-- Structure de la table `commande`
--

DROP TABLE IF EXISTS `commande`;
CREATE TABLE IF NOT EXISTS `commande` (
  `id_commande` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_acheteur` int DEFAULT NULL,
  `id_annonce` int DEFAULT NULL,
  `montant_total` decimal(10,2) DEFAULT NULL,
  `commission` decimal(10,2) DEFAULT NULL,
  `date_commande` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_commande`),
  UNIQUE KEY `id_commande` (`id_commande`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `depot_box`
--

DROP TABLE IF EXISTS `depot_box`;
CREATE TABLE IF NOT EXISTS `depot_box` (
  `id_depot` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_annonce` int DEFAULT NULL,
  `id_box` int DEFAULT NULL,
  `code_ouverture` varchar(10) DEFAULT NULL,
  `code_barres_pro` varchar(50) DEFAULT NULL,
  `date_depot` date DEFAULT NULL,
  `date_retrait` date DEFAULT NULL,
  PRIMARY KEY (`id_depot`),
  UNIQUE KEY `id_depot` (`id_depot`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `dictionnaire`
--

DROP TABLE IF EXISTS `dictionnaire`;
CREATE TABLE IF NOT EXISTS `dictionnaire` (
  `id_trad` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `langue` varchar(5) DEFAULT NULL,
  `cle` varchar(100) DEFAULT NULL,
  `valeur` text,
  PRIMARY KEY (`id_trad`),
  UNIQUE KEY `id_trad` (`id_trad`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `document`
--

DROP TABLE IF EXISTS `document`;
CREATE TABLE IF NOT EXISTS `document` (
  `id_document` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `type_doc` varchar(50) DEFAULT NULL,
  `url_pdf` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id_document`),
  UNIQUE KEY `id_document` (`id_document`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `documents_legaux`
--

DROP TABLE IF EXISTS `documents_legaux`;
CREATE TABLE IF NOT EXISTS `documents_legaux` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` bigint UNSIGNED NOT NULL,
  `type_document` enum('KBIS','SIREN','DIPLOME','PIECE_IDENTITE') COLLATE utf8mb4_unicode_ci NOT NULL,
  `chemin_fichier` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `statut_document` enum('En attente','Validé','Rejeté') COLLATE utf8mb4_unicode_ci DEFAULT 'En attente',
  `date_soumission` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Déchargement des données de la table `documents_legaux`
--

INSERT INTO `documents_legaux` (`id`, `user_id`, `type_document`, `chemin_fichier`, `statut_document`, `date_soumission`) VALUES
(7, 1, 'KBIS', 'uploads/kbis_entreprise_1.pdf', 'En attente', '2026-04-02 20:43:16'),
(8, 6, 'KBIS', 'uploads/cni_recto_verso.jpg', 'En attente', '2026-04-02 20:43:16'),
(9, 7, 'KBIS', 'uploads/justificatif_pro.png', 'En attente', '2026-04-02 20:43:16'),
(10, 9, 'KBIS', 'uploads/justificatif_pfffro.png', 'En attente', '2026-04-02 20:43:16'),
(11, 11, 'KBIS', 'uploads/justificatif_pffro.png', 'En attente', '2026-04-02 20:43:16'),
(12, 11, 'KBIS', 'uploads/justificatif_fpro.png', 'En attente', '2026-04-02 20:43:16');

-- --------------------------------------------------------

--
-- Structure de la table `evenement`
--

DROP TABLE IF EXISTS `evenement`;
CREATE TABLE IF NOT EXISTS `evenement` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_salarie` int DEFAULT NULL,
  `titre` varchar(150) DEFAULT NULL,
  `description` text,
  `date_debut` timestamp NULL DEFAULT NULL,
  `date_fin` timestamp NULL DEFAULT NULL,
  `nb_places` int DEFAULT NULL,
  `statut_validation` varchar(20) DEFAULT NULL,
  `format` varchar(20) DEFAULT NULL,
  `type` enum('evenement','formation','atelier','reunion') DEFAULT NULL,
  `lieu` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '0',
  `prix` int DEFAULT '0',
  `image_url` varchar(255) DEFAULT NULL,
  `pdf_url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_event` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `evenement`
--

INSERT INTO `evenement` (`id`, `id_salarie`, `titre`, `description`, `date_debut`, `date_fin`, `nb_places`, `statut_validation`, `format`, `type`, `lieu`) VALUES
(1, 1, 'Atelier Couture Débutant', 'Apprenez à rapiécer vos vêtements au lieu de les jeter.', '2026-03-20 13:00:00', '2026-03-20 16:00:00', 8, 'en attente', 'Présentiel', 'evenement', '0'),
(2, 1, 'dsfgds', 'sdFsD', '2026-04-09 08:00:00', '2026-04-09 11:00:00', 0, 'en attente', 'Présentiel', 'formation', 'sdfdf'),
(3, 1, 'sdgsd', 'sdgfsd', '2026-04-16 08:00:00', '2026-04-16 11:00:00', 0, 'en attente', 'Présentiel', 'formation', 'sdf');

-- --------------------------------------------------------

--
-- Structure de la table `historique_conteneurs`
--

DROP TABLE IF EXISTS `historique_conteneurs`;
CREATE TABLE IF NOT EXISTS `historique_conteneurs` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `conteneur_id` bigint UNSIGNED NOT NULL,
  `annonce_id` bigint UNSIGNED NOT NULL,
  `particulier_id` bigint UNSIGNED NOT NULL,
  `professionnel_id` bigint UNSIGNED DEFAULT NULL,
  `code_ouverture` varchar(10) NOT NULL,
  `code_barre_recuperation` varchar(50) NOT NULL,
  `date_reservation` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `date_depot_effective` timestamp NULL DEFAULT NULL,
  `date_retrait_effective` timestamp NULL DEFAULT NULL,
  `etat_objet_depot` varchar(50) DEFAULT NULL,
  `etat` enum('SUPPRIME','EN COURS','RECUPERE') DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `code_barre_recuperation` (`code_barre_recuperation`),
  KEY `fk_conteneur` (`conteneur_id`),
  KEY `fk_annonce` (`annonce_id`),
  KEY `fk_particulier` (`particulier_id`),
  KEY `fk_professionnel` (`professionnel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `inscription`
--

DROP TABLE IF EXISTS `inscription`;
CREATE TABLE IF NOT EXISTS `inscription` (
  `id_inscription` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `id_event` int DEFAULT NULL,
  `date_inscrip` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_inscription`),
  UNIQUE KEY `id_inscription` (`id_inscription`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `languages`
--

DROP TABLE IF EXISTS `languages`;
CREATE TABLE IF NOT EXISTS `languages` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(5) NOT NULL,
  `name` varchar(50) NOT NULL,
  `is_default` tinyint(1) DEFAULT '0',
  `is_active` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `languages`
--

INSERT INTO `languages` (`id`, `code`, `name`, `is_default`, `is_active`) VALUES
(1, 'fr', 'Français', 0, 1),
(2, 'en', 'English', 0, 1);

-- --------------------------------------------------------

--
-- Structure de la table `log_connexion`
--

DROP TABLE IF EXISTS `log_connexion`;
CREATE TABLE IF NOT EXISTS `log_connexion` (
  `id_log` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `ip` varchar(45) DEFAULT NULL,
  `date_connexion` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_log`),
  UNIQUE KEY `id_log` (`id_log`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `message_forum`
--

DROP TABLE IF EXISTS `message_forum`;
CREATE TABLE IF NOT EXISTS `message_forum` (
  `id_message` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_topic` int DEFAULT NULL,
  `id_user` int DEFAULT NULL,
  `contenu` text,
  `est_modere` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id_message`),
  UNIQUE KEY `id_message` (`id_message`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `notification`
--

DROP TABLE IF EXISTS `notification`;
CREATE TABLE IF NOT EXISTS `notification` (
  `id_notif` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `contenu` text,
  `est_lu` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id_notif`),
  UNIQUE KEY `id_notif` (`id_notif`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `paiement`
--

DROP TABLE IF EXISTS `paiement`;
CREATE TABLE IF NOT EXISTS `paiement` (
  `id_paiement` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_commande` int DEFAULT NULL,
  `stripe_id` varchar(255) DEFAULT NULL,
  `statut` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id_paiement`),
  UNIQUE KEY `id_paiement` (`id_paiement`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `partenaire`
--

DROP TABLE IF EXISTS `partenaire`;
CREATE TABLE IF NOT EXISTS `partenaire` (
  `id_partenaire` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `nom_entreprise` varchar(150) DEFAULT NULL,
  `type_partenariat` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id_partenaire`),
  UNIQUE KEY `id_partenaire` (`id_partenaire`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `plan_abo`
--

DROP TABLE IF EXISTS `plan_abo`;
CREATE TABLE IF NOT EXISTS `plan_abo` (
  `id_plan` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `nom` varchar(100) DEFAULT NULL,
  `prix` decimal(10,2) DEFAULT NULL,
  `duree_mois` int DEFAULT NULL,
  PRIMARY KEY (`id_plan`),
  UNIQUE KEY `id_plan` (`id_plan`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `projet_pro`
--

DROP TABLE IF EXISTS `projet_pro`;
CREATE TABLE IF NOT EXISTS `projet_pro` (
  `id_projet` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `titre` varchar(150) DEFAULT NULL,
  `desc_etapes` text,
  `url_photo_avant` varchar(255) DEFAULT NULL,
  `url_photo_apres` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id_projet`),
  UNIQUE KEY `id_projet` (`id_projet`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `ressource_pedagogique`
--

DROP TABLE IF EXISTS `ressource_pedagogique`;
CREATE TABLE IF NOT EXISTS `ressource_pedagogique` (
  `id_ressource` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_salarie` int DEFAULT NULL,
  `titre` varchar(150) DEFAULT NULL,
  `url_fichier` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id_ressource`),
  UNIQUE KEY `id_ressource` (`id_ressource`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `topic_forum`
--

DROP TABLE IF EXISTS `topic_forum`;
CREATE TABLE IF NOT EXISTS `topic_forum` (
  `id_topic` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `titre` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id_topic`),
  UNIQUE KEY `id_topic` (`id_topic`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `translations`
--

DROP TABLE IF EXISTS `translations`;
CREATE TABLE IF NOT EXISTS `translations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `lang_code` varchar(5) DEFAULT NULL,
  `msg_key` varchar(100) DEFAULT NULL,
  `msg_value` text,
  PRIMARY KEY (`id`),
  KEY `lang_code` (`lang_code`)
) ENGINE=InnoDB AUTO_INCREMENT=314 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `translations`
--

INSERT INTO `translations` (`id`, `lang_code`, `msg_key`, `msg_value`) VALUES
(1, 'fr', 'backend.error.invalid_id', 'ID invalide'),
(2, 'en', 'backend.error.invalid_id', 'Invalid ID'),
(3, 'fr', 'backend.error.json_decode', 'Impossible de décoder le JSON'),
(4, 'en', 'backend.error.json_decode', 'Unable to decode JSON payload'),
(5, 'fr', 'backend.error.internal', 'Erreur interne du serveur'),
(6, 'en', 'backend.error.internal', 'Internal server error'),
(7, 'fr', 'backend.error.email_used', 'Cet email est déjà utilisé'),
(8, 'en', 'backend.error.email_used', 'This email is already in use'),
(9, 'fr', 'backend.error.category_used', 'Ce libellé de catégorie existe déjà'),
(10, 'en', 'backend.error.category_used', 'This category label already exists'),
(11, 'fr', 'backend.error.not_found', 'Ressource introuvable'),
(12, 'en', 'backend.error.not_found', 'Resource not found'),
(13, 'fr', 'backend.error.invalid_id', 'ID invalide'),
(14, 'en', 'backend.error.invalid_id', 'Invalid ID'),
(15, 'fr', 'backend.error.json_decode', 'Impossible de décoder le JSON'),
(16, 'en', 'backend.error.json_decode', 'Unable to decode JSON payload'),
(17, 'fr', 'backend.error.internal', 'Erreur interne du serveur'),
(18, 'en', 'backend.error.internal', 'Internal server error'),
(19, 'fr', 'backend.error.email_used', 'Cet email est déjà utilisé'),
(20, 'en', 'backend.error.email_used', 'This email is already in use'),
(21, 'fr', 'backend.error.category_used', 'Ce libellé de catégorie existe déjà'),
(22, 'en', 'backend.error.category_used', 'This category label already exists'),
(23, 'fr', 'backend.error.not_found', 'Ressource introuvable'),
(24, 'en', 'backend.error.not_found', 'Resource not found'),
(25, 'fr', 'backend.auth.invalid_creds', 'Email ou mot de passe incorrect'),
(26, 'en', 'backend.auth.invalid_creds', 'Invalid email or password'),
(27, 'fr', 'backend.auth.no_token', 'Accès refusé : Token manquant'),
(28, 'en', 'backend.auth.no_token', 'Access denied: Missing token'),
(29, 'fr', 'backend.auth.invalid_token', 'Accès refusé : Token invalide ou expiré'),
(30, 'en', 'backend.auth.invalid_token', 'Access denied: Invalid or expired token'),
(31, 'fr', 'backend.success.validated', 'Validé avec succès'),
(32, 'en', 'backend.success.validated', 'Successfully validated'),
(33, 'fr', 'backend.success.refused', 'Refusé avec succès'),
(34, 'en', 'backend.success.refused', 'Successfully refused'),
(35, 'fr', 'backend.auth.invalid_creds', 'Email ou mot de passe incorrect'),
(36, 'en', 'backend.auth.invalid_creds', 'Invalid email or password'),
(37, 'fr', 'backend.auth.no_token', 'Accès refusé : Token manquant'),
(38, 'en', 'backend.auth.no_token', 'Access denied: Missing token'),
(39, 'fr', 'backend.auth.invalid_token', 'Accès refusé : Token invalide ou expiré'),
(40, 'en', 'backend.auth.invalid_token', 'Access denied: Invalid or expired token'),
(41, 'fr', 'backend.success.validated', 'Validé avec succès'),
(42, 'en', 'backend.success.validated', 'Successfully validated'),
(43, 'fr', 'backend.success.refused', 'Refusé avec succès'),
(44, 'en', 'backend.success.refused', 'Successfully refused'),
(45, 'fr', 'backoffice.hero.title', 'Centre de commande ReNova'),
(46, 'en', 'backoffice.hero.title', 'ReNova Command Center'),
(47, 'fr', 'backoffice.hero.desc', 'Gérez utilisateurs, validations, finances et logistique depuis ce panneau centralisé.'),
(48, 'en', 'backoffice.hero.desc', 'Manage users, validations, finances, and logistics from this centralized panel.'),
(49, 'fr', 'backoffice.kpi.users', 'Utilisateurs inscrits'),
(50, 'en', 'backoffice.kpi.users', 'Registered Users'),
(51, 'fr', 'backoffice.kpi.pending', 'Validations en attente'),
(52, 'en', 'backoffice.kpi.pending', 'Pending Validations'),
(53, 'fr', 'backoffice.tabs.all', 'Tout'),
(54, 'en', 'backoffice.tabs.all', 'All'),
(55, 'fr', 'backoffice.tabs.ads', 'Annonces'),
(56, 'en', 'backoffice.tabs.ads', 'Ads'),
(57, 'fr', 'backoffice.tabs.events', 'Événements'),
(58, 'en', 'backoffice.tabs.events', 'Events'),
(59, 'fr', 'backoffice.hero.title', 'Centre de commande ReNova'),
(60, 'en', 'backoffice.hero.title', 'ReNova Command Center'),
(61, 'fr', 'backoffice.hero.desc', 'Gérez utilisateurs, validations, finances et logistique depuis ce panneau centralisé.'),
(62, 'en', 'backoffice.hero.desc', 'Manage users, validations, finances, and logistics from this centralized panel.'),
(63, 'fr', 'backoffice.kpi.users', 'Utilisateurs inscrits'),
(64, 'en', 'backoffice.kpi.users', 'Registered Users'),
(65, 'fr', 'backoffice.kpi.pending', 'Validations en attente'),
(66, 'en', 'backoffice.kpi.pending', 'Pending Validations'),
(67, 'fr', 'backoffice.tabs.all', 'Tout'),
(68, 'en', 'backoffice.tabs.all', 'All'),
(69, 'fr', 'backoffice.tabs.ads', 'Annonces'),
(70, 'en', 'backoffice.tabs.ads', 'Ads'),
(71, 'fr', 'backoffice.tabs.events', 'Événements'),
(72, 'en', 'backoffice.tabs.events', 'Events'),
(73, 'fr', 'backoffice.btn.approve', '✓ Approuver'),
(74, 'en', 'backoffice.btn.approve', '✓ Approve'),
(75, 'fr', 'backoffice.btn.refuse', '✕ Refuser'),
(76, 'en', 'backoffice.btn.refuse', '✕ Refuse'),
(77, 'fr', 'backoffice.btn.create_user', '＋ Créer un compte'),
(78, 'en', 'backoffice.btn.create_user', '＋ Create Account'),
(79, 'fr', 'backoffice.btn.create_cat', '＋ Créer une catégorie'),
(80, 'en', 'backoffice.btn.create_cat', '＋ Create Category'),
(81, 'fr', 'backoffice.search.placeholder', 'Rechercher un utilisateur…'),
(82, 'en', 'backoffice.search.placeholder', 'Search for a user...'),
(83, 'fr', 'backoffice.btn.approve', '✓ Approuver'),
(84, 'en', 'backoffice.btn.approve', '✓ Approve'),
(85, 'fr', 'backoffice.btn.refuse', '✕ Refuser'),
(86, 'en', 'backoffice.btn.refuse', '✕ Refuse'),
(87, 'fr', 'backoffice.btn.create_user', '＋ Créer un compte'),
(88, 'en', 'backoffice.btn.create_user', '＋ Create Account'),
(89, 'fr', 'backoffice.btn.create_cat', '＋ Créer une catégorie'),
(90, 'en', 'backoffice.btn.create_cat', '＋ Create Category'),
(91, 'fr', 'backoffice.search.placeholder', 'Rechercher un utilisateur…'),
(92, 'en', 'backoffice.search.placeholder', 'Search for a user...'),
(93, 'fr', 'backoffice.hero.eyebrow', 'Back Office'),
(94, 'en', 'backoffice.hero.eyebrow', 'Admin Panel'),
(95, 'fr', 'backoffice.btn.process_validations', 'Traiter les validations'),
(96, 'en', 'backoffice.btn.process_validations', 'Process Validations'),
(97, 'fr', 'backoffice.btn.manage_users', 'Gérer les utilisateurs'),
(98, 'en', 'backoffice.btn.manage_users', 'Manage Users'),
(99, 'fr', 'backoffice.kpi.pending_badge', 'En attente'),
(100, 'en', 'backoffice.kpi.pending_badge', 'Pending'),
(101, 'fr', 'backoffice.kpi.revenue', 'Revenus ce mois (Stripe)'),
(102, 'en', 'backoffice.kpi.revenue', 'Revenue this month (Stripe)'),
(103, 'fr', 'backoffice.kpi.todo', 'À traiter'),
(104, 'en', 'backoffice.kpi.todo', 'To Do'),
(105, 'fr', 'backoffice.kpi.reports', 'Signalements critiques'),
(106, 'en', 'backoffice.kpi.reports', 'Critical Reports'),
(107, 'fr', 'backoffice.users.title', 'Gestion des Utilisateurs'),
(108, 'en', 'backoffice.users.title', 'User Management'),
(109, 'fr', 'backoffice.role.all', 'Tous les rôles'),
(110, 'en', 'backoffice.role.all', 'All Roles'),
(111, 'fr', 'backoffice.role.user', 'Particulier'),
(112, 'en', 'backoffice.role.user', 'Individual'),
(113, 'fr', 'backoffice.role.pro', 'Professionnel'),
(114, 'en', 'backoffice.role.pro', 'Professional'),
(115, 'fr', 'backoffice.role.admin', 'Admin'),
(116, 'en', 'backoffice.role.admin', 'Admin'),
(117, 'fr', 'backoffice.role.staff', 'Salarié'),
(118, 'en', 'backoffice.role.staff', 'Staff'),
(119, 'fr', 'backoffice.validations.title', 'Validations'),
(120, 'en', 'backoffice.validations.title', 'Validations'),
(121, 'fr', 'backoffice.tabs.pro', 'Comptes Pro'),
(122, 'en', 'backoffice.tabs.pro', 'Pro Accounts'),
(123, 'fr', 'backoffice.tabs.content', 'Contenus'),
(124, 'en', 'backoffice.tabs.content', 'Content'),
(125, 'fr', 'backoffice.categories.title', 'Catégories'),
(126, 'en', 'backoffice.categories.title', 'Categories'),
(127, 'fr', 'backoffice.modal.edit_user.title', 'Modifier le compte'),
(128, 'en', 'backoffice.modal.edit_user.title', 'Edit Account'),
(129, 'fr', 'backoffice.modal.edit_user.desc', 'Modifiez les informations et les droits de l\'utilisateur.'),
(130, 'en', 'backoffice.modal.edit_user.desc', 'Modify user information and permissions.'),
(131, 'fr', 'backoffice.form.firstname', 'Prénom'),
(132, 'en', 'backoffice.form.firstname', 'First Name'),
(133, 'fr', 'backoffice.form.lastname', 'Nom'),
(134, 'en', 'backoffice.form.lastname', 'Last Name'),
(135, 'fr', 'backoffice.form.email', 'Email'),
(136, 'en', 'backoffice.form.email', 'Email'),
(137, 'fr', 'backoffice.form.role', 'Rôle'),
(138, 'en', 'backoffice.form.role', 'Role'),
(139, 'fr', 'backoffice.form.status', 'Statut'),
(140, 'en', 'backoffice.form.status', 'Status'),
(141, 'fr', 'backoffice.form.password', 'Mot de passe'),
(142, 'en', 'backoffice.form.password', 'Password'),
(143, 'fr', 'backoffice.form.initial_status', 'Statut initial'),
(144, 'en', 'backoffice.form.initial_status', 'Initial Status'),
(145, 'fr', 'backoffice.form.label', 'Libellé'),
(146, 'en', 'backoffice.form.label', 'Label'),
(147, 'fr', 'backoffice.status.active', 'Actif'),
(148, 'en', 'backoffice.status.active', 'Active'),
(149, 'fr', 'backoffice.status.suspended', 'Suspendu'),
(150, 'en', 'backoffice.status.suspended', 'Suspended'),
(151, 'fr', 'backoffice.status.banned', 'Banni'),
(152, 'en', 'backoffice.status.banned', 'Banned'),
(153, 'fr', 'backoffice.btn.save_changes', 'Enregistrer les modifications'),
(154, 'en', 'backoffice.btn.save_changes', 'Save Changes'),
(155, 'fr', 'backoffice.modal.new_user.title', 'Créer un compte'),
(156, 'en', 'backoffice.modal.new_user.title', 'Create Account'),
(157, 'fr', 'backoffice.modal.new_user.desc', 'Créer manuellement un compte sur la plateforme.'),
(158, 'en', 'backoffice.modal.new_user.desc', 'Manually create an account on the platform.'),
(159, 'fr', 'backoffice.btn.submit_create_user', 'Créer le compte ✓'),
(160, 'en', 'backoffice.btn.submit_create_user', 'Create Account ✓'),
(161, 'fr', 'backoffice.modal.new_cat.title', 'Créer une catégorie'),
(162, 'en', 'backoffice.modal.new_cat.title', 'Create a Category'),
(163, 'fr', 'backoffice.modal.new_cat.desc', 'Créer manuellement une catégorie d\'annonces sur la plateforme.'),
(164, 'en', 'backoffice.modal.new_cat.desc', 'Manually create an ad category on the platform.'),
(165, 'fr', 'backoffice.btn.submit_create_cat', 'Créer la catégorie ✓'),
(166, 'en', 'backoffice.btn.submit_create_cat', 'Create Category ✓'),
(167, 'fr', 'backoffice.hero.eyebrow', 'Back Office'),
(168, 'en', 'backoffice.hero.eyebrow', 'Admin Panel'),
(169, 'fr', 'backoffice.btn.process_validations', 'Traiter les validations'),
(170, 'en', 'backoffice.btn.process_validations', 'Process Validations'),
(171, 'fr', 'backoffice.btn.manage_users', 'Gérer les utilisateurs'),
(172, 'en', 'backoffice.btn.manage_users', 'Manage Users'),
(173, 'fr', 'backoffice.kpi.pending_badge', 'En attente'),
(174, 'en', 'backoffice.kpi.pending_badge', 'Pending'),
(175, 'fr', 'backoffice.kpi.revenue', 'Revenus ce mois (Stripe)'),
(176, 'en', 'backoffice.kpi.revenue', 'Revenue this month (Stripe)'),
(177, 'fr', 'backoffice.kpi.todo', 'À traiter'),
(178, 'en', 'backoffice.kpi.todo', 'To Do'),
(179, 'fr', 'backoffice.kpi.reports', 'Signalements critiques'),
(180, 'en', 'backoffice.kpi.reports', 'Critical Reports'),
(181, 'fr', 'backoffice.users.title', 'Gestion des Utilisateurs'),
(182, 'en', 'backoffice.users.title', 'User Management'),
(183, 'fr', 'backoffice.role.all', 'Tous les rôles'),
(184, 'en', 'backoffice.role.all', 'All Roles'),
(185, 'fr', 'backoffice.role.user', 'Particulier'),
(186, 'en', 'backoffice.role.user', 'Individual'),
(187, 'fr', 'backoffice.role.pro', 'Professionnel'),
(188, 'en', 'backoffice.role.pro', 'Professional'),
(189, 'fr', 'backoffice.role.admin', 'Admin'),
(190, 'en', 'backoffice.role.admin', 'Admin'),
(191, 'fr', 'backoffice.role.staff', 'Salarié'),
(192, 'en', 'backoffice.role.staff', 'Staff'),
(193, 'fr', 'backoffice.validations.title', 'Validations'),
(194, 'en', 'backoffice.validations.title', 'Validations'),
(195, 'fr', 'backoffice.tabs.pro', 'Comptes Pro'),
(196, 'en', 'backoffice.tabs.pro', 'Pro Accounts'),
(197, 'fr', 'backoffice.tabs.content', 'Contenus'),
(198, 'en', 'backoffice.tabs.content', 'Content'),
(199, 'fr', 'backoffice.categories.title', 'Catégories'),
(200, 'en', 'backoffice.categories.title', 'Categories'),
(201, 'fr', 'backoffice.modal.edit_user.title', 'Modifier le compte'),
(202, 'en', 'backoffice.modal.edit_user.title', 'Edit Account'),
(203, 'fr', 'backoffice.modal.edit_user.desc', 'Modifiez les informations et les droits de l\'utilisateur.'),
(204, 'en', 'backoffice.modal.edit_user.desc', 'Modify user information and permissions.'),
(205, 'fr', 'backoffice.form.firstname', 'Prénom'),
(206, 'en', 'backoffice.form.firstname', 'First Name'),
(207, 'fr', 'backoffice.form.lastname', 'Nom'),
(208, 'en', 'backoffice.form.lastname', 'Last Name'),
(209, 'fr', 'backoffice.form.email', 'Email'),
(210, 'en', 'backoffice.form.email', 'Email'),
(211, 'fr', 'backoffice.form.role', 'Rôle'),
(212, 'en', 'backoffice.form.role', 'Role'),
(213, 'fr', 'backoffice.form.status', 'Statut'),
(214, 'en', 'backoffice.form.status', 'Status'),
(215, 'fr', 'backoffice.form.password', 'Mot de passe'),
(216, 'en', 'backoffice.form.password', 'Password'),
(217, 'fr', 'backoffice.form.initial_status', 'Statut initial'),
(218, 'en', 'backoffice.form.initial_status', 'Initial Status'),
(219, 'fr', 'backoffice.form.label', 'Libellé'),
(220, 'en', 'backoffice.form.label', 'Label'),
(221, 'fr', 'backoffice.status.active', 'Actif'),
(222, 'en', 'backoffice.status.active', 'Active'),
(223, 'fr', 'backoffice.status.suspended', 'Suspendu'),
(224, 'en', 'backoffice.status.suspended', 'Suspended'),
(225, 'fr', 'backoffice.status.banned', 'Banni'),
(226, 'en', 'backoffice.status.banned', 'Banned'),
(227, 'fr', 'backoffice.btn.save_changes', 'Enregistrer les modifications'),
(228, 'en', 'backoffice.btn.save_changes', 'Save Changes'),
(229, 'fr', 'backoffice.modal.new_user.title', 'Créer un compte'),
(230, 'en', 'backoffice.modal.new_user.title', 'Create Account'),
(231, 'fr', 'backoffice.modal.new_user.desc', 'Créer manuellement un compte sur la plateforme.'),
(232, 'en', 'backoffice.modal.new_user.desc', 'Manually create an account on the platform.'),
(233, 'fr', 'backoffice.btn.submit_create_user', 'Créer le compte ✓'),
(234, 'en', 'backoffice.btn.submit_create_user', 'Create Account ✓'),
(235, 'fr', 'backoffice.modal.new_cat.title', 'Créer une catégorie'),
(236, 'en', 'backoffice.modal.new_cat.title', 'Create a Category'),
(237, 'fr', 'backoffice.modal.new_cat.desc', 'Créer manuellement une catégorie d\'annonces sur la plateforme.'),
(238, 'en', 'backoffice.modal.new_cat.desc', 'Manually create an ad category on the platform.'),
(239, 'fr', 'backoffice.btn.submit_create_cat', 'Créer la catégorie ✓'),
(240, 'en', 'backoffice.btn.submit_create_cat', 'Create Category ✓'),
(241, 'fr', 'backoffice.ads.ad', 'Annonce'),
(242, 'en', 'backoffice.ads.ad', 'Ad'),
(243, 'fr', 'backoffice.ads.published_on', 'Publiée le'),
(244, 'en', 'backoffice.ads.published_on', 'Published on'),
(245, 'fr', 'backoffice.ads.ad_tag', 'Annonce'),
(246, 'en', 'backoffice.ads.ad_tag', 'Listing'),
(247, 'fr', 'backoffice.ads.no_ads', 'Aucune annonce en attente.'),
(248, 'en', 'backoffice.ads.no_ads', 'No pending ads.'),
(249, 'fr', 'backoffice.events.event', 'Event'),
(250, 'en', 'backoffice.events.event', 'Event'),
(251, 'fr', 'backoffice.events.organizer', 'Organisateur'),
(252, 'en', 'backoffice.events.organizer', 'Organizer'),
(253, 'fr', 'backoffice.events.seats', 'Places :'),
(254, 'en', 'backoffice.events.seats', 'Seats :'),
(255, 'fr', 'backoffice.events.on_date', 'Le'),
(256, 'en', 'backoffice.events.on_date', 'On'),
(257, 'fr', 'backoffice.events.event_tag', 'Événement'),
(258, 'en', 'backoffice.events.event_tag', 'Event'),
(259, 'fr', 'backoffice.events.no_events', 'Aucun événement en attente.'),
(260, 'en', 'backoffice.events.no_events', 'No pending events.'),
(261, 'fr', 'backoffice.common.loading', 'Chargement'),
(262, 'en', 'backoffice.common.loading', 'Loading'),
(263, 'fr', 'backoffice.common.in_dev', 'En cours de dev'),
(264, 'en', 'backoffice.common.in_dev', 'Work in progress'),
(265, 'fr', 'backoffice.users.accounts_badge', 'comptes'),
(266, 'en', 'backoffice.users.accounts_badge', 'accounts'),
(267, 'fr', 'backoffice.users.th_user', 'Utilisateur'),
(268, 'en', 'backoffice.users.th_user', 'User'),
(269, 'fr', 'backoffice.users.th_score', 'Score'),
(270, 'en', 'backoffice.users.th_score', 'Score'),
(271, 'fr', 'backoffice.users.th_actions', 'Actions'),
(272, 'en', 'backoffice.users.th_actions', 'Actions'),
(273, 'fr', 'backoffice.users.no_users_found', 'Aucun utilisateur trouvé.'),
(274, 'en', 'backoffice.users.no_users_found', 'No users found.'),
(275, 'fr', 'backoffice.users.search_error', 'Erreur lors de la recherche.'),
(276, 'en', 'backoffice.users.search_error', 'Error during search.'),
(277, 'fr', 'backoffice.users.confirm_delete', 'Supprimer cet utilisateur ?'),
(278, 'en', 'backoffice.users.confirm_delete', 'Delete this user?'),
(279, 'fr', 'backoffice.users.alert_empty', 'Veuillez remplir tous les champs.'),
(280, 'en', 'backoffice.users.alert_empty', 'Please fill in all fields.'),
(281, 'fr', 'backoffice.users.success_update', 'Utilisateur mis à jour avec succès.'),
(282, 'en', 'backoffice.users.success_update', 'User successfully updated.'),
(298, 'fr', 'backoffice.users.th_document', 'Document'),
(299, 'fr', 'backoffice.users.th_status', 'Statut'),
(300, 'fr', 'backoffice.users.doc_view', 'Voir'),
(301, 'fr', 'backoffice.users.doc_none', 'Aucun'),
(302, 'fr', 'backoffice.users.status_approved', 'Validé'),
(303, 'fr', 'backoffice.users.status_rejected', 'Rejeté'),
(304, 'fr', 'backoffice.users.action_approve', 'Approuver'),
(305, 'fr', 'backoffice.users.action_refuse', 'Refuser'),
(306, 'en', 'backoffice.users.th_document', 'Document'),
(307, 'en', 'backoffice.users.th_status', 'Status'),
(308, 'en', 'backoffice.users.doc_view', 'View'),
(309, 'en', 'backoffice.users.doc_none', 'None'),
(310, 'en', 'backoffice.users.status_approved', 'Approved'),
(311, 'en', 'backoffice.users.status_rejected', 'Rejected'),
(312, 'en', 'backoffice.users.action_approve', 'Approve'),
(313, 'en', 'backoffice.users.action_refuse', 'Reject');

-- --------------------------------------------------------

--
-- Structure de la table `upcycling_score`
--

DROP TABLE IF EXISTS `upcycling_score`;
CREATE TABLE IF NOT EXISTS `upcycling_score` (
  `id_score` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `points_totaux` int DEFAULT '0',
  `poids_evite_kg` decimal(10,2) DEFAULT '0.00',
  PRIMARY KEY (`id_score`),
  UNIQUE KEY `id_score` (`id_score`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `utilisateur`
--

DROP TABLE IF EXISTS `utilisateur`;
CREATE TABLE IF NOT EXISTS `utilisateur` (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `role` varchar(50) NOT NULL,
  `nom` varchar(100) DEFAULT NULL,
  `prenom` varchar(100) DEFAULT NULL,
  `email` varchar(150) NOT NULL,
  `mot_de_passe` varchar(255) NOT NULL,
  `tutoriel_vu` tinyint(1) DEFAULT '0',
  `type_statut` varchar(50) DEFAULT NULL,
  `nom_entreprise` varchar(150) DEFAULT NULL,
  `siret` varchar(14) DEFAULT NULL,
  `score` int NOT NULL DEFAULT '0',
  `validation` enum('En attente','Validé','Rejeté') DEFAULT 'En attente',
  `motif_refus` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_user` (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `utilisateur`
--

INSERT INTO `utilisateur` (`id`, `role`, `nom`, `prenom`, `email`, `mot_de_passe`, `tutoriel_vu`, `type_statut`, `nom_entreprise`, `siret`, `score`, `validation`, `motif_refus`) VALUES
(1, 'Salarié', 'BOUDJENANE', 'Amel', 'amelboudjenanee@icloud.com', '$2a$10$TfC9LhGzA1ZMvU7X77vqb.d7bDdeIkWiTTMTp6WrinpAov2SWNYlO', 1, NULL, NULL, NULL, 0, 'Validé', NULL),
(6, 'Salarié', 'Martinnnn', 'Emma', 'emma.martin@mail.com', '', 1, NULL, NULL, NULL, 0, 'Rejeté', 'dddd'),
(7, 'Utilisateur', 'Bernard', 'Hugo', 'hugo.bernard@mail.com', '', 0, NULL, NULL, NULL, 0, 'Rejeté', NULL),
(9, 'Utilisateur', 'Robert', 'Nathan', 'nathan.robert@mail.com', '', 0, NULL, NULL, NULL, 0, 'Rejeté', NULL),
(10, 'Utilisateur', 'Richard', 'Lea', 'lea.richard@mail.com', '', 1, NULL, NULL, NULL, 0, 'Rejeté', NULL),
(11, 'Prestataire', 'Durand', 'Tom', 'tom.durand@mail.com', '', 1, NULL, NULL, NULL, 0, 'Rejeté', NULL),
(18, 'Utilisateur', 'ASEFAERF', 'WEDGFRF', 'thewarrow.bdj@gmail.com', '$2a$10$TfC9LhGzA1ZMvU7X77vqb.d7bDdeIkWiTTMTp6WrinpAov2SWNYlO', 0, 'Validé', NULL, NULL, 0, 'En attente', NULL);

--
-- Contraintes pour les tables déchargées
--

--
-- Contraintes pour la table `documents_legaux`
--
ALTER TABLE `documents_legaux`
  ADD CONSTRAINT `documents_legaux_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `utilisateur` (`id`) ON DELETE CASCADE;

--
-- Contraintes pour la table `historique_conteneurs`
--
ALTER TABLE `historique_conteneurs`
  ADD CONSTRAINT `fk_annonce` FOREIGN KEY (`annonce_id`) REFERENCES `annonce` (`id`),
  ADD CONSTRAINT `fk_conteneur` FOREIGN KEY (`conteneur_id`) REFERENCES `box_conteneur` (`id`),
  ADD CONSTRAINT `fk_particulier` FOREIGN KEY (`particulier_id`) REFERENCES `utilisateur` (`id`),
  ADD CONSTRAINT `fk_professionnel` FOREIGN KEY (`professionnel_id`) REFERENCES `utilisateur` (`id`);

--
-- Contraintes pour la table `translations`
--
ALTER TABLE `translations`
  ADD CONSTRAINT `translations_ibfk_1` FOREIGN KEY (`lang_code`) REFERENCES `languages` (`code`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
