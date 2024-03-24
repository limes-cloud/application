/*
 Navicat Premium Data Transfer

 Source Server         : dev
 Source Server Type    : MySQL
 Source Server Version : 80200
 Source Host           : localhost:3306
 Source Schema         : user_center

 Target Server Type    : MySQL
 Target Server Version : 80200
 File Encoding         : 65001

 Date: 24/03/2024 17:08:17
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for agreement_content
-- ----------------------------
DROP TABLE IF EXISTS `agreement_content`;
CREATE TABLE `agreement_content` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `name` varchar(32) NOT NULL COMMENT '协议名称',
  `status` tinyint(1) NOT NULL COMMENT '协议状态',
  `content` blob NOT NULL COMMENT '协议内容',
  `description` varchar(128) NOT NULL COMMENT '协议描述',
  PRIMARY KEY (`id`),
  KEY `idx_agreement_content_created_at` (`created_at`),
  KEY `idx_agreement_content_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4  COMMENT='协议内容';

-- ----------------------------
-- Records of agreement_content
-- ----------------------------
BEGIN;
INSERT INTO `agreement_content` VALUES (1, 1708614504, 1708614504, '测试协议', 1, 0x3C703EE6B58BE8AF95E58D8FE8AEAEE6B58BE8AF95E58D8FE8AEAEE6B58BE8AF95E58D8FE8AEAEE6B58BE8AF95E58D8FE8AEAE3C2F703E0A3C703EE6B58BE8AF95E58D8FE8AEAEE6B58BE8AF95E58D8FE8AEAEE6B58BE8AF95E58D8FE8AEAE3C2F703E, '测试协议测试协议');
COMMIT;

-- ----------------------------
-- Table structure for agreement_scene
-- ----------------------------
DROP TABLE IF EXISTS `agreement_scene`;
CREATE TABLE `agreement_scene` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `keyword` varchar(32) NOT NULL COMMENT '场景标识',
  `name` varchar(32) NOT NULL COMMENT '场景名称',
  `description` varchar(128) NOT NULL COMMENT '场景描述',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`),
  KEY `idx_agreement_scene_created_at` (`created_at`),
  KEY `idx_agreement_scene_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4  COMMENT='协议场景';

-- ----------------------------
-- Records of agreement_scene
-- ----------------------------
BEGIN;
INSERT INTO `agreement_scene` VALUES (1, 1708616170, 1708617119, 'login', '登录', '登录场景');
COMMIT;

-- ----------------------------
-- Table structure for agreement_scene_content
-- ----------------------------
DROP TABLE IF EXISTS `agreement_scene_content`;
CREATE TABLE `agreement_scene_content` (
  `scene_id` int unsigned DEFAULT NULL,
  `agreement_id` int unsigned DEFAULT NULL,
  `content_id` int unsigned NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`content_id`),
  KEY `fk_agreement_scene_content_scene` (`scene_id`),
  CONSTRAINT `fk_agreement_scene_content_content` FOREIGN KEY (`content_id`) REFERENCES `agreement_content` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_agreement_scene_content_scene` FOREIGN KEY (`scene_id`) REFERENCES `agreement_scene` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_agreement_scene_scene_contents` FOREIGN KEY (`scene_id`) REFERENCES `agreement_scene` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='协议场景-内容';

-- ----------------------------
-- Records of agreement_scene_content
-- ----------------------------
BEGIN;
INSERT INTO `agreement_scene_content` VALUES (1, NULL, 1);
COMMIT;

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `keyword` varchar(32) NOT NULL COMMENT '应用标识',
  `logo` varchar(128) NOT NULL COMMENT '应用logo',
  `name` varchar(32) NOT NULL COMMENT '应用名称',
  `status` tinyint(1) NOT NULL COMMENT '应用状态',
  `user_fields` varchar(1024) DEFAULT NULL COMMENT '用户字段',
  `version` varchar(32) DEFAULT NULL COMMENT '应用版本',
  `copyright` varchar(128) DEFAULT NULL COMMENT '应用版权',
  `allow_registry` tinyint(1) NOT NULL COMMENT '是否允许注册',
  `description` varchar(128) DEFAULT NULL COMMENT '应用描述',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`),
  KEY `idx_app_created_at` (`created_at`),
  KEY `idx_app_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4  COMMENT='应用信息';

-- ----------------------------
-- Records of app
-- ----------------------------
BEGIN;
INSERT INTO `app` VALUES (2, 1708708901, 1710054434, 'party-affairs', '24acef3dbe2cc8eb776008fc133e4f73338a3644a6581763f35f3ffc71d22641', '引路灯', 1, NULL, '', '', 0, '');
COMMIT;

-- ----------------------------
-- Table structure for app_channel
-- ----------------------------
DROP TABLE IF EXISTS `app_channel`;
CREATE TABLE `app_channel` (
  `app_id` int unsigned NOT NULL COMMENT '主键ID',
  `channel_id` int unsigned NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`app_id`,`channel_id`),
  KEY `fk_app_channel_channel` (`channel_id`),
  CONSTRAINT `fk_app_channel_app` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_app_channel_channel` FOREIGN KEY (`channel_id`) REFERENCES `channel` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='应用信息';

-- ----------------------------
-- Records of app_channel
-- ----------------------------
BEGIN;
INSERT INTO `app_channel` VALUES (2, 2);
INSERT INTO `app_channel` VALUES (2, 3);
INSERT INTO `app_channel` VALUES (2, 4);
COMMIT;

-- ----------------------------
-- Table structure for app_field
-- ----------------------------
DROP TABLE IF EXISTS `app_field`;
CREATE TABLE `app_field` (
  `app_id` int unsigned NOT NULL COMMENT '主键ID',
  `field_id` int unsigned NOT NULL COMMENT '主键ID',
  PRIMARY KEY (`app_id`,`field_id`),
  KEY `fk_app_field_field` (`field_id`),
  CONSTRAINT `fk_app_field_app` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_app_field_field` FOREIGN KEY (`field_id`) REFERENCES `field` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='应用信息';

-- ----------------------------
-- Records of app_field
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for auth
-- ----------------------------
DROP TABLE IF EXISTS `auth`;
CREATE TABLE `auth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `user_id` int unsigned NOT NULL COMMENT '用户id',
  `channel_id` int unsigned NOT NULL COMMENT '渠道id',
  `auth_id` varchar(64) DEFAULT NULL COMMENT '渠道授权ID',
  `union_id` varchar(64) DEFAULT NULL COMMENT '渠道联合ID',
  `channel_token` varchar(64) DEFAULT NULL COMMENT '渠道token',
  `channel_expire_at` bigint DEFAULT NULL COMMENT '渠道token过期时间',
  `jwt_token` varchar(1024) DEFAULT NULL COMMENT '平台Token',
  `jwt_expire_at` bigint DEFAULT NULL COMMENT '过期时间',
  `login_at` bigint DEFAULT NULL COMMENT '最近登录时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uc` (`user_id`,`channel_id`),
  UNIQUE KEY `ua` (`channel_id`,`auth_id`),
  KEY `idx_auth_created_at` (`created_at`),
  CONSTRAINT `fk_auth_channel` FOREIGN KEY (`channel_id`) REFERENCES `channel` (`id`),
  CONSTRAINT `fk_auth_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_user_auths` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4  COMMENT='用户授权';

-- ----------------------------
-- Records of auth
-- ----------------------------
BEGIN;
INSERT INTO `auth` VALUES (1, 1708753245, 1, 3, NULL, NULL, NULL, 0, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfaWQiOjIsImFwcF9rZXl3b3JkIjoicGFydHktYWZmYWlycyIsImNoYW5uZWxfaWQiOjMsImV4cCI6MTcxMDA1OTU0OSwiaWF0IjoxNzEwMDUyMzQ4LCJuYmYiOjE3MTAwNTIzNDgsInVzZXJfaWQiOjF9.EHKk720Hcz4Wa0RSq72Sm5UX0M5qxnqlGZ3Tgc9fh9E', 1710059548, 1710052348);
INSERT INTO `auth` VALUES (3, 1708754544, 3, 3, NULL, NULL, NULL, 0, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfaWQiOjIsImFwcF9rZXl3b3JkIjoicGFydHktYWZmYWlycyIsImNoYW5uZWxfaWQiOjMsImV4cCI6MTcwODk2NjQ1NSwiaWF0IjoxNzA4OTU5MjU0LCJuYmYiOjE3MDg5NTkyNTQsInVzZXJfaWQiOjN9.BBUz7BojtcAxME4nC-Y47Im1yjR8w-5fXfU9H40jFAU', 1708966454, 1708959254);
INSERT INTO `auth` VALUES (7, 1710054591, 1, 4, '17871318', NULL, NULL, 0, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfaWQiOjIsImFwcF9rZXl3b3JkIjoicGFydHktYWZmYWlycyIsImNoYW5uZWxfaWQiOjQsImV4cCI6MTcxMDI0MjE3MCwiaWF0IjoxNzEwMjM0OTY5LCJuYmYiOjE3MTAyMzQ5NjksInVzZXJfaWQiOjF9.P9CpWL6kp2ch2lnuR7GgFddZmWgjXqew1dAhxm_06U0', 1710242169, 1710234969);
COMMIT;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for channel
-- ----------------------------
DROP TABLE IF EXISTS `channel`;
CREATE TABLE `channel` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `logo` varchar(128) NOT NULL COMMENT '渠道logo',
  `platform` char(32) NOT NULL COMMENT '渠道标识',
  `name` varchar(32) NOT NULL COMMENT '渠道名称',
  `ak` varchar(32) DEFAULT NULL COMMENT '渠道ak',
  `sk` varchar(32) DEFAULT NULL COMMENT '渠道sk',
  `extra` varchar(256) DEFAULT NULL COMMENT '渠道状态',
  `status` tinyint(1) NOT NULL COMMENT '渠道状态',
  PRIMARY KEY (`id`),
  UNIQUE KEY `platform` (`platform`),
  KEY `idx_channel_created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4  COMMENT='授权渠道';

-- ----------------------------
-- Records of channel
-- ----------------------------
BEGIN;
INSERT INTO `channel` VALUES (2, 1708622662, '5118bb9a26458eed86525a00b02c8bba8299dfa98244239858c50b0be431a069', 'captcha', '验证码', '', '', '', 1);
INSERT INTO `channel` VALUES (3, 1708622671, '97252c871797d84d7f582df166ed07711834ec0c675a0d171df39770f3a93960', 'password', '密码', '', '', '', 1);
INSERT INTO `channel` VALUES (4, 1708622689, '5118bb9a26458eed86525a00b02c8bba8299dfa98244239858c50b0be431a069', 'yb', '易班', 'e4750b34230b48e1', 'b0891a7f6018e5a76b085e3cb9548edd', ' ', 1);
COMMIT;

-- ----------------------------
-- Table structure for field
-- ----------------------------
DROP TABLE IF EXISTS `field`;
CREATE TABLE `field` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `keyword` varchar(32) NOT NULL COMMENT '字段标识',
  `type` varchar(32) NOT NULL COMMENT '字段类型',
  `name` varchar(32) NOT NULL COMMENT '字段名称',
  `description` varchar(128) DEFAULT NULL COMMENT '字段描述',
  PRIMARY KEY (`id`),
  UNIQUE KEY `keyword` (`keyword`),
  KEY `idx_field_updated_at` (`updated_at`),
  KEY `idx_field_created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4  COMMENT='应用信息';

-- ----------------------------
-- Records of field
-- ----------------------------
BEGIN;
INSERT INTO `field` VALUES (1, 1708658338, 1708658353, 'test', 'bool', '测试字段1', 'test1');
COMMIT;

-- ----------------------------
-- Table structure for gorm_init
-- ----------------------------
DROP TABLE IF EXISTS `gorm_init`;
CREATE TABLE `gorm_init` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `init` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ;

-- ----------------------------
-- Records of gorm_init
-- ----------------------------
BEGIN;
INSERT INTO `gorm_init` VALUES (1, 0);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `phone` char(15) DEFAULT NULL COMMENT '手机',
  `email` varchar(64) DEFAULT NULL COMMENT '邮箱',
  `username` char(32) DEFAULT NULL COMMENT '账号',
  `password` varchar(256) DEFAULT NULL COMMENT '密码',
  `nick_name` varchar(32) DEFAULT NULL COMMENT '昵称',
  `real_name` varchar(32) DEFAULT NULL COMMENT '真实姓名',
  `avatar` varchar(128) DEFAULT NULL COMMENT '头像',
  `gender` enum('F','M','U') DEFAULT 'U' COMMENT '昵称',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `disable_desc` varchar(128) DEFAULT NULL COMMENT '禁用原因',
  `from` varchar(128) NOT NULL COMMENT '用户来源标识',
  `from_desc` varchar(128) NOT NULL COMMENT '用户来源',
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_user_created_at` (`created_at`),
  KEY `idx_user_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4  COMMENT='用户信息';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 1708753245, 1710234957, NULL, '1280291001@qq.com', NULL, '', '456', '123', '', 'U', 1, NULL, 'party-affairs', '引路灯');
INSERT INTO `user` VALUES (3, 1708754544, 1708958375, NULL, NULL, 'asd123', '$2a$10$fPevC.gl5h/HO5yLC4VDh.eGFSJcEUpRa71EUvRO7BFEpfmbMaUsK', '1', '1', '', 'M', 1, NULL, 'party-affairs', '引路灯');
INSERT INTO `user` VALUES (4, 1708872985, 1708872985, '18286219255', '615@qq.com', NULL, '', '用户1', '用户1', '', 'M', 1, NULL, '', '');
INSERT INTO `user` VALUES (5, 1708872985, 1708872985, '18286219256', '616@qq.com', NULL, '', '用户2', '用户2', '', 'F', 1, NULL, '', '');
COMMIT;

-- ----------------------------
-- Table structure for user_app
-- ----------------------------
DROP TABLE IF EXISTS `user_app`;
CREATE TABLE `user_app` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `user_id` int unsigned NOT NULL COMMENT '用户id',
  `app_id` int unsigned NOT NULL COMMENT '应用id',
  `login_at` bigint DEFAULT NULL COMMENT '最近登录',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ua` (`user_id`,`app_id`),
  KEY `idx_user_app_created_at` (`created_at`),
  KEY `fk_user_app_app` (`app_id`),
  CONSTRAINT `fk_user_app_app` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`),
  CONSTRAINT `fk_user_user_apps` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4  COMMENT='用户应用';

-- ----------------------------
-- Records of user_app
-- ----------------------------
BEGIN;
INSERT INTO `user_app` VALUES (1, 1708754544, 3, 2, 1710051764);
INSERT INTO `user_app` VALUES (5, 1710052290, 1, 2, 1710234969);
COMMIT;

-- ----------------------------
-- Table structure for user_extra
-- ----------------------------
DROP TABLE IF EXISTS `user_extra`;
CREATE TABLE `user_extra` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `user_id` int unsigned NOT NULL COMMENT '用户id',
  `keyword` varchar(32) NOT NULL COMMENT '关键字',
  `value` varchar(1024) NOT NULL COMMENT '扩展值',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk` (`user_id`,`keyword`),
  KEY `idx_user_extra_created_at` (`created_at`),
  CONSTRAINT `fk_user_user_extras` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='用户扩展信息';

-- ----------------------------
-- Records of user_extra
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
