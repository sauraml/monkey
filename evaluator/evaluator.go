package evaluator

import (
	"github.com/sauraml/monkey/ast"
	"github.com/sauraml/monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Satements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)
	}

	return result
}
