package zentohtml

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ZenHtml(tab string, zenSpl eleArr) (string, int) {
	var (
		zenTextHtml string
		st          Stack
		tag         string
		recnt       int
	)
	if zenSpl[0].name == "^" {
		return "", 1
	}
	if zenSpl[0].flag == eleFlag {
		tag = tab + zenSpl[0].RepPreelement()
		st.Push(zenSpl[0])
	}
LOOP:
	for i := 1; i < len(zenSpl); i++ {
		switch zenSpl[i].name {
		case ">": //child tag Recursion solve
			if zenSpl[i-1].flag == eleFlag {
				zenTextHtml += tag
			}
			zenTmp, cntTmp := ZenHtml(tab+"\t", zenSpl[i+1:])
			zenTextHtml += "\n" + zenTmp
			i += cntTmp
		case "*": //repeat case
			i += 1
			rtab := ""
			count, _ := strconv.Atoi(zenSpl[i].name)
			if zenSpl[i+1].name == ">" {
				zenTmp, cntTmp := ZenHtml(tab+"\t", zenSpl[i+2:])
				tag += "\n" + zenTmp
				i += cntTmp + 1
				rtab = tab
			}
			if ele, err := st.Pop(); err == nil {
				tag += rtab + ele.RepBackelement() + "\n"
			}
			zenTextHtml += strings.Repeat(tag, count)
		case "+": //brother case
			rtab := tab
			if zenSpl[i-1].flag == eleFlag {
				zenTextHtml += tag
				rtab = ""
			}
			if ele, err := st.Pop(); err == nil {
				zenTextHtml += rtab + ele.RepBackelement() + "\n"
			}
			i += 1
			if zenSpl[i].flag == eleFlag {
				tag = tab + zenSpl[i].RepPreelement()
				st.Push(zenSpl[i])
			}
		case "^": //return to its father node
			if zenSpl[i-1].flag == eleFlag {
				zenTextHtml += tag
				if ele, err := st.Pop(); err == nil {
					zenTextHtml += ele.RepBackelement() + "\n"
				}
			}
			recnt = i + 1
			break LOOP
		}
	}
	if st.IsEmpty() == false {
		ele, _ := st.Pop()
		zenTextHtml += tab + ele.RepBackelement() + "\n"
	}
	return zenTextHtml, recnt
}

func (zenText ZenObj) ChangeToHtml() string {
	zenSpl := zenText.Split()
	zenTextHtml, _ := ZenHtml("", zenSpl)
	return zenTextHtml
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
