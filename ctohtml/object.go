package ctohtml

const (
	opStr      string = "#>[]*.^!+{},"
	illegalStr string = "\t\n \r"
)

type Stack []string
type Str []string
type zenObj string
