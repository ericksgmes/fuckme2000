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

func (p *Parser) ParsePrint() []string {
	var r []string
	p.advance()
	p.advance()
	left := p.current.Literal
	if p.current.Type == lexer.INT {
		r = append(r, fmt.Sprintf("PUSH %s", left))
	} else {
		r = append(r, fmt.Sprintf("LOAD %s", left))
	}
	p.advance()
	if p.current.Type == lexer.PLUS {
		p.advance()
		if p.current.Type == lexer.IDENT {
			right := p.current.Literal
			r = append(
				r, 
				fmt.Sprintf("LOAD %s", right),
				"ADD", 
				"PRINT",
			)
			return r
		} else {
			right := p.current.Literal
			r = append(
				r, 
				fmt.Sprintf("PUSH %s", right),
				"ADD", 
				"PRINT",
			)
			return r
		}
	} else {
		r = append(r, "PRINT")
		return r
	}
}

func (p *Parser) ParseProgram() []string {
    var instructions []string

    for p.current.Type != lexer.EOF {
        switch p.current.Type {
        case lexer.LET:
            instructions = append(instructions, p.ParseLet()...)
        case lexer.PRINT:
            instructions = append(instructions, p.ParsePrint()...)
        default:
            panic("unexpected token: " + p.current.Type)
        }
    }

    return instructions
}
