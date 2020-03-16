package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	httpBroker "github.com/dipress/cards/internal/broker/http"
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

	// Setup database connection.
	db, err := sql.Open("postgres", *dsn)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer db.Close()

	// Migrate schema.
	if err := schema.Migrate(db); err != nil {
		if errors.Cause(err) != migrate.ErrNoChange {
			log.Fatalf("failed to migrate schema: %v", err)
		}
	}

	// Services
	services := setupServices(db)

	// Setup server.
	srv := setupServer(*addr, services)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("filed to serve http: %v", err)
	}

}

func setupServer(addr string, services *httpBroker.Services) *http.Server {
	return httpBroker.NewServer(addr, services)
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
