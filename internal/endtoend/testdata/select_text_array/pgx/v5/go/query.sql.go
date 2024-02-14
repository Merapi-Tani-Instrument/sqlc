// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
)

const selectTextArray = `-- name: SelectTextArray :many
SELECT $1::TEXT[]
`

func (q *Queries) SelectTextArray(ctx context.Context, dollar_1 []string) ([][]string, error) {
	rows, err := q.db.Query(ctx, selectTextArray, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items [][]string
	for rows.Next() {
		var column_1 []string
		if err := rows.Scan(&column_1); err != nil {
			return nil, err
		}
		items = append(items, column_1)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}