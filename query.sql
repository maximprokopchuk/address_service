-- name: GetAddress :one
SELECT * FROM address
WHERE id = $1 LIMIT 1;

-- name: ListAddressesForParent :many
SELECT * FROM address
WHERE parent_id = $1
ORDER BY name;

-- name: GetAddressesByParent :many
SELECT * FROM address
WHERE parent_id IS NULL
ORDER BY name;

-- name: CreateAddress :one
INSERT INTO address (
  name, type, parent_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteAddress :exec
DELETE FROM address
WHERE id = $1;