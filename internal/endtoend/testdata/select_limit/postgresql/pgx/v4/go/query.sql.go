// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const fooLimit = `-- name: FooLimit :many
SELECT a FROM foo
LIMIT $1
`

func (q *Queries) FooLimit(ctx context.Context, limit int32) ([]sql.NullString, error) {
	rows, err := q.db.Query(ctx, fooLimit, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var a sql.NullString
		if err := rows.Scan(&a); err != nil {
			return nil, err
		}
		items = append(items, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const fooLimitOffset = `-- name: FooLimitOffset :many
SELECT a FROM foo
LIMIT $1 OFFSET $2
`

type FooLimitOffsetParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) FooLimitOffset(ctx context.Context, arg FooLimitOffsetParams) ([]sql.NullString, error) {
	rows, err := q.db.Query(ctx, fooLimitOffset, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var a sql.NullString
		if err := rows.Scan(&a); err != nil {
			return nil, err
		}
		items = append(items, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
