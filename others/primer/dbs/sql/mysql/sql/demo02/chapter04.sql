############################
# 全国计算机等级考试
# 二级教程 - MySQL数据库程序设计
############################

# 使用`db_school`
USE db_school;

############################
# 第四章 数据查询
############################

############################
# 4.2 单表查询
############################

############################
# 4.2.1 选择字段
############################

# 查询指定字段
# 例 4.1 - 查询所有班级的班级编号、所属学院和班级名称。
SELECT classNo, department, className
FROM tb_class;

# 例 4.2 - 从班级表`tb_class`中查询出所有的院系名称。
SELECT department FROM tb_class;

# 查询所有字段
# 例 4.3 - 查询全体学生的详细信息。
SELECT * FROM tb_student;

# 查询经过计算的值
# 例 4.4 - 查询全体学生的姓名，性别和年龄。
SELECT studentName, sex, CURDATE()-YEAR(birthday)
FROM tb_student;

# 定义字段的别名
# 例 4.5 - 查询全体学生的姓名、性别和年龄，要求给目标列表达式取别名。
SELECT studentName '学生', sex '性别', CURDATE()-YEAR(birthday) AS '年龄'
FROM tb_student;

############################
# 4.2.2 选择指定字段
############################

# 比较大小
# 例 4.6 - 查询课时大于等于48学时的课程名称及学分。
SELECT courseName '课程名称', credit AS '学分'
FROM tb_course
WHERE courseHour>=48;

# 例 4.7 - 查询少数民族学生的姓名、性别、籍贯和民族。
SELECT studentName '姓名', sex '性别', native '籍贯', nation '民族'
FROM tb_student
WHERE nation!='汉';

# 带BETWEEN..AND关键字的范围查询
# 例 4.8 - 查询出生日期在'1997-01-01'和'1997-12-31'之间的学生姓名、性别和出生日期。
SELECT studentName '姓名', sex '性别', birthday '出生日期'
FROM tb_student
WHERE birthday BETWEEN '1997-01-01' AND '1997-12-31';

# 例 4.9 - 查询出生日期不在'1997-01-01'和'1997-12-31'之间的学生姓名、性别和出生日期。
SELECT studentName '姓名', sex '性别', birthday '出生日期'
FROM tb_student
WHERE birthday NOT BETWEEN '1997-01-01' AND '1997-12-31';

# 带IN关键字的集合查询
# 例 4.10 - 查询籍贯是北京、天津和上海的学生信息。
SELECT *
FROM tb_student
WHERE native IN ('北京', '天津', '上海');

# 例 4.11 - 查询籍贯不是北京、天津和上海的学生信息。
SELECT *
FROM tb_student
WHERE native NOT IN ('北京', '天津', '上海');

# 带LIKE关键字的字符串匹配查询
# 例 4.12 - 查询学号为'2013110201'的学生的详细情况。
SELECT *
FROM tb_student
WHERE studentNo='2013110201';

SELECT *
FROM tb_student
WHERE studentNo LIKE '2013110201';

# 例 4.13 - 查询所有姓'王'的学生的学号、姓名和班号。
SELECT studentNo '学号', studentName '姓名', classNo '班号'
FROM tb_student
WHERE studentName LIKE '王%';

# 例 4.14 - 查询所有不姓'王'的学生的学号、姓名和班号。
SELECT studentNo '学号', studentName '姓名', classNo '班号'
FROM tb_student
WHERE studentName NOT LIKE '王%';

# 例 4.15 - 查询姓名中包含'林'字的学生学号、姓名和班号。
SELECT studentNo '学号', studentName '姓名', classNo '班号'
FROM tb_student
WHERE studentName LIKE '%林%';

# 例 4.16 - 查询姓'王'且姓名长度为三个中文字的学生的学号、姓名和班号。
SELECT studentNo '学号', studentName '姓名', classNo '班号'
FROM tb_student
WHERE studentName LIKE '王__';

# 例 4.17 - 查询课程名称中含有下划线'_'的课程信息。
SELECT *
FROM tb_course
WHERE courseName LIKE '%#_%' ESCAPE '#';

# 使用正则表达式的查询
# 例 4.18 - 查询课程名称中带有中文'系统'的课程信息。
SELECT *
FROM tb_course
WHERE courseName REGEXP '系统';

SELECT *
FROM tb_course
WHERE courseName LIKE '%系统%';

# 例 4.19 - 查询课程名称中含有'管理'、'信息'或'系统'中文字符的所有课程信息。
SELECT *
FROM tb_course
WHERE courseName REGEXP '管理|信息|系统';

# 带IS NULL关键字的空值查询
# 例 4.20 - 查询缺少先行课的课程信息。
SELECT *
FROM tb_course
WHERE priorCourse IS NULL;

# 例 4.21 - 查询所有有先行课的课程信息。
SELECT *
FROM tb_course
WHERE priorCourse IS NOT NULL;

# 带AND或OR的多条件查询
# 例 4.22 - 查询学分大于等于3且学时数大于32的课程名称、学分和学时数。
SELECT courseName '课程名称', credit '学分', courseHour '学时数'
FROM tb_course
WHERE credit>=3 AND courseHour>32;

# 例 4.23 - 查询籍贯是北京或上海的学生的姓名、籍贯和民族。
SELECT studentName '姓名', native '籍贯', nation '民族'
FROM tb_student
WHERE native='北京' OR native='上海';

# 例 4.24 - 查询籍贯是北京或湖南的少数民族男生的姓名、籍贯和民族。
SELECT studentName '姓名', native '籍贯', nation '民族'
FROM tb_student
WHERE (native='北京' OR native='湖南') AND nation!='汉' AND sex='男';

############################
# 4.2.3 对查询结果排序
############################
# 例 4.25 - 查询学生的姓名、籍贯和民族，并将查询结果按姓名升序排列。
SELECT studentName '姓名', native '籍贯', nation '民族'
FROM tb_student
ORDER BY studentName;

# 例 4.26 - 查询学生选课成绩大于85分的学号、课程号和成绩信息，并将查询结果先按学号升序排列，再按成绩降序排列。
SELECT studentNo '学号', courseNo '课程号', score '成绩'
FROM tb_score
WHERE score>85
ORDER BY studentNo, score DESC;

############################
# 4.2.4 限制查询结果的数量
############################

# 例 4.27 - 查询成绩排名第3至第5的学生学号、课程号和成绩。
SELECT studentNo '学生学号', courseNo '课程号', score '成绩'
FROM tb_score
ORDER BY score DESC
LIMIT 2,3;

SELECT studentNo '学生学号', courseNo '课程号', score '成绩'
FROM tb_score
ORDER BY score DESC
LIMIT 3 OFFSET 2;


############################
# 4.3 分组聚合查询
############################

############################
# 4.3.1 使用聚合函数查询
############################
# 例 4.28 - 查询学生总人数。
SELECT COUNT(*)
FROM tb_student;

# 例 4.29 - 查询选修了课程的学生总人数。
SELECT COUNT(DISTINCT studentNo)
FROM tb_score;

# 例 4.30 - 计算选修课程编号为'21001'的学生平均成绩。
SELECT AVG(score)
FROM tb_score
WHERE courseNo='21001';

# 例 4.31 - 计算选修课程编号为'21001'的学生最高分。
SELECT MAX(score)
FROM tb_score
WHERE courseNo='21001';

############################
# 4.3.2 分组聚合查询
############################
# 例 4.32 - 查询各个课程号及相关的选课人数。
SELECT courseNo '课程号', COUNT(studentNo) '选课人数'
FROM tb_score
GROUP BY courseNo;

# 例 4.33 - 查询每个学生的选课门数、平均分和最高分。
SELECT studentNo '学生学号', COUNT(courseNo) '选课门数', AVG(score) '平均分', MAX(score) '最高分'
FROM tb_score
GROUP BY studentNo;

# 例 4.34 - 查询平均分在80分以上的每个同学的选课门数、平均分和最高分。
SELECT studentNo '学生学号', COUNT(courseNo) '选课门数', AVG(score) '平均分', MAX(score) '最高分'
FROM tb_score
GROUP BY studentNo
HAVING AVG(score)>=80;

# 例 4.35 - 查询有2门以上（含2门）课程成绩大于88分的学生学号及（88分以上的）课程数。
SELECT studentNo '学生学号', COUNT(courseNo) '选课门数'
FROM tb_score
WHERE score>88
GROUP BY studentNo
HAVING COUNT(courseNo)>=2;

