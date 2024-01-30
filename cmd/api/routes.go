package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/movies/:id", app.getOneMovie)
	router.HandlerFunc(http.MethodGet, "/movies", app.getAllMovies)

	router.HandlerFunc(http.MethodGet, "/genres", app.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/genres/:genre_id/movies", app.getAllMoviesByGenres)

	router.HandlerFunc(http.MethodPost, "/admin/movies/add", app.addMovie)
	router.HandlerFunc(http.MethodPost, "/admin/movies/edit", app.editMovie)

	return app.enableCORS(router)
}
