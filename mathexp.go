package mathexp

import (
	"encoding/json"
)

type MathExp struct {
	ExpWrapper *ConditionGroupSpec
}

func New(expJson []byte) (*MathExp, error) {
	var rootCondGrpSpec ConditionGroupSpec
	if err := json.Unmarshal(expJson, &rootCondGrpSpec); err != nil {
		return nil, err
	}
	if ok, err := rootCondGrpSpec.isValid(); !ok {
		return nil, err
	}
	return &MathExp{ExpWrapper: &rootCondGrpSpec}, nil
}

func (mexp *MathExp) Evaluate(args map[string]interface{}) (map[string]interface{}, error) {
	vars := allVars(mexp.ExpWrapper)
	next := verifyBeforeEvalute(vars, args)
	if !next {
		return nil, ErrArgsMismatch
	}

	vals, err := getEffectiveValues(vars, args)
	if err != nil {
		return nil, err
	}
	outs, err := mexp.ExpWrapper.evaluate(vals)
	return outs, err
}
