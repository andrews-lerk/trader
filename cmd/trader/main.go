package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"trader/internal/config"
	internal "trader/internal/controller/http"
	"trader/internal/logging"
)

func main() {
	// load config and setup logging
	cfg := config.MustLoad()
	logger := logging.SetupLogging(&cfg.Logger)

	logger.Info("Setup logging")

	// Init router
	router := internal.NewRouter(logger)

	logger.Info(
		"Starting server",
		slog.String("Host", cfg.Api.Host),
		slog.Int("Port", cfg.Api.Port),
	)
	srv := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", cfg.Api.Host, cfg.Api.Port),
		Handler:      router,
		ReadTimeout:  cfg.Api.Timeout,
		WriteTimeout: cfg.Api.Timeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Error("Failed to start server")
	}
	logger.Warn("Shutdown")
}
