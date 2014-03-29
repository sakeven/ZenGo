package zentohtml

import (
	"errors"
)

func (stack Stack) Len() int {
	return len(stack)
}
func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) IsEmpty() bool {
	if stack.Len() == 0 {
		return true
	}
	return false
}

func (stack *Stack) Push(x elemen) {
	*stack = append(*stack, x)
}

func (stack Stack) Top() (elemen, error) {
	if stack.IsEmpty() {
		return *(new(elemen)), errors.New("can't Top an empty stack")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (elemen, error) {
	thestack := *stack
	if thestack.IsEmpty() {
		return *(new(elemen)), errors.New("can't Pop an empty stack")
	}
	*stack = thestack[:len(thestack)-1]
	return thestack[len(thestack)-1], nil
}
