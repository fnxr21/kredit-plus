-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               9.1.0 - MySQL Community Server - GPL
-- Server OS:                    Linux
-- HeidiSQL Version:             12.6.0.6765
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for kredit_mst
CREATE DATABASE IF NOT EXISTS `kredit_mst` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `kredit_mst`;

-- Dumping structure for table kredit_mst.asset
CREATE TABLE IF NOT EXISTS `asset` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(100) NOT NULL COMMENT 'mobil,rumah',
  `type` varchar(100) NOT NULL COMMENT 'Properti,Elektronik,Kendaraan,Investasi',
  `amount` decimal(20,2) NOT NULL,
  `partner_id` bigint unsigned DEFAULT NULL COMMENT 'partner could be null ,add null when this is from customer or conventional partner',
  `customer_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_asset_deleted_at` (`deleted_at`),
  KEY `fk_asset_partner` (`partner_id`),
  KEY `fk_asset_customer` (`customer_id`),
  CONSTRAINT `fk_asset_customer` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`),
  CONSTRAINT `fk_asset_partner` FOREIGN KEY (`partner_id`) REFERENCES `partner` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table kredit_mst.asset: ~0 rows (approximately)

-- Dumping structure for table kredit_mst.credit_limit
CREATE TABLE IF NOT EXISTS `credit_limit` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tenor_months` double NOT NULL COMMENT 'lama bulan pelunasan',
  `limit_amount` decimal(15,2) NOT NULL COMMENT 'limit transaction',
  PRIMARY KEY (`id`),
  KEY `idx_credit_limit_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table kredit_mst.credit_limit: ~0 rows (approximately)

-- Dumping structure for table kredit_mst.customer
CREATE TABLE IF NOT EXISTS `customer` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `phone_number` varchar(100) NOT NULL,
  `nik` varchar(100) NOT NULL,
  `full_name` varchar(100) NOT NULL,
  `legal_name` varchar(100) NOT NULL,
  `birthplace` varchar(100) NOT NULL,
  `dob` date NOT NULL,
  `salary` decimal(15,2) NOT NULL,
  `image_ktp` varchar(255) NOT NULL COMMENT 'pathfile',
  `image_selfie` varchar(255) NOT NULL COMMENT 'pathfile',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_customer_username` (`username`),
  UNIQUE KEY `uni_customer_email` (`email`),
  UNIQUE KEY `uni_customer_nik` (`nik`),
  KEY `idx_customer_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table kredit_mst.customer: ~0 rows (approximately)

-- Dumping structure for table kredit_mst.my_user
CREATE TABLE IF NOT EXISTS `my_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `phone_number` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_my_user_username` (`username`),
  UNIQUE KEY `uni_my_user_email` (`email`),
  KEY `idx_my_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table kredit_mst.my_user: ~1 rows (approximately)
INSERT IGNORE INTO `my_user` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `password`, `phone_number`, `email`) VALUES
	(1, '2024-12-19 00:17:50.000', NULL, NULL, 'admin', '$2a$10$1Rff5M8Vq7A7gTBfi.T2puamxaxaMluWbm0gyFFKK7cnF90eXsNgq', '0893213124', 'test@gmail');

-- Dumping structure for table kredit_mst.partner
CREATE TABLE IF NOT EXISTS `partner` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `phone_number` varchar(100) NOT NULL,
  `address` varchar(100) NOT NULL,
  `partner_bank_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_partner_email` (`email`),
  KEY `idx_partner_deleted_at` (`deleted_at`),
  KEY `fk_partner_partner_bank` (`partner_bank_id`),
  CONSTRAINT `fk_partner_partner_bank` FOREIGN KEY (`partner_bank_id`) REFERENCES `partner_bank` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table kredit_mst.partner: ~0 rows (approximately)

