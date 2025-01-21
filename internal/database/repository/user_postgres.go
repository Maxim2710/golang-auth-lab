package repository

import (
	"github.com/Maxim2710/golang-auth-lab/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserById(id int) (*models.User, error) {
	var user models.User
	query := `SELECT username, email, created_at FROM users WHERE id = $1`
	err := r.db.Get(&user, query, id)
	return &user, err
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE email = $1`
	err := r.db.Get(&user, query, email)
	return &user, err
}

func (r *UserRepository) UpdatePasswordByEmail(email string, newPassword string) error {
	query := `UPDATE users SET password = $1 WHERE email = $2`
	_, err := r.db.Exec(query, newPassword, email)
	return err
}
