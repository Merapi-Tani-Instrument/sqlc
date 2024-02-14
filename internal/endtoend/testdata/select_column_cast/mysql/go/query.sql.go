// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
)

const selectColumnCast = `-- name: SelectColumnCast :many
SELECT CAST(bar AS UNSIGNED) FROM foo
`

func (q *Queries) SelectColumnCast(ctx context.Context) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, selectColumnCast)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var bar int64
		if err := rows.Scan(&bar); err != nil {
			return nil, err
		}
		items = append(items, bar)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}