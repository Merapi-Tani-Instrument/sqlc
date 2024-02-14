// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
)

const schemaScopedCreate = `-- name: SchemaScopedCreate :one
INSERT INTO foo.bar (id, name) VALUES ($1, $2) RETURNING id
`

type SchemaScopedCreateParams struct {
	ID   int32
	Name string
}

func (q *Queries) SchemaScopedCreate(ctx context.Context, arg SchemaScopedCreateParams) (int32, error) {
	row := q.db.QueryRow(ctx, schemaScopedCreate, arg.ID, arg.Name)
	var id int32
	err := row.Scan(&id)
	return id, err
}
