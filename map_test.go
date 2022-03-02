package aki

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mapTestCases = []struct {
	name   string
	m      M[int, int]
	keys   []int
	values []int
}{
	{
		name:   "nil",
		m:      nil,
		keys:   []int{},
		values: []int{},
	},
	{
		name:   "empty",
		m:      make(map[int]int),
		keys:   []int{},
		values: []int{},
	},
	{
		name:   "one_value",
		m:      M[int, int]{1: 2},
		keys:   []int{1},
		values: []int{2},
	},
	{
		name:   "multiple_value",
		m:      M[int, int]{1: 2, 3: 4, 5: 6, 7: 8},
		keys:   []int{1, 3, 5, 7},
		values: []int{2, 4, 6, 8},
	},
}

func TestM_Values(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.values, tt.m.Values())
		})
	}
}

func TestM_Keys(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.keys, tt.m.Keys())
		})
	}
}
