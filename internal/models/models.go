// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package models

import (
	"time"

	"github.com/jackc/pgtype"
)

type Hub struct {
	ID   int64       `db:"id" json:"id"`
	Name pgtype.Text `db:"name" json:"name"`
	Bio  pgtype.Text `db:"bio" json:"bio"`
}

type User struct {
	ID        pgtype.Text `db:"id" json:"id"`
	Name      pgtype.Text `db:"name" json:"name"`
	Bio       pgtype.Text `db:"bio" json:"bio"`
	UpdatedAt time.Time   `db:"updated_at" json:"updated_at"`
}
