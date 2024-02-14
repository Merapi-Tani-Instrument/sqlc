// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countAlertReportBy = `-- name: CountAlertReportBy :many
select DATE_TRUNC($1,ts)::text as datetime,coalesce(count,0) as count from 
(
    SELECT  DATE_TRUNC($1,eventdate) as hr ,count(*)
    FROM    alertreport
    where eventdate between $2 and $3
    GROUP BY   1
) AS cnt 
right outer join ( SELECT ts FROM generate_series ( $2, $3, CONCAT('1 ',$1)::interval) AS ts ) as dte
on DATE_TRUNC($1, ts ) = cnt.hr 
order by 1 asc
`

type CountAlertReportByParams struct {
	DateTrunc   string
	Eventdate   pgtype.Date
	Eventdate_2 pgtype.Date
}

type CountAlertReportByRow struct {
	Datetime string
	Count    int64
}

func (q *Queries) CountAlertReportBy(ctx context.Context, arg CountAlertReportByParams) ([]CountAlertReportByRow, error) {
	rows, err := q.db.Query(ctx, countAlertReportBy, arg.DateTrunc, arg.Eventdate, arg.Eventdate_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CountAlertReportByRow
	for rows.Next() {
		var i CountAlertReportByRow
		if err := rows.Scan(&i.Datetime, &i.Count); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
