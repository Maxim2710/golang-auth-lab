package models

import "time"

type User struct {
	ID        int       `db:"id" json:"id"`
	Username  string    `db:"username" json:"username" binding:"required"`
	Email     string    `db:"email" json:"email" binding:"required"`
	Password  string    `db:"password,omitempty" json:"password" binding:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
