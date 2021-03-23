package tokenizer

import (
	"github.com/nicholasvuono/mere/internal/pkg/token"
)

type Tokenizer struct {
	input string // input string
	index int    // points to current char
	next  int    // points to next char
	char  byte   // current char
}

// create a new tokenizer anf read first char inthe input string
func NewTokenizer(input string) *Tokenizer {
	t := &Tokenizer{input: input}
	t.readChar()
	return t
}

// read character at the current index of the tokenizer input string
func (t *Tokenizer) readChar() {
	if t.next >= len(t.input) {
		t.char = 0
	} else {
		t.char = t.input[t.next]
	}
	t.index = t.next
	t.next += 1
}

// return new token for current character and move to the next character in the input
func (t *Tokenizer) NextToken() token.Token {
	var tok token.Token
	switch t.char {
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: string(t.char)}
	case '(':
		tok = token.Token{Type: token.LBRCK, Literal: string(t.char)}
	case ')':
		tok = token.Token{Type: token.RBRCK, Literal: string(t.char)}
	case '"':
		tok = token.Token{Type: token.STR, Literal: t.readStr()}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isLetter(t.char) {
			return token.Token{Type: token.Lookup(tok.Literal), Literal: t.readIdent()}
		} else if isDigit(t.char) {
			return token.Token{Type: token.NUM, Literal: t.readNum()}
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: string(t.char)}
		}
	}
	t.readChar()
	return tok
}

func (t *Tokenizer) readStr() string {
	index := t.index + 1
	for {
		t.readChar()
		if t.char == '"' {
			break
		}
	}
	return t.input[index:t.index]
}

func (t *Tokenizer) readIdent() string {
	index := t.index
	for isLetter(t.char) {
		t.readChar()
	}
	return t.input[index:t.index]
}

func (t *Tokenizer) readNum() string {
	index := t.index
	for isDigit(t.char) {
		t.readChar()
	}
	return t.input[index:t.index]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
