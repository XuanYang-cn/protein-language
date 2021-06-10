package lexer

import (
	"testing"

	"protine/token"
)

func TestNextToken(t *testing.T) {
	t.Run("Test Simple Token", func(t *testing.T) {
		input := "=+*/-(){},;"

		tests := []struct {
			expectedType    token.Type
			expectedLiteral string
		}{
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
			{token.TIMES, "*"},
			{token.DIVID, "/"},
			{token.MINUS, "-"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.COMMA, ","},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		l := New(input)

		for i, tt := range tests {
			tok := l.NextToken()
			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
					i, tt.expectedType, tok.Type)
			}

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
					i, tt.expectedLiteral, tok.Literal)
			}
		}
	})

	t.Run("Test complex code", func(t *testing.T) {
		input := `var five int = 5;
var ten int = 10;

var add = fn(a int, b int) int {
    return a + b;
};

var result int = add(five, ten);
`
		tests := []struct {
			expectedType    token.Type
			expectedLiteral string
		}{
			{token.VAR, "var"},
			{token.IDENT, "five"},
			{token.INT, "int"},
			{token.ASSIGN, "="},
			{token.INT8, "5"},
			{token.SEMICOLON, ";"},

			{token.VAR, "var"},
			{token.IDENT, "ten"},
			{token.INT, "int"},
			{token.ASSIGN, "="},
			{token.INT8, "10"},
			{token.SEMICOLON, ";"},

			{token.VAR, "var"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "fn"},
			{token.LPAREN, "("},
			{token.IDENT, "a"},
			{token.INT, "int"},
			{token.COMMA, ","},
			{token.IDENT, "b"},
			{token.INT, "int"},
			{token.RPAREN, ")"},
			{token.INT, "int"},
			{token.LBRACE, "{"},
			{token.RETURN, "return"},
			{token.IDENT, "a"},
			{token.PLUS, "+"},
			{token.IDENT, "b"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},

			{token.VAR, "var"},
			{token.IDENT, "result"},
			{token.INT, "int"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "five"},
			{token.COMMA, ","},
			{token.IDENT, "ten"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
		}

		l := New(input)

		for i, tt := range tests {
			tok := l.NextToken()
			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
					i, tt.expectedType, tok.Type)
			}

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
					i, tt.expectedLiteral, tok.Literal)
			}
		}

	})

}
