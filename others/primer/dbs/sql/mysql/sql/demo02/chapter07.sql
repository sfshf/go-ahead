############################
# 全国计算机等级考试
# 二级教程 - MySQL数据库程序设计
############################

# 使用数据库`db_school`
USE `db_school`;

############################
# 第七章 视图
############################

# 例 7.1 - 在数据库`db_school`中创建视图`v_student`，要求该视图包含客户信息表`tb_student`中所有男生的信息，
# 并且要求保证今后对该视图数据的修改都必须符合学生性别为男性这个条件。
CREATE OR REPLACE VIEW `db_school`.`v_student`
AS
SELECT * FROM `db_school`.`tb_student` WHERE sex='男'
WITH CHECK OPTION;

# 例 7.2 - 在数据库`db_school`中创建视图`db_school.v_score_avg`，
# 要求该视图包含表`tb_score`中所有学生的学号和平均成绩，并按学号`studentNo`进行排序。
CREATE OR REPLACE VIEW `db_school`.`v_score_avg`(studentNo, score_avg)
AS
SELECT studentNo, AVG(score) FROM tb_score GROUP BY studentNo;

# 例 7.3 - 举一个例子，针对前面的SELECT语句的`规则1~8`，最好能涵盖几个；也可以举一个产生错误的例子。
# 创建视图语句中的`SELECT语句`，存在以下一些限制：
# 1、定义视图的用户除了要求被授予`CREATE VIEW`的权限之外，还必须被授予可以操作视图所涉及的基础表或其他视图的相关权限，
# 例如，由`SELECT语句`选择的每一列上的某些权限。
# 2、`SELECT语句`不能包含`FROM子句`中的子查询。
# 3、`SELECT语句`不能引用系统变量或用户变量。
# 4、`SELECT语句`不能引用预处理语句参数。
# 5、在`SELECT语句`中引用的表或视图必须存在。但是，创建完视图后，可以删除视图定义中所引用的基础表或源视图。
# 若想检查视图定义是否存在这类问题，可使用`CHECK TABLE语句`。
# 6、若`SELECT语句`中所引用的不是当前的数据库的基础表或源视图时，需要在该表或视图前加上数据库的名称作为限定前缀。
# 7、在由`SELECT语句`构造的视图定义中，允许使用`ORDER BY子句`。但是，如果从特定视图进行了选择，
# 而该视图使用了自己的`ORDER BY语句`，则视图定义中的`ORDER BY子句`将被忽略。
# 8、对于`SELECT语句`中的其他选项或子句，若所创建的视图中也包含了这些选项，则语句执行效果未定义。
# 例如，如果在视图定义中包含了`LIMIT子句`，而`SELECT语句`也使用了自己的`LIMIT子句`，那么MySQL对使用哪一个`LIMIT语句`未做定义。

# 例 7.4 - 举一个例子，说明`WITH CHECK OPTION`中的内容。
# 针对数据库`db_school`中的表`db_score`，使用`WITH CHECK OPTION子句`创建视图`v_score`，要求该视图包含表`tb_score`中
# 所有score<90的学生学号、课程号和成绩信息；
# 分别使用`WITH LOCAL CHECK OPTION`、`WITH CASCADED CHECK OPTION`子句创建
# 视图`v_score_local`和`v_score_cascaded`，要求该视图包含表`tb_score`中所有score>80的学生学号、课程号和成绩信息。
CREATE OR REPLACE VIEW `db_school`.`v_score`
AS
SELECT studentNo, courseNo, score FROM tb_score WHERE score<90
WITH CHECK OPTION;

CREATE OR REPLACE VIEW `db_school`.`v_score_local`
AS
SELECT * FROM v_score WHERE score>80
WITH LOCAL CHECK OPTION;

CREATE OR REPLACE VIEW `db_school`.`v_score_cascaded`
AS
SELECT * FROM v_score WHERE score>80
WITH CASCADED CHECK OPTION;

