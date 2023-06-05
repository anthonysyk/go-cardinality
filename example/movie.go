package example

import (
	_ "embed"
	"encoding/json"
)

//go:embed movies.json
var MoviesJson []byte

type Movie struct {
	Title           string   `json:"title"`
	Year            int      `json:"year"`
	Cast            []string `json:"cast"`
	Genres          []string `json:"genres"`
	Href            string   `json:"href"`
	Extract         string   `json:"extract"`
	Thumbnail       string   `json:"thumbnail"`
	ThumbnailWidth  int      `json:"thumbnail_width"`
	ThumbnailHeight int      `json:"thumbnail_height"`
}

func GetMovies() ([]Movie, error) {
	var movies []Movie
	err := json.Unmarshal(MoviesJson, &movies)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
