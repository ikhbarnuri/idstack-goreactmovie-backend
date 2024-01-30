package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"idstack-goreactmovie-backend/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

type MoviePayload struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        string `json:"year"`
	ReleaseDate string `json:"release_date"`
	Runtime     string `json:"runtime"`
	Rating      string `json:"rating"`
	MPPAARating string `json:"mppaa_rating"`
}

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

func (app *Application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.Models.DB.GetGenreAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, genres, "genres")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *Application) getAllMoviesByGenres(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	genreId, err := strconv.Atoi(params.ByName("genre_id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	movies, err := app.Models.DB.All(genreId)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "genres")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *Application) addMovie(w http.ResponseWriter, r *http.Request) {
	var payload MoviePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	var movie models.Movie

	movie.Id, _ = strconv.Atoi(payload.Id)
	movie.Title = payload.Title
	movie.Description = payload.Description
	movie.ReleaseDate, _ = time.Parse("2006-01-02", payload.ReleaseDate)
	movie.Year = movie.ReleaseDate.Year()
	movie.Runtime, _ = strconv.Atoi(payload.Runtime)
	movie.Rating, _ = strconv.Atoi(payload.Rating)
	movie.MPPAARating = payload.MPPAARating
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	err = app.Models.DB.InsertMovie(movie)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	type jsonRes struct {
		Ok bool `json:"ok"`
	}

	ok := jsonRes{
		Ok: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}
}

func (app *Application) editMovie(w http.ResponseWriter, r *http.Request) {
	var payload MoviePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	var movie models.Movie

	id, _ := strconv.Atoi(payload.Id)
	singleMovie, _ := app.Models.DB.Get(id)
	movie = *singleMovie

	movie.Id, _ = strconv.Atoi(payload.Id)
	movie.Title = payload.Title
	movie.Description = payload.Description
	movie.ReleaseDate, _ = time.Parse("2006-01-02", payload.ReleaseDate)
	movie.Year = movie.ReleaseDate.Year()
	movie.Runtime, _ = strconv.Atoi(payload.Runtime)
	movie.Rating, _ = strconv.Atoi(payload.Rating)
	movie.MPPAARating = payload.MPPAARating
	movie.UpdatedAt = time.Now()

	err = app.Models.DB.UpdateMovie(movie)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	type jsonRes struct {
		Ok bool `json:"ok"`
	}

	ok := jsonRes{
		Ok: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}
}

func (app *Application) deleteMovie(w http.ResponseWriter, r *http.Request) {
	var payload MoviePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	id, _ := strconv.Atoi(payload.Id)
	err = app.Models.DB.DeleteMovie(id)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	type jsonRes struct {
		Ok bool `json:"ok"`
	}

	ok := jsonRes{
		Ok: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}
}
