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

// IData ...
type IData interface {
	extKey() string
	compareTo(rhs IData) int
	toBoolean() bool
	toInt() int
	toDouble() float64
	toString() string
}

// Expr ...
type Expr interface {
	isResolvable() bool

	// evaluate resolves the expr and return data, usually it
	// - returns FALSE if internal data is empty
	// - returns internal data if context is nil or look-up miss
	// - recursively resolve left node, right node and node list
	// - perform logic operation (e.g. comparison, regex match etc) with
	//   with respect to left, right and list nodes
	evaluate(context map[string]IData) IData
}

// ContextMap is a builder to construct context map
type ContextMap struct {
	m map[string]IData
}

// NewContextMap returns a ContextMap instance
func NewContextMap() ContextMap {
	return ContextMap{
		m: make(map[string]IData),
	}
}

// AddInt adds int type context
func (ctx ContextMap) AddInt(key string, val int) ContextMap {
	ctx.m[key] = IntData{val: val}
	return ctx
}

// AddDouble add float64 type context
func (ctx ContextMap) AddDouble(key string, val float64) ContextMap {
	ctx.m[key] = DoubleData{val: val}
	return ctx
}

// AddString add string type context
func (ctx ContextMap) AddString(key string, val string) ContextMap {
	ctx.m[key] = StringData{val: val}
	return ctx
}

// AddBoolean adds boolean type context
func (ctx ContextMap) AddBoolean(key string, val bool) ContextMap {
	ctx.m[key] = BoolData{val: val}
	return ctx
}

// AddCustom adds extended data type
func (ctx ContextMap) AddCustom(ctxKey, ctxValRaw, typeName string) ContextMap {
	if create, ok := getExtDataTypes()[typeName]; ok {
		ctx.m[ctxKey] = create(ctxValRaw, typeName)
	}
	return ctx
}

// Build builds the context map
func (ctx ContextMap) Build() map[string]IData {
	return ctx.m
}

// Compile ...
func Compile(str string) Expr {
	lexer := NewLexerOfString(str)
	// NOTE: here is a hack, `LureParserImpl` is the implementation details
	// that we are not suppose to use. We use it anyway because the generated
	// interface `func LureParse(Lurelex LureLexer) int` does not return any
	// thing useful.
	parser := &LureParserImpl{}
	parser.Parse(*lexer)
	return parser.lval.savePoint.get(0)
}

func main() {
}
