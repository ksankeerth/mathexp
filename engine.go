package mathexp

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type CondOp string

type ConditionSpec struct {
	Op string
	V1 interface{}
	V2 interface{}
}

func (cond *ConditionSpec) String() string {
	return fmt.Sprintf("%#v", cond)
}

func (cond *ConditionSpec) isValid() bool {
	valid := cond != nil
	if valid {
		types, ok := CondOps[cond.Op]
		if ok {
			v1Typ := reflect.TypeOf(cond.V1)
			v2Typ := reflect.TypeOf(cond.V2)
			if v1Typ == v2Typ {
				return isTypeAllowed(types, v1Typ)

			}
		}
	}
	// TODO : need to consider that one value can be input and const
	// TODO : verify all the inputs are available in vars
	// TODO : same approach followed in expression.isValid() can be used
	return valid
}

func (cond *ConditionSpec) evaluate(vals map[string]*Value) (bool, error) {
	//TODO : type checks for v1 v2
	// TODO : type evaluate for bool, numbers(int and float) , string and cs
	// Always it's expected that v1 and v2 are belongs to same type

	// Algorithm
	// if v1 or v2 is ConditionSpec, then both should be ConditionSpect if return an error
	// Next check v1 and v2 are value or reference?
	// v1 or v2 are reference => resolve
	// v1 or v2 are values => just move to next step
	// after resolve the values, typeof v1Val and typeof v2Val should be the same
	// finally evaluate

	if v1Cond, ok := cond.V1.(ConditionSpec); ok {
		if v2Cond, ok := cond.V2.(ConditionSpec); ok {
			v1Match, err := v1Cond.evaluate(vals)
			if err != nil {
				return false, err
			}
			v2Match, err := v2Cond.evaluate(vals)
			if err != nil {
				return false, err
			}
			return condEval(cond.Op, v1Match, v2Match)
		} else {
			return false, errors.New("both should be condition")
		}
	}
	v1Str, v1StrOk := cond.V1.(string)
	v2Str, v2StrOk := cond.V2.(string)

	var v1Val interface{}
	var v2Val interface{}
	v1Val = v1Str
	v2Val = v2Str

	if v1StrOk && v2StrOk {
		resolved := resolveRef(vals, v1Str, v2Str)
		if rv, ok := resolved[v1Str]; ok {
			v1Val = rv.val
		}
		if rv, ok := resolved[v2Str]; ok {
			v2Val = rv.val
		}
		return condEval(cond.Op, v1Val, v2Val)
	} else if v1StrOk {
		resolved := resolveRef(vals, v1Str, v2Str)
		if rv, ok := resolved[v1Str]; ok {
			v1Val = rv.val
		}
		return condEval(cond.Op, v1Val, v2Val)
	} else if v2StrOk {
		resolved := resolveRef(vals, v1Str, v2Str)

		if rv, ok := resolved[v2Str]; ok {
			v2Val = rv.val
		}
		return condEval(cond.Op, v1Val, v2Val)
	}

	return condEval(cond.Op, cond.V1, cond.V2)

}

func resolveRef(vals map[string]*Value, vars ...string) map[string]*Value {
	resolved := make(map[string]*Value)
	for _, v := range vars {
		val, ok := vals[v]
		if ok {
			resolved[v] = val
		}
	}
	return resolved
}