# 在这里，视图`v_score_local`和`v_score_cascaded`是根据视图`v_score`定义的。
# 视图`v_score_local`具有`LOCAL`检查选项，因此，仅会针对其自身检查对插入项进行测试；
# 视图`v_score_cascaded`含有`CASCADED`检查选项，因此，不仅会针对它自己的检查对插入项进行测试，也会针对基本视图`v_score`的检查对插入项进行测试。
# 通过下列插入语句可以清楚地分辨彼此之间的差异：
INSERT INTO db_school.v_score_local VALUES ('2013110101', '21005', 90);

INSERT INTO db_school.v_score_cascaded VALUES ('2013110101', '21005', 90);

# 例 7.5 - 删除数据库`db_school`中的视图`v_student`。
DROP VIEW IF EXISTS `db_school`.`v_student`;

# 例 7.6 - 使用`ALTER VIEW语句`修改数据库`db_school`中的视图`v_student`的定义，
# 要求该视图包含学生表`tb_student`中性别为'男'、民族为'汉'的学生的学号、姓名和所属班级，
# 并且要求保证今后对该视图数据的修改都必须符合学生性别为'男'、民族为'汉'这个条件。
ALTER VIEW `db_school`.`v_student`
AS
SELECT studentNo, studentName, classNo
FROM tb_student
WHERE sex='男' AND nation='汉'
WITH CHECK OPTION;

# 例 7.7 - 使用`CREATE OR REPLACE VIEW语句`修改数据库`db_school`中的视图`v_student`的定义，
# 要求该视图包含学生表`tb_student`中性别为'男'、民族为'汉'的学生，
# 并且要求保证今后对该视图数据的修改都必须符合学生性别为'男'、民族为'汉'这个条件。
CREATE OR REPLACE VIEW `db_school`.`v_student`
AS
SELECT *
FROM tb_student
WHERE sex='男' AND nation='汉'
WITH CHECK OPTION;

# 例 7.8 - 查看数据库`db_school`中视图`v_course`的定义。
SHOW CREATE VIEW `db_school`.`v_course`\G

# 例 7.9 - 在数据库`db_school`中，向视图`v_student`中插入下面一条记录：('2014310108', '周明', '男', '1997-08-16', '辽宁', '汉', 'IS1401')。
INSERT INTO db_school.v_student VALUES ('2014310108', '周明', '男', '1997-08-16', '辽宁', '汉', 'IS1401');

# 例 7.10 - 将视图`v_student`中所有学生的`native`列更新为'河南'。
UPDATE v_student
SET native='河南';

# 例 7.11 - 删除视图`v_student`中姓名为'周明'的学生信息。
DELETE FROM v_student
WHERE studentName='周明';

# 例 7.12 - 在视图`v_student`中查找`classNo`为'CS1401'的学生学号和姓名。
SELECT studentNo, studentName
FROM v_student
WHERE classNo='CS1401';

# 练习题2 - 在数据库`db_school`中创建视图`v_score`，要求该视图包含成绩表`tb_score`中所有成绩在90分以上的成绩信息，
# 并且要求保证今后对该视图数据的修改都必须符合成绩大于90这个条件。
CREATE OR REPLACE VIEW `db_school`.`v_score`
AS
SELECT *
FROM tb_score
WHERE score>90
WITH CHECK OPTION;

# 练习题3 - 在视图`v_score`中查找`classNo`为'21002'的学生的学号和成绩。
SELECT v_score.studentNo, v_score.score
FROM v_score INNER JOIN tb_student
ON v_score.studentNo=tb_student.studentNo
WHERE classNo='21002';

# 练习题4 - 在数据库`db_school`中，向视图`v_score`中插入下面一条记录：('2014310101', '31005', 95)。
INSERT INTO v_score VALUES ('2014310101', '31005', 95);

# 练习题5 - 删除视图`v_score`中学号为'2014310101'的学生成绩信息。
DELETE FROM v_score WHERE studentNo='2014310101';
