package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mymmrac/aki"
)

func TestToM(t *testing.T) {
	m1 := map[int]float64{1: 2}
	m2 := M[int, float64]{3: 4}
	var m3 map[int]float64
	var m4 M[int, float64]

	t.Run("map", func(t *testing.T) {
		assert.Equal(t, M[int, float64](m1), ToM(m1))
	})

	t.Run("m_map", func(t *testing.T) {
		assert.Equal(t, m2, ToM(m2))
	})

	t.Run("nil", func(t *testing.T) {
		assert.Equal(t, m4, ToM(m3))
	})
}

var mapTestCases = []struct {
	name                   string
	m                      M[int, float64]
	keys                   []int
	values                 []float64
	filterPredicate        Predicate[int, float64]
	filteredMap            M[int, float64]
	filterKeyPredicate     PredicateByKey[int]
	filteredKeyMap         M[int, float64]
	filteredValuePredicate PredicateByValue[float64]
	filteredValueMap       M[int, float64]
	entries                []Entry[int, float64]
}{
	{
		name:                   "nil",
		m:                      nil,
		keys:                   []int{},
		values:                 []float64{},
		filterPredicate:        func(_ int, _ float64) bool { return true },
		filteredMap:            nil,
		filterKeyPredicate:     func(_ int) bool { return true },
		filteredKeyMap:         nil,
		filteredValuePredicate: func(_ float64) bool { return true },
		filteredValueMap:       nil,
		entries:                []Entry[int, float64]{},
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
		entries:                []Entry[int, float64]{},
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
		entries: []Entry[int, float64]{
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
		entries: []Entry[int, float64]{
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

func TestValues(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.values, Values(tt.m))
		})
	}
}

func TestKeys(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.keys, Keys(tt.m))
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

func TestFilter(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, map[int]float64(tt.filteredMap), Filter(tt.m, tt.filterPredicate))
		})
	}
}

func TestM_FilterSelf(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.m.Copy()
			assert.Equal(t, tt.filteredMap, m.FilterSelf(tt.filterPredicate))
			assert.Equal(t, tt.filteredMap, m)
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

func TestFilterByKey(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, map[int]float64(tt.filteredKeyMap), FilterByKey(tt.m, tt.filterKeyPredicate))
		})
	}
}

func TestM_FilterSelfByKey(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.m.Copy()
			assert.Equal(t, tt.filteredKeyMap, m.FilterSelfByKey(tt.filterKeyPredicate))
			assert.Equal(t, tt.filteredKeyMap, m)
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

func TestFilterByValue(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, map[int]float64(tt.filteredValueMap), FilterByValue(tt.m, tt.filteredValuePredicate))
		})
	}
}

func TestM_FilterSelfByValue(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.m.Copy()
			assert.Equal(t, tt.filteredValueMap, m.FilterSelfByValue(tt.filteredValuePredicate))
			assert.Equal(t, tt.filteredValueMap, m)
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

func TestEntries(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.entries, Entries(tt.m))
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

func TestFillEntries(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			m := M[int, float64]{}
			FillEntries(m, tt.entries)

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

func TestFillEntry(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			m := M[int, float64]{}
			FillEntry(m, tt.entries...)

			if tt.m == nil {
				assert.Equal(t, M[int, float64]{}, m)
				return
			}

			assert.Equal(t, m, tt.m)
		})
	}
}

func TestFromEntries(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			m := FromEntries(tt.entries)

			if tt.m == nil {
				assert.Equal(t, M[int, float64]{}, m)
				return
			}

			assert.Equal(t, tt.m, m)
		})
	}
}

func TestFromEntry(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			m := FromEntry(tt.entries...)

			if tt.m == nil {
				assert.Equal(t, M[int, float64]{}, m)
				return
			}

			assert.Equal(t, tt.m, m)
		})
	}
}

func TestM_Copy(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.m, tt.m.Copy())
		})
	}
}

func TestCopy(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, map[int]float64(tt.m), Copy(tt.m))
		})
	}
}

var twoMapsTestCases = []struct {
	name      string
	this      M[int, float64]
	other     M[int, float64]
	merge     M[int, float64]
	mergeLeft M[int, float64]
}{
	{
		name:      "nil-nil",
		this:      nil,
		other:     nil,
		merge:     nil,
		mergeLeft: nil,
	},
	{
		name:      "nil-empty",
		this:      nil,
		other:     M[int, float64]{},
		merge:     nil,
		mergeLeft: nil,
	},
	{
		name:      "empty-nil",
		this:      M[int, float64]{},
		other:     nil,
		merge:     M[int, float64]{},
		mergeLeft: M[int, float64]{},
	},
	{
		name:      "empty-empty",
		this:      M[int, float64]{},
		other:     M[int, float64]{},
		merge:     M[int, float64]{},
		mergeLeft: M[int, float64]{},
	},
	{
		name:      "one-nil",
		this:      M[int, float64]{1: 2},
		other:     nil,
		merge:     M[int, float64]{1: 2},
		mergeLeft: M[int, float64]{1: 2},
	},
	{
		name:      "nil-one",
		this:      nil,
		other:     M[int, float64]{1: 2},
		merge:     M[int, float64]{1: 2},
		mergeLeft: M[int, float64]{1: 2},
	},
	{
		name:      "same-same",
		this:      M[int, float64]{1: 2},
		other:     M[int, float64]{1: 2},
		merge:     M[int, float64]{1: 2},
		mergeLeft: M[int, float64]{1: 2},
	},
	{
		name:      "diff",
		this:      M[int, float64]{1: 2},
		other:     M[int, float64]{3: 4},
		merge:     M[int, float64]{1: 2, 3: 4},
		mergeLeft: M[int, float64]{1: 2, 3: 4},
	},
	{
		name:      "diff-with-same",
		this:      M[int, float64]{1: 2},
		other:     M[int, float64]{1: 5, 3: 4},
		merge:     M[int, float64]{1: 5, 3: 4},
		mergeLeft: M[int, float64]{1: 2, 3: 4},
	},
}

func TestM_Merge(t *testing.T) {
	for _, tt := range twoMapsTestCases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.this == nil && len(tt.other) != 0 {
				assert.Panics(t, func() {
					tt.this.Merge(tt.other)
				})
				return
			}

			assert.Equal(t, tt.merge, tt.this.Merge(tt.other))
		})
	}
}

func TestNewEntry(t *testing.T) {
	e1 := Entry[int, float64]{
		Key:   1,
		Value: 2,
	}

	assert.Equal(t, e1, NewEntry(e1.Key, e1.Value))
}

type cloneableFloat float64

func (c cloneableFloat) Clone() aki.Cloneable[cloneableFloat] {
	return c
}

func TestToMCloneable(t *testing.T) {
	m1 := map[int]cloneableFloat{1: 2}
	m2 := MCloneable[int, cloneableFloat, cloneableFloat]{3: cloneableFloat(4)}
	var m3 map[int]cloneableFloat
	var m4 MCloneable[int, cloneableFloat, cloneableFloat]

	t.Run("map", func(t *testing.T) {
		assert.Equal(t, MCloneable[int, cloneableFloat, cloneableFloat](m1), ToMCloneable[int, cloneableFloat](m1))
	})

	t.Run("m_map", func(t *testing.T) {
		assert.Equal(t, m2, ToMCloneable[int, cloneableFloat](m2))
	})

	t.Run("nil", func(t *testing.T) {
		assert.Equal(t, m4, ToMCloneable[int, cloneableFloat](m3))
	})
}