func condEval(op string, v1 interface{}, v2 interface{}) (bool, error) {
	switch op {
	case "and":
		switch v1.(type) {
		case string:
			return true, nil

		case bool:
			v1Bool, v1ok := v1.(bool)
			v2Bool, v2ok := v2.(bool)
			if v1ok && v2ok {
				return v1Bool && v2Bool, nil
			}
		case float64:
			return true, nil
		case int:
			return true, nil
		}
	case "or":
		switch v1.(type) {
		case string:
			return true, nil
		case bool:
			v1Bool, v1ok := v1.(bool)
			v2Bool, v2ok := v2.(bool)
			if v1ok && v2ok {
				return v1Bool || v2Bool, nil
			}
		case float64:
			return true, nil
		case int:
			return true, nil

		}
	case "eq":
		switch v1.(type) {
		case string:
			v1Str, v1ok := v1.(string)
			v2Str, v2ok := v2.(string)
			if v1ok && v2ok {
				return v1Str == v2Str, nil
			}
		case bool:
			v1Bool, v1ok := v1.(bool)
			v2Bool, v2ok := v2.(bool)
			if v1ok && v2ok {
				return v1Bool == v2Bool, nil
			}
		case float64:
			v1Float, v1ok := v1.(float64)
			v2Float, v2ok := v2.(float64)
			if v1ok && v2ok {
				return v1Float == v2Float, nil
			}
		case int:
			v1Int, v1ok := v1.(int)
			v2Int, v2ok := v2.(int)
			if v1ok && v2ok {
				return v1Int == v2Int, nil
			}
		}
	case "neq":
		switch v1.(type) {
		case string:
			v1Str, v1ok := v1.(string)
			v2Str, v2ok := v2.(string)
			if v1ok && v2ok {
				return v1Str != v2Str, nil
			}
		case bool:
			v1Bool, v1ok := v1.(bool)
			v2Bool, v2ok := v2.(bool)
			if v1ok && v2ok {
				return v1Bool != v2Bool, nil
			}
		case float64:
			v1Float, v1ok := v1.(float64)
			v2Float, v2ok := v2.(float64)
			if v1ok && v2ok {
				return v1Float != v2Float, nil
			}
		case int:
			v1Int, v1ok := v1.(int)
			v2Int, v2ok := v2.(int)
			if v1ok && v2ok {
				return v1Int != v2Int, nil
			}
		}
	case "lt":
		switch v1.(type) {
		case string:
			v1Str, v1ok := v1.(string)
			v2Str, v2ok := v2.(string)
			if v1ok && v2ok {
				return strings.Compare(v1Str, v2Str) < 0, nil
			}
		case bool:
			v1Bool, v1ok := v1.(bool)
			v2Bool, v2ok := v2.(bool)
			if v1ok && v2ok {
				return v1Bool == false && v2Bool == true, nil
			}
		case float64:
			v1Float, v1ok := v1.(float64)
			v2Float, v2ok := v2.(float64)
			if v1ok && v2ok {
				return v1Float < v2Float, nil
			}
		case int:
			v1Int, v1ok := v1.(int)
			v2Int, v2ok := v2.(int)
			if v1ok && v2ok {
				return v1Int < v2Int, nil
			}
		}
	case "lte":
		switch v1.(type) {
		case string:
			v1Str, v1ok := v1.(string)
			v2Str, v2ok := v2.(string)
			if v1ok && v2ok {
				return strings.Compare(v1Str, v2Str) <= 0, nil
			}
		case bool:
			v1Bool, v1ok := v1.(bool)
			v2Bool, v2ok := v2.(bool)
			if v1ok && v2ok {
				return v1Bool == false && v2Bool == true, nil
			}
		case float64:
			v1Float, v1ok := v1.(float64)
			v2Float, v2ok := v2.(float64)
			if v1ok && v2ok {
				return v1Float <= v2Float, nil
			}
		case int:
			v1Int, v1ok := v1.(int)
			v2Int, v2ok := v2.(int)
			if v1ok && v2ok {
				return v1Int <= v2Int, nil
			}
		}
	case "gt":
		switch v1.(type) {
		case string:
			v1Str, v1ok := v1.(string)
			v2Str, v2ok := v2.(string)
			if v1ok && v2ok {
				return strings.Compare(v1Str, v2Str) > 0, nil
			}
		case bool:
			v1Bool, v1ok := v1.(bool)
			v2Bool, v2ok := v2.(bool)
			if v1ok && v2ok {
				return v1Bool == true && v2Bool == false, nil
			}
		case float64:
			v1Float, v1ok := v1.(float64)
			v2Float, v2ok := v2.(float64)
			if v1ok && v2ok {
				return v1Float > v2Float, nil
			}
		case int:
			v1Int, v1ok := v1.(int)
			v2Int, v2ok := v2.(int)
			if v1ok && v2ok {
				return v1Int > v2Int, nil
			}
		}
	case "gte":
		switch v1.(type) {
		case string:
			v1Str, v1ok := v1.(string)
			v2Str, v2ok := v2.(string)
			if v1ok && v2ok {
				return strings.Compare(v1Str, v2Str) >= 0, nil
			}
		case bool:
			v1Bool, v1ok := v1.(bool)
			v2Bool, v2ok := v2.(bool)
			if v1ok && v2ok {
				return v1Bool == true && v2Bool == false, nil
			}
		case float64:
			v1Float, v1ok := v1.(float64)
			v2Float, v2ok := v2.(float64)
			if v1ok && v2ok {
				return v1Float >= v2Float, nil
			}
		case int:
			v1Int, v1ok := v1.(int)
			v2Int, v2ok := v2.(int)
			if v1ok && v2ok {
				return v1Int >= v2Int, nil
			}
		}
	}
	return false, nil
}

