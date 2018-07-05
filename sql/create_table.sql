CREATE TABLE `t_device` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "自增主键",
	`device_id` BIGINT UNSIGNED NOT NULL COMMENT "设备id",
	`token` CHAR (40) NOT NULL COMMENT "设备登录的token",
	`device_type` UNSIGNED thinint NOT NULL COMMENT "设备类型,0:Android；1:IOS；2：Windows;3:Web",
	`model` thinint NOT NULL COMMENT "机型",
	`version` CHAR (10) NOT NULL COMMENT "设备版本",
	`user_id` BIGINT UNSIGNED NOT NULL COMMENT "账户id",
	`online_state` thinint UNSIGNED NOT NULL "在线状态，0：不在线；1：在线",
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "更新时间",
	PRIMARY KEY `pk_id` (`id`),
	UNIQUE INDEX `uk_device_id` (`device_id`),
	INDEX `idx_user_id` (`user_id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT "设备";

CREATE TABLE `t_user` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "自增主键",
	`user_id` BIGINT UNSIGNED NOT NULL COMMENT "账户id",
	`number` CHAR (20) NOT NULL COMMENT "手机号",
	`nickname` CHAR (20) NOT NULL COMMENT "昵称",
	`pwd` CHAR (20) NOT NULL COMMENT "密码",
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "更新时间",
	PRIMARY KEY `pk_id` (`id`),
	UNIQUE INDEX `uk_user_id` (`user_id`),
	UNIQUE INDEX `uk_user_name` (`user_name`)
) ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT "账户";

CREATE TABLE `t_friend` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "自增主键",
	`user_id` BIGINT UNSIGNED NOT NULL COMMENT "账户id",
	`friend` BIGINT UNSIGNED NOT NULL COMMENT "好友账户id",
	`label` CHAR (20) NOT NULL COMMENT "备注，标签",
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "更新时间",
	PRIMARY KEY `pk_id` (`id`),
	UNIQUE INDEX `uk_user_id` (`user_id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT "好友关系";

CREATE TABLE `t_group` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "自增主键",
	`group_id` BIGINT UNSIGNED NOT NULL COMMENT "组id",
	`group_name` BIGINT UNSIGNED NOT NULL COMMENT "组名",
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "更新时间",
	PRIMARY KEY `pk_id` (`id`),
	UNIQUE INDEX `uk_grop_id` (`grop_id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT "群组";

CREATE TABLE `t_group_user` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "自增主键",
	`group_id` BIGINT UNSIGNED NOT NULL COMMENT "组id",
	`user_id` BIGINT UNSIGNED NOT NULL COMMENT "组名",
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "更新时间",
	PRIMARY KEY `pk_id` (`id`),
	UNIQUE INDEX `uk_grop_id` (`grop_id`),
	INDEX `idx_user_id` (`user_id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT "群组成员";

CREATE TABLE `t_message` (
	 `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "自增主键",
	`sender` BIGINT UNSIGNED NOT NULL COMMENT "发送者账户id",
	`receiver_type` BIGINT UNSIGNED NOT NULL COMMENT "接收者账户id",
	`receiver` BIGINT UNSIGNED NOT NULL COMMENT "接收者账户id",
	`group_id` BIGINT UNSIGNED NULL COMMENT "群组id,如果为null,为单聊信息，如果不为null，则为群聊信息",
	`type` thinint UNSIGNED NOT NULL COMMENT "消息类型,0：文本；1：语音；2：图片",
	`content` BLOB NOT NULL COMMENT "内容",
	`sync_key` BIGINT UNSIGNED COMMENT "消息同步序列",
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间",
	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "更新时间",
	PRIMARY KEY `pk_id` (`id`),
	INDEX `idx_sync_key` (`sync_key`)
) ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT "消息";