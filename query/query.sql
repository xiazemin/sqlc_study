-- name: GetOneAuthor :one
SELECT * FROM authors where  bio=? and id in (?)  and name in (?)  limit 1;


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

