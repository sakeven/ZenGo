package ctohtml

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
var Tag elemen

func (zenText zenObj) Split() Str {
	var (
		zenSplit          Str
		zenSpl            []elemen
		phase             string
		leve              int
		eleflag, attrflag bool
	)
	zenText += "!"
	for i := 0; i < len(zenText); i += 1 {
		char := string(zenText[i])
		if strings.Index(opStr, char) != -1 {
			switch char {
			case "[":
				attrflag = true
			case ">":
				zenSpl[len(zenSpl)-1].attr = Arr
				Arr = Arr[0:0]
				eleflag = true
				leve += 1
			case "^":
				leve -= 1
			case "{":
				i += 1
				cnt := 0
				phase, cnt = zenText[i:].getValue()
			}
			if phase != "" {
				if eleflag {
					Tag.name = phase
					Tag.flag = eleFlag
					Arr = append(Arr, Tag)
					eleflag = false
				} else if attrflag {
					Tag.name = phase
					Tag.flag = attrFlag
					attrflag = false
				}
				zenSplit = append(zenSplit, phase)
				phase = ""
			}
			if char == "^" || (zenSplit[len(zenSplit)-1] != "^" && char == "+") {
				zenSplit = append(zenSplit, "!")
			}
			zenSplit = append(zenSplit, char)

		} else if strings.Index(illegalStr, char) == -1 {
			phase += char
		}
	}
	for i := 0; i < leve; i += 1 {
		zenSplit = append(zenSplit, "^")
	}
	return zenSplit
}
