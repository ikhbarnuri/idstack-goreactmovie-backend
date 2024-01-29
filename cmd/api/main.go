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

type Config struct {
	Port int
	Env  string
}

type AppStatus struct {
	Status       string `json:"status"`
	Environtment string `json:"environtment"`
	Version      string `json:"version"`
}

type Application struct {
	Config
	Logger *log.Logger
}

func main() {
	var config Config

	flag.IntVar(&config.Port, "port", 4000, "Server port to listen on ")
	flag.StringVar(&config.Env, "env", "development", "Applicaton environtment (development|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		Config: config,
		Logger: logger,
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

	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
