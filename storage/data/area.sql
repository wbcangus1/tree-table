CREATE TABLE `area`
(
    `id`          int         NOT NULL AUTO_INCREMENT COMMENT '区域ID',
    `parent_id`   int                  DEFAULT NULL COMMENT '父区域ID',
    `name`        varchar(100) NOT NULL COMMENT '名称',
    `level`       tinyint     NOT NULL COMMENT '层级',
    `sort`        int         NOT NULL COMMENT '同层级排序',
    `path`        varchar(255)         DEFAULT NULL COMMENT '节点路径',
    `description` text COMMENT '描述',
    `tags`        text COMMENT '标签Json数组',
    `created_at`  datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_area_name` (`name`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 100001
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;