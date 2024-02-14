// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
)

const deleteBarByID = `-- name: DeleteBarByID :one
DELETE FROM bars WHERE id = $1 RETURNING id, name
`

func (q *Queries) DeleteBarByID(ctx context.Context, id int32) (Bar, error) {
	row := q.db.QueryRow(ctx, deleteBarByID, id)
	var i Bar
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteExclusionByID = `-- name: DeleteExclusionByID :one
DELETE FROM exclusions WHERE id = $1 RETURNING id, name
`

func (q *Queries) DeleteExclusionByID(ctx context.Context, id int32) (Exclusions, error) {
	row := q.db.QueryRow(ctx, deleteExclusionByID, id)
	var i Exclusions
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteMyDataByID = `-- name: DeleteMyDataByID :one
DELETE FROM my_data WHERE id = $1 RETURNING id, name
`

func (q *Queries) DeleteMyDataByID(ctx context.Context, id int32) (MyData, error) {
	row := q.db.QueryRow(ctx, deleteMyDataByID, id)
	var i MyData
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
