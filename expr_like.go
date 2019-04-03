package lure

import "regexp"

// ExprLike represents an identity node in the AST
type ExprLike struct {
	left  Expr
	right Expr
}

// evaluate returns:
// - FALSE if left or right or context is nil
// - a boolData that tells if right matches left
func (me ExprLike) evaluate(context map[string]IData) IData {
	if me.left == nil || me.right == nil || context == nil {
		return boolDataFalse
	}
	// return true for  "<left> like <right>" if
	leftData := me.left.evaluate(context)
	rightData := me.right.evaluate(context)
	if matched, err := regexp.MatchString(rightData.toString(), leftData.toString()); err != nil {
		return BoolData{val: matched}
	}
	return boolDataFalse
}

func (me ExprLike) isResolvable() bool {
	return true
}
