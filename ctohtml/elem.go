package ctohtml

import (
	"errors"
)

func (e elemen) isAttr() bool {
	return e.flag
}

func (e elemen) repAttr() (attr string, err error) {
	if e.isAttr() {
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
	if e.isAttr() == false {
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
	if e.isAttr() == false {
		tag = "</" + e.name + ">"
	} else {
		err = errors.New("Not an element")
	}
	return tag, err
}
