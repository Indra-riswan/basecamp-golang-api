-- MySQL dump 10.13  Distrib 8.0.29, for Linux (x86_64)
--
-- Host: localhost    Database: basecamp-golang-api
-- ------------------------------------------------------
-- Server version	8.0.30-0ubuntu0.20.04.2

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `brands`
--

DROP TABLE IF EXISTS `brands`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `brands` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nama` char(50) DEFAULT NULL,
  `password` longtext NOT NULL,
  `website` char(100) DEFAULT NULL,
  `marketing` char(50) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `nohp` bigint DEFAULT NULL,
  `gambar` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_brands_nama` (`nama`),
  UNIQUE KEY `idx_brands_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `brands`
--

LOCK TABLES `brands` WRITE;
/*!40000 ALTER TABLE `brands` DISABLE KEYS */;
INSERT INTO `brands` VALUES (1,'Lepaskuncilombok','$2a$04$XvTSgkHgmfacFhbhJJb4ReySLJIJx4GAGYumP3krTCLs12KMuLyiK','lepaskuncilombok.com','Dino','lepaskuncilombok@gmail.com',6285971837872,'lepaskuncilombok.jpeg'),(2,'Nonemtrans','$2a$04$T8/fE5yvUm8WX1SoyiW9M.KHcUWxc9el1Kyh2YqNU3TQx2QiAIoY.','nonemtrans.com','Erwin','nonemtrans@gmail.com',6287750333343,'nonemtrans.jpeg'),(3,'Mataramrentcar','$2a$04$73YZOPF7mXXfH1BeTsHcG.g3RpNXQHzFo471HuERItPyoGMDOn.hy','mataramrentcar.com','Bagong','mataramrentcar@gmail.com',6285333502625,'mataramrentcar.jpeg'),(4,'Lokatrans','$2a$04$kx/tIpCXkgKkn0JSfxLOaeRjqv3.jxLF7Ix4IEv0JeAED2jlPqSha','wisatadilombok.com','Loka','wisatadilombok@gmail.com',6285338881800,'lokatrans.jpeg'),(5,'Adventourid','$2a$04$GemUUrMK24TuAWJ5xi0EyuA1cF9K1GdaimioeLC5//7dNDQy3/C4C','adventourid.com','Iwan','adventourlombok@gmail.com',6285333987570,'adventourid.jpeg'),(6,'Indotravelin','$2a$04$Zd5PevAxXtIRSt/rMR4Gbe5rB9jHj9FrACd4Juw4fDozXh2Dq8NzK','indotravelin.com','Yogi','indotravelin@gmail.com',6287844781533,'indotravelin.jpeg'),(7,'Baletrans','$2a$04$Be1eFepRZtUs6k.2o5lLh.A68tVVQd1lmiwypt5h4lhgwgDam8toC','baletrans.com','fajar herla','baletranslombok@gmail.com',6285971837872,'baletrans.jpg');
/*!40000 ALTER TABLE `brands` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mitras`
--

DROP TABLE IF EXISTS `mitras`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mitras` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nama` char(50) DEFAULT NULL,
  `password` longtext NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_mitras_nama` (`nama`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mitras`
--

