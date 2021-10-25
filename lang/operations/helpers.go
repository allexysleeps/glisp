package operations

import "glisp/lang/shared"

func argValue(scope *shared.Scope, eval shared.Evaluator, arg shared.ExpArgument) shared.Value {
	var val shared.Value
	switch arg.Type() {
	case shared.TypeValue:
		val = arg.(shared.ArgValue).Value
	case shared.TypeVariable:
		val = scope.Get(arg.(shared.ArgVariable).Value)
	case shared.TypeExp:
		val = eval(scope, *arg.(shared.ArgExp).Value)
	}
	return val
}
