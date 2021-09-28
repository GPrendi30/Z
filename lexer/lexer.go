package lexer

import (
	"Z/token"
)

//TODO make l.ch a rune. add support for Unicode.

// Lexer is the lexical analyzer
// input as a string, is the source code that we want to analyze
// position is the current position, while readPosition is where the lexer should read next
// ch is the char under examination
type Lexer struct {
	input string
	position int
	readPosition int
	ch  byte
}

// New returns a new lexer from the given input
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar is a function of the lexer that reads the following char.
func (l* Lexer) readChar() {

	// if the position is beyond the len of the input, parse no more bytes
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l* Lexer) NextToken() token.Token {

	var tok token.Token
	l.skipWhiteSpace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.NewToken(token.EQ, string(ch)+string(l.ch), l.position)
		} else {
			tok = token.NewToken(token.ASSIGN, string(l.ch), l.position)
		}
	case ';':
		tok = token.NewToken(token.SEMICOLON, string(l.ch), l.position)
	case '(':
		tok = token.NewToken(token.LPAREN, string(l.ch), l.position)
	case ')':
		tok = token.NewToken(token.RPAREN, string(l.ch), l.position)
	case ',':
		tok = token.NewToken(token.COMMA, string(l.ch), l.position)
	case '+':
		tok = token.NewToken(token.PLUS, string(l.ch), l.position)
	case '-':
		tok = token.NewToken(token.MINUS, string(l.ch), l.position)
	case '*':
		tok = token.NewToken(token.ASTERISK, string(l.ch), l.position)
	case '/':
		tok = token.NewToken(token.SLASH, string(l.ch), l.position)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.NewToken(token.NOT_EQ, string(ch)+string(l.ch), l.position)
		} else {
		tok = token.NewToken(token.BANG, string(l.ch), l.position)
		}
	case '<':
		tok = token.NewToken(token.LT, string(l.ch), l.position)
	case '>':
		tok = token.NewToken(token.GT, string(l.ch), l.position)
	case '{':
		tok = token.NewToken(token.LBRACE, string(l.ch), l.position)
	case '}':
		tok = token.NewToken(token.RBRACE, string(l.ch), l.position)
	case 0:
		tok = token.NewToken(token.EOF, "", len(l.input)-1)
	default:
		if isLetter(l.ch) {
			tok.Position = l.position
			tok.Text = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Text)

			return tok
		} else if isDigit(l.ch) {
			tok.Position = l.position
			tok.Type = token.INT
			tok.Text = l.readNumber()

			return tok
		} else {
			//defer panic("Error: Illegal character at position")
			errorPos , errorByte := l.illegalError()
			return token.NewToken(token.ILLEGAL, string(errorByte), errorPos)
		}
	}

	l.readChar()
	return tok
}

//to be removed
func (l *Lexer) illegalError() (int, byte){
	errorPos := l.position
	errorByte := l.ch
	l.position = len(l.input)
	l.ch = 0
	return errorPos, errorByte
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

