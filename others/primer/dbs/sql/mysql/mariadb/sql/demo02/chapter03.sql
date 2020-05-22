############################
# 全国计算机等级考试
# 二级教程 - MySQL数据库程序设计
############################

############################
# 第三章 数据定义
############################

# 创建`db_school`数据库
CREATE DATABASE IF NOT EXISTS `db_school`
DEFAULT CHARACTER SET `GB2312`
DEFAULT COLLATE `GB2312_chinese_ci`;

# 使用`db_school`数据库
USE db_school;

# 创建`tb_student`表
CREATE TABLE IF NOT EXISTS `tb_student`
(
`studentNo`   CHAR(10),
`studentName` VARCHAR(20)   NOT NULL,
`sex`         CHAR(2)       NOT NULL,
`birthday`    DATE,
`native`      VARCHAR(20),
`nation`      VARCHAR(10)   DEFAULT '汉',
`classNo`     VARCHAR(10),
CONSTRAINT `PK_student` PRIMARY KEY (studentNo)
) ENGINE=InnoDB;

# 创建`tb_class`表
CREATE TABLE IF NOT EXISTS `tb_class`
(
`classNo`       CHAR(6) PRIMARY KEY,
`className`     VARCHAR(20) NOT NULL,
`department`    VARCHAR(30) NOT NULL,
`grade`         SMALLINT,
`classNum`      TINYINT,
CONSTRAINT `UQ_class` UNIQUE (`className`)
) ENGINE=InnoDB;

# 创建`tb_course`表
CREATE TABLE IF NOT EXISTS `tb_course`
(
`courseNo`      CHAR(6),
`courseName`    VARCHAR(20)     NOT NULL,
`credit`        INT             NOT NULL,
`courseHour`    INT             NOT NULL,
`term`          CHAR(2),
`priorCourse`   CHAR(6),
CONSTRAINT `PK_course` PRIMARY KEY (`courseNo`),
CONSTRAINT `CK_course` CHECK (`credit`=`courseHour`/16)
) ENGINE=InnoDB;

# 创建`tb_score`表
CREATE TABLE IF NOT EXISTS `tb_score`
(
`studentNo`     CHAR(10),
`courseNo`      CHAR(6),
`score`         FLOAT       CHECK(score>=0 AND score<=100),
CONSTRAINT `PK_score` PRIMARY KEY (studentNo, courseNo)
) ENGINE=InnoDB;

# 添加`外键`
ALTER TABLE `tb_student` ADD CONSTRAINT `FK_student` FOREIGN KEY (`classNo`) REFERENCES `tb_class`(`classNo`);
ALTER TABLE `tb_course` ADD CONSTRAINT `FK_course` FOREIGN KEY (`priorCourse`) REFERENCES `tb_course`(`courseNo`);
ALTER TABLE `tb_score` ADD CONSTRAINT `FK_score1` FOREIGN KEY (`studentNo`) REFERENCES `tb_student`(`studentNo`);
ALTER TABLE `tb_score` ADD CONSTRAINT `FK_score2` FOREIGN KEY (`courseNo`) REFERENCES `tb_course`(`courseNo`);

# 插入所有数据
INSERT INTO `db_school`.`tb_class`(`classNo`, `className`, `department`, `grade`, `classNum`) VALUES
('AC1301', '会计 13-1 班', '会计学院', 2013, 35),
('AC1302', '会计 13-2 班', '会计学院', 2013, 35),
('CS1401', '计算机 14-1 班', '计算机学院', 2014, 35),
('IS1301', '信息系统 13-1 班', '信息学院', 2013, NULL),
('IS1401', '信息系统 14-1 班', '信息学院', NULL, 30);

INSERT INTO `db_school`.`tb_student`(`studentNO`, `studentName`, `sex`, `birthday`, `native`, `nation`, `classNo`) VALUES
('2013110101', '张晓勇', '男', '1997-12-11', '山西', '汉', 'AC1301'),
('2013110103', '王一敏', '女', '1996-03-25', '河北', '汉', 'AC1301'),
('2013110201', '江山', '女', '1996-09-17', '内蒙古', '锡伯', 'AC1302'),
('2013110202', '李明', '男', '1996-01-14', '广西', '壮', 'AC1302'),
('2013310101', '黄菊', '女', '1995-09-30', '北京', '汉', 'IS1301'),
('2013310102', '林海', '男', '1991-01-14', '南宁', '汉', 'IS1301'),
('2013310103', '吴昊', '男', '1995-11-18', '河北', '汉', 'IS1301'),
('2014210101', '刘涛', '男', '1997-04-03', '湖南', '侗', 'CS1401'),
('2014210102', '郭志坚', '男', '1997-02-21', '上海', '汉', 'CS1401'),
('2014310101', '王林', '男', '1996-10-09', '河南', '汉', 'IS1401'),
('2014310102', '李怡然', '女', '1996-12-31', '辽宁', '汉', 'IS1401');

INSERT INTO `db_school`.`tb_course`(`courseNo`, `courseName`, `credit`, `courseHour`, `term`, `priorCourse`) VALUES
('11003', '管理学', 2, 32, '2', NULL),
('11005', '会计学', 3, 48, '3', NULL),
('21001', '计算机基础', 3, 48, '1', NULL),
('21002', 'OFFICE高级应用', 3, 48, '2', '21001'),
('21004', '程序设计', 4, 64, '2', '21001'),
('21005', '数据库', 4, 64, '4', '21004'),
('21006', '操作系统', 4, 64, '5', '21001'),
('31001', '管理信息系统', 3, 48, '3', '21004'),
('31002', '信息系统_分析与设计', 2, 32, '4', '31001'),
('31005', '项目管理', 3, 48, '5', '31001');

INSERT INTO `db_school`.`tb_score`(`studentNo`, `courseNo`, `score`) VALUES
('2013110101', '11003', 90),
('2013110101', '21001', 86),
('2013110103', '11003', 89),
('2013110103', '21001', 88),
('2013110201', '11003', 78),
('2013110201', '21001', 92),
('2013110202', '11003', 82),
('2013110202', '21001', 85),
('2013310101', '21004', 83),
('2013310101', '31002', 68),
('2013310103', '21004', 80),
('2013310103', '31002', 76),
('2014210101', '21002', 93),
('2014210101', '21004', 89),
('2014210102', '21002', 95),
('2014210102', '21004', 88),
('2014310101', '21001', 79),
('2014310101', '21004', 80),
('2014310102', '21001', 91),
('2014310102', '21004', 87);
