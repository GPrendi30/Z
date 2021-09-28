package lexer

import (
	"Z/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "set five = 5; set ten = 10; " +
		"set add = fun(x, y) { x + y;}; " +
		"set result = add(five, ten);" +
		"if, else" +
		"! * / > <" +
		"== !="

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.SET, "set"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.SET, "set"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.SET, "set"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fun"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.SET, "set"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.COMMA, ","},
		{token.ELSE, "else"},
		{token.BANG, "!"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.GT, ">"},
		{token.LT, "<"},
		{token.EQ, "=="},
		{token.NOT_EQ, "!="},


	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Text != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Text)
		}
	}
}