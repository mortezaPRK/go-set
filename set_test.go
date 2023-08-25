package set_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mortezaPRK/go-set"
	"github.com/stretchr/testify/require"
)

func TestFromSlice(t *testing.T) {
	t.Parallel()

	require.ElementsMatch(t, set.FromSlice([]int{1, 2, 3}).ToList(), set.New[int](0).Add(1, 2, 3).ToList())
}

func TestAdd(t *testing.T) {
	t.Parallel()

	s := set.New[int](0)

	ns := s.Add(1, 1, 2, 3, 4, 4, 5)

	require.ElementsMatch(t, s.ToList(), []int{1, 2, 3, 4, 5})
	require.Same(t, s, ns)
}

func TestRemove(t *testing.T) {
	t.Parallel()

	type matrix struct {
		source   *set.Set[int]
		needle   int
		expected *set.Set[int]
	}

	for _, tc := range []matrix{
		{
			source:   set.New[int](0).Add(1, 2),
			needle:   2,
			expected: set.New[int](0).Add(1),
		},
		{
			source:   set.New[int](0).Add(1, 2),
			needle:   3,
			expected: set.New[int](0).Add(1, 2),
		},
	} {
		tc := tc
		name := fmt.Sprintf("remvoing %d from %s should result in %s", tc.needle, toString(tc.source), toString(tc.expected))
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			final := tc.source.Remove(tc.needle)

			require.Same(t, tc.source, final)
			require.ElementsMatch(t, tc.source.ToList(), tc.expected.ToList())
		})
	}
}

func TestHas(t *testing.T) {
	t.Parallel()

	type matrix struct {
		source   *set.Set[int]
		needle   int
		expected bool
	}

	for _, tc := range []matrix{
		{
			source:   set.New[int](0).Add(1, 2),
			needle:   2,
			expected: true,
		},
		{
			source:   set.New[int](0).Add(1, 2),
			needle:   3,
			expected: false,
		},
	} {
		tc := tc
		name := fmt.Sprintf("%s has %d: %t", toString(tc.source), tc.needle, tc.expected)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			has := tc.source.Has(tc.needle)

			require.Equal(t, tc.expected, has)
		})
	}
}

func TestClone(t *testing.T) {
	t.Parallel()

	s := set.New[int](0).Add(1, 2, 3, 4, 5)

	ns := s.Clone()

	require.NotSame(t, s, ns)
	require.Equal(t, s.Len(), ns.Len())
	require.ElementsMatch(t, s.ToList(), ns.ToList())
}

func TestMerge(t *testing.T) {
	t.Parallel()

	s_0 := set.New[int](0).Add(0, 1)
	s_1 := set.New[int](0).Add(2, 3)
	s_2 := set.New[int](0).Add(2, 3, 4)
	s_3 := set.New[int](0).Add(5, 6)

	ns := s_0.Merge(s_1, s_2, s_3)

	require.Same(t, s_0, ns)
	require.ElementsMatch(t, s_0.ToList(), []int{0, 1, 2, 3, 4, 5, 6})
}

func TestIsDisjoint(t *testing.T) {
	t.Parallel()

	type matrix struct {
		source   *set.Set[int]
		target   *set.Set[int]
		expected bool
	}

	for _, tc := range []matrix{
		{
			source:   set.New[int](0).Add(2, 4),
			target:   set.New[int](0).Add(2, 4),
			expected: false,
		},
		{
			source:   set.New[int](0).Add(2, 4),
			target:   set.New[int](0).Add(2, 5),
			expected: false,
		},
		{
			source:   set.New[int](0).Add(2, 4),
			target:   set.New[int](0).Add(1, 3),
			expected: true,
		},
		{
			source:   set.New[int](0).Add(2, 4),
			target:   set.New[int](0),
			expected: true,
		},
	} {
		tc := tc
		name := fmt.Sprintf("%s and %s has no common element: %t", toString(tc.source), toString(tc.target), tc.expected)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			isDisjoint := tc.source.IsDisjoint(tc.target)

			require.Equal(t, tc.expected, isDisjoint)
		})
	}
}

func TestIsSubset(t *testing.T) {
	t.Parallel()

	type matrix struct {
		source   *set.Set[int]
		target   *set.Set[int]
		expected bool
	}

	for _, tc := range []matrix{
		{
			source:   set.New[int](0).Add(2, 4),
			target:   set.New[int](0).Add(2, 4),
			expected: true,
		},
		{
			source:   set.New[int](0).Add(2, 4),
			target:   set.New[int](0).Add(2, 3, 4),
			expected: true,
		},
		{
			source:   set.New[int](0).Add(2, 4),
			target:   set.New[int](0).Add(2, 3),
			expected: false,
		},
		{
			source:   set.New[int](0).Add(2, 4),
			target:   set.New[int](0).Add(1, 3),
			expected: false,
		},
	} {
		tc := tc
		name := fmt.Sprintf("%s is subset of %s: %t", toString(tc.source), toString(tc.target), tc.expected)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			isSubset := tc.source.IsSubset(tc.target)

			require.Equal(t, tc.expected, isSubset)
		})
	}
}

