package zentohtml

// import (
// 	"errors"
// 	//"fmt"
// )

func (e elemen) Typeof() string {
	switch e.flag {
	case textFlag:
		return "Text"
	case opFlag:
		return "Operation"
	case attrFlag:
		return "Attribute"
	case eleFlag:
		return "Element"
	case valueFlag:
		return "Value"
	case mulFlag:
		return "Repeate"
	}
	return ""
}

func (e elemen) RepAttr() (attr string) {
	if e.flag == attrFlag {
		attr = e.name + "=\""
		for _, value := range e.val {
			attr += value + " "
		}
		attr += "\""
	}
	return attr
}

func (e elemen) RepPreelement() (tag string) {
	if e.flag == eleFlag {
		tag = "<" + e.name
		for _, attr := range e.attr {
			str := attr.RepAttr()
			tag += " " + str
		}
		tag += ">"
	}
	return tag
}

func (e elemen) RepBackelement() (tag string) {
	if e.flag == eleFlag {
		tag = "</" + e.name + ">"
	}
	return tag
}
