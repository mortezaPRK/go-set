// Package set provides a Set data structure.
package set

import (
	"fmt"
	"strings"
)

type valType struct{}

var val valType

type Set[T comparable] struct {
	c map[T]valType
}

func New[T comparable](size int) *Set[T] {
	return &Set[T]{c: make(map[T]valType, size)}
}

func From[T comparable](sl ...T) *Set[T] {
	return New[T](len(sl)).Add(sl...)
}

func FromSlice[T comparable](sl []T) *Set[T] {
	return New[T](len(sl)).Add(sl...)
}

// Add adds items to set.
func (s *Set[T]) Add(items ...T) *Set[T] {
	for i := range items {
		s.c[items[i]] = val
	}

	return s
}

// Remove deletes items from set.
func (s *Set[T]) Remove(items ...T) *Set[T] {
	for i := range items {
		delete(s.c, items[i])
	}

	return s
}

// Has returns true if item is in set.
func (s *Set[T]) Has(item T) bool {
	_, ok := s.c[item]

	return ok
}

// Clone creates a shallow copy of the set.
func (s *Set[T]) Clone() *Set[T] {
	out := New[T](len(s.c))

	for k := range s.c {
		out.c[k] = val
	}

	return out
}

// Merge merges all sets into current set.
func (s *Set[T]) Merge(others ...*Set[T]) *Set[T] {
	for _, other := range others {
		for k := range other.c {
			s.c[k] = val
		}
	}

	return s
}

// IsDisjoint returns true if set has no elements in common with other.
func (s *Set[T]) IsDisjoint(other *Set[T]) bool {
	smaller, bigger := s.c, other.c
	if len(s.c) > len(other.c) {
		smaller, bigger = bigger, smaller
	}

	for k := range smaller {
		if _, ok := bigger[k]; ok {
			return false
		}
	}

	return true
}

// IsSubset returns true if every element in the set is in other.
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	for k := range s.c {
		if _, ok := other.c[k]; !ok {
			return false
		}
	}

	return true
}

// Len returns the number of elements in set.
func (s *Set[T]) Len() int {
	return len(s.c)
}

// Intersection return a new set with elements common to the set and all others.
func (s *Set[T]) Intersection(others ...*Set[T]) *Set[T] {
	others = append(others, s)
	smallest, _ := findSmallestAndBigestIndex(others)
	needle := others[smallest]
	copy(others[smallest:], others[smallest+1:])
	others = others[:len(others)-1]

	out := New[T](len(needle.c))

	for k := range needle.c {
		shouldAdd := true
		for _, set := range others {
			if _, ok := set.c[k]; !ok {
				shouldAdd = false
				break
			}
		}

		if shouldAdd {
			out.c[k] = val
		}
	}

	return out
}

// Diff return a new set with elements in the set that are not in the others.
func (s *Set[T]) Diff(others ...*Set[T]) *Set[T] {
	out := New[T](len(s.c))

	for k := range s.c {
		shouldAdd := true
		for _, other := range others {
			if _, ok := other.c[k]; ok {
				shouldAdd = false
				break
			}
		}
		if shouldAdd {
			out.c[k] = val
		}
	}

	return out
}

// Pop remove and return an arbitrary element from the set.
func (s *Set[T]) Pop() *T {
	for k := range s.c {
		delete(s.c, k)
		return &k
	}

	return nil
}

// ToList returns a list of items in set.
func (s *Set[T]) ToList() []T {
	out := make([]T, 0, len(s.c))
	for k := range s.c {
		out = append(out, k)
	}
	return out
}

// ForEach calls fn for each item in set.
func (s *Set[T]) ForEach(fn func(item T)) {
	for k := range s.c {
		fn(k)
	}
}

// Equal returns true if set is equal to other.
func (s *Set[T]) Equal(other *Set[T]) bool {
	if len(s.c) != len(other.c) {
		return false
	}

	for k := range s.c {
		if _, ok := other.c[k]; !ok {
			return false
		}
	}

	return true
}

// String returns a string representation of set elements.
func (s *Set[T]) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("[%d]{", len(s.c)))
	loopDelimiter := ""
	for i := range s.c {
		sb.WriteString(loopDelimiter)
		loopDelimiter = ","
		sb.WriteString(fmt.Sprint(i))
	}
	sb.WriteString("}")
	return sb.String()
}
