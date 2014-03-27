package stack

import (
	"errors"
)

func (stack Stack) Top() (string, error) {
	if stack.IsEmpty() {
		return "", errors.New("can't Top an empty stack")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() error {
	thestack := *stack
	if thestack.IsEmpty() {
		return errors.New("can't Pop an empty stack")
	}
	*stack = thestack[:len(thestack)-1]
	return nil
}
