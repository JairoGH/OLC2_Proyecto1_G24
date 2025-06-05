parser grammar VGrammar;

options {
    tokenVocab = VLexer;
}

// Punto de entrada del programa
programa: (stmt)* EOF;

// Sentencias válidas en el lenguaje
stmt
    : decl_stmt PUNTO_Y_COMA?                             #StmtDecl
    | assign_stmt PUNTO_Y_COMA?                           #StmtAssign
    | flow_stmt                                           #StmtFlow
    | func_call PUNTO_Y_COMA?                            #StmtFuncCall
    | if_stmt                                             #StmtIf
    | switch_stmt                                         #StmtSwitch
    | for_stmt                                            #StmtFor
    | func_decl                                           #StmtFuncDecl
    | struct_decl                                         #StmtStructDecl
    ;

// Declaración de variables

decl_stmt
    : RW_MUT ID OP_DECLARATION expr                      #InferDecl
    | RW_MUT ID RW_INT                                   #TypedDeclInt
    | RW_MUT ID RW_STRING                                #TypedDeclString
    | RW_MUT ID RW_FLOAT64                               #TypedDeclFloat
    | RW_MUT ID RW_BOOL                                  #TypedDeclBool
    | RW_MUT ID RW_RUNE                                  #TypedDeclRune
    ;

// Asignación de valores
assign_stmt
    : ID OP_ASSIGN expr                                   #AssignSimple
    | ID OP_ADD_ASSIGN expr                               #AssignAdd
    | ID OP_SUB_ASSIGN expr                               #AssignSub
    ;

// Expresiones
expr
    : literal                                             #ExprLiteral
    | ID                                                  #ExprId
    | expr OP_SUMA expr                                   #ExprAdd
    | expr OP_RESTA expr                                  #ExprSub
    | expr OP_MULT expr                                   #ExprMul
    | expr OP_DIV expr                                    #ExprDiv
    | expr OP_MOD expr                                    #ExprMod
    | expr OP_IGUAL expr                                  #ExprEqual
    | expr OP_DIFERENTE expr                              #ExprNotEqual
    | expr OP_MENORQ expr                                 #ExprLess
    | expr OP_MAYORQ expr                                 #ExprGreater
    | expr OP_MENORIGUAL expr                             #ExprLessEq
    | expr OP_MAYORIGUAL expr                             #ExprGreaterEq
    | expr OP_AND expr                                    #ExprAnd
    | expr OP_OR expr                                     #ExprOr
    | OP_NOT expr                                         #ExprNot
    | PAR_IZQ expr PAR_DER                                #ExprParen
    ;

literal
    : INT_LITERAL                                         #LitInt
    | FLOAT_LITERAL                                       #LitFloat
    | STRING_LITERAL                                      #LitString
    | RUNE_LITERAL                                        #LitRune
    | RW_TRUE                                             #LitTrue
    | RW_FALSE                                            #LitFalse
    | RW_NIL                                              #LitNil
    ;

// Sentencias de flujo
flow_stmt
    : RW_BREAK                                            #StmtBreak
    | RW_CONTINUE                                         #StmtContinue
    | RW_RETURN expr?                                     #StmtReturn
    ;

// Función embebida o llamada a función
func_call
    : ID PAR_IZQ (expr (COMA expr)*)? PAR_DER         # FunctionCall
    | RW_PRINT PAR_IZQ (expr (COMA expr)*)? PAR_DER   # PrintCall
    | RW_PRINTLN PAR_IZQ (expr (COMA expr)*)? PAR_DER # PrintLnCall
    ;

// If - else
if_stmt
    : RW_IF expr LLAVE_IZQ stmt* LLAVE_DER (RW_ELSE LLAVE_IZQ stmt* LLAVE_DER)?
    ;

// Switch - case
switch_stmt
    : RW_SWITCH expr LLAVE_IZQ switch_case* default_case? LLAVE_DER
    ;

switch_case
    : RW_CASE expr DOS_PUNTOS stmt*
    ;

default_case
    : RW_DEFAULT DOS_PUNTOS stmt*
    ;

// For loop
for_stmt
    : RW_FOR ID RW_IN expr LLAVE_IZQ stmt* LLAVE_DER
    ;

// Declaración de funciones
func_decl
    : RW_FN (ID | RW_MAIN) PAR_IZQ (param (COMA param)*)? PAR_DER (RW_INT | RW_STRING | RW_FLOAT64 | RW_BOOL | RW_RUNE)? LLAVE_IZQ stmt* LLAVE_DER
    ;

param
    : ID RW_INT
    | ID RW_STRING
    | ID RW_FLOAT64
    | ID RW_BOOL
    | ID RW_RUNE
    ;

// Declaración de structs
struct_decl
    : RW_STRUCT ID LLAVE_IZQ struct_field* LLAVE_DER
    ;

struct_field
    : RW_MUT ID RW_INT
    | RW_MUT ID RW_STRING
    | RW_MUT ID RW_FLOAT64
    | RW_MUT ID RW_BOOL
    | RW_MUT ID RW_RUNE
    ;
