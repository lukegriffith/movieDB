package movies

import (
	"errors"
	"fmt"
)

var (
	moviesList              []Movie
	MovieTitleValidationErr = errors.New("Movie title invalid")
)

type Title string

func NewTitle(title string) (error, Title) {
	if len(title) <= 0 {
		return MovieTitleValidationErr, ""
	}
	return nil, Title(title)
}

type Movie struct {
	Title Title
}

type MovieRequestProps struct {
	Title string
}

func init() {
	moviesList = make([]Movie, 0)
}

func NewMovie(req MovieRequestProps) (error, Movie) {
	err, title := NewTitle(req.Title)
	if err != nil {
		return fmt.Errorf("new movie error: %w", err), Movie{}
	}
	return nil, Movie{title}
}

func AddMovie(req MovieRequestProps) error {
	err, movie := NewMovie(req)
	if err != nil {
		return fmt.Errorf("Add movie error: %w", err)
	}
	moviesList = append(moviesList, movie)
	return nil
}

func GetMovies() *[]Movie {
	return &moviesList
}
