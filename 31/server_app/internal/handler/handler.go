package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/kuzminprog/skillbox_golang/tree/main/30/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/create", h.CreateUser)
	router.Post("/make_friends", h.MakeFriends)
	router.Delete("/user", h.DeleteUser)
	router.Get("/friends/{user_id}", h.GetFriends)
	router.Put("/{user_id}", h.UpdateAge)

	return router
}
