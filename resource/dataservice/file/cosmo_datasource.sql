/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.6.49-log : Database - cosmo_datasource
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE /*!32312 IF NOT EXISTS*/`cosmo_datasource` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `cosmo_datasource`;

/*Table structure for table `ds_data_source` */

DROP TABLE IF EXISTS `ds_data_source`;

CREATE TABLE `ds_data_source` (
    `id` bigint(15) NOT NULL AUTO_INCREMENT,
    `created_by` varchar(255) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(255) DEFAULT NULL COMMENT '更新者',
    `created_time` datetime DEFAULT NULL COMMENT '创建日期',
    `updated_time` datetime DEFAULT NULL COMMENT '更新日期',
    `name` varchar(128) NOT NULL COMMENT '数据源名称',
    `description` varchar(255) DEFAULT NULL COMMENT '数据源描述',
    `type` varchar(2) NOT NULL COMMENT '数据源类型',
    `data_position` varchar(2) DEFAULT NULL COMMENT '数据位置',
    `database_name` varchar(128) DEFAULT NULL COMMENT '数据库名称',
    `database_mode` varchar(2) DEFAULT NULL COMMENT '数据库模式(1:单机;2:集群)',
    `cluster_name` varchar(128) DEFAULT NULL COMMENT '集群名称',
    `share_mode` varchar(2) NOT NULL COMMENT '共享方式',
    `share_user_group_id` varchar(128) DEFAULT NULL COMMENT '共享用户组(多个用,分隔)',
    `owner` varchar(255) NOT NULL COMMENT '负责人(多个用,分隔)',
    `conn_url` varchar(1024) NOT NULL COMMENT '连接串',
    `user_name` varchar(64) DEFAULT NULL COMMENT '用户名',
    `password` text COMMENT '密码',
    `conn_status` varchar(2) DEFAULT NULL COMMENT '连接状态',
    `tenant_id` varchar(255) DEFAULT NULL COMMENT '租户id',
    `log_path` varchar(255) DEFAULT NULL COMMENT '日志路径(仅日志类型存储)',
    `hive_defaultFS` varchar(255) DEFAULT NULL COMMENT 'HIVE JDBC',
    `backup_JDBCUrl` varchar(500) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8 COMMENT='数据源';

/*Table structure for table `ubyfiuaa_ods_sdh_datasource_ds_data_source` */

DROP TABLE IF EXISTS `ubyfiuaa_ods_sdh_datasource_ds_data_source`;

CREATE TABLE `ubyfiuaa_ods_sdh_datasource_ds_data_source` (
    `id` bigint(15) NOT NULL AUTO_INCREMENT,
    `created_by` varchar(255) DEFAULT NULL COMMENT '创建者',
    `updated_by` varchar(255) DEFAULT NULL COMMENT '更新者',
    `created_time` datetime DEFAULT NULL COMMENT '创建日期',
    `updated_time` datetime DEFAULT NULL COMMENT '更新日期',
    `name` varchar(128) NOT NULL COMMENT '数据源名称',
    `description` varchar(255) DEFAULT NULL COMMENT '数据源描述',
    `type` varchar(2) NOT NULL COMMENT '数据源类型',
    `data_position` varchar(2) DEFAULT NULL COMMENT '数据位置',
    `database_name` varchar(128) DEFAULT NULL COMMENT '数据库名称',
    `database_mode` varchar(2) DEFAULT NULL COMMENT '数据库模式(1:单机;2:集群)',
    `cluster_name` varchar(128) DEFAULT NULL COMMENT '集群名称',
    `share_mode` varchar(2) NOT NULL COMMENT '共享方式',
    `share_user_group_id` varchar(128) DEFAULT NULL COMMENT '共享用户组(多个用,分隔)',
    `owner` varchar(255) NOT NULL COMMENT '负责人(多个用,分隔)',
    `conn_url` varchar(1024) NOT NULL COMMENT '连接串',
    `user_name` varchar(64) DEFAULT NULL COMMENT '用户名',
    `password` text COMMENT '密码',
    `conn_status` varchar(2) DEFAULT NULL COMMENT '连接状态',
    `tenant_id` varchar(255) DEFAULT NULL COMMENT '租户id',
    `log_path` varchar(255) DEFAULT NULL COMMENT '日志路径(仅日志类型存储)',
    `hive_defaultFS` varchar(255) DEFAULT NULL COMMENT 'HIVE JDBC',
    `backup_JDBCUrl` varchar(500) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='数据源';

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;