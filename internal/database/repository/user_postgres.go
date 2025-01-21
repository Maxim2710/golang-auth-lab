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
