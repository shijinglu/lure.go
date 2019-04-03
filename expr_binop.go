package lure

type binOpType int

const (
	binOpTypeEQ binOpType = iota
	binOpTypeNE
	binOpTypeGT
	binOpTypeLT
	binOpTypeGE
	binOpTypeLE
	binOpTypeAND
	binOpTypeOR
)

// ExprBinOp represents an identity node in the AST
type ExprBinOp struct {
	opTk  binOpType
	left  Expr
	right Expr
}

// evaluate returns:
// - FALSE if left or right or context is nil
// - a boolData that tells if right matches left
func (me ExprBinOp) evaluate(context map[string]IData) IData {
	if me.left == nil || me.right == nil || context == nil {
		return boolDataFalse
	}
	// return true for  "<left> like <right>" if
	leftData := me.left.evaluate(context)
	rightData := me.right.evaluate(context)
	cmp := leftData.compareTo(rightData)
	var ret bool
	switch me.opTk {
	case binOpTypeEQ:
		ret = cmp == 0
	case binOpTypeNE:
		ret = cmp != 0
	case binOpTypeGT:
		ret = cmp > 0
	case binOpTypeLT:
		ret = cmp < 0
	case binOpTypeGE:
		ret = cmp >= 0
	case binOpTypeLE:
		ret = cmp <= 0
	case binOpTypeAND:
		ret = leftData.toBoolean() && rightData.toBoolean()
	case binOpTypeOR:
		ret = leftData.toBoolean() || rightData.toBoolean()
	}

	if ret {
		return boolDataTrue
	}
	return boolDataFalse
}

func (me ExprBinOp) isResolvable() bool {
	return true
}
