# Collections

Iterable generic collections for go v1.23+.

Experimental.

You can use these with go1.22 IF you use `GOEXPERIMENT=rangefunc`


## Set

```go
s := NewSet("1", "a", "b")
for i := range s.Iter() {
    fmt.Println(i)
}
```

## To maybe do

Data structures:
 * Queue, Deque
 * LinkedList
 * Map? for concurrent/immutable/sorted variants?

Variants:
 * Concurrent*
 * Immutable*
 * Sorted*
