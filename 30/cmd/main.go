package main

import (
	user_app "github.com/kuzminprog/skillbox_golang/tree/main/30"
	"github.com/kuzminprog/skillbox_golang/tree/main/30/internal/handler"
	"github.com/kuzminprog/skillbox_golang/tree/main/30/internal/repository"
	"github.com/kuzminprog/skillbox_golang/tree/main/30/internal/service"
	"github.com/rs/zerolog/log"
)

func main() {

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	server := new(user_app.Server)
	go func() {
		if err := server.Run("8080", handler.InitRouters()); err != nil {
			log.Err(err).Msg("Server is not running")
		}
	}()

}
