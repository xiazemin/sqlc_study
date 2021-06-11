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


使用聚合函数的情况下

结论：如果一个字段是not null ，只有在数据库里这个字段是空的情况下会报错，问题：是不是应该加column的同时给历史数据加默认值？也可以不加

如果NOT NULL，但是没有插入数据，且没有命中会返回 NULL （出现在新增字段的场景，或者数据库初始化的场景） //只有这种场景会报错
                            命中了 返回默认值
              如果插入数据了，没有命中会返回默认值
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


