package collection

import (
	"testing"

	"github.com/matryer/is"
)

func TestOrdered(t *testing.T) {
	s := NewOrderedSet("1", "b", "c")
	is := is.New(t)
	is.True(s.Contains("b"))
	is.True(!s.Contains("2"))
	is.Equal(3, s.Len())
	is.Equal([]string{"1", "b", "c"}, s.Slice())
}
