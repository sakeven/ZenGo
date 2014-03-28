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

func (zenText zenObj) Split() Str {
	var (
		zenSplit Str
		phase    string
		leve     int
	)
	zenText += "!"
	for i := 0; i < len(zenText); i += 1 {
		char := string(zenText[i])
		if strings.Index(opStr, char) != -1 {
			if phase != "" {
				zenSplit = append(zenSplit, phase)
				phase = ""
			}
			if char == "^" || (zenSplit[len(zenSplit)-1] != "^" && char == "+") {
				zenSplit = append(zenSplit, "!")
			}
			zenSplit = append(zenSplit, char)
			switch char {
			case ">":
				leve += 1
			case "^":
				leve -= 1
			case "{":
				i += 1
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
