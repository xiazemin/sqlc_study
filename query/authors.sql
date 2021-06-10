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
  id,name,bio
) VALUES (
  ?,?, ? 
);

/* name: DeleteAuthor :exec */
DELETE FROM authors
WHERE id = ?;

/* name: DeleteAuthorIn :exec */
DELETE FROM authors
WHERE id in (?);

/* name: GetTotalSize :one */
SELECT sum(size) from authors WHERE id in (?);