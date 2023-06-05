package pretty

import (
	"fmt"
	"testing"
)

func TestMapToEnum(t *testing.T) {
	m := map[interface{}]int{
		"Action":    162,
		"Adventure": 43,
		"Animated":  79,
		"Western":   13,
	}
	enumStr := MapToEnum(m)
	fmt.Println(enumStr)
}
