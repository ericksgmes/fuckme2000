package lexer

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
    LET       = "LET"
    IDENT     = "IDENT"
    ASSIGN    = "="
    INT       = "INT"
    PLUS      = "+"
    SEMICOLON = ";"
    PRINT     = "PRINT"
    LPAREN    = "("
    RPAREN    = ")"
    EOF       = "EOF"
)
