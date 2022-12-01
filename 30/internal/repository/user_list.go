package repository

import user_app "github.com/kuzminprog/skillbox_golang/tree/main/30"

type User struct {
	name    string
	age     int8
	friends map[int]bool
}

type Users map[int]*User

type UserListDB struct {
	db *DataBase
}

func NewUserListDB(db *DataBase) *UserListDB {
	return &UserListDB{db: db}
}

func (r *UserListDB) CreateUser(user user_app.RequestCreate) (string, error)
func (r *UserListDB) MakeFriends(sourceId string, targetId string) (string, error)
func (r *UserListDB) DeleteUser(id string) (string, error)
func (r *UserListDB) GetFriends(id string) ([]string, error)
func (r *UserListDB) UpdateAge(id string, age string) (string, error)
