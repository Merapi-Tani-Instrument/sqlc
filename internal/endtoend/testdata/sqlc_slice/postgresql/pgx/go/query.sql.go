// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
)

const funcParamIdent = `-- name: FuncParamIdent :many
SELECT name FROM foo WHERE name = $1
  AND id IN ($2)
`

type FuncParamIdentParams struct {
	Slug       string
	Favourites []int32
}

func (q *Queries) FuncParamIdent(ctx context.Context, arg FuncParamIdentParams) ([]string, error) {
	rows, err := q.db.Query(ctx, funcParamIdent, arg.Slug, arg.Favourites)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const funcParamString = `-- name: FuncParamString :many
SELECT name FROM foo WHERE name = $1
  AND id IN ($2)
`

type FuncParamStringParams struct {
	Slug       string
	Favourites []int32
}

func (q *Queries) FuncParamString(ctx context.Context, arg FuncParamStringParams) ([]string, error) {
	rows, err := q.db.Query(ctx, funcParamString, arg.Slug, arg.Favourites)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
