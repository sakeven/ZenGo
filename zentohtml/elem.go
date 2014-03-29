package zentohtml

import (
	"errors"
	//"fmt"
)

func (e elemen) typeof() string {
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

func (e elemen) repAttr() (attr string, err error) {
	if e.flag == attrFlag {
		attr = e.name + "=\""
		for _, value := range e.val {
			attr += value + " "
		}
		attr += "\""
	} else {
		err = errors.New("Not an attribute")
	}
	return attr, err
}

func (e elemen) repPreelement() (tag string, err error) {
	if e.flag == eleFlag {
		tag = "<" + e.name
		for _, attr := range e.attr {
			str, _ := attr.repAttr()
			tag += " " + str
		}
		tag += ">"
	} else {
		err = errors.New("Not an element")
	}
	//fmt.Println("123", tag)
	return tag, err
}

func (e elemen) repBackelement() (tag string, err error) {
	if e.flag == eleFlag {
		tag = "</" + e.name + ">"
	} else {
		err = errors.New("Not an element")
	}
	return tag, err
}
