// internal/parser/parser.go
package parser

import (
	"fmt"
	"fuckme2000/internal/lexer"
)

type Parser struct {
    tokens  []lexer.Token
    pos     int
    current lexer.Token
}

func New(tokens []lexer.Token) *Parser {
    p := &Parser{
        tokens: tokens,
        pos:    -1,
    }
    p.advance()
    return p
}

func (p *Parser) advance() {
    p.pos++
    if p.pos < len(p.tokens) {
        p.current = p.tokens[p.pos]
    } else {
        p.current = lexer.Token{Type: "EOF", Literal: ""}
    }
}

func (p *Parser) ParseLet() []string {
	p.advance()
	ident := p.current.Literal
	p.advance()
	p.advance()
	value := p.current.Literal
	p.advance()
	return []string{
		fmt.Sprintf("PUSH %s", value),
		fmt.Sprintf("STORE %s", ident),
	}
}
