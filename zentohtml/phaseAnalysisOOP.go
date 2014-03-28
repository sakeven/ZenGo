package zentohtml

import (
	"strings"
)

func (zenText zenObj) ToHtml() string { //text convert to HTML
	var htmlText string
	m := map[rune]string{'<': "&lt;", '>': "&gt;", '"': "&quot;", ' ': "&nbsp;", '&': "&amp;", '©': "&copy;", '®': "&reg;"}
	for _, char := range zenText {
		if esc, ok := m[char]; ok {
			htmlText += esc
		} else {
			htmlText += string(char)
		}
	}
	return htmlText
}

func (zenText zenObj) getValue() (string, int) {
	cnt, phase := 0, ""
	for i, char := range zenText {
		schar := string(char)
		if schar == "{" {
			cnt += 1
		} else if schar == "}" {
			if cnt == 1 {
				cnt = i
				break
			}
			cnt -= 1
		}
		phase += schar
	}
	return phase, cnt
}

func (zenText zenObj) Split() eleArr {
	var (
		zenSplit               Str
		zenSpl, attr, ele, num eleArr
		phase                  string
		leve                   int
		flag                   bool
	)
	for i := 0; i < len(zenText); i += 1 {
		char := string(zenText[i])
		if strings.Index(opStr, char) != -1 {
			if strings.Index(endStr, char) != -1 {
				if flag == valueFlag {
					attr.val = append(Tag.val, phase)
				} else if flag == attrFlag {
					attr.name = phase
					ele.attr = append(ele.attr, arr)
				} else if flag == eleFlag {
					ele.name = phase
					ele.flag = eleFlag
				} else if flag == mulFalg {
					ele.name = phase
					ele.flag = mulFlag
				}
				flag = nonFlag
			}
			switch char { //begin flag
			case "+", "^", "*":
				zenSpl = append(zenSpl, ele)
				flag = eleFlag
				ele = *(new(elemen))
			case "[", ",":
				flag = attrFlag
				arr = *(new(elemen))
			case "#":
				flag = valueFlag
				Tag.name = "id"
			case ".":
				flag = valueFlag
				Tag.name = "class"
			case "{":
				cnt := 0
				phase, cnt = zenText[i:].getValue()
				arr.val = append(arr.val, phase)
				flag = nonFlag
				i += cnt
			case "*":
				flag = mulFalg
			default:
				flag = nonFlag
			}
			if char == "+" || char == ">" || char == "^" || char == "*" {
				op := elemen{name: char, flag: opFlag}
				zenSpl = append(zenSpl, op)
			}
		} else if strings.Index(illegalStr, char) == -1 {
			phase += char
		}
	}
	for i := 0; i < leve; i += 1 {
		zenSpl = append(zenSpl, elemen{name: "^", flag: opFlag})
	}
	return zenSpl
}
