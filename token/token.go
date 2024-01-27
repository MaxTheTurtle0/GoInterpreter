package token

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // Identifiers + literals
    IDENT = "IDENT" // add, foobar, x, y, ...
    INT = "INT" // 1234603

    // Operators
    ASSIGN = "="
    PLUS = "+"
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*" 
    SLASH = "/" 
    DOUBLE_ASTERISK = "**"

    // Bitwise Operators
    BITWISE_AND = "&"
    BITWISE_OR = "|"
    BITWISE_XOR_NOT = "^"
    BITWISE_CLEAR = "&^"
    BITWISE_LEFT_SHIFT = "<<"
    BITWISE_RIGHT_SHIFT = ">>"
    
    LT = "<"
    GT = ">"
    LT_EQ = "<=" 
    GT_EQ = ">="
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

var keywords = map[string]TokenType{
    "fn": FUNCTION,
    "let": LET,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}
