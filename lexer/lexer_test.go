package lexer

import "testing"

func TestNextToken_Simple(t *testing.T) {
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

func TestNextToken_Advanced(t *testing.T) {
    input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5
5 < 10 > 5;

if (5 < 10) {
    return true;
} else {
    return false;
}

5 == 10; 5 != 10`

    tests := []struct {
        expectedType TokenType
        expectedLiteral string
    }{
        {LET, "let"},
        {IDENTIFIER, "five"},
        {ASSIGN, "="},
        {INT, "5"},
        {SEMICOLON, ";"},
        {LET, "let"},
        {IDENTIFIER, "ten"},
        {ASSIGN, "="},
        {INT, "10"},
        {SEMICOLON, ";"},
        {LET, "let"},
        {IDENTIFIER, "add"},
        {ASSIGN, "="},
        {FUNCTION, "fn"},
        {LPAREN, "("},
        {IDENTIFIER, "x"},
        {COMMA, ","},
        {IDENTIFIER, "y"},
        {RPAREN, ")"},
        {LBRACE, "{"},
        {IDENTIFIER, "x"},
        {PLUS, "+"},
        {IDENTIFIER, "y"},
        {SEMICOLON, ";"},
        {RBRACE, "}"},
        {SEMICOLON, ";"},
        {LET, "let"},
        {IDENTIFIER, "result"},
        {ASSIGN, "="},
        {IDENTIFIER, "add"},
        {LPAREN, "("},
        {IDENTIFIER, "five"},
        {COMMA, ","},
        {IDENTIFIER, "ten"},
        {RPAREN, ")"},
        {SEMICOLON, ";"},
        {BANG, "!"},
        {MINUS, "-"},
        {SLASH, "/"},
        {ASTERISK, "*"},
        {INT, "5"},
        {INT, "5"},
        {LT, "<"},
        {INT, "10"},
        {GT, ">"},
        {INT, "5"},
        {SEMICOLON, ";"},
        {IF, "if"},
        {LPAREN, "("},
        {INT, "5"},
        {LT, "<"},
        {INT, "10"},
        {RPAREN, ")"},
        {LBRACE, "{"},
        {RETURN, "return"},
        {TRUE, "true"},
        {SEMICOLON, ";"},
        {RBRACE, "}"},
        {ELSE, "else"},
        {LBRACE, "{"},
        {RETURN, "return"},
        {FALSE, "false"},
        {SEMICOLON, ";"},
        {RBRACE, "}"},
        {INT, "5"},
        {EQ, "=="},
        {INT, "10"},
        {SEMICOLON, ";"},
        {INT, "5"},
        {NOT_EQ, "!="},
        {INT, "10"},
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