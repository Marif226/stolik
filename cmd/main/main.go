package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/marif226/stolik/internal/config"
	"github.com/marif226/stolik/internal/user"
	"github.com/marif226/stolik/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create roter")

	router := httprouter.New()

	cfg := config.GetConfig()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

// Start starts the application
func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("detect app path")

		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")
		path.Join(appDir, "app.sock")
		logger.Debugf("socket path: %s", socketPath)

		logger.Info("create unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket %s", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	}

	if listenErr != nil {
		log.Fatal(listenErr)
	}
	
	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	logger.Infof("server is listening port %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)

	logger.Fatal(server.Serve(listener))
}