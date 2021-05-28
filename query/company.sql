/* name: ListCompanyById :many */
select * from company where id in (?);

/* name: GetCompanyName :one */
select companyName from company where companyName =?;

/* name: InsertMulti :execresult */
insert into company (id,name,companyName) values (?,?,?),(?,?,?);

/*  name: CompanysRange :execresult */
select * from company where id > ? AND  id < ?;

/*  name: Companys :execresult */
select * from company where id BETWEEN ? AND ?;
