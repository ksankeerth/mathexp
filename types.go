package mathif

type Op string

type CondOp Op

const (
	OpEq   = CondOp("eq")
	OpNEq  = CondOp("neq")
	OpLt   = CondOp("lt")
	OpGt   = CondOp("gt")
	OpLtEq = CondOp("lteq")
	OpGtEq = CondOp("gteq")
)

type MathOp Op

const (
	OpAdd        = MathOp("add")
	OpMinus      = MathOp("sub")
	OpMultiply   = MathOp("mul")
	OpDivide     = MathOp("div")
	OpPercentage = MathOp("perc")
)

type CondConjOp Op

const (
	OpAnd = CondConjOp("and")
	OpOr  = CondConjOp("or")
	OpNot = CondConjOp("not")
)

type ConditionSpec struct {
	op Op
	v1 interface{}
	v2 interface{}
}

// Todo : return true if the
func (cs *ConditionSpec) Verify() bool {

}

func (cs *ConditionSpec) Eval() bool {
	// Todo check types
	// if float
	// if string
	// if int
	// bool
	// if ConditionSpec ev
}

// Todo : comment
// It will contain
type ConditionGroupSpec struct {
	cond     *ConditionSpec
	parent   *ConditionGroupSpec
	vars     []*VariableSpec
	exps     []*ExpresionSpec
	subConds []*ConditionGroupSpec
}

type ExpresionSpec struct {
	opCode MathOp
	v1     VariableSpec
	v2     VariableSpec
}

type VarType string

const (
	Const  = VarType("const")
	ExpOut = VarType("expout")
	Ref    = VarType("ref")
	In     = VarType("in")
)

type VariableSpec struct {
	VarType VarType
	name    string
	symbol  string
	value   interface{}
	// TODO : add unit information
}
