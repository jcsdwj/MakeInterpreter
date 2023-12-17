/**
  @author: 伍萬
  @since: 2023/11/29
  @desc: //TODO
**/

package parser

import (
	"fmt"
	ast2 "lexer/ast"
	"lexer/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements go=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast2.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral() should be 'let' got=%s", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast2.LetStatement)
	if !ok {
		t.Errorf("s should be *ast2.LetStatement got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value should be %s got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() should be %s got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements go=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast2.ReturnStatement)
		if !ok {
			t.Errorf("stmt should be *ast2.ReturnStatement got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral() should be'return' got=%q",
				returnStmt.TokenLiteral())
			continue
		}
	}
}

// 整数字面量
func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not contain enough statements go=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast2.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not *ast2.ExpressionStatement got=%T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast2.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiterak.got=%T", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value should be 5 got=%d", literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s.got=%s", "5", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statement does not contain %d statements.got = %d\n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast2.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement got=%T",
				program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast2.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'.got=%s", tt.operator, exp.Operator)
		}

		if !testInegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}
func testInegerLiteral(t *testing.T, il ast2.Expression, value int64) bool {
	integ, ok := il.(*ast2.IntegerLiteral)
	if !ok {
		t.Errorf("il not *asct.IntegerLiterak.got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d.got=%d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d.got=%s", value, integ.TokenLiteral())
		return false
	}
	return true
}
