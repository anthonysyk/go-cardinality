package schema

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetValue(t *testing.T) {
	type Item struct {
		IntValue         int
		Int64Value       int64
		StringValue      string
		SliceStringValue []string
		SliceIntValue    []int
	}
	tests := []struct {
		label         string
		item          Item
		field         Field
		expectedValue interface{}
	}{
		{
			label: "int",
			item:  Item{IntValue: 2020},
			field: Field{
				Name: "IntValue",
				Type: reflect.Int,
			},
			expectedValue: 2020,
		},
		{
			label: "int64",
			item:  Item{Int64Value: 2020},
			field: Field{
				Name: "Int64Value",
				Type: reflect.Int64,
			},
			expectedValue: int64(2020),
		},
		{
			label: "string",
			item:  Item{StringValue: "Action"},
			field: Field{
				Name: "StringValue",
				Type: reflect.Int64,
			},
			expectedValue: "Action",
		},
		{
			label: "slice_string",
			item:  Item{SliceStringValue: []string{"Action", "Thriller", "Comedy"}},
			field: Field{
				Name: "SliceStringValue",
				Type: reflect.Slice,
			},
			expectedValue: []interface{}{"Action", "Thriller", "Comedy"},
		},
		{
			label: "slice_int",
			item:  Item{SliceIntValue: []int{2020, 2021, 2023}},
			field: Field{
				Name: "SliceIntValue",
				Type: reflect.Slice,
			},
			expectedValue: []interface{}{2020, 2021, 2023},
		},
	}
	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			result := test.field.GetValue(test.item)
			assert.Equal(t, test.expectedValue, result)
		})
	}
}
