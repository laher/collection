package collections

import (
	"iter"
)

type OrderedSet[E comparable] struct {
	set   *Set[E]
	slice []E
}

// use the compiler to confirm that a set respects the Collection interface
var _ Collection[string] = &OrderedSet[string]{}

func NewOrderedSet[E comparable](vals ...E) *OrderedSet[E] {
	s := &OrderedSet[E]{}
	s.init(vals...)
	return s
}

func (s *OrderedSet[E]) init(vals ...E) {
	s.set = NewSet(vals...)
	s.slice = vals
}

func (s *OrderedSet[E]) Add(vals ...E) {
	for _, v := range vals {
		if _, ok := s.set.items[v]; !ok {
			s.slice = append(s.slice, v)
			s.set.items[v] = struct{}{}
		}
	}
}

func (s *OrderedSet[E]) Len() int {
	return len(s.slice)
}

func (s *OrderedSet[E]) IsEmpty() bool {
	return len(s.slice) == 0
}

func (s *OrderedSet[E]) Contains(v E) bool {
	_, ok := s.set.items[v]
	return ok
}

func (s *OrderedSet[E]) Clear() {
	clear(s.set.items)
	clear(s.slice)
}

func (s *OrderedSet[E]) Slice() []E {
	return s.slice
}

func (s *OrderedSet[E]) Iter() iter.Seq[E] {
	return iter.Seq[E](func(yield func(E) bool) {
		for _, r := range s.slice {
			if !yield(r) {
				return
			}
		}
	})
}
