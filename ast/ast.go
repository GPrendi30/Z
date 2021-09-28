package ast

import "Z/token"

type Node interface {
	TokenLiteral() string
}

type Expression interface {
	Node
	expressionNode()
}


type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Text
}



