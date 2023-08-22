/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.6.49-log : Database - cosmo_gateway
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`cosmo_gateway` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `cosmo_gateway`;

/*Table structure for table `api_path` */

DROP TABLE IF EXISTS `api_path`;

CREATE TABLE `api_path` (
    `route_id` varchar(255) NOT NULL,
    `path` varchar(255) DEFAULT NULL,
    `service_path` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `app_rd_statistics_day` */

DROP TABLE IF EXISTS `app_rd_statistics_day`;

CREATE TABLE `app_rd_statistics_day` (
    `app_id` bigint(20) NOT NULL,
    `rd_id` varchar(50) NOT NULL,
    `time` int(11) NOT NULL,
    `fail_cnt` bigint(20) DEFAULT NULL,
    `total` bigint(20) DEFAULT NULL,
    PRIMARY KEY (`app_id`,`rd_id`,`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `app_rd_statistics_hour` */

DROP TABLE IF EXISTS `app_rd_statistics_hour`;

CREATE TABLE `app_rd_statistics_hour` (
    `app_id` bigint(20) NOT NULL,
    `rd_id` varchar(50) NOT NULL,
    `time` int(11) NOT NULL,
    `fail_cnt` bigint(20) DEFAULT NULL,
    `total` bigint(20) DEFAULT NULL,
    PRIMARY KEY (`app_id`,`rd_id`,`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `app_route_define_auth` */

DROP TABLE IF EXISTS `app_route_define_auth`;

CREATE TABLE `app_route_define_auth` (
    `app_id` bigint(20) NOT NULL,
    `rd_id` varchar(255) NOT NULL,
    `expire_time` bigint(20) DEFAULT NULL,
    `_status` tinyint(1) NOT NULL,
    PRIMARY KEY (`app_id`,`rd_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `application` */

DROP TABLE IF EXISTS `application`;

CREATE TABLE `application` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `creation_date` datetime DEFAULT NULL,
    `modified_date` datetime DEFAULT NULL,
    `is_active` tinyint(1) NOT NULL,
    `_key` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `secret` varchar(255) NOT NULL,
    `whitelist` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `UK_ltx1h78gnwdiol6kww33dmeu3` (`_key`),
    UNIQUE KEY `UK_lspnba25gpku3nx3oecprrx8c` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=95 DEFAULT CHARSET=utf8;

/*Table structure for table `application_route_defines` */

DROP TABLE IF EXISTS `application_route_defines`;

CREATE TABLE `application_route_defines` (
    `apps_id` bigint(20) NOT NULL,
    `route_defines_id` varchar(50) NOT NULL,
    PRIMARY KEY (`apps_id`,`route_defines_id`),
    KEY `FKpis86gykg4jikeux9ys3mbt6k` (`route_defines_id`),
    CONSTRAINT `FKp8v3qfu35lr902yye4k6p9y` FOREIGN KEY (`apps_id`) REFERENCES `application` (`id`),
    CONSTRAINT `FKpis86gykg4jikeux9ys3mbt6k` FOREIGN KEY (`route_defines_id`) REFERENCES `route_define` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `disable_route_define` */

DROP TABLE IF EXISTS `disable_route_define`;

CREATE TABLE `disable_route_define` (
    `_method` int(11) NOT NULL,
    `rd_id` varchar(255) NOT NULL,
    PRIMARY KEY (`_method`,`rd_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `rate_limit_rule` */

DROP TABLE IF EXISTS `rate_limit_rule`;

CREATE TABLE `rate_limit_rule` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `app_id` bigint(20) NOT NULL,
    `duration_secs` bigint(20) NOT NULL,
    `is_active` tinyint(1) NOT NULL,
    `_limit` int(11) NOT NULL,
    `rd_id` varchar(50) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `UKhjgtx2531yfv7exa0x19tjvds` (`app_id`,`rd_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `route_define` */

DROP TABLE IF EXISTS `route_define`;

CREATE TABLE `route_define` (
    `id` varchar(50) NOT NULL,
    `creation_date` datetime DEFAULT NULL,
    `modified_date` datetime DEFAULT NULL,
    `enable` tinyint(1) NOT NULL,
    `path` varchar(255) NOT NULL,
    `retryable` tinyint(1) NOT NULL,
    `service_id` varchar(80) DEFAULT NULL,
    `service_path` varchar(255) NOT NULL,
    `strip_prefix` tinyint(1) NOT NULL,
    `url` varchar(255) DEFAULT NULL,
    `wipe_service_path` tinyint(1) NOT NULL,
    `_order` int(11) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `UKdru7jh1c98pdrp3anhpci1aqs` (`path`,`service_path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `route_path` */

DROP TABLE IF EXISTS `route_path`;

CREATE TABLE `route_path` (
    `route_id` varchar(255) NOT NULL,
    `path` varchar(255) DEFAULT NULL,
    `service_path` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;