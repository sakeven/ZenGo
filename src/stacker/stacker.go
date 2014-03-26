package main

import (
	"fmt"
	"stacker/stack"
)

func main() {
	var haystack stack.Stack
	haystack.Push("hay")
	haystack.Push(1)
	haystack.Push(82.3)
	fmt.Println(haystack.Cap())
	for {
		item, err := haystack.Top()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(item)
		haystack.Pop()
	}
}
