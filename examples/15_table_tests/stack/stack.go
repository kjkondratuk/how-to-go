package stack

type stack struct {
	elements []interface{}
}

type Stack interface {
	Push(val interface{})
	Pop() interface{}
}

func NewStack() Stack {
	return &stack{
		elements: make([]interface{}, 0),
	}
}

func (s *stack) Push(val interface{}) {
	s.elements = append(s.elements, val)
}

func (s *stack) Pop() interface{} {
	v := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return v
}
