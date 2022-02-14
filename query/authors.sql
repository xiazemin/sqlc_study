-- name: GetAuthorsInCompanyById :many
SELECT * FROM authors where company_id in ( select id from company where id in (?) );

-- name: GetAuthorsInCompany :many
SELECT * FROM authors where company_id in ( select id from company where id in (?) and name in (?) );

-- name: GetAuthorsInOneCompany :many
SELECT * FROM authors where company_id in ( select id from company where id = ? );

-- name: GetOneAuthor :one
SELECT * FROM authors where  id in (?)  and bio=? and  name in (?) and company_id in (?) limit 1; 


-- name: ListAuthors :many
SELECT * FROM authors where  bio=? and id in (?)  and name in (?)  ORDER BY name;


/* name: ListAllAuthors :many */
SELECT * FROM authors
ORDER BY name;

/* name: CreateAuthor :execresult */
INSERT INTO authors (
  id,name,bio,company_id
) VALUES (
  ?,?, ?,1 
);

/* name: DeleteAuthor :exec */
DELETE FROM authors
WHERE id = ?;

/* name: DeleteAuthorIn :exec */
DELETE FROM authors
WHERE id in (?);

/* name: GetTotalSize :one */
SELECT ifnull(sum(size),0) from authors WHERE id in (?);

/* name: GetTotalSizeNullIn :one */
SELECT sum(size) from authors WHERE id in (?);

/* name: GetTotalSize1NullIn :one */
SELECT sum(size1) from authors WHERE id in (?);

/* name: GetTotalSizeNull :one */
SELECT sum(size) from authors WHERE id =?;

/* name: GetTotalSize1Null :one */
SELECT max(size1) from authors WHERE id = ?;

/* name: GetMaxID :one */
SELECT MAX(id) FROM authors ;

/* name: GetLargestID :one */
SELECT id FROM authors order by id desc limit 1;

/* name: GetMax_empty_col :one */
select max(empty_col) from authors where id =? ;

/* name: GetMax_default_col :one */
select max(default_col) from authors where id =? ;

/* name: GetMax_default_col1 :one */
select max(default_col1) from authors where id =? ;

/* name: ListAuthorsOmit :many omits:[name,bio]*/
select * from authors where id=? and name=? or bio=?;