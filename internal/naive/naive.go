package naive

import (
	"github.com/anthonysyk/go-cardinality/internal/schema"
)

func DistinctCount[T any](empty T, items []T, fields ...string) schema.Fields {
	filteredFields := schema.FilterFields(empty, fields...)

	for _, item := range items {
		for _, field := range filteredFields {
			value := field.GetValue(item)
			field.AddValues(value)
		}
	}
	return filteredFields
}
