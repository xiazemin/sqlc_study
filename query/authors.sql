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

/* name: GetTotalSizeNull :one */
SELECT sum(size) from authors WHERE id in (?);

/* name: GetMaxID :one */
SELECT MAX(id) FROM authors ;