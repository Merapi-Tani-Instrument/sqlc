// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
)

const getDate = `-- name: GetDate :one
SELECT  from CURRENT_DATE
`

type GetDateRow struct {
}

func (q *Queries) GetDate(ctx context.Context) (GetDateRow, error) {
	row := q.db.QueryRow(ctx, getDate)
	var i GetDateRow
	err := row.Scan()
	return i, err
}
