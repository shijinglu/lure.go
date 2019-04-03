package lure

import (
	"testing"
)

func TestLexer_Lex(t *testing.T) {

	lexer0 := NewLexerOfString(`CITY_ID >= 3 && (VERION == "v3.1.4" || PI < 3.14) && env in ("qa", "dev") `)
	type args struct {
		lval *LureSymType
	}
	defaultArg := args{
		lval: &LureSymType{},
	}

	tests := []struct {
		name    string
		lexer   *Lexer
		args    args
		want    int
		wantErr bool
	}{
		{
			name:  "CITY_ID is an identifier",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_IDENTITY_LITERAL,
		},
		{
			name:  ">= is an op",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_GE,
		},
		{
			name:  "3 is an integer",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_INT_LITERAL,
		},
		{
			name:  "&& is an op",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_AND_LOGIC,
		},
		{
			name:  "left bracket (",
			lexer: lexer0,
			args:  defaultArg,
			want:  int('('),
		},
		{
			name:  "VERION is an identifier",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_IDENTITY_LITERAL,
		},
		{
			name:  "== again",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_EQ,
		},
		{
			name:  "v3.1.4 is a raw string",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_STRING_LITERAL,
		},
		{
			name:  "|| is an op",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_OR_LOGIC,
		},
		{
			name:  "PI is an indentifier",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_IDENTITY_LITERAL,
		},
		{
			name:  "< is an op",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_LT,
		},
		{
			name:  "3.14 is a float number",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_DOUBLE_LITERAL,
		},
		{
			name:  "right bracket )",
			lexer: lexer0,
			args:  defaultArg,
			want:  int(')'),
		},
		{
			name:  "&& again",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_AND_LOGIC,
		},
		{
			name:  "env is an identifier",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_IDENTITY_LITERAL,
		},
		{
			name:  "in is an op",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_IN,
		},
		{
			name:  "left bracket 2",
			lexer: lexer0,
			args:  defaultArg,
			want:  int('('),
		},
		{
			name:  "qa is a raw string",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_STRING_LITERAL,
		},
		{
			name:  ", is a char",
			lexer: lexer0,
			args:  defaultArg,
			want:  int(','),
		},
		{
			name:  "dev is a raw string",
			lexer: lexer0,
			args:  defaultArg,
			want:  TK_STRING_LITERAL,
		},
		{
			name:  "right bracket 2",
			lexer: lexer0,
			args:  defaultArg,
			want:  int(')'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.lexer.Lex(tt.args.lval)
			if got != tt.want {
				t.Errorf("Lexer.Lex() = %v, want %v", got, tt.want)
			}
		})
	}
}
