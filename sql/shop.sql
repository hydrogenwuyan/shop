DROP DATABASE IF EXISTS `shop`;

set names utf8;
CREATE DATABASE IF NOT EXISTS `booksys` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

USE shop;

-- --------------------------------------------------
--  Table Structure for `t_user_entity`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `t_user_entity` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id` bigint NOT NULL COMMENT '用户id',
    `user_name` varchar(50) NOT NULL COMMENT '用户名',
    `password` varchar(300) NOT NULL COMMENT '密码',
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY(`id`)
) ENGINE=InnoDB COMMENT='user entity' DEFAULT CHARSET=utf8;
CREATE UNIQUE INDEX `t_user_entity_user_id` ON `t_user_entity` (`user_id`);
