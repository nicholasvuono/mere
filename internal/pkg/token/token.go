package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers + literals
	IDENTIFIER = "IDENTIFIER" //add, foo, bar, x, y, etc.
	NUMBER     = "NUMBER"     // 11234412
	STRING     = "STRING"
	BOOLEAN    = "BOOLEAN"

	//Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"

	LESS_THAN    = "<"
	GREATER_THAN = ">"

	EQUALS    = "=="
	NOT_EQUAL = "!="

	//Delimiters
	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHESIS = ")"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSEIF   = "ELSEIF"
	ELSE     = "ELSE"
	WHILE    = "WHILE"
	RETURN   = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"func":    FUNCTION,
	"let":     LET,
	"true":    TRUE,
	"false":   FALSE,
	"if":      IF,
	"else if": ELSEIF,
	"else":    ELSE,
	"while":   WHILE,
	"return":  RETURN,
}

func LookupIdentifier(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	}
	return IDENTIFIER
}
