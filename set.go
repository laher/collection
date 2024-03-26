package collections

import (
	"fmt"
	"iter"
)

type Set[E comparable] struct {
	items map[E]struct{}
}

// use the compiler to confirm that a set respects the Collection interface
var _ Collection[string] = &Set[string]{}

func NewSet[E comparable](vals ...E) *Set[E] {
	s := &Set[E]{}
	s.init(vals...)
	return s
}

func (s *Set[E]) init(vals ...E) {
	s.items = make(map[E]struct{}, len(vals))
	for _, v := range vals {
		s.items[v] = struct{}{}
	}
}

func (s *Set[E]) Add(vals ...E) {
	for _, v := range vals {
		s.items[v] = struct{}{}
	}
}

func (s *Set[E]) Clear() {
	clear(s.items)
}

func (s *Set[E]) Len() int {
	return len(s.items)
}

func (s *Set[E]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Set[E]) Contains(v E) bool {
	_, ok := s.items[v]
	return ok
}

func (s *Set[E]) Slice() []E {
	result := make([]E, 0, len(s.items))
	for v := range s.items {
		result = append(result, v)
	}
	return result
}

func (s *Set[E]) String() string {
	return fmt.Sprintf("%v", s.Slice())
}

func (s *Set[E]) Iter() iter.Seq[E] {
	return iter.Seq[E](func(yield func(E) bool) {
		for r := range s.items {
			if !yield(r) {
				return
			}
		}
	})
}
