package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	user_app "github.com/kuzminprog/skillbox_golang/tree/main/30"
	"github.com/kuzminprog/skillbox_golang/tree/main/30/internal/handler"
	"github.com/kuzminprog/skillbox_golang/tree/main/30/internal/repository"
	"github.com/kuzminprog/skillbox_golang/tree/main/30/internal/service"
	"github.com/rs/zerolog/log"
)

func main() {

	db, err := repository.NewDataBase("----")
	if err != nil {
		log.Error().Msg("wer")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	server := new(user_app.Server)
	go func() {
		if err := server.Run("8080", handler.InitRouters()); err != nil {
			log.Err(err).Msg("Server is not running")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		log.Err(err).Msg("Server did not shut down correctly")
	}

}
