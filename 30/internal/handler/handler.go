package handler

import (
	"30/internal/repository/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var (
	users = make(storage.Users)
)

// PostCreateHandler - processes for POST request.
// Gets name of age and list of friends from the request.
// Creates a user.
// If successful the response is 201.
// Response is sent with the user id.
func PostCreateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: POST /create")

	var request RequestCreate
	err := getDataFromRequest(&request, r)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	id := users.Add(request.Name, request.Age, request.Friends)
	responseMessage := fmt.Sprintf("The user %v with ID %v was created\n", request.Name, id)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(responseMessage))
	log.Println("Response: 201 Created")
}

// PostFriendsHandler - processes for POST request.
// Gets two id's from the request. Makes these id's friends.
// If successful the response is 200.
func PostFriendsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: POST /make_friends")

	var request RequestMakeFriend
	err := getDataFromRequest(&request, r)

	w.WriteHeader(http.StatusInternalServerError)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	if request.SourceID == request.TargetID {
		message := "Two users are the same person"
		log.Println(message)
		w.Write([]byte(message))
		return
	}

	if !users.HasId(request.SourceID) && !users.HasId(request.TargetID) {
		w.Write([]byte("No user found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	users[request.SourceID].AddFriend(request.TargetID)
	users[request.TargetID].AddFriend(request.SourceID)

	responseMessage := fmt.Sprintf("%v and %v are now friends\n",
		request.SourceID, request.TargetID)

	w.Write([]byte(responseMessage))
	log.Println("Response: 200 Ok")
}

// DeleteHandler - processes for DELETE request.
// Gets two id's from the request. Deletes a user.
// Handles Connection: close.
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: DELETE /user")

	if r.Header.Get("Connection") == "close" {
		w.Header().Add("Connection", "close")
		log.Println("Connection close")
		defer r.Body.Close()
	}

	var request RequestDeleteUser
	err := getDataFromRequest(&request, r)

	w.WriteHeader(http.StatusInternalServerError)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	if !users.HasId(request.TargetID) {
		w.Write([]byte("No user found"))
		return
	}

	name := users[request.TargetID].GetName()
	users.Delete(request.TargetID)

	responseMessage := fmt.Sprintf("User %v deleted", name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseMessage))
	log.Println("Response: 200 Ok")
}

// GetHandler - processes for GET {user_id} request.
// Gets the user id from the URL.
// Sends a response with a list of friends.
// If successful the response is 200.
func GetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: GET /friends/user_id")
	id := chi.URLParam(r, "user_id")

	if !users.HasId(id) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No user found"))
		return
	}

	friends := users[id].GetFriends()

	var responseMessage string
	for _, friend_id := range friends {
		responseMessage += users[friend_id].GetName()
		responseMessage += "\n"
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseMessage))
	log.Println("Response: 200 Ok")
}

// PutHandler - processes for PUT {user_id} request.
// Gets the user id from the URL.
// Gets two age from the request.
// Updates the user's age.
// If successful the response is 200.
func PutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request: PUT /user_id")

	var request RequestUserId
	err := getDataFromRequest(&request, r)

	w.WriteHeader(http.StatusInternalServerError)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	id := chi.URLParam(r, "user_id")
	if !users.HasId(id) {
		w.Write([]byte("No user found"))
		return
	}

	users[id].SetAge(request.NewAge)
	log.Printf("user %v is set to a new age\n", id)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Age has been successfully updated"))
	log.Println("Response: 200 Ok")
}
