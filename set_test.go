package collections

import (
	"testing"

	"github.com/matryer/is"
)

func TestSet(t *testing.T) {
	is := is.New(t)
	s := NewSet("1", "b", "c")
	is.Equal(3, s.Len())
	is.Equal(3, len(s.Slice()))
	is.True(s.Contains("b"))
	is.True(!s.Contains("2"))
}