LOCK TABLES `mitras` WRITE;
/*!40000 ALTER TABLE `mitras` DISABLE KEYS */;
INSERT INTO `mitras` VALUES (1,'mitra 1','$2a$04$GvwKUUIYy/EdJWVrB7vlmeckP1GrIra9kJWJRLqqBq09vTYasN/6y'),(2,'mitra 2','$2a$04$sBkBH0y9/3Z1IxvF/GJowOz2BHiYF59Wpj1ze6o9e.ZJYXb9mE3Q6'),(3,'mitra 3','$2a$04$0gQxPvYNfMZsJ.y/ZP6aJOy1woCcrdqxn6tVrIPz/OYwjiKQrD4vi');
/*!40000 ALTER TABLE `mitras` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mobils`
--

DROP TABLE IF EXISTS `mobils`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mobils` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `unit` char(50) DEFAULT NULL,
  `nopol` char(20) DEFAULT NULL,
  `gambar` char(50) DEFAULT NULL,
  `isready` tinyint(1) DEFAULT '1',
  `own` char(50) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_mobils_mitra` (`own`),
  CONSTRAINT `fk_mobils_mitra` FOREIGN KEY (`own`) REFERENCES `mitras` (`nama`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mobils`
--

LOCK TABLES `mobils` WRITE;
/*!40000 ALTER TABLE `mobils` DISABLE KEYS */;
INSERT INTO `mobils` VALUES (1,'Xenia','DR 1234 SS','xenia.jpg',1,'mitra 3'),(2,'Avanza','DR 5678 SS','avanza.jpg',1,'mitra 3'),(3,'Agya','DR 91011 SS','agya.jpg',1,'mitra 3'),(4,'Alphard','DR 1213 SS','alphard.jpg',1,'mitra 1'),(5,'Ertiga','DR 1415 SS','ertiga.jpg',1,'mitra 1'),(6,'Agya','DR 1617 SS','agya.jpg',1,'mitra 2'),(7,'Reborn','DR 1819 SS','reborn.jpg',1,'mitra 2');
/*!40000 ALTER TABLE `mobils` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaksis`
--

DROP TABLE IF EXISTS `transaksis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transaksis` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `brand` char(50) DEFAULT NULL,
  `mitra` char(50) DEFAULT NULL,
  `unit` char(20) DEFAULT NULL,
  `nopol` char(50) DEFAULT NULL,
  `user` char(50) DEFAULT NULL,
  `checker` char(50) DEFAULT NULL,
  `status` char(20) DEFAULT NULL,
  `driver` char(20) DEFAULT NULL,
  `out` char(10) DEFAULT NULL,
  `in` char(10) DEFAULT NULL,
  `day` bigint DEFAULT NULL,
  `perhari` bigint DEFAULT NULL,
  `op` bigint DEFAULT NULL,
  `total` bigint unsigned NOT NULL,
  `total_op` bigint unsigned NOT NULL,
  `sisa_profit` bigint unsigned NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_mitras_penjualan` (`mitra`),
  KEY `fk_brands_penjualan` (`brand`),
  CONSTRAINT `fk_brands_penjualan` FOREIGN KEY (`brand`) REFERENCES `brands` (`nama`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_mitras_penjualan` FOREIGN KEY (`mitra`) REFERENCES `mitras` (`nama`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaksis`
--

LOCK TABLES `transaksis` WRITE;
/*!40000 ALTER TABLE `transaksis` DISABLE KEYS */;
INSERT INTO `transaksis` VALUES (1,'Baletrans','mitra 1','Alphard','DR 1234 SH','costumer 1','Frenky','lepas kunci','','11.00 AM','11.00 AM',2,4500000,3500000,9000000,7000000,2000000,'2022-07-17 13:44:45.245','2022-07-17 13:44:45.245'),(2,'Baletrans','mitra 2','Avanza','DR 2345 SH','costumer 1','Frenky','lepas kunci','','11.00 AM','11.00 AM',2,4500000,3500000,9000000,7000000,2000000,'2022-07-17 13:45:23.386','2022-07-17 13:45:23.386'),(5,'Lepaskuncilombok','mitra 2','Ayla','DR 9101 SH','costumer 4','Frenky','lepas kunci','','09.00 AM','09.00 AM',5,450000,350000,2250000,1750000,500000,'2022-07-17 13:50:26.958','2022-07-17 13:50:26.958'),(6,'Adventourid','mitra 3','Agya','DR 5678 SH','costumer 3','Frenky','lepas kunci','','11.40 AM','11.40 AM',6,450000,350000,2700000,2100000,600000,'2022-07-17 13:51:33.593','2022-07-17 13:51:33.593');
/*!40000 ALTER TABLE `transaksis` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-08-30 16:14:06
