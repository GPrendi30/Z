package ast

import "Z/token"

type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Text
}

