package shared

import (
	"strconv"
	"strings"
)

type ValType string

const (
	TypeNull ValType = "null"
	TypeNum  ValType = "num"
	TypeBool ValType = "bool"
	TypeStr  ValType = "str"
)

type Value interface {
	Type() ValType
	IsNull() bool
	NumVal() float64
	BoolVal() bool
	StrVal() string
}

type value struct {
	t       ValType
	numVal  float64
	boolVal bool
	strVal  string
}

func (v value) Type() ValType { return v.t }
func (v value) IsNull() bool  { return v.t == TypeNull }

func (v value) NumVal() float64 {
	switch v.t {
	case TypeNum:
		return v.numVal
	case TypeStr:
		num, err := strconv.ParseFloat(v.strVal, 64)
		if err != nil {
			return 0
		}
		return num
	case TypeBool:
		if v.boolVal {
			return 1
		}
		return 0
	}
	return 0
}

func (v value) BoolVal() bool {
	switch v.t {
	case TypeBool:
		return v.boolVal
	case TypeStr:
		if v.strVal != "" {
			return true
		}
		return false
	case TypeNum:
		if v.numVal > 0 {
			return true
		}
		return false
	}
	return false
}

func (v value) StrVal() string {
	switch v.t {
	case TypeStr:
		return v.strVal
	case TypeBool:
		if v.boolVal {
			return "true"
		}
		return "false"
	case TypeNum:
		return strconv.FormatFloat(v.numVal, 'f', -1, 64)
	}
	return ""
}

func CreateValue(s string) (Value, bool) {
	num, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return value{t: TypeNum, numVal: num}, true
	}
	if s == "true" {
		return value{t: TypeBool, boolVal: true}, true
	}
	if s == "false" {
		return value{t: TypeBool, boolVal: false}, true
	}
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return value{t: TypeStr, strVal: s}, true
	}
	return value{}, false
}

func CreateValueOfType(t ValType, val interface{}) Value {
	var v Value
	switch t {
	case TypeStr:
		v = value{t: t, strVal: val.(string)}
	case TypeNum:
		v = value{t: t, numVal: val.(float64)}
	case TypeBool:
		v = value{t: t, boolVal: val.(bool)}
	case TypeNull:
		v = value{t: t}
	}
	return v
}
