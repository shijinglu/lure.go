package lure

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// Lexer implements the LureLexer interface
// s contains the text to lex
type Lexer struct {
	s   *scanner.Scanner
	str string
	pos int
}

// NewLexerOfString ...
func NewLexerOfString(str string) *Lexer {
	var s scanner.Scanner
	s.Init(strings.NewReader(str))
	s.Filename = "__expr"
	return &Lexer{
		s:   &s,
		str: str,
		pos: 0,
	}
}

// Lex ...
func (l Lexer) Lex(lval *LureSymType) int {
	var acc string
	fmt.Printf("[%p] pos: %s, l.s.peek(): %s \n", &(l.s), l.s.Pos(), scanner.TokenString(l.s.Peek()))
	for tok := l.s.Scan(); tok != scanner.EOF; tok = l.s.Scan() {
		switch tok {
		case scanner.Ident:
			tokStr := l.s.TokenText()
			switch strings.ToLower(tokStr) {
			case "in":
				return TK_IN
			case "between":
				return TK_BETWEEN
			case "and":
				return TK_AND_KEYWORD
			case "or":
				return TK_AND_KEYWORD
			case "not":
				return TK_NOT
			case "md5mod":
				return TK_MD5MOD
			}
			lval.strVal = tokStr
			return TK_IDENTITY_LITERAL
		case scanner.Int:
			iVal, err := strconv.Atoi(l.s.TokenText())
			if err != nil {
				return 0
			}
			lval.intVal = iVal
			return TK_INT_LITERAL
		case scanner.Float:
			fVal, err := strconv.ParseFloat(l.s.TokenText(), 64)
			if err != nil {
				l.Error(err.Error())
				return -1
			}
			lval.doubleVal = fVal
			return TK_DOUBLE_LITERAL
		case scanner.RawString, scanner.Char, scanner.String:
			quoateStr := l.s.TokenText()
			lval.strVal = quoateStr[1 : len(quoateStr)-1]
			return TK_STRING_LITERAL
		default:
			// ">", "<", ">=", "<=", "==", "!=", "&&", "||"
			acc += string(tok)
			lval.strVal = acc
			switch acc {
			case "=", "!", "&", "|":
				// continue
			case "==":
				lval.intVal = TK_EQ
				return TK_EQ
			case ">=":
				lval.intVal = TK_GE
				return TK_GE
			case "<=":
				lval.intVal = TK_LE
				return TK_LE
			case "!=":
				lval.intVal = TK_NE
				return TK_NE
			case "&&":
				lval.intVal = TK_AND_LOGIC
				return TK_AND_LOGIC
			case "||":
				lval.intVal = TK_OR_LOGIC
				return TK_OR_LOGIC
			case ">":
				if l.s.Peek() != '=' {
					lval.intVal = TK_GT
					return TK_GT
				}
			case "<":
				if l.s.Peek() != '=' {
					lval.intVal = TK_LT
					return TK_LT
				}
			default:
				lval.intVal = int(tok)
				return int(tok)
			}
		}
		fmt.Printf("[%v] %s: %s\n", l.s.Position, l.s.TokenText(), scanner.TokenString(tok))
	}
	return 0
}

func (l Lexer) Error(msg string) {
	_ = fmt.Errorf("error parsing %s: reason: %s", l.str, msg)
}
