package lexer

import (
	"fmt"
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
`
	tests := []struct {
		expectedToken token.Token
	}{
		{expectedToken: token.Token{TokenType: token.LET, Literal: "let"}},
		{expectedToken: token.Token{TokenType: token.IDENTIFIER, Literal: "five"}},
		{expectedToken: token.Token{TokenType: token.ASSIGN, Literal: "="}},
		{expectedToken: token.Token{TokenType: token.INT, Literal: "5"}},
		{expectedToken: token.Token{TokenType: token.SEMICOLON, Literal: ";"}},
		{expectedToken: token.Token{TokenType: token.LET, Literal: "let"}},
		{expectedToken: token.Token{TokenType: token.IDENTIFIER, Literal: "ten"}},
		{expectedToken: token.Token{TokenType: token.ASSIGN, Literal: "="}},
		{expectedToken: token.Token{TokenType: token.INT, Literal: "10"}},
		{expectedToken: token.Token{TokenType: token.SEMICOLON, Literal: ";"}},
		{expectedToken: token.Token{TokenType: token.LET, Literal: "let"}},
		{expectedToken: token.Token{TokenType: token.IDENTIFIER, Literal: "add"}},
		{expectedToken: token.Token{TokenType: token.ASSIGN, Literal: "="}},
		{expectedToken: token.Token{TokenType: token.FUNCTION, Literal: "fn"}},
		{expectedToken: token.Token{TokenType: token.LEFT_PARENTHESIS, Literal: "("}},
		{expectedToken: token.Token{TokenType: token.IDENTIFIER, Literal: "x"}},
		{expectedToken: token.Token{TokenType: token.COMMA, Literal: ","}},
		{expectedToken: token.Token{TokenType: token.IDENTIFIER, Literal: "y"}},
		{expectedToken: token.Token{TokenType: token.RIGHT_PARENTHISIS, Literal: ")"}},
		{expectedToken: token.Token{TokenType: token.LEFT_BRACE, Literal: "{"}},
		{expectedToken: token.Token{TokenType: token.IDENTIFIER, Literal: "x"}},
		{expectedToken: token.Token{TokenType: token.PLUS, Literal: "+"}},
		{expectedToken: token.Token{TokenType: token.IDENTIFIER, Literal: "y"}},
		{expectedToken: token.Token{TokenType: token.SEMICOLON, Literal: ";"}},
		{expectedToken: token.Token{TokenType: token.RIGHT_BRACE, Literal: "}"}},
		{expectedToken: token.Token{TokenType: token.SEMICOLON, Literal: ";"}},
		{expectedToken: token.Token{TokenType: token.LET, Literal: "let"}},
		{expectedToken: token.Token{TokenType: token.IDENTIFIER, Literal: "result"}},
		{expectedToken: token.Token{TokenType: token.ASSIGN, Literal: "="}},
		{expectedToken: token.Token{TokenType: token.IDENTIFIER, Literal: "add"}},
		{expectedToken: token.Token{TokenType: token.LEFT_PARENTHESIS, Literal: "("}},
		{expectedToken: token.Token{TokenType: token.IDENTIFIER, Literal: "five"}},
		{expectedToken: token.Token{TokenType: token.COMMA, Literal: ","}},
		{expectedToken: token.Token{TokenType: token.IDENTIFIER, Literal: "ten"}},
		{expectedToken: token.Token{TokenType: token.RIGHT_PARENTHISIS, Literal: ")"}},
		{expectedToken: token.Token{TokenType: token.SEMICOLON, Literal: ";"}},
		{expectedToken: token.Token{TokenType: token.BANG, Literal: "!"}},
		{expectedToken: token.Token{TokenType: token.MINUS, Literal: "-"}},
		{expectedToken: token.Token{TokenType: token.SLASH, Literal: "/"}},
		{expectedToken: token.Token{TokenType: token.ASTERISK, Literal: "*"}},
		{expectedToken: token.Token{TokenType: token.INT, Literal: "5"}},
		{expectedToken: token.Token{TokenType: token.SEMICOLON, Literal: ";"}},
		{expectedToken: token.Token{TokenType: token.INT, Literal: "5"}},
		{expectedToken: token.Token{TokenType: token.LESS_THAN, Literal: "<"}},
		{expectedToken: token.Token{TokenType: token.INT, Literal: "10"}},
		{expectedToken: token.Token{TokenType: token.GREATER_THAN, Literal: ">"}},
		{expectedToken: token.Token{TokenType: token.INT, Literal: "5"}},
		{expectedToken: token.Token{TokenType: token.SEMICOLON, Literal: ";"}},
		{expectedToken: token.Token{TokenType: token.IF, Literal: "if"}},
		{expectedToken: token.Token{TokenType: token.LEFT_PARENTHESIS, Literal: "("}},
		{expectedToken: token.Token{TokenType: token.INT, Literal: "5"}},
		{expectedToken: token.Token{TokenType: token.LESS_THAN, Literal: "<"}},
		{expectedToken: token.Token{TokenType: token.INT, Literal: "10"}},
		{expectedToken: token.Token{TokenType: token.RIGHT_PARENTHISIS, Literal: ")"}},
		{expectedToken: token.Token{TokenType: token.LEFT_BRACE, Literal: "{"}},
		{expectedToken: token.Token{TokenType: token.RETURN, Literal: "return"}},
		{expectedToken: token.Token{TokenType: token.TRUE, Literal: "true"}},
		{expectedToken: token.Token{TokenType: token.SEMICOLON, Literal: ";"}},
		{expectedToken: token.Token{TokenType: token.RIGHT_BRACE, Literal: "}"}},
		{expectedToken: token.Token{TokenType: token.ELSE, Literal: "else"}},
		{expectedToken: token.Token{TokenType: token.LEFT_BRACE, Literal: "{"}},
		{expectedToken: token.Token{TokenType: token.RETURN, Literal: "return"}},
		{expectedToken: token.Token{TokenType: token.FALSE, Literal: "false"}},
		{expectedToken: token.Token{TokenType: token.SEMICOLON, Literal: ";"}},
		{expectedToken: token.Token{TokenType: token.RIGHT_BRACE, Literal: "}"}},
		{expectedToken: token.Token{TokenType: token.INT, Literal: "10"}},
		{expectedToken: token.Token{TokenType: token.EQUALS, Literal: "=="}},
		{expectedToken: token.Token{TokenType: token.INT, Literal: "10"}},
		{expectedToken: token.Token{TokenType: token.SEMICOLON, Literal: ";"}},
		{expectedToken: token.Token{TokenType: token.INT, Literal: "10"}},
		{expectedToken: token.Token{TokenType: token.NOT_EQUALS, Literal: "!="}},
		{expectedToken: token.Token{TokenType: token.INT, Literal: "9"}},
		{expectedToken: token.Token{TokenType: token.SEMICOLON, Literal: ";"}},
		{expectedToken: token.Token{TokenType: token.EOF, Literal: ""}},
	}

	l := New(input)

	for i, test := range tests {
		currentToken := l.NextToken()
		testPassed, errMsg := validateToken(i, currentToken, test.expectedToken)
		if !testPassed {
			t.Fatalf(errMsg)
		}
	}
}

func validateToken(testNumber int, currentToken token.Token, expectedToken token.Token) (bool, string) {
	if currentToken.TokenType != expectedToken.TokenType {
		return false, fmt.Sprintf(
			"test %d - incorrect token type returned. expected: %q, found: %q",
			testNumber, expectedToken.TokenType, currentToken.TokenType,
		)
	}
	if currentToken.Literal != expectedToken.Literal {
		return false, fmt.Sprintf(
			"test %d - incorrect literal returned. expected: %q, found: %q",
			testNumber, expectedToken.Literal, currentToken.Literal,
		)
	}
	return true, ""
}
