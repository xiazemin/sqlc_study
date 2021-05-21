/* name: ListCompanyById :many */
select * from company where id in (?);

/* name: GetCompanyName :one */
select companyName from company where companyName =?;
