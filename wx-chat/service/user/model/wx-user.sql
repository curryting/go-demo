CREATE TABLE `wx_user` (
                        `id` bigint NOT NULL AUTO_INCREMENT,
                        `username` varchar(255)  NOT NULL DEFAULT '' COMMENT 'username',
                        `password` varchar(255)  NOT NULL DEFAULT '' COMMENT 'password',
                        `gender` tinyint(1) DEFAULT '0' NOT NULL COMMENT '0-un status｜1-male｜2-female',
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `username_unique` (`username`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;