-- Dumping structure for table kredit_mst.partner_bank
CREATE TABLE IF NOT EXISTS `partner_bank` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `bank_account` varchar(100) NOT NULL,
  `account_holder_name` varchar(100) NOT NULL,
  `bank_name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_partner_bank_bank_account` (`bank_account`),
  KEY `idx_partner_bank_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table kredit_mst.partner_bank: ~0 rows (approximately)

-- Dumping structure for table kredit_mst.transaction_detail
CREATE TABLE IF NOT EXISTS `transaction_detail` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `contract_number` varchar(200) NOT NULL,
  `otr` decimal(15,2) DEFAULT NULL COMMENT '2 jenis: white godds(kulkas) dan total biaya adminstrasi dll(mobil)',
  `admin_fee` decimal(15,2) DEFAULT NULL,
  `installment_amount` decimal(15,2) DEFAULT NULL COMMENT 'total cicilan perbulan',
  `interest_amount` decimal(15,2) DEFAULT NULL COMMENT 'biaya++ (bunga)',
  `payment` enum('pending','accept','reject') NOT NULL DEFAULT 'pending',
  `credit_limit_id` bigint unsigned DEFAULT NULL,
  `partner_bank_id` bigint unsigned NOT NULL,
  `asset_id` bigint unsigned DEFAULT NULL,
  `partner_id` bigint unsigned DEFAULT NULL COMMENT 'partner could be null ,add null when this is from customer or conventional partner',
  `customer_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_transaction_detail_contract_number` (`contract_number`),
  KEY `idx_transaction_detail_deleted_at` (`deleted_at`),
  KEY `fk_transaction_detail_credit_limit` (`credit_limit_id`),
  KEY `fk_transaction_detail_partner_bank` (`partner_bank_id`),
  KEY `fk_transaction_detail_asset` (`asset_id`),
  KEY `fk_transaction_detail_partner` (`partner_id`),
  KEY `fk_transaction_detail_customer` (`customer_id`),
  CONSTRAINT `fk_transaction_detail_asset` FOREIGN KEY (`asset_id`) REFERENCES `asset` (`id`),
  CONSTRAINT `fk_transaction_detail_credit_limit` FOREIGN KEY (`credit_limit_id`) REFERENCES `credit_limit` (`id`),
  CONSTRAINT `fk_transaction_detail_customer` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`),
  CONSTRAINT `fk_transaction_detail_partner` FOREIGN KEY (`partner_id`) REFERENCES `partner` (`id`),
  CONSTRAINT `fk_transaction_detail_partner_bank` FOREIGN KEY (`partner_bank_id`) REFERENCES `partner_bank` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table kredit_mst.transaction_detail: ~0 rows (approximately)

-- Dumping structure for table kredit_mst.transaction_payment
CREATE TABLE IF NOT EXISTS `transaction_payment` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `transaction_detail_id` bigint unsigned DEFAULT NULL,
  `payment_date` datetime(3) NOT NULL,
  `amount` decimal(15,2) DEFAULT NULL,
  `status` enum('pending','completed','failed') NOT NULL DEFAULT 'pending',
  `partner_id` bigint unsigned DEFAULT NULL COMMENT 'partner could be null ,add null when this is from customer or conventional partner',
  `customer_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_transaction_payment_deleted_at` (`deleted_at`),
  KEY `fk_transaction_payment_customer` (`customer_id`),
  KEY `fk_transaction_payment_transaction_detail` (`transaction_detail_id`),
  KEY `fk_transaction_payment_partner` (`partner_id`),
  CONSTRAINT `fk_transaction_payment_customer` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`),
  CONSTRAINT `fk_transaction_payment_partner` FOREIGN KEY (`partner_id`) REFERENCES `partner` (`id`),
  CONSTRAINT `fk_transaction_payment_transaction_detail` FOREIGN KEY (`transaction_detail_id`) REFERENCES `transaction_detail` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Dumping data for table kredit_mst.transaction_payment: ~0 rows (approximately)

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
