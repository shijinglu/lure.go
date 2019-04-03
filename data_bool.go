package lure

var (
	boolDataFalse = BoolData{val: false}
	boolDataTrue  = BoolData{val: true}
)

// BoolData ...
type BoolData struct {
	val bool
}

func (BoolData) extKey() string {
	return ""
}

func (me BoolData) compareTo(rhs IData) int {
	return bool2Int(me.val) - bool2Int(rhs.toBoolean())
}

func (me BoolData) toBoolean() bool {
	return me.val
}

func (me BoolData) toInt() int {
	return bool2Int(me.val)
}

func (me BoolData) toDouble() float64 {
	return float64(me.toInt())
}

func (me BoolData) toString() string {
	if me.val {
		return "true"
	}
	return "false"
}

func bool2Int(val bool) int {
	if val {
		return 1
	}
	return 0
}
