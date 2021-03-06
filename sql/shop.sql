DROP DATABASE IF EXISTS `shop`;

set names utf8;
CREATE DATABASE IF NOT EXISTS `shop` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

USE shop;

-- --------------------------------------------------
--  Table Structure for `t_user_entity`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `t_user_entity` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `userId` bigint NOT NULL COMMENT '用户id',
    `password` varchar(300) NOT NULL COMMENT '密码',
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY(`id`)
) ENGINE=InnoDB COMMENT='user entity' DEFAULT CHARSET=utf8;
CREATE UNIQUE INDEX `t_user_entity_user_id` ON `t_user_entity` (`userId`);

-- --------------------------------------------------
--  Table Structure for `t_inventory_entity`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `t_inventory_entity` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `shopId` bigint NOT NULL COMMENT '商品id',
    `num` bigint NOT NULL COMMENT '商品数量',
    `money` bigint NOT NULL COMMENT '商品单价 单位:分',
    `version` bigint NOT NULL COMMENT '版本号',
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY(`id`)
) ENGINE=InnoDB COMMENT='inventory entity' DEFAULT CHARSET=utf8;
CREATE UNIQUE INDEX `t_inventory_entity_shop_id` ON `t_inventory_entity` (`shopId`);

insert t_inventory_entity values(1,1,100,10,1,0,0,0);
insert t_inventory_entity values(2,2,100,10,1,0,0,0);

-- --------------------------------------------------
--  Table Structure for `t_order_entity`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `t_order_entity` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
    `status` int(11) NOT NULL COMMENT '订单状态',
    `shopId` bigint NOT NULL COMMENT '商品id',
    `num` bigint NOT NULL COMMENT '商品数量',
    `userId` bigint NOT NULL COMMENT '用户id',
    `money` bigint NOT NULL COMMENT '订单金额',
    `updateTime` bigint(20) DEFAULT 0 COMMENT "更新时间",
    `createTime` bigint(20) DEFAULT 0 COMMENT "创建时间",
    `deleteTime` bigint(20)  DEFAULT 0 COMMENT "删除时间",
    PRIMARY KEY(`id`)
) ENGINE=InnoDB COMMENT='order entity' DEFAULT CHARSET=utf8;