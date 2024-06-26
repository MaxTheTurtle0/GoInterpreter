package lexer

import "go/interpreter/token"

type Lexer struct {
    input string
    position int // current position in input (points to current char)
    readPosition int // current reading position in input (after current char)
    ch byte // current char under examination
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token 
   
    l.skipWhitespace()
    
    peekChar := l.peekChar()
    switch l.ch {
    case '=':
        switch peekChar {
        case '=':
            tok = l.newTwoCharToken(token.EQ, l.ch)
        default:
            tok = newToken(token.ASSIGN, l.ch)
        }
    case '+':
        tok = newToken(token.PLUS, l.ch)
    case '-':
        tok = newToken(token.MINUS, l.ch)
    case '!':
        switch peekChar {
        case '=':
            tok = l.newTwoCharToken(token.NOT_EQ, l.ch)
        default:
            tok = newToken(token.BANG, l.ch)
        }
    case '/':
        switch peekChar { 
        case '/':
            l.skipComment()
            return l.NextToken()
        case '*':
            l.skipMultilineComment()
            return l.NextToken()
        default:
            tok = newToken(token.SLASH, l.ch)
        }
    case '*':
        switch peekChar {
        case '*':
            tok = l.newTwoCharToken(token.DOUBLE_ASTERISK, l.ch)
        default:
            tok = newToken(token.ASTERISK, l.ch)
        }
    case '<':
        switch peekChar {
        case '=':
            tok = l.newTwoCharToken(token.LT_EQ, l.ch)
        case '<':
            tok = l.newTwoCharToken(token.BITWISE_LEFT_SHIFT, l.ch)
        default:
            tok = newToken(token.LT, l.ch)
        }
    case '>':
        switch peekChar {
        case '=':
            tok = l.newTwoCharToken(token.GT_EQ, l.ch)
        case '>':
            tok = l.newTwoCharToken(token.BITWISE_RIGHT_SHIFT, l.ch)
        default:
            tok = newToken(token.GT, l.ch)
        }
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
    case ',':
        tok = newToken(token.COMMA, l.ch)
    case '(':
        tok = newToken(token.LPAREN, l.ch)
    case ')':
        tok = newToken(token.RPAREN, l.ch) 
    case '{':
        tok = newToken(token.LBRACE, l.ch)
    case '}':
        tok = newToken(token.RBRACE, l.ch)
    case '%':
        tok = newToken(token.MODULO, l.ch)
    case '&':
        switch peekChar {
        case '^':
            tok = l.newTwoCharToken(token.BITWISE_CLEAR, l.ch)
        default:
            tok = newToken(token.BITWISE_AND, l.ch)
        }
    case '|':
        tok = newToken(token.BITWISE_OR, l.ch)
    case '^':
        tok = newToken(token.BITWISE_XOR_NOT, l.ch) 
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    default: 
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Type = token.LookupIdent(tok.Literal)
            return tok
        } else if isDigit(l.ch) {
            tok.Type = token.INT
            tok.Literal = l.readNumber()
            return tok
        } else {
            tok = newToken(token.ILLEGAL, l.ch)
        }
    }

    l.readChar()
    return tok
}

func (l *Lexer) peekChar() byte {
    if l.readPosition >= len(l.input) {
        return 0
    } else {
        return l.input[l.readPosition]
    }
}

func (l* Lexer) skipMultilineComment() {
    for l.ch != 0 {
        l.readChar()
        if l.ch == '*' && l.peekChar() == '/' {
            l.readChar()
            l.readChar()
            break
        }
    }
}

func (l* Lexer) skipComment() {
    for l.ch != '\n' && l.ch != 0 {
        l.readChar()
    }
}

func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '-'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

func (l *Lexer) newTwoCharToken(tokenType token.TokenType, ch byte) token.Token {
    l.readChar() 
    return token.Token{Type: tokenType, Literal: string(ch) + string(l.ch)}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}
