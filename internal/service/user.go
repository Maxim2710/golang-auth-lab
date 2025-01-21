package service

import (
	"errors"
	"github.com/Maxim2710/golang-auth-lab/internal/database/repository"
	"github.com/Maxim2710/golang-auth-lab/internal/models"
	"github.com/Maxim2710/golang-auth-lab/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserById(id int) (*models.User, error) {
	return s.repo.GetUserById(id)
}

func (s *UserService) UpdatePassword(token string, password string, password2 string) error {
	email, err := utils.ParseToken(token)
	if err != nil {
		return errors.New("invalid token")
	}

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("password incorrect. Please try again")
	}

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password2), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("could not hash password")
	}

	if err := s.repo.UpdatePasswordByEmail(email, string(hashed_password)); err != nil {
		return errors.New("failed to update password")
	}

	return nil
}
