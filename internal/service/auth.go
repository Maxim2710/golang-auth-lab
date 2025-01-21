package service

import (
	"errors"
	"github.com/Maxim2710/golang-auth-lab/internal/database/repository"
	"github.com/Maxim2710/golang-auth-lab/internal/models"
	"github.com/Maxim2710/golang-auth-lab/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(username, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}

func (s *AuthService) AuthenticateUser(email string, password string) (token string, err error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Invalid password")
	}

	token, err = utils.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
