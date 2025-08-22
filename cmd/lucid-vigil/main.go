package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lucid-vigil/pkg/actions"
	"github.com/lucid-vigil/pkg/api"
	"github.com/lucid-vigil/pkg/config"
	"github.com/lucid-vigil/pkg/events"
	"github.com/lucid-vigil/pkg/logger"
	monitors "github.com/lucid-vigil/pkg/monitors/detector"
	"github.com/lucid-vigil/pkg/monitors/scheduler"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Load configuration first
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	// Initialize logger based on config
	logger.InitLogger(cfg.LogLevel)

	log.Info().Msg("Lucid Vigil application starting...")
	log.Info().Msgf("Configuration loaded: LogLevel=%s, APIPort=%s, ActionsEnabled=%t",
		cfg.LogLevel, cfg.APIPort, cfg.Actions.Enabled)

	// Create action dispatcher
	actionDispatcher := actions.NewActionDispatcher(cfg.Actions.Enabled)

	// Create a context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())

	// Set up a channel to listen for OS signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to handle graceful shutdown
	go func() {
		sig := <-sigChan
		log.Info().Msgf("Received signal: %s. Shutting down gracefully...", sig)
		cancel()
	}()

	// Start API server in a goroutine
	go api.StartAPIServer(cfg.APIPort)

	// Initialize and start the scheduler
	sched := scheduler.NewScheduler(cfg)

	// Register monitors with action dispatcher support
	registerMonitorWithActions := func(monitor scheduler.Monitor, monitorConfig *config.MonitorConfig) {
		// Set up action integration if monitor supports it
		if baseMonitor, ok := monitor.(interface {
			SetActionDispatcher(*actions.ActionDispatcher)
			SetConfiguredActions([]string)
		}); ok {
			baseMonitor.SetActionDispatcher(actionDispatcher)
			if monitorConfig != nil {
				baseMonitor.SetConfiguredActions(monitorConfig.Actions)
			}
		}
		sched.RegisterMonitor(monitor)
	}

	// Register all monitors
	registerMonitorWithActions(monitors.NewAnalyzerMonitor(zerolog.Logger{}, events.NewEventBus(zerolog.Logger{}, 1024)))
	registerMonitorWithActions(monitors.NewDetectorMonitor(zerolog.Logger{}, events.NewEventBus(zerolog.Logger{}, 1024)))
	registerMonitorWithActions(monitors.NewSchedulerMonitor(zerolog.Logger{}, events.NewEventBus(zerolog.Logger{}, 1024)))
	registerMonitorWithActions(monitors.NewSentinel(zerolog.Logger{}, events.NewEventBus(zerolog.Logger{}, 1024)))
	registerMonitorWithActions(monitors.NewSentryMonitor(zerolog.Logger{}, events.NewEventBus(zerolog.Logger{}, 1024)))
	registerMonitorWithActions(monitors.NewScribeMonitor(zerolog.Logger{}, events.NewEventBus(zerolog.Logger{}, 1024)))

	// Start all configured monitors
	sched.Start(ctx)

	// Wait for shutdown signal
	<-ctx.Done()

	log.Info().Msg("Lucid Vigil application stopped.")
	time.Sleep(1 * time.Second) // Give some time for cleanup
}
