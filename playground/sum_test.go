package playground

import (
	"fmt"
	"strings"
	"testing"
	"text/scanner"
)

func testSimple(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

}

func testScanner(t *testing.T) {
	const src = `CITY_ID == 3 && VERION = 'v3.1.4'`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}
}
