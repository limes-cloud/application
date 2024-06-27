/*
 Navicat Premium Data Transfer

 Source Server     : dev
 Source Server Type    : MySQL
 Source Server Version : 80200
 Source Host       : localhost:3306
 Source Schema     : usercenter

 Target Server Type    : MySQL
 Target Server Version : 80200
 File Encoding     : 65001

 Date: 27/06/2024 16:00:50
*/

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `logo` varchar(128) NOT NULL COMMENT '图标',
    `keyword` char(32) NOT NULL COMMENT '标识',
    `name` varchar(32) NOT NULL COMMENT '名称',
    `status` tinyint(1) NOT NULL COMMENT '状态',
    `version` varchar(32) DEFAULT NULL COMMENT '版本',
    `copyright` varchar(128) DEFAULT NULL COMMENT '版权',
    `extra` tinytext COMMENT '扩展信息',
    `allow_registry` tinyint(1) NOT NULL COMMENT '允许注册',
    `description` varchar(128) DEFAULT NULL COMMENT '描述',
    `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
    `updated_at` bigint unsigned NOT NULL COMMENT '修改时间',
    `disable_desc` tinytext COMMENT '禁用原因',
    PRIMARY KEY (`id`),
    UNIQUE KEY `keyword` (`keyword`),
    KEY `idx_app_created_at` (`created_at`),
    KEY `idx_app_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='应用信息';

-- ----------------------------
-- Records of app
-- ----------------------------
BEGIN;
INSERT INTO `app` VALUES (1, '36e2e87f7b73219343da52a28ba47eec', 'PartyAffairs', '测试应用', 1, 'v1.0', 'lime.qlime.cn', '', 1, '统一应用管理中心示例客户端', 1718438283, 1718771171, NULL);
COMMIT;

-- ----------------------------
-- Table structure for app_channel
-- ----------------------------
DROP TABLE IF EXISTS `app_channel`;
CREATE TABLE `app_channel` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `app_id` bigint unsigned NOT NULL COMMENT '应用id',
    `channel_id` bigint unsigned NOT NULL COMMENT '渠道id',
    PRIMARY KEY (`id`),
    UNIQUE KEY `app_id` (`app_id`,`channel_id`),
    KEY `channel_id` (`channel_id`),
    CONSTRAINT `app_channel_ibfk_1` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`) ON DELETE CASCADE,
    CONSTRAINT `app_channel_ibfk_2` FOREIGN KEY (`channel_id`) REFERENCES `channel` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COMMENT='应用渠道信息';

-- ----------------------------
-- Records of app_channel
-- ----------------------------
BEGIN;
INSERT INTO `app_channel` VALUES (13, 1, 1);
INSERT INTO `app_channel` VALUES (14, 1, 2);
INSERT INTO `app_channel` VALUES (15, 1, 3);
INSERT INTO `app_channel` VALUES (16, 1, 4);
INSERT INTO `app_channel` VALUES (17, 1, 5);
COMMIT;

-- ----------------------------
-- Table structure for app_field
-- ----------------------------
DROP TABLE IF EXISTS `app_field`;
CREATE TABLE `app_field` (
      `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
      `app_id` bigint unsigned NOT NULL COMMENT '应用id',
      `field_id` bigint unsigned NOT NULL COMMENT '字段id',
      PRIMARY KEY (`id`),
      UNIQUE KEY `app_id` (`app_id`,`field_id`),
      KEY `field_id` (`field_id`),
      CONSTRAINT `app_field_ibfk_1` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`) ON DELETE CASCADE,
      CONSTRAINT `app_field_ibfk_2` FOREIGN KEY (`field_id`) REFERENCES `field` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='应用字段信息';

-- ----------------------------
-- Records of app_field
-- ----------------------------
BEGIN;
INSERT INTO `app_field` VALUES (7, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for auth
-- ----------------------------
DROP TABLE IF EXISTS `auth`;
CREATE TABLE `auth` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
     `user_id` bigint unsigned NOT NULL,
     `app_id` bigint unsigned NOT NULL,
     `status` tinyint(1) DEFAULT '1' COMMENT '状态',
     `disable_desc` varchar(128) DEFAULT NULL COMMENT '禁用原因',
     `token` varchar(512) DEFAULT NULL COMMENT '用户token',
     `logged_at` bigint NOT NULL DEFAULT '0' COMMENT '登陆时间',
     `expired_at` bigint NOT NULL DEFAULT '0' COMMENT '过期时间',
     `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
     PRIMARY KEY (`id`),
     UNIQUE KEY `ua` (`user_id`,`app_id`),
     KEY `created_at` (`created_at`),
     KEY `app_id` (`app_id`),
     CONSTRAINT `auth_ibfk_1` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`),
     CONSTRAINT `auth_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COMMENT='用户登录信息';

