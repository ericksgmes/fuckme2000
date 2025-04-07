package lexer

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    IDENT   = "IDENT"
    INT     = "INT"

    ASSIGN  = "="
    PLUS    = "+"

    LET     = "LET"
    PRINT   = "PRINT"
    SEMICOLON = ";"
)