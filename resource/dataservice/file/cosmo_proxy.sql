/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.6.49-log : Database - cosmo_proxy
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`cosmo_proxy` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `cosmo_proxy`;

/*Table structure for table `route_define` */

DROP TABLE IF EXISTS `route_define`;

CREATE TABLE `route_define` (
    `id` varchar(20) NOT NULL,
    `path` varchar(128) DEFAULT NULL,
    `service_id` varchar(128) DEFAULT NULL,
    `url` varchar(80) DEFAULT NULL,
    `strip_prefix` tinyint(1) DEFAULT NULL,
    `retryable` tinyint(1) DEFAULT NULL,
    `api_name` varchar(128) DEFAULT NULL,
    `enabled` tinyint(1) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;