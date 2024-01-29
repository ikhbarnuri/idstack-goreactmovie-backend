package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"idstack-goreactmovie-backend/models"
	"net/http"
	"strconv"
	"time"
)

func (app *Application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.Logger.Println(errors.New("Invalid id paramater"))
	}

	movie := models.Movie{
		Id:          id,
		Title:       "some movie title",
		Description: "some description",
		Year:        25,
		ReleaseDate: time.Date(1990, 01, 01, 01, 0, 0, 0, time.Local),
		Runtime:     112,
		Rating:      5,
		MPPAARating: "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	app.writeJSON(w, http.StatusOK, movie, "movie")
}
