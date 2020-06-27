############################
# 全国计算机等级考试
# 二级教程 - MySQL数据库程序设计
############################

# 使用数据库`db_school`
USE `db_school`;

############################
# 第六章 索引
############################

# 例 6.1 - 在创建新表的同时创建普通索引。
# 为了测试使用CREATE TABLE时建立索引的方法，建立一个表`tb_student1`包括7个字段，用于测试索引的创建、查看等。
# 要求在创建表的同时，在`studentName`字段上建立普通索引。
CREATE TABLE `tb_student1`
(
`studentNo` CHAR(10) NOT NULL,
`studentName` VARCHAR(20) NOT NULL,
`sex` CHAR(2) NOT NULL,
`birthday` DATE,
`native` VARCHAR(20),
`nation` VARCHAR(10) DEFAULT '汉',
`classNo` CHAR(6),
INDEX `idx_studName`(`studentName`)
);

# 例 6.2 - 创建新表时，建立唯一性索引。
# 建立一个表`tb_student2`，要求在建立表`tb_student2`的同时，在`studentNo`字段上建立唯一性索引。
CREATE TABLE `tb_student2`
(
`studentNo` CHAR(10) NOT NULL,
`studentName` VARCHAR(20) NOT NULL,
`sex` CHAR(2) NOT NULL,
`birthday` DATE,
`native` VARCHAR(20),
`nation` VARCHAR(10) DEFAULT '汉',
`classNo` CHAR(6),
CONSTRAINT `uq_studentNo` UNIQUE (`studentNo`)
);

# 例 6.3 - 在创建新表的同时建立主键索引。
# 在MySQL中，创建表时，若指定表的主键，系统自动建立主键索引。若建立外键，系统亦自动建立索引。
# 为了便于查看创建索引的效果，我们分别建立两个表`tb_score1`及`tb_score2`。两个表的结构完全相同，区别仅在于创建`tb_score2`语句中定义了表的主键和外键。
# 以创建`tb_score1`表为例，查看该表上所建立的索引。
CREATE TABLE `tb_score1`
(
`studentNo` CHAR(10),
`courseNo` CHAR(5),
`score` FLOAT
);

SHOW INDEXES FROM tb_score1\G

CREATE TABLE `tb_score2`
(
`studentNo` CHAR(10),
`courseNo` CHAR(5),
`score` FLOAT,
CONSTRAINT `PK_score2` PRIMARY KEY(`studentNo`, `courseNo`),
CONSTRAINT `FK_score21` FOREIGN KEY(`studentNo`) REFERENCES `tb_student`(`studentNo`),
CONSTRAINT `FK_score22` FOREIGN KEY(`courseNo`) REFERENCES `tb_course`(`courseNo`)
);

SHOW INDEXES FROM tb_score2\G

# 例 6.4 - 创建普通索引。
# 在数据库`db_school`的学生表`tb_student`上建立一个普通索引，索引字段是学号`studentNo`。
CREATE INDEX `index_stu`
ON `tb_student`(`studentNo`);

# 例 6.5 - 创建基于字段值前缀字符的索引。
# 在数据库`db_school`中课程表`tb_course`上建立一个索引，要求按课程名称`courseName`字段值前三个字符建立降序索引。
CREATE INDEX `index_course`
ON `tb_course`(`courseName`(3) DESC);

# 例 6.6 - 创建组合索引。
# 在数据库`db_school`中表`tb_book`上建立图书类别（升序）和书名（降序）的组合序列，索引名称为`index_book`。
-- CREATE INDEX `index_book`
-- ON `tb_book`(`bclassNo`, `bookName` DESC);

# 例 6.7 - 使用ALTER TABLE语句建立普通索引。
# 使用ALTER TABLE语句在`例6.1`所建立的`tb_student1`表`studentName`列上建立一个普通索引。
ALTER TABLE `tb_student1` ADD INDEX `idx_studentName`(`studentName`);

# 例 6.8 - 删除例`例6.7`中所创建的索引`idx_studentName`。
DROP INDEX `idx_studentName` ON `tb_student1`;

# 例 6.9 - 使用ALTER TABLE语句删除索引。
ALTER TABLE `tb_student1` DROP INDEX `idx_studName`;

# 练习题1 - 分别用ALTER TABLE及CREATE INDEX语句在表`tb_student`上建立主键索引。
ALTER TABLE `tb_student` DROP PRIMARY KEY;
ALTER TABLE `tb_student` ADD PRIMARY KEY `pk_studentNo`(`studentNo`);
-- DROP INDEX `pk_studentNo` ON `tb_student`;

# 练习题2 - 在创建表`tb_score`的同时，建立学号、课程号的组合索引。
DROP TABLE IF EXISTS `tb_score`;

CREATE TABLE IF NOT EXISTS `tb_score`
(
`studentNo`     CHAR(10),
`courseNo`      CHAR(6),
`score`         FLOAT       CHECK(score>=0 AND score<=100),
INDEX `idx_score`(`studentNo`, `courseNo`)
) ENGINE=InnoDB;

# 练习题3 - 删除上题创建在`tb_score`表上的索引。
ALTER TABLE `tb_score` DROP INDEX IF EXISTS `idx_score`;

DROP INDEX IF EXISTS `idx_score` ON `tb_score`;
