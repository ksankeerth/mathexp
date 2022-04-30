package mathexp

import "encoding/json"

type MathExp struct {
	ExpWrapper *ConditionGroupSpec
}

func New(expJson []byte) (*MathExp, error) {
	var rootCondGrpSpec ConditionGroupSpec
	if err := json.Unmarshal(expJson, &rootCondGrpSpec); err != nil {
		return nil, err
	}
	return &MathExp{ExpWrapper: &rootCondGrpSpec}, nil
}
