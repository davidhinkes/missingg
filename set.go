package missingg

import (
	"fmt"
	"reflect"
)

// Set provides set operations.
type Set interface {
	Insert(x interface{})
	Erase(x interface{}) bool
	Has(x interface{}) bool

	// Returns all elements in a slice.
	AsSlice() []interface{}
}

// NewSet creates a new Set.
func NewSet(exampleElement interface{}) Set {
	return &setImplmentation{
		ty:       reflect.TypeOf(exampleElement),
		elements: make(map[interface{}]struct{}),
	}
}

// NewSetWithCapacity creates a new set with a capacity hint.
func NewSetWithCapacity(exampleElement interface{}, capacity int) Set {
	return &setImplmentation{
		ty:       reflect.TypeOf(exampleElement),
		elements: make(map[interface{}]struct{}, capacity),
	}
}

type setImplmentation struct {
	ty       reflect.Type
	elements map[interface{}]struct{}
}

func (s *setImplmentation) isTypeOK(x interface{}) bool {
	return reflect.TypeOf(x) == s.ty
}

func (s *setImplmentation) ensureType(x interface{}) {
	if !s.isTypeOK(x) {
		panic(fmt.Sprintf("Got %v, want %v", reflect.TypeOf(x), s.ty))
	}
}

func (s *setImplmentation) Has(x interface{}) bool {
	s.ensureType(x)
	_, has := s.elements[x]
	return has
}

func (s *setImplmentation) Insert(x interface{}) {
	s.ensureType(x)
	s.elements[x] = struct{}{}
}

func (s *setImplmentation) Erase(x interface{}) bool {
	s.ensureType(x)
	if !s.Has(x) {
		return false
	}
	delete(s.elements, x)
	return true
}

func (s *setImplmentation) AsSlice() []interface{} {
	slice := make([]interface{}, 0, len(s.elements))
	for k, _ := range s.elements {
		slice = append(slice, k)
	}
	return slice
}
