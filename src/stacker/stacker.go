package main

import (
	"fmt"
	"stacker/stack"
)

func main() {
	var haystack stack.Stack
	haystack.Push("hay")
	haystack.Push("sda")
	haystack.Push("asd")
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
