package main

import (
	"zentohtml"
)

func main() {
	var str zentohtml.ZenObj = `html>
	head>
		meta*2
		+style#id[type{textcs\as{da} fas},sa]
		+title
		^
	+body[style]>
		p>
			font
			^
		+div>
			p*4>
			a[href]^`
	str.Split()
}
