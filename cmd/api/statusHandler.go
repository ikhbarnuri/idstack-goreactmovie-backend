package main

import (
	"encoding/json"
	"net/http"
)

func (app *Application) statusHandler(writer http.ResponseWriter, request *http.Request) {
	currentStatus := AppStatus{
		Status:       "Online",
		Environtment: app.Config.Env,
		Version:      version,
	}

	res, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		app.Logger.Println(err)

	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}
