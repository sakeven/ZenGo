package stack

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
