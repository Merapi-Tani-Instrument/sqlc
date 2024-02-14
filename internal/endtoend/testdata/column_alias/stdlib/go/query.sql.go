// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
	"time"
)

const getUsers = `-- name: GetUsers :many
SELECT 
    users.id,
    users.fname,
    users.lname,
    users.email,
    users.created_at,
    rank_email,
    rank_fname,
    rank_lname,
    similarity
FROM 
    users, 
    to_tsvector(users.email || users.fname || users.lname) document,
    to_tsquery($1::TEXT) query,
    NULLIF(ts_rank(to_tsvector(users.email), query), 0) rank_email,
    NULLIF(ts_rank(to_tsvector(users.fname), query), 0) rank_fname,
    NULLIF(ts_rank(to_tsvector(users.lname), query), 0) rank_lname,
    SIMILARITY($1::TEXT, users.email || users.fname || users.lname) similarity
WHERE query @@ document OR similarity > 0
ORDER BY rank_email, rank_lname, rank_fname, similarity DESC NULLS LAST
`

type GetUsersRow struct {
	ID         int32
	Fname      string
	Lname      string
	Email      string
	CreatedAt  time.Time
	RankEmail  sql.NullFloat64
	RankFname  sql.NullFloat64
	RankLname  sql.NullFloat64
	Similarity sql.NullFloat64
}

func (q *Queries) GetUsers(ctx context.Context, searchTerm string) ([]GetUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getUsers, searchTerm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersRow
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Fname,
			&i.Lname,
			&i.Email,
			&i.CreatedAt,
			&i.RankEmail,
			&i.RankFname,
			&i.RankLname,
			&i.Similarity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