# 例 4.36 - 查询所有学生选课的平均成绩，但只有当平均成绩大于80的情况下才输出。
SELECT AVG(score)
FROM tb_score
HAVING AVG(score)>80;


############################
# 4.4 连接查询
############################

############################
# 4.4.1 交叉链接
############################
# 例 4.37 - 查询学生表与成绩表的交叉连接。
SELECT *
FROM tb_student CROSS JOIN tb_score;

SELECT *
FROM tb_student, tb_score;

############################
# 4.4.2 内连接
############################
# 等值与非等值连接
# 例 4.38 - 查询每个学生选修课程的情况。
SELECT tb_student.*, tb_score.*
FROM tb_student, tb_score
WHERE tb_student.studentNo=tb_score.studentNo;

SELECT tb_student.*, tb_score.*
FROM tb_student INNER JOIN tb_score
ON tb_student.studentNo=tb_score.studentNo;

# 例 4.39 - 查询会计学院全体同学的学号、姓名、籍贯、班级编号和所在班级名称。
SELECT tb_student.studentNo '学生学号', tb_student.studentName '姓名', tb_student.native '籍贯', tb_student.classNo '班级编号', tb_class.className '班级名称'
FROM tb_student INNER JOIN tb_class
ON tb_student.classNo=tb_class.classNo
WHERE tb_class.department='会计学院';

# 例 4.40 - 查询选修了课程名称为'程序设计'的学生学号、姓名和成绩。
SELECT tb_student.studentNo '学生学号', tb_student.studentName '学生姓名', tb_score.score '成绩'
FROM tb_student INNER JOIN tb_score INNER JOIN tb_course
ON tb_student.studentNo=tb_score.studentNo AND tb_score.courseNo=tb_course.courseNo
WHERE tb_course.courseName='程序设计';

# 自连接
# 例 4.41 - 查询与'数据库'这门课学分相同的课程信息。
SELECT c2.*
FROM tb_course c1, tb_course c2
WHERE c2.credit=c1.credit AND c1.courseName='数据库' AND c2.courseName!='数据库';

SELECT c2.*
FROM tb_course c1 INNER JOIN tb_course c2
ON c2.credit=c1.credit
WHERE c1.courseName='数据库' AND c2.courseName!='数据库';

# 自然连接
# 例 4.42 - 用自然链接查询每个学生及其选修课程的情况，要求显示学生学号、姓名、选修的课程号和成绩。
SELECT tb_student.studentNo '学号', tb_student.studentName '姓名', tb_score.courseNo '选修课程号', tb_score.score '成绩'
FROM tb_student NATURAL JOIN tb_score;

############################
# 4.4.3 外连接
############################
# 左外连接
# 例 4.43 - 使用左外链接查询所有学生及其选修课程的情况，包括没有选修课程的学生，要求显示学号、姓名、性别、班号、选修的课程号和成绩。
SELECT stu.studentNo '学号', stu.studentName '姓名', stu.sex '性别', stu.classNo '班级号', sco.courseNo '选修课程号', sco.score '成绩'
FROM tb_student AS stu LEFT OUTER JOIN tb_score sco
ON stu.studentNo=sco.studentNo;

# 右外连接
# 例 4.44 - 使用右外链接查询所有学生及其选修课程的情况，包括没有选修课程的学生，要求显示学号、姓名、性别、班号、选修的课程号和成绩。
SELECT stu.studentNo '学号', stu.studentName '姓名', stu.sex '性别', stu.classNo '班级号', sco.courseNo '选修课程号', sco.score '成绩'
FROM tb_score sco RIGHT OUTER JOIN tb_student AS stu
ON stu.studentNo=sco.studentNo;


############################
# 4.5 子查询
############################

############################
# 4.5.1 带IN关键字的子查询
############################
# 例 4.45 - 查询选修了课程的学生姓名。
SELECT studentName
FROM tb_student
WHERE studentNo IN (SELECT DISTINCT studentNo
                    FROM tb_score);

# 例 4.46 - 查询没有选修过课程的学生姓名。
SELECT studentName
FROM tb_student
WHERE studentNo NOT IN (SELECT DISTINCT studentNo
                        FROM tb_score);

