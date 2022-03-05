package aki

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mapTestCases = []struct {
	name   string
	m      M[int, float64]
	keys   []int
	values []float64
	fp     MPredicate[int, float64]
	fm     M[int, float64]
	fkp    MPredicateByKey[int]
	fkm    M[int, float64]
	fvp    MPredicateByValue[float64]
	fvm    M[int, float64]
}{
	{
		name:   "nil",
		m:      nil,
		keys:   []int{},
		values: []float64{},
		fp:     func(_ int, _ float64) bool { return true },
		fm:     map[int]float64{},
		fkp:    func(_ int) bool { return true },
		fkm:    map[int]float64{},
		fvp:    func(_ float64) bool { return true },
		fvm:    map[int]float64{},
	},
	{
		name:   "empty",
		m:      map[int]float64{},
		keys:   []int{},
		values: []float64{},
		fp:     func(_ int, _ float64) bool { return true },
		fm:     map[int]float64{},
		fkp:    func(_ int) bool { return true },
		fkm:    map[int]float64{},
		fvp:    func(_ float64) bool { return true },
		fvm:    map[int]float64{},
	},
	{
		name:   "one_value",
		m:      M[int, float64]{1: 2},
		keys:   []int{1},
		values: []float64{2},
		fp:     func(_ int, _ float64) bool { return false },
		fm:     map[int]float64{},
		fkp:    func(_ int) bool { return false },
		fkm:    map[int]float64{},
		fvp:    func(_ float64) bool { return false },
		fvm:    map[int]float64{},
	},
	{
		name:   "multiple_value",
		m:      M[int, float64]{1: 2, 3: 4, 5: 6, 7: 8},
		keys:   []int{1, 3, 5, 7},
		values: []float64{2, 4, 6, 8},
		fp:     func(key int, value float64) bool { return key == 1 || value == 4 },
		fm:     map[int]float64{1: 2, 3: 4},
		fkp:    func(key int) bool { return key == 1 },
		fkm:    map[int]float64{1: 2},
		fvp:    func(value float64) bool { return value == 4 },
		fvm:    map[int]float64{3: 4},
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
			assert.Equal(t, tt.fm, tt.m.Filter(tt.fp))
		})
	}
}

func TestMFilter(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, map[int]float64(tt.fm), MFilter(tt.m, tt.fp))
		})
	}
}

func TestM_FilterByKey(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.fkm, tt.m.FilterByKey(tt.fkp))
		})
	}
}

func TestMFilterByKey(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, map[int]float64(tt.fkm), MFilterByKey(tt.m, tt.fkp))
		})
	}
}

func TestM_FilterByValue(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.fvm, tt.m.FilterByValue(tt.fvp))
		})
	}
}

func TestMFilterByValue(t *testing.T) {
	for _, tt := range mapTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, map[int]float64(tt.fvm), MFilterByValue(tt.m, tt.fvp))
		})
	}
}
