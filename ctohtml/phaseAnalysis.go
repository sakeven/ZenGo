package ctohtml

import (
	// "fmt"
	"strings"
)

func (phase zenObj) ToHtml() string { //text convert to HTML
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

func (zenText zenObj) Split() Str {
	var (
		zenSplit Str
		phase    string
		leve     int
	)
	for i := 0; i < len(zenText); i += 1 {
		char := string(zenText[i])
		if strings.Index(opStr, char) != -1 {
			if phase != "" {
				zenSplit = append(zenSplit, phase)
				phase = ""
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
	if phase != "" {
		zenSplit = append(zenSplit, phase)
	}
	for i := 0; i < leve; i += 1 {
		zenSplit = append(zenSplit, "^")
	}
	return zenSplit
}
