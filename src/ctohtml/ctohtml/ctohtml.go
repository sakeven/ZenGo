package ctohtml

import (
	"fmt"
	"stacker/stack"
	"strconv"
)

func ToHtml(phase string) string { //text convert to HTML
	var htmlphase string
	for _, char := range phase {
		switch char {
		case '<':
			htmlphase += "&lt;"
		case '>':
			htmlphase += "&gt;"
		case '"':
			htmlphase += "&quot;"
		case ' ':
			htmlphase += "&nbsp;"
		case '&':
			htmlphase += "&amp;"
		case '©':
			htmlphase += "&copy;"
		case '®':
			htmlphase += "&reg;"
		default:
			htmlphase += string(char)
		}
	}
	return htmlphase
}

type Str []string

const opStr string = "#>[]*.^!+"

func Split(zenText string) (Str, int) {
	var (
		zenSplit Str
		phase    string
		nopflag  bool
		leve     int
	)
	for _, char := range zenText {
		nopflag = true
		for i := 0; i < len(opStr); i += 1 {
			if int(char) == int(opStr[i]) {
				if char == '>' {
					leve += 1
				} else if char == '^' {
					leve -= 1
				}
				if phase != "" {
					zenSplit = append(zenSplit, phase)
				}
				zenSplit = append(zenSplit, string(char))
				phase = ""
				nopflag = false
				break
			}
		}
		if nopflag {
			phase += string(char)
		}
	}
	if phase != "" {
		zenSplit = append(zenSplit, phase)
	}
	for i := 0; i < leve; i += 1 {
		zenSplit = append(zenSplit, "^")
	}
	return zenSplit, leve
}

func ZenHtml(tab string, zenSplit Str) (string, int) {
	var (
		zenTextHtml, back          string
		flagId, flagClass, flagMul bool
		st                         stack.Stack
		tag                        string
		recnt                      int
	)
	if zenSplit[0] == "^" {
		return "", 1
	}
	zenTextHtml = tab + "<" + zenSplit[0]
	tag = zenTextHtml
	back = tab + "</" + zenSplit[0] + ">\n"
	st.Push(back)
LOOP:
	for i := 1; i < len(zenSplit); i += 1 {
		switch zenSplit[i] {
		case "#":
			i += 1
			if flagId {
				tag += " " + zenSplit[i]
				zenTextHtml += " " + zenSplit[i]
			} else {
				zenTextHtml += " id=\"" + zenSplit[i]
				tag += " id=\"" + zenSplit[i]
				flagId = true
			}

		case ">": //child tag Recursion solve
			if flagId || flagClass {
				zenTextHtml += "\">\n"
				flagId = false
				flagClass = false
			} else {
				zenTextHtml += ">\n"
			}
			zenTmp, cntTmp := ZenHtml(tab+"\t", zenSplit[i+1:])
			zenTextHtml += zenTmp
			i += cntTmp
		case ".":
			i += 1
			if flagClass {
				tag += " " + zenSplit[i]
				zenTextHtml += " " + zenSplit[i]
			} else if flagId {
				tag += "\" class=\"" + zenSplit[i]
				zenTextHtml += "\" class=\"" + zenSplit[i]
				flagClass = true
			} else {
				flagClass = true
				tag += " class=\"" + zenSplit[i]
				zenTextHtml += " class=\"" + zenSplit[i]
			}
		case "*":
			i += 1
			flagMul = true
			cnt, _ := strconv.Atoi(zenSplit[i])
			if flagId || flagClass {
				tag += "\">"
				flagId = false
				flagClass = false
				zenTextHtml += "\">"
			} else {
				zenTextHtml += ">"
				tag += ">"
			}
			if zenSplit[i+1] == ">" {
				i += 1
				zenTmp, cntTmp := ZenHtml(tab+"\t", zenSplit[i+1:])
				zenTextHtml += "\n" + zenTmp
				tag += "\n" + zenTmp
				i += cntTmp
			}
			ele, _ := st.Top()
			if phase, ok := ele.(string); ok {
				tag += phase
				zenTextHtml += phase
			}
			st.Pop()
			for j := 0; j < cnt-1; j++ {
				zenTextHtml += tag
			}
		case "+":
			i += 1
			ele, _ := st.Top()
			if phase, ok := ele.(string); ok {
				zenTextHtml += phase
				st.Pop()
			}
			zenTextHtml += tab + "<" + zenSplit[i]
			tag = tab + "<" + zenSplit[i]
			back = tab + "</" + zenSplit[i] + ">\n"
			st.Push(back)
		case "!": //leaf node flag
			if flagMul {
				continue
			}
			if flagId || flagClass {
				zenTextHtml += "\">"
				flagId = false
				flagClass = false
			} else {
				zenTextHtml += ">"
			}

		case "^": //return to its father node
			recnt = i + 1
			break LOOP
		}
	}
	for st.IsEmpty() == false {
		ele, _ := st.Top()
		if phase, ok := ele.(string); ok {
			zenTextHtml += phase
		}
		st.Pop()
	}
	return zenTextHtml, recnt
}

func ChangeToHtml(zenText string) string {
	zenSplit, _ := Split(zenText)
	// for i, str := range zenSplit {
	// 	fmt.Println(i, str)
	// }
	zenTextHtml, _ := ZenHtml("", zenSplit)
	fmt.Println()
	return zenTextHtml
}
