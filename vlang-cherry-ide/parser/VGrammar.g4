parser grammar VGrammar;

options {
	tokenVocab = VLexer;
	
}
// Parser Rules
program: (stmt)* main_func? EOF?;


main_func: RW_FN RW_MAIN PAR_IZQ PAR_DER LLAVE_IZQ stmt* LLAVE_DER # FuncionMain;

stmt:
	decl_stmt
	| assign_stmt
	| transfer_stmt
	| if_stmt
	| switch_stmt
	| while_stmt
	| for_stmt
	| guard_stmt
	| func_call
	| vector_func
	| func_dcl
	| strct_dcl;

decl_stmt: 
    RW_MUT ID OP_ASSIGN expr                 # DeclararInferenciaMut
  | ID OP_ASSIGN expr                        # DeclararInferencia
  | RW_MUT ID type OP_DECLARACION expr       # DeclaraTipoValor
  | RW_MUT ID type                           # DeclararTipo
  | ID type OP_DECLARACION expr              # DeclararSinMutValor
  | ID type                                  # DeclararSinMutTipo
  | RW_MUT? ID OP_DECLARACION  type expr?      # DeclararVector
  ;

vector_expr:
	LLAVE_IZQ (expr (COMA expr)*)? COMA? LLAVE_DER # ListaItemsVector;

vector_item: id_pattern (CORCHETE_IZQ expr CORCHETE_DER)+ # VectorItem;

vector_prop: vector_item PUNTO id_pattern # PropVector;

vector_func: vector_item PUNTO func_call # FuncionVector;

repeating:
	(vector_type | matrix_type) PAR_IZQ ID DOS_PUNTOS expr COMA ID DOS_PUNTOS expr PAR_DER;

var_type: RW_MUT;

type: ID | RW_INT | RW_FLOAT64 | RW_STRING | RW_BOOL | vector_type | matrix_type;

vector_type: 
      CORCHETE_IZQ CORCHETE_DER type                                       # VectorSimple    // []int
    | CORCHETE_IZQ CORCHETE_DER CORCHETE_IZQ CORCHETE_DER type             # MatrizDoble     // [][]int
    ;

matrix_type: aux_matrix_type | CORCHETE_IZQ CORCHETE_IZQ ID CORCHETE_DER CORCHETE_DER;

aux_matrix_type: CORCHETE_IZQ matrix_type CORCHETE_DER;

assign_stmt: 
	 vector_item OP_DECLARACION expr    									# AssignVectorItem
	|id_pattern OP_DECLARACION expr											# AsignacionDirecta
	| id_pattern op = (OP_INCREMENTO | OP_DECREMENTO) expr				# AsignacionAritmetica
	| vector_item op = (OP_INCREMENTO | OP_DECREMENTO | OP_DECLARACION) expr	# AsignacionVector;

id_pattern: ID (PUNTO ID)* # ID_Patron;
  
literal:
	INT_LITERAL		    # IntLiteral 
	| FLOAT_LITERAL		# FloatLiteral
	| STRING_LITERAL	# StringLiteral 
	| BOOL_LITERAL		# BoolLiteral   
	| NIL_LITERAL		# NilLiteral; 

expr:
	PAR_IZQ expr PAR_DER									# ParentecisExp 
	| func_call											# FunctionCallExp 
	| id_pattern										# IdExp 
	| vector_item										# VectorItemExp 
	| vector_prop										# VectorPropExp 
	| vector_func										# VectorFuncExp 
	| literal											# LiteralExp 
	| vector_expr										# VectorExp 
	| repeating											# RepeatingExp
	| struct_vector										# StructExp
	| struct_init                                       # StructInitExp
	| struct_method_call                      			# StructMethodExp  
	| op = (OP_NOT | OP_RESTA) expr							 # ExpUnary 	
	| left = expr op = (OP_MULT | OP_DIV | OP_MOD) right = expr	# ExpBinario 
	| left = expr op = (OP_SUMA | OP_RESTA) right = expr		# ExpBinario 
	| left = expr op = (
		OP_MENORQ
		| OP_MENORIGUAL
		| OP_MAYORQ
		| OP_MAYORIGUAL
	) right = expr												 	# ExpBinario 
	| left = expr op = (OP_IGUAL | OP_DIFERENTE) right = expr	# ExpBinario 
	| left = expr op = OP_AND right = expr								# ExpBinario 
	| left = expr op = OP_OR right = expr								# ExpBinario;
  

if_stmt: if_chain (RW_ELSE if_chain)* else_stmt? # IFstmt;

if_chain: RW_IF expr LLAVE_IZQ stmt* LLAVE_DER # IFcadena;
else_stmt: RW_ELSE LLAVE_IZQ stmt* LLAVE_DER # ElseStmt; 
 
switch_stmt:
	RW_SWITCH expr LLAVE_IZQ switch_case* default_case? LLAVE_DER # SwitchStmt;

switch_case: RW_CASE expr DOS_PUNTOS stmt* # SwitchCase;

default_case: RW_DEFAULT DOS_PUNTOS stmt* # DefaultCase;

while_stmt: RW_FOR expr LLAVE_IZQ stmt* LLAVE_DER # WhileStmt;

for_stmt:
	RW_FOR ID RW_IN (expr | range) LLAVE_IZQ stmt* LLAVE_DER # ForStmt;

range: expr PUNTO PUNTO PUNTO expr # RangoNum;

guard_stmt:
	RW_GUARD expr RW_ELSE LLAVE_IZQ stmt* LLAVE_DER # GuardStmt;

transfer_stmt:
	RW_RETURN expr?	# ReturnStmt
	| RW_BREAK		# BreakStmt
	| RW_CONTINUE	# ContinueStmt;

func_call: id_pattern PAR_IZQ arg_list? PAR_DER # LlamarFuncion;

// external names -> num: value, num2: value2
arg_list: func_arg (COMA func_arg)* COMA? # ArgList;
func_arg: (ID DOS_PUNTOS)? (ANPERSAND)? (id_pattern | expr) # FuncionArg; // 

func_dcl:
    RW_FN ID PAR_IZQ param_list? PAR_DER type? LLAVE_IZQ stmt* LLAVE_DER                                  # FuncionDeclerada
    | RW_FN PAR_IZQ method_receiver PAR_DER ID PAR_IZQ param_list? PAR_DER type? LLAVE_IZQ stmt* LLAVE_DER    # MetodoStruct
    ;

method_receiver: ID OP_MULT ID # MethodReceiver;

param_list: func_param (COMA func_param)* COMA? # ParamList;
func_param: ID? ID  RW_INOUT? type # FuncParam;

// ESTRUCTURAS
strct_dcl: RW_STRUCT ID LLAVE_IZQ struct_prop* LLAVE_DER # DeclararStruct;

struct_prop:
    type ID (OP_DECLARACION expr)?              # StructAttr        
    | RW_MUTATING? func_dcl                     # StructFunc       
    ;
struct_vector: CORCHETE_IZQ ID CORCHETE_DER PAR_IZQ PAR_DER # StructVector;

struct_init: ID LLAVE_IZQ struct_init_list? LLAVE_DER # StructInit;

struct_init_list: struct_init_field (COMA struct_init_field)* COMA? # StructInitList;

struct_method_call: id_pattern PUNTO ID PAR_IZQ arg_list? PAR_DER # StructMethodCall;

struct_init_field: ID DOS_PUNTOS expr # StructInitField; //REVISAR SI YA NO SE USA