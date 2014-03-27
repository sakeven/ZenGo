package ctohtml

const (
	opStr      string = "#>[]*.^!+{},"
	illegalStr string = "\t\n \r"
)

type Stack []string
type Str []string
type zenObj string
type value string
type attribute struct {
	name string
	val  value
}
type elemen struct {
	name string
	val  []value
	attr []attribute
}
