-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : ven. 20 mars 2026 à 22:51
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
  `statut_validation` varchar(20) DEFAULT 'En attente',
  `code_postal` varchar(10) DEFAULT NULL,
  `ville` varchar(100) DEFAULT NULL,
  `projet_potentiel` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_annonce` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `annonce`
--

INSERT INTO `annonce` (`id`, `id_user`, `id_categorie`, `titre`, `description`, `type`, `prix`, `statut_validation`, `code_postal`, `ville`, `projet_potentiel`) VALUES
(1, 1, 3, 'Tondeuse à gazon électrique', 'Je vends ma tondeuse en très bon état, servie 3 fois.', 'Vente', 45.50, 'valide', '93420', 'Villepinte', NULL);

-- --------------------------------------------------------

--
-- Structure de la table `article_news`
--

DROP TABLE IF EXISTS `article_news`;
CREATE TABLE IF NOT EXISTS `article_news` (
  `id_article` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `id_salarie` int DEFAULT NULL,
  `titre` varchar(150) DEFAULT NULL,
  `contenu` text,
  `type` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id_article`),
  UNIQUE KEY `id_article` (`id_article`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Structure de la table `box_conteneur`
--

DROP TABLE IF EXISTS `box_conteneur`;
CREATE TABLE IF NOT EXISTS `box_conteneur` (
  `id_box` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `localisation` varchar(255) DEFAULT NULL,
  `etat` varchar(50) DEFAULT 'Disponible',
  PRIMARY KEY (`id_box`),
  UNIQUE KEY `id_box` (`id_box`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

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
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_event` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `evenement`
--

INSERT INTO `evenement` (`id`, `id_salarie`, `titre`, `description`, `date_debut`, `date_fin`, `nb_places`, `statut_validation`, `format`) VALUES
(1, 1, 'Atelier Couture Débutant', 'Apprenez à rapiécer vos vêtements au lieu de les jeter.', '2026-03-20 13:00:00', '2026-03-20 16:00:00', 8, 'valide', 'Présentiel');

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
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_user` (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Déchargement des données de la table `utilisateur`
--

INSERT INTO `utilisateur` (`id`, `role`, `nom`, `prenom`, `email`, `mot_de_passe`, `tutoriel_vu`, `type_statut`, `nom_entreprise`, `siret`, `score`) VALUES
(1, 'Administrateur', 'BOUDJENANE', 'Amel', 'amelboudjenanee@icloud.com', '$2y$10$cm5C6PeLqY5q/OTKazT/KO68drOeCYJRqgwWtfzPWkEo1quHMvlN.', 1, NULL, NULL, NULL, 0),
(6, 'Salarié', 'Martinnnn', 'Emma', 'emma.martin@mail.com', '', 1, NULL, NULL, NULL, 0),
(7, 'Utilisateur', 'Bernard', 'Hugo', 'hugo.bernard@mail.com', '', 0, NULL, NULL, NULL, 0),
(9, 'Utilisateur', 'Robert', 'Nathan', 'nathan.robert@mail.com', '', 0, NULL, NULL, NULL, 0),
(10, 'Utilisateur', 'Richard', 'Lea', 'lea.richard@mail.com', '', 1, NULL, NULL, NULL, 0),
(11, 'Prestataire', 'Durand', 'Tom', 'tom.durand@mail.com', '', 1, NULL, NULL, NULL, 0);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
