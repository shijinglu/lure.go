package lure

// ExprFunction represents an identity node in the AST
type ExprFunction struct {
	left Expr
	list *ExprList
}

// evaluate resolves function expression against the context
// Use left node string to search (case sensitive) for the functional in the extensions.
// If hit miss, return false, else use the functional to eval.
func (me ExprFunction) evaluate(context map[string]IData) IData {
	if me.left == nil || context == nil {
		return boolDataFalse
	}
	// return true for  "<left> like <right>" if
	var params []IData
	if me.list != nil {
		for _, xp := range me.list.exprs {
			if xp.isResolvable() {
				params = append(params, xp.evaluate(context))
			}
		}
	}
	leftData := me.left.evaluate(context)
	funcName := leftData.toString()
	if f, ok := getFunctions()[funcName]; ok {
		return f.derive(params)
	}
	return boolDataFalse
}

func (me ExprFunction) isResolvable() bool {
	return true
}
