############################
# 全国计算机等级考试
# 二级教程 - MySQL数据库程序设计
############################

# 使用`db_school`数据库
USE db_school;

############################
# 第五章 数据更新
############################

# 例 5.1 - 向表`tb_student`中插入一条新记录('2014210103', '王玲', '女', '1998-02-21', '安徽', '汉', 'CS1401')。
INSERT INTO tb_student VALUES ('2014210103', '王玲', '女', '1998-02-21', '安徽', '汉', 'CS1401');

# 例 5.2 - 向表`tb_student`中插入一条新记录('2013110102', '赵婷婷', '女', '1996-11-30', '天津', '汉', 'AC1301')。
INSERT INTO tb_student VALUES ('2013110102', '赵婷婷', '女', '1996-11-30', '天津', '汉', 'AC1301');

# 例 5.3 - 向表`tb_student`中插入一条新记录，学号为'2013110203'，姓名为'孟颖'，性别为'女'，出生日期为'1997-03-20'，籍贯为'上海'，民族为'汉'，班号为'AC1302'。
INSERT INTO tb_student(studentNo, studentName, sex, birthday, native, nation, classNo) VALUES
('2013110203', '孟颖', '女', '1997-03-20', '上海', '汉', 'AC1302');

# 例 5.4 - 向数据库`db_school`的表`tb_student`中插入一条新记录，学号为'2014310103'，姓名为'孙新'，性别为'男'，民族为'傣'，班号为'IS1401'。
INSERT INTO db_school.tb_student(studentNo, studentName, sex, nation, classNo) VALUES
('2014310103', '孙新', '男', '傣', 'IS1401');

# 例 5.5 - 在表`tb_student`中插入三条新记录：学号为'2014310104'，姓名为'陈卓卓'，性别为'女'；学号为'2014310105'，姓名为'马丽'，性别为'女'；学号为'2014310106'，姓名为'许江'，性别为'男'。
INSERT INTO tb_student(studentNo, studentName, sex) VALUES
('2014310104', '陈卓卓', '女'),
('2014310105', '马丽', '女'),
('2014310106', '许江', '男');

# 例 5.6 - 不指定插入字段列表，向数据库`db_school`的表`tb_student`中插入两条记录('2014310107', '赵鹏', '男', '1997-10-16', '吉林', '朝鲜', 'IS1401')，('2014310108', '李菊', '女', '1998-01-24', '河北', '汉', 'IS1401')。
INSERT INTO db_school.tb_student VALUES
('2014310107', '赵鹏', '男', '1997-10-16', '吉林', '朝鲜', 'IS1401'),
('2014310108', '李菊', '女', '1998-01-24', '河北', '汉', 'IS1401');

# 例 5.7 - 假设要为表`tb_student`制作一个备份表`tb_student_copy`，两个表结构完全一致，现使用`INSERT...SELECT`语句将表`tb_student`中的数据备份到表`tb_student_copy`中。
CREATE TABLE `tb_student_copy` (
  `studentNo` char(10) NOT NULL,
  `studentName` varchar(20) NOT NULL,
  `sex` char(2) NOT NULL,
  `birthday` date DEFAULT NULL,
  `native` varchar(20) DEFAULT NULL,
  `nation` varchar(10) DEFAULT '汉',
  `classNo` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`studentNo`),
  KEY `FK_student_copy` (`classNo`),
  CONSTRAINT `FK_student_copy` FOREIGN KEY (`classNo`) REFERENCES `tb_class` (`classNo`)
) ENGINE=InnoDB DEFAULT CHARSET=gb2312;

INSERT INTO tb_student_copy SELECT * FROM tb_student;

# 例 5.8 - 当前表`tb_student_copy`中已经存在这样一条数据记录：('2013110101', '张晓勇', '男', '1997-12-11', '山西', '汉', 'AC1301')，
# 其中该表中`studentNo`是主键，现向该表中再次插入一行数据：('2013110101', '周旭', '男', '1996-10-01', '湖南', '汉', 'AC1301')。
REPLACE INTO tb_student_copy VALUES ('2013110101', '周旭', '男', '1996-10-01', '湖南', '汉', 'AC1301');

# 例 5.9 - 将表`tb_student`中学号为'2014210101'的学生姓名修改为'黄涛'，籍贯修改为'湖北'，民族修改为缺省值。
UPDATE tb_student
SET studentName='黄涛', native='湖北', nation=NULL
WHERE studentNo='2014210101';

# 例 5.10 - 将成绩表`tb_score`中所有学生的成绩提高5%。
UPDATE tb_score
SET score=score*1.05;

# 例 5.11 - 将选修'程序设计'这门课程的学生成绩置零。
UPDATE tb_score
SET score=0
WHERE courseNo=(SELECT courseNo FROM tb_course WHERE courseName='程序设计');

SELECT studentNo, tb_score.courseNo, tb_course.courseName, tb_score.score
FROM tb_score INNER JOIN tb_course
ON tb_score.courseNo=tb_course.courseNo
WHERE tb_course.courseName='程序设计';

# 例 5.12 - 删除表`tb_student`中姓名为'王一敏'的学生信息。
DELETE FROM tb_score
WHERE studentNo=(SELECT studentNo FROM tb_student WHERE studentName='王一敏');
DELETE FROM tb_student
WHERE studentName='王一敏';

# 例 5.13 - 将'程序设计'这门课程的所有选课记录删除。
DELETE FROM tb_score
WHERE courseNo=(SELECT courseNo FROM tb_course WHERE courseName='程序设计');

# 例 5.14 - 删除所有学生的选课记录。
DELETE FROM tb_score;
DELETE FROM tb_student;

# 例 5.15 - 使用TRUNCATE语句删除数据表`tb_student`的备份`tb_student_copy`中的所有记录。
# 如果要删除表中的所有记录，还可以使用TRUNCATE语句。
# TRUNCATE语句将直接删除原来的表并重新创建一个表，而不是逐行删除表中的记录，因此执行速度会比DELETE操作更快。
TRUNCATE TABLE tb_student_copy;
