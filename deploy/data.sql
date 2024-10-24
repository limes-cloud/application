
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `usercenter`
--

-- --------------------------------------------------------

--
-- 表的结构 `app`
--

CREATE TABLE `app` (
                       `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                       `logo` varchar(128) NOT NULL COMMENT '图标',
                       `keyword` char(32) NOT NULL COMMENT '标识',
                       `name` varchar(32) NOT NULL COMMENT '名称',
                       `status` tinyint(1) NOT NULL COMMENT '状态',
                       `version` varchar(32) DEFAULT NULL COMMENT '版本',
                       `copyright` varchar(128) DEFAULT NULL COMMENT '版权',
                       `extra` tinytext COMMENT '扩展信息',
                       `allow_registry` tinyint(1) NOT NULL COMMENT '允许注册',
                       `description` varchar(128) DEFAULT NULL COMMENT '描述',
                       `created_at` bigint(20) UNSIGNED NOT NULL COMMENT '创建时间',
                       `updated_at` bigint(20) UNSIGNED NOT NULL COMMENT '修改时间',
                       `disable_desc` tinytext COMMENT '禁用原因'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用信息';

--
-- 转存表中的数据 `app`
--

INSERT INTO `app` (`id`, `logo`, `keyword`, `name`, `status`, `version`, `copyright`, `extra`, `allow_registry`, `description`, `created_at`, `updated_at`, `disable_desc`) VALUES
                                                                                                                                                                                (1, 'd88d2daa847ab52c05e0cefa1920ec6f', 'PartyAffairs', '信号灯', 1, 'v1.0', 'Copyright © 2024 qlime.cn. All rights reserved.', '', 0, '党性教育不掉线，组织生活不缺席。', 1718438283, 1729087016, NULL),
                                                                                                                                                                                (2, 'd7772aea6932855cb53f159fd289f63e', 'Poverty', '资助系统', 1, 'v1.0', 'lime.cn', NULL, 0, '资助系统', 1723030370, 1728653368, NULL);

-- --------------------------------------------------------

--
-- 表的结构 `app_channel`
--

CREATE TABLE `app_channel` (
                               `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                               `app_id` bigint(20) UNSIGNED NOT NULL COMMENT '应用id',
                               `channel_id` bigint(20) UNSIGNED NOT NULL COMMENT '渠道id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用渠道信息';

--
-- 转存表中的数据 `app_channel`
--

INSERT INTO `app_channel` (`id`, `app_id`, `channel_id`) VALUES
                                                             (54, 1, 1),
                                                             (55, 1, 2),
                                                             (56, 1, 3),
                                                             (48, 2, 1),
                                                             (49, 2, 2),
                                                             (50, 2, 3);

-- --------------------------------------------------------

--
-- 表的结构 `app_field`
--

CREATE TABLE `app_field` (
                             `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                             `app_id` bigint(20) UNSIGNED NOT NULL COMMENT '应用id',
                             `field_id` bigint(20) UNSIGNED NOT NULL COMMENT '字段id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用字段信息';

-- --------------------------------------------------------

--
-- 表的结构 `auth`
--

CREATE TABLE `auth` (
                        `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                        `user_id` bigint(20) UNSIGNED NOT NULL,
                        `app_id` bigint(20) UNSIGNED NOT NULL,
                        `status` tinyint(1) DEFAULT '1' COMMENT '状态',
                        `disable_desc` varchar(128) DEFAULT NULL COMMENT '禁用原因',
                        `token` varchar(512) DEFAULT NULL COMMENT '用户token',
                        `logged_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '登陆时间',
                        `expired_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '过期时间',
                        `created_at` bigint(20) UNSIGNED NOT NULL COMMENT '创建时间',
                        `setting` text COMMENT '用户设置'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户登录信息';


-- --------------------------------------------------------

--
-- 表的结构 `channel`
--

CREATE TABLE `channel` (
                           `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                           `logo` varchar(128) NOT NULL COMMENT '图标',
                           `keyword` char(32) NOT NULL COMMENT '标识',
                           `name` varchar(32) NOT NULL COMMENT '名称',
                           `ak` varchar(32) DEFAULT NULL COMMENT 'ak',
                           `sk` varchar(32) DEFAULT NULL COMMENT 'sk',
                           `extra` tinytext COMMENT '扩展信息',
                           `status` tinyint(1) NOT NULL COMMENT '渠道状态',
                           `created_at` bigint(20) UNSIGNED NOT NULL COMMENT '创建时间',
                           `updated_at` bigint(20) UNSIGNED NOT NULL COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='授权渠道';

--
-- 转存表中的数据 `channel`
--

INSERT INTO `channel` (`id`, `logo`, `keyword`, `name`, `ak`, `sk`, `extra`, `status`, `created_at`, `updated_at`) VALUES
                                                                                                                       (1, 'd19faa31f4a04b52f802a465edf50d18', 'password', '密码', '', NULL, NULL, 1, 1718392194, 1720365719),
                                                                                                                       (2, '1f27444925877922d71110d993edf590', 'yb', '易班', 'dbd3b98ef20bd760', 'acc44f108e1a0c2e5c63a837462fc7fd', NULL, 1, 1718702913, 1720376982),
                                                                                                                       (3, '385d37202ae8f08cd8ba429eb51b5422', 'email', '邮箱', NULL, NULL, NULL, 1, 1718702924, 1718793318),
                                                                                                                       (4, '2252554cf6309d2e53e95a5d40458d17', 'mp', '微信', 'wx819e8d8e719671ee', '740862670088146f3920e206963bd77f', NULL, 0, 1718731281, 1720424101),
                                                                                                                       (5, '49dc09f716382dd3f460daaba2649939', 'qq', 'QQ', '1109922231', 'I7UHeWvpYXCZGbUo', NULL, 0, 1718731468, 1720424103);

-- --------------------------------------------------------

--
-- 表的结构 `extra`
--

CREATE TABLE `extra` (
                         `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                         `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户',
                         `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '标识',
                         `value` varchar(1024) NOT NULL COMMENT '值',
                         `created_at` bigint(20) UNSIGNED NOT NULL COMMENT '创建时间',
                         `updated_at` bigint(20) UNSIGNED NOT NULL COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户字段信息';

-- --------------------------------------------------------

--
-- 表的结构 `feedback`
--

CREATE TABLE `feedback` (
                            `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                            `app_id` bigint(20) UNSIGNED NOT NULL COMMENT '应用id',
                            `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户id',
                            `category_id` bigint(20) UNSIGNED NOT NULL COMMENT '分类id',
                            `title` varchar(128) NOT NULL COMMENT '标题',
                            `content` text NOT NULL COMMENT '内容',
                            `status` char(32) NOT NULL COMMENT '状态',
                            `images` tinytext COMMENT '图片',
                            `contact` char(32) DEFAULT NULL COMMENT '联系方式',
                            `device` text NOT NULL COMMENT '设备',
                            `platform` char(32) NOT NULL COMMENT '平台',
                            `version` varchar(32) DEFAULT NULL COMMENT '版本',
                            `md5` varchar(64) NOT NULL COMMENT 'md5值',
                            `processed_by` bigint(20) UNSIGNED DEFAULT NULL COMMENT '处理人',
                            `processed_result` varchar(256) DEFAULT NULL COMMENT '处理结果',
                            `created_at` bigint(20) UNSIGNED NOT NULL COMMENT '创建时间',
                            `updated_at` bigint(20) UNSIGNED NOT NULL COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='反馈信息';

-- --------------------------------------------------------

--
-- 表的结构 `feedback_category`
--

CREATE TABLE `feedback_category` (
                                     `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                                     `name` varchar(32) NOT NULL COMMENT '分类名称',
                                     `created_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='反馈分类';

--
-- 转存表中的数据 `feedback_category`
--

INSERT INTO `feedback_category` (`id`, `name`, `created_at`) VALUES
                                                                 (1, '功能异常', 1720025501),
                                                                 (3, '产品建议', 1720025937),
                                                                 (4, '界面错位', 1720879750),
                                                                 (5, '加载缓慢', 1720879761);

-- --------------------------------------------------------

--
-- 表的结构 `field`
--

CREATE TABLE `field` (
                         `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                         `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '字段标识',
                         `type` char(32) NOT NULL COMMENT '字段类型',
                         `name` varchar(64) NOT NULL COMMENT '字段名称',
                         `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '字段状态',
                         `description` varchar(128) DEFAULT NULL COMMENT '字段描述',
                         `created_at` bigint(20) UNSIGNED NOT NULL COMMENT '创建时间',
                         `updated_at` bigint(20) UNSIGNED NOT NULL COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='字段信息';

-- --------------------------------------------------------

--
-- 表的结构 `gorm_init`
--

CREATE TABLE `gorm_init` (
                             `id` int(10) UNSIGNED NOT NULL,
                             `init` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `gorm_init`
--

INSERT INTO `gorm_init` (`id`, `init`) VALUES
    (1, 1);

-- --------------------------------------------------------

--
-- 表的结构 `oauth`
--

CREATE TABLE `oauth` (
                         `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                         `user_id` bigint(20) UNSIGNED DEFAULT NULL,
                         `channel_id` bigint(20) UNSIGNED NOT NULL,
                         `auth_id` varchar(64) DEFAULT NULL COMMENT '渠道授权ID',
                         `union_id` varchar(64) DEFAULT NULL COMMENT '渠道联合ID',
                         `token` varchar(1024) DEFAULT NULL COMMENT '渠道token',
                         `logged_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '登陆时间',
                         `expired_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '过期时间',
                         `created_at` bigint(20) UNSIGNED NOT NULL COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户渠道授权信息';

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
                        `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
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
                        `created_at` bigint(20) UNSIGNED NOT NULL COMMENT '创建时间',
                        `updated_at` bigint(20) UNSIGNED NOT NULL COMMENT '修改时间',
                        `deleted_at` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息';

--
-- 表的结构 `userinfo`
--

CREATE TABLE `userinfo` (
                            `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                            `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户',
                            `field_id` bigint(20) UNSIGNED NOT NULL COMMENT '标识ID',
                            `keyword` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '标识',
                            `value` varchar(1024) NOT NULL COMMENT '值',
                            `created_at` bigint(20) UNSIGNED NOT NULL COMMENT '创建时间',
                            `updated_at` bigint(20) UNSIGNED NOT NULL COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户字段信息';

-- --------------------------------------------------------

--
-- 表的结构 `user_app`
--

CREATE TABLE `user_app` (
                            `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                            `user_id` bigint(20) UNSIGNED NOT NULL,
                            `app_id` bigint(20) UNSIGNED NOT NULL,
                            `status` tinyint(1) DEFAULT '1' COMMENT '状态',
                            `disable_desc` varchar(128) DEFAULT NULL COMMENT '禁用原因',
                            `setting` tinytext COMMENT '用户设置',
                            `token` varchar(512) DEFAULT NULL COMMENT '用户token',
                            `logged_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '登陆时间',
                            `expired_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '过期时间',
                            `created_at` bigint(20) UNSIGNED NOT NULL COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户应用信息';


--
-- 转储表的索引
--

--
-- 表的索引 `app`
--
ALTER TABLE `app`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `keyword` (`keyword`),
  ADD KEY `idx_app_created_at` (`created_at`),
  ADD KEY `idx_app_updated_at` (`updated_at`);

--
-- 表的索引 `app_channel`
--
ALTER TABLE `app_channel`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `app_id` (`app_id`,`channel_id`),
  ADD KEY `channel_id` (`channel_id`);

--
-- 表的索引 `app_field`
--
ALTER TABLE `app_field`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `app_id` (`app_id`,`field_id`),
  ADD KEY `field_id` (`field_id`);

--
-- 表的索引 `auth`
--
ALTER TABLE `auth`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `ua` (`user_id`,`app_id`),
  ADD KEY `created_at` (`created_at`),
  ADD KEY `app_id` (`app_id`);

--
-- 表的索引 `channel`
--
ALTER TABLE `channel`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `keyword` (`keyword`),
  ADD KEY `idx_channel_created_at` (`created_at`),
  ADD KEY `idx_channel_updated_at` (`updated_at`);

--
-- 表的索引 `extra`
--
ALTER TABLE `extra`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uk` (`user_id`,`keyword`),
  ADD KEY `keyword` (`keyword`);

--
-- 表的索引 `feedback`
--
ALTER TABLE `feedback`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `md5` (`md5`),
  ADD KEY `created_at` (`created_at`),
  ADD KEY `updated_at` (`updated_at`),
  ADD KEY `app_id` (`app_id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `category_id` (`category_id`);

--
-- 表的索引 `feedback_category`
--
ALTER TABLE `feedback_category`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- 表的索引 `field`
--
ALTER TABLE `field`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `keyword` (`keyword`),
  ADD KEY `idx_field_updated_at` (`updated_at`),
  ADD KEY `idx_field_created_at` (`created_at`);

--
-- 表的索引 `gorm_init`
--
ALTER TABLE `gorm_init`
    ADD PRIMARY KEY (`id`);

--
-- 表的索引 `oauth`
--
ALTER TABLE `oauth`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`channel_id`),
  ADD UNIQUE KEY `channel_id` (`channel_id`,`auth_id`),
  ADD KEY `created_at` (`created_at`);

--
-- 表的索引 `user`
--
ALTER TABLE `user`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `phone` (`phone`,`deleted_at`),
  ADD UNIQUE KEY `email` (`email`,`deleted_at`),
  ADD UNIQUE KEY `username` (`username`,`deleted_at`),
  ADD KEY `idx_user_created_at` (`created_at`),
  ADD KEY `idx_user_updated_at` (`updated_at`);

--
-- 表的索引 `userinfo`
--
ALTER TABLE `userinfo`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uk` (`user_id`,`keyword`),
  ADD KEY `keyword` (`keyword`),
  ADD KEY `userinfo_ibfk_2` (`field_id`);

--
-- 表的索引 `user_app`
--
ALTER TABLE `user_app`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `ua` (`user_id`,`app_id`),
  ADD KEY `created_at` (`created_at`),
  ADD KEY `fk_user_app_app` (`app_id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `app`
--
ALTER TABLE `app`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `app_channel`
--
ALTER TABLE `app_channel`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=57;

--
-- 使用表AUTO_INCREMENT `app_field`
--
ALTER TABLE `app_field`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=9;

--
-- 使用表AUTO_INCREMENT `auth`
--
ALTER TABLE `auth`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=783;

--
-- 使用表AUTO_INCREMENT `channel`
--
ALTER TABLE `channel`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `extra`
--
ALTER TABLE `extra`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID';

--
-- 使用表AUTO_INCREMENT `feedback`
--
ALTER TABLE `feedback`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `feedback_category`
--
ALTER TABLE `feedback_category`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `field`
--
ALTER TABLE `field`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `gorm_init`
--
ALTER TABLE `gorm_init`
    MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `oauth`
--
ALTER TABLE `oauth`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=189;

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=1338;

--
-- 使用表AUTO_INCREMENT `userinfo`
--
ALTER TABLE `userinfo`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=11;

--
-- 使用表AUTO_INCREMENT `user_app`
--
ALTER TABLE `user_app`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=142;

--
-- 限制导出的表
--

--
-- 限制表 `app_channel`
--
ALTER TABLE `app_channel`
    ADD CONSTRAINT `app_channel_ibfk_1` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `app_channel_ibfk_2` FOREIGN KEY (`channel_id`) REFERENCES `channel` (`id`) ON DELETE CASCADE;

--
-- 限制表 `app_field`
--
ALTER TABLE `app_field`
    ADD CONSTRAINT `app_field_ibfk_1` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `app_field_ibfk_2` FOREIGN KEY (`field_id`) REFERENCES `field` (`id`) ON DELETE CASCADE;

--
-- 限制表 `auth`
--
ALTER TABLE `auth`
    ADD CONSTRAINT `auth_ibfk_1` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`),
  ADD CONSTRAINT `auth_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;

--
-- 限制表 `extra`
--
ALTER TABLE `extra`
    ADD CONSTRAINT `user_field_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `user_field_ibfk_2` FOREIGN KEY (`keyword`) REFERENCES `field` (`keyword`) ON UPDATE CASCADE;

--
-- 限制表 `feedback`
--
ALTER TABLE `feedback`
    ADD CONSTRAINT `feedback_ibfk_1` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `feedback_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `feedback_ibfk_3` FOREIGN KEY (`category_id`) REFERENCES `feedback_category` (`id`);

--
-- 限制表 `oauth`
--
ALTER TABLE `oauth`
    ADD CONSTRAINT `oauth_ibfk_1` FOREIGN KEY (`channel_id`) REFERENCES `channel` (`id`),
  ADD CONSTRAINT `oauth_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;

--
-- 限制表 `userinfo`
--
ALTER TABLE `userinfo`
    ADD CONSTRAINT `userinfo_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `userinfo_ibfk_2` FOREIGN KEY (`field_id`) REFERENCES `field` (`id`),
  ADD CONSTRAINT `userinfo_ibfk_3` FOREIGN KEY (`keyword`) REFERENCES `field` (`keyword`) ON UPDATE CASCADE;

--
-- 限制表 `user_app`
--
ALTER TABLE `user_app`
    ADD CONSTRAINT `fk_user_app_app` FOREIGN KEY (`app_id`) REFERENCES `app` (`id`),
  ADD CONSTRAINT `fk_user_user_apps` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
