############################
# 全国计算机等级考试
# 二级教程 - MySQL数据库程序设计
############################

# 使用数据库`db_school`
USE `db_school`;

############################
# 第八章 触发器
############################

# 例 8.1 - 在数据库`db_school`的表`tb_student`中创建一个触发器`tb_student_insert_trigger`，用于每次向表`tb_student`中插入一行数据时将学生变量`str`的值设置为"one student added!"。
# 例 8.2 - 删除数据库`db_school`中的触发器`tb_student_insert_trigger`。
# 例 8.3 - 在数据库`db_school`的表`tb_student`中重新创建触发器`tb_student_insert_trigger`，用于每次向表`tb_student`中插入一行数据时将学生变量`str`的值设置为新插入学生的学号。
# 例 8.4 - 在数据库`db_school`的表`tb_student`中创建一个触发器`tb_student_update_trigger`，用于每次更新表`tb_student`时将表中`nation`列的值设置为`native`列的值。

# 练习题 1 - 在数据库`db_test`的表`content`中创建一个触发器`content_delete_trigger`，用于每次当删除表`content`中一行数据时将用户变量`str`设置为"old content deleted!"。
# 练习题 2 - 在数据库`db_score`的表`tb_score`中创建触发器`tb_score_insert_trigger`，用于每次向表`tb_score`插入一行数据时将成绩变量`str`的值设置为"new score record added!"。
# 练习题 3 - 在数据库`db_score`的表`tb_score`中创建一个触发器`tb_score_update_trigger`，用于每次更新表`tb_score`时，将该表中`score`列的值在原值的基础上加1。
# 练习题 4 - 删除数据库`db_score`中的触发器`tb_score_insert_trigger`。
