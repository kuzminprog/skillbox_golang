package service

import (
	user_app "github.com/kuzminprog/skillbox_golang/tree/main/30"
	"github.com/kuzminprog/skillbox_golang/tree/main/30/internal/repository"
)

type UserService struct {
	repo repository.User
}

// NewUserService sends information to the repository and returns a response
func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

// CreateUser sends information to the repository and returns a response
func (s *UserService) CreateUser(user user_app.RequestCreate) (string, error) {
	return s.repo.CreateUser(user)
}

// MakeFriends sends information to the repository and returns a response
func (s *UserService) MakeFriends(sourceId, targetId string) (string, error) {
	return s.repo.MakeFriends(sourceId, targetId)
}

// DeleteUser sends information to the repository and returns a response
func (s *UserService) DeleteUser(id string) (string, error) {
	return s.repo.DeleteUser(id)
}

// GetFriends sends information to the repository and returns a response
func (s *UserService) GetFriends(id string) ([]string, error) {
	return s.repo.GetFriends(id)
}

// UpdateAge sends information to the repository and returns a response
func (s *UserService) UpdateAge(id, age string) (string, error) {
	return s.repo.UpdateAge(id, age)
}
