package collection

import (
	"testing"

	"github.com/matryer/is"
)

func TestUnion(t *testing.T) {
	a := NewSet(1)
	b := NewOrderedSet(5, 3, 6)
	is := is.New(t)

	x := Union[int](a, b)
	is.Equal([]int{1, 5, 3, 6}, x.Slice())
}
