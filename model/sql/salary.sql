CREATE TABLE `salaries` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
    `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
    `employee_id` bigint(20) unsigned NOT NULL COMMENT '员工 ID,关联 employees 表的 ID',
    `enterprise_code` varchar(255) NOT NULL COMMENT '企业识别码',
    `month` varchar(255) NOT NULL COMMENT '月份,例如 2024-04',
    `salary` int(11) NOT NULL COMMENT '实发工资',
    PRIMARY KEY (`id`),
    KEY `idx_salaries_deleted_at` (`deleted_at`),
    KEY `idx_salaries_employee_id_enterprise_code` (`employee_id`, `enterprise_code`),
    CONSTRAINT `fk_salaries_employee_id` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='员工工资表';