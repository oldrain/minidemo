
CREATE TABLE `user` (
  `id` bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;


CREATE TABLE `customer_token` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `partner_id` varchar(128) NOT NULL DEFAULT '0' COMMENT '合作商ID',
  `app_id` varchar(128) NOT NULL DEFAULT '0' COMMENT '应用ID',
  `customer_code` varchar(40) NOT NULL DEFAULT '0' COMMENT '客户ID',
  `token` varchar(1024) NOT NULL DEFAULT '' COMMENT 'token',
  `client_name` varchar(40) NOT NULL DEFAULT '' COMMENT '客户端名称(mini-program)',
  `login_time` datetime NOT NULL COMMENT '登录时间',
  `deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除(0=否; 1=是)',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `created_by` varchar(20) NOT NULL DEFAULT '' COMMENT '创建人',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_by` varchar(20) NOT NULL DEFAULT '' COMMENT '更新人',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_customer_code` (`customer_code`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='帐号token记录表';
