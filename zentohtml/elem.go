package zentohtml

import (
	"errors"
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
	case mulFalg:
		return "Repeate"
	}
}

func (e elemen) repAttr() (attr string, err error) {
	if e.flag == attrFlag {
		attr = e.name + "=\""
		for _, value := range e.val {
			attr += e.val + " "
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
			tag += " " + attr.repAttr()
		}
		tag += ">"
	} else {
		err = errors.New("Not an element")
	}
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
