package lure

// DataType represents all acceptable types in the context
type DataType = int

// Enum for DataTypes
const (
	TypeBool DataType = iota + 1
	TypeInt
	TypeDouble
	TypeString
	TypeCustom
)

// Data holds data
type Data struct {
	boolVal        bool
	intVal         int
	doubleVal      float64
	stringVal      string
	genericVal     interface{}
	genericValType DataType
}

// Expr ...
type Expr interface {
}

// ExprList ...
type ExprList interface {
}

func (xpList ExprList) get(i int) Expr {
	return nil
}

func exprListAppend(list *ExprList, expr *Expr) *ExprList {
	return nil
}

func exprListOfExpr(expr *Expr) *ExprList {
	return nil
}

func exprBinOp(lhs *Expr, op int, rhs *Expr) *Expr {
	return nil
}

func exprIn(lhs *Expr, op int, list *ExprList) *Expr {
	return nil
}

func exprFunction0(fname string) *Expr {
	return nil
}
func exprFunction(fname string, list *ExprList) *Expr {
	return nil
}
func exprBetween(xp *Expr, xpMin *Expr, xpMax *Expr) *Expr {
	return nil
}
func exprUnaryOp(op int, xp *Expr) *Expr {
	return nil
}
func exprOfInt(val int) *Expr {
	return nil
}
func exprOfDouble(val float64) *Expr {
	return nil
}
func exprOfString(val string) *Expr {
	return nil
}
func exprOfIdentity(val string) *Expr {
	return nil
}

// Compile ...
func Compile(str string) *Expr {
	lexer := NewLexerOfString(str)
	parser := &LureParserImpl{}
	parser.Parse(*lexer)
	return parser.lval.exprList.get(0)
}

func main() {
}
