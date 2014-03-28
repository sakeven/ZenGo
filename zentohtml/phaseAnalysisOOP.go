package zentohtml

import (
	"strings"
)

func (zenText zenObj) ToHtml() string { //text convert to HTML
	var htmlText string
	for _, char := range zenText {
		switch char {
		case '<':
			htmlText += "&lt;"
		case '>':
			htmlText += "&gt;"
		case '"':
			htmlText += "&quot;"
		case ' ':
			htmlText += "&nbsp;"
		case '&':
			htmlText += "&amp;"
		case '©':
			htmlText += "&copy;"
		case '®':
			htmlText += "&reg;"
		default:
			htmlText += string(char)
		}
	}
	return htmlText
}

func (zenText zenObj) getValue() (string, int) {
	i = 1
	cnt := 0
	phase = ""
	for {
		char = string(zenText[i])
		if char == "{" {
			cnt += 1
		} else if char == "}" {
			if cnt == 0 {
				i -= 1
				break
			}
			cnt -= 1
		}
		phase += char
		i += 1
	}
	return phase, i
}

var Arr []elemen
var arr elemen
var ele elemen

func (zenText zenObj) Split() Str {
	var (
		zenSplit Str
		zenSpl   []elemen
		phase    string
		leve     int
		flag     bool
	)
	zenText += "!"
	for i := 0; i < len(zenText); i += 1 {
		char := string(zenText[i])
		if strings.Index(opStr, char) != -1 {
			switch char {
			case ">", "#", "*", "+", "[", "^", ".", ",": //end flag
				if flag == valueFlag {
					arr.val = append(Tag.val, phase)
				} else if flag == attrFlag {
					arr.name = phase
					ele.attr = append(ele.attr, arr)
				} else if flag == eleFlag {
					ele.name = phase
					ele.flag = eleFlag
				} else if flag == mulFalg {
					arr.name = phase
					arr.flag = mulFlag
				}
				flag = nonFlag
			}
			switch char { //begin flag
			case ">", "+":
				zenSpl = append(zenSpl, ele)
				flag = eleFlag
				ele = *(new(elemen))
			case "[", ",":
				flag = attrFlag
				arr = *(new(elemen))
			case "{", ".", "#":
				flag = valueFlag
			case "*":
				flag = mulFalg
			default:
				flag = nonFlag
			}
			switch char {
			case "#":
				Tag.name = "id"
			case ".":
				Tag.name = "class"
			case "{":
				i += 1
				cnt := 0
				phase, cnt = zenText[i:].getValue()
				if flag == valueFlag {
					Tag.val = append(Tag.val, phase)
				}
				flag = nonFlag
				i += cnt
			}
			if char == "+" || char == ">" || char == "^" {
				op := elemen{name: char, flag: opFlag}
				zenSpl = append(zenSpl, op)
			}
		} else if strings.Index(illegalStr, char) == -1 {
			phase += char
		}
	}
	for i := 0; i < leve; i += 1 {
		zenSplit = append(zenSplit, "^")
	}
	return zenSplit
}
