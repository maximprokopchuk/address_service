// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAddress = `-- name: CreateAddress :one
INSERT INTO address (
  name, type, parent_id
) VALUES (
  $1, $2, $3
)
RETURNING id, type, name, parent_id
`

type CreateAddressParams struct {
	Name     string
	Type     string
	ParentID pgtype.Int4
}

func (q *Queries) CreateAddress(ctx context.Context, arg CreateAddressParams) (Address, error) {
	row := q.db.QueryRow(ctx, createAddress, arg.Name, arg.Type, arg.ParentID)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Name,
		&i.ParentID,
	)
	return i, err
}

const deleteAddress = `-- name: DeleteAddress :exec
DELETE FROM address
WHERE id = $1
`

func (q *Queries) DeleteAddress(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAddress, id)
	return err
}

const getAddress = `-- name: GetAddress :one
SELECT id, type, name, parent_id FROM address
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAddress(ctx context.Context, id int64) (Address, error) {
	row := q.db.QueryRow(ctx, getAddress, id)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Name,
		&i.ParentID,
	)
	return i, err
}

const getAddressesByParent = `-- name: GetAddressesByParent :many
SELECT id, type, name, parent_id FROM address
WHERE parent_id IS NULL
ORDER BY name
`

func (q *Queries) GetAddressesByParent(ctx context.Context) ([]Address, error) {
	rows, err := q.db.Query(ctx, getAddressesByParent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Address
	for rows.Next() {
		var i Address
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Name,
			&i.ParentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAddressesForParent = `-- name: ListAddressesForParent :many
SELECT id, type, name, parent_id FROM address
WHERE parent_id = $1
ORDER BY name
`

func (q *Queries) ListAddressesForParent(ctx context.Context, parentID pgtype.Int4) ([]Address, error) {
	rows, err := q.db.Query(ctx, listAddressesForParent, parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Address
	for rows.Next() {
		var i Address
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Name,
			&i.ParentID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