type VarSpec struct {
	Type  string
	Name  string
	Sym   string
	Value interface{}
	// TODO : add unit information
}

func (vs *VarSpec) String() string {
	return fmt.Sprintf("%#v", vs)
}

func (vs *VarSpec) isValid() bool {
	valid := false
	if vs != nil {
		switch vs.Type {
		case VarTypIn:
			valid = reflect.TypeOf(vs.Value) == reflect.TypeOf(nil)
		case VarTypOut:

			valid = reflect.TypeOf(vs.Value) == reflect.TypeOf(nil)
		case VarTypConst:
			valid = vs.Value != nil
		case VarTypExpOut:
			// valid = reflect.TypeOf(vs.Value) == reflect.TypeOf(nil)
		}
		// TODO : check regex for name
		// TODO : check regex for sym
	}
	return valid
}

type ExpresionSpec struct {
	Op  string
	V1  string
	V2  string
	Out string
}

func (es *ExpresionSpec) evaluate(values map[string]*Value) error {
	var err error

	v1Val, v1Ok := values[es.V1]
	v2Val, v2Ok := values[es.V2]

	if v1Ok && v2Ok {
		out := mathEval(es.Op, v1Val, v2Val)
		values[es.Out] = &Value{
			val:     out.val,
			typ:     reflect.TypeOf(out.val),
			sym:     values[es.Out].sym,
			varType: values[es.Out].varType,
		}

	} else {
		err = errors.New("Input values are not found")
	}

	return err
}

func mathEval(op string, v1 *Value, v2 *Value) *Value {
	// TODO type checking
	var out *Value
	switch op {
	case "add":
		v1Float, v1ok := v1.val.(float64)
		v2Float, v2ok := v2.val.(float64)
		if v1ok && v2ok {
			out = &Value{
				val: v1Float + v2Float,
			}
		} else {
			out = &Value{val: 0}
		}
	case "sub":
		v1Float, v1ok := v1.val.(float64)
		v2Float, v2ok := v2.val.(float64)
		if v1ok && v2ok {
			out = &Value{
				val: v1Float - v2Float,
			}
		} else {
			out = &Value{val: 0}
		}
	case "mul":
		v1Float, v1ok := v1.val.(float64)
		v2Float, v2ok := v2.val.(float64)
		if v1ok && v2ok {
			out = &Value{
				val: v1Float * v2Float,
			}
		} else {
			out = &Value{val: 0}
		}
	case "div":
		v1Float, v1ok := v1.val.(float64)
		v2Float, v2ok := v2.val.(float64)
		if v1ok && v2ok && v2Float != 0.0 {
			out = &Value{
				val: v1Float / v2Float,
			}
		} else {
			out = &Value{val: 0}
		}
	}
	return out
}

func (es *ExpresionSpec) String() string {
	return fmt.Sprintf("%#v", es)
}

func (es *ExpresionSpec) isValid(vars []*VarSpec) bool {
	valid := false
	if es != nil {
		if reflect.TypeOf(es.V1) == reflect.TypeOf(es.V2) {
			if strings.Contains(strings.Join(MathOps, ","), es.Op+",") {
				stop := false
				v1Valid := false
				v2Valid := false
				outValid := false

				iterateVars(vars, &stop, func(vs *VarSpec) {

					if !v1Valid && vs.Sym == es.V1 && vs.Type != VarTypOut {
						v1Valid = true
					}
					if !v2Valid && vs.Sym == es.V2 && vs.Type != VarTypOut {
						v2Valid = true
					}
					if !outValid && vs.Sym == es.Out && (vs.Type == VarTypOut || vs.Type == VarTypExpOut) {
						outValid = true
					}
					if v1Valid && v2Valid && outValid {
						stop = true
						valid = true
					}

				})

			}
		}
	}
	return valid
}

func iterateVars(vars []*VarSpec, stop *bool, matcher func(value *VarSpec)) {
	for _, vs := range vars {
		if *stop {
			return
		}
		matcher(vs)
	}
}

