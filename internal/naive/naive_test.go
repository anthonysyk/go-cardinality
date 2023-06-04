package naive

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistinctCountTypes(t *testing.T) {
	type Item struct {
		Int int
	}
	tests := []struct {
		label     string
		fieldName string
		items     []Item
		expected  map[interface{}]int
	}{
		{
			label:     "nominal_int",
			fieldName: "Int",
			items: []Item{
				{Int: 2020}, {Int: 2020},
				{Int: 2021},
				{Int: 2022}, {Int: 2022}, {Int: 2022},
				{Int: 2023},
			},
			expected: map[interface{}]int{2020: 2, 2021: 1, 2022: 3, 2023: 1},
		},
	}
	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			fields := DistinctCount(Item{}, test.items, test.fieldName)
			result, err := fields.GetField(test.fieldName)
			assert.NoError(t, err)
			assert.EqualValues(t, test.expected, result.DistinctValues)
		})
	}
}
