package scanner

import (
	"fmt"
	"GoScript/src/parseerror"
	"GoScript/src/token"
	"strconv"
)

var keywords = map[string]token.Type{
	"void":      token.VOID,
	"final":      token.FINAL,
	"class":    token.CLASS,
	"el":     token.EL,
	"false":    token.FALSE,
	"repeat":      token.REPEAT,
	"fn":      token.FN,
	"if":       token.IF,
	"nil":      token.NIL,
	"print":    token.PRINT,
	"println":      token.PRINTLN,
	"rtn":   token.RTN,
	"superClass":    token.SUPERCLASS,
	"this":     token.THIS,
	"true":     token.TRUE,
	"ask":      token.ASK,
	"while":    token.WHILE,
	"break":    token.BREAK,
	"lambda":   token.LAMBDA,
	"include":  token.INCLUDE,
	"continue": token.CONTINUE,
}


type Scanner struct {
	source  string
	start   int
	current int
	line    int
	tokens  []token.Token
}


func New(source string) Scanner {
	scanner := Scanner{source: source, line: 1, tokens: make([]token.Token, 0)}
	return scanner
}



func (sc *Scanner) ScanTokens() []token.Token {
	for !sc.isAtEnd() {
		
		sc.start = sc.current
		sc.scanToken()
	}
	sc.tokens = append(sc.tokens, token.Token{Type: token.EOF})
	return sc.tokens
}

func (sc *Scanner) makeToken(tp token.Type) token.Token {
	lexeme := sc.source[sc.start:sc.current]
	return token.Token{Type: tp, Lexeme: lexeme, Line: sc.line}
}

func (sc *Scanner) addToken(tp token.Type) {
	sc.addTokenWithLiteral(tp, nil)
}

func (sc *Scanner) addTokenWithLiteral(tp token.Type, literal interface{}) {
	text := sc.source[sc.start:sc.current]
	sc.tokens = append(sc.tokens, token.Token{Type: tp, Lexeme: text, Literal: literal, Line: sc.line})
}

func (sc *Scanner) scanString() {
	for sc.peek() != '"' && !sc.isAtEnd() {
		if sc.peek() == '\n' {
			sc.line++
		}
		sc.advance()
	}

	
	if sc.isAtEnd() {
		parseerror.LogMessage(sc.line, "Unterminated string")
		return
	}

	
	sc.advance()

	
	value := sc.source[sc.start+1 : sc.current-1]
	sc.addTokenWithLiteral(token.STRING, value)
}

func (sc *Scanner) scanNumber() {
	for sc.isDigit(sc.peek()) {
		sc.advance()
	}

	
	if sc.peek() == '.' && sc.isDigit(sc.peekNext()) {
		sc.advance() 
		for sc.isDigit(sc.peek()) {
			sc.advance()
		}
	}

	number, err := strconv.ParseFloat(sc.source[sc.start:sc.current], 64)
	if err != nil {
		panic("Invalid number format")
	} else {
		sc.addTokenWithLiteral(token.NUMBER, number)
	}
}

func (sc *Scanner) scanIdentifier() {
	for sc.isAlphaNumeric(sc.peek()) {
		sc.advance()
	}

	
	text := sc.source[sc.start:sc.current]
	tp, ok := keywords[text]
	if ok {
		sc.addToken(tp)
	} else {
		sc.addToken(token.IDENTIFIER)
	}
}

func (sc *Scanner) scanToken() {
	c := sc.advance()

	switch c {
	case '(':
		sc.addToken(token.LEFTPAREN)
	case ')':
		sc.addToken(token.RIGHTPAREN)
	case '{':
		sc.addToken(token.LEFTBRACE)
	case '}':
		sc.addToken(token.RIGHTBRACE)
	case ',':
		sc.addToken(token.COMMA)
	case '.':
		sc.addToken(token.DOT)
	case '-':
		sc.addToken(token.MINUS)
	case '+':
		sc.addToken(token.PLUS)
	case '?':
		sc.addToken(token.QMARK)
	case ':':
		if sc.match('=') {
			sc.addToken(token.COLONEQUAL)
		} else {
			sc.addToken(token.COLON)
		}
	case ';':
		sc.addToken(token.SEMICOLON)
	case '^':
		sc.addToken(token.POWER)
	case '*':
		sc.addToken(token.STAR)
	case '!':
		if sc.match('=') {
			sc.addToken(token.BANGEQUAL)
		} else if sc.match('!') {
			for sc.peek() != '\n' && !sc.isAtEnd() {
				sc.advance()
			}
 		} else if sc.match('*') {
			for sc.peek() != '*' && sc.peek() + 1 != '!' && !sc.isAtEnd() {
				sc.advance()
			}
 		} else {
			sc.addToken(token.BANG)
		}
	case '=':
		if sc.match('=') {
			sc.addToken(token.EQUALEQUAL)
		} else {
			sc.addToken(token.EQUAL)
		}
	case '<':
		if sc.match('=') {
			sc.addToken(token.LESSEQUAL)
		} else {
			sc.addToken(token.LESS)
		}
	case '>':
		if sc.match('=') {
			sc.addToken(token.GREATEREQUAL)
		} else {
			sc.addToken(token.GREATER)
		}
	case '/':
		sc.addToken(token.SLASH)
	case '\n':
		sc.line++
	case ' ', '\r', '\t':
		
	case '"':
		sc.scanString()
	default:
		if sc.isDigit(c) {
			sc.scanNumber()
		} else if sc.isAlpha(c) {
			sc.scanIdentifier()
		} else {
			parseerror.LogMessage(sc.line, fmt.Sprintf("Unexpected character: %c", c))
		}
	}
}

func (sc *Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (sc *Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func (sc *Scanner) isAlphaNumeric(c byte) bool {
	return sc.isAlpha(c) || sc.isDigit(c)
}

func (sc *Scanner) isAtEnd() bool {
	return sc.current >= len(sc.source)
}


func (sc *Scanner) advance() byte {
	sc.current++
	return sc.source[sc.current-1]
}

func (sc *Scanner) match(expected byte) bool {
	if sc.isAtEnd() {
		return false
	}
	if sc.source[sc.current] != expected {
		return false
	}
	sc.current++
	return true
}

func (sc *Scanner) peek() byte {
	if sc.isAtEnd() {
		return 0
	}
	return sc.source[sc.current]
}

func (sc *Scanner) peekNext() byte {
	if sc.current+1 >= len(sc.source) {
		return 0
	}
	return sc.source[sc.current+1]
}
