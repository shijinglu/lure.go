package lure

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// VersionData ...
type VersionData struct {
	major, minor, patch int
}

// NewVersionFromString parses version from string (`v3.2.1`)
func NewVersionFromString(ver string) (VersionData, error) {

	splits := strings.Split(strings.TrimLeft(ver, "vV"), ".")
	verArr := make([]int, 3)
	for i := 0; i < 3 && i < len(splits); i++ {
		n, e := strconv.Atoi(splits[i])
		if e != nil {
			return VersionData{}, e
		}
		verArr[i] = n
	}
	return VersionData{
		major: verArr[0],
		minor: verArr[1],
		patch: verArr[2],
	}, nil
}

func (VersionData) extKey() string {
	return "semver"
}

// compareTo compares me with rhs
// if both are valid return meaningful comparison int
// if rhs is nil, return 1
// if rhs is not a version, try convert and then compare
// if rhs cannot be casted to version, return INT_MIN
func (me VersionData) compareTo(rhs IData) int {
	if rhs == nil {
		return 1
	}
	if rv, ok := rhs.(VersionData); ok {
		if me.major != rv.major {
			return me.major - rv.major
		}
		if me.minor != rv.minor {
			return me.minor - rv.minor
		}
		if me.patch != rv.patch {
			return me.patch - rv.patch
		}
		return 0
	}
	if rv, err := NewVersionFromString(rhs.toString()); err == nil {
		return me.compareTo(rv)
	}
	return math.MinInt32
}

func (me VersionData) toBoolean() bool {
	return !(me.major == 0 && me.minor == 0 && me.patch == 0)
}

func (me VersionData) toInt() int {
	if i, err := strconv.Atoi(fmt.Sprintf("%d%d%d", me.major, me.minor, me.patch)); err == nil {
		return i
	}
	return 0
}

func (me VersionData) toDouble() float64 {
	return float64(me.toInt())
}

func (me VersionData) toString() string {
	return fmt.Sprintf("v%d.%d.%d", me.major, me.minor, me.patch)
}

// VersionDataConstructor constructs VersionData from raw string
func VersionDataConstructor(raw string, extKey string) IData {
	if ver, err := NewVersionFromString(raw); err == nil {
		return ver
	}
	return VersionData{
		major: 0,
		minor: 0,
		patch: 0,
	}
}
