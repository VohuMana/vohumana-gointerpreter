package lexer

import (
    "github.com/vohumana/vohumana-gointerpreter/token"
)

// Lexer is used to read Moneky source 
type Lexer struct {
    // Monkey source
    input string
    // Current position in input (points to current char)
    position int 
    // Current reading position in input (after current char)
    readPosition int 
    // Current character being examined
    char byte
}

// New creates a new Lexer to be used
func New(input string) *Lexer {
    lex := &Lexer{input: input}
    lex.readChar()
    return lex
}

func (l *Lexer) readChar() {
    // Check if we are at the end of input
    if l.readPosition >= len(l.input) {
        l.char = 0
    } else {
        l.char = l.input[l.readPosition]
    }

    // Advance the position
    l.position = l.readPosition
    l.readPosition++
}

func (l *Lexer) peekChar() byte {
    if l.readPosition >= len(l.input) {
        return 0
    }

    return l.input[l.readPosition]
}

// NextToken parses the next token in the input
func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    l.skipWhitespace()

    switch l.char {
        case '=':
            if l.peekChar() == '=' {
                ch := l.char
                l.readChar()
                tok = token.Token{ Type: token.EQ, Literal: string(ch) + string(l.char)}
            } else {
                tok = token.NewToken(token.ASSIGN, l.char)
            }
        case ';':
            tok = token.NewToken(token.SEMICOLON, l.char)
        case '(':
            tok = token.NewToken(token.LPAREN, l.char)
        case ')':
            tok = token.NewToken(token.RPAREN, l.char)
        case ',':
            tok = token.NewToken(token.COMMA, l.char)
        case '+':
            tok = token.NewToken(token.PLUS, l.char)
        case '-':
            tok = token.NewToken(token.MINUS, l.char)
        case '{':
            tok = token.NewToken(token.LBRACE, l.char)
        case '}':
            tok = token.NewToken(token.RBRACE, l.char)
        case '!':
            if l.peekChar() == '=' {
                ch := l.char
                l.readChar()
                tok = token.Token{ Type: token.NOT_EQ, Literal: string(ch) + string(l.char)}
            } else {
                tok = token.NewToken(token.BANG, l.char)
            }
        case '/':
            tok = token.NewToken(token.SLASH, l.char)
        case '*':
            tok = token.NewToken(token.ASTERISK, l.char)
        case '<':
            tok = token.NewToken(token.LT, l.char)
        case '>':
            tok = token.NewToken(token.GT, l.char)
        case 0:
            tok.Literal = ""
            tok.Type = token.EOF
        default:
            if isLetter(l.char) {
                tok.Literal = l.readIdentifier()
                tok.Type = token.LookupIdentifier(tok.Literal)
                return tok
            } else if isDigit(l.char) {
                tok.Type = token.INT
                tok.Literal = l.readNumber()
                return tok
            }

            tok = token.NewToken(token.ILLEGAL, l.char)
    }

    l.readChar()
    return tok
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.char) {
        l.readChar()
    }

    return l.input[position: l.position]
}

func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.char) {
        l.readChar()
    }

    return l.input[position : l.position]
}

func (l *Lexer) skipWhitespace() {
    for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
        l.readChar()
    }
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}
