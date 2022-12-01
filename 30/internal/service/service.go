package service

import (
	user_app "github.com/kuzminprog/skillbox_golang/tree/main/30"
	"github.com/kuzminprog/skillbox_golang/tree/main/30/internal/repository"
)

type User interface {
	CreateUser(user user_app.RequestCreate) (string, error)
	MakeFriends(sourceId string, targetId string) (string, error)
	DeleteUser(id string) (string, error)
	GetFriends(id string) ([]string, error)
	UpdateAge(id string, age string) (string, error)
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.UserList),
	}
}
