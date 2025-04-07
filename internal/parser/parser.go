package parser

import (
	"fuckme2000/internal/lexer"
	"fuckme2000/internal/vm"
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

func (p *Parser) ParseLet() []vm.Instruction {
	p.advance() // IDENT
	ident := p.current.Literal

	p.advance() // ASSIGN
	p.advance() // INT
	value := p.current.Literal

	p.advance() // SEMICOLON 
	
	return []vm.Instruction{
		{Op: "PUSH", Arg: value},
		{Op: "STORE", Arg: ident},
	}
}

func (p *Parser) ParsePrint() []vm.Instruction {
	var r []vm.Instruction

	p.advance() // LPAREN
	p.advance() // IDENT or INT

	left := p.current.Literal
	if p.current.Type == lexer.INT {
		r = append(r, vm.Instruction{Op: "PUSH", Arg: left})
	} else {
		r = append(r, vm.Instruction{Op: "LOAD", Arg: left})
	}

	p.advance() // PLUS or RPAREN

	if p.current.Type == lexer.PLUS {
		p.advance() // IDENT or INT
		right := p.current.Literal

		if p.current.Type == lexer.IDENT {
			r = append(r, vm.Instruction{Op: "LOAD", Arg: right})
		} else {
			r = append(r, vm.Instruction{Op: "PUSH", Arg: right})
		}

		r = append(r, vm.Instruction{Op: "ADD"})
		p.advance() // RPAREN
		p.advance() // SEMICOLON

		r = append(r, vm.Instruction{Op: "PRINT"})
		return r
	}

	// No addition
	p.advance() // RPAREN
	p.advance() // SEMICOLON

	r = append(r, vm.Instruction{Op: "PRINT"})
	return r
}

func (p *Parser) ParseProgram() []vm.Instruction {
	var instructions []vm.Instruction

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
