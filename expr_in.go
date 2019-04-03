package lure

// ExprIn represents an identity node in the AST
type ExprIn struct {
	left Expr
	list *ExprList
}

// evaluate:
// - returns FALSE if left or context is nil
// - resolve all nodes in the list and return true if a match is found
// - if the already resolved sets contains the value, return TRUE
func (me ExprIn) evaluate(context map[string]IData) IData {
	if me.left == nil || context == nil || me.list == nil {
		return boolDataFalse
	}
	leftData := me.left.evaluate(context)
	// test if the resolved set contains the value
	m := me.list.getResolvedSet()
	if _, ok := m[leftData.toString()]; ok {
		return boolDataTrue
	}
	// loop over unresolved list, resolve and return TRUE if match found
	if me.list.isResolvable() {
		for _, xp := range me.list.getUnResolvedList() {
			if xp.evaluate(context).compareTo(leftData) == 0 {
				return boolDataTrue
			}
		}
	}
	return boolDataFalse
}

func (me ExprIn) isResolvable() bool {
	return true
}
