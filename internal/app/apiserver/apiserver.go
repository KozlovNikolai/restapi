package apiserver

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/KozlovNikolai/restapi/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer log.Fatal(db.Close())

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)

	fmt.Printf("config.BindAddr: %v\n", config.BindAddr)
	fmt.Printf("config.DatabaseURL: %v\n", config.DatabaseURL)
	fmt.Printf("config.IdleTimout: %v\n", config.IdleTimout)
	fmt.Printf("config.Timeout: %v\n", config.Timeout)
	fmt.Printf("config.LogLevel: %v\n", config.LogLevel)

	s := &http.Server{
		Addr:         config.BindAddr,
		Handler:      srv,
		ReadTimeout:  config.Timeout,
		WriteTimeout: config.Timeout,
		IdleTimeout:  config.IdleTimout,
	}

	return s.ListenAndServe()
	//return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
