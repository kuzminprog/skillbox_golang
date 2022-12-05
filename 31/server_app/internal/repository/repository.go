package repository

import user_app "github.com/kuzminprog/skillbox_golang/tree/main/30"

type User interface {
	CreateUser(user user_app.RequestCreate) (string, error)
	MakeFriends(sourceId, targetId string) (string, error)
	DeleteUser(id string) (string, error)
	GetFriends(id string) ([]string, error)
	UpdateAge(id, age string) (string, error)
	getUser(id string) (*UserItem, error)
	setUser(id string, user *UserItem) error
}

type Repository struct {
	User
}

func NewRepository(db *DataBase) *Repository {
	return &Repository{
		User: NewUserDB(db),
	}
}

// sliceToMap converts data from a slice into a map.
// Returns the map
func sliceToMap(slice []string) (map[string]bool, error) {
	dict := make(map[string]bool, len(slice))

	for _, item := range slice {
		dict[item] = true
	}

	return dict, nil
}

// mapToSlice converts data from the map into a slice.
// Returns a slice
func mapToSlice(dict map[string]bool) ([]string, error) {
	var slice []string

	for keys := range dict {
		slice = append(slice, keys)
	}

	return slice, nil
}
