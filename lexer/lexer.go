package lexer

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

func (l *Lexer) NextToken() Token {
    var tok Token

    switch l.char {
        case '=':
            tok = newToken(ASSIGN, l.char)
        case ';':
            tok = newToken(SEMICOLON, l.char)
        case '(':
            tok = newToken(LPAREN, l.char)
        case ')':
            tok = newToken(RPAREN, l.char)
        case ',':
            tok = newToken(COMMA, l.char)
        case '+':
            tok = newToken(PLUS, l.char)
        case '{':
            tok = newToken(LBRACE, l.char)
        case '}':
            tok = newToken(RBRACE, l.char)
        case 0:
            tok.Literal = ""
            tok.Type = EOF
    }

    l.readChar()
    return tok
}
