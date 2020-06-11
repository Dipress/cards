package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httpBroker "github.com/dipress/cards/internal/broker/http"
	"github.com/dipress/cards/internal/broker/http/logger"
	"github.com/dipress/cards/internal/card"
	"github.com/dipress/cards/internal/storage/postgres"
	"github.com/dipress/cards/internal/storage/postgres/schema"
	"github.com/dipress/cards/internal/validation"
	"github.com/mattes/migrate"
	"github.com/pkg/errors"
)

func main() {
	var (
		addr = flag.String("addr", ":8080", "address of http server")
		dsn  = flag.String("dsn", "", "postgres database DSN")
	)

	flag.Parse()

	// Logger initialize.
	logger, err := logger.New(
		logger.WithSensitiveFields([]string{
			"password",
			"token",
		}),
	)

	// Setup database connection.
	logger.Info("connecting to db", nil)
	db, err := sql.Open("postgres", *dsn)
	if err != nil {
		logger.Fatal(fmt.Errorf("failed to create db: %w", err), nil)
	}

	if err := db.Ping(); err != nil {
		logger.Fatal(fmt.Errorf("failed to connect db: %w", err), nil)
	}

	defer db.Close()
	logger.Info("connection to db established", nil)

	// Migrate schema.
	if err := schema.Migrate(db); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Fatal(fmt.Errorf("failed to migrate schema: %w", err), nil)
		}
	}

	// Make a channel for errors.
	errChan := make(chan error)

	// Services
	services := setupServices(db)

	// Setup server.
	srv := setupServer(*addr, logger, services)

	go func() {
		logger.Info(fmt.Sprintf("starting %s server", srv.Addr), nil)
		if err := srv.ListenAndServe(); err != nil {
			errChan <- fmt.Errorf("launch server %s: %w", srv.Addr, err)
		}
	}()

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		logger.Fatal(err, nil)
	case <-osSignals:
		if err := srv.Shutdown(context.TODO()); err != nil {
			logger.Fatal(fmt.Errorf("stop server %s: %w", srv.Addr, err), nil)
		}
	}
}

func setupServer(addr string, logger *logger.Logger, services *httpBroker.Services) *http.Server {
	return httpBroker.NewServer(addr, logger, services)
}

func setupServices(db *sql.DB) *httpBroker.Services {
	// Repositories.
	cardRepo := postgres.NewCardRepository(db)

	// Servives.
	cardService := card.NewService(cardRepo, &validation.Card{})

	services := httpBroker.Services{
		Card: cardService,
	}

	return &services
}
