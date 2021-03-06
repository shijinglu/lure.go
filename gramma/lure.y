/**
 * Gramma fro Lu's Rule Engine (LURE)
 * example:
 *  USER_TAGS IN ['admin', 'tester']
 *  CITY_ID IN (1, 2, 3)
 *  VERSION BETWEEN v1.2.3 AND v2.1.4
 *  VERSION >= v3.2.1
 *  DEVICE_OS == 'ios'
 *  RISK_SCORE > 0.85
 *  P_VALUE < 0.05
 */
%{
package lure

%}

%union {
	boolVal   bool
	intVal    int
	doubleVal float64
	strVal    string
	expr      Expr
	exprList  *ExprList
	savePoint *ExprList
}

// %parse-param {ExprList **rootExprList}

%left TK_AND_LOGIC
%left TK_OR_LOGIC
%left TK_BETWEEN
%left TK_EQ TK_NE TK_GT TK_LT TK_GE TK_LE
%left TK_IN TK_NOT
%nonassoc BINOPX


%token <intVal> TK_EQ TK_NE TK_GT TK_LT TK_GE TK_LE 
%token <intVal> TK_AND_KEYWORD TK_AND_LOGIC TK_OR_LOGIC TK_IN TK_NOT TK_NOTIN
%token <intVal> TK_INT_LITERAL
%token <doubleVal> TK_DOUBLE_LITERAL
%token <strVal> TK_STRING_LITERAL TK_MD5MOD TK_STRCMP TK_IDENTITY_LITERAL

%type <strVal> function_name
%type <expr> expr literal_value
%type <exprList> expr_list
%type <intVal> cmp_op


%%
expr_list
 : expr                 { $$ = exprListOfExpr($1); saveListOfExpr(&(Lurercvr.lval), $$); }
 | expr_list ',' expr   { $$ = exprListAppend($1, $3); }
 ;

expr
 : literal_value                                               { $$ = $1; }
 | expr cmp_op expr %prec BINOPX                               { $$ = exprBinOp($1, $2, $3); }
 | expr TK_IN '(' expr_list ')'                                { $$ = exprIn($1, TK_IN, $4); }
 | expr TK_NOT TK_IN '(' expr_list ')'                         { $$ = exprIn($1, TK_NOTIN, $5); }
 | function_name '(' ')'                                       { $$ = exprFunction0($1); }
 | function_name '(' expr_list ')'                             { $$ = exprFunction($1, $3); }
 | expr TK_BETWEEN expr TK_AND_KEYWORD expr %prec BINOPX       { $$ = exprBetween($1, $3, $5); }
 | expr TK_AND_LOGIC expr                                      { $$ = exprBinOp($1, $2, $3); }
 | expr TK_OR_LOGIC expr                                       { $$ = exprBinOp($1, $2, $3); }
 | TK_NOT expr                                                 { $$ = exprUnaryOp($1, $2); }
 | '(' expr ')'                                                { $$ = $2; }
 ;

cmp_op
 : TK_EQ
 | TK_NE
 | TK_GT
 | TK_LT
 | TK_GE
 | TK_LE
 ;

literal_value
 : TK_INT_LITERAL           { $$ = exprOfInt($1); }
 | TK_DOUBLE_LITERAL        { $$ = exprOfDouble($1); }
 | TK_STRING_LITERAL        { $$ = exprOfString($1); }
 | TK_IDENTITY_LITERAL      { $$ = exprOfIdentity($1); }
 ;

function_name
 : TK_MD5MOD
 | TK_STRCMP
 ;

%%

