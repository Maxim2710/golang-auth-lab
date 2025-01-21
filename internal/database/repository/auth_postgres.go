package repository

import (
	"github.com/Maxim2710/golang-auth-lab/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users VALUES ($1, $2, $3) RETURNING id, created_at`
	return r.db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
}

func (r *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE email = $1`
	err := r.db.Get(&user, query, email)
	return &user, err
}
