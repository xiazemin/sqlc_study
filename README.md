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