func TestLen(t *testing.T) {
	t.Parallel()

	type matrix struct {
		source      *set.Set[int]
		expectedLen int
	}

	for _, tc := range []matrix{
		{
			source:      set.New[int](0),
			expectedLen: 0,
		},
		{
			source:      set.New[int](0).Add(1, 2, 3),
			expectedLen: 3,
		},
	} {
		tc := tc
		name := fmt.Sprintf("length of %s should be %d", toString(tc.source), tc.expectedLen)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			length := tc.source.Len()

			require.Equal(t, tc.expectedLen, length)
		})
	}
}

func TestIntersection(t *testing.T) {
	t.Parallel()

	type matrix struct {
		source   *set.Set[int]
		targets  []*set.Set[int]
		expected []int
	}

	for _, tc := range []matrix{
		{
			source: set.New[int](0).Add(1, 2, 3),
			targets: []*set.Set[int]{
				set.New[int](0).Add(3),
				set.New[int](0).Add(1, 3),
				set.New[int](0).Add(2, 4),
			},
			expected: []int{},
		},
		{
			source: set.New[int](0).Add(1, 2, 3),
			targets: []*set.Set[int]{
				set.New[int](0).Add(1, 2, 3),
			},
			expected: []int{1, 2, 3},
		},
		{
			source: set.New[int](0).Add(1, 2, 3),
			targets: []*set.Set[int]{
				set.New[int](0).Add(1, 2),
				set.New[int](0).Add(1, 3),
				set.New[int](0).Add(1, 4, 5),
				set.New[int](0).Add(1, 4, 5, 6),
				set.New[int](0).Add(1, 4, 5, 6, 7),
			},
			expected: []int{1},
		},
	} {
		tc := tc
		name := fmt.Sprintf("intersecting %s with %s", toString(tc.source), toString(tc.targets...))
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			source := tc.source.Clone()

			intersection := tc.source.Intersection(tc.targets...)

			require.NotSame(t, intersection, tc.source)
			require.ElementsMatch(t, intersection.ToList(), tc.expected)
			require.ElementsMatch(t, tc.source.ToList(), source.ToList())
		})
	}
}

func TestDiff(t *testing.T) {
	t.Parallel()

	type matrix struct {
		source   *set.Set[int]
		targets  []*set.Set[int]
		expected []int
	}

	for _, tc := range []matrix{
		{
			source: set.New[int](0).Add(1, 2, 3),
			targets: []*set.Set[int]{
				set.New[int](0).Add(3),
				set.New[int](0).Add(1),
				set.New[int](0).Add(2, 1),
			},
			expected: []int{},
		},
		{
			source: set.New[int](0).Add(1, 2, 3),
			targets: []*set.Set[int]{
				set.New[int](0).Add(1, 2, 3),
			},
			expected: []int{},
		},
		{
			source: set.New[int](0).Add(1, 2, 3),
			targets: []*set.Set[int]{
				set.New[int](0).Add(1, 4, 5),
				set.New[int](0).Add(3, 6, 7),
				set.New[int](0).Add(4),
			},
			expected: []int{2},
		},
	} {
		tc := tc
		name := fmt.Sprintf("difference between %s and %s", toString(tc.source), toString(tc.targets...))
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			source := tc.source.Clone()

			difference := tc.source.Diff(tc.targets...)

			require.NotSame(t, difference, tc.source)
			require.ElementsMatch(t, tc.source.ToList(), source.ToList())
			require.ElementsMatch(t, difference.ToList(), tc.expected)
		})
	}
}

func TestPop(t *testing.T) {
	t.Parallel()

	s := set.New[int](0).Add(1, 2)

	element := s.Pop()
	require.NotNil(t, element)
	require.Contains(t, []int{1, 2}, *element)

	anotherElement := s.Pop()
	require.NotNil(t, anotherElement)
	require.Contains(t, []int{1, 2}, *anotherElement)

	require.NotEqual(t, *element, *anotherElement)

	nilElement := s.Pop()
	require.Nil(t, nilElement)
}

func TestToList(t *testing.T) {
	t.Parallel()

	s := set.New[int](0).Add(1, 2, 4)

	elements := s.ToList()

	require.ElementsMatch(t, []int{1, 2, 4}, elements)
}

func toString(sets ...*set.Set[int]) string {
	allItems := make([]string, 0, len(sets))
	for _, s := range sets {
		items := make([]string, 0, s.Len())
		for i := range s.ToList() {
			items = append(items, fmt.Sprintf("%d", i))
		}
		allItems = append(allItems, "["+strings.Join(items, ",")+"]")
	}
	return strings.Join(allItems, ",")
}
