package internal

import "fmt"

type Token struct {
	t       TokenType
	lexeme  string
	literal interface{}
	line    int
}

func (t Token) String() string {
	return fmt.Sprintf("%v %v %v %d", t.t, t.lexeme, t.literal, t.line)
}
