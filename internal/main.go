package main

import (
	"fmt"

	"github.com/mymmrac/aki/is"
	"github.com/mymmrac/aki/maps"
)

func main() {
	m := maps.Map[string, int]{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	//m.FillEntry(maps.Entry[string, int]{"f", 4})
	m.FillEntry(maps.NewEntry(is.Or("", "f"), 4))

	fmt.Println(m)
}
