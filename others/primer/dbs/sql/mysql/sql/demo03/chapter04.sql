############################
# 全国计算机等级考试
# 二级教程 - MySQL数据库程序设计
############################

############################
# 第四章 练习题
############################

# 三、编程题
# 1、查询供应零件号为'P1'的供应商号码。
SELECT DISTINCT SNO '供应商编号'
FROM SP
WHERE PNO='P1';

# 2、查询供货量在300~500之间的所有供货情况。
SELECT *
FROM SP
WHERE QTY BETWEEN 300 AND 500;

# 3、查询供应红色零件的供应商号码和供应商名称。
SELECT DISTINCT S.SNO '供应商编号', S.SNAME '供应商名称'
FROM S NATURAL JOIN SP NATURAL JOIN P
WHERE P.COLOR='Red';

# 4、查询重量在15之下，Paris供应商供应的零件代码和零件名。
SELECT P.PNO '零件编号', P.PNAME '零件名'
FROM P NATURAL JOIN SP NATURAL JOIN S
WHERE P.WEIGHT<=15 AND S.CITY='Paris';

# 5、查询由London供应商供应的零件名称。
SELECT P.PNAME '零件名称'
FROM P NATURAL JOIN SP NATURAL JOIN S
WHERE S.CITY='London';

# 6、查询不供应红色零件的供应商名称。
SELECT DISTINCT S.SNAME '供应商名称'
FROM S NATURAL JOIN SP NATURAL JOIN P
WHERE P.COLOR!='Red';

# 7、查询供应商S3没有供应的零件名称。
SELECT PNAME '零件名称'
FROM P
WHERE PNO NOT IN (SELECT PNO
                  FROM SP
                  WHERE SNO='S3');

# 8、查询供应零件代码为P1和P2两种零件的供应商名称。
SELECT SNAME '供应商名称'
FROM S NATURAL JOIN SP
WHERE PNO='P1' AND SNO IN (SELECT SNO
                             FROM SP
                             WHERE PNO='P2');

# 9、查询与零件名Nut颜色相同的零件代码和零件名称。
SELECT P2.PNO '零件编码', P2.PNAME '零件名称'
FROM P P1 INNER JOIN P P2
ON P2.COLOR=P1.COLOR
WHERE P1.PNAME='Nut' AND P2.PNAME!='Nut';

# 10、查询供应了全部零件的供应商名称。
# 没有一种零件是他没有供应的。
SELECT S.SNAME '供应商名称'
FROM S
WHERE NOT EXISTS (SELECT *
                  FROM P
                  WHERE NOT EXISTS (SELECT *
                                    FROM SP
                                    WHERE SP.PNO=P.PNO
                                    AND SP.SNO=S.SNO));
