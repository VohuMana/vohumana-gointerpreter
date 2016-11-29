package lexer

import "testing"

func TestNextToken(t *testing.T) {
    input := `=+(){},;`

    tests := []struct {
        expectedType TokenType
        expectedLiteral string
    }{
        {ASSIGN, "="},
        {PLUS, "+"},
        {LPAREN, "("},
        {RPAREN, ")"},
        {LBRACE, "{"},
        {RBRACE, "}"},
        {COMMA, ","},
        {SEMICOLON, ";"},
        {EOF, ""},
    }

    lex := New(input)

    for i, tt := range tests {
        tok := lex.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatal("tests[%d] - tokentype is wrong. Expected=%q, got=%q", i, tt.expectedType, tok.Type)
        }

        if tok.Literal != tt.expectedLiteral {
            t.Fatal("tests[%d] - literal is wrong. Expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
        }
    }
}