parser grammar vGrammar;

options {
  tokenVocab=vLexer;
}

// Punto de entrada
programa: funcionPrincipal EOF;

funcionPrincipal
    : RW_FN RW_MAIN PAR_IZQ PAR_DER bloque
    ;

bloque
    : LLAVE_IZQ sentencia* LLAVE_DER
    ;

sentencia
    : declaracion
    | asignacion
    | ifStatement
    | RW_PRINT PAR_IZQ expresion PAR_DER
    ;

declaracion
    : RW_MUT ID OP_DECLARATION expresion
    ;

asignacion
    : ID OP_ASSIGN expresion
    ;

ifStatement
    : RW_IF expresion bloque (RW_ELSE ifStatement)? (RW_ELSE bloque)?
    ;

expresion
    : expresion OP_SUMA expresion
    | expresion OP_RESTA expresion
    | expresion OP_MULT expresion
    | expresion OP_DIV expresion
    | expresion OP_MOD expresion
    | PAR_IZQ expresion PAR_DER
    | INT_LITERAL
    | FLOAT_LITERAL
    | STRING_LITERAL
    | RUNE_LITERAL
    | RW_TRUE
    | RW_FALSE
    | ID
    ;