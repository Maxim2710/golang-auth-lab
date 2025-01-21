package service

import (
	"github.com/Maxim2710/golang-auth-lab/internal/database/repository"
	"github.com/Maxim2710/golang-auth-lab/internal/models"
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
