package token

import "fmt"


type Type string


const (
	
	LEFTPAREN  = "("
	RIGHTPAREN = ")"
	LEFTBRACE  = "{"
	RIGHTBRACE = "}"
	COMMA      = ","
	DOT        = "."
	MINUS      = "-"
	PLUS       = "+"
	SEMICOLON  = ";"
	SLASH      = "/"
	STAR       = "*"
	QMARK      = "?"
	COLON      = ":"
	HASH       = "#"
	
	BANG         = "!"
	AND			 = "&"
	OR			 = "|"
	BANGEQUAL    = "!="
	EQUAL        = "="
	EQUALEQUAL   = "=="
	COLONEQUAL     = ":="
	GREATER      = ">"
	GREATEREQUAL = ">="
	LESS         = "<"
	LESSEQUAL    = "<="
	POWER        = "^"
	
	IDENTIFIER = "IDENT"
	STRING     = "STRING"
	NUMBER     = "NUMBER"
	
	VOID	 = "void"
	FINAL	 = "final"
	ASK      = "ask"
	CLASS    = "class"
	EL       = "el"
	FALSE    = "false"
	FN       = "fn"
	LAMBDA	 = "lambda"
	IF       = "if"
	NIL      = "nil"
	PRINT    = "print"
	PRINTLN  = "println"
	RTN      = "rtn"
	SUPERCLASS    = "superClass"
	THIS     = "this"
	TRUE     = "true"
	REPEAT   = "repeat"
	INCLUDE  = "include"
	WHILE    = "while"
	BREAK    = "break"
	CONTINUE = "continue"
	IMPLEMENTS = "implements"
	EOF      = "eof"
	INVALID  = "__INVALID__"
)


type Token struct {
	Type    Type
	Lexeme  string
	Literal interface{}
	Line    int
}

func (token *Token) String() string {
	return fmt.Sprintf("%s %s %v", token.Type, token.Lexeme, token.Literal)
}
