package lexer

import (
	"lexer/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	I := New(input)
	for i, tt := range tests {
		tok := I.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tok.Type expected %s, got %s",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tok.Literal expected %s, got %s",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNexToken2(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;

let add = fn(x,y){
	x+y;
};

let result = add(five,ten);
!-/*5;
5 < 10 > 5
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
	}

	I := New(input)
	for i, tt := range tests {
		tok := I.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tok.Type expected %s, got %s",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tok.Literal expected %s, got %s",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
