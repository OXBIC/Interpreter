package main

import "fmt"

// Nuestro Lenguage Ficticio
// program → expr*
//
// expr → assignment | infixExpr | int
//
// assignment → id = expr ;
//
// infixExp → expr infixOp expr ;
//
// infixOp → + | - | * | /

// A program is a list of zero or more expressions.
// An expression is either an assignment, an infix expression, or an integer. And so on…
func main() {
	fmt.Println("Hello World")
}
