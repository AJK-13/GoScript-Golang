package ast

import (
	"fmt"
	"GoScript/src/token"
	"strings"
)

type Node interface {
	String() string
}

type Expr interface {
	Node
}

type Binary struct {
	Expr
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (b *Binary) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(b.Operator.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(b.Left.String())
	sb.WriteString(" ")
	sb.WriteString(b.Right.String())
	sb.WriteString(")")
	return sb.String()
}

type Grouping struct {
	Expr
	Expression Expr
}

func (g *Grouping) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(g.Expression.String())
	sb.WriteString(")")
	return sb.String()
}

type Literal struct {
	Expr
	Value interface{}
}

func (l *Literal) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%v", l.Value))
	return sb.String()
}

type Unary struct {
	Expr
	Operator token.Token
	Right    Expr
}

func (u *Unary) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(u.Operator.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(u.Right.String())
	sb.WriteString(")")
	return sb.String()
}

type Ternary struct {
	Expr
	Condition Expr
	QMark     token.Token
	Then      Expr
	Colon     token.Token
	Else      Expr
}


func (t *Ternary) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(t.Condition.String())
	sb.WriteString(" ")
	sb.WriteString(t.QMark.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(t.Then.String())
	sb.WriteString(" ")
	sb.WriteString(t.Colon.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(t.Else.String())
	sb.WriteString(")")
	return sb.String()
}

type Assign struct {
	Expr
	Name     token.Token
	Value    Expr
	EnvIndex int
	EnvDepth int
}

func (a *Assign) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("=")
	sb.WriteString(" ")
	sb.WriteString(a.Name.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(a.Value.String())
	sb.WriteString(")")
	return sb.String()
}

type Variable struct {
	Expr
	Name     token.Token
	EnvIndex int
	EnvDepth int
}
func (v *Variable) String() string {
	var sb strings.Builder
	sb.WriteString(v.Name.Lexeme)
	return sb.String()
}

type Stmt interface {
	Node
}

type Block struct {
	Stmt
	Statements []Stmt
	EnvSize    int
}

func (b *Block) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	for _, stmt := range b.Statements {
		sb.WriteString(stmt.String())
	}
	sb.WriteString(")")
	return sb.String()
}

type Expression struct {
	Stmt
	Expression Expr
}

func (e *Expression) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(e.Expression.String())
	sb.WriteString(")")
	return sb.String()
}

type Print struct {
	Stmt
	Expression Expr
}
type Println struct {
	Stmt
	Expression Expr
}
type Ask struct {
	Stmt
	Expression Expr
}
type AskNum struct {
	Stmt
	Expression Expr
}
type Lambda struct {
	Stmt
	Expression Expr
}
func (p *Print) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("print")
	sb.WriteString(" ")
	sb.WriteString(p.Expression.String())
	sb.WriteString(")")
	return sb.String()
}

func (p *Println) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("println")
	sb.WriteString(" ")
	sb.WriteString(p.Expression.String())
	sb.WriteString(")")
	return sb.String()
}

type Var struct {
	Stmt
	Name        token.Token
	Initializer Expr
	EnvIndex    int
}
type Final struct {
	Stmt
	Name        token.Token
	Initializer Expr
	EnvIndex    int
}

func (v *Var) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("var")
	sb.WriteString(" ")
	sb.WriteString(v.Name.Lexeme)
	sb.WriteString(" ")
	if v.Initializer != nil {
		sb.WriteString(v.Initializer.String())
	} else {
		sb.WriteString("nil")
	}
	sb.WriteString(")")
	return sb.String()
}

type If struct {
	Stmt
	Condition  Expr
	ThenBranch Stmt
	ElseBranch Stmt
}

func (i *If) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("if")
	sb.WriteString(" ")
	sb.WriteString(i.Condition.String())
	sb.WriteString(" ")
	sb.WriteString(i.ThenBranch.String())
	sb.WriteString(" ")
	sb.WriteString(i.ElseBranch.String())
	sb.WriteString(")")
	return sb.String()
}

type For struct {
	Stmt
	Initializer Expr
	Condition   Expr
	Increment   Expr
	Statement   Stmt
}

func (f *For) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("for")
	sb.WriteString(" ")
	sb.WriteString("(")
	sb.WriteString(f.Initializer.String())
	sb.WriteString(")")
	sb.WriteString(" ")
	sb.WriteString("(")
	sb.WriteString(f.Condition.String())
	sb.WriteString(")")
	sb.WriteString(" ")
	sb.WriteString("(")
	sb.WriteString(f.Increment.String())
	sb.WriteString(")")
	sb.WriteString(" ")
	sb.WriteString(f.Statement.String())
	sb.WriteString(")")
	return sb.String()
}

