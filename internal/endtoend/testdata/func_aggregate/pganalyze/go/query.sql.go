// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
)

const percentile = `-- name: Percentile :one
select percentile_disc(0.5) within group (order by authors.name)
from authors
`

func (q *Queries) Percentile(ctx context.Context) (string, error) {
	row := q.db.QueryRowContext(ctx, percentile)
	var percentile_disc string
	err := row.Scan(&percentile_disc)
	return percentile_disc, err
}
