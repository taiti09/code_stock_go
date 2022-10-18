// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Stack struct {
	data []int
	size int
}

func main() {
	

}

func NewStack(cap int) *Stack {
	return &Stack{
		data: make([]int, 0, cap),
		size: 0,
	}
}

func (s *Stack) push(p int) {
	s.data = append(s.data, p)
	s.size++
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		return -1
	}
	res := s.data[len(s.data)-1]
	if s.size < 2 {
		s.size = 0
		return res
	}
	s.data = s.data[:s.size-1]
	s.size = s.size - 1

	return res
}

func (s *Stack) isEmpty() bool {
	if s.size == 0 {
		return true
	}

	return false
}

func (s *Stack) front() int {
	if s.isEmpty() {
		return -1
	}
	return s.data[s.size-1]
}

func (s *Stack) string() string {
	return fmt.Sprint(s.data)
}
