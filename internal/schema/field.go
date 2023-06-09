package schema

import (
	"errors"
	"fmt"
	"github.com/anthonysyk/go-cardinality/internal/pretty"
	"reflect"
)

type Field struct {
	Name           string
	Type           reflect.Kind
	DistinctValues map[interface{}]int
}

func NewField(fieldName string, fieldType reflect.Kind) *Field {
	return &Field{
		Name:           fieldName,
		Type:           fieldType,
		DistinctValues: make(map[interface{}]int),
	}
}

func (f *Field) AddValues(value interface{}) {
	if IsSlice(value) {
		sliceValue := reflect.ValueOf(value)
		for i := 0; i < sliceValue.Len(); i++ {
			fieldValue := sliceValue.Index(i).Interface()
			f.DistinctValues[fieldValue]++
		}
		return
	}
	f.DistinctValues[value]++
}

func (f *Field) GetValue(e interface{}) interface{} {
	return getValue(e, f.Name)
}

func (f *Field) PrettyPrint() {
	fmt.Println(pretty.MapToEnum(f.DistinctValues))
}

type Fields []*Field

func (fs Fields) GetField(name string) (*Field, error) {
	for _, f := range fs {
		if f.Name == name {
			return f, nil
		}
	}
	return nil, errors.New("field name is empty")
}

func FilterFields[T any](empty T, fields ...string) Fields {
	pType := reflect.TypeOf(empty)
	var filteredFields Fields
	for i := 0; i < pType.NumField(); i++ {
		field := pType.Field(i)
		for _, f := range fields {
			if f == field.Name {
				filteredFields = append(filteredFields, NewField(field.Name, field.Type.Kind()))
			}
		}
	}
	return filteredFields
}
