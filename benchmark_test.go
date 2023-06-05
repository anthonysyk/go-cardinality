package go_cardinality

import (
	"github.com/anthonysyk/go-cardinality/example"
	"github.com/anthonysyk/go-cardinality/internal/naive"
	"testing"
)

func BenchmarkNaive(b *testing.B) {
	movies, _ := example.GetMovies()
	naive.DistinctCount(example.Movie{}, movies, "Year", "Genres")
}
