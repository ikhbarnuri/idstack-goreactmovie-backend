package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *Application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	get, err := app.Models.DB.Get(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	movie := get

	app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *Application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.Models.DB.All()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
