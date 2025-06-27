package set

import (
	"fmt"
	"iter"
	"maps"
	"slices"
)

// Set is a finite set of elements of some comparable type.
//
// The zero value of Set is safe to use. However, concurrent use is not
// safe.
type Set[E comparable] struct {
	m map[E]struct{}
}

// OF returns a set consisting of the given list of elements.
func Of[E comparable](v ...E) Set[E] {
	s := Set[E]{m: make(map[E]struct{})}
	for _, a := range v {
		s.m[a] = struct{}{}
	}
	return s
}

// Add adds an element to set. It reports whether adding the element
// made the set any larger.
func (s *Set[E]) Add(x E) bool {
	if s.m == nil {
		s.m = make(map[E]struct{})
	}
	_, ok := s.m[x]
	s.m[x] = struct{}{}
	return !ok
}

// Contains returns whether an element is in the set.
func (s *Set[E]) Contains(x E) bool {
	_, ok := s.m[x]
	return ok
}

// Delete removes an element from the set. It reports whether the
// element was actually in the set.
func (s *Set[E]) Delete(x E) bool {
	_, ok := s.m[x]
	delete(s.m, x)
	return ok
}

// Len returns the number of elements in the set.
func (s *Set[E]) Len() int {
	return len(s.m)
}

// String returns a human-readable representation of the set's
// elements.
func (s *Set[E]) String() string {
	return fmt.Sprintf("set%s", fmt.Sprint(slices.Collect(s.All())))
}

// All returns an iterator over the elements of the set. As for
// built-in maps, the order is unpredictable.
func (s *Set[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for x := range s.m {
			if !yield(x) {
				return
			}
		}
	}
}

// Clone returns a copy of the set. The elements are copied in a
// shallow way, by assignment.
func (s *Set[E]) Clone() Set[E] {
	return Set[E]{m: maps.Clone(s.m)}
}

// Union returns a new set containing all of the elements in the two
// given sets. Later changes to membership of the these sets will not
// affect the union.
func Union[A comparable](s1, s2 *Set[A]) Set[A] {
	var s Set[A]
	for x := range s1.m {
		s.Add(x)
	}
	for x := range s2.m {
		s.Add(x)
	}
	return s
}
