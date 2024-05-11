CREATE TABLE `employees` (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
     `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
     `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
     `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
     `enterprise_code` varchar(255) NOT NULL COMMENT '企业识别码',
     `name` varchar(255) NOT NULL COMMENT '员工姓名',
     `gender` bigint(20) NOT NULL COMMENT '员工性别',
     `age` bigint(20) NOT NULL COMMENT '员工年龄',
     `avatar` varchar(255) NOT NULL COMMENT '员工照片',
     `mobile` varchar(255) NOT NULL COMMENT '手机号',
     `introduce` text NOT NULL COMMENT '备注信息',
     `id_card_num` varchar(255) NOT NULL COMMENT '身份证号',
     `bank_card_num` varchar(255) NOT NULL COMMENT '银行卡号',
     PRIMARY KEY (`id`),
     KEY `idx_employees_deleted_at` (`deleted_at`),
     KEY `idx_employees_enterprise_code` (`enterprise_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='员工信息表';