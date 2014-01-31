package missingg

import (
	"fmt"
)

func ExampleSet() {
	s := NewSet(0)
	s.Insert(1)
	s.Insert(2)
	s.Insert(2)
	s.Insert(3)
	s.Insert(7)
	sl := s.AsSlice()
	Sort(sl, func(a, b int) bool { return a < b })
	fmt.Println(sl)
	// Output: [1 2 3 7]
}
