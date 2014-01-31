package missingg

import (
	"reflect"
	"sort"
)

type st struct {
	vxs reflect.Value
	vf  reflect.Value
}

func (s st) Len() int {
	return s.vxs.Len()
}

func (s st) Swap(i, j int) {
	vi := s.vxs.Index(i)
	vj := s.vxs.Index(j)
	t := reflect.ValueOf(vi.Interface())
	vi.Set(vj)
	vj.Set(t)
}

func (s st) Less(i, j int) bool {
	vi := s.vxs.Index(i)
	vj := s.vxs.Index(j)
	return s.vf.Call([]reflect.Value{vi, vj})[0].Bool()
}

// Sort sorts a slice given a compare function.
// xs is a slice
// compare is a less function of the form func(a, b T) bool
func Sort(xs, compare interface{}) {
	sort.Sort(st{reflect.ValueOf(xs), reflect.ValueOf(compare)})
}
