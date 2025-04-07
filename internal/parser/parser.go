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

	var instrs []vm.Instruction

	p.advance() // First value (IDENT or INT)
	left := p.current.Literal

	if p.current.Type == lexer.INT {
		instrs = append(instrs, vm.Instruction{Op: "PUSH", Arg: left})
	} else if p.current.Type == lexer.IDENT {
		instrs = append(instrs, vm.Instruction{Op: "LOAD", Arg: left})
	} else {
		panic("Expected value after '='")
	}

	p.advance() // PLUS or SEMICOLON

	if p.current.Type == lexer.PLUS {
		p.advance() // right operand
		right := p.current.Literal

		if p.current.Type == lexer.INT {
			instrs = append(instrs, vm.Instruction{Op: "PUSH", Arg: right})
		} else if p.current.Type == lexer.IDENT {
			instrs = append(instrs, vm.Instruction{Op: "LOAD", Arg: right})
		} else {
			panic("Expected value after '+'")
		}

		instrs = append(instrs, vm.Instruction{Op: "ADD"})
		p.advance() // SEMICOLON
	}

	if p.current.Type != lexer.SEMICOLON {
		panic("Expected ';' after let statement")
	}

	p.advance() // move past ;

	instrs = append(instrs, vm.Instruction{Op: "STORE", Arg: ident})

	return instrs
}

func (p *Parser) ParsePrint() []vm.Instruction {
	var r []vm.Instruction

	p.advance() // LPAREN

	// Capture left side
	p.advance() // IDENT or INT
	left := p.current.Literal
	if p.current.Type == lexer.INT {
		r = append(r, vm.Instruction{Op: "PUSH", Arg: left})
	} else if p.current.Type == lexer.IDENT {
		r = append(r, vm.Instruction{Op: "LOAD", Arg: left})
	} else {
		panic("expected value inside print()")
	}

	p.advance() // Either PLUS or RPAREN

	if p.current.Type == lexer.PLUS {
		// Handle addition
		p.advance() // right operand
		right := p.current.Literal

		if p.current.Type == lexer.INT {
			r = append(r, vm.Instruction{Op: "PUSH", Arg: right})
		} else if p.current.Type == lexer.IDENT {
			r = append(r, vm.Instruction{Op: "LOAD", Arg: right})
		} else {
			panic("expected number or variable after '+'")
		}

		r = append(r, vm.Instruction{Op: "ADD"})

		p.advance() // RPAREN
		if p.current.Type != lexer.RPAREN {
			panic("expected ')' after print expression")
		}

		p.advance() // SEMICOLON
		if p.current.Type != lexer.SEMICOLON {
			panic("expected ';' after print()")
		}

		p.advance() // Move past SEMICOLON
		r = append(r, vm.Instruction{Op: "PRINT"})
		return r
	}

	// If there's no PLUS (just print(x) or print(5))
	if p.current.Type != lexer.RPAREN {
		panic("expected ')' after value in print()")
	}
	p.advance() // SEMICOLON
	if p.current.Type != lexer.SEMICOLON {
		panic("expected ';' after print()")
	}
	p.advance() // Move past SEMICOLON

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
