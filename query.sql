-- name: GetAddress :one
SELECT * FROM address
WHERE id = $1 LIMIT 1;

-- name: ListAddressesForParent :many
SELECT * FROM address
WHERE parent = $1
ORDER BY name;

-- name: ListTopLevelAddresses :many
SELECT * FROM address
WHERE parent IS NULL
ORDER BY name;

-- name: CreateAddress :one
INSERT INTO address (
  name, type, parent
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteAddress :exec
DELETE FROM address
WHERE id = $1;