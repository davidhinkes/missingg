package missingg_test

import (
	"fmt"
	. "github.com/davidhinkes/missingg"
)

func ExampleSet() {
	s := NewSet(0)
	s.Insert(1)
	s.Insert(2)
	s.Insert(2)
	s.Insert(3)
	s.Insert(7)
	sl := s.AsSlice()
	Sort(sl, func(a, b interface{}) bool { return a.(int) < b.(int) })
	fmt.Println(sl)
	// Output: [1 2 3 7]
}
