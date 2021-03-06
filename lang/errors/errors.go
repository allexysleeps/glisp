package errors

import (
	"fmt"
	"log"
)

type errType string

const (
	ErrUndefined   errType = "ErrUndefined"
	ErrArgAmount   errType = "ErrArgAmount"
	ErrWrongSyntax errType = "ErrWrongSyntax"
)

type Err struct {
	T         errType
	Text      string
	Operation string
	Stack     *Err
}

func (e *Err) Print() {
	head := fmt.Sprintf("[%s]: \"%s\" Stack:\n%s", e.T, e.Text, e.Operation)
	body := fmt.Sprintf("\n%s", e.Operation)

	stackErr := e.Stack

	for stackErr != nil {
		if stackErr.Stack == nil {
			head = fmt.Sprintf("[%s]: \"%s\" Stack:", stackErr.T, stackErr.Text)
		}

		body += fmt.Sprintf("\n%s", stackErr.Operation)
		stackErr = stackErr.Stack
	}
	log.Printf("%s%s", head, body)
}

func CreateErrStack(operation string, stack *Err) *Err {
	return &Err{
		Operation: operation,
		Stack:     stack,
		T:         "",
		Text:      "",
	}
}

func CreateRootError(t errType, text, operation string) *Err {
	return &Err{
		T:         t,
		Text:      text,
		Operation: operation,
	}
}
