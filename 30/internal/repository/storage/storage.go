package storage

import (
	"30/internal/repository/user"
	"fmt"
	"log"
	"strconv"
)

type Users map[string]*user.User

// getId - function for generating identifiers.
var (
	getId = idWrapper()
)

// idWrapper - returns a function for generating identifiers.
func idWrapper() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

// Add - adds a new user to the Users type.
// Gets name, age and slice of friends
// Makes friends of users specified in the slice.
// Returns the id of the new user.
func (u Users) Add(name string, age string, friends []string) string {
	id := strconv.Itoa(getId())

	friends = u.fixFriendsArray(friends)
	u[id] = user.NewUser()
	u[id].SetAll(name, age, friends)

	for _, friendId := range friends {

		if friendId == id {
			log.Printf("User %v cannot be his own friend\n", id)
			continue
		}

		if !u.HasId(id) {
			log.Printf("User %v is not added as a friend because he does not exist\n", id)
			continue
		}

		u[friendId].AddFriend(id)
	}
	return id
}

// Delete - deletes the user with the id from the Users type.
// Deletes this user from other users
func (u Users) Delete(id string) {
	friends := u[id].GetFriends()

	for _, friendId := range friends {
		u[friendId].DeleteFriend(id)
	}

	delete(u, id)
}

// HasId - checks if there is a user with the id in users
// If it exists it returns true
func (u Users) HasId(id string) bool {
	_, ok := u[id]
	if !ok {
		message := fmt.Sprintf("The user %v does not exist\n", id)
		log.Print(message)
		return false
	}
	return true
}

// fixFriendsArray - Clears the slice from users that do not exist in Users
// Returns a new slice
func (u Users) fixFriendsArray(friends []string) []string {
	var existingFriends []string
	for _, id := range friends {
		if u.HasId(id) {
			existingFriends = append(existingFriends, id)
		}
	}
	return existingFriends
}
