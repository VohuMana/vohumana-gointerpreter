package token

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
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"

    // Comparison Operators
    LT = "<"
    GT = ">"
    EQ = "=="
    NOT_EQ = "!="

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
    TRUE = "TRUE"
    FALSE = "FALSE"
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"

)

func NewToken(tokenType TokenType, ch byte) Token {
    return Token {Type: tokenType, Literal: string(ch)}
}

var keywords = map[string]TokenType {
    "fn" : FUNCTION,
    "let": LET,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
}

func LookupIdentifier(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }

    return IDENTIFIER
}