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

// Add items to set.
func (s *Set[T]) Add(items ...T) *Set[T] {
	for i := range items {
		s.c[items[i]] = val
	}

	return s
}

// Remove items from set.
func (s *Set[T]) Remove(items ...T) *Set[T] {
	for i := range items {
		delete(s.c, items[i])
	}

	return s
}

// Check if item is in set.
func (s *Set[T]) Has(item T) bool {
	_, ok := s.c[item]

	return ok
}

// Return a shallow copy of set.
func (s *Set[T]) Clone() *Set[T] {
	out := New[T](len(s.c))

	for k := range s.c {
		out.c[k] = val
	}

	return out
}

// Merge others into the set.
func (s *Set[T]) Merge(others ...*Set[T]) *Set[T] {
	for _, other := range others {
		for k := range other.c {
			s.c[k] = val
		}
	}

	return s
}

// Return True if the set has no elements in common with other.
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

// Test whether every element in the set is in other.
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	for k := range s.c {
		if _, ok := other.c[k]; !ok {
			return false
		}
	}

	return true
}

// Return the number of elements in set.
func (s *Set[T]) Len() int {
	return len(s.c)
}

// Return a new set with elements common to the set and all others.
func (s *Set[T]) Intersection(others ...*Set[T]) *Set[T] {
	sets := append([]*Set[T]{s}, others...)
	smallest, _ := findSmallestAndBigestIndex(sets)
	needle := sets[smallest]
	copy(sets[smallest:], sets[smallest+1:])
	sets = sets[:len(sets)-1]

	out := New[T](len(needle.c))

	for k := range needle.c {
		shouldAdd := true
		for _, set := range sets {
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

// Return a new set with elements in the set that are not in the others.
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

// Remove and return an arbitrary element from the set.
func (s *Set[T]) Pop() *T {
	for k := range s.c {
		delete(s.c, k)
		return &k
	}

	return nil
}

// Return a list of items in set.
func (s *Set[T]) ToList() []T {
	out := make([]T, 0, len(s.c))
	for k := range s.c {
		out = append(out, k)
	}
	return out
}

// Call fn for each item in set.
func (s *Set[T]) ForEach(fn func(item T)) {
	for k := range s.c {
		fn(k)
	}
}

// Return a string representation of set elements.
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
