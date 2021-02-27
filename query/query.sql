-- name: GetAuthor :one
SELECT * FROM authors WHERE id=? and name=? LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors where id in ? ORDER BY name;

-- name: CreateAuthor :exec
INSERT INTO authors (
  name, bio
) VALUES (
  ?, ?
);

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = ?;