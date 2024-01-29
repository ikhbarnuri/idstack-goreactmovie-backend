package main

import (
	"encoding/json"
	"net/http"
)

func (app *Application) writeJSON(w http.ResponseWriter, status int, data any, wrap string) error {
	wrapper := make(map[string]any)

	wrapper[wrap] = data

	res, err := json.Marshal(wrapper)
	if err != nil {
		app.Logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(res)
	if err != nil {
		return err
	}

	return nil
}

func (app *Application) errorJSON(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	errMessage := jsonError{
		Message: err.Error(),
	}

	app.writeJSON(w, http.StatusBadRequest, errMessage, "error")
}
