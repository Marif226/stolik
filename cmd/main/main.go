package main

import (
	"net"
	"net/http"
	"time"
	"github.com/julienschmidt/httprouter"
	"github.com/marif226/stolik/internal/user"
	"github.com/marif226/stolik/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create roter")

	router := httprouter.New()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)
}

// Start starts the application
func start(router *httprouter.Router) {
	logger := logging.GetLogger()

	logger.Info("start application")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	logger.Info("server is listening port 0.0.0.0:1234")

	logger.Fatal(server.Serve(listener))
}