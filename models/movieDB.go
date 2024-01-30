package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `select * from movies where id=$1`
	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie
	err := row.Scan(
		&movie.Id,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Rating,
		&movie.Runtime,
		&movie.MPPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	query = `select
				mg.id, mg.movie_id, mg.genre_id, g.genre_name
			from
				movies_genres mg
				left join genres g on (g.id = mg.genre_id)
			where
				mg.movie_id = $1
		`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	genres := make(map[int]string)
	for rows.Next() {
		var mg MovieGenre
		err := rows.Scan(
			&mg.Id,
			&mg.MovieId,
			&mg.GenreId,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[mg.Id] = mg.Genre.GenreName
	}

	movie.MovieGenre = genres

	return &movie, nil
}

func (m *DBModel) All(genre ...int) ([]*Movie, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	where := ""

	if len(genre) > 0 {
		where = fmt.Sprintf(
			"WHERE id IN ("+
				"SELECT "+
				"movie_id "+
				"FROM "+
				"movies_genres "+
				"WHERE "+
				"genre_id = %d)",
			genre[0],
		)
	}

	query := fmt.Sprintf(`SELECT * FROM movies %s ORDER BY title`, where)
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next() {
		var movie Movie
		err = rows.Scan(
			&movie.Id,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Rating,
			&movie.Runtime,
			&movie.MPPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		genreQuery := `select
				mg.id, mg.movie_id, mg.genre_id, g.genre_name
			from
				movies_genres mg
				left join genres g on (g.id = mg.genre_id)
			where
				mg.movie_id = $1
		`

		genreRows, _ := m.DB.QueryContext(ctx, genreQuery, movie.Id)
		defer genreRows.Close()

		genres := make(map[int]string)
		for genreRows.Next() {
			var mg MovieGenre
			err := genreRows.Scan(
				&mg.Id,
				&mg.MovieId,
				&mg.GenreId,
				&mg.Genre.GenreName,
			)
			if err != nil {
				return nil, err
			}
			genres[mg.GenreId] = mg.Genre.GenreName
		}
		genreRows.Close()
		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}

	return movies, nil
}

func (m *DBModel) GetGenreAll() ([]*Genre, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	query := `SELECT 
				*
			FROM 
				genres 
			ORDER BY 
				genre_name	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []*Genre

	for rows.Next() {
		var genre Genre
		err = rows.Scan(
			&genre.Id,
			&genre.GenreName,
			&genre.CreatedAt,
			&genre.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		genres = append(genres, &genre)
	}

	return genres, nil
}
