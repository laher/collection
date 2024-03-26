package collections

func Union[E comparable](s, s2 Collection[E]) *Set[E] {
	result := NewSet(s.Slice()...)
	result.Add(s2.Slice()...)
	return result
}

func Intersection[E comparable](s, s2 Collection[E]) *Set[E] {
	result := NewSet[E]()
	for v := range s.Iter() {
		if s2.Contains(v) {
			result.Add(v)
		}
	}
	return result
}
