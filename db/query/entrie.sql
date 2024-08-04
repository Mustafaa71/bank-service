-- name: CreateEntrie :one
INSERT INTO entrie (
  accoount_id,
  amount 
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetEntrie :one
SELECT * FROM entrie
WHERE id = $1 LIMIT 1;

-- name: ListEntrie :many
SELECT * FROM entrie
WHERE accoount_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

