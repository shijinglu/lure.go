package lure

func saveListOfExpr(lureSym *LureSymType, list *ExprList) {
	lureSym.savePoint = list
}

func exprListAppend(list *ExprList, expr Expr) *ExprList {
	if list == nil {
		return exprListOfExpr(expr)
	}
	list.add(expr)
	return list
}

func exprListOfExpr(expr Expr) *ExprList {
	list := newExprList()
	list.add(expr)
	return &list
}

// getBinOpType maps XX_TK to binOpTypeXX
func getBinOpType(opTk int) binOpType {
	switch opTk {
	case TK_AND_LOGIC:
		return binOpTypeAND
	case TK_OR_LOGIC:
		return binOpTypeOR
	case TK_EQ:
		return binOpTypeEQ
	case TK_NE:
		return binOpTypeNE
	case TK_GT:
		return binOpTypeGT
	case TK_LT:
		return binOpTypeLT
	case TK_GE:
		return binOpTypeGE
	case TK_LE:
		return binOpTypeLE
	default:
		return binOpType(opTk)
	}
}

func exprBinOp(lhs Expr, op int, rhs Expr) Expr {
	return ExprBinOp{
		opTk:  getBinOpType(op),
		left:  lhs,
		right: rhs,
	}
}

func exprIn(lhs Expr, op int, list *ExprList) Expr {
	x := ExprIn{
		left: lhs,
		list: list,
	}
	if op == TK_NOTIN {
		return ExprBinOp{
			opTk:  binOpTypeEQ,
			left:  x,
			right: literalFromBoolean(false),
		}
	}
	return x
}

func exprFunction0(fname string) Expr {
	return ExprFunction{
		left: ExprIdentity{raw: fname},
		list: nil,
	}
}

func exprFunction(fname string, list *ExprList) Expr {
	return ExprFunction{
		left: ExprIdentity{raw: fname},
		list: list,
	}
}

// exprBetween expresses xpMin <= xp <= xpMax
func exprBetween(xp Expr, xpMin Expr, xpMax Expr) Expr {
	return ExprBinOp{
		opTk: binOpTypeAND,
		left: ExprBinOp{
			opTk:  binOpTypeGE,
			left:  xp,
			right: xpMin,
		},
		right: ExprBinOp{
			opTk:  binOpTypeLE,
			left:  xp,
			right: xpMax,
		},
	}
}

func exprUnaryOp(op int, xp Expr) Expr {
	switch op {
	case TK_NOT:
		return ExprBinOp{
			opTk:  binOpTypeEQ,
			left:  xp,
			right: literalFromBoolean(false),
		}
	default:
		return nil
	}
}

func exprOfInt(val int) Expr {
	return literalFromInt(val)
}

func exprOfDouble(val float64) Expr {
	return literalFromDouble(val)
}

func exprOfString(val string) Expr {
	return literalFromString(val)
}

func exprOfIdentity(val string) Expr {
	return ExprIdentity{raw: val}
}
