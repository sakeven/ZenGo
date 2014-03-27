package phaseAnalysis

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

const opStr string = "#>[]*.^!+{},"

func Split(zenText string) Str {
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
		if nopflag && char != '\n' && char != '\r' && char != ' ' && char != '	' {
			phase += string(char)
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
