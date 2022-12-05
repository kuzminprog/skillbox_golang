package main

import (
	"flag"
	"net/http"
	"os"
	proxyapp "proxy"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	port := flag.String("port", "9000", "string")
	flag.Parse()

	initLogging()

	if err := initConfig(); err != nil {
		log.Err(err).Msg("Configuration file not loaded")
	}
	proxy := proxyapp.NewProxy()

	h := proxy
	http.Handle("/", h)

	log.Info().Msg("Starting proxy")
	server := &http.Server{
		Addr:    ":" + *port,
		Handler: h,
	}
	log.Info().Msg("Proxy started")
	log.Err(server.ListenAndServe()).Msg("Proxy is not running")
}

// initConfig() initializes configuration files
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config_proxy")
	return viper.ReadInConfig()
}

// initLogging() sets the logging settings
func initLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "02.01.2006 15:04:05"}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Logger()
}
