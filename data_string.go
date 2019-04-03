package lure

import (
	"strconv"
	"strings"
)

// StringData ...
type StringData struct {
	val string
}

func (StringData) extKey() string {
	return ""
}

func (me StringData) compareTo(rhs IData) int {
	return strings.Compare(me.val, rhs.toString())
}

// toBoolean cast string to boolean, if current string
// is parseable, it return the parsed boolean, return
// false otherwise
func (me StringData) toBoolean() bool {
	if b, err := strconv.ParseBool(me.val); err != nil {
		return b
	}
	return false
}

func (me StringData) toInt() int {
	if v, err := strconv.Atoi(me.val); err != nil {
		return v
	}
	return 0
}

func (me StringData) toDouble() float64 {
	if v, err := strconv.ParseFloat(me.val, 64); err != nil {
		return v
	}
	return 0.0
}

func (me StringData) toString() string {
	return me.val
}
