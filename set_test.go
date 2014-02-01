package missingg

import (
	"testing"

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
	Sort(sl, func(a, b interface{}) bool { return a.(int) < b.(int) })
	fmt.Println(sl)
	// Output: [1 2 3 7]
}

func TestTypeFailure(t *testing.T) {
	type dumbStruct struct {
		x string
	}
	rows := []struct {
		setType interface{}
		value   interface{}
		ok      bool
	}{
		{0, 0, true},
		{0, 1, true},
		{"", "hi", true},
		{dumbStruct{}, dumbStruct{}, true},
		{int64(0), 1, false},
		{"", 5, false},
		{dumbStruct{}, &dumbStruct{}, false},
	}
	for _, r := range rows {
		s := NewSet(r.setType).(*setImplmentation)
		if ok := s.isTypeOK(r.value); ok != r.ok {
			t.Errorf("Value (%v): Want: %v Got: %v", r.value, r.ok, ok)
		}
		if r.ok {
			s.Insert(r.value)
			s.Insert(r.value)
			if len(s.AsSlice()) != 1 {
				t.Errorf("Duplicates found for value %v", r.value)
			}
		}
	}
}
