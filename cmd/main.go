package main

import (
	"fmt"
	"net/http"
	"os"
	server "relat/internal/adapter/http"
	"relat/internal/platform/config"
	"relat/internal/platform/logger"
)

func main() {
	log, err := logger.NewLogger()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := config.Bootstrap(); err != nil {
		log.Error("error bootstraping config", "err", err)
		os.Exit(1)
	}

	srv := server.NewServer()
	mux := srv.MountHandlers()

	log.Info(fmt.Sprintf("server listening on %s", config.EnvServerAddr))
	if err := http.ListenAndServe(config.EnvServerAddr, mux); err != nil {
		log.Error("error in server listening", "err", err)
		os.Exit(1)
	}
}
