// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Bar struct {
	C pgtype.Text
	D pgtype.Text
}

type Foo struct {
	A pgtype.Text
	B pgtype.Text
}
