// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package users

import (
	"database/sql"
)

type User struct {
	ID   int64
	Name string
	Bio  sql.NullString
}
