-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : jeu. 02 avr. 2026 à 09:55
-- Version du serveur : 5.7.40
-- Version de PHP : 8.2.0

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
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) DEFAULT NULL,
  `id_plan` int(11) DEFAULT NULL,
  `date_debut` datetime DEFAULT NULL,
  `date_fin` datetime DEFAULT NULL,
  `statut` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  KEY `id_plan` (`id_plan`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `annonce`
--

DROP TABLE IF EXISTS `annonce`;
CREATE TABLE IF NOT EXISTS `annonce` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) DEFAULT NULL,
  `id_categorie` int(11) DEFAULT NULL,
  `titre` varchar(255) DEFAULT NULL,
  `description` text,
  `type` enum('Don','Vente','Service') DEFAULT NULL,
  `prix` decimal(10,2) DEFAULT NULL,
  `statut_validation` enum('En attente','Validee','Refusee') DEFAULT NULL,
  `code_postal` varchar(20) DEFAULT NULL,
  `ville` varchar(100) DEFAULT NULL,
  `profit_potentiel` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  KEY `id_categorie` (`id_categorie`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `annonce`
--

INSERT INTO `annonce` (`id`, `id_user`, `id_categorie`, `titre`, `description`, `type`, `prix`, `statut_validation`, `code_postal`, `ville`, `profit_potentiel`) VALUES
(1, 3, 1, 'Chaise en bois vintage', 'Ancienne chaise en chêne, parfaite pour un projet d\'upcycling. À venir chercher sur place.', 'Don', '0.00', 'Validee', '75011', 'Paris', '0.00'),
(2, 4, 3, 'Lot de chutes de bois massif', 'Belles planches de noyer récupérées suite à un chantier. Idéal pour petits meubles.', 'Vente', '25.50', 'Validee', '69003', 'Lyon', '5.00'),
(3, 5, 4, 'Réparation de petit électroménager', 'Je propose mes services pour réparer vos grille-pains, mixeurs et micro-ondes.', 'Service', '15.00', 'En attente', '31000', 'Toulouse', '2.50'),
(4, 3, 1, 'Canapé convertible abîmé', 'Le tissu est déchiré mais la structure est intacte. Pour bricoleur.', 'Don', '0.00', 'Refusee', '75018', 'Paris', '0.00');

-- --------------------------------------------------------

--
-- Structure de la table `article_news`
--

DROP TABLE IF EXISTS `article_news`;
CREATE TABLE IF NOT EXISTS `article_news` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_salarie` int(11) DEFAULT NULL,
  `titre` varchar(255) DEFAULT NULL,
  `contenu` text,
  `type` enum('News','Conseil') DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_salarie` (`id_salarie`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `box_conteneur`
--

DROP TABLE IF EXISTS `box_conteneur`;
CREATE TABLE IF NOT EXISTS `box_conteneur` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `localisation` varchar(255) DEFAULT NULL,
  `etat` enum('Disponible','Plein','Maintenance') DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `campagne_pub`
--

DROP TABLE IF EXISTS `campagne_pub`;
CREATE TABLE IF NOT EXISTS `campagne_pub` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_partenaire` int(11) DEFAULT NULL,
  `nom_campagne` varchar(255) DEFAULT NULL,
  `budget` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_partenaire` (`id_partenaire`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `categorie`
--

DROP TABLE IF EXISTS `categorie`;
CREATE TABLE IF NOT EXISTS `categorie` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `libelle` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `categorie`
--

INSERT INTO `categorie` (`id`, `libelle`) VALUES
(1, 'Mobilier'),
(2, 'Électroménager'),
(3, 'Matériaux de construction'),
(4, 'Outils et Bricolage'),
(5, 'Vêtements et Textiles');

-- --------------------------------------------------------

--
-- Structure de la table `commande`
--

DROP TABLE IF EXISTS `commande`;
CREATE TABLE IF NOT EXISTS `commande` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_acheteur` int(11) DEFAULT NULL,
  `id_annonce` int(11) DEFAULT NULL,
  `montant_total` decimal(10,2) DEFAULT NULL,
  `commission` decimal(5,2) DEFAULT NULL,
  `date_commande` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_acheteur` (`id_acheteur`),
  KEY `id_annonce` (`id_annonce`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `depot_box`
--

DROP TABLE IF EXISTS `depot_box`;
CREATE TABLE IF NOT EXISTS `depot_box` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_annonce` int(11) DEFAULT NULL,
  `id_box` int(11) DEFAULT NULL,
  `code_ouverture` varchar(50) DEFAULT NULL,
  `code_bonne_reception` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_annonce` (`id_annonce`),
  KEY `id_box` (`id_box`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `dictionnaire`
--

DROP TABLE IF EXISTS `dictionnaire`;
CREATE TABLE IF NOT EXISTS `dictionnaire` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `langue` varchar(50) DEFAULT NULL,
  `mot` varchar(255) DEFAULT NULL,
  `valeur` text,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `document`
--

DROP TABLE IF EXISTS `document`;
CREATE TABLE IF NOT EXISTS `document` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) DEFAULT NULL,
  `type_doc` enum('Facture','Contrat','Devis') DEFAULT NULL,
  `url_pdf` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `evenement`
--

DROP TABLE IF EXISTS `evenement`;
CREATE TABLE IF NOT EXISTS `evenement` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_salarie` int(11) DEFAULT NULL,
  `titre` varchar(255) DEFAULT NULL,
  `description` text,
  `date_debut` datetime DEFAULT NULL,
  `date_fin` datetime DEFAULT NULL,
  `nb_places` int(11) DEFAULT NULL,
  `statut_validation` varchar(50) DEFAULT NULL,
  `format` enum('En Ligne','Presentiel') DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_salarie` (`id_salarie`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `evenement`
--

INSERT INTO `evenement` (`id`, `id_salarie`, `titre`, `description`, `date_debut`, `date_fin`, `nb_places`, `statut_validation`, `format`) VALUES
(1, 2, 'Atelier Upcycling : Customiser ses meubles', 'Apprenez les bases du ponçage et de la peinture sur meuble ancien avec nos experts.', '2026-04-15 14:00:00', '2026-04-15 17:00:00', 15, 'Valide', 'Presentiel'),
(2, 2, 'Webinaire : Les bases de l\'économie circulaire', 'Conférence gratuite en ligne pour comprendre comment réduire ses déchets.', '2026-04-20 18:30:00', '2026-04-20 19:30:00', 100, 'Valide', 'En Ligne'),
(3, 2, 'Initiation à la réparation vélo', 'Venez avec votre vélo, nous vous apprenons à réparer les pannes courantes.', '2026-05-10 10:00:00', '2026-05-10 12:30:00', 10, 'En attente', 'Presentiel');

-- --------------------------------------------------------

--
-- Structure de la table `inscription`
--

DROP TABLE IF EXISTS `inscription`;
CREATE TABLE IF NOT EXISTS `inscription` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) DEFAULT NULL,
  `id_event` int(11) DEFAULT NULL,
  `date_inscription` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  KEY `id_event` (`id_event`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `log_connexion`
--

DROP TABLE IF EXISTS `log_connexion`;
CREATE TABLE IF NOT EXISTS `log_connexion` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) DEFAULT NULL,
  `ip` varchar(50) DEFAULT NULL,
  `date_connexion` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `message_forum`
--

DROP TABLE IF EXISTS `message_forum`;
CREATE TABLE IF NOT EXISTS `message_forum` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_theme` int(11) DEFAULT NULL,
  `id_user` int(11) DEFAULT NULL,
  `contenu` text,
  `est_resolu` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_theme` (`id_theme`),
  KEY `id_user` (`id_user`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `notification`
--

DROP TABLE IF EXISTS `notification`;
CREATE TABLE IF NOT EXISTS `notification` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) DEFAULT NULL,
  `contenu` text,
  `est_lu` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `paiement`
--

DROP TABLE IF EXISTS `paiement`;
CREATE TABLE IF NOT EXISTS `paiement` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_commande` int(11) DEFAULT NULL,
  `montant` decimal(10,2) DEFAULT NULL,
  `statut_paiement` varchar(50) DEFAULT NULL,
  `date_paiement` datetime DEFAULT NULL,
  `id_stripe` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_commande` (`id_commande`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `partenaire`
--

DROP TABLE IF EXISTS `partenaire`;
CREATE TABLE IF NOT EXISTS `partenaire` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nom_entreprise` varchar(255) DEFAULT NULL,
  `type_partenariat` enum('Pub','Sponsor','Fournisseur') DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `plan_abo`
--

DROP TABLE IF EXISTS `plan_abo`;
CREATE TABLE IF NOT EXISTS `plan_abo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nom` varchar(255) DEFAULT NULL,
  `prix` decimal(10,2) DEFAULT NULL,
  `duree_mois` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `projet_pro`
--

DROP TABLE IF EXISTS `projet_pro`;
CREATE TABLE IF NOT EXISTS `projet_pro` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) DEFAULT NULL,
  `titre` varchar(255) DEFAULT NULL,
  `desc_etapes` text,
  `url_photo_avant` varchar(255) DEFAULT NULL,
  `url_photo_apres` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `ressource_pedagogique`
--

DROP TABLE IF EXISTS `ressource_pedagogique`;
CREATE TABLE IF NOT EXISTS `ressource_pedagogique` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_salarie` int(11) DEFAULT NULL,
  `titre` varchar(255) DEFAULT NULL,
  `type_ressource` varchar(100) DEFAULT NULL,
  `url_fichier` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_salarie` (`id_salarie`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `theme_forum`
--

DROP TABLE IF EXISTS `theme_forum`;
CREATE TABLE IF NOT EXISTS `theme_forum` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) DEFAULT NULL,
  `titre` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `upcycling_score`
--

DROP TABLE IF EXISTS `upcycling_score`;
CREATE TABLE IF NOT EXISTS `upcycling_score` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) DEFAULT NULL,
  `points_totaux` int(11) DEFAULT NULL,
  `poids_evite_kg` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Structure de la table `utilisateur`
--

DROP TABLE IF EXISTS `utilisateur`;
CREATE TABLE IF NOT EXISTS `utilisateur` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role` enum('Admin','Salarie','Pro','Particulier') DEFAULT NULL,
  `nom` varchar(255) DEFAULT NULL,
  `prenom` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `mot_de_passe` varchar(255) DEFAULT NULL,
  `tutoriel_vu` tinyint(1) DEFAULT '0',
  `type_statut` varchar(255) DEFAULT NULL,
  `nom_entreprise` varchar(255) DEFAULT NULL,
  `siret` varchar(50) DEFAULT NULL,
  `score` int(11) DEFAULT '0',
  `date_naissance` date DEFAULT NULL,
  `validation` enum('En attente','Validé','Rejeté') DEFAULT 'En attente',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

--
-- Déchargement des données de la table `utilisateur`
--

INSERT INTO `utilisateur` (`id`, `role`, `nom`, `prenom`, `email`, `mot_de_passe`, `tutoriel_vu`, `type_statut`, `nom_entreprise`, `siret`, `score`, `date_naissance`, `validation`) VALUES
(1, 'Admin', 'Dujardin', 'Jean', 'admin@renova.fr', '', 1, NULL, NULL, NULL, 0, NULL, NULL),
(2, 'Admin', 'Laurent', 'Sophie', 's.laurent@renova.fr', '', 1, NULL, NULL, NULL, 0, NULL, NULL),
(3, 'Particulier', 'Dupont', 'Marie', 'marie.dupont@email.com', '$2y$10$ExempleHashPassword123', 0, NULL, NULL, NULL, 0, NULL, NULL),
(4, 'Pro', 'Morin', 'Jean', 'contact@ateliermorin.fr', '', 1, NULL, NULL, NULL, 0, NULL, NULL),
(5, 'Particulier', 'Petit', 'Lucas', 'lucas.petit@gmail.com', '$2y$10$ExempleHashPassword123', 1, NULL, NULL, NULL, 0, NULL, NULL);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
