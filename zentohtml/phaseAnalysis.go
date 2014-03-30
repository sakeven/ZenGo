package zentohtml

import (
	"fmt"
	"strings"
)

func (zenText ZenObj) ToHtml() string { //text convert to HTML
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

func (zenText ZenObj) getText() (string, int) {
	cnt, phase, begChar, endChar := 0, "", "{", "}"
	for i, char := range zenText {
		schar := string(char)
		if schar == begChar {

			cnt += 1
			if cnt == 1 {
				continue
			}
		} else if schar == endChar {
			if cnt == 1 {
				cnt = i
				break
			}
			cnt -= 1
		}
		if cnt >= 1 {
			phase += schar
		}
	}
	return phase, cnt
}

func (zenText ZenObj) getValue() (string, int) {
	cnt, flag, phase, quoChar := 0, true, "", "\""
	for i, char := range zenText {
		schar := string(char)
		if schar == quoChar && flag {
			flag = false
			continue
		} else if schar == quoChar && flag == false {
			cnt = i
			break
		}
		phase += schar
	}
	return phase, cnt
}

func (zenText ZenObj) Split() eleArr {
	var (
		attr, ele elemen
		zenSpl    eleArr
		phase     string
		leve      int
		flag      int = eleFlag
	)
	count := strings.Count(string(zenText), ">") - strings.Count(string(zenText), "^")
	zenText += ZenObj(strings.Repeat("^", count+1))
	for i := 0; i < len(zenText); i += 1 {
		char := string(zenText[i])
		if strings.Index(opStr, char) != -1 {
			if strings.Index(endStr, char) != -1 && phase != "" {
				if flag == valueFlag {
					attr.val = append(attr.val, phase)
					ele.attr = append(ele.attr, attr)
				} else if flag == attrFlag {
					attr.name, attr.flag = phase, attrFlag
					ele.attr = append(ele.attr, attr)
				} else if flag == eleFlag {
					ele.name, ele.flag = phase, eleFlag
				} else if flag == mulFlag {
					ele.name, ele.flag = phase, mulFlag
				}
				flag = nonFlag
			}
			switch char { //begin flag
			case "+", ">", "*", "^":
				if ele.flag != nonFlag {
					zenSpl = append(zenSpl, ele)
				}
				if char == "*" {
					flag = mulFlag
				} else {
					flag = eleFlag
				}
				ele, attr = *(new(elemen)), *(new(elemen))
			case "[", ",":
				flag = attrFlag
				attr = *(new(elemen))
			case "#":
				flag = valueFlag
				attr.name, attr.flag = "id", attrFlag
			case ".":
				flag = valueFlag
				attr.name, attr.flag = "class", attrFlag
			case "=":
				if phase == "" {
					continue
				}
				cnt := 0
				i += 1
				attr.name, attr.flag = phase, attrFlag
				phase, cnt = zenText[i:].getValue()
				attr.val = append(attr.val, phase)
				ele.attr = append(ele.attr, attr)
				flag = nonFlag
				i += cnt
			case "{":
				cnt := 0
				ele.flag = textFlag
				phase, cnt = zenText[i:].getText()
				ele.name = phase
				flag = nonFlag
				i += cnt
				fmt.Println("sad", string(zenText[i]))
			default:
				flag = nonFlag
			}
			phase = ""
			if strings.Index(illegalOp, char) == -1 {
				zenSpl = append(zenSpl, elemen{name: char, flag: opFlag})
			}
		} else if strings.Index(illegalStr, char) == -1 {
			phase += char
		}
	}
	for i := 0; i < leve; i += 1 {
		zenSpl = append(zenSpl, elemen{name: "^", flag: opFlag})
	}
	for _, fs := range zenSpl {
		fmt.Println(fs.name)
	}
	return zenSpl
}
