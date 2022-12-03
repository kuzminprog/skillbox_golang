package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	user_app "github.com/kuzminprog/skillbox_golang/tree/main/30"
	"github.com/rs/zerolog/log"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user user_app.RequestCreate
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("POST: New user %v", string(content)))

	err = json.Unmarshal(content, &user)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.User.CreateUser(user)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	message := fmt.Sprintf("The user %v was created with the id %v", user.Name, id)
	newMessageResponse(w, http.StatusCreated, message)
}

func (h *Handler) MakeFriends(w http.ResponseWriter, r *http.Request) {
	var friends user_app.RequestMakeFriend
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("POST: Make friends %v", string(content)))

	err = json.Unmarshal(content, &friends)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	message, err := h.services.User.MakeFriends(friends.SourceID, friends.TargetID)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newMessageResponse(w, http.StatusOK, message)

}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user user_app.RequestDeleteUser
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("DELETE: %v", string(content)))

	err = json.Unmarshal(content, &user)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	message, err := h.services.User.DeleteUser(user.TargetID)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newMessageResponse(w, http.StatusOK, message)
}

func (h *Handler) GetFriends(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "user_id")

	log.Info().Msg(fmt.Sprintf("GET: friends %v", string(id)))

	friends, err := h.services.User.GetFriends(id)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	message := strings.Join(friends, ";")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func (h *Handler) UpdateAge(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "user_id")

	var age user_app.RequestAge
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal(content, &age)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	message, err := h.services.User.UpdateAge(id, age.NewAge)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newMessageResponse(w, http.StatusOK, message)
}
