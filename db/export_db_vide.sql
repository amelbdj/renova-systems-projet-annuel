-- MySQL dump 10.13  Distrib 8.0.46, for Linux (x86_64)
--
-- Host: localhost    Database: pa2026
-- ------------------------------------------------------
-- Server version	8.0.46

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `pa2026`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `pa2026` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `pa2026`;

--
-- Table structure for table `abonnement`
--

DROP TABLE IF EXISTS `abonnement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `abonnement` (
  `id_abonnement` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `id_plan` int DEFAULT NULL,
  `date_debut` date DEFAULT NULL,
  `date_fin` date DEFAULT NULL,
  `statut` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id_abonnement`),
  UNIQUE KEY `id_abonnement` (`id_abonnement`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `annonce`
--

DROP TABLE IF EXISTS `annonce`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `annonce` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
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
  `image` varchar(255) DEFAULT NULL,
  `is_sponsored` tinyint(1) DEFAULT '0',
  `commission_prelevee` decimal(10,2) DEFAULT '0.00',
  `id_box` int DEFAULT NULL,
  `photo_url` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `statut_vente` enum('EN VENTE','EN ATTENTE DEPOT','RESERVEE','EN BOX','EN ATTENTE DE RECUPERATION','RECUPERE','VENDU') DEFAULT 'EN VENTE',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_annonce` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `article_news`
--

DROP TABLE IF EXISTS `article_news`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `article_news` (
  `id_article` bigint unsigned NOT NULL AUTO_INCREMENT,
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
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `box`
--

DROP TABLE IF EXISTS `box`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `box` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_conteneur` int NOT NULL,
  `numero` int NOT NULL,
  `taille` varchar(10) NOT NULL,
  `statut` varchar(50) DEFAULT 'libre',
  `code_secret` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_conteneur` (`id_conteneur`),
  CONSTRAINT `box_ibfk_1` FOREIGN KEY (`id_conteneur`) REFERENCES `conteneur` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `box_conteneur`
--

DROP TABLE IF EXISTS `box_conteneur`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `box_conteneur` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `localisation` varchar(255) DEFAULT NULL,
  `etat` enum('MAINTENANCE','LIBRE','OCCUPE','') DEFAULT 'LIBRE',
  `type_materiau_accepte` varchar(50) DEFAULT NULL,
  `capacite` int DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_box` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `campagne_pub`
--

DROP TABLE IF EXISTS `campagne_pub`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `campagne_pub` (
  `id_campagne` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_partenaire` int DEFAULT NULL,
  `titre` varchar(150) DEFAULT NULL,
  `budget` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`id_campagne`),
  UNIQUE KEY `id_campagne` (`id_campagne`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `categorie`
--

DROP TABLE IF EXISTS `categorie`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categorie` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `libelle` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_categorie` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `conteneur`
--

DROP TABLE IF EXISTS `conteneur`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `conteneur` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nom` varchar(100) NOT NULL,
  `adresse` varchar(255) NOT NULL,
  `date_creation` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `depot_box`
--

DROP TABLE IF EXISTS `depot_box`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `depot_box` (
  `id_depot` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_annonce` int DEFAULT NULL,
  `id_box` int DEFAULT NULL,
  `code_ouverture` varchar(10) DEFAULT NULL,
  `code_barres_pro` varchar(50) DEFAULT NULL,
  `date_depot` date DEFAULT NULL,
  `date_retrait` date DEFAULT NULL,
  PRIMARY KEY (`id_depot`),
  UNIQUE KEY `id_depot` (`id_depot`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `dictionnaire`
--

DROP TABLE IF EXISTS `dictionnaire`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `dictionnaire` (
  `id_trad` bigint unsigned NOT NULL AUTO_INCREMENT,
  `langue` varchar(5) DEFAULT NULL,
  `cle` varchar(100) DEFAULT NULL,
  `valeur` text,
  PRIMARY KEY (`id_trad`),
  UNIQUE KEY `id_trad` (`id_trad`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `document`
--

DROP TABLE IF EXISTS `document`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `document` (
  `id_document` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `type_doc` varchar(50) DEFAULT NULL,
  `url_pdf` varchar(255) DEFAULT NULL,
  `id_commande` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id_document`),
  UNIQUE KEY `id_document` (`id_document`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `documents_legaux`
--

DROP TABLE IF EXISTS `documents_legaux`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `documents_legaux` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `type_document` enum('KBIS','SIREN','DIPLOME','PIECE_IDENTITE') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `chemin_fichier` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `statut_document` enum('En attente','Validé','Rejeté') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT 'En attente',
  `date_soumission` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `documents_legaux_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `utilisateur` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `etapes_projet`
--

DROP TABLE IF EXISTS `etapes_projet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `etapes_projet` (
  `id_etape` int NOT NULL AUTO_INCREMENT,
  `id_projet` int NOT NULL,
  `titre` varchar(150) NOT NULL,
  `description` text,
  `statut` varchar(50) DEFAULT 'a_faire',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_etape`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `evenement`
--

DROP TABLE IF EXISTS `evenement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `evenement` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_salarie` int DEFAULT NULL,
  `titre` varchar(150) DEFAULT NULL,
  `description` text,
  `date_debut` timestamp NULL DEFAULT NULL,
  `date_fin` timestamp NULL DEFAULT NULL,
  `nb_places` int DEFAULT NULL,
  `statut_validation` varchar(20) DEFAULT NULL,
  `format` varchar(20) DEFAULT NULL,
  `type` enum('evenement','formation','atelier','reunion') DEFAULT NULL,
  `lieu` varchar(50) DEFAULT '0',
  `prix` int DEFAULT '0',
  `image_url` varchar(255) DEFAULT '',
  `plan_cours` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_event` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `historique_conteneurs`
--

DROP TABLE IF EXISTS `historique_conteneurs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `historique_conteneurs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `conteneur_id` bigint unsigned NOT NULL,
  `annonce_id` bigint unsigned NOT NULL,
  `acheteur_id` bigint unsigned NOT NULL,
  `vendeur_id` bigint unsigned DEFAULT NULL,
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
  KEY `fk_particulier` (`acheteur_id`),
  KEY `fk_professionnel` (`vendeur_id`),
  CONSTRAINT `fk_annonce` FOREIGN KEY (`annonce_id`) REFERENCES `annonce` (`id`),
  CONSTRAINT `fk_conteneur` FOREIGN KEY (`conteneur_id`) REFERENCES `box_conteneur` (`id`),
  CONSTRAINT `fk_particulier` FOREIGN KEY (`acheteur_id`) REFERENCES `utilisateur` (`id`),
  CONSTRAINT `fk_professionnel` FOREIGN KEY (`vendeur_id`) REFERENCES `utilisateur` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `inscription`
--

DROP TABLE IF EXISTS `inscription`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `inscription` (
  `id_inscription` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `id_event` int DEFAULT NULL,
  `date_inscrip` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_inscription`),
  UNIQUE KEY `id_inscription` (`id_inscription`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `languages`
--

DROP TABLE IF EXISTS `languages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `languages` (
  `id` int NOT NULL AUTO_INCREMENT,
  `code` varchar(5) NOT NULL,
  `name` varchar(50) NOT NULL,
  `is_default` tinyint(1) DEFAULT '0',
  `is_active` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `log_connexion`
--

DROP TABLE IF EXISTS `log_connexion`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `log_connexion` (
  `id_log` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `ip` varchar(45) DEFAULT NULL,
  `date_connexion` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_log`),
  UNIQUE KEY `id_log` (`id_log`)
) ENGINE=InnoDB AUTO_INCREMENT=200 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message` (
  `id` int NOT NULL AUTO_INCREMENT,
  `annonce_id` bigint unsigned NOT NULL,
  `expediteur_id` bigint unsigned NOT NULL,
  `destinataire_id` bigint unsigned NOT NULL,
  `contenu` text NOT NULL,
  `lu` tinyint(1) DEFAULT '0',
  `date_envoi` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_message_expediteur` (`expediteur_id`),
  KEY `fk_message_destinataire` (`destinataire_id`),
  KEY `idx_conversation` (`annonce_id`,`expediteur_id`,`destinataire_id`),
  CONSTRAINT `fk_message_annonce` FOREIGN KEY (`annonce_id`) REFERENCES `annonce` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_message_destinataire` FOREIGN KEY (`destinataire_id`) REFERENCES `utilisateur` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_message_expediteur` FOREIGN KEY (`expediteur_id`) REFERENCES `utilisateur` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `message_forum`
--

DROP TABLE IF EXISTS `message_forum`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `message_forum` (
  `id_message` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_topic` int DEFAULT NULL,
  `id_user` int DEFAULT NULL,
  `contenu` text,
  `est_modere` tinyint(1) DEFAULT '0',
  `date_creation` datetime DEFAULT CURRENT_TIMESTAMP,
  `est_signale` tinyint(1) DEFAULT '0',
  `statut` varchar(50) DEFAULT 'public',
  PRIMARY KEY (`id_message`),
  UNIQUE KEY `id_message` (`id_message`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `notification`
--

DROP TABLE IF EXISTS `notification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `notification` (
  `id_notif` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `contenu` text,
  `est_lu` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id_notif`),
  UNIQUE KEY `id_notif` (`id_notif`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `order`
--

DROP TABLE IF EXISTS `order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order` (
  `id_commande` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_acheteur` int DEFAULT NULL,
  `id_annonce` int DEFAULT NULL,
  `montant_total` decimal(10,2) DEFAULT NULL,
  `commission` decimal(10,2) DEFAULT NULL,
  `date_commande` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `type` varchar(20) DEFAULT 'annonce',
  PRIMARY KEY (`id_commande`),
  UNIQUE KEY `id_commande` (`id_commande`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `paiement`
--

DROP TABLE IF EXISTS `paiement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `paiement` (
  `id_paiement` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_commande` int DEFAULT NULL,
  `stripe_id` varchar(255) DEFAULT NULL,
  `statut` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id_paiement`),
  UNIQUE KEY `id_paiement` (`id_paiement`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `partenaire`
--

DROP TABLE IF EXISTS `partenaire`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `partenaire` (
  `id_partenaire` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nom_entreprise` varchar(150) DEFAULT NULL,
  `type_partenariat` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id_partenaire`),
  UNIQUE KEY `id_partenaire` (`id_partenaire`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `plan_abo`
--

DROP TABLE IF EXISTS `plan_abo`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `plan_abo` (
  `id_plan` bigint unsigned NOT NULL AUTO_INCREMENT,
  `nom` varchar(100) DEFAULT NULL,
  `prix` decimal(10,2) DEFAULT NULL,
  `duree_mois` int DEFAULT NULL,
  PRIMARY KEY (`id_plan`),
  UNIQUE KEY `id_plan` (`id_plan`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `projet_pro`
--

DROP TABLE IF EXISTS `projet_pro`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `projet_pro` (
  `id_projet` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `titre` varchar(150) DEFAULT NULL,
  `desc_etapes` text,
  `url_photo_avant` varchar(255) DEFAULT NULL,
  `url_photo_apres` varchar(255) DEFAULT NULL,
  `statut` varchar(50) DEFAULT 'en_cours',
  `co2_evite` decimal(10,2) DEFAULT '0.00',
  PRIMARY KEY (`id_projet`),
  UNIQUE KEY `id_projet` (`id_projet`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `ressource_pedagogique`
--

DROP TABLE IF EXISTS `ressource_pedagogique`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ressource_pedagogique` (
  `id_ressource` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_salarie` int DEFAULT NULL,
  `titre` varchar(150) DEFAULT NULL,
  `url_fichier` varchar(255) DEFAULT NULL,
  `id_event` int DEFAULT NULL,
  PRIMARY KEY (`id_ressource`),
  UNIQUE KEY `id_ressource` (`id_ressource`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `topic_forum`
--

DROP TABLE IF EXISTS `topic_forum`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `topic_forum` (
  `id_topic` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `titre` varchar(200) DEFAULT NULL,
  `date_creation` datetime DEFAULT CURRENT_TIMESTAMP,
  `est_signale` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id_topic`),
  UNIQUE KEY `id_topic` (`id_topic`)
) ENGINE=InnoDB AUTO_INCREMENT=106 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `translations`
--

DROP TABLE IF EXISTS `translations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `translations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `lang_code` varchar(5) DEFAULT NULL,
  `msg_key` varchar(100) DEFAULT NULL,
  `msg_value` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_lang_key` (`lang_code`,`msg_key`),
  CONSTRAINT `translations_ibfk_1` FOREIGN KEY (`lang_code`) REFERENCES `languages` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=2917 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `upcycling_score`
--

DROP TABLE IF EXISTS `upcycling_score`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `upcycling_score` (
  `id_score` bigint unsigned NOT NULL AUTO_INCREMENT,
  `id_user` int DEFAULT NULL,
  `points_totaux` int DEFAULT '0',
  `poids_evite_kg` decimal(10,2) DEFAULT '0.00',
  PRIMARY KEY (`id_score`),
  UNIQUE KEY `id_score` (`id_score`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `utilisateur`
--

DROP TABLE IF EXISTS `utilisateur`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `utilisateur` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `role` enum('Utilisateur','Salarié','Administrateur','Pro') NOT NULL,
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
  `stripe_account_id` varchar(255) DEFAULT NULL,
  `stripe_verif_completed` tinyint(1) DEFAULT '0',
  `est_premium` tinyint(1) NOT NULL DEFAULT '0',
  `stripe_customer_id` varchar(255) DEFAULT NULL,
  `plan_abo` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_user` (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=15060 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-07-05 20:38:46