-- ----------------------------
-- Records of auth
-- ----------------------------
BEGIN;
INSERT INTO `auth` VALUES (5, 7, 1, 0, '1', '', 0, 0, 1718608014);
INSERT INTO `auth` VALUES (10, 12, 1, 1, NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBLZXl3b3JkIjoiUGFydHlBZmZhaXJzIiwiZXhwIjoxNzE5MDY2OTk5LCJpYXQiOjE3MTkwNTk3OTgsIm5iZiI6MTcxOTA1OTc5OCwidXNlcklkIjoxMn0.UkY0-W46tpl1uhMBi1V2Jbt3MFe-0T2YeizHSUMICsI', 1719059798, 1719066998, 1718947235);
INSERT INTO `auth` VALUES (12, 14, 1, 1, NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBLZXl3b3JkIjoiUGFydHlBZmZhaXJzIiwiZXhwIjoxNzE4OTU5NTg3LCJpYXQiOjE3MTg5NTIzODYsIm5iZiI6MTcxODk1MjM4NiwidXNlcklkIjoxNH0.Aii44hgDh6tI9YqtvYk5RwihU1vYQDVotwniF9d3fBw', 1718952386, 1718959586, 1718952345);
INSERT INTO `auth` VALUES (13, 15, 1, 1, NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBLZXl3b3JkIjoiUGFydHlBZmZhaXJzIiwiZXhwIjoxNzE4OTYwMTc0LCJpYXQiOjE3MTg5NTI5NzMsIm5iZiI6MTcxODk1Mjk3MywidXNlcklkIjoxNX0.Wj4Z0474i3_sbiE7s8Qfq3MJuHxxfkrIdhoMtTAy59c', 1718952973, 1718960173, 1718952973);
COMMIT;

-- ----------------------------
-- Table structure for channel
-- ----------------------------
DROP TABLE IF EXISTS `channel`;
CREATE TABLE `channel` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `logo` varchar(128) NOT NULL COMMENT '图标',
    `keyword` char(32) NOT NULL COMMENT '标识',
    `name` varchar(32) NOT NULL COMMENT '名称',
    `ak` varchar(32) DEFAULT NULL COMMENT 'ak',
    `sk` varchar(32) DEFAULT NULL COMMENT 'sk',
    `extra` tinytext COMMENT '扩展信息',
    `status` tinyint(1) NOT NULL COMMENT '渠道状态',
    `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
    `updated_at` bigint unsigned NOT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `keyword` (`keyword`),
    KEY `idx_channel_created_at` (`created_at`),
    KEY `idx_channel_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='授权渠道';

-- ----------------------------
-- Records of channel
-- ----------------------------
BEGIN;
INSERT INTO `channel` VALUES (1, '2a0786fe9127b8116bc30ed2ce9581e2', 'password', '密码', '1', NULL, NULL, 1, 1718392194, 1718393021);
INSERT INTO `channel` VALUES (2, '1f27444925877922d71110d993edf590', 'yb', '易班', NULL, NULL, NULL, 1, 1718702913, 1718702916);
INSERT INTO `channel` VALUES (3, '385d37202ae8f08cd8ba429eb51b5422', 'email', '邮箱', NULL, NULL, NULL, 1, 1718702924, 1718793318);
INSERT INTO `channel` VALUES (4, '2252554cf6309d2e53e95a5d40458d17', 'mp', '微信', 'wx819e8d8e719671ee', '740862670088146f3920e206963bd77f', NULL, 1, 1718731281, 1718731472);
INSERT INTO `channel` VALUES (5, '49dc09f716382dd3f460daaba2649939', 'qq', 'QQ', '1109922231', 'I7UHeWvpYXCZGbUo', NULL, 1, 1718731468, 1718731474);
COMMIT;

-- ----------------------------
-- Table structure for extra
-- ----------------------------
DROP TABLE IF EXISTS `extra`;
CREATE TABLE `extra` (
      `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
      `user_id` bigint unsigned NOT NULL COMMENT '用户',
      `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '标识',
      `value` varchar(1024) NOT NULL COMMENT '值',
      `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
      `updated_at` bigint unsigned NOT NULL COMMENT '修改时间',
      PRIMARY KEY (`id`),
      UNIQUE KEY `uk` (`user_id`,`keyword`),
      KEY `keyword` (`keyword`),
      CONSTRAINT `user_field_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
      CONSTRAINT `user_field_ibfk_2` FOREIGN KEY (`keyword`) REFERENCES `field` (`keyword`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户字段信息';

-- ----------------------------
-- Records of extra
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for field
-- ----------------------------
DROP TABLE IF EXISTS `field`;
CREATE TABLE `field` (
      `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
      `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '字段标识',
      `type` char(32) NOT NULL COMMENT '字段类型',
      `name` varchar(64) NOT NULL COMMENT '字段名称',
      `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '字段状态',
      `description` varchar(128) DEFAULT NULL COMMENT '字段描述',
      `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
      `updated_at` bigint unsigned NOT NULL COMMENT '修改时间',
      PRIMARY KEY (`id`),
      UNIQUE KEY `keyword` (`keyword`),
      KEY `idx_field_updated_at` (`updated_at`),
      KEY `idx_field_created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='字段信息';

-- ----------------------------
-- Records of field
-- ----------------------------
BEGIN;
INSERT INTO `field` VALUES (1, 'name', 'string', '名称', 1, '存储名字', 1718432322, 1718432326);
COMMIT;

-- ----------------------------
-- Table structure for gorm_init
-- ----------------------------
DROP TABLE IF EXISTS `gorm_init`;
CREATE TABLE `gorm_init` (
      `id` int unsigned NOT NULL AUTO_INCREMENT,
      `init` tinyint(1) DEFAULT NULL,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of gorm_init
-- ----------------------------
BEGIN;
INSERT INTO `gorm_init` VALUES (1, 1);
COMMIT;

-- ----------------------------
-- Table structure for oauth
-- ----------------------------
DROP TABLE IF EXISTS `oauth`;
CREATE TABLE `oauth` (
      `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
      `user_id` bigint unsigned DEFAULT NULL,
      `channel_id` bigint unsigned NOT NULL,
      `auth_id` varchar(64) DEFAULT NULL COMMENT '渠道授权ID',
      `union_id` varchar(64) DEFAULT NULL COMMENT '渠道联合ID',
      `token` varchar(1024) DEFAULT NULL COMMENT '渠道token',
      `logged_at` bigint NOT NULL DEFAULT '0' COMMENT '登陆时间',
      `expired_at` bigint NOT NULL DEFAULT '0' COMMENT '过期时间',
      `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
      PRIMARY KEY (`id`),
      UNIQUE KEY `user_id` (`user_id`,`channel_id`),
      UNIQUE KEY `channel_id` (`channel_id`,`auth_id`),
      KEY `created_at` (`created_at`),
      CONSTRAINT `oauth_ibfk_1` FOREIGN KEY (`channel_id`) REFERENCES `channel` (`id`),
      CONSTRAINT `oauth_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COMMENT='用户渠道授权信息';

-- ----------------------------
-- Records of oauth
-- ----------------------------
BEGIN;
INSERT INTO `oauth` VALUES (1, 12, 4, 'oNhHw0CrhR2Kcp2a_CqzpSKtD4k0', '', '81_ucRLKwGtbj3dNGmqtwhAxHQcto9iZvTxz8jFzbXc7aJcsahUFlJacgpMwaWO3wEZiQ2t2QP3MHyPjcrYXgEkMGF5dOE2OBfG32YilPJG3b1zdrN49rI18aTSksgPLVaABAWSC', 1718992610, 1, 1718990354);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
     `phone` char(15) DEFAULT NULL COMMENT '手机',
     `email` varchar(64) DEFAULT NULL COMMENT '邮箱',
     `username` char(32) DEFAULT NULL COMMENT '账号',
     `password` varchar(256) DEFAULT NULL COMMENT '密码',
     `nick_name` varchar(32) DEFAULT NULL COMMENT '昵称',
     `real_name` varchar(32) DEFAULT NULL COMMENT '真实姓名',
     `avatar` varchar(128) DEFAULT NULL COMMENT '头像',
     `gender` enum('F','M','U') DEFAULT 'U' COMMENT '性别',
     `status` tinyint(1) DEFAULT '1' COMMENT '状态',
     `disable_desc` varchar(128) DEFAULT NULL COMMENT '禁用原因',
     `from` varchar(128) NOT NULL COMMENT '来源标识',
     `from_desc` varchar(128) NOT NULL COMMENT '来源描述',
     `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
     `updated_at` bigint unsigned NOT NULL COMMENT '修改时间',
     `deleted_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
     PRIMARY KEY (`id`),
     UNIQUE KEY `phone` (`phone`,`deleted_at`),
     UNIQUE KEY `email` (`email`,`deleted_at`),
     UNIQUE KEY `username` (`username`,`deleted_at`),
     KEY `idx_user_created_at` (`created_at`),
     KEY `idx_user_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (6, '18286219254', NULL, NULL, NULL, '', '测试用户', NULL, 'M', 0, NULL, '', '', 1718463260, 1718464434, 1718464628);
INSERT INTO `user` VALUES (7, '18286219254', NULL, NULL, NULL, 'd9e94004', '测试', NULL, 'M', 1, NULL, '', '', 1718464680, 1718770840, 0);
INSERT INTO `user` VALUES (12, NULL, NULL, 'a12345678', '$2a$10$F9cp1TM96d2AOPgrLTTb1OcQr0zKZ5ZsRa6QQqJ8rx/Cw9lX1JnbG', '8a68e6b3', NULL, 'b373234319cc81c55ddd81b8de001f11', 'M', 1, NULL, 'PartyAffairs', '测试应用', 1718947235, 1719070909, 0);
INSERT INTO `user` VALUES (14, NULL, '1280291001@qq.com', NULL, NULL, '用户-e34b7b85', NULL, NULL, 'F', 1, NULL, 'PartyAffairs', '测试应用', 1718952345, 1718952345, 0);
INSERT INTO `user` VALUES (15, NULL, NULL, 'a123456789', '$2a$10$szw0Hric1UMa86QPH39Y3eGj6z/gF5na2GrkKbvlYUWM5lRtkL1JG', '用户-24883d1a', NULL, NULL, 'F', 1, NULL, 'PartyAffairs', '测试应用', 1718952973, 1718952973, 0);
COMMIT;

-- ----------------------------
-- Table structure for user_app
-- ----------------------------
DROP TABLE IF EXISTS `user_app`;
CREATE TABLE `user_app` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
     `user_id` bigint unsigned NOT NULL,
     `app_id` bigint unsigned NOT NULL,
     `status` tinyint(1) DEFAULT '1' COMMENT '状态',
     `disable_desc` varchar(128) DEFAULT NULL COMMENT '禁用原因',
     `setting` tinytext COMMENT '用户设置',
     `token` varchar(512) DEFAULT NULL COMMENT '用户token',
     `logged_at` bigint NOT NULL DEFAULT '0' COMMENT '登陆时间',
     `expired_at` bigint NOT NULL DEFAULT '0' COMMENT '过期时间',
     `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
     PRIMARY KEY (`id`),
     UNIQUE KEY `ua` (`user_id`,`app_id`),
     KEY `created_at` (`created_at`),
     KEY `fk_user_app_app` (`app_id`),
     CONSTRAINT `fk_user_app_app` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`),
     CONSTRAINT `fk_user_user_apps` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户应用信息';

-- ----------------------------
-- Records of user_app
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for userinfo
-- ----------------------------
DROP TABLE IF EXISTS `userinfo`;
CREATE TABLE `userinfo` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
     `user_id` bigint unsigned NOT NULL COMMENT '用户',
     `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '标识',
     `value` varchar(1024) NOT NULL COMMENT '值',
     `created_at` bigint unsigned NOT NULL COMMENT '创建时间',
     `updated_at` bigint unsigned NOT NULL COMMENT '修改时间',
     PRIMARY KEY (`id`),
     UNIQUE KEY `uk` (`user_id`,`keyword`),
     KEY `keyword` (`keyword`),
     CONSTRAINT `userinfo_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
     CONSTRAINT `userinfo_ibfk_2` FOREIGN KEY (`keyword`) REFERENCES `field` (`keyword`) ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='用户字段信息';

-- ----------------------------
-- Records of userinfo
-- ----------------------------
BEGIN;
INSERT INTO `userinfo` VALUES (2, 7, 'name', '1', 1718627394, 1718627394);
COMMIT;