type ConditionGroupSpec struct {
	Cond               *ConditionSpec
	Vars               []*VarSpec
	Expressions        []*ExpresionSpec
	SubConditionGroups []*ConditionGroupSpec
	parent             *ConditionGroupSpec
}

type traveler func(cg *ConditionGroupSpec, vars []*VarSpec) bool

func (cg *ConditionGroupSpec) traverse(vars *[]*VarSpec, traveler traveler) {

	if *vars == nil {
		*vars = make([]*VarSpec, len(cg.Vars))
		copy(*vars, cg.Vars)
	} else {
		*vars = append(*vars, cg.Vars...)
	}
	stop := traveler(cg, *vars)
	if cg.SubConditionGroups == nil || stop {
		return
	}
	for _, subCg := range cg.SubConditionGroups {
		subCg.traverse(vars, traveler)
	}
}

func (cg *ConditionGroupSpec) isRoot() bool {
	return cg.parent == nil
}

func (cg *ConditionGroupSpec) isValid() (bool, error) {
	valid := false
	var err error
	var vars []*VarSpec
	cg.traverse(&vars, func(cg *ConditionGroupSpec, vars []*VarSpec) bool {

		for _, vs := range cg.Vars {
			if !vs.isValid() {
				err = errors.New("Var " + vs.String() + " is not valid")
				return true
			}
		}
		for _, exp := range cg.Expressions {
			if !exp.isValid(vars) {
				err = errors.New("Expression " + exp.String() + " is not valid")
				return true
			}
		}
		if !cg.isRoot() {
			if !cg.Cond.isValid() {
				err = errors.New("Condition " + cg.Cond.String() + " is not valid")
				return true
			}
		}
		valid = true
		return false
	})
	//TODO : return a descriptive error message
	return valid, err
}

func allVars(cg *ConditionGroupSpec) []*VarSpec {
	var allVars []*VarSpec
	cg.traverse(&allVars, func(cg *ConditionGroupSpec, vars []*VarSpec) bool {
		return false
	})
	return allVars
}

func verifyBeforeEvalute(vars []*VarSpec, args map[string]interface{}) bool {
	requiredArgs := 0
	for _, vs := range vars {
		if vs.Type == VarTypIn {
			if _, ok := args[vs.Sym]; ok {
				requiredArgs++
			}
		}
	}
	return requiredArgs == len(args)
}

func (cg *ConditionGroupSpec) evaluate(values map[string]*Value) (map[string]interface{}, error) {

	next := cg.isRoot()
	if !next {
		ok, err := cg.Cond.evaluate(values)
		if ok {
			next = true
		}
		return nil, err
	}

	// Condition has matched or RootNode
	if next {
		for _, exp := range cg.Expressions {
			err := exp.evaluate(values)
			if err != nil {
				// TODO : add some reasons or error message
				return nil, err
			}
		}

		for _, scg := range cg.SubConditionGroups {
			_, err := scg.evaluate(values)
			if err != nil {
				return nil, err
			}
		}
	}
	if cg.isRoot() {
		outMap := make(map[string]interface{})
		for _, vs := range cg.Vars {
			if vs.Type == VarTypOut {
				outMap[vs.Sym] = values[vs.Sym].val
			}
		}
		return outMap, nil
	}
	return nil, nil
}

type Value struct {
	sym     string
	val     interface{}
	typ     reflect.Type
	varType string
}

func getEffectiveValues(vars []*VarSpec, args map[string]interface{}) (map[string]*Value, error) {
	effectiveVals := make(map[string]*Value)
	stop := false
	var err error
	iterateVars(vars, &stop, func(vs *VarSpec) {
		if vs.Type == VarTypConst {
			effectiveVals[vs.Sym] = &Value{
				sym:     vs.Sym,
				val:     vs.Value,
				typ:     reflect.TypeOf(vs.Value),
				varType: vs.Type,
			}
		} else if vs.Type == VarTypIn {
			v, ok := args[vs.Sym]
			if !ok {
				err = errors.New(vs.Sym + " was not provided")
				stop = true
			}
			effectiveVals[vs.Sym] = &Value{
				sym:     vs.Sym,
				val:     v,
				typ:     reflect.TypeOf(v),
				varType: vs.Type,
			}

		} else {
			effectiveVals[vs.Sym] = &Value{
				sym:     vs.Sym,
				varType: vs.Type,
			}
		}
	})
	return effectiveVals, err
}
