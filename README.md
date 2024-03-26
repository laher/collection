# Collections

[![Go Reference](https://pkg.go.dev/badge/github.com/laher/collections.svg)](https://pkg.go.dev/github.com/laher/collections)

Iterable generic collections for go v1.23+.

Experimental.

You can use these with go1.22 IF you use `GOEXPERIMENT=rangefunc`

## Example

```go
s := NewSet("1", "a", "b")
for i := range s.Iter() {
    fmt.Println(i)
}
```

## Done so far

 * Concrete Types: Set, OrderedSet
 * Interfaces: Collection
 * Funcs: Union(), Intersection()

## To maybe do

Data structures:
 * Queue, Deque
 * LinkedList
 * Map? for concurrent/immutable/sorted variants?

Variants:
 * Ordered*
 * Concurrent*
 * Immutable*
 * Sorted*
