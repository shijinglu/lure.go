package lure

import "strconv"

// DoubleData ...
type DoubleData struct {
	val float64
}

func (DoubleData) extKey() string {
	return ""
}

func (me DoubleData) compareTo(rhs IData) int {
	cmp := me.val - rhs.toDouble()
	if cmp > 0 {
		return 1
	}
	if cmp < 0 {
		return -1
	}
	return 0
}

func (me DoubleData) toBoolean() bool {
	if me.val == 0 {
		return false
	}
	return true
}

func (me DoubleData) toInt() int {
	return int(me.val)
}

func (me DoubleData) toDouble() float64 {
	return me.val
}

func (me DoubleData) toString() string {
	return strconv.FormatFloat(me.val, 'g', -1, 64)
}
