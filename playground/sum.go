package playground

import (
	"fmt"
	"strings"
	"text/scanner"
	"unicode"
)

// Sum ...
func Sum(a, b int) int {
	return a + b
}

func main() {
	const src = "%var1 var2%"

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "default"

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}

	fmt.Println()
	s.Init(strings.NewReader(src))
	s.Filename = "percent"

	// treat leading '%' as part of an identifier
	s.IsIdentRune = func(ch rune, i int) bool {
		return ch == '%' && i == 0 || unicode.IsLetter(ch) || unicode.IsDigit(ch) && i > 0
	}

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}

}
