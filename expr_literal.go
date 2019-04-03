package lure

type literalType int

// ExprLiteral represents an identity node in the AST
type ExprLiteral struct {
	data IData
}

func literalFromString(val string) ExprLiteral {
	return ExprLiteral{
		data: StringData{val: val},
	}
}

func literalFromBoolean(val bool) ExprLiteral {
	return ExprLiteral{
		data: BoolData{val: val},
	}
}

func literalFromDouble(val float64) ExprLiteral {
	return ExprLiteral{
		data: DoubleData{val: val},
	}
}

func literalFromInt(val int) ExprLiteral {
	return ExprLiteral{
		data: IntData{val: val},
	}
}

func (me ExprLiteral) evaluate(context map[string]IData) IData {
	if nil == me.data {
		return boolDataFalse
	}
	return me.data
}

func (me ExprLiteral) isResolvable() bool {
	return false
}
