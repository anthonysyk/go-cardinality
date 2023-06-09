package schema

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewField(t *testing.T) {
	field := NewField("Value", reflect.String)
	assert.Equal(t, "Value", field.Name)
	assert.Equal(t, reflect.String, field.Type)
	assert.NotNil(t, field.DistinctValues)
}

func TestFilterFields(t *testing.T) {
	type Item struct {
		A int
		B string
		C []string
	}
	fields := FilterFields(Item{}, "A", "C")
	fieldA, err := fields.GetField("A")
	assert.NoError(t, err)
	_, err = fields.GetField("B")
	assert.Error(t, err)
	fieldC, err := fields.GetField("C")
	assert.NoError(t, err)
	assert.Contains(t, fields, fieldA)
	assert.Contains(t, fields, fieldC)
	assert.Len(t, fields, 2)
}

func TestField_AddValues(t *testing.T) {
	tests := []struct {
		label     string
		field     *Field
		addValues func(field *Field)
		expected  map[interface{}]int
	}{
		{
			label: "single_value_int",
			field: NewField("A", reflect.Int),
			addValues: func(field *Field) {
				field.AddValues(1)
			},
			expected: map[interface{}]int{1: 1},
		},
		{
			label: "slice_values_int",
			field: NewField("A", reflect.Int),
			addValues: func(field *Field) {
				field.AddValues([]interface{}{1, 2, 3})
			},
			expected: map[interface{}]int{1: 1, 2: 1, 3: 1},
		},
		{
			label: "slice_values_slice",
			field: NewField("A", reflect.Slice),
			addValues: func(field *Field) {
				field.AddValues([]interface{}{1, 2, 3})
			},
			expected: map[interface{}]int{1: 1, 2: 1, 3: 1},
		},
	}

	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			test.addValues(test.field)
			assert.Len(t, test.field.DistinctValues, len(test.expected))
			assert.Equal(t, test.expected, test.field.DistinctValues)
		})
	}
}

func TestField_GetValue(t *testing.T) {
	type Item struct {
		A int
		B int
	}

	item := Item{A: 2020, B: 2021}
	aValue := NewField("A", reflect.Int).GetValue(item)
	assert.Equal(t, item.A, aValue)
	bValue := NewField("B", reflect.Int).GetValue(item)
	assert.Equal(t, item.B, bValue)
}
