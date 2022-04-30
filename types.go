package mathexp

type ConditionSpec struct {
	Op string
	V1 interface{}
	V2 interface{}
}

type VarSpec struct {
	Type  string
	Name  string
	Sym   string
	Value interface{}
	// TODO : add unit information
}

type ExpresionSpec struct {
	Op  string
	V1  string
	V2  string
	Out string
}

type ConditionGroupSpec struct {
	Cond               *ConditionSpec
	Vars               []*VarSpec
	Expressions        []*ExpresionSpec
	SubConditionGroups []*ConditionGroupSpec
}
