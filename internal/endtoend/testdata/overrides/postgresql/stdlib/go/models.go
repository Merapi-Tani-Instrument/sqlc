// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package override

import (
	"github.com/lib/pq"
	"github.com/sqlc-dev/sqlc-testdata/pkg"
)

type Foo struct {
	Other   string
	Total   int64
	Tags    []string
	ByteSeq []byte
	Retyped pkg.CustomType
	Langs   pq.StringArray
}