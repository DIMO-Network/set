package set

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	var s Set[int]

	if !s.Add(1) {
		t.Errorf("the value 1 was not in the set, but Add returned false")
	}
	if !s.Add(5) {
		t.Errorf("the value 5 was not in the set, but Add returned false")
	}
	if s.Add(1) {
		t.Errorf("the value 1 was already in the set, but Add returned true")
	}
}

func TestContains(t *testing.T) {
	var s Set[int]
	s.Add(1)
	s.Add(5)

	if !s.Contains(1) {
		t.Errorf("set should have reported containing 1")
	}
	if !s.Contains(5) {
		t.Errorf("set should have reported containing 5")
	}
	if s.Contains(8) {
		t.Errorf("set erroneously reported containing 8")
	}
}

func TestDelete(t *testing.T) {
	var s Set[int]
	s.Add(1)
	s.Add(5)
	s.Delete(5)

	if s.Contains(5) {
		t.Errorf("set reports containing 5, but it should have been removed")
	}
	if !s.Contains(1) {
		t.Errorf("1 has been added and not removed, but the set is claiming to not contain it")
	}
}

func TestConstruction(t *testing.T) {
	s := Of(5, 6)

	if !s.Contains(5) {
		t.Errorf("set was initialized with 5, but reports not containing 5")
	}
	if !s.Contains(6) {
		t.Errorf("set was initialized with 6, but reports not containing 6")
	}
}

func TestAll(t *testing.T) {
	s := Of(5, 6, 1)

	out := slices.Sorted(s.All())
	expect := []int{1, 5, 6}

	if !slices.Equal(expect, out) {
		t.Errorf("expected iterator to run through %v, but we got %v", expect, out)
	}
}

func TestUnion(t *testing.T) {
	s1 := Of(1, 4, -2)
	s2 := Of(-6, 1)

	u := Union(&s1, &s2)

	if u.Len() != 4 {
		t.Errorf("expected 4 elements in the union, but got %d", u.Len())
	}

	if !u.Contains(1) {
		t.Errorf("union should contain 1")
	}
	if !u.Contains(-6) {
		t.Errorf("union should contain -6")
	}
}

func TestStringer(t *testing.T) {
	s1 := Of(1, 4, -2)
	// var s2 Set[int]

	str := s1.String()

	rightPrefix := strings.HasPrefix(str, "set[")
	if !rightPrefix {
		t.Errorf("expected \"set[\" at the start of string %q", str)
	}

	rightSuffix := strings.HasSuffix(str, "]")
	if !rightSuffix {
		t.Errorf("expected \"]\" at the end of %q", str)
	}

	if rightPrefix && rightSuffix {
		var temp []int
		middle := str[4 : len(str)-1]
		fmt.Println("XD", middle)
		for i, item := range strings.Split(middle, " ") {
			integ, err := strconv.Atoi(item)
			if err != nil {
				t.Fatalf("couldn't parse index %d element %q as an integer: %s", i, item, err)
			} else {
				temp = append(temp, integ)
			}
		}

		if len(temp) != 3 {
			t.Errorf("should have 3 elements, but string representation showed %d", len(temp))
		}

		expectedEls := []int{-2, 1, 4}

		var strMissing []int
		for _, el := range expectedEls {
			if !slices.Contains(temp, el) {
				strMissing = append(strMissing, el)
			}
		}
		if len(strMissing) != 0 {
			t.Errorf("string representation missing elements %s", fmt.Sprint(strMissing))
		}

		var strExtra []int
		for _, el := range expectedEls {
			if !slices.Contains(expectedEls, el) {
				strExtra = append(strExtra, el)
			}
		}
		if len(strExtra) != 0 {
			t.Errorf("string representation contained extra elements %s", fmt.Sprint(strExtra))
		}
	}
}
