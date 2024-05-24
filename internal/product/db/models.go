// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Product struct {
	ID          int64
	CreatedAt   pgtype.Timestamp
	Title       string
	Description string
	Img         string
	Size        int32
	SizeUs      pgtype.Int4
	SizeEu      pgtype.Int4
	Price       pgtype.Numeric
	Sex         pgtype.Int4
	Kind        pgtype.Int4
}