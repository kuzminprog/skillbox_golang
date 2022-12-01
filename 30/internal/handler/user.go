package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

}

func (h *Handler) MakeFriends(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) GetFriends(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) UpdateAge(w http.ResponseWriter, r *http.Request) {

}
