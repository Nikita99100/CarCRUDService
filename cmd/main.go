package main

import (
	"flag"
	"github.com/Nikita99100/CarCRUDService/config"
	"github.com/Nikita99100/CarCRUDService/handler"
	"github.com/Nikita99100/CarCRUDService/pkg/os"
	"github.com/Nikita99100/CarCRUDService/server"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

const (
	defaultConfigPath     = "deployment/configs.toml"
	serverShutdownTimeout = 30 * time.Second
)

func main() {
	// Parse flags
	configPath := flag.String("config", defaultConfigPath, "configuration file path")
	flag.Parse()

	// Parse and set configs
	cfg, err := config.Parse(*configPath)
	if err != nil {
		log.Fatalf("failed to parse the config file: %v", err)
	}

	// Create a handler object
	hdlr := handler.Handler{}

	// Create a rest server gateway and handle http requests
	router := echo.New()
	rest := server.Rest{
		Router:  router,
		Handler: &hdlr,
	}
	rest.Route()

	// Start an http server and remember to shut it down
	go server.Start(router, cfg.Port)
	defer server.Stop(router, serverShutdownTimeout)

	<-os.NotifyAboutExit()
}
