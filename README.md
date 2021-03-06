https://segmentfault.com/a/1190000022529075?utm_source=tag-newest

localhost:sqlc xiazemin$ ln -s ../sqlc_study/query/ .
localhost:sqlc xiazemin$ ln -s ../sqlc_study/schema/ .
localhost:sqlc xiazemin$ ln -s ../sqlc_study/sqlc.yaml/ .

mysql> insert into company (id,name) values (0,"a");
mysql> insert into company (id,name) values (1,"a");
mysql> insert into authors (id,name,bio,company_id) values(15,'Brian','Brian',1);

 ~/go/bin/sqlc generate
