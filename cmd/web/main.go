package main

import (
	"flag"
	"html/template"
	"log"
	"os"
)

const version = "1.0.0"
const cssVersion = "1"

type config struct {
	port int
	env string
	api string
	db struct {
		// data source name
		dsn string
	}
	stripe struct {
		secret string
		key string
	}
}

type application struct {
	config config
	infoLog *log.Logger
	errorLog *log.Logger
	templateCache map[string]*template.Template
	version string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development | production)")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to api")

	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	tc := make(map[string]*template.Template)

	app := &application{
		config: cfg,
		infoLog: infoLog,
		errorLog: errorLog,
		templateCache: tc,
		version: version,
	}
}