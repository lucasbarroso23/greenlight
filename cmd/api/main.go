package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// this struct holds all configuration settings for our application
type config struct {
	port int
	env  string
}

// this struct is responsible to group the dependecies for our HTTP handlers, middlewares and helpers
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	// Reading comnmand-line flags in to the config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initializing a new logger witch writes messages to the standard out stream
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// application config struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// setting up a new servemux to our handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// setup http server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start http server
	logger.Printf("starting %server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
