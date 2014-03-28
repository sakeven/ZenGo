package ctohtml

const (
	opStr      string = "#>[]*.^!+{},"
	illegalStr string = "\t\n \r"
)

type Stack []string
type Str []string
type zenObj string
type elemen struct {
	name string
	val  []string
	attr []attribute
}
