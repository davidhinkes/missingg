package missingg_test

import (
	"fmt"
	. "github.com/davidhinkes/missingg"
)

func ExampleSort() {
	s := []int{5, 4, 7, 9}
	Sort(s, func(a, b int) bool { return a < b })
	fmt.Println(s)
	// Output: [4 5 7 9]
}
