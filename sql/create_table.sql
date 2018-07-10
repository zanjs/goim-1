CREATE TABLE `t_device` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "自增主键",
	`device_id` BIGINT UNSIGNED NOT NULL COMMENT "设备id",
	`token` CHAR ( 40 ) NOT NULL COMMENT "设备登录的token",
	`device_type` TINYINT UNSIGNED NOT NULL COMMENT "设备类型,0:Android；1:IOS；2：Windows;3:Web",
	`model` TINYINT NOT NULL COMMENT "机型",
	`version` CHAR ( 10 ) NOT NULL COMMENT "设备版本",
	`user_id` BIGINT UNSIGNED NOT NULL COMMENT "账户id",
	`state` TINYINT UNSIGNED NOT NULL COMMENT "在线状态，0：不在线；1：在线",
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "更新时间",
	PRIMARY KEY `pk_id` ( `id` ),
	UNIQUE INDEX `uk_device_id` ( `device_id` ),
	INDEX `idx_user_id` ( `user_id` )
) COMMENT "设备";
CREATE TABLE `t_user` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "自增主键",
	`user_id` BIGINT UNSIGNED NOT NULL COMMENT "账户id",
	`number` VARCHAR ( 20 ) NOT NULL COMMENT "手机号",
	`nickname` VARCHAR ( 20 ) NOT NULL COMMENT "昵称",
	`password` VARCHAR ( 20 ) NOT NULL COMMENT "密码",
	`sex` TINYINT NOT NULL COMMENT "性别，1:男；2:女",
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "更新时间",
	PRIMARY KEY `pk_id` ( `id` ),
	UNIQUE INDEX `uk_user_id` ( `user_id` ),
	UNIQUE INDEX `uk_number` ( `number` )
) COMMENT "账户";
CREATE TABLE `t_friend` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "自增主键",
	`user_id` BIGINT UNSIGNED NOT NULL COMMENT "账户id",
	`friend` BIGINT UNSIGNED NOT NULL COMMENT "好友账户id",
	`label` CHAR ( 20 ) NOT NULL COMMENT "备注，标签",
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "更新时间",
	PRIMARY KEY `pk_id` ( `id` ),
	UNIQUE INDEX `uk_user_id` ( `user_id` )
) COMMENT "好友关系";
CREATE TABLE `t_group` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "自增主键",
	`group_id` BIGINT UNSIGNED NOT NULL COMMENT "组id",
	`group_name` BIGINT UNSIGNED NOT NULL COMMENT "组名",
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "更新时间",
	PRIMARY KEY `pk_id` ( `id` ),
	UNIQUE INDEX `uk_group_id` ( `group_id` )
) COMMENT "群组";
CREATE TABLE `t_group_user` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "自增主键",
	`group_id` BIGINT UNSIGNED NOT NULL COMMENT "组id",
	`user_id` BIGINT UNSIGNED NOT NULL COMMENT "用户id",
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "更新时间",
	PRIMARY KEY `pk_id` ( `id` ),
	UNIQUE INDEX `uk_group_id` ( `group_id` ),
	INDEX `idx_user_id` ( `user_id` )
) COMMENT "群组成员关系";
CREATE TABLE `t_message` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "自增主键",
	`user_id` BIGINT UNSIGNED NOT NULL COMMENT "用户id",
	`sender` BIGINT UNSIGNED NOT NULL COMMENT "发送者账户id",
	`receiver_type` BIGINT UNSIGNED NOT NULL COMMENT "接收者账户id",
	`receiver` BIGINT UNSIGNED NOT NULL COMMENT "接收者id,如果是单聊信息，则为user_id，如果是群组消息，则为group_id",
	`type` TINYINT UNSIGNED NOT NULL COMMENT "消息类型,0：文本；1：语音；2：图片",
	`content` BLOB NOT NULL COMMENT "内容",
	`sync_key` BIGINT UNSIGNED COMMENT "消息同步序列",
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "更新时间",
	PRIMARY KEY `pk_id` ( `id` ),
INDEX `idx_sync_key` ( `sync_key` )
) COMMENT "消息";