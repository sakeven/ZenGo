package ctohtml

const (
	opStr      string = "#>[]*.^!+{},"
	illegalStr string = "\t\n \r"
	eleFlag    int    = 1
	attrFlag   int    = 2
	valueFlag  int    = 3
)

type Stack []string
type Str []string
type zenObj string
