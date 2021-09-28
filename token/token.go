package token

type Token struct {
	Type TokenType
	Text string
	Position int
}

func NewToken(t TokenType, text string, pos int) Token {
	return Token{
		Type: t,
		Text: text,
		Position: pos,
	}
}
