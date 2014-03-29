package main

import (
	//"fmt"
	"zentohtml"
)

func main() {
	var str zentohtml.ZenObj = `
html>
	head>
		meta*2
		>p
		^
		+style#id[type{textcs\as{da} fas},sa]
		+title
		^
	+body[style]>
		p>
			font
			^
		+div>
			p*4>
			a[href]
	`
	str.Split()
}
