package main

import (
	"log"
	"net/http"

	"github.com/sarch-3/stream-to-display/internal/api"
	"github.com/sarch-3/stream-to-display/internal/config"
	"github.com/sarch-3/stream-to-display/internal/service"
)

func main() {
	cfg := config.Load()
	playerService := service.NewPlayerService(cfg.SocketPath)
	handler := api.NewHandler(playerService)

	log.Printf("server started on %s", cfg.Addr())
	if err := http.ListenAndServe(cfg.Addr(), handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
