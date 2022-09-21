package user

import (
	"log"
)

type User struct {
	name    string
	age     string
	friends map[string]bool
}

// NewUser - сreates an unfilled new user.
// Returns the user.
func NewUser() *User {
	user := new(User)
	user.friends = make(map[string]bool)
	return user
}

// AddFriend - аdds an id to a user's friends list.
// If the user is already a friend, leaves everything as it is.
func (u *User) AddFriend(id string) {
	_, ok := u.friends[id]
	if ok {
		log.Printf("User %v already has friend %v\n", u.name, id)
		return
	}
	u.friends[id] = true
}

// DeleteFriend - deletes the id from the user's friends list.
// If the user is not on the list, leaves everything as it is.
func (u *User) DeleteFriend(id string) {
	_, ok := u.friends[id]
	if !ok {
		log.Printf("User %v does not have friend %v\n", u.name, id)
		return
	}
	delete(u.friends, id)
}

// GetName - returns the user name.
func (u *User) GetName() string {
	return u.name
}

// GetAge - returns the user's age.
func (u *User) GetAge() string {
	return u.age
}

// GetFriends - returns a list of friends.
// From a map makes a slice.
// Returns a slice of friends.
func (u *User) GetFriends() []string {
	var friends []string
	for id := range u.friends {
		friends = append(friends, id)
	}
	return friends
}

// SetName - assigns the user a name.
func (u *User) SetName(name string) {
	u.name = name
}

// SetAge - assigns an age to the user.
func (u *User) SetAge(age string) {
	u.age = age
}

// SetFriends - assigns a list of friends to the user.
// From a slice makes a map.
func (u *User) SetFriends(friends []string) {
	u.friends = make(map[string]bool)
	for _, friendId := range friends {
		u.AddFriend(friendId)
	}
}

// SetAll - assigns user name, age and list of friends
func (u *User) SetAll(name string, age string, friends []string) {
	u.name = name
	u.age = age
	u.SetFriends(friends)
}
