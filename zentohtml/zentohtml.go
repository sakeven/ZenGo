package zentohtml

import (
	"bufio"
	"io"
	"strconv"
)

func attrCase(zenSplit Str) (string, int) {
	var strTmp string
	i := -1
LOOP:
	for {
		i++
		switch zenSplit[i] {
		case "{":
			i += 1
			strTmp += zenSplit[i]
		case ",":
			strTmp += "\""
		case "]":
			strTmp += "\""
			break LOOP
		case "}":
		default:
			strTmp += " " + zenSplit[i] + "=\""
		}
	}
	return strTmp, i + 1
}

func ZenHtml(tab string, zenSplit Str) (string, int) {
	var (
		zenTextHtml, back          string
		flagId, flagClass, flagMul bool
		st                         Stack
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
	for i := 1; i < len(zenSplit); i++ {
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
			phase, _ := st.Pop()
			tag += phase
			zenTextHtml += phase
			for j := 0; j < cnt-1; j++ {
				zenTextHtml += tag
			}
		case "+":
			i += 1
			phase, _ := st.Pop()
			zenTextHtml += phase
			zenTextHtml += tab + "<" + zenSplit[i]
			tag = tab + "<" + zenSplit[i]
			back = tab + "</" + zenSplit[i] + ">\n"
			st.Push(back)
		case "!": //leaf node flag
			if flagMul {
				flagMul = false
				continue
			}
			if flagId || flagClass {
				zenTextHtml += "\">"
				flagId = false
				flagClass = false
			} else {
				zenTextHtml += ">"
			}
		case "[":
			if flagId || flagClass {
				tag += "\""
				flagId = false
				flagClass = false
				zenTextHtml += "\""
			}
			strTmp, cntTmp := attrCase(zenSplit[i+1:])
			zenTextHtml += strTmp
			tag += strTmp
			i += cntTmp
		case "^": //return to its father node
			recnt = i + 1
			break LOOP
		}
	}
	for st.IsEmpty() == false {
		phase, _ := st.Pop()
		zenTextHtml += phase
	}
	return zenTextHtml, recnt
}

func (zenText ZenObj) ChangeToHtml() string {
	//zenSplit := zenText.Split()
	//zenTextHtml, _ := ZenHtml("", zenSplit)
	return "" //zenTextHtml
}

func FileToHtml(inFile io.Reader, outFile io.Writer) (err error) {
	var zenText string
	ifp := bufio.NewReader(inFile)
	ofp := bufio.NewWriter(outFile)
	defer func() {
		if err == nil {
			err = ofp.Flush()
		}
	}()
	for {
		line, ok := ifp.ReadString('\n')
		zenText += line
		if ok != nil {
			break
		}
	}
	zenTextHtml := ZenObj(zenText).ChangeToHtml()
	ofp.WriteString(zenTextHtml)
	return nil
}
