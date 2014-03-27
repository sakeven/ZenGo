package ctohtml

import (
	"errors"
)

type Str []string

const opStr string = "#>[]*.^!+{},"

type Stack []string

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

func (stack *Stack) Push(x string) {
	*stack = append(*stack, x)
}

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
