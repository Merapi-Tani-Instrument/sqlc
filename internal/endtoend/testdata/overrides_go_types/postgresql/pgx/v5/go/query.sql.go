// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package override

import (
	"context"

	uuid "github.com/gofrs/uuid"
)

const loadFoo = `-- name: LoadFoo :many
SELECT id, other_id, age, balance, bio, about FROM foo WHERE id = $1
`

func (q *Queries) LoadFoo(ctx context.Context, id uuid.UUID) ([]Foo, error) {
	rows, err := q.db.Query(ctx, loadFoo, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Foo
	for rows.Next() {
		var i Foo
		if err := rows.Scan(
			&i.ID,
			&i.OtherID,
			&i.Age,
			&i.Balance,
			&i.Bio,
			&i.About,
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