package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"idstack-goreactmovie-backend/models"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn string
	}
}

type AppStatus struct {
	Status       string `json:"status"`
	Environtment string `json:"environtment"`
	Version      string `json:"version"`
}

type Application struct {
	Config
	Logger *log.Logger
	Models models.Models
}

func main() {
	var config Config

	flag.IntVar(&config.Port, "port", 4000, "Server port to listen on ")
	flag.StringVar(&config.Env, "env", "development", "Applicaton environtment (development|production)")
	flag.StringVar(&config.Db.Dsn, "dsn", "postgres://postgres:postgres@localhost/gomoviereact?sslmode=disable", "Postgre connection config")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	open, err := openDB(config)
	if err != nil {
		logger.Fatal(err)
	}
	defer func(open *sql.DB) {
		err := open.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(open)

	app := &Application{
		Config: config,
		Logger: logger,
		Models: models.NewModels(open),
	}

	fmt.Println("Server is running...")

	//http.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
	//	currentStatus := AppStatus{
	//		Status:       "Online",
	//		Environtment: config.Env,
	//		Version:      version,
	//	}
	//
	//	res, err := json.MarshalIndent(currentStatus, "", "\t")
	//	if err != nil {
	//		log.Println(err)
	//
	//	}
	//
	//	writer.Header().Set("Content-Type", "application/json")
	//	writer.WriteHeader(http.StatusOK)
	//	writer.Write(res)
	//})

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting server on port %d", config.Port)

	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func openDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Db.Dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
