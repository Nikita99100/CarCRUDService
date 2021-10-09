package main

import (
	"github.com/Nikita99100/CarCRUDService/handler"
	"github.com/Nikita99100/CarCRUDService/server"
	"github.com/labstack/echo/v4"
	"time"
)

const (
	serverShutdownTimeout = 30 * time.Second
	port                  = "80"
)

func main() {
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
	server.Start(router, port)
	defer server.Stop(router, serverShutdownTimeout)
}
