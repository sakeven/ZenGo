package main

import (
	"ctohtml/ctohtml"
	"fmt"
)

func main() {
	var str string
	str = "div#dt#name.item.iteam>a>p.iteam>li>a.href*3>p!"
	str = ctohtml.ChangeToHtml(str)
	fmt.Println(str)
	str = "pt>div#dt#name.item.iteam>p.iteam>a*2!^^+h>div*3!^^+a.href!"
	str = ctohtml.ChangeToHtml(str)
	fmt.Println(str)
	str = "div!"
	str = ctohtml.ChangeToHtml(str)
	fmt.Println(str)
}
