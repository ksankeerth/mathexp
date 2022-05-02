package mathexp

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

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
	return valid
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

				iterateVars(vars, &stop, func(value interface{}) {
					vs, ok := value.(*VarSpec)
					if ok {
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
					}
				})

			}
		}
	}
	return valid
}

func iterateVars(vars []*VarSpec, stop *bool, matcher func(value interface{})) {
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

func verifyBeforeEvalute(cg *ConditionGroupSpec, args map[string]interface{}) bool {
	requiredArgs := 0
	for _, vs := range allVars(cg) {
		if vs.Type == VarTypIn {
			if _, ok := args[vs.Sym]; ok {
				requiredArgs++
			}
		}
	}
	return requiredArgs == len(args)
}

// func (cg *ConditionGroupSpec) evaluate(args map[string]interface{}) (map[string]interface{}, error) {
// 	if ok := verifyBeforeEvalute(cg, args); !ok {
// 		return nil, ErrArgsMismatch
// 	}
// 	cg.traverse(new([]*VarSpec), func(cg *ConditionGroupSpec, vars []*VarSpec) bool {

// 	})
// }
