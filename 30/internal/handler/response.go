package handler

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func newMessageResponse(w http.ResponseWriter, statusCode int, response string) {
	w.WriteHeader(statusCode)

	log.Info().Msg(fmt.Sprintf("Sending: %v", response))
	w.Write([]byte(response))
}
