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

cosmo_datasource.sql: |-
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


cosmo_gateway.sql: |-
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


cosmo_proxy.sql: |-
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

sdh_auth.sql: |-
CREATE DATABASE /*!32312 IF NOT EXISTS*/`sdh_auth` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `sdh_auth`;

/*Table structure for table `cluster` */

DROP TABLE IF EXISTS `cluster`;

CREATE TABLE `cluster` (
  `id` bigint(16) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL COMMENT '集群名称',
  `short_name` varchar(50) DEFAULT NULL COMMENT '集群短名称Id',
  `created_by` varchar(50) DEFAULT NULL,
  `created_time` bigint(13) DEFAULT NULL,
  `updated_by` varchar(50) DEFAULT NULL,
  `updated_time` bigint(13) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `tenant_id` varchar(255) DEFAULT NULL COMMENT '租户id',
  `path` varchar(255) DEFAULT NULL COMMENT '集群路径',
  `hue_url` varchar(255) DEFAULT NULL COMMENT 'hueUrl',
  `host_name` varchar(255) DEFAULT NULL COMMENT '主机名称',
  `ip` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL COMMENT '类型',
  `username` varchar(255) DEFAULT NULL COMMENT '用户名',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `dispatch` int(2) DEFAULT NULL COMMENT '1kangroo 2车联',
  `set_url` varchar(255) DEFAULT NULL COMMENT '配置URL',
  `deleted` tinyint(1) DEFAULT '0' COMMENT '删除标识',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_short_name` (`short_name`)
) ENGINE=InnoDB AUTO_INCREMENT=1415179545997602818 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `cluster_setting_file_relation` */

DROP TABLE IF EXISTS `cluster_setting_file_relation`;

CREATE TABLE `cluster_setting_file_relation` (
  `id` bigint(16) unsigned NOT NULL AUTO_INCREMENT,
  `cluster_id` bigint(16) DEFAULT NULL COMMENT '集群ID',
  `file_id` bigint(16) DEFAULT NULL COMMENT '文件ID',
  `created_by` varchar(20) DEFAULT NULL COMMENT '创建人',
  `created_time` bigint(13) DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(20) DEFAULT NULL COMMENT '更新人',
  `updated_time` bigint(13) DEFAULT NULL COMMENT '更新时间',
  `file_name` varchar(255) DEFAULT NULL COMMENT '文件名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `group` */

DROP TABLE IF EXISTS `group`;

CREATE TABLE `group` (
  `id` bigint(16) NOT NULL AUTO_INCREMENT,
  `group_name` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `description` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `created_by` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `created_time` bigint(13) DEFAULT NULL,
  `updated_by` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `updated_time` bigint(13) DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT '0',
  `tenant_id` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `code` varchar(20) COLLATE utf8_bin DEFAULT NULL COMMENT '是否是管理员',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_group_name` (`group_name`,`tenant_id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Table structure for table `group_data_relation` */

DROP TABLE IF EXISTS `group_data_relation`;

CREATE TABLE `group_data_relation` (
  `id` bigint(16) NOT NULL AUTO_INCREMENT,
  `group_id` bigint(16) DEFAULT NULL COMMENT '用户组id',
  `data_id` bigint(16) DEFAULT NULL COMMENT '数据源id',
  `created_time` bigint(13) DEFAULT NULL,
  `created_by` varchar(50) DEFAULT NULL,
  `updated_time` bigint(13) DEFAULT NULL,
  `updated_by` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1394554722697367555 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `node` */

DROP TABLE IF EXISTS `node`;

CREATE TABLE `node` (
  `id` bigint(16) NOT NULL COMMENT '节点id',
  `name` varchar(50) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `host` varchar(255) DEFAULT NULL,
  `source` bigint(16) DEFAULT NULL COMMENT '集群ID',
  `created_by` varchar(255) DEFAULT NULL,
  `created_time` bigint(13) DEFAULT NULL,
  `updated_by` varchar(50) DEFAULT NULL,
  `updated_time` bigint(13) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Table structure for table `resource` */

DROP TABLE IF EXISTS `resource`;

CREATE TABLE `resource` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) DEFAULT NULL,
  `code` varchar(100) DEFAULT NULL,
  `type` int(11) DEFAULT NULL COMMENT '1menu2page3button4data5url',
  `name` varchar(50) DEFAULT NULL,
  `url` varchar(100) DEFAULT NULL,
  `method` varchar(20) DEFAULT NULL,
  `description` varchar(50) DEFAULT NULL,
  `created_time` bigint(20) DEFAULT NULL,
  `created_by` varchar(50) DEFAULT NULL,
  `updated_time` bigint(20) DEFAULT NULL,
  `updated_by` varchar(50) DEFAULT NULL,
  `order` int(11) DEFAULT NULL,
  `can_set` tinyint(1) DEFAULT '1' COMMENT '是否可以设置',
  `is_pub` tinyint(1) DEFAULT '0' COMMENT '是否是公共的接口',
  `request_url` varchar(255) DEFAULT NULL COMMENT '请求接口',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ux_resource_code` (`code`),
  UNIQUE KEY `ux_resource_path` (`url`),
  KEY `index_parent_id` (`parent_id`) USING BTREE,
  KEY `index_is_pub` (`is_pub`) USING BTREE,
  KEY `index_type` (`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1339111450366447619 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `role` */

DROP TABLE IF EXISTS `role`;

CREATE TABLE `role` (
  `id` bigint(16) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(50) DEFAULT NULL COMMENT '角色名称',
  `description` varchar(255) DEFAULT NULL COMMENT '角色说明',
  `created_by` varchar(50) DEFAULT NULL COMMENT '创建人',
  `created_time` bigint(13) DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(50) DEFAULT NULL COMMENT '更新人',
  `updated_time` bigint(13) DEFAULT NULL COMMENT '更新时间',
  `tenant_id` varchar(50) DEFAULT NULL COMMENT '所属租户id',
  `code` varchar(50) DEFAULT NULL COMMENT '平台管理员SUPER_GROUP,租户管理员TENANT_ADMIN',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4;

/*Table structure for table `role_resource_relation` */

DROP TABLE IF EXISTS `role_resource_relation`;

CREATE TABLE `role_resource_relation` (
  `id` bigint(16) NOT NULL AUTO_INCREMENT,
  `role_id` bigint(16) DEFAULT NULL,
  `resource_id` bigint(16) DEFAULT NULL,
  `created_time` bigint(13) DEFAULT NULL,
  `created_by` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `updated_time` bigint(13) DEFAULT NULL,
  `updated_by` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_role_resource` (`role_id`,`resource_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1416998050304749570 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Table structure for table `task` */

DROP TABLE IF EXISTS `task`;

CREATE TABLE `task` (
  `id` varchar(255) NOT NULL COMMENT 'md5加密id',
  `sql_content` text COMMENT 'sql内容',
  `status` enum('NEW','DONE','ERROR','running_to_file','running_to_ftp') NOT NULL COMMENT '状态类型',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '修改时间',
  `file_path` varchar(255) DEFAULT NULL COMMENT '文件路径',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='api表';

/*Table structure for table `tenant` */

DROP TABLE IF EXISTS `tenant`;

CREATE TABLE `tenant` (
  `id` varchar(255) COLLATE utf8_bin NOT NULL,
  `name` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '租户名称',
  `description` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '简要描述',
  `created_by` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '创建人',
  `created_time` bigint(13) DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '更新人',
  `updated_time` bigint(13) DEFAULT NULL COMMENT '最后更新时间',
  `short_name` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '来源短名称',
  `status` varchar(10) COLLATE utf8_bin DEFAULT NULL COMMENT '来源状态',
  `default_access_rule` varchar(10) COLLATE utf8_bin DEFAULT NULL,
  `company_id` varchar(20) COLLATE utf8_bin DEFAULT NULL COMMENT '公司id',
  `contact_a_name` varchar(20) COLLATE utf8_bin DEFAULT NULL,
  `contact_a_tel` varchar(20) COLLATE utf8_bin DEFAULT NULL,
  `contact_a_position` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `contact_b_name` varchar(20) COLLATE utf8_bin DEFAULT NULL,
  `contact_b_tel` varchar(20) COLLATE utf8_bin DEFAULT NULL,
  `contact_b_position` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `company_name` varchar(50) COLLATE utf8_bin DEFAULT NULL COMMENT '公司名称',
  `company_short_name` varchar(20) COLLATE utf8_bin DEFAULT NULL COMMENT '公司短名称',
  `deleted` tinyint(1) DEFAULT NULL COMMENT '是否删除',
  `group` int(11) DEFAULT NULL,
  `maintainer_openid` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `nature` int(11) DEFAULT NULL,
  `priority` int(255) DEFAULT NULL,
  `used` tinyint(1) DEFAULT NULL COMMENT '是否常用',
  `can_set` tinyint(1) DEFAULT '0' COMMENT '是否可以进行设置和编辑',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` bigint(16) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `saic_id` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '统一认证平台ID',
  `dept` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '部门',
  `telephone` varchar(20) COLLATE utf8_bin DEFAULT NULL COMMENT '手机号',
  `username` varchar(50) COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `email` varchar(50) COLLATE utf8_bin DEFAULT NULL COMMENT '邮箱',
  `created_by` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '创建人',
  `created_time` bigint(13) DEFAULT NULL COMMENT '创建时间',
  `updated_by` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '更新人',
  `updated_time` bigint(13) DEFAULT NULL COMMENT '更新时间',
  `status` int(1) DEFAULT '-1' COMMENT '-3从mq插入的-2停用-1未启用0未激活1已启用',
  `account_start_expired` bigint(13) DEFAULT NULL COMMENT '账号有效期的开始时间',
  `account_expired` bigint(13) DEFAULT NULL COMMENT '账号过期时间',
  `domain_account` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '域账号',
  `deleted` tinyint(1) DEFAULT '0' COMMENT '是否删除',
  `password` varchar(50) COLLATE utf8_bin DEFAULT NULL COMMENT '密码',
  `auth_type` varchar(100) DEFAULT '0',
  `uuc_user_id` BIGINT(16) DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_telephone` (`telephone`)
) ENGINE=InnoDB AUTO_INCREMENT=177 DEFAULT CHARSET=utf8 COLLATE=utf8_bin COMMENT='用户表';

/*Table structure for table `user_group_relation` */

DROP TABLE IF EXISTS `user_group_relation`;

CREATE TABLE `user_group_relation` (
  `id` bigint(16) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(16) DEFAULT NULL,
  `group_id` bigint(16) DEFAULT NULL,
  `created_by` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `created_time` bigint(13) DEFAULT NULL,
  `updated_by` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `updated_time` bigint(13) DEFAULT NULL,
  `tenant_id` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1394554906470797320 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Table structure for table `user_role_relation` */

DROP TABLE IF EXISTS `user_role_relation`;

CREATE TABLE `user_role_relation` (
  `id` bigint(16) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(16) DEFAULT NULL,
  `role_id` bigint(16) DEFAULT NULL,
  `created_time` bigint(13) DEFAULT NULL COMMENT '创建时间',
  `created_by` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '创建人',
  `updated_time` bigint(13) DEFAULT NULL,
  `updated_by` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `tenant_id` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `index_user_id` (`user_id`,`role_id`,`tenant_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2161 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Table structure for table `user_tenant_relation` */

DROP TABLE IF EXISTS `user_tenant_relation`;

CREATE TABLE `user_tenant_relation` (
  `id` bigint(16) NOT NULL AUTO_INCREMENT,
  `tenant_id` varchar(255) COLLATE utf8_bin NOT NULL,
  `user_id` bigint(16) NOT NULL,
  `created_by` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '创建人',
  `created_time` bigint(13) DEFAULT NULL COMMENT '时间戳，创建时间',
  `updated_by` varchar(255) COLLATE utf8_bin DEFAULT NULL COMMENT '更新人',
  `updated_time` bigint(13) DEFAULT NULL COMMENT '更新时间',
  `role` varchar(50) COLLATE utf8_bin DEFAULT NULL COMMENT '用户对应租户的级别',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=255 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


sdh_auth_data.sql: |-
use sdh_auth;
insert into `role` (`id`, `name`, `description`, `created_by`, `created_time`, `updated_by`, `updated_time`, `tenant_id`, `code`) values('-1','新数据市场管理员','与租户关联的市场管理员','system',NULL,'system',NULL,NULL,'NEW_MARKET_ADMIN');
insert into `role` (`id`, `name`, `description`, `created_by`, `created_time`, `updated_by`, `updated_time`, `tenant_id`, `code`) values('0','市场管理员','市场管理员','system',NULL,'system',NULL,NULL,'MARKET_ADMIN');
UPDATE `role` SET id = 0 WHERE id = 1;
insert into `role` (`id`, `name`, `description`, `created_by`, `created_time`, `updated_by`, `updated_time`, `tenant_id`, `code`) values('1','平台管理员','具有分配租户管理员，配置数据，功能等权限','default_user_name','1594293043994','default_user_name','1594293043994',NULL,'SUPER_ADMIN');
insert into `role` (`id`, `name`, `description`, `created_by`, `created_time`, `updated_by`, `updated_time`, `tenant_id`, `code`) values('2','租户管理员','具有该租户下的所有功能权限','default_user_name','1594293065318','???','1596776532885',NULL,'TENANT_ADMIN');
insert into `role` (`id`, `name`, `description`, `created_by`, `created_time`, `updated_by`, `updated_time`, `tenant_id`, `code`) values('3','应用管理员','具有应用的所有权限','default_user_name','1594293080726','default_user_name','1594293080726',NULL,NULL);
insert into `role` (`id`, `name`, `description`, `created_by`, `created_time`, `updated_by`, `updated_time`, `tenant_id`, `code`) values('4','数据开发','具有平台各模块的部分操作权限','default_user_name','1594293103863','default_user_name','1594293103863',NULL,NULL);
insert into `role` (`id`, `name`, `description`, `created_by`, `created_time`, `updated_by`, `updated_time`, `tenant_id`, `code`) values('5','业务开发','具有平台模块查看以及部分应用的权限','default_user_name','1594293128376','default_user_name','1594293128376',NULL,NULL);
insert into `role` (`id`, `name`, `description`, `created_by`, `created_time`, `updated_by`, `updated_time`, `tenant_id`, `code`) values('6','运维','具有各模块查看以及任务运维权限','default_user_name','1594293153284','default_user_name','1594293153284',NULL,NULL);
insert into `role` (`id`, `name`, `description`, `created_by`, `created_time`, `updated_by`, `updated_time`, `tenant_id`, `code`) values('7','访客','具有部分模块的查看权限，不具备操作权限','default_user_name','1594293175925','default_user_name','1594293175925',NULL,NULL);
insert into `role` (`id`, `name`, `description`, `created_by`, `created_time`, `updated_by`, `updated_time`, `tenant_id`, `code`) values('19','数据服务开发','具有数据源管理、数据服务、个人中心的权限','default_user_name','1596770055882','default_user_name','1596770159076','027cf661-2c58-40ac-9073-1ebd18eec8c4',NULL);
insert into `role` (`id`, `name`, `description`, `created_by`, `created_time`, `updated_by`, `updated_time`, `tenant_id`, `code`) values('20','实时数据开发','具有数据源管理、实时开发模块的权限','default_user_name','1596776267662','default_user_name','1596776267662','027cf661-2c58-40ac-9073-1ebd18eec8c4',NULL);
insert into `tenant` (`id`, `name`, `description`, `created_by`, `created_time`, `updated_by`, `updated_time`, `short_name`, `status`, `default_access_rule`, `company_id`, `contact_a_name`, `contact_a_tel`, `contact_a_position`, `contact_b_name`, `contact_b_tel`, `contact_b_position`, `company_name`, `company_short_name`, `deleted`, `group`, `maintainer_openid`, `nature`, `priority`, `used`, `can_set`) values('027cf661-2c58-40ac-9073-1ebd18eec8c4','卡奥斯大数据服务','卡奥斯大数据服务租户\n负责人： 陈辉','admin','1593656755686',NULL,'1593656755686','saicdt','1','1',NULL,'王的','13333333333',NULL,NULL,NULL,NULL,NULL,NULL,'0',NULL,NULL,NULL,NULL,'1','0');
insert into `user` ( `saic_id`, `dept`, `telephone`, `username`, `email`, `created_by`, `created_time`, `updated_by`, `updated_time`, `status`, `account_start_expired`, `account_expired`, `domain_account`, `deleted`, `password`) values(NULL,'9月测试','15295756791','data_service_test','133@133.com','test','1632375362472','default_user_name','1671100118604','1',NULL,NULL,NULL,'0',NULL);
insert into `user_tenant_relation` (`tenant_id`, `user_id`, `created_by`, `created_time`, `updated_by`, `updated_time`, `role`) values('027cf661-2c58-40ac-9073-1ebd18eec8c4','1','test','1629337652786','test','1629337652781',NULL);
insert into `user_role_relation` ( `user_id`, `role_id`, `created_time`, `created_by`, `updated_time`, `updated_by`, `tenant_id`) values('1','2','1629337652791','test','1629337652789','test','027cf661-2c58-40ac-9073-1ebd18eec8c4');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1','0',NULL,'1','首页','/home','get','首页','1595300932367','system',NULL,NULL,'2','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('2','1','/dataSourceMgt','1','数据源管理',NULL,'get','数据源管理','1595300932307','system',NULL,NULL,'21','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('3','2',NULL,'2','通用数据源','/dataSourceMgt','get','通用数据源','1595300932303','system',NULL,NULL,'211','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('4','3','dataSourceMgtCreate','3','新建数据源',NULL,'post','新建数据源','1595300932235','system',NULL,NULL,'2111','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('5','3','dataSourceMgtEdit','3','编辑数据源',NULL,'put','编辑数据源','1595300932301','system',NULL,NULL,'2112','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('6','3','dataSourceMgtDelete','3','删除数据源',NULL,'delete','删除数据源','1595300932302','system',NULL,NULL,'2113','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('7','1','/dataService','1','数据服务',NULL,'get','数据服务','1595300932319','system',NULL,NULL,'22','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('8','7',NULL,'2','API概况','/dataService','get','API概况','1595300932310','system',NULL,NULL,'221','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('9','7',NULL,'2','API管理','/dataService/manage','get','API管理','1595300932312','system',NULL,NULL,'222','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('10','7',NULL,'2','API市场','/dataService/store','get','API市场','1595300932313','system',NULL,NULL,'223','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('11','7',NULL,'2','应用中心','/dataService/center','get','应用中心','1595300932318','system',NULL,NULL,'224','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('12','11',NULL,'2','应用管理','/dataService/center/manage','get','应用管理','1595300932315','system',NULL,NULL,'2241','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('13','11',NULL,'2','应用API','/dataService/center/api','get','应用API','1595300932316','system',NULL,NULL,'2242','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('14','1','/RealTime','1','实时数据处理',NULL,'get','实时数据处理','1595300932324','system',NULL,NULL,'23','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('15','14',NULL,'2','实时计算','/RealTime','get','实时计算','1595300932322','system',NULL,NULL,'231','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('16','1','/dataAssets','2','数据资产',NULL,'get','数据资产','1595300932327','system',NULL,NULL,'24','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('17','1','systemManage','1','系统管理',NULL,'get','系统管理','1595300932362','system',NULL,NULL,'25','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('18','17',NULL,'2','租户管理','/systemManage/tenantManage','get','租户管理','1595300932329','system',NULL,NULL,'251','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('19','17',NULL,'2','用户管理','/systemManage/userManage','get','用户管理','1595300932340','system',NULL,NULL,'252','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('20','19','userManageCreate','3','创建用户',NULL,'post','创建用户','1595300932334','system',NULL,NULL,'2521','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('21','19','userManageEdit','3','编辑用户资料',NULL,'put','编辑用户资料','1595300932337','system',NULL,NULL,'2522','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('22','19','userManageDelete','3','删除用户',NULL,'delete','删除用户','1595300932338','system',NULL,NULL,'2523','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('23','19','userManageEnable','3','启用-停用',NULL,'put','启用-停用','1595300932339','system',NULL,NULL,'2524','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('24','17',NULL,'2','角色管理','/systemManage/roleManage','get','角色管理','1595300932349','system',NULL,NULL,'253','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('25','24','roleManageCreate','3','创建角色',NULL,'post','创建角色','1595300932341','system',NULL,NULL,'2531','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('26','24','roleManageEdit','3','编辑角色',NULL,'put','编辑角色','1595300932344','system',NULL,NULL,'2532','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('27','24','roleManageDelete','3','删除角色',NULL,'delete','删除角色','1595300932345','system',NULL,NULL,'2533','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('28','24','roleManageFunAuth','3','设置角色权限',NULL,'post','设置角色权限','1595300932347','system',NULL,NULL,'2534','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('29','24','roleManageAddUser','3','添加用户',NULL,'post','添加用户','1595300932348','system',NULL,NULL,'2535','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('30','17',NULL,'2','用户组管理','/systemManage/groupManage','get','用户组管理','1595300932359','system',NULL,NULL,'254','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('31','30','groupManageCreate','3','创建用户组',NULL,'post','创建用户组','1595300932351','system',NULL,NULL,'2541','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('32','30','groupManageEdit','3','编辑用户组',NULL,'put','编辑用户组','1595300932353','system',NULL,NULL,'2542','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('33','30','groupManageDelete','3','删除用户组',NULL,'delete','删除用户组','1595300932355','system',NULL,NULL,'2543','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('34','30','groupManageDataAuth','3','数据源权限管理',NULL,'get','数据源权限管理','1595300932356','system',NULL,NULL,'2544','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('35','30','groupManageAddUser','3','用户组添加用户',NULL,'get','用户组添加用户','1595300932358','system',NULL,NULL,'2545','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('36','1','/outline','1','离线开发',NULL,'get','离线开发','1595300932364','system',NULL,NULL,'26','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('37','1','/personal','1','个人中心','/personal','get','个人中心','1595300932366','system',NULL,NULL,'27','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('38','14',NULL,'2','实时数据传输','/RealTime/trans','get','实时数据传输','1597376558267','system',NULL,NULL,'2811','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('39','14',NULL,'2','实时数据计算','/RealTime/compute','get','实时数据计算','1597376558312','system',NULL,NULL,'2812','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('42','7',NULL,'2','API数据源','/dataService/datasource','get','API数据源','1597376558315','system',NULL,NULL,'2821','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('43','7',NULL,'2','API来源管理','/dataService/channel','get','API来源管理','1597376558317','system',NULL,NULL,'2822','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('44','37','/personal/apply/api','1','我的申请','','get','我的申请','1597376558320','system',NULL,NULL,'2831','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('45','44',NULL,'2','API权限','/personal/apply/api','get','API权限','1597376558318','system',NULL,NULL,'28311','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('46','37',NULL,'1','审批授权','/personal/approve','get','审批授权','1597376558324','system',NULL,NULL,'2832','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('47','46',NULL,'2','API审批授权','/personal/approve/api','get','API审批授权','1597376558322','system',NULL,NULL,'28321','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('48','17',NULL,'2','集群管理','/systemManage/colonyManage','get','集群管理','1595300932359','system',NULL,NULL,'255','0','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('49','9','api_apiMgt_creact','3','新增api',NULL,'post','新增api','1598497769023','system',NULL,NULL,'2221','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('50','9','api_apiMgt_edit','3','编辑api',NULL,'put','编辑api','1598497769066','system',NULL,NULL,'2222','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('51','9','api_apiMgt_del','3','删除api',NULL,'delete','删除api','1598497769068','system',NULL,NULL,'2223','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('52','9','api_apiMgt_on','3','发布api',NULL,'patch','发布api','1598497769069','system',NULL,NULL,'2224','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('53','9','api_apiMgt_auth','3','授权api',NULL,'post','授权api','1598497769071','system',NULL,NULL,'2225','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('54','9','api_apiMgt_test','3','测试api',NULL,'post','测试api','1598497769073','system',NULL,NULL,'2226','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('55','9','api_apiMgt_off','3','下线api',NULL,'patch','下线api','1598497769074','system',NULL,NULL,'2227','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('56','9','api_apiMgt_submit','3','提交api',NULL,'patch','提交api','1598497769076','system',NULL,NULL,'2228','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('57','9','api_apiMgt_withdraw','3','撤回api',NULL,'patch','撤回api','1598497769077','system',NULL,NULL,'2229','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('58','9','api_apiMgt_group','3','分组新增修改',NULL,'post','分组新增修改','1598497769079','system',NULL,NULL,'2230','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('59','10','api_apiStore_test','3','api测试',NULL,'post','api测试','1598497769080','system',NULL,NULL,'2231','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('60','10','api_apiStore_apply','3','api申请',NULL,'post','api申请','1598497769082','system',NULL,NULL,'2232','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('61','12','api_appMgt_create','3','新建app',NULL,'post','新建app','1598497769083','system',NULL,NULL,'22411','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('62','12','api_appMgt_edit','3','编辑app',NULL,'put','编辑app','1598497769085','system',NULL,NULL,'22412','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('63','12','api_appMgt_del','3','删除app',NULL,'delete','删除app','1598497769086','system',NULL,NULL,'22413','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('64','13','api_appApi_test','3','测试api',NULL,'post','测试api','1598497769087','system',NULL,NULL,'22421','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('65','13','api_appApi_cancelAuth','3','取消授权api',NULL,'get','取消授权api','1598497769090','system',NULL,NULL,'22422','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('66','42','api_apiDs_selectDs','3','选择数据源',NULL,'post','选择数据源','1598497769091','system',NULL,NULL,'28211','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('67','43','api_apiChannel_create','3','api来源新增',NULL,'post','api来源新增','1598497769092','system',NULL,NULL,'28221','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('68','43','api_apiChannel_edit','3','api来源编辑',NULL,'post','api来源编辑','1598497769093','system',NULL,NULL,'28222','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('69','43','api_apiChannel_del','3','api来源删除',NULL,'delete','api来源删除','1598497769094','system',NULL,NULL,'28223','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('70','36',NULL,'0','任务管理','/outline/taskmeta','get','任务管理','1598927560396','system',NULL,NULL,'261','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('71','36',NULL,'0','任务运维','/outline/taskrun','get','任务运维','1598927560439','system',NULL,NULL,'262','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('72','16',NULL,'2','数据字典','/dataAssets/dictionary','get','数据字典',NULL,NULL,NULL,NULL,'241','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('73','39',NULL,'3','元数据管理','/RealTime/compute/metadataMgt','get','元数据管理',NULL,NULL,NULL,NULL,'28122','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('74','39',NULL,'3','任务管理','/RealTime/compute/admin','get','任务管理',NULL,NULL,NULL,NULL,'28123','1','0','');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('75','17',NULL,'2','数据市场配置管理','/systemManage/dataMarketManage','get','数据市场配置管理',NULL,NULL,NULL,NULL,'256','0','0',NULL);
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('76','17',NULL,'2','数据市场用户','/systemManage/dataMarketUser','get','数据市场用户',NULL,NULL,NULL,NULL,NULL,'0','0',NULL);
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('77','17',NULL,'2','超级管理员管理','/systemManage/adminUserManage','get','超级管理员管理',NULL,NULL,NULL,NULL,NULL,'0','0',NULL);
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('78','0',NULL,'4',NULL,NULL,'get',NULL,NULL,'system',NULL,NULL,NULL,'0','1','/api-auth/user/getUserInfo');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('84','1','/sdh-cboard','1','数据可视化','/sdh-cboard','get',NULL,NULL,'system',NULL,NULL,NULL,'1','1',NULL);
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419806748674','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016634','system','1608104016635',NULL,NULL,'0','0','/api-manager/manager/app/private');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419848691713','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016636','system','1608104016637',NULL,NULL,'0','0','/api-manager/manager/channel/select');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419852886017','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016637','system','1608104016638',NULL,NULL,'0','0','/api-manager/manager/api');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419857080322','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016638','system','1608104016639',NULL,NULL,'0','0','/api-manager/manager/groups/select');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419861274626','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016640','system','1608104016640',NULL,NULL,'0','0','/api-manager/manager/api/{id}/on');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419873857537','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016644','system','1608104016644',NULL,NULL,'0','0','/api-manager/manager/api/{id}/off');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419886440450','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016645','system','1608104016646',NULL,NULL,'0','0','/api-manager/manager/api/{id}/publish');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419894829057','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016648','system','1608104016649',NULL,NULL,'0','0','/api-manager/data-market/dataSet/api/apply');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419907411970','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016651','system','1608104016651',NULL,NULL,'0','0','/api-manager/manager/api/{id}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419911606273','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016652','system','1608104016652',NULL,NULL,'0','0','/api-manager/manager/api/authorize');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419915800578','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016653','system','1608104016653',NULL,NULL,'0','0','/api-manager/manager/treeGroups');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419919994882','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016655','system','1608104016655',NULL,NULL,'0','0','/api-manager/manager/groups');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419928383489','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016656','system','1608104016657',NULL,NULL,'0','0','/api-manager/manager/groups/{id}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419936772097','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016658','system','1608104016658',NULL,NULL,'0','0','/api-manager/manager/datasources');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419940966402','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016660','system','1608104016661',NULL,NULL,'0','0','/api-datasource/metadata/tables/{dataSoureName}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419953549313','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016663','system','1608104016664',NULL,NULL,'0','0','/api-datasource/metadata/fields/{dataSoureName}/{tableName}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419970326529','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016666','system','1608104016666',NULL,NULL,'0','0','/api-manager/manager/api-test');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419978715137','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016668','system','1608104016668',NULL,NULL,'0','0','/api-manager/manager/swagger/preview');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419987103746','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016669','system','1608104016669',NULL,NULL,'0','0','/api-manager/manager/swagger/import/{service_id}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419987103747','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016670','system','1608104016671',NULL,NULL,'0','0','/api-datasource/metadata/datasources');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419995492354','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016671','system','1608104016671',NULL,NULL,'0','0','/api-manager/manager/api/{apiId}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419995492355','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016672','system','1608104016672',NULL,NULL,'0','0','/api-manager/manager/app-api');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111419999686658','9',NULL,'4',NULL,NULL,'get',NULL,'1608104016673','system','1608104016673',NULL,NULL,'0','0','/api-manager/manager/gateway-admin/statistic/day');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111423384489985','45',NULL,'4',NULL,NULL,'get',NULL,'1608104017480','system','1608104017481',NULL,NULL,'0','0','/api-manager/manager/app/private');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111423392878593','45',NULL,'4',NULL,NULL,'get',NULL,'1608104017482','system','1608104017482',NULL,NULL,'0','0','/api-manager/manager/api-apply/myApproved');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111423397072898','45',NULL,'4',NULL,NULL,'get',NULL,'1608104017482','system','1608104017482',NULL,NULL,'0','0','/api-manager/manager/api-apply/myInApproval');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111423397072899','45',NULL,'4',NULL,NULL,'get',NULL,'1608104017483','system','1608104017483',NULL,NULL,'0','0','/api-manager/manager/api-apply/{apiId}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111423401267202','45',NULL,'4',NULL,NULL,'get',NULL,'1608104017484','system','1608104017484',NULL,NULL,'0','0','/api-manager/manager/api-apply');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111424642781185','24',NULL,'4',NULL,NULL,'get',NULL,'1608104017780','system','1608104017781',NULL,NULL,'0','0','/api-auth/role');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111424651169794','24',NULL,'4',NULL,NULL,'get',NULL,'1608104017782','system','1608104017782',NULL,NULL,'0','0','/api-auth/user/role/{id}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111424655364097','24',NULL,'4',NULL,NULL,'get',NULL,'1608104017783','system','1608104017783',NULL,NULL,'0','0','/api-auth/role/addUser');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111424659558402','24',NULL,'4',NULL,NULL,'get',NULL,'1608104017784','system','1608104017784',NULL,NULL,'0','0','/api-auth/resource');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111424667947010','24',NULL,'4',NULL,NULL,'get',NULL,'1608104017785','system','1608104017785',NULL,NULL,'0','0','/api-auth/role/addResource');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111424667947011','24',NULL,'4',NULL,NULL,'get',NULL,'1608104017786','system','1608104017786',NULL,NULL,'0','0','/api-auth/role/update');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111424672141313','24',NULL,'4',NULL,NULL,'get',NULL,'1608104017786','system','1608104017786',NULL,NULL,'0','0','/api-auth/role/delete');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111424672141314','24',NULL,'4',NULL,NULL,'get',NULL,'1608104017787','system','1608104017787',NULL,NULL,'0','0','/api-auth/role/add');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111426236616705','30',NULL,'4',NULL,NULL,'get',NULL,'1608104018160','system','1608104018161',NULL,NULL,'0','0','/api-auth/group');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111426245005313','30',NULL,'4',NULL,NULL,'get',NULL,'1608104018162','system','1608104018162',NULL,NULL,'0','0','/api-auth/user/group/{id}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111426249199617','30',NULL,'4',NULL,NULL,'get',NULL,'1608104018163','system','1608104018163',NULL,NULL,'0','0','/api-auth/group/addUser');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111426253393922','30',NULL,'4',NULL,NULL,'get',NULL,'1608104018163','system','1608104018163',NULL,NULL,'0','0','/api-datasource/metadata/admin/datasources');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111426253393923','30',NULL,'4',NULL,NULL,'get',NULL,'1608104018164','system','1608104018164',NULL,NULL,'0','0','/api-auth/group/addDataAuth');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111426257588226','30',NULL,'4',NULL,NULL,'get',NULL,'1608104018165','system','1608104018165',NULL,NULL,'0','0','/api-auth/group/update');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111426261782529','30',NULL,'4',NULL,NULL,'get',NULL,'1608104018165','system','1608104018165',NULL,NULL,'0','0','/api-auth/group/delete');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111427717206017','3',NULL,'4',NULL,NULL,'get',NULL,'1608104018513','system','1608104018514',NULL,NULL,'0','0','/api-auth/user/getUserByTenant');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111427725594626','3',NULL,'4',NULL,NULL,'get',NULL,'1608104018515','system','1608104018515',NULL,NULL,'0','0','/api-auth/cluster');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111427729788930','3',NULL,'4',NULL,NULL,'get',NULL,'1608104018515','system','1608104018515',NULL,NULL,'0','0','/api-datasource/dsDataSource');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111427729788931','3',NULL,'4',NULL,NULL,'get',NULL,'1608104018516','system','1608104018516',NULL,NULL,'0','0','/api-datasource/dsDataSource/{id}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111427733983234','3',NULL,'4',NULL,NULL,'get',NULL,'1608104018516','system','1608104018517',NULL,NULL,'0','0','/api-datasource/metadata/test/connection');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111428979691521','19',NULL,'4',NULL,NULL,'get',NULL,'1608104018814','system','1608104018815',NULL,NULL,'0','0','/api-auth/user');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111428988080130','19',NULL,'4',NULL,NULL,'get',NULL,'1608104018815','system','1608104018815',NULL,NULL,'0','0','/api-auth/group');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111428988080131','19',NULL,'4',NULL,NULL,'get',NULL,'1608104018816','system','1608104018816',NULL,NULL,'0','0','/api-auth/role');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111428992274433','19',NULL,'4',NULL,NULL,'get',NULL,'1608104018817','system','1608104018817',NULL,NULL,'0','0','/api-auth/tenant/getQueryTenant');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111428996468737','19',NULL,'4',NULL,NULL,'get',NULL,'1608104018818','system','1608104018818',NULL,NULL,'0','0','/api-auth/user/delete');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111429000663042','19',NULL,'4',NULL,NULL,'get',NULL,'1608104018818','system','1608104018818',NULL,NULL,'0','0','/api-auth/user/update');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111429000663043','19',NULL,'4',NULL,NULL,'get',NULL,'1608104018819','system','1608104018819',NULL,NULL,'0','0','/api-auth/user/add');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111430506418177','38',NULL,'4',NULL,NULL,'get',NULL,'1608104019178','system','1608104019179',NULL,NULL,'0','0','/api-mirror/jobs/findList');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111430514806786','38',NULL,'4',NULL,NULL,'get',NULL,'1608104019179','system','1608104019179',NULL,NULL,'0','0','/api-datasource/metadata/datasources');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111430514806787','38',NULL,'4',NULL,NULL,'get',NULL,'1608104019180','system','1608104019180',NULL,NULL,'0','0','/api-datasource/metadata/tables/{db}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111430519001089','38',NULL,'4',NULL,NULL,'get',NULL,'1608104019180','system','1608104019180',NULL,NULL,'0','0','/api-datasource/metadata/fields/{db}/{table}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111430519001090','38',NULL,'4',NULL,NULL,'get',NULL,'1608104019181','system','1608104019181',NULL,NULL,'0','0','/api-mirror/jobs/insertJobs');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111430523195394','38',NULL,'4',NULL,NULL,'get',NULL,'1608104019181','system','1608104019182',NULL,NULL,'0','0','/api-mirror/jobs/findJobById');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111430523195395','38',NULL,'4',NULL,NULL,'get',NULL,'1608104019182','system','1608104019182',NULL,NULL,'0','0','/api-mirror/jobs/deleteJob');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111430527389698','38',NULL,'4',NULL,NULL,'get',NULL,'1608104019183','system','1608104019183',NULL,NULL,'0','0','/api-mirror/jobs/startJob');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111430531584002','38',NULL,'4',NULL,NULL,'get',NULL,'1608104019184','system','1608104019184',NULL,NULL,'0','0','/api-mirror/jobs/restartJob');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111430535778305','38',NULL,'4',NULL,NULL,'get',NULL,'1608104019184','system','1608104019184',NULL,NULL,'0','0','/api-mirror/jobs/stopJob');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111432364494849','8',NULL,'4',NULL,NULL,'get',NULL,'1608104019621','system','1608104019622',NULL,NULL,'0','0','/api-manager/home/navigate');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111432372883458','8',NULL,'4',NULL,NULL,'get',NULL,'1608104019623','system','1608104019623',NULL,NULL,'0','0','/api-manager/home/app-api');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433316601857','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019848','system','1608104019849',NULL,NULL,'0','0','/api-manager/manager/datasources');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433324990466','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019849','system','1608104019849',NULL,NULL,'0','0','/api-auth/user/getUserByTenant');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433324990467','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019850','system','1608104019850',NULL,NULL,'0','0','/api-auth/cluster');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433329184769','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019850','system','1608104019850',NULL,NULL,'0','0','/api-map/data/map/ListTable');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433329184770','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019851','system','1608104019851',NULL,NULL,'0','0','/api-datasource/metadata/datasources');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433333379073','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019851','system','1608104019851',NULL,NULL,'0','0','/api-map/data/map/tables/{dsId}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433333379074','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019852','system','1608104019852',NULL,NULL,'0','0','/api-map/data/map/fields/{tableId}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433337573377','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019852','system','1608104019852',NULL,NULL,'0','0','/api-map/data/map/field');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433337573378','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019853','system','1608104019853',NULL,NULL,'0','0','/api-map/data/map/table/add');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433341767682','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019853','system','1608104019853',NULL,NULL,'0','0','/api-map/data/map/field/import');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433341767683','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019854','system','1608104019854',NULL,NULL,'0','0','/api-map/data/map/table/import');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433345961986','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019854','system','1608104019854',NULL,NULL,'0','0','/api-map/data/map/queryTableStructure');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433345961987','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019855','system','1608104019855',NULL,NULL,'0','0','/api-map/data/map/queryTableInformation');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433350156289','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019855','system','1608104019855',NULL,NULL,'0','0','/api-map/data/map/table/chart');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111433350156290','72',NULL,'4',NULL,NULL,'get',NULL,'1608104019856','system','1608104019856',NULL,NULL,'0','0','/api-map/data/map/queryTableSimpleData');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111435795435521','43',NULL,'4',NULL,NULL,'get',NULL,'1608104020439','system','1608104020440',NULL,NULL,'0','0','/api-manager/manager/channel');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111435803824130','43',NULL,'4',NULL,NULL,'get',NULL,'1608104020440','system','1608104020440',NULL,NULL,'0','0','/api-manager/manager/channel/select');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111435803824131','43',NULL,'4',NULL,NULL,'get',NULL,'1608104020441','system','1608104020441',NULL,NULL,'0','0','/api-manager/manager/channel/{channelId}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111435808018433','43',NULL,'4',NULL,NULL,'get',NULL,'1608104020441','system','1608104020441',NULL,NULL,'0','0','/api-manager/manager/treeChannel');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111435808018434','43',NULL,'4',NULL,NULL,'get',NULL,'1608104020442','system','1608104020442',NULL,NULL,'0','0','/api-manager/manager/treeChannel/{channelId}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111437129224194','13',NULL,'4',NULL,NULL,'get',NULL,'1608104020757','system','1608104020758',NULL,NULL,'0','0','/api-manager/manager/app-api');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111437137612801','13',NULL,'4',NULL,NULL,'get',NULL,'1608104020759','system','1608104020759',NULL,NULL,'0','0','/api-manager/manager/api/cancel-auth/{apiId}/{appId}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111437141807105','13',NULL,'4',NULL,NULL,'get',NULL,'1608104020759','system','1608104020759',NULL,NULL,'0','0','/api-datasource/metadata/datasources');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111437146001410','13',NULL,'4',NULL,NULL,'get',NULL,'1608104020760','system','1608104020760',NULL,NULL,'0','0','/api-manager/manager/api/{apiId}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111437146001411','13',NULL,'4',NULL,NULL,'get',NULL,'1608104020760','system','1608104020761',NULL,NULL,'0','0','/api-manager/manager/gateway-admin/statistic/day');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111438408486914','47',NULL,'4',NULL,NULL,'get',NULL,'1608104021061','system','1608104021062',NULL,NULL,'0','0','/api-manager/manager/api-apply');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111438412681218','47',NULL,'4',NULL,NULL,'get',NULL,'1608104021063','system','1608104021063',NULL,NULL,'0','0','/api-manager/manager/api-apply/{apiId}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111438416875521','47',NULL,'4',NULL,NULL,'get',NULL,'1608104021064','system','1608104021064',NULL,NULL,'0','0','/api-manager/manager/api-apply/approve');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111438421069825','47',NULL,'4',NULL,NULL,'get',NULL,'1608104021064','system','1608104021065',NULL,NULL,'0','0','/api-manager/manager/api-apply/reject');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111439574503426','10',NULL,'4',NULL,NULL,'get',NULL,'1608104021339','system','1608104021340',NULL,NULL,'0','0','/api-manager/manager/app/private');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111439578697729','10',NULL,'4',NULL,NULL,'get',NULL,'1608104021341','system','1608104021341',NULL,NULL,'0','0','/api-manager/manager/api-market');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111439582892033','10',NULL,'4',NULL,NULL,'get',NULL,'1608104021341','system','1608104021341',NULL,NULL,'0','0','/api-manager/manager/channel/select');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111439582892034','10',NULL,'4',NULL,NULL,'get',NULL,'1608104021342','system','1608104021342',NULL,NULL,'0','0','/api-manager/manager/api-apply');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111440740519937','18',NULL,'4',NULL,NULL,'get',NULL,'1608104021618','system','1608104021619',NULL,NULL,'0','0','/api-auth/tenant');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111440748908546','18',NULL,'4',NULL,NULL,'get',NULL,'1608104021619','system','1608104021619',NULL,NULL,'0','0','/api-auth/tenant/addAdmin');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111440748908547','18',NULL,'4',NULL,NULL,'get',NULL,'1608104021620','system','1608104021620',NULL,NULL,'0','0','/api-auth/tenant/add');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111441814261762','77',NULL,'4',NULL,NULL,'get',NULL,'1608104021874','system','1608104021874',NULL,NULL,'0','0','/api-auth/user');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111441818456065','77',NULL,'4',NULL,NULL,'get',NULL,'1608104021875','system','1608104021875',NULL,NULL,'0','0','/api-auth/user/delete');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111441822650369','77',NULL,'4',NULL,NULL,'get',NULL,'1608104021875','system','1608104021875',NULL,NULL,'0','0','/api-auth/user/update');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111441822650370','77',NULL,'4',NULL,NULL,'get',NULL,'1608104021875','system','1608104021876',NULL,NULL,'0','0','/api-auth/role/addAdmin');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111442963501057','48',NULL,'4',NULL,NULL,'get',NULL,'1608104022149','system','1608104022150',NULL,NULL,'0','0','/api-auth/cluster/getAll');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111442976083969','48',NULL,'4',NULL,NULL,'get',NULL,'1608104022150','system','1608104022151',NULL,NULL,'0','0','/api-auth/cluster/{id}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111442976083970','48',NULL,'4',NULL,NULL,'get',NULL,'1608104022151','system','1608104022151',NULL,NULL,'0','0','/api-auth/tenant');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111442980278274','48',NULL,'4',NULL,NULL,'get',NULL,'1608104022151','system','1608104022151',NULL,NULL,'0','0','/api-auth/cluster');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111444070797313','75',NULL,'4',NULL,NULL,'get',NULL,'1608104022412','system','1608104022412',NULL,NULL,'0','0','/api-auth/user');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111444074991618','75',NULL,'4',NULL,NULL,'get',NULL,'1608104022413','system','1608104022413',NULL,NULL,'0','0','/api-auth/tenant/getTenantList');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111444079185922','75',NULL,'4',NULL,NULL,'get',NULL,'1608104022413','system','1608104022413',NULL,NULL,'0','0','/api-auth/user/add');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111444079185923','75',NULL,'4',NULL,NULL,'get',NULL,'1608104022414','system','1608104022414',NULL,NULL,'0','0','/api-auth/user/deleteMarket');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111444083380225','75',NULL,'4',NULL,NULL,'get',NULL,'1608104022414','system','1608104022414',NULL,NULL,'0','0','/api-auth/user/update');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111445320699906','42',NULL,'4',NULL,NULL,'get',NULL,'1608104022710','system','1608104022711',NULL,NULL,'0','0','/api-manager/manager/ds');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111445329088513','42',NULL,'4',NULL,NULL,'get',NULL,'1608104022711','system','1608104022712',NULL,NULL,'0','0','/api-auth/user/getUserByTenant');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111445333282817','42',NULL,'4',NULL,NULL,'get',NULL,'1608104022713','system','1608104022713',NULL,NULL,'0','0','/api-datasource/metadata/datasources');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111446356692994','12',NULL,'4',NULL,NULL,'get',NULL,'1608104022957','system','1608104022958',NULL,NULL,'0','0','/api-manager/manager/app/private');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111446365081601','12',NULL,'4',NULL,NULL,'get',NULL,'1608104022959','system','1608104022959',NULL,NULL,'0','0','/api-manager/manager/app/api-used/{apiId}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111446373470209','12',NULL,'4',NULL,NULL,'get',NULL,'1608104022960','system','1608104022961',NULL,NULL,'0','0','/api-manager/manager/app/{apiId}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111446377664513','12',NULL,'4',NULL,NULL,'get',NULL,'1608104022962','system','1608104022962',NULL,NULL,'0','0','/api-manager/manager/app');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111447484960770','76',NULL,'4',NULL,NULL,'get',NULL,'1608104023225','system','1608104023226',NULL,NULL,'0','0','/api-auth/user');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111447489155073','76',NULL,'4',NULL,NULL,'get',NULL,'1608104023227','system','1608104023227',NULL,NULL,'0','0','/api-auth/tenant/getMarketTenant');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111447493349377','76',NULL,'4',NULL,NULL,'get',NULL,'1608104023228','system','1608104023228',NULL,NULL,'0','0','/api-auth/user/add');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111447497543682','76',NULL,'4',NULL,NULL,'get',NULL,'1608104023229','system','1608104023229',NULL,NULL,'0','0','/api-auth/user/deleteMarket');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111447505932290','76',NULL,'4',NULL,NULL,'get',NULL,'1608104023230','system','1608104023231',NULL,NULL,'0','0','/api-auth/user/update');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111448684531713','73',NULL,'4',NULL,NULL,'get',NULL,'1608104023512','system','1608104023513',NULL,NULL,'0','0','/api-hermes/metadata/findFinkTableList');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111448692920322','73',NULL,'4',NULL,NULL,'get',NULL,'1608104023513','system','1608104023513',NULL,NULL,'0','0','/api-datasource/metadata/datasources');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111448692920323','73',NULL,'4',NULL,NULL,'get',NULL,'1608104023514','system','1608104023514',NULL,NULL,'0','0','/api-datasource/metadata/tables/kafaka_ds');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111448697114625','73',NULL,'4',NULL,NULL,'get',NULL,'1608104023514','system','1608104023514',NULL,NULL,'0','0','/api-hermes/metadata/findFinkTableType');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111448697114626','73',NULL,'4',NULL,NULL,'get',NULL,'1608104023514','system','1608104023514',NULL,NULL,'0','0','/api-hermes/metadata/createFlinkTable');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111448701308930','73',NULL,'4',NULL,NULL,'get',NULL,'1608104023515','system','1608104023515',NULL,NULL,'0','0','/api-hermes/metadata/selectFlinkTableById');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111448701308931','73',NULL,'4',NULL,NULL,'get',NULL,'1608104023515','system','1608104023515',NULL,NULL,'0','0','/api-hermes/metadata/updateFlinkTableById');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111448701308932','73',NULL,'4',NULL,NULL,'get',NULL,'1608104023516','system','1608104023516',NULL,NULL,'0','0','/api-hermes/metadata/deleteFlinkTable');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450311921665','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023900','system','1608104023900',NULL,NULL,'0','0','/api-hermes/taskInfo/findAllTask');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450316115969','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023901','system','1608104023901',NULL,NULL,'0','0','/api-hermes/taskInfo/updateTaskStatus');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450320310273','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023901','system','1608104023901',NULL,NULL,'0','0','/api-hermes/taskInfo/deleteTaskById');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450320310274','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023901','system','1608104023901',NULL,NULL,'0','0','/api-auth/cluster');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450320310275','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023902','system','1608104023902',NULL,NULL,'0','0','/api-hermes/taskInfo/selectTaskById');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450324504578','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023902','system','1608104023903',NULL,NULL,'0','0','/api-file/upload/file');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450324504579','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023903','system','1608104023903',NULL,NULL,'0','0','/api-hermes/taskInfo/saveFlinkTask');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450328698882','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023903','system','1608104023903',NULL,NULL,'0','0','/api-hermes/taskInfo/updateFlinkTaskById');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450328698883','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023904','system','1608104023904',NULL,NULL,'0','0','/api-datasource/metadata/datasources');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450332893186','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023905','system','1608104023905',NULL,NULL,'0','0','/api-datasource/metadata/tables/kafaka_ds');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450337087489','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023905','system','1608104023905',NULL,NULL,'0','0','/api-hermes/metadata/selectTableListByDataSource');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450341281794','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023906','system','1608104023906',NULL,NULL,'0','0','/api-hermes/metadata/selectFlinkTableById');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450341281795','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023907','system','1608104023907',NULL,NULL,'0','0','/api-hermes/metadata/selectRedisTableListByDataSource');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450345476097','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023907','system','1608104023907',NULL,NULL,'0','0','/api-datasource/metadata/fields/{db}/{table}');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450345476098','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023908','system','1608104023908',NULL,NULL,'0','0','/api-hermes/metadata/selectHbaseByDataSource');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450345476099','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023908','system','1608104023909',NULL,NULL,'0','0','/api-hermes-cron/flinkSqlTask/validateSql');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450353864706','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023910','system','1608104023910',NULL,NULL,'0','0','/api-hermes/taskInfo/saveFlinkSqlTask');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450353864707','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023911','system','1608104023911',NULL,NULL,'0','0','/api-hermes/taskInfo/updateFlinkSqlTaskById');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450358059009','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023911','system','1608104023911',NULL,NULL,'0','0','/api-hermes-cron/flinkSqlTask/alterTable');
insert into `resource` (`id`, `parent_id`, `code`, `type`, `name`, `url`, `method`, `description`, `created_time`, `created_by`, `updated_time`, `updated_by`, `order`, `can_set`, `is_pub`, `request_url`) values('1339111450366447618','74',NULL,'4',NULL,NULL,'get',NULL,'1608104023912','system','1608104023912',NULL,NULL,'0','0','/api-file/download/file');