############################
# 4.5.2 带比较运算符的子查询
############################
# 例 4.47 - 查询班级'计算机 14-1 班'所有学生的学号、姓名。
SELECT studentNo '学号', studentName '姓名'
FROM tb_student
WHERE classNo=(SELECT classNo FROM tb_class WHERE className='计算机 14-1 班');

# 例 4.48 - 查询与'李明'在同一个班学习的学生学号、姓名和班号。
SELECT studentNo '学号', studentName '姓名', classNo '班号'
FROM tb_student
WHERE classNo=(SELECT classNo FROM tb_student WHERE studentName='李明') AND studentName!='李明';

# 例 4.49 - 查询男生中比`某个`女生出生`年份晚`的学生姓名和出生年份。
SELECT studentName '姓名', birthday '出生日期'
FROM tb_student
WHERE sex='男' AND YEAR(birthday) > ANY(SELECT YEAR(birthday)
                                  FROM tb_student
                                  WHERE sex='女');

# 例 4.50 - 查询男生中比`所有`女生出生`年份晚`的学生姓名和出生年份。
SELECT studentName '姓名', birthday '出生日期'
FROM tb_student
WHERE sex='男' AND YEAR(birthday) > ALL(SELECT YEAR(birthday)
                                  FROM tb_student
                                  WHERE sex='女');

############################
# 4.5.3 带EXISTS关键字的子查询
############################
# 例 4.51 - 查询选修了课程号为'31002'的学生姓名。
SELECT studentName '学生姓名'
FROM tb_student
WHERE EXISTS (SELECT *
              FROM tb_score
              WHERE tb_score.studentNo=tb_student.studentNo AND courseNo='31002');

# 例 4.52 - 查询没有选修课程号为'31002'的学生姓名。
SELECT studentName '学生姓名'
FROM tb_student
WHERE NOT EXISTS (SELECT *
                  FROM tb_score
                  WHERE tb_score.studentNo=tb_student.studentNo AND courseNo='31002');

# 例 4.53 - 查询查询选修了全部课程的学生姓名。
# 否定的否定，即肯定。
SELECT studentName '学生姓名'
FROM tb_student
WHERE NOT EXISTS (SELECT *
                  FROM tb_course
                  WHERE NOT EXISTS (SELECT *
                                    FROM tb_score
                                    WHERE studentNo=tb_student.studentNo AND courseNo=tb_course.courseNo));


############################
# 4.6 联合查询（UNION）
############################
# 例 4.54 - 使用'UNION'查询选修了'管理学'或'计算机基础'的学生学号。
SELECT studentNo '学生学号'
FROM tb_score NATURAL JOIN tb_course
WHERE tb_course.courseName='管理学'
UNION
SELECT studentNo '学生编号'
FROM tb_score NATURAL JOIN tb_course
WHERE tb_course.courseName='计算机基础';

SELECT DISTINCT studentNo '学生学号'
FROM tb_score NATURAL JOIN tb_course
WHERE tb_course.courseName='管理学' OR tb_course.courseName='计算机基础';

# 例 4.55 - 使用'UNION ALL'查询选修了'管理学'或'计算机基础'的学生学号。
SELECT studentNo '学生学号'
FROM tb_score NATURAL JOIN tb_course
WHERE tb_course.courseName='管理学'
UNION ALL
SELECT studentNo '学生学号'
FROM tb_score NATURAL JOIN tb_course
WHERE tb_course.courseName='计算机基础';

# 例 4.56 - 查询选修了'计算机基础'和'管理学'的学生学号。
SELECT studentNo '学生学号'
FROM tb_score NATURAL JOIN tb_course
WHERE tb_course.courseName='计算机基础'
AND studentNo IN (SELECT studentNo
                  FROM tb_score NATURAL JOIN tb_course
                  WHERE tb_course.courseName='管理学');

# 例 4.57 - 查询选修了'计算机基础'但没有选修'管理学'的学生学号。
SELECT studentNo '学生学号'
FROM tb_score NATURAL JOIN tb_course
WHERE tb_course.courseName='计算机基础'
AND studentNo NOT IN (SELECT studentNo
                  FROM tb_score NATURAL JOIN tb_course
                  WHERE tb_course.courseName='管理学');
