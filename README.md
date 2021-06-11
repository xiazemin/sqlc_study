支持in 普通语法 和子查询的sqlc
https://github.com/xiazemin/sqlc

localhost:sqlc xiazemin$ ln -s ../sqlc_study/query/ .
localhost:sqlc xiazemin$ ln -s ../sqlc_study/schema/ .
localhost:sqlc xiazemin$ ln -s ../sqlc_study/sqlc.yaml/ .

mysql> insert into company (id,name) values (0,"a");
mysql> insert into company (id,name) values (1,"a");
mysql> insert into authors (id,name,bio,company_id) values(15,'Brian','Brian',1);

 ~/go/bin/sqlc generate


生成Interface
    emit_interface: true

生成mock
go generate ./...

参考：
https://segmentfault.com/a/1190000022529075?utm_source=tag-newest


mysql> select * from authors;
+----+-----------------+-------------------------------------------------------------------------+------------+------+-----------+-------------+-------+--------------+
| id | name            | bio                                                                     | company_id | size | empty_col | default_col | size1 | default_col1 |
+----+-----------------+-------------------------------------------------------------------------+------------+------+-----------+-------------+-------+--------------+
|  1 | a               | a                                                                       |          1 |    1 |      NULL |           0 |  NULL |            1 |
|  2 | b               | b                                                                       |          2 | NULL |      NULL |           0 |  NULL |            0 |
| 10 | Brian Kernighan | Co-author of The C Programming Language and The Go Programming Language |          1 | NULL |      NULL |           0 |  NULL |            0 |
| 11 | Brian Kernighan | Co-author of The C Programming Language and The Go Programming Language |          1 | NULL |      NULL |           0 |  NULL |            0 |
| 12 | Brian Kernighan | Co-author of The C Programming Language and The Go Programming Language |          1 | NULL |      NULL |           0 |  NULL |            0 |
| 13 | Brian Kernighan | Co-author of The C Programming Language and The Go Programming Language |          1 | NULL |      NULL |           0 |  NULL |            0 |
| 14 | Brian Kernighan | Co-author of The C Programming Language and The Go Programming Language |          1 | NULL |      NULL |           0 |  NULL |            0 |
| 15 | Brian Kernighan | Co-author of The C Programming Language and The Go Programming Language |          1 | NULL |      NULL |           0 |  NULL |            0 |
| 16 | Brian Kernighan | Co-author of The C Programming Language and The Go Programming Language |          1 | NULL |      NULL |           0 |  NULL |            0 |
| 17 | Brian Kernighan | Co-author of The C Programming Language and The Go Programming Language |          1 | NULL |      NULL |           0 |  NULL |            0 |
| 18 | Brian Kernighan | Co-author of The C Programming Language and The Go Programming Language |          1 | NULL |      NULL |           0 |  NULL |            0 |
| 19 | Brian Kernighan | Co-author of The C Programming Language and The Go Programming Language |          1 | NULL |      NULL |           0 |  NULL |            0 |
| 20 | Brian Kernighan | Co-author of The C Programming Language and The Go Programming Language |          1 | NULL |      NULL |           0 |  NULL |            0 |
| 21 | Brian Kernighan | Co-author of The C Programming Language and The Go Programming Language |          1 | NULL |      NULL |           0 |  NULL |            0 |
| 22 | Brian Kernighan | Co-author of The C Programming Language and The Go Programming Language |          1 | NULL |      NULL |           0 |  NULL |            0 |
+----+-----------------+-------------------------------------------------------------------------+------------+------+-----------+-------------+-------+--------------+


使用聚合函数的情况下


如果NOT NULL，没有命中会返回 NULL ,且走了索引才会，in 不会
             命中了 返回默认值
如果DEFAULT NULL，但是没有插入数据，且没有命中会返回NULL  用了sql.null不会报错
                            命中了 返回NULL
              如果插入数据了，没有命中会返回NULL
                            命中了 返回NULL




mysql> select max(default_col) from authors where id=11234567890;
+------------------+
| max(default_col) |
+------------------+
|             NULL |
+------------------+
1 row in set (0.00 sec)

mysql> select max(default_col) from authors where id=1;
+------------------+
| max(default_col) |
+------------------+
|                0 |
+------------------+
1 row in set (0.00 sec)

mysql> select max(empty_col) from authors where id=11234567890;
+----------------+
| max(empty_col) |
+----------------+
|           NULL |
+----------------+
1 row in set (0.00 sec)

mysql> select max(empty_col) from authors where id=1;
+----------------+
| max(empty_col) |
+----------------+
|           NULL |
+----------------+
1 row in set (0.00 sec)



/**
mysql> select empty_col from authors where id=1;
+-----------+
| empty_col |
+-----------+
|      NULL |
+-----------+
1 row in set (0.01 sec)

mysql> select empty_col from authors where id=11234567890;
Empty set (0.00 sec)

mysql> select default_col from authors where id=1;
+-------------+
| default_col |
+-------------+
|           0 |
+-------------+
1 row in set (0.00 sec)

mysql> select default_col from authors where id=11234567890;
Empty set (0.00 sec)

mysql> 
*/

GetMax_default_col 0 sql: Scan error on column index 0, name "max(default_col)": converting NULL to int32 is unsupported
2021/06/11 18:33:26 CreateAuthor error:Error 1364: Field 'default_col' doesn't have a default value


