package service

import (
	user_app "github.com/kuzminprog/skillbox_golang/tree/main/30"
	"github.com/kuzminprog/skillbox_golang/tree/main/30/internal/repository"
)

type UserService struct {
	repo repository.UserList
}

func NewUserService(repo repository.UserList) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user user_app.RequestCreate) (string, error)
func (s *UserService) MakeFriends(sourceId string, targetId string) (string, error)
func (s *UserService) DeleteUser(id string) (string, error)
func (s *UserService) GetFriends(id string) ([]string, error)
func (s *UserService) UpdateAge(id string, age string) (string, error)
