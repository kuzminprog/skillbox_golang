package server

import (
	"30/internal/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// StartServer - starts the server
// Routes the request
// Returns an error in the event of a failure
func StartServer() error {
	fmt.Println("Starting Server...")

	router := chi.NewRouter()

	router.Post("/create", handler.PostCreateHandler)
	router.Post("/make_friends", handler.PostFriendsHandler)
	router.Delete("/user", handler.DeleteHandler)
	router.Get("/friends/{user_id}", handler.GetHandler)
	router.Put("/{user_id}", handler.PutHandler)

	err := http.ListenAndServe("localhost:8080", router)
	return err
}
