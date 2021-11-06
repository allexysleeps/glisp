package shared

type varType uint

const (
	VarPrimitive varType = iota
	VarFn
)

type Variable interface {
	Type() varType
	Name() string
	Value() Value
	Exec() (*Expression, []string)
}

type primVar struct {
	t     varType
	name  string
	value Value
}

func (v *primVar) Type() varType {
	return v.t
}

func (v *primVar) Name() string {
	return v.name
}

func (v *primVar) Value() Value {
	return v.value
}

func (v *primVar) Exec() (*Expression, []string) {
	return nil, []string{}
}

type funcVar struct {
	t    varType
	name string
	exp  *Expression
	args []string
}

func (v *funcVar) Type() varType {
	return v.t
}

func (v *funcVar) Name() string {
	return v.name
}

func (v *funcVar) Value() Value {
	return nil
}

func (v *funcVar) Exec() (*Expression, []string) {
	return v.exp, v.args
}

func CreateValueVar(name string, value Value) Variable {
	return &primVar{t: VarPrimitive, name: name, value: value}
}

func CreateFunctionVar(name string, exp *Expression, args []string) Variable {
	return &funcVar{t: VarFn, name: name, exp: exp, args: args}
}
