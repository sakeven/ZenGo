package zentohtml

const (
	opStr      string = "#>[]*.^!+{},"
	illegalOp  string = "#[].!{},"
	illegalStr string = "\t\n \r"
	endStr     string = ">#*+[^,]"
)
const (
	nonFlag = iota
	textFlag
	opFlag
	eleFlag
	attrFlag
	valueFlag
	mulFlag
)

type Str []string
type ZenObj string
type elemen struct {
	name string
	val  []string
	attr []elemen
	flag int
}
type Stack []elemen
type eleArr []elemen
