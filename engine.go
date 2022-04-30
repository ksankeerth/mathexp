package mathexp

type ConditionSpec struct {
	Op string
	V1 interface{}
	V2 interface{}
}

func (cond *ConditionSpec) isValid() bool {
	// TODO : check Op in the allowed list
	// TODO : check v1 and v2 are valid for Op
}

type VarSpec struct {
	Type  string
	Name  string
	Sym   string
	Value interface{}
	// TODO : add unit information
}


func (vs *VarSpec) isValid() bool {
	// TODO : check  Type is in the allowed list
	// TODO : check regex for name
	// TODO : check regex for sym
	// TODO : check  is the value suitable for the Type
}

type ExpresionSpec struct {
	Op  string
	V1  string
	V2  string
	Out string
}

func (es *ExpresionSpec) isValid(vars []*VarSpec) bool {
	// TODO : check Op is allowed list 
	// TODO: check v1, v2 and Out exists in vars
}

type ConditionGroupSpec struct {
	Cond               *ConditionSpec
	Vars               []*VarSpec
	Expressions        []*ExpresionSpec
	SubConditionGroups []*ConditionGroupSpec
	parent             *ConditionGroupSpec
}

type traveler func(cg *ConditionGroupSpec, vars []*VarSpec)

func (cg *ConditionGroupSpec) traverse(vars []*VarSpec, traveler traveler) {

	if vars == nil {
		vars = make([]*VarSpec, len(cg.Vars))
		copy(vars, cg.Vars)
	} else {
		vars = append(vars, cg.Vars...)
	}
	traveler(cg, vars)
	if cg.SubConditionGroups == nil {
		return
	}
	for _, subCg := range cg.SubConditionGroups {
		subCg.traverse(vars, traveler)
	}
}

func (cg *ConditionGroupSpec) isRoot() bool {
	return cg.parent == nil
}

func verify
