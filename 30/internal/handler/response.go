package handler

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

// newMessageResponse sends a message.
// It gets http status and message.
// Writes the response to http.ResponseWriter
func newMessageResponse(w http.ResponseWriter, statusCode int, response string) {
	w.WriteHeader(statusCode)

	log.Info().Msg(fmt.Sprintf("Sending: %v", response))
	w.Write([]byte(response))
}
