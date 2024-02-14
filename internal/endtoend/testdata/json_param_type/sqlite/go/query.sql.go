// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const findByAddress = `-- name: FindByAddress :one
SELECT id, metadata FROM "user" WHERE "metadata"->>'address1' = ?1 LIMIT 1
`

func (q *Queries) FindByAddress(ctx context.Context, metadata sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, findByAddress, metadata)
	var i User
	err := row.Scan(&i.ID, &i.Metadata)
	return i, err
}