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