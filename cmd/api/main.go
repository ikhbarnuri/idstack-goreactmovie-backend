package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
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

func main() {
	var config Config

	flag.IntVar(&config.Port, "port", 4000, "Server port to listen on ")
	flag.StringVar(&config.Env, "env", "development", "Applicaton environtment (development|production)")
	flag.Parse()

	fmt.Println("Server is running...")

	http.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
		currentStatus := AppStatus{
			Status:       "Online",
			Environtment: config.Env,
			Version:      version,
		}

		res, err := json.MarshalIndent(currentStatus, "", "\t")
		if err != nil {
			log.Println(err)

		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(res)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	if err != nil {
		log.Println(err)
	}
}
