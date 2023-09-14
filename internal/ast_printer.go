package internal

import "fmt"

type AstPrinter struct {
}

func (astPrinter AstPrinter) Print(expr Expr) any {
	return expr.accept(astPrinter)
}

func (astPrinter AstPrinter) visitBinaryExpr(expr Binary) any {
	return astPrinter.parenthesize(expr.operator.lexeme, expr.left, expr.right)
}

func (astPrinter AstPrinter) visitGroupingExpr(expr Grouping) any {
	return astPrinter.parenthesize("group", expr.expresion)
}

func (astPrinter AstPrinter) visitLiteralExpr(expr Literal) any {
	if expr.value == nil {
		return "nil"
	}
	return expr.value.String()
}

func (astPrinter AstPrinter) visitUnaryExpr(expr Unary) any {
	return astPrinter.parenthesize(expr.operator.lexeme, expr.right)
}

func (astPrinter AstPrinter) parenthesize(name string, exprs ...Expr) string {
	str := "("

	str += name

	for _, expr := range exprs {
		str += " " + fmt.Sprintf("%v", expr.accept(astPrinter))
	}

	str += ")"
	return str
}
