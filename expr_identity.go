package lure

// ExprIdentity represents an identity node in the AST
type ExprIdentity struct {
	raw string
}

func (me ExprIdentity) evaluate(context map[string]IData) IData {
	if me.raw == "" {
		return boolDataFalse
	}
	if context == nil {
		return StringData{val: me.raw}
	}
	if val, ok := context[me.raw]; ok {
		return val
	}
	return StringData{val: me.raw}
}

func (me ExprIdentity) isResolvable() bool {
	return true
}
