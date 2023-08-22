create user d3os identified by 'd3os@cosmoplat';
grant all on *.* to d3os@'%' identified by 'd3os@cosmoplat' with grant option;
flush privileges;

/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.6.49-log : Database - cosmo_api_manager
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE /*!32312 IF NOT EXISTS*/`cosmo_api_manager` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `cosmo_api_manager`;

/*Table structure for table `api` */

DROP TABLE IF EXISTS `api`;

CREATE TABLE `api` (
  `tenant_id` varchar(255) NOT NULL,
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `uname` varchar(255) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `b_path` varchar(255) DEFAULT NULL COMMENT '后端path',
  `ds_id` bigint(15) DEFAULT NULL COMMENT '数据源ID',
  `table_name` varchar(255) DEFAULT NULL COMMENT '表名',
  `method` varchar(20) DEFAULT NULL COMMENT '请求方法',
  `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NULL DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `paging` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否分页',
  `query_string` text,
  `type` varchar(20) DEFAULT NULL COMMENT 'api类型;original:生成API,inter:注册API,external:第三方',
  `status` tinyint(4) NOT NULL DEFAULT '2' COMMENT '1:online,2:offline,9:停用',
  `service_id` varchar(80) DEFAULT NULL,
  `service_path` varchar(255) DEFAULT NULL,
  `route_id` varchar(255) DEFAULT NULL COMMENT '网关路由ID',
  `market_status` tinyint(4) DEFAULT '0' COMMENT 'api市场状态;0:未上架,1:上架,9:被下架',
  `description` varchar(255) DEFAULT NULL,
  `ds_type` int(11) DEFAULT NULL COMMENT '数据源类型',
  `group_id` int(11) DEFAULT NULL COMMENT '分组ID',
  `market_test` tinyint(1) DEFAULT '0' COMMENT '发布市场后是否可以测试',
  `auth_type` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `path_2` (`path`,`type`,`tenant_id`,`method`,`service_path`),
  KEY `service_id` (`service_id`,`method`,`b_path`)
) ENGINE=InnoDB AUTO_INCREMENT=169 DEFAULT CHARSET=utf8;

/*Table structure for table `api_application_apply` */

DROP TABLE IF EXISTS `api_application_apply`;

CREATE TABLE `api_application_apply` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `application_id` int(11) NOT NULL COMMENT '应用ID',
  `api_id` bigint(11) NOT NULL COMMENT 'API_ID',
  `api_limit` bigint(11) DEFAULT NULL COMMENT 'api调用次数上限',
  `type` varchar(255) DEFAULT NULL,
  `start_time` datetime DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '结束时间',
  `comment` text COMMENT '申请说明',
  `description` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT 'NEW' COMMENT '审批状态;new,approved,reject,deleted',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NULL DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `api_ext` */

DROP TABLE IF EXISTS `api_ext`;

CREATE TABLE `api_ext` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) DEFAULT NULL,
  `type` varchar(100) DEFAULT NULL,
  `template` text COMMENT '模板',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;

/*Table structure for table `api_online_pre` */

DROP TABLE IF EXISTS `api_online_pre`;

CREATE TABLE `api_online_pre` (
  `tenant_id` varchar(255) NOT NULL,
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `uname` varchar(255) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `b_path` varchar(255) DEFAULT NULL COMMENT '后端path',
  `ds_id` bigint(15) DEFAULT NULL COMMENT '数据源ID',
  `table_name` varchar(255) DEFAULT NULL COMMENT '表名',
  `method` varchar(20) DEFAULT NULL COMMENT '请求方法',
  `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NULL DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `paging` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否分页',
  `query_string` text,
  `type` varchar(20) DEFAULT NULL COMMENT 'api类型;original:生成API,inter:注册API,external:第三方',
  `status` tinyint(4) NOT NULL DEFAULT '2' COMMENT '1:online,2:offline,9:停用',
  `service_id` varchar(80) DEFAULT NULL,
  `service_path` varchar(255) DEFAULT NULL,
  `route_id` varchar(255) DEFAULT NULL COMMENT '网关路由ID',
  `market_status` tinyint(4) DEFAULT '0' COMMENT 'api市场状态;0:未上架,1:上架,9:被下架',
  `description` varchar(255) DEFAULT NULL,
  `ds_type` int(11) DEFAULT NULL COMMENT '数据源类型',
  `group_id` int(11) DEFAULT NULL COMMENT '分组ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `path_2` (`path`,`type`,`tenant_id`,`method`),
  KEY `service_id` (`service_id`,`method`,`b_path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `api_param` */

DROP TABLE IF EXISTS `api_param`;

CREATE TABLE `api_param` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) DEFAULT NULL,
  `field` varchar(255) DEFAULT NULL COMMENT '字段名',
  `name` varchar(255) DEFAULT NULL COMMENT '字段别名',
  `position` varchar(25) DEFAULT NULL COMMENT 'path,query,body',
  `required` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否必输',
  `operator` varchar(100) DEFAULT NULL COMMENT 'eq|like',
  `type` varchar(100) DEFAULT NULL COMMENT 'path,body,heade',
  `def_value` varchar(255) DEFAULT NULL,
  `sample_value` varchar(255) DEFAULT NULL COMMENT '测试用例',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=104 DEFAULT CHARSET=utf8;

/*Table structure for table `api_param_ext` */

DROP TABLE IF EXISTS `api_param_ext`;

CREATE TABLE `api_param_ext` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) DEFAULT NULL,
  `template` text COMMENT 'json模板',
  `sample_value` text COMMENT '测试用例',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=253 DEFAULT CHARSET=utf8;

/*Table structure for table `api_param_ext_online_pre` */

DROP TABLE IF EXISTS `api_param_ext_online_pre`;

CREATE TABLE `api_param_ext_online_pre` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) DEFAULT NULL,
  `template` text COMMENT 'json模板',
  `sample_value` text COMMENT '测试用例',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `api_param_online_pre` */

DROP TABLE IF EXISTS `api_param_online_pre`;

CREATE TABLE `api_param_online_pre` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) DEFAULT NULL,
  `field` varchar(255) DEFAULT NULL COMMENT '字段名',
  `name` varchar(255) DEFAULT NULL COMMENT '字段别名',
  `position` varchar(25) DEFAULT NULL COMMENT 'path,query,body',
  `required` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否必输',
  `operator` varchar(100) DEFAULT NULL COMMENT 'eq|like',
  `type` varchar(100) DEFAULT NULL COMMENT 'path,body,heade',
  `def_value` varchar(255) DEFAULT NULL,
  `sample_value` varchar(255) DEFAULT NULL COMMENT '测试用例',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `api_path` */

DROP TABLE IF EXISTS `api_path`;

CREATE TABLE `api_path` (
  `id` bigint(11) NOT NULL,
  `path` varchar(255) DEFAULT NULL,
  `service_path` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `api_result` */

DROP TABLE IF EXISTS `api_result`;

CREATE TABLE `api_result` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) DEFAULT NULL,
  `field` varchar(255) DEFAULT NULL COMMENT '字段名',
  `name` varchar(255) DEFAULT NULL COMMENT '字段别名',
  `type` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `sample_value` varchar(255) DEFAULT NULL COMMENT '测试用例',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=654 DEFAULT CHARSET=utf8;

/*Table structure for table `api_result_bak` */

DROP TABLE IF EXISTS `api_result_bak`;

CREATE TABLE `api_result_bak` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) DEFAULT NULL,
  `field` varchar(255) DEFAULT NULL COMMENT '字段名',
  `name` varchar(255) DEFAULT NULL COMMENT '字段别名',
  `type` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `sample_value` varchar(255) DEFAULT NULL COMMENT '测试用例',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `api_result_copy` */

DROP TABLE IF EXISTS `api_result_copy`;

CREATE TABLE `api_result_copy` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) DEFAULT NULL,
  `field` varchar(255) DEFAULT NULL COMMENT '字段名',
  `name` varchar(255) DEFAULT NULL COMMENT '字段别名',
  `type` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `sample_value` varchar(255) DEFAULT NULL COMMENT '测试用例',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `api_result_ext` */

DROP TABLE IF EXISTS `api_result_ext`;

CREATE TABLE `api_result_ext` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) DEFAULT NULL,
  `template` text COMMENT 'json模板',
  `sample_value` mediumtext COMMENT '测试用例',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=253 DEFAULT CHARSET=utf8;

/*Table structure for table `api_result_ext_online_pre` */

DROP TABLE IF EXISTS `api_result_ext_online_pre`;

CREATE TABLE `api_result_ext_online_pre` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) DEFAULT NULL,
  `template` text COMMENT 'json模板',
  `sample_value` mediumtext COMMENT '测试用例',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `api_result_online_pre` */

DROP TABLE IF EXISTS `api_result_online_pre`;

CREATE TABLE `api_result_online_pre` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) DEFAULT NULL,
  `field` varchar(255) DEFAULT NULL COMMENT '字段名',
  `name` varchar(255) DEFAULT NULL COMMENT '字段别名',
  `type` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `sample_value` varchar(255) DEFAULT NULL COMMENT '测试用例',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `api_script` */

DROP TABLE IF EXISTS `api_script`;

CREATE TABLE `api_script` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) DEFAULT NULL,
  `code_type` varchar(30) DEFAULT NULL,
  `tenant_id` varchar(60) DEFAULT NULL,
  `source_code` text,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `sample_value` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `api_test_result` */

DROP TABLE IF EXISTS `api_test_result`;

CREATE TABLE `api_test_result` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `api_id` bigint(11) NOT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `request_param` text,
  `result` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `api_third_part_auth_param` */

DROP TABLE IF EXISTS `api_third_part_auth_param`;

CREATE TABLE `api_third_part_auth_param` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) DEFAULT NULL,
  `param_name` varchar(255) DEFAULT NULL COMMENT '参数名',
  `param_value` varchar(255) DEFAULT NULL COMMENT '参数值',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8;

/*Table structure for table `api_third_part_param` */

DROP TABLE IF EXISTS `api_third_part_param`;

CREATE TABLE `api_third_part_param` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) DEFAULT NULL,
  `param_name` varchar(255) DEFAULT NULL COMMENT '参数名',
  `param_value` varchar(255) DEFAULT NULL COMMENT '参数值',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `app_car_binding_1` */

DROP TABLE IF EXISTS `app_car_binding_1`;

CREATE TABLE `app_car_binding_1` (
  `customer_id` bigint(12) DEFAULT NULL,
  `owner_name` varchar(400) DEFAULT NULL,
  `owner_mp` varchar(4000) DEFAULT NULL,
  `id_card_number` varchar(4000) DEFAULT NULL,
  `province` varchar(200) DEFAULT NULL,
  `mailing_city` varchar(60) DEFAULT NULL,
  `org_name` varchar(150) DEFAULT NULL,
  `dealer_code` varchar(50) DEFAULT NULL,
  `create_date` timestamp NULL DEFAULT NULL,
  `update_date` timestamp NULL DEFAULT NULL,
  `vin` varchar(20) DEFAULT NULL,
  KEY `vin` (`vin`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `app_car_binding_2` */

DROP TABLE IF EXISTS `app_car_binding_2`;

CREATE TABLE `app_car_binding_2` (
  `customer_id` int(12) DEFAULT NULL,
  `owner_name` varchar(200) DEFAULT NULL,
  `owner_mp` varchar(4000) DEFAULT NULL,
  `id_card_number` varchar(4000) DEFAULT NULL,
  `province` varchar(200) DEFAULT NULL,
  `mailing_city` varchar(60) DEFAULT NULL,
  `org_name` varchar(150) DEFAULT NULL,
  `dealer_code` varchar(9) DEFAULT NULL,
  `create_date` timestamp NULL DEFAULT NULL,
  `update_date` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `app_car_binding_3` */

DROP TABLE IF EXISTS `app_car_binding_3`;

CREATE TABLE `app_car_binding_3` (
  `brand_name` varchar(60) DEFAULT NULL,
  `series_name` varchar(60) DEFAULT NULL,
  `model_name` varchar(150) DEFAULT NULL,
  `color_name` varchar(60) DEFAULT NULL,
  `delivery_date` timestamp NULL DEFAULT NULL,
  `vin` char(17) DEFAULT NULL,
  `material_number` varchar(15) DEFAULT NULL,
  `engine_number` varchar(64) DEFAULT NULL,
  `create_date` timestamp NULL DEFAULT NULL,
  `update_date` timestamp NULL DEFAULT NULL,
  `brand_id` varchar(20) DEFAULT NULL,
  KEY `vin` (`vin`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `app_car_binding_4` */

DROP TABLE IF EXISTS `app_car_binding_4`;

CREATE TABLE `app_car_binding_4` (
  `brand_name` varchar(60) DEFAULT NULL,
  `owner_mp` varchar(4000) DEFAULT NULL,
  `series_name` varchar(60) DEFAULT NULL,
  `model_name` varchar(150) DEFAULT NULL,
  `color_name` varchar(60) DEFAULT NULL,
  `delivery_date` timestamp NULL DEFAULT NULL,
  `vin` char(17) DEFAULT NULL,
  `material_number` varchar(15) DEFAULT NULL,
  `create_date` timestamp NULL DEFAULT NULL,
  `update_date` timestamp NULL DEFAULT NULL,
  `brand_id` varchar(20) DEFAULT NULL,
  `engine_number` varchar(40) DEFAULT NULL
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Table structure for table `app_rd_statistics_day_copy1` */

DROP TABLE IF EXISTS `app_rd_statistics_day_copy1`;

CREATE TABLE `app_rd_statistics_day_copy1` (
  `app_id` bigint(20) NOT NULL,
  `rd_id` varchar(50) NOT NULL,
  `time` int(11) NOT NULL,
  `fail_cnt` bigint(20) DEFAULT NULL,
  `total` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`app_id`,`rd_id`,`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Table structure for table `app_rd_statistics_hour` */

DROP TABLE IF EXISTS `app_rd_statistics_hour`;

CREATE TABLE `app_rd_statistics_hour` (
  `app_id` bigint(20) NOT NULL,
  `rd_id` varchar(50) NOT NULL,
  `time` int(11) NOT NULL,
  `fail_cnt` bigint(20) DEFAULT NULL,
  `total` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`app_id`,`rd_id`,`time`),
  KEY `app_id` (`app_id`,`rd_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Table structure for table `app_rd_statistics_hour_copy1` */

DROP TABLE IF EXISTS `app_rd_statistics_hour_copy1`;

CREATE TABLE `app_rd_statistics_hour_copy1` (
  `app_id` bigint(20) NOT NULL,
  `rd_id` varchar(50) NOT NULL,
  `time` int(11) NOT NULL,
  `fail_cnt` bigint(20) DEFAULT NULL,
  `total` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`app_id`,`rd_id`,`time`),
  KEY `app_id` (`app_id`,`rd_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Table structure for table `application` */

DROP TABLE IF EXISTS `application`;

CREATE TABLE `application` (
  `tenant_id` varchar(255) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `gateway_id` bigint(20) NOT NULL COMMENT '网关应用ID',
  `name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NULL DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `app_key` varchar(255) DEFAULT NULL COMMENT '网关APP_KEY',
  `is_default` int(2) DEFAULT '0' COMMENT '是否为租户默认应用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

/*Table structure for table `application_api_relation` */

DROP TABLE IF EXISTS `application_api_relation`;

CREATE TABLE `application_api_relation` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `application_id` int(11) NOT NULL COMMENT '应用ID',
  `api_id` bigint(11) NOT NULL,
  `api_limit` bigint(11) DEFAULT NULL COMMENT '调用次数限制',
  `type` varchar(255) DEFAULT NULL,
  `start_time` datetime DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '结束时间',
  `description` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NULL DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=110 DEFAULT CHARSET=utf8;

/*Table structure for table `billing_plan` */

DROP TABLE IF EXISTS `billing_plan`;

CREATE TABLE `billing_plan` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `request_limit` int(11) DEFAULT NULL COMMENT '请求限制，-1: 无限制',
  `desciption` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `book` */

DROP TABLE IF EXISTS `book`;

CREATE TABLE `book` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(50) DEFAULT NULL,
  `pageCount` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `channel` */

DROP TABLE IF EXISTS `channel`;

CREATE TABLE `channel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `channel_name` varchar(255) NOT NULL COMMENT '渠道名称',
  `channel_type` varchar(255) DEFAULT NULL COMMENT '渠道类型',
  `description` varchar(255) DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `tree_channel_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `channel_api_mapping` */

DROP TABLE IF EXISTS `channel_api_mapping`;

CREATE TABLE `channel_api_mapping` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) NOT NULL,
  `channel_id` int(11) NOT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `contract` */

DROP TABLE IF EXISTS `contract`;

CREATE TABLE `contract` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `partner_id` int(11) NOT NULL,
  `billing_plan_id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `start_at` timestamp NULL DEFAULT NULL,
  `end_at` timestamp NULL DEFAULT NULL,
  `amount` int(11) DEFAULT NULL COMMENT '合同金额',
  `party_a` varchar(255) DEFAULT NULL COMMENT '甲方负责人',
  `party_b` varchar(255) DEFAULT NULL COMMENT '乙方负责人',
  `content` text,
  `desciption` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_cboard_relation` */

DROP TABLE IF EXISTS `data_cboard_relation`;

CREATE TABLE `data_cboard_relation` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `dataset_id` int(11) DEFAULT NULL COMMENT '数据产品ID',
  `cboard_id` bigint(20) DEFAULT NULL COMMENT '报表ID',
  `created_by` varchar(255) DEFAULT NULL,
  `created_time` timestamp NULL DEFAULT NULL,
  `updated_by` varchar(50) DEFAULT NULL,
  `updated_time` timestamp NULL DEFAULT NULL,
  `path` varchar(50) DEFAULT NULL COMMENT '路径',
  `name` varchar(50) DEFAULT NULL COMMENT '数据报表名称',
  `tenant_id` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='数据集绑定的报表关系表';

/*Table structure for table `data_cboard_visit` */

DROP TABLE IF EXISTS `data_cboard_visit`;

CREATE TABLE `data_cboard_visit` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) DEFAULT NULL COMMENT '访问用户id',
  `cboard_id` bigint(20) DEFAULT NULL COMMENT '被访问的报表id',
  `created_by` varchar(50) DEFAULT NULL COMMENT '访问人姓名',
  `created_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '访问时间',
  `updated_by` varchar(50) DEFAULT NULL,
  `updated_time` timestamp NULL DEFAULT NULL,
  `tenant_id` varchar(50) DEFAULT NULL COMMENT '访问者租户id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='申请人访问报表记录表';

/*Table structure for table `data_products_apply` */

DROP TABLE IF EXISTS `data_products_apply`;

CREATE TABLE `data_products_apply` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(60) DEFAULT NULL,
  `name` varchar(60) DEFAULT NULL,
  `channel_ids` varchar(255) DEFAULT NULL,
  `tenant_ids` varchar(255) DEFAULT NULL,
  `ds_type` varchar(60) DEFAULT NULL,
  `ds_id` int(11) DEFAULT NULL,
  `data_set_id` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `data_dic_path` varchar(255) DEFAULT NULL,
  `data_dic_name` varchar(255) DEFAULT NULL,
  `data_file_path` varchar(255) DEFAULT NULL,
  `data_file_name` varchar(255) DEFAULT NULL,
  `sample_data_path` varchar(255) DEFAULT NULL,
  `sample_data_name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` varchar(60) DEFAULT NULL,
  `updated_by` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `data_set_id` (`data_set_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_products_apply_copy1` */

DROP TABLE IF EXISTS `data_products_apply_copy1`;

CREATE TABLE `data_products_apply_copy1` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(60) DEFAULT NULL,
  `name` varchar(60) DEFAULT NULL,
  `channel_ids` varchar(255) DEFAULT NULL,
  `tenant_ids` varchar(255) DEFAULT NULL,
  `ds_type` varchar(60) DEFAULT NULL,
  `ds_id` int(11) DEFAULT NULL,
  `data_set_id` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `data_dic_path` varchar(255) DEFAULT NULL,
  `data_dic_name` varchar(255) DEFAULT NULL,
  `data_file_path` varchar(255) DEFAULT NULL,
  `data_file_name` varchar(255) DEFAULT NULL,
  `sample_data_path` varchar(255) DEFAULT NULL,
  `sample_data_name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` varchar(60) DEFAULT NULL,
  `updated_by` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `data_set_id` (`data_set_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_products_apply_keyword` */

DROP TABLE IF EXISTS `data_products_apply_keyword`;

CREATE TABLE `data_products_apply_keyword` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_products_apply_id` int(11) NOT NULL COMMENT '数据产品申请id',
  `key_word` varchar(255) DEFAULT NULL COMMENT '关键字',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_products_apply_keyword_copy1` */

DROP TABLE IF EXISTS `data_products_apply_keyword_copy1`;

CREATE TABLE `data_products_apply_keyword_copy1` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_products_apply_id` int(11) NOT NULL COMMENT '数据产品申请id',
  `key_word` varchar(255) DEFAULT NULL COMMENT '关键字',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_products_ext` */

DROP TABLE IF EXISTS `data_products_ext`;

CREATE TABLE `data_products_ext` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(60) DEFAULT NULL,
  `dp_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `ds_type` int(11) DEFAULT NULL,
  `ds_id` int(11) DEFAULT NULL,
  `ds_name` varchar(255) DEFAULT NULL,
  `method` varchar(255) DEFAULT NULL,
  `table_name` varchar(255) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_time` datetime DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_products_ext_copy1` */

DROP TABLE IF EXISTS `data_products_ext_copy1`;

CREATE TABLE `data_products_ext_copy1` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(60) DEFAULT NULL,
  `dp_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `ds_type` int(11) DEFAULT NULL,
  `ds_id` int(11) DEFAULT NULL,
  `ds_name` varchar(255) DEFAULT NULL,
  `method` varchar(255) DEFAULT NULL,
  `table_name` varchar(255) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_time` datetime DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_rinse_tag` */

DROP TABLE IF EXISTS `data_rinse_tag`;

CREATE TABLE `data_rinse_tag` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(100) NOT NULL COMMENT '名称',
  `business_table` varchar(64) NOT NULL COMMENT '表名',
  `field_name` varchar(100) DEFAULT NULL COMMENT '字段名称',
  `state_code` varchar(8) DEFAULT NULL COMMENT '字段类型',
  `tag_type` varchar(4) DEFAULT NULL COMMENT '标签类型',
  `definition` varchar(512) DEFAULT NULL COMMENT '标签定义',
  `number` varchar(32) DEFAULT NULL COMMENT '编码',
  `usage_scenario` varchar(512) DEFAULT NULL COMMENT '使用场景',
  `generative_rule` varchar(256) DEFAULT NULL COMMENT '生成规则',
  `first_catalog` varchar(32) DEFAULT NULL COMMENT '一级目录',
  `catalog` varchar(32) DEFAULT NULL COMMENT '二级目录',
  `type_catalog` varchar(32) DEFAULT NULL COMMENT '类型目录',
  `update_mode` varchar(64) DEFAULT NULL COMMENT '更新方式',
  `open_state` varchar(8) DEFAULT NULL COMMENT '开放状态',
  `principal` varchar(64) DEFAULT NULL COMMENT '标签负责人',
  `principal_id` varchar(64) DEFAULT NULL COMMENT '标签负责人ID',
  `status` varchar(64) DEFAULT NULL COMMENT '审批状态',
  `data_source_id` int(11) DEFAULT NULL COMMENT '数据源ID',
  `online_status` varchar(4) DEFAULT NULL COMMENT '上线状态',
  `create_by` varchar(64) NOT NULL COMMENT '创建者',
  `create_date` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(64) NOT NULL COMMENT '更新者',
  `update_date` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `version` varchar(32) DEFAULT '0' COMMENT '版本号',
  `del_flag` char(1) NOT NULL DEFAULT '0' COMMENT '删除标记',
  `tag_value` varchar(512) DEFAULT '' COMMENT '标签值',
  `remark` varchar(512) DEFAULT '' COMMENT '备注',
  `is_arithmetic` varchar(4) DEFAULT 'N' COMMENT '是否算法标签 （Y：是；N：不是）',
  `feedback_id` varchar(32) DEFAULT '' COMMENT '反馈ID',
  `sort` int(8) NOT NULL COMMENT '排序',
  `required_fields` varchar(1) DEFAULT '0' COMMENT '必填字段',
  `column_accuracy` varchar(20) DEFAULT NULL COMMENT '字段精度',
  `column_max_value` varchar(20) DEFAULT NULL COMMENT '字段最大值',
  `column_min_value` varchar(20) DEFAULT NULL COMMENT '字段最小值',
  `number_calls` int(11) DEFAULT '0' COMMENT '调用次数',
  PRIMARY KEY (`id`),
  KEY `data_rinse_tag_business_table` (`business_table`) USING BTREE,
  KEY `data_rinse_tag_field_name` (`field_name`) USING BTREE,
  KEY `data_rinse_tag_del_flag` (`del_flag`) USING BTREE,
  KEY `data_rinse_tag_name` (`name`) USING BTREE,
  KEY `business_table` (`business_table`,`field_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标签清洗表';

/*Table structure for table `data_rinse_tag_detail` */

DROP TABLE IF EXISTS `data_rinse_tag_detail`;

CREATE TABLE `data_rinse_tag_detail` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `field_name` varchar(100) NOT NULL COMMENT '字段名称',
  `sub_value` varchar(255) NOT NULL COMMENT '字段值',
  `rinse_tag_id` int(11) NOT NULL COMMENT '元标签ID',
  `business_table` varchar(64) NOT NULL COMMENT '表名',
  `data_source_id` int(11) DEFAULT NULL COMMENT '数据源ID',
  `value` text COMMENT '字段名称',
  `create_by` varchar(64) NOT NULL COMMENT '创建者',
  `create_date` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(64) NOT NULL COMMENT '更新者',
  `update_date` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `version` varchar(32) DEFAULT '0' COMMENT '版本号',
  `del_flag` char(1) NOT NULL DEFAULT '0' COMMENT '删除标记',
  KEY `business_table` (`business_table`,`field_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标签清洗表明细';

/*Table structure for table `data_rinse_tag_detail_tmp` */

DROP TABLE IF EXISTS `data_rinse_tag_detail_tmp`;

CREATE TABLE `data_rinse_tag_detail_tmp` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `field_name` varchar(100) NOT NULL COMMENT '字段名称',
  `sub_value` varchar(255) NOT NULL COMMENT '字段值',
  `rinse_tag_id` int(11) NOT NULL COMMENT '元标签ID',
  `business_table` varchar(64) NOT NULL COMMENT '表名',
  `data_source_id` int(11) DEFAULT NULL COMMENT '数据源ID',
  `value` text COMMENT '字段名称',
  `create_by` varchar(64) NOT NULL COMMENT '创建者',
  `create_date` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(64) NOT NULL COMMENT '更新者',
  `update_date` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `version` varchar(32) DEFAULT '0' COMMENT '版本号',
  `del_flag` char(1) NOT NULL DEFAULT '0' COMMENT '删除标记'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标签清洗表明细';

/*Table structure for table `data_set` */

DROP TABLE IF EXISTS `data_set`;

CREATE TABLE `data_set` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_set_name` varchar(255) NOT NULL,
  `data_set` text COMMENT '数据集(json格式)',
  `description` text,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_api` */

DROP TABLE IF EXISTS `data_set_api`;

CREATE TABLE `data_set_api` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `api_id` int(11) NOT NULL COMMENT 'API_ID',
  `description` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_api_apply` */

DROP TABLE IF EXISTS `data_set_api_apply`;

CREATE TABLE `data_set_api_apply` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` bigint(20) DEFAULT NULL COMMENT '序列号',
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `api_id` int(11) NOT NULL COMMENT 'API_ID',
  `status` varchar(255) DEFAULT 'NEW' COMMENT '审批状态;new,approved,reject,deleted',
  `description` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  `type` varchar(50) DEFAULT NULL COMMENT '是否是cboard',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_api_apply_copy1` */

DROP TABLE IF EXISTS `data_set_api_apply_copy1`;

CREATE TABLE `data_set_api_apply_copy1` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` bigint(20) DEFAULT NULL COMMENT '序列号',
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `api_id` int(11) NOT NULL COMMENT 'API_ID',
  `status` varchar(255) DEFAULT 'NEW' COMMENT '审批状态;new,approved,reject,deleted',
  `description` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_api_copy1` */

DROP TABLE IF EXISTS `data_set_api_copy1`;

CREATE TABLE `data_set_api_copy1` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `api_id` int(11) NOT NULL COMMENT 'API_ID',
  `description` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_api_copy2` */

DROP TABLE IF EXISTS `data_set_api_copy2`;

CREATE TABLE `data_set_api_copy2` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `api_id` int(11) NOT NULL COMMENT 'API_ID',
  `description` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_channel` */

DROP TABLE IF EXISTS `data_set_channel`;

CREATE TABLE `data_set_channel` (
  `tenant_id` varchar(50) NOT NULL,
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `channel_id` int(11) NOT NULL,
  PRIMARY KEY (`tenant_id`,`data_set_id`,`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_channel_copy1` */

DROP TABLE IF EXISTS `data_set_channel_copy1`;

CREATE TABLE `data_set_channel_copy1` (
  `tenant_id` varchar(50) NOT NULL,
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `channel_id` int(11) NOT NULL,
  PRIMARY KEY (`tenant_id`,`data_set_id`,`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_channel_copy2` */

DROP TABLE IF EXISTS `data_set_channel_copy2`;

CREATE TABLE `data_set_channel_copy2` (
  `tenant_id` varchar(50) NOT NULL,
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `channel_id` int(11) NOT NULL,
  PRIMARY KEY (`tenant_id`,`data_set_id`,`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_copy1` */

DROP TABLE IF EXISTS `data_set_copy1`;

CREATE TABLE `data_set_copy1` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_set_name` varchar(255) NOT NULL,
  `channel_id` int(11) NOT NULL DEFAULT '0',
  `data_set` text COMMENT '数据集(json格式)',
  `description` text,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_copy2` */

DROP TABLE IF EXISTS `data_set_copy2`;

CREATE TABLE `data_set_copy2` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_set_name` varchar(255) NOT NULL,
  `data_set` text COMMENT '数据集(json格式)',
  `description` text,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_dic` */

DROP TABLE IF EXISTS `data_set_dic`;

CREATE TABLE `data_set_dic` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `col_name` varchar(255) DEFAULT NULL,
  `col_type` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_dic_copy1` */

DROP TABLE IF EXISTS `data_set_dic_copy1`;

CREATE TABLE `data_set_dic_copy1` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `col_name` varchar(255) DEFAULT NULL,
  `col_type` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_enterprise_user` */

DROP TABLE IF EXISTS `data_set_enterprise_user`;

CREATE TABLE `data_set_enterprise_user` (
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `enterprise_user_id` varchar(50) NOT NULL DEFAULT '' COMMENT '数据市场企业会员表ID',
  PRIMARY KEY (`data_set_id`,`enterprise_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_enterprise_user_copy1` */

DROP TABLE IF EXISTS `data_set_enterprise_user_copy1`;

CREATE TABLE `data_set_enterprise_user_copy1` (
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `enterprise_user_id` varchar(50) NOT NULL DEFAULT '' COMMENT '数据市场企业会员表ID',
  PRIMARY KEY (`data_set_id`,`enterprise_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_keyword` */

DROP TABLE IF EXISTS `data_set_keyword`;

CREATE TABLE `data_set_keyword` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `key_word` varchar(255) DEFAULT NULL COMMENT '关键字',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_keyword_copy1` */

DROP TABLE IF EXISTS `data_set_keyword_copy1`;

CREATE TABLE `data_set_keyword_copy1` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `key_word` varchar(255) DEFAULT NULL COMMENT '关键字',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_order` */

DROP TABLE IF EXISTS `data_set_order`;

CREATE TABLE `data_set_order` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_num` varchar(60) DEFAULT NULL COMMENT '订单编号',
  `application_id` int(11) NOT NULL COMMENT '应用ID',
  `data_set_id` int(11) NOT NULL COMMENT '数据集id',
  `api_limit` bigint(11) DEFAULT NULL COMMENT 'api调用次数上限',
  `type` varchar(255) DEFAULT NULL,
  `start_time` datetime DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '结束时间',
  `comment` text COMMENT '申请说明',
  `description` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT 'NEW' COMMENT '审批状态;new,approved,reject,deleted',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  `apply_sign_file_id` int(11) DEFAULT NULL COMMENT '数据申请方签名文件id',
  `provider_sign_file_id` int(11) DEFAULT NULL COMMENT '数据提供方签名文件id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_order_api_relation` */

DROP TABLE IF EXISTS `data_set_order_api_relation`;

CREATE TABLE `data_set_order_api_relation` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` int(11) DEFAULT NULL,
  `api_id` bigint(12) DEFAULT NULL,
  `tenant_id` varchar(200) DEFAULT NULL COMMENT '租户id',
  `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_order_api_relation_copy1` */

DROP TABLE IF EXISTS `data_set_order_api_relation_copy1`;

CREATE TABLE `data_set_order_api_relation_copy1` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` int(11) DEFAULT NULL,
  `api_id` bigint(12) DEFAULT NULL,
  `tenant_id` varchar(200) DEFAULT NULL COMMENT '租户id',
  `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_order_cboard_relation` */

DROP TABLE IF EXISTS `data_set_order_cboard_relation`;

CREATE TABLE `data_set_order_cboard_relation` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` int(11) DEFAULT NULL,
  `cboard_id` bigint(12) DEFAULT NULL,
  `tenant_id` varchar(200) DEFAULT NULL COMMENT '租户id',
  `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='申请的订单报表关系表';

/*Table structure for table `data_set_order_result_relation` */

DROP TABLE IF EXISTS `data_set_order_result_relation`;

CREATE TABLE `data_set_order_result_relation` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` int(11) NOT NULL COMMENT '结果与订单的关系id',
  `application_id` int(11) DEFAULT NULL,
  `data_set_id` int(11) NOT NULL COMMENT '数据集与api关系ID',
  `api_id` int(11) NOT NULL COMMENT 'API_ID',
  `api_result_id` int(11) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_order_result_relation_copy1` */

DROP TABLE IF EXISTS `data_set_order_result_relation_copy1`;

CREATE TABLE `data_set_order_result_relation_copy1` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `order_id` int(11) NOT NULL COMMENT '结果与订单的关系id',
  `application_id` int(11) DEFAULT NULL,
  `data_set_id` int(11) NOT NULL COMMENT '数据集与api关系ID',
  `api_id` int(11) NOT NULL COMMENT 'API_ID',
  `api_result_id` int(11) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_result` */

DROP TABLE IF EXISTS `data_set_result`;

CREATE TABLE `data_set_result` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_set_id` int(11) NOT NULL COMMENT '数据集与api关系ID',
  `application_id` int(11) DEFAULT NULL,
  `api_id` int(11) NOT NULL COMMENT 'API_ID',
  `api_result_id` int(11) DEFAULT NULL COMMENT 'API返回字段id',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_result_copy1` */

DROP TABLE IF EXISTS `data_set_result_copy1`;

CREATE TABLE `data_set_result_copy1` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_set_id` int(11) NOT NULL COMMENT '数据集与api关系ID',
  `application_id` int(11) DEFAULT NULL,
  `api_id` int(11) NOT NULL COMMENT 'API_ID',
  `api_result_id` int(11) DEFAULT NULL COMMENT 'API返回字段id',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_sample` */

DROP TABLE IF EXISTS `data_set_sample`;

CREATE TABLE `data_set_sample` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_set_id` int(11) NOT NULL,
  `data_key` varchar(255) NOT NULL,
  `data_value` varchar(255) NOT NULL,
  `description` text,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  `line_no` int(11) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_sample_copy1` */

DROP TABLE IF EXISTS `data_set_sample_copy1`;

CREATE TABLE `data_set_sample_copy1` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_set_id` int(11) NOT NULL,
  `data_key` varchar(255) NOT NULL,
  `data_value` varchar(255) NOT NULL,
  `description` text,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  `line_no` int(11) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_sort` */

DROP TABLE IF EXISTS `data_set_sort`;

CREATE TABLE `data_set_sort` (
  `data_set_id` int(11) NOT NULL COMMENT '和dataset 关联的字段',
  `visit_rank` int(11) DEFAULT NULL,
  `click_rank` int(11) DEFAULT NULL,
  `final_rank` float(11,2) DEFAULT NULL COMMENT '最终排名',
  PRIMARY KEY (`data_set_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_sort_copy1` */

DROP TABLE IF EXISTS `data_set_sort_copy1`;

CREATE TABLE `data_set_sort_copy1` (
  `data_set_id` int(11) NOT NULL COMMENT '和dataset 关联的字段',
  `visit_rank` int(11) DEFAULT NULL,
  `click_rank` int(11) DEFAULT NULL,
  `final_rank` float(11,2) DEFAULT NULL COMMENT '最终排名',
  PRIMARY KEY (`data_set_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_visit` */

DROP TABLE IF EXISTS `data_set_visit`;

CREATE TABLE `data_set_visit` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_set_visit_copy1` */

DROP TABLE IF EXISTS `data_set_visit_copy1`;

CREATE TABLE `data_set_visit_copy1` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `data_set_id` int(11) NOT NULL COMMENT '数据集ID',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_source` */

DROP TABLE IF EXISTS `data_source`;

CREATE TABLE `data_source` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` bigint(15) NOT NULL,
  `name` varchar(25) DEFAULT NULL,
  `url` varchar(250) DEFAULT NULL,
  `user_name` varchar(250) DEFAULT NULL,
  `password` text,
  `type` varchar(20) DEFAULT NULL,
  `service_url` varchar(500) DEFAULT NULL,
  `service_port` int(5) NOT NULL,
  `service_status` int(1) NOT NULL,
  `container_id` varchar(255) DEFAULT NULL,
  `definitions` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `data_tag_export` */

DROP TABLE IF EXISTS `data_tag_export`;

CREATE TABLE `data_tag_export` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` varchar(20) CHARACTER SET utf8 NOT NULL COMMENT '主键',
  `type_name` varchar(255) DEFAULT NULL COMMENT '标签类型',
  `name` varchar(50) DEFAULT NULL COMMENT '数据标签名',
  `description` varchar(255) DEFAULT NULL COMMENT '标签值描述',
  `rule` varchar(255) DEFAULT NULL COMMENT '生成规则',
  `belong_module_all_name` varchar(255) DEFAULT NULL COMMENT '所属模块负责人',
  `class_all_name` varchar(50) DEFAULT NULL COMMENT '多级类目',
  `rule_label_name` varchar(255) DEFAULT NULL COMMENT '规则样式标识',
  `table_type_name` varchar(150) DEFAULT NULL COMMENT '使用场景数据',
  `comments` varchar(255) DEFAULT NULL COMMENT '备注',
  `generate_method_name` varchar(255) DEFAULT NULL COMMENT '生成方式',
  `deploy_status_name` varchar(50) DEFAULT NULL COMMENT '部署状态',
  `update_rate` varchar(255) DEFAULT NULL COMMENT '更新频率',
  `is_algorithm_name` varchar(255) DEFAULT NULL COMMENT '是否算法标签',
  `accuracy` int(20) DEFAULT NULL COMMENT '标签值精度',
  `created_st_time` varchar(100) DEFAULT NULL COMMENT '创建时间(工具链平台)',
  `updated_st_time` varchar(100) DEFAULT NULL COMMENT '更新时间(工具链平台)',
  `position_name` varchar(255) DEFAULT NULL COMMENT '数据源位置',
  `column_desc` varchar(255) DEFAULT NULL COMMENT '字段描述',
  `database_name` varchar(255) DEFAULT NULL COMMENT '所属库名',
  `table_name` varchar(50) DEFAULT NULL COMMENT '所属表名',
  `column_name` varchar(50) DEFAULT NULL COMMENT '所属字段名',
  `last_fields_all_name` varchar(255) DEFAULT NULL COMMENT '上一级字段',
  `ip` varchar(255) DEFAULT NULL COMMENT 'ip',
  `port` varchar(255) DEFAULT NULL COMMENT '端口',
  `is_visible` int(1) DEFAULT NULL COMMENT '1:全局可见 0:全局不可见',
  `created_time` datetime DEFAULT NULL COMMENT '创建时间',
  `created_by` varchar(255) DEFAULT NULL COMMENT '创建人',
  `updated_time` datetime DEFAULT NULL COMMENT '更新时间',
  `updated_by` varchar(255) DEFAULT NULL COMMENT '更新人',
  `is_deleted` int(1) DEFAULT NULL COMMENT '删除标志(是否删除  1：是  0：否)',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `table_name` (`table_name`,`column_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

/*Table structure for table `enterprise_data_set_relation` */

DROP TABLE IF EXISTS `enterprise_data_set_relation`;

CREATE TABLE `enterprise_data_set_relation` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `enterprise_user_id` int(11) DEFAULT NULL,
  `data_set_id` int(11) DEFAULT NULL,
  `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `enterprise_user_id` (`enterprise_user_id`,`data_set_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `enterprise_entry_apply` */

DROP TABLE IF EXISTS `enterprise_entry_apply`;

CREATE TABLE `enterprise_entry_apply` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(60) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL COMMENT '企业名称',
  `contacts` varchar(255) DEFAULT NULL COMMENT '联系人名称',
  `mobile` varchar(255) DEFAULT NULL COMMENT '联系人电话',
  `email` varchar(255) DEFAULT NULL COMMENT '联系人邮箱',
  `agreement_path` varchar(255) DEFAULT NULL COMMENT '文件路径',
  `agreement_name` varchar(255) DEFAULT NULL COMMENT '文件名',
  `status` varchar(60) DEFAULT NULL COMMENT '状态new,approved,reject,deleted',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `comment` varchar(255) DEFAULT NULL COMMENT '审批描述',
  `created_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `enterprise_user` */

DROP TABLE IF EXISTS `enterprise_user`;

CREATE TABLE `enterprise_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `enterprise_user_id` varchar(200) DEFAULT NULL COMMENT '企业会员名',
  `enterprise_user_name` varchar(200) DEFAULT NULL,
  `circulation_management_user_id` varchar(200) DEFAULT NULL,
  `circulation_management_user_name` varchar(200) DEFAULT NULL COMMENT '流通管理员id',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `tenant_id` varchar(200) DEFAULT NULL COMMENT '租户id',
  `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `groups` */

DROP TABLE IF EXISTS `groups`;

CREATE TABLE `groups` (
  `tenant_id` varchar(255) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group_name` varchar(255) NOT NULL COMMENT '分组名称',
  `created_by` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `tree_group_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8;

/*Table structure for table `instance_api_relation` */

DROP TABLE IF EXISTS `instance_api_relation`;

CREATE TABLE `instance_api_relation` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `ds_id` bigint(20) NOT NULL,
  `instance_name` varchar(255) DEFAULT NULL,
  `service_port` int(11) DEFAULT NULL,
  `service_status` int(11) DEFAULT NULL,
  `service_url` varchar(255) DEFAULT NULL,
  `container_id` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='实例数据源关系表';

/*Table structure for table `mail_record` */

DROP TABLE IF EXISTS `mail_record`;

CREATE TABLE `mail_record` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(60) DEFAULT NULL COMMENT '类型',
  `recipient_to` varchar(600) DEFAULT NULL COMMENT '接收人',
  `recipient_cc` varchar(600) DEFAULT NULL COMMENT '抄送人',
  `status` int(6) DEFAULT '0' COMMENT '状态',
  `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `mail_relation` */

DROP TABLE IF EXISTS `mail_relation`;

CREATE TABLE `mail_relation` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(60) DEFAULT NULL COMMENT '业务类型',
  `realtion_id` varchar(60) DEFAULT NULL COMMENT '关联业务id',
  `applicant_email` varchar(255) DEFAULT NULL COMMENT '申请人邮箱',
  `status` int(5) DEFAULT NULL COMMENT '邮件状态 0待发送 1已发送 -1异常发送',
  `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `mail_template` */

DROP TABLE IF EXISTS `mail_template`;

CREATE TABLE `mail_template` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(255) DEFAULT NULL COMMENT '类型',
  `subject` varchar(255) DEFAULT NULL COMMENT '主题',
  `message` varchar(800) DEFAULT NULL COMMENT '内容',
  `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `market_user` */

DROP TABLE IF EXISTS `market_user`;

CREATE TABLE `market_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(80) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '用户名',
  `telephone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `email` varchar(200) DEFAULT NULL COMMENT '邮箱',
  `tenant_id` varchar(200) DEFAULT NULL COMMENT '租户id',
  `tenant_name` varchar(200) DEFAULT NULL COMMENT '租户',
  `department` varchar(255) DEFAULT NULL COMMENT '部门',
  `effective_days` int(11) DEFAULT NULL COMMENT '有效天数',
  `effective_date` datetime DEFAULT NULL COMMENT '生效时间',
  `expire_date` datetime DEFAULT NULL COMMENT '到期时间',
  `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `partner` */

DROP TABLE IF EXISTS `partner`;

CREATE TABLE `partner` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(255) NOT NULL COMMENT '租户id',
  `name` varchar(255) DEFAULT NULL,
  `desciption` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `push_data` */

DROP TABLE IF EXISTS `push_data`;

CREATE TABLE `push_data` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `Mobile_phone_brand` varchar(50) DEFAULT NULL COMMENT '手机品牌',
  `Mobile_phone_model` varchar(50) DEFAULT NULL COMMENT '手机机型',
  `operating_system` varchar(50) DEFAULT NULL COMMENT '操作系统',
  `ssystem_language` varchar(50) DEFAULT NULL COMMENT '系统语言',
  `Operator` varchar(50) DEFAULT NULL COMMENT '运营商',
  `Resident_Province` varchar(50) DEFAULT NULL COMMENT '常驻省份',
  `Resident_City` varchar(50) DEFAULT NULL COMMENT '常驻城市',
  `Hometown` varchar(50) DEFAULT NULL COMMENT '家乡地',
  `age` varchar(50) DEFAULT NULL COMMENT '年龄',
  `sex` varchar(50) DEFAULT NULL COMMENT '性别',
  `Consumption_level` varchar(50) DEFAULT NULL COMMENT '消费水平',
  `local_tyrant` varchar(50) DEFAULT NULL COMMENT '土豪',
  `Home_buyer` varchar(50) DEFAULT NULL COMMENT '家庭主购者',
  `marriage` varchar(50) DEFAULT NULL COMMENT '婚姻',
  `occupation` varchar(50) DEFAULT NULL COMMENT '职业',
  `teacher` varchar(50) DEFAULT NULL COMMENT '教师',
  `programmer` varchar(50) DEFAULT NULL COMMENT '程序员',
  `doctor` varchar(50) DEFAULT NULL COMMENT '医生',
  `Truck_driver` varchar(50) DEFAULT NULL COMMENT '货车司机',
  `Online_car_driver` varchar(50) DEFAULT NULL COMMENT '网约车司机',
  `Residential_area` varchar(50) DEFAULT NULL COMMENT '居住地区',
  `Residential_city_level` varchar(50) DEFAULT NULL COMMENT '居住城市级别',
  `house_moving` varchar(50) DEFAULT NULL COMMENT '搬家',
  `Domestic_tourism` varchar(50) DEFAULT NULL COMMENT '国内旅游',
  `Outbound_tour` varchar(50) DEFAULT NULL COMMENT '出境游',
  `have_car` varchar(50) DEFAULT NULL COMMENT '有车',
  `HCar_brand` varchar(50) DEFAULT NULL COMMENT '有车品牌',
  `Preparation_for_pregnancy` varchar(50) DEFAULT NULL COMMENT '备孕',
  `pregnancy` varchar(50) DEFAULT NULL COMMENT '孕期',
  `Parents_of_children_aged_0_1` varchar(50) DEFAULT NULL COMMENT '0-1岁小孩父母',
  `Parents_of_children_aged_0_3` varchar(50) DEFAULT NULL COMMENT '0-3岁小孩父母',
  `Parents_of_children_aged_3_6` varchar(50) DEFAULT NULL COMMENT '3-6岁小孩父母',
  `Parents_of_children_aged_0_6` varchar(50) DEFAULT NULL COMMENT '0-6岁小孩父母',
  `Parents_of_children_aged_3_14` varchar(50) DEFAULT NULL COMMENT '3-14岁小孩父母',
  `Parents_of_children_aged_0_14` varchar(50) DEFAULT NULL COMMENT '0-14岁小孩父母',
  `mum` varchar(50) DEFAULT NULL COMMENT '妈妈',
  `Mother_and_baby` varchar(50) DEFAULT NULL COMMENT '母婴',
  `Parents_of_primary_students` varchar(50) DEFAULT NULL COMMENT '小学生家长',
  `Parents_of_middle_students` varchar(50) DEFAULT NULL COMMENT '中学生家长',
  `second_child` varchar(50) DEFAULT NULL COMMENT '二胎',
  `Stock_Exchange` varchar(50) DEFAULT NULL COMMENT '股票交易',
  `Investment_Financing` varchar(50) DEFAULT NULL COMMENT '投资理财',
  `Bookkeeping` varchar(50) DEFAULT NULL COMMENT '记账',
  `bank` varchar(50) DEFAULT NULL COMMENT '银行',
  `Credit_Card` varchar(50) DEFAULT NULL COMMENT '信用卡',
  `lottery` varchar(50) DEFAULT NULL COMMENT '彩票',
  `p2p` varchar(50) DEFAULT NULL COMMENT '网贷p2p',
  `Integrated_e_commerce` varchar(50) DEFAULT NULL COMMENT '综合电商',
  `Overseas_shopping` varchar(50) DEFAULT NULL COMMENT '海淘',
  `express` varchar(50) DEFAULT NULL COMMENT '快递',
  `liquor` varchar(50) DEFAULT NULL COMMENT '酒类',
  `Group_buying` varchar(50) DEFAULT NULL COMMENT '团购',
  `preschool_education` varchar(50) DEFAULT NULL COMMENT '学前教育',
  `Primary_secondary_education` varchar(50) DEFAULT NULL COMMENT '中小学教育',
  `Adult_Education` varchar(50) DEFAULT NULL COMMENT '成人教育',
  `English_training` varchar(50) DEFAULT NULL COMMENT '英语培训',
  `Training_abroad` varchar(50) DEFAULT NULL COMMENT '出国培训',
  `Examination_certificate` varchar(50) DEFAULT NULL COMMENT '考试证书',
  `Language_learning` varchar(50) DEFAULT NULL COMMENT '语言学习',
  `Online_education` varchar(50) DEFAULT NULL COMMENT '在线教育',
  `Business_travel` varchar(50) DEFAULT NULL COMMENT '商务旅行',
  `Leisure_travel` varchar(50) DEFAULT NULL COMMENT '休闲旅行',
  `Hotel_accommodation` varchar(50) DEFAULT NULL COMMENT '酒店·住宿',
  `outdoors` varchar(50) DEFAULT NULL COMMENT '户外',
  `Travel_around` varchar(50) DEFAULT NULL COMMENT '周边游',
  `Car_rental` varchar(50) DEFAULT NULL COMMENT '租车',
  `buy_car` varchar(50) DEFAULT NULL COMMENT '买车',
  `Car_brand` varchar(50) DEFAULT NULL COMMENT '买车品牌',
  `Driving_test` varchar(50) DEFAULT NULL COMMENT '驾考',
  `Love_car` varchar(50) DEFAULT NULL COMMENT '爱车',
  `Driving_behalf_others` varchar(50) DEFAULT NULL COMMENT '代驾',
  `Violation_inquiry` varchar(50) DEFAULT NULL COMMENT '违章查询',
  `Automobile_maintenance` varchar(50) DEFAULT NULL COMMENT '汽车养护',
  `used_car` varchar(50) DEFAULT NULL COMMENT '二手车',
  `News_information` varchar(50) DEFAULT NULL COMMENT '新闻·资讯',
  `magazine` varchar(50) DEFAULT NULL COMMENT '杂志',
  `novel` varchar(50) DEFAULT NULL COMMENT '小说',
  `Listen_books` varchar(50) DEFAULT NULL COMMENT '听书',
  `cartoon` varchar(50) DEFAULT NULL COMMENT '漫画',
  `Funny` varchar(50) DEFAULT NULL COMMENT '搞笑',
  `Pictorial` varchar(50) DEFAULT NULL COMMENT '画报',
  `Sports_information` varchar(50) DEFAULT NULL COMMENT '体育资讯',
  `karaoke` varchar(50) DEFAULT NULL COMMENT 'K歌',
  `music_player` varchar(50) DEFAULT NULL COMMENT '音乐播放器',
  `radio_station` varchar(50) DEFAULT NULL COMMENT '电台',
  `Ringtone` varchar(50) DEFAULT NULL COMMENT '铃声',
  `Video_player` varchar(50) DEFAULT NULL COMMENT '视频播放器',
  `Live_TV` varchar(50) DEFAULT NULL COMMENT '电视直播',
  `Online_video` varchar(50) DEFAULT NULL COMMENT '在线视频',
  `Online_music` varchar(50) DEFAULT NULL COMMENT '在线音乐',
  `chat` varchar(50) DEFAULT NULL COMMENT '聊天',
  `Single_friends` varchar(50) DEFAULT NULL COMMENT '单身交友',
  `marriage_love` varchar(50) DEFAULT NULL COMMENT '婚恋',
  `micro_blog` varchar(50) DEFAULT NULL COMMENT '微博',
  `community` varchar(50) DEFAULT NULL COMMENT '社区',
  `Restaurant_recommendation` varchar(50) DEFAULT NULL COMMENT '餐厅推荐',
  `take_out_food` varchar(50) DEFAULT NULL COMMENT '外卖',
  `Fresh_delivery` varchar(50) DEFAULT NULL COMMENT '生鲜配送',
  `Looking_for_Aunt` varchar(50) DEFAULT NULL COMMENT '找阿姨',
  `laundry` varchar(50) DEFAULT NULL COMMENT '洗衣',
  `cook` varchar(50) DEFAULT NULL COMMENT '做饭',
  `massage` varchar(50) DEFAULT NULL COMMENT '按摩',
  `film` varchar(50) DEFAULT NULL COMMENT '电影',
  `Pets` varchar(50) DEFAULT NULL COMMENT '宠物',
  `renting` varchar(50) DEFAULT NULL COMMENT '租房',
  `buy_house` varchar(50) DEFAULT NULL COMMENT '买房',
  `home_decoration` varchar(50) DEFAULT NULL COMMENT '家装',
  `wedding` varchar(50) DEFAULT NULL COMMENT '婚礼',
  `Exercise_and_fitness` varchar(50) DEFAULT NULL COMMENT '运动健身',
  `Lose_weight` varchar(50) DEFAULT NULL COMMENT '减肥',
  `Health_medical_treatment` varchar(50) DEFAULT NULL COMMENT '健康医疗',
  `Health_management` varchar(50) DEFAULT NULL COMMENT '健康管理',
  `Shooting_beautification` varchar(50) DEFAULT NULL COMMENT '拍摄美化',
  `Picture_sharing` varchar(50) DEFAULT NULL COMMENT '图片分享',
  `Camera` varchar(50) DEFAULT NULL COMMENT '摄像',
  `Beauty_cosmetics` varchar(50) DEFAULT NULL COMMENT '美容美妆',
  `Parenting_community` varchar(50) DEFAULT NULL COMMENT '育儿社区',
  `Menstrual_management` varchar(50) DEFAULT NULL COMMENT '经期管理',
  `Travel_by_taxi` varchar(50) DEFAULT NULL COMMENT '打车出行',
  `Food_recipes` varchar(50) DEFAULT NULL COMMENT '美食菜谱',
  `Map_navigation` varchar(50) DEFAULT NULL COMMENT '地图导航',
  `Weather_inquiry` varchar(50) DEFAULT NULL COMMENT '天气查询',
  `Storage_cloud_disk` varchar(50) DEFAULT NULL COMMENT '存储云盘',
  `Office_software` varchar(50) DEFAULT NULL COMMENT '办公软件',
  `note` varchar(50) DEFAULT NULL COMMENT '笔记',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `Leisure_intelligence` varchar(50) DEFAULT NULL COMMENT '休闲益智',
  `Children_intelligence` varchar(50) DEFAULT NULL COMMENT '儿童益智',
  `Racing` varchar(50) DEFAULT NULL COMMENT '竞速',
  `Sports` varchar(50) DEFAULT NULL COMMENT '体育',
  `Action_shooting` varchar(50) DEFAULT NULL COMMENT '动作射击',
  `Chess` varchar(50) DEFAULT NULL COMMENT '棋牌',
  `Cosplay` varchar(50) DEFAULT NULL COMMENT '角色扮演',
  `business_strategy` varchar(50) DEFAULT NULL COMMENT '经营策略',
  `Network_game` varchar(50) DEFAULT NULL COMMENT '网络游戏',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `statistics` */

DROP TABLE IF EXISTS `statistics`;

CREATE TABLE `statistics` (
  `tenant_id` varchar(255) NOT NULL COMMENT '租户ID',
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '表自增ID',
  `statistics_item` varchar(255) DEFAULT NULL COMMENT '统计项',
  `item_id` int(11) DEFAULT NULL COMMENT '相关业务ID',
  `count` bigint(22) DEFAULT NULL COMMENT '统计数',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `sub_problem` */

DROP TABLE IF EXISTS `sub_problem`;

CREATE TABLE `sub_problem` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `compay_name` varchar(255) DEFAULT NULL,
  `atten_name` varchar(255) DEFAULT NULL,
  `atten_phone` bigint(15) DEFAULT NULL,
  `atten_type` varchar(255) DEFAULT NULL,
  `atten_content` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `swagger_api_pre_import` */

DROP TABLE IF EXISTS `swagger_api_pre_import`;

CREATE TABLE `swagger_api_pre_import` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `uname` varchar(255) DEFAULT NULL,
  `path` varchar(255) DEFAULT NULL,
  `b_path` varchar(255) DEFAULT NULL COMMENT '后端path',
  `ds_id` int(11) DEFAULT NULL,
  `table_name` varchar(255) DEFAULT NULL,
  `method` varchar(20) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `paging` tinyint(1) NOT NULL DEFAULT '0',
  `query_string` text,
  `type` varchar(20) DEFAULT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1: online, 2: offline',
  `service_id` varchar(80) DEFAULT NULL,
  `service_path` varchar(255) DEFAULT NULL,
  `ext` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `path` (`path`,`method`,`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `swagger_source` */

DROP TABLE IF EXISTS `swagger_source`;

CREATE TABLE `swagger_source` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `url` varchar(250) DEFAULT NULL,
  `definitions` text,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UK` (`url`,`tenant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `swagger_source_online_pre` */

DROP TABLE IF EXISTS `swagger_source_online_pre`;

CREATE TABLE `swagger_source_online_pre` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `url` varchar(250) DEFAULT NULL,
  `definitions` text,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UK` (`url`,`tenant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `third_part_api` */

DROP TABLE IF EXISTS `third_part_api`;

CREATE TABLE `third_part_api` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `path` varchar(255) NOT NULL DEFAULT '',
  `auth_id` int(11) DEFAULT NULL,
  `app_id` int(11) DEFAULT NULL,
  `b_path` varchar(255) NOT NULL DEFAULT '',
  `api_url` varchar(250) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `third_path_application` */

DROP TABLE IF EXISTS `third_path_application`;

CREATE TABLE `third_path_application` (
  `tenant_id` varchar(255) DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `partner_id` int(11) DEFAULT NULL COMMENT '合作伙id',
  `contract_id` int(11) DEFAULT NULL COMMENT '合同id',
  `desciption` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `tree_channel` */

DROP TABLE IF EXISTS `tree_channel`;

CREATE TABLE `tree_channel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `channel_name` varchar(255) NOT NULL COMMENT '渠道名称',
  `parent_id` int(11) DEFAULT NULL COMMENT '父Channel',
  `description` varchar(255) DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

/*Table structure for table `tree_channel_copy1` */

DROP TABLE IF EXISTS `tree_channel_copy1`;

CREATE TABLE `tree_channel_copy1` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `channel_name` varchar(255) NOT NULL COMMENT '渠道名称',
  `parent_id` int(11) DEFAULT NULL COMMENT '父Channel',
  `description` varchar(255) DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Table structure for table `tree_groups` */

DROP TABLE IF EXISTS `tree_groups`;

CREATE TABLE `tree_groups` (
  `tenant_id` varchar(255) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group_name` varchar(255) NOT NULL COMMENT '分组名称',
  `parent_id` int(11) DEFAULT NULL COMMENT '父Group',
  `created_by` varchar(255) DEFAULT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  `updated_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;