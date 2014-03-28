package zentohtml

const (
	opStr      string = "#>[]*.^!+{},"
	illegalStr string = "\t\n \r"
)
const (
	nonFlag = iota
	textFlag
	opFlag
	eleFlag
	attrFlag
	valueFlag
	mulFalg
)

type Stack []string
type Str []string
type zenObj string
type elemen struct {
	name string
	val  []string
	attr []attribute
	flag int
}
