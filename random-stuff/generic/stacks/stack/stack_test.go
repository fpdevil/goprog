package stack_test

import (
	"testing"

	"github.com/fpdevil/goprog/random-stuff/generic/stacks/stack"
)

func TestStack(t *testing.T) {
	count := 1
	var s stack.Stack
	assertTrue(t, s.Len() == 0, "expected empty Stack", count)
	count++
	assertTrue(t, s.Cap() == 0, "expected empty Stack", count)
	count++
	assertTrue(t, s.IsEmpty(), "expected empty Stack", count)
	count++
	value, err := s.Pop()
	assertTrue(t, value == nil, "expected nil value", count)
	count++
	assertTrue(t, err != nil, "expected error", count)
	count++
	value1, err := s.Top()
	assertTrue(t, value1 == nil, "expected nil value", count)
	count++
	assertTrue(t, err != nil, "expected error", count)
	count++
	s.Push(1)
	s.Push(1)
	s.Push("three")
	assertTrue(t, s.Len() == 3, "expected empty stack", count)
	count++
	assertTrue(t, s.IsEmpty() == false, "expected nonempty Stack", count)
	count++
	value2, err := s.Pop()
	assertEqualString(t, value2.(string), "three", "unexpected text", count)
	count++
	assertTrue(t, err == nil, "no error expected", count)
	count++
	value3, err := s.Top()
	assertTrue(t, value3 == 2, "unexpected number", count)
	count++
	assertTrue(t, err == nil, "no error expected", count)
	count++
	s.Pop()
	assertTrue(t, s.Len() == 1, "expected nonempty Stack", count)
	count++
	assertTrue(t, s.IsEmpty() == false, "expected nonempty Stack",
		count)
	count++
	value4, err := s.Pop()
	assertTrue(t, value4 == 1, "unexpected number", count)
	count++
	assertTrue(t, err == nil, "no error expected", count)
	count++
	assertTrue(t, s.Len() == 0, "expected empty Stack", count)
	count++
	assertTrue(t, s.IsEmpty(), "expected empty Stack", count)
	count++
}

// assertTrue calls testing.T.Error() with the provided message if the
// condition is false
func assertTrue(t *testing.T, condition bool, message string, id int) {
	if !condition {
		t.Errorf("#%d: %s", id, message)
	}
}

// assertEqualString() calls testing.T.Error() with the given message if
// the given strings are not equal.
func assertEqualString(t *testing.T, a, b string, message string, id int) {
	if a != b {
		t.Errorf("#%d: %s \"%s\" !=\n\"%s\"", id, message, a, b)
	}
}
