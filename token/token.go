package token

type TokenType string

type Token struct {
	TokenType TokenType
	Literal   string
}

const (
	// Special
	UNKNOWN = "UNKNOWN"
	EOF     = "EOF"

	// Identifiers
	IDENTIFIER = "IDENTIFIER"

	// literals
	INT = "INT"

	// Operators
	ASSIGN       = "="
	PLUS         = "PLUS"
	MINUS        = "-"
	BANG         = "!"
	ASTERISK     = "*"
	SLASH        = "/"
	LESS_THAN    = "<"
	GREATER_THAN = ">"
	EQUALS       = "=="
	NOT_EQUALS   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHISIS = ")"
	LEFT_BRACE        = "{"
	RIGHT_BRACE       = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func GetTokenTypeForWord(word string) TokenType {
	token, isKeyword := keywords[word]
	if isKeyword {
		return token
	}
	return IDENTIFIER
}
