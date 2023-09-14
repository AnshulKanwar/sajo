package internal

import "fmt"

// TODO: use Generics
type Expr interface {
	accept(visitor ExprVisitor) any
}

type Binary struct {
	left     Expr
	operator Token
	right    Expr
}

type Grouping struct {
	expresion Expr
}

type Literal struct {
	value interface{ fmt.Stringer }
}

type Unary struct {
	operator Token
	right    Expr
}

func (binary Binary) accept(visitor ExprVisitor) any {
	return visitor.visitBinaryExpr(binary)
}

func (grouping Grouping) accept(visitor ExprVisitor) any {
	return visitor.visitGroupingExpr(grouping)
}

func (literal Literal) accept(visitor ExprVisitor) any {
	return visitor.visitLiteralExpr(literal)
}

func (unary Unary) accept(visitor ExprVisitor) any {
	return visitor.visitUnaryExpr(unary)
}

// TODO: use Generics
type ExprVisitor interface {
	visitBinaryExpr(expr Binary) any
	visitGroupingExpr(expr Grouping) any
	visitLiteralExpr(expr Literal) any
	visitUnaryExpr(expr Unary) any
}
