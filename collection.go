package collection

import "iter"

type Collection[E comparable] interface {
	Add(...E)
	Clear()
	Contains(E) bool
	Iter() iter.Seq[E]
	Len() int
	IsEmpty() bool
	Slice() []E
}
