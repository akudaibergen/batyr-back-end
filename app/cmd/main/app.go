package main

import (
	"app/initializer"
	"app/internal/objects"
	objectDB "app/internal/objects/db"
	"app/pkg/handlers/metric"
	"app/pkg/logging"
	"app/pkg/shutdown"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("Logger initialized")

	logger.Println("Initializing router")
	router := httprouter.New()

	logger.Println("Initializing health metric")
	metricHandler := metric.Handler{Logger: logger}
	metricHandler.Register(router)

	logger.Println("Initializing config")
	projectDir, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}

	config, err := initializer.LoadConfig(projectDir)
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	postgresDB := initializer.ConnectDB(&config)

	objectStorage := objectDB.NewStorage(postgresDB, logger)
	objectService := objects.NewService(objectStorage, logger)
	objectHandler := objects.Handler{
		Logger:  logger,
		Service: objectService,
	}
	objectHandler.Register(router)

	// Middleware for adding CORS headers
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Allow access from any origin
			w.Header().Set("Access-Control-Allow-Origin", "*")
			// Allow only GET, POST, OPTIONS, PUT, DELETE methods
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
			// Allow only specified headers
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			// Respond to OPTIONS preflight requests
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Continue with the next handler
			next.ServeHTTP(w, r)
		})
	}

	// Start the server
	start(router, logger, config, corsMiddleware)
}

func start(router http.Handler, logger logging.Logger, cfg initializer.Config, middleware func(http.Handler) http.Handler) {
	var listener net.Listener
	var err error

	if cfg.ListenType == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		socketPath := path.Join(appDir, "app.sock")
		logger.Infof("Socket path: %s", socketPath)

		logger.Info("Creating and listening on Unix socket")
		listener, err = net.Listen("unix", socketPath)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		logger.Infof("Binding application to host: %s and port: %s", cfg.ListenBindIp, cfg.ListenPort)

		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.ListenBindIp, cfg.ListenPort))
		if err != nil {
			logger.Fatal(err)
		}
	}

	// Create a new server with CORS middleware
	server := &http.Server{
		Handler:      middleware(router),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Graceful shutdown
	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM},
		server)

	logger.Println("Application initialized and started")

	// Start serving requests
	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("Server shutdown")
		default:
			logger.Fatal(err)
		}
	}
}
