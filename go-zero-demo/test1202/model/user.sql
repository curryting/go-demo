CREATE TABLE `tbUserConfig` (
     `iId` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户自增id',
     `sUserName` varchar(64) NOT NULL COMMENT '用户名',
     `sPassword` varchar(32) NOT NULL COMMENT '密码',
     `dtCreateTime` datetime NOT NULL COMMENT '创建时间',
     `dtUpdateTime` datetime NOT NULL COMMENT '修改时间',
     PRIMARY KEY (`iId`),
     UNIQUE KEY `sUserName` (`sUserName`)
) ENGINE=InnoDB CHARSET=gbk COMMENT='用户配置表';