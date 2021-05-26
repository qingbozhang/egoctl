DROP TABLE IF EXISTS `enum`;
CREATE TABLE `enum` (
                        `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
                        `group_key` varchar(255) NOT NULL COMMENT '组名key，唯一标识符',
                        `group_title` varchar(255) NOT NULL COMMENT '名称',
                        `key` int(11) NOT NULL COMMENT 'ant键',
                        `title` varchar(255) NOT NULL COMMENT 'ant名称',
                        `ctime` bigint(20) NOT NULL COMMENT '创建时间',
                        `utime` bigint(20) NOT NULL COMMENT '更新时间',
                        `dtime` bigint(20) NOT NULL COMMENT '删除时间',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;
DROP TABLE IF EXISTS `send`;
CREATE TABLE `send` (
                        `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
                        `strategy_id` int(11) NOT NULL COMMENT '策略id',
                        `uid` int(11) NOT NULL COMMENT '用户',
                        `channel_content_id` int(11) DEFAULT NULL COMMENT '通道内容',
                        `status` int(11) NOT NULL COMMENT '状态',
                        `send_status` int(11) NOT NULL COMMENT '发送状态 1 初始化，2 发送中，2超时，3成功，4失败',
                        `reply` longtext NOT NULL COMMENT '发送内容',
                        `op_uid` int(11) NOT NULL COMMENT '操作uid',
                        `ctime` bigint(20) NOT NULL COMMENT '创建时间',
                        `channel_type_id` int(11) NOT NULL COMMENT '通道类型',
                        `plan_send_time` bigint(20) NOT NULL COMMENT '计划发送时间',
                        `exec_send_time` bigint(20) NOT NULL COMMENT '执行发送时间',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4;