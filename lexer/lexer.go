package lexer

import (
	"fmt"
	"monkey/token"
)

type Lexer struct {
	input       string
	position    int
	currentChar byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input, position: -1}
	return lexer
}

func (l *Lexer) readChar() {
	if l.position+1 >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.position+1]
	}
	fmt.Print(string(l.currentChar))
	l.position++
}

func (l *Lexer) lookAheadChar() byte {
	if l.position+1 >= len(l.input) {
		return 0
	} else {
		return l.input[l.position+1]
	}
}

func (l *Lexer) NextToken() token.Token {
	var nextToken token.Token

	l.skipWhitespace()
	l.readChar()

	switch l.currentChar {
	case '=':
		if l.lookAheadChar() == '=' {
			literal := string(l.currentChar)
			l.readChar()
			literal += string(l.currentChar)
			nextToken.Literal = literal
			nextToken.TokenType = token.EQUALS
		} else {
			nextToken = newToken(token.ASSIGN, l.currentChar)
		}
	case ';':
		nextToken = newToken(token.SEMICOLON, l.currentChar)
	case '(':
		nextToken = newToken(token.LEFT_PARENTHESIS, l.currentChar)
	case ')':
		nextToken = newToken(token.RIGHT_PARENTHISIS, l.currentChar)
	case ',':
		nextToken = newToken(token.COMMA, l.currentChar)
	case '+':
		nextToken = newToken(token.PLUS, l.currentChar)
	case '{':
		nextToken = newToken(token.LEFT_BRACE, l.currentChar)
	case '}':
		nextToken = newToken(token.RIGHT_BRACE, l.currentChar)
	case '-':
		nextToken = newToken(token.MINUS, l.currentChar)
	case '!':
		if l.lookAheadChar() == '=' {
			literal := string(l.currentChar)
			l.readChar()
			literal += string(l.currentChar)
			nextToken.Literal = literal
			nextToken.TokenType = token.NOT_EQUALS
		} else {
			nextToken = newToken(token.BANG, l.currentChar)
		}
	case '*':
		nextToken = newToken(token.ASTERISK, l.currentChar)
	case '/':
		nextToken = newToken(token.SLASH, l.currentChar)
	case '<':
		nextToken = newToken(token.LESS_THAN, l.currentChar)
	case '>':
		nextToken = newToken(token.GREATER_THAN, l.currentChar)
	case 0:
		nextToken.Literal = ""
		nextToken.TokenType = token.EOF
	default:
		if isDigit(l.currentChar) {
			nextToken.Literal = l.readNumber()
			nextToken.TokenType = token.INT
		} else if isLetter(l.currentChar) {
			nextToken.Literal = l.readWord()
			nextToken.TokenType = token.GetTokenTypeForWord(nextToken.Literal)
		} else {
			nextToken = newToken(token.UNKNOWN, l.currentChar)
		}
	}
	return nextToken
}

func isWhiteSpaceChar(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func (l *Lexer) skipWhitespace() {
	for isWhiteSpaceChar(l.lookAheadChar()) {
		l.readChar()
	}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' ||
		'A' <= char && char <= 'Z' ||
		char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) readWord() string {
	wordStartPosition := l.position
	for isLetter(l.lookAheadChar()) || isDigit(l.lookAheadChar()) {
		l.readChar()
	}
	wordEndPosition := l.position
	return l.input[wordStartPosition : wordEndPosition+1]
}

func (l *Lexer) readNumber() string {
	numberStartPosition := l.position
	for isDigit(l.lookAheadChar()) {
		l.readChar()
	}
	numberEndPostition := l.position
	return l.input[numberStartPosition : numberEndPostition+1]
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{TokenType: tokenType, Literal: string(char)}
}
