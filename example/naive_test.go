package example

import (
	"github.com/anthonysyk/go-cardinality/internal/naive"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistinctCountNaive(t *testing.T) {
	movies, err := GetMovies()
	assert.NoError(t, err)
	fields := naive.DistinctCount(Movie{}, movies, "Year", "Genres")
	years, err := fields.GetField("Year")
	assert.NoError(t, err)
	assert.Equal(t, 275, years.DistinctValues[2020])
	assert.Equal(t, 360, years.DistinctValues[2021])
	assert.Equal(t, 326, years.DistinctValues[2022])
	assert.Equal(t, 192, years.DistinctValues[2023])
	genres, err := fields.GetField("Genres")
	assert.NoError(t, err)
	assert.Equal(t, 162, genres.DistinctValues["Action"])
	assert.Equal(t, 43, genres.DistinctValues["Adventure"])
	assert.Equal(t, 79, genres.DistinctValues["Animated"])
	assert.Equal(t, 59, genres.DistinctValues["Biography"])
}
