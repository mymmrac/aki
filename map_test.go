package aki

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mapTestCases = []struct {
	name                   string
	m                      M[int, float64]
	keys                   []int
	values                 []float64
	filterPredicate        MPredicate[int, float64]
	filteredMap            M[int, float64]
	filterKeyPredicate     MPredicateByKey[int]
	filteredKeyMap         M[int, float64]
	filteredValuePredicate MPredicateByValue[float64]
	filteredValueMap       M[int, float64]
	entries                []MEntry[int, float64]
}{
	{
		name:                   "nil",
		m:                      nil,
		keys:                   []int{},
		values:                 []float64{},
		filterPredicate:        func(_ int, _ float64) bool { return true },
		filteredMap:            map[int]float64{},
		filterKeyPredicate:     func(_ int) bool { return true },
		filteredKeyMap:         map[int]float64{},
		filteredValuePredicate: func(_ float64) bool { return true },
		filteredValueMap:       map[int]float64{},
		entries:                []MEntry[int, float64]{},
	},
	{
		name:                   "empty",
		m:                      map[int]float64{},
		keys:                   []int{},
		values:                 []float64{},
		filterPredicate:        func(_ int, _ float64) bool { return true },
		filteredMap:            map[int]float64{},
		filterKeyPredicate:     func(_ int) bool { return true },
		filteredKeyMap:         map[int]float64{},
		filteredValuePredicate: func(_ float64) bool { return true },
		filteredValueMap:       map[int]float64{},
		entries:                []MEntry[int, float64]{},
	},
	{
		name:                   "one_value",
		m:                      M[int, float64]{1: 2},
		keys:                   []int{1},
		values:                 []float64{2},
		filterPredicate:        func(_ int, _ float64) bool { return false },
		filteredMap:            map[int]float64{},
		filterKeyPredicate:     func(_ int) bool { return false },
		filteredKeyMap:         map[int]float64{},
		filteredValuePredicate: func(_ float64) bool { return false },
		filteredValueMap:       map[int]float64{},
		entries: []MEntry[int, float64]{
			{Key: 1, Value: 2},
		},
	},
	{
		name:                   "multiple_value",
		m:                      M[int, float64]{1: 2, 3: 4, 5: 6, 7: 8},
		keys:                   []int{1, 3, 5, 7},
		values:                 []float64{2, 4, 6, 8},
		filterPredicate:        func(key int, value float64) bool { return key == 1 || value == 4 },
		filteredMap:            map[int]float64{1: 2, 3: 4},
		filterKeyPredicate:     func(key int) bool { return key == 1 },
		filteredKeyMap:         map[int]float64{1: 2},
		filteredValuePredicate: func(value float64) bool { return value == 4 },
		filteredValueMap:       map[int]float64{3: 4},
		entries: []MEntry[int, float64]{
			{Key: 1, Value: 2},
			{Key: 3, Value: 4},
			{Key: 5, Value: 6},
			{Key: 7, Value: 8},
		},
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

func TestMValues(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.values, MValues(tt.m))
		})
	}
}

func TestMKeys(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.keys, MKeys(tt.m))
		})
	}
}

func TestM_Filter(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.filteredMap, tt.m.Filter(tt.filterPredicate))
		})
	}
}

func TestMFilter(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, map[int]float64(tt.filteredMap), MFilter(tt.m, tt.filterPredicate))
		})
	}
}

func TestM_FilterByKey(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.filteredKeyMap, tt.m.FilterByKey(tt.filterKeyPredicate))
		})
	}
}

func TestMFilterByKey(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, map[int]float64(tt.filteredKeyMap), MFilterByKey(tt.m, tt.filterKeyPredicate))
		})
	}
}

func TestM_FilterByValue(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.filteredValueMap, tt.m.FilterByValue(tt.filteredValuePredicate))
		})
	}
}

func TestMFilterByValue(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, map[int]float64(tt.filteredValueMap), MFilterByValue(tt.m, tt.filteredValuePredicate))
		})
	}
}

func TestM_Entries(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.entries, tt.m.Entries())
		})
	}
}

func TestMEntries(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.entries, MEntries(tt.m))
		})
	}
}

func TestM_FillEntries(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			m := M[int, float64]{}
			m.FillEntries(tt.entries)

			if tt.m == nil {
				assert.Equal(t, M[int, float64]{}, m)
				return
			}
			assert.Equal(t, m, tt.m)
		})
	}
}

func TestMFillEntries(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			m := M[int, float64]{}
			MFillEntries(m, tt.entries)

			if tt.m == nil {
				assert.Equal(t, M[int, float64]{}, m)
				return
			}
			assert.Equal(t, m, tt.m)
		})
	}
}

func TestM_FillEntry(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			m := M[int, float64]{}
			m.FillEntry(tt.entries...)

			if tt.m == nil {
				assert.Equal(t, M[int, float64]{}, m)
				return
			}

			assert.Equal(t, m, tt.m)
		})
	}
}

func TestMFillEntry(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			m := M[int, float64]{}
			MFillEntry(m, tt.entries...)

			if tt.m == nil {
				assert.Equal(t, M[int, float64]{}, m)
				return
			}

			assert.Equal(t, m, tt.m)
		})
	}
}

func TestMFromEntries(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			m := MFromEntries(tt.entries)

			if tt.m == nil {
				assert.Equal(t, M[int, float64]{}, m)
				return
			}

			assert.Equal(t, tt.m, m)
		})
	}
}

func TestMFromEntry(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			m := MFromEntry(tt.entries...)

			if tt.m == nil {
				assert.Equal(t, M[int, float64]{}, m)
				return
			}

			assert.Equal(t, tt.m, m)
		})
	}
}
