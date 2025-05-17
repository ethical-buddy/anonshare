package core

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"anonshare/config"
	"anonshare/core/logger"
)

type App struct {
	ctx    context.Context
	cancel context.CancelFunc
	Config *config.Config
}

// NewApp initializes the app with context and config
func NewApp() *App {
	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config:", err)
		os.Exit(1)
	}

	logger.InitLogger(cfg.Debug)

	return &App{
		ctx:    ctx,
		cancel: cancel,
		Config: cfg,
	}
}

// Run starts the main loop and listens for shutdown
func (a *App) Run() {
	logger.Info("AnonShare starting up...")

	// TODO: Init other modules (node identity, metadata, etc.)

	a.waitForShutdown()
}

// waitForShutdown blocks until interrupt signal is received
func (a *App) waitForShutdown() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
	logger.Warn("Interrupt received. Shutting down...")
	a.Shutdown()
}

// Shutdown gracefully terminates app
func (a *App) Shutdown() {
	a.cancel()
	time.Sleep(500 * time.Millisecond) // Give modules time to cleanup
	logger.Info("AnonShare shut down cleanly.")
}

