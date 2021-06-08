package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT" // add, x, y
	INT   = "INT"   // default int8

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	TIMES  = "*"
	DIVID  = "/"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Saved words
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
)
