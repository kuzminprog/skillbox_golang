package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	user_app "github.com/kuzminprog/skillbox_golang/tree/main/30"
)

type UserItem struct {
	Name    string          `json:"name"`
	Age     int8            `json:"age"`
	Friends map[string]bool `json:"friends"`
}

type UserDB struct {
	db *DataBase
}

func NewUserDB(db *DataBase) *UserDB {
	return &UserDB{db: db}
}

func (r *UserDB) CreateUser(user user_app.RequestCreate) (string, error) {
	age, err := strconv.Atoi(user.Age)
	if err != nil {
		return "", err
	}

	friends, err := sliceToMap(user.Friends)
	if err != nil {
		return "", err
	}

	userItem := &UserItem{
		Name:    user.Name,
		Age:     int8(age),
		Friends: friends,
	}

	r.db.lastId++
	id := strconv.Itoa(r.db.lastId)

	for friendId := range friends {
		friend, err := r.getUser(friendId)
		if err != nil {
			return "", nil
		}

		friend.Friends[id] = true

		err = r.setUser(friendId, friend)
		if err != nil {
			return "", nil
		}
	}

	err = r.setUser(id, userItem)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("New user id: %v", id), nil
}

func (r *UserDB) MakeFriends(sourceId, targetId string) (string, error) {
	sourceUser, err := r.getUser(sourceId)
	if err != nil {
		return "", err
	}

	targetUser, err := r.getUser(targetId)
	if err != nil {
		return "", err
	}

	sourceUser.Friends[targetId] = true
	targetUser.Friends[sourceId] = true

	err = r.setUser(sourceId, sourceUser)
	if err != nil {
		return "", err
	}

	err = r.setUser(targetId, targetUser)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v и %v теперь друзья", targetId, sourceId), nil
}

func (r *UserDB) DeleteUser(id string) (string, error) {
	user, err := r.getUser(id)
	if err != nil {
		return "", err
	}

	for friendId := range user.Friends {

		friend, err := r.getUser(friendId)
		if err != nil {
			return "", err
		}

		delete(friend.Friends, id)

		err = r.setUser(friendId, friend)
		if err != nil {
			return "", err
		}
	}

	ctx := context.Background()
	err = r.db.rdb.Del(ctx, id).Err()
	if err != nil {
		return "Del", err
	}

	return fmt.Sprintf("%v удален", user.Name), nil
}

func (r *UserDB) GetFriends(id string) ([]string, error) {
	user, err := r.getUser(id)
	if err != nil {
		return nil, err
	}

	friends, err := mapToSlice(user.Friends)
	if err != nil {
		return nil, err
	}

	return friends, nil
}

func (r *UserDB) UpdateAge(id, age string) (string, error) {
	user, err := r.getUser(id)
	if err != nil {
		return "", err
	}

	userAge, err := strconv.Atoi(age)
	if err != nil {
		return "", err
	}

	user.Age = int8(userAge)

	err = r.setUser(id, user)
	if err != nil {
		return "", err
	}

	return "Возраст пользователя успешно обновлён", nil
}

func (r *UserDB) getUser(id string) (*UserItem, error) {
	ctx := context.Background()

	data, err := r.db.rdb.Get(ctx, id).Result()
	if err != nil {
		return nil, err
	}

	var user UserItem
	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (r *UserDB) setUser(id string, user *UserItem) error {
	ctx := context.Background()

	data, err := json.Marshal(&user)
	if err != nil {
		return err
	}

	err = r.db.rdb.Set(ctx, id, data, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
