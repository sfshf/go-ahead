############################
# 全国计算机等级考试
# 二级教程 - MySQL数据库程序设计
############################

############################
# 第三章 练习题
############################

# 创建数据库`db_sp`
CREATE DATABASE IF NOT EXISTS `db_sp`
DEFAULT CHARACTER SET `utf8`
DEFAULT COLLATE `utf8_general_ci`;

USE `db_sp`;

# 创建供应商表`S`
CREATE TABLE IF NOT EXISTS `S`
(
`SNO` VARCHAR(25) COMMENT '供应商编号',
`SNAME` VARCHAR(25) NOT NULL COMMENT '供应商名称',
`STATUS` INT(10) COMMENT '状态',
`CITY` VARCHAR(25) COMMENT '所在城市',
CONSTRAINT `PK_S` PRIMARY KEY (`SNO`),
CONSTRAINT `UQ_S` UNIQUE (`SNAME`),
CONSTRAINT `CK_S` CHECK (`CITY`!='London' OR `STATUS`=20)
) ENGINE=InnoDB;

# 创建零件表`P`
CREATE TABLE IF NOT EXISTS `P`
(
`PNO` VARCHAR(25) COMMENT '零件编号',
`PNAME` VARCHAR(25) COMMENT '零件名称',
`COLOR` VARCHAR(25) COMMENT '颜色',
`WEIGHT` DOUBLE COMMENT '重量',
CONSTRAINT `PK_P` PRIMARY KEY (`PNO`),
CONSTRAINT `CK_P` CHECK (`COLOR` IN ('Red', 'Yellow', 'Green', 'Blue'))
) ENGINE=InnoDB;

# 创建供应情况表`SP`
CREATE TABLE IF NOT EXISTS `SP`
(
`SNO` VARCHAR(25) COMMENT '供应商编号',
`PNO` VARCHAR(25) COMMENT '零件编号',
`QTY` INT(25) COMMENT '供应量',
CONSTRAINT `PK_SP` PRIMARY KEY (`SNO`, `PNO`)
) ENGINE=InnoDB;

# 添加`外键`
ALTER TABLE `SP` ADD CONSTRAINT `FK_SP_S` FOREIGN KEY (`SNO`) REFERENCES `S`(`SNO`);
ALTER TABLE `SP` ADD CONSTRAINT `FK_SP_P` FOREIGN KEY (`PNO`) REFERENCES `P`(`PNO`);

# 插入数据
INSERT INTO `S`(`SNO`, `SNAME`, `STATUS`, `CITY`) VALUES
('S1', 'Smith', 20, 'London'),
('S2', 'Jones', 10, 'Paris'),
('S3', 'Blake', 30, 'Paris'),
('S4', 'Clark', 20, 'London'),
('S5', 'Adams', 30, 'Athens'),
('S6', 'Brown', NULL, 'New York');

INSERT INTO `P`(`PNO`, `PNAME`, `COLOR`, `WEIGHT`) VALUES
('P1', 'Nut', 'Red', 12),
('P2', 'Bolt', 'Green', 17),
('P3', 'Screw', 'Blue', 17),
('P4', 'Screw', 'Red', 14),
('P5', 'Cam', 'Blue', 12),
('P6', 'Cog', 'Red', 19);

INSERT INTO `SP`(`SNO`, `PNO`, `QTY`) VALUES
('S1', 'P1', 200),
('S1', 'P4', 700),
('S1', 'P5', 400),
('S2', 'P1', 200),
('S2', 'P2', 200),
('S2', 'P3', 500),
('S2', 'P4', 600),
('S2', 'P5', 400),
('S2', 'P6', 800),
('S3', 'P3', 200),
('S3', 'P4', 500),
('S4', 'P2', 300),
('S4', 'P5', 300),
('S5', 'P1', 100),
('S5', 'P6', 200),
('S5', 'P2', 100),
('S5', 'P3', 200),
('S5', 'P5', 400);
