package repository

import user_app "github.com/kuzminprog/skillbox_golang/tree/main/30"

type UserList interface {
	CreateUser(user user_app.RequestCreate) (string, error)
	MakeFriends(sourceId string, targetId string) (string, error)
	DeleteUser(id string) (string, error)
	GetFriends(id string) ([]string, error)
	UpdateAge(id string, age string) (string, error)
}

type Repository struct {
	UserList
}

func NewRepository(db *DataBase) *Repository {
	return &Repository{
		// UserList: NewUserListDB(db),
	}
}
