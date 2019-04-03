package lure

import "strconv"

// IntData ...
type IntData struct {
	val int
}

func (IntData) extKey() string {
	return ""
}

func (me IntData) compareTo(rhs IData) int {
	return me.val - rhs.toInt()
}

func (me IntData) toBoolean() bool {
	if me.val == 0 {
		return false
	}
	return true
}

func (me IntData) toInt() int {
	return int(me.val)
}

func (me IntData) toDouble() float64 {
	return float64(me.val)
}

func (me IntData) toString() string {
	return strconv.FormatInt(int64(me.val), 10)
}
