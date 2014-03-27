package main

import (
	"fmt"
)

func main() {
	str := ">paos"
	s := ""
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '>', '^', '!', '{', '}', '[', ']', '+':
			s += "1"
			fallthrough
		case '>':
			s += "2"
		}
	}
	fmt.Println(s)
}
