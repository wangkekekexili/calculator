package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	v := s.Pop()
	if v != 2 {
		t.Fatalf("Expected to get %v; got %v", 2, v)
	}
	v = s.Pop()
	if v != 1 {
		t.Fatalf("Expected to get %v; got %v", 1, v)
	}
	v = s.Pop()
	if v != nil {
		t.Fatalf("Expected to get nil; got %v", v)
	}
}
