package main

import (
	"fmt"
)

type TokenType uint16

const (
	//
	YA                = iota
	KEYWORD TokenType = iota
	IDENTIFIER
)

//???
type Token interface {
	Type() TokenType
	Lexeme() string
}

type Match struct {
	toketype TokenType
	lexeme   string
}

type IntergerConstant struct {
	Token
	value uint64
}

func (m *Match) Type() TokenType {
	return m.toketype
}
func (m *Match) Lexeme() string {
	return m.lexeme
}

func (i *IntergerConstant) Value() uint64 {
	return i.value
}
func main() {
	// fmt.Println("ya", YA)
	t := IntergerConstant{&Match{KEYWORD, "wizard"}, 2}
	fmt.Println(t.Type(), t.Lexeme(), t.Value())
	x := Token(t)
	fmt.Println(x.Type(), x.Lexeme())
}
