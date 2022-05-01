package mathexp

import (
	"errors"
	"reflect"
)

var CondOps = map[string][]reflect.Type{
	"and": {reflect.TypeOf(ConditionSpec{})},
	"or":  {reflect.TypeOf(ConditionSpec{})},
	"eq":  {reflect.TypeOf(reflect.String), reflect.TypeOf(reflect.Float64), reflect.TypeOf(reflect.Int64), reflect.TypeOf(reflect.Bool)},
	"neq": {reflect.TypeOf(reflect.String), reflect.TypeOf(reflect.Float64), reflect.TypeOf(reflect.Int64), reflect.TypeOf(reflect.Bool)},
	"lt":  {reflect.TypeOf(reflect.String), reflect.TypeOf(reflect.Float64), reflect.TypeOf(reflect.Int64)},
	"lte": {reflect.TypeOf(reflect.String), reflect.TypeOf(reflect.Float64), reflect.TypeOf(reflect.Int64)},
	"gt":  {reflect.TypeOf(reflect.String), reflect.TypeOf(reflect.Float64), reflect.TypeOf(reflect.Int64)},
	"gte": {reflect.TypeOf(reflect.String), reflect.TypeOf(reflect.Float64), reflect.TypeOf(reflect.Int64)},
}

var MathOps = []string{"add", "sub", "mul", "div"}

const (
	VarTypIn     = "in"
	VarTypOut    = "out"
	VarTypConst  = "const"
	VarTypExpOut = "expout"
)

var (
	ErrExpNotValid = errors.New("Errors in the expression")
)
