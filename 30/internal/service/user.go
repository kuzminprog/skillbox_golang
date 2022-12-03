package service

import (
	user_app "github.com/kuzminprog/skillbox_golang/tree/main/30"
	"github.com/kuzminprog/skillbox_golang/tree/main/30/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user user_app.RequestCreate) (string, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) MakeFriends(sourceId, targetId string) (string, error) {
	return s.repo.MakeFriends(sourceId, targetId)
}
func (s *UserService) DeleteUser(id string) (string, error) {
	return s.repo.DeleteUser(id)
}

func (s *UserService) GetFriends(id string) ([]string, error) {
	return s.repo.GetFriends(id)
}

func (s *UserService) UpdateAge(id, age string) (string, error) {
	return s.repo.UpdateAge(id, age)
}
