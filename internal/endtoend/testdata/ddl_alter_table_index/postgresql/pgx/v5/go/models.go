// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Measurement struct {
	CityID    int32
	Logdate   pgtype.Date
	Peaktemp  pgtype.Int4
	Unitsales pgtype.Int4
}

type MeasurementY2006m02 struct {
	CityID    int32
	Logdate   pgtype.Date
	Peaktemp  pgtype.Int4
	Unitsales pgtype.Int4
}