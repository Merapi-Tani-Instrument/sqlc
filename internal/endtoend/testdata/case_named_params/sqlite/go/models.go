// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package querytest

import (
	"database/sql"
)

type Author struct {
	ID       int64
	Username interface{}
	Email    interface{}
	Name     string
	Bio      sql.NullString
}
