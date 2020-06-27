############################
# 全国计算机等级考试
# 二级教程 - MySQL数据库程序设计
############################

############################
# 第五章 练习题
############################

# 使用数据库`db_sp`
USE db_sp;

# 三、应用题
# 2、请使用UPDATE语句将数据库`db_sp`的表`P`中蓝色零件的重量增加20%。
UPDATE P SET WEIGHT=WEIGHT*1.2 WHERE COLOR='Blue';

# 3、请使用DELETE语句将数据库`db_sp`的表`S`中状态为空值的供应商信息删除。
DELETE FROM S
WHERE STATUS IS NULL;

# 4、请使用DELETE语句删除数据库`db_sp`中没有供应零件的供应商信息。
DELETE FROM S
WHERE SNO NOT IN (SELECT DISTINCT SNO FROM SP);
