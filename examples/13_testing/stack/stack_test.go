package stack

import "testing"

func TestNewStack(t *testing.T) {
	// should create a new stack
	s := NewStack()

	if s == nil {
		t.Errorf("The stack should be created when the constructor is called.")
	}
}

func TestStack_Pop_Push(t *testing.T) {
	s := NewStack()

	// should panic when there is nothing to pop, to make sure it's created empty
	assertPanic(t, func() {
		s.Pop()
	})

	// Should pop single item
	s.Push("some value")
	s1 := s.Pop().(string)
	if s1 != "some value" {
		t.Errorf("The popped item should be the same one pushed for a single element.")
	}

	// If multiple items are pushed, the last one should be popped first
	s.Push("first element")
	s.Push("second element")
	s.Push("third element")

	s2 := s.Pop().(string)
	if s2 != "third element" {
		t.Errorf("The last element put on the stack should be the first off.")
	}
}

// assertPanic: recovers from a panic with an error message if there was no panic
func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Should have panicked!")
		}
	}()
	f()
}
