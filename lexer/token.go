package lexer

// TokenType is a string that represents what the type of the token is
type TokenType string

// Token is a struct that holds info about each token found
type Token struct {
    Type TokenType
    Literal string
    // TODO: Add line and column here
}

const (
    // Special cases
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // Identifiers and literals
    IDENTIFIER = "IDENT"
    INT = "INT"
    
    // Operators
    ASSIGN = "="
    PLUS = "+"

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"
    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"
)

func newToken(tokenType TokenType, ch byte) Token {
    return Token {Type: tokenType, Literal: string(ch)}
}