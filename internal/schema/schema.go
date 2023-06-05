package schema

import (
	"reflect"
)

func GetValue(e interface{}, field string) interface{} {
	r := reflect.ValueOf(e)
	value := reflect.Indirect(r).FieldByName(field)

	if value.Kind() == reflect.Slice {
		numElements := value.Len()
		values := make([]interface{}, numElements)
		for i := 0; i < numElements; i++ {
			values[i] = value.Index(i).Interface()
		}
		return values
	}

	return value.Interface()
}

func IsSlice(value interface{}) bool {
	valueType := reflect.TypeOf(value)
	return valueType.Kind() == reflect.Slice
}