type While struct {
	Stmt
	Condition Expr
	Statement Stmt
}

func (w *While) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("while")
	sb.WriteString(" ")
	sb.WriteString(w.Condition.String())
	sb.WriteString(" ")
	sb.WriteString(w.Statement.String())
	sb.WriteString(")")
	return sb.String()
}

type Logical struct {
	Expr
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (l *Logical) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(l.Operator.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(l.Left.String())
	sb.WriteString(" ")
	sb.WriteString(l.Right.String())
	sb.WriteString(")")
	return sb.String()
}

type Call struct {
	Callee    Expr
	Paren     token.Token
	Arguments []Expr
}

func (c *Call) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("call")
	sb.WriteString(" ")
	sb.WriteString(c.Callee.String())
	sb.WriteString(" ")
	for _, e := range c.Arguments {
		sb.WriteString(e.String())
		sb.WriteString(" ")
	}
	sb.WriteString(")")
	return sb.String()
}

type Function struct {
	Name          token.Token
	Params        []token.Token
	Body          []Stmt
	EnvSize       int
	EnvIndex      int
	IsClassMethod bool
}

func (f *Function) IsProperty() bool {
	return f.Params == nil
}

func (f *Function) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("fn")
	sb.WriteString(" ")
	sb.WriteString(f.Name.Lexeme)
	sb.WriteString(" ")
	sb.WriteString("(")
	for _, p := range f.Params {
		sb.WriteString(p.Lexeme)
		sb.WriteString(" ")
	}
	sb.WriteString(")")
	sb.WriteString(" ")
	sb.WriteString("(")
	for _, stmt := range f.Body {
		sb.WriteString(stmt.String())
		sb.WriteString(" ")
	}
	sb.WriteString(")")
	sb.WriteString(")")
	return sb.String()
}

type Return struct {
	Stmt
	Keyword token.Token
	Value   Expr
}

func (r *Return) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("rtn")
	sb.WriteString(" ")
	sb.WriteString(r.Value.String())
	sb.WriteString(" ")
	sb.WriteString(")")
	return sb.String()
}

type Break struct {
	Stmt
	Token token.Token
}

func (b *Break) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("break")
	sb.WriteString(")")
	return sb.String()
}

type Continue struct {
	Stmt
	Token token.Token
}

func (c *Continue) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("continue")
	sb.WriteString(")")
	return sb.String()
}

type Class struct {
	Stmt
	Name         token.Token
	Methods      []*Function
	ClassMethods []*Function
	EnvIndex     int
	SuperClass   *Variable
}

func (c *Class) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("class")
	sb.WriteString("")
	sb.WriteString(c.Name.Lexeme)
	sb.WriteString("")
	for _, f := range c.Methods {
		sb.WriteString(f.String())
		sb.WriteString(" ")
	}
	sb.WriteString("")
	for _, f := range c.ClassMethods {
		sb.WriteString(f.String())
		sb.WriteString(" ")
	}
	sb.WriteString(")")
	return sb.String()
}

type Get struct {
	Expr
	Name       token.Token
	Expression Expr
}

func (g *Get) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(".")
	sb.WriteString(" ")
	sb.WriteString(g.Expression.String())
	sb.WriteString(" ")
	sb.WriteString(g.Name.Lexeme)
	sb.WriteString(")")
	return sb.String()
}

type Set struct {
	Expr
	Object Expr
	Name   token.Token
	Value  Expr
}

func (s *Set) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("set")
	sb.WriteString(" ")
	sb.WriteString(s.Object.String())
	sb.WriteString(" ")
	sb.WriteString(s.Name.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(s.Value.String())
	sb.WriteString(")")
	return sb.String()
}

type This struct {
	Expr
	Keyword  token.Token
	EnvIndex int
	EnvDepth int
}

func (t *This) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("this")
	sb.WriteString(" ")
	sb.WriteString(t.Keyword.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(")")
	return sb.String()
}

type Super struct {
	Expr
	Keyword  token.Token
	Method   token.Token
	EnvIndex int
	EnvDepth int
}

func (s *Super) String() string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString("super")
	sb.WriteString(" ")
	sb.WriteString(s.Method.Lexeme)
	sb.WriteString(" ")
	sb.WriteString(")")
	return sb.String()
}
