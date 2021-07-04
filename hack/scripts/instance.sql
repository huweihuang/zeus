CREATE TABLE IF NOT EXISTS `t_instance` (
  `id`                 INT(11)      NOT NULL AUTO_INCREMENT,
  `job_id`             VARCHAR(64)  NOT NULL DEFAULT ''     COMMENT '任务ID'
  `name`               VARCHAR(64)  NOT NULL DEFAULT ''     COMMENT '实例名字',
  `namespace`          VARCHAR(64)  NOT NULL DEFAULT ''     COMMENT '命名空间',
  `host_id`            VARCHAR(64)  NOT NULL DEFAULT ''     COMMENT '母机ID',
  `replicas`           INT(11)      NOT NULL DEFAULT 0      COMMENT '副本数',
  `image`              VARCHAR(128) NOT NULL DEFAULT ''     COMMENT '镜像',
  `status`             BOOLEAN      NOT NULL DEFAULT false  COMMENT '实例状态',
  `job_state`          VARCHAR(64)  NOT NULL DEFAULT ''     COMMENT '任务状态',
  `created_time`       TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modify_time`        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uind_name` (`name`,`host_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
