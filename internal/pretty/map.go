package pretty

import (
	"fmt"
	"sort"
	"strings"
)

func MapToEnum(m map[interface{}]int) string {
	// Sort the map entries by occurrence in descending order
	type pair struct {
		Key   interface{}
		Value int
	}
	pairs := make([]pair, 0, len(m))
	for key, value := range m {
		pairs = append(pairs, pair{key, value})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	// Build the enum elements
	elements := make([]string, len(pairs))
	for i, p := range pairs {
		elements[i] = fmt.Sprintf("%s = %d", p.Key, p.Value)
	}

	// Format the enum
	return strings.Join(elements, "\n")
}
