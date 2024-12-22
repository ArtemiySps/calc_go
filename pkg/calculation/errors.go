package calculation

import "errors"

var (
	ErrDivisionByZero      = errors.New("division by zero")
	ErrUnexpectedSymbol    = errors.New("unexpected symbol")
	ErrTwoOperationsInARow = errors.New("two operations in a row")
	ErrOperationAtTheEnd   = errors.New("operation at the end of expression")
	ErrBracketsProblem     = errors.New("the number of open and closed brackets is different")
	ErrNoNumbers           = errors.New("no numbers in expresion")
	ErrServerError         = errors.New("unknown error")
)
