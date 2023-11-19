# GoSet

GoSet is a simple set implementation in Go. It is based on the built-in map type.

> The `set` is not concurrency safe

## Installation

```bash
go get github.com/mortezaPRK/go-set

```

## Usage:

### common use cases:
```go
// Create a new set with capacity 5:
s := set.New[int](5)
// s: {}

s.Add(1, 2, 3)
// s: {1, 2, 3}

s.Remove(2)
// s: {1, 3}

s.Has(1)
// true

ns := s.Clone()
s.Remove(1)
// ns: {1, 3}
// s: {3}
```

### Helpers:

```go
set.From("a", "b", "d")
// s: {a, b, d}

s := set.FromSlice([]int{1, 2, 3})
// s: {1, 2, 3}

s.Pop()
// 1 or 2 or 3 (randomly)

s.ForEach(func(e interface{}) {
    fmt.Println(e)
})
// 3
// 1
// 2

s.ToList()
// []int{2, 3, 1}
```


### Interacting with multiple sets:

```go
// Merge multiple sets:
s_0 := set.New[int](2).Add(0, 1)
s_1 := set.New[int](2).Add(2, 3)
s_0.Merge(s_1)
// s_0: {0, 1, 2, 3}
// s_1: {2, 3}

// Check if two sets are disjoint (have no common elements):
set.From(1,2,3).IsDisjoint(set.From(4,5,6))
// true

// Check if the set is a subset of the other set:
set.From(2,4).IsSubset(set.From(1,2,4,6))
// true

// Get the intersection of multiple sets:
set.From(2,4,5).Intersection(
    set.From(1,4,6), 
    set.From(2,4), 
    set.From(1,6)
)
// {4}

// Get the difference between multiple sets: (elements in the set but not in the others)
set.From(2,4,5).Diff(
    set.From(1,4,6),
    set.From(2,4),
    set.From(1,6),
)
// {5}

// Check equality between two sets:
set.From(1,2,3).Equal(set.From(3,2,1))
// true
```
