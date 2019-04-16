package lure

// ExprList manages a list of Expr for related operations
// - exprs is the overall expressions added
// - unresolved tracks the exprs that need context for resolution
// - resolved tracks all the nonresolvable nodes (like literals)
type ExprList struct {
	exprs      []Expr
	unresolved []Expr
	resolved   map[string]IData
}

func newExprList() ExprList {
	return ExprList{
		exprs:      nil,
		unresolved: nil,
		resolved:   make(map[string]IData),
	}
}

func (me *ExprList) get(i int) Expr {
	if i < len(me.exprs) {
		return me.exprs[i]
	}
	return nil
}

func (me *ExprList) getResolvedSet() map[string]IData {
	return me.resolved
}

func (me *ExprList) getUnResolvedList() []Expr {
	return me.unresolved
}

func (me *ExprList) add(xp Expr) {
	me.exprs = append(me.exprs, xp)
	if xp.isResolvable() {
		me.unresolved = append(me.unresolved, xp)
	} else {
		data := xp.evaluate(nil)
		me.resolved[data.toString()] = data
	}
}

// isResolvable tells if any expr is resolvable
func (me *ExprList) isResolvable() bool {
	return len(me.unresolved) > 0
}
