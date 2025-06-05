lexer grammar VLexer;
// =====================
// PALABRAS RESERVADAS
// =====================
RW_MAIN:            'main';
RW_FN:              'fn';
RW_MUT:             'mut';
RW_STRUCT:          'struct';
RW_IF:              'if';
RW_ELSE:            'else';
RW_SWITCH:          'switch';
RW_CASE:            'case';
RW_DEFAULT:         'default';
RW_FOR:             'for';
RW_IN:              'in';
RW_BREAK:           'break';
RW_CONTINUE:        'continue';
RW_RETURN:          'return';
RW_TRUE:            'true';
RW_FALSE:           'false';
RW_NIL:             'nil';

RW_INT:             'int';
RW_FLOAT64:         'float64';
RW_STRING:          'string';
RW_BOOL:            'bool';
RW_RUNE:            'rune';

// =====================
// FUNCIONES EMBEBIDAS
// =====================
RW_PRINT:           'print';
RW_PRINTLN:         'println';
RW_ATOI:            'Atoi';
RW_PARSEFLOAT:      'parseFloat';
RW_TYPEOF:          'typeof';
RW_APPEND:          'append';
RW_LEN:             'len';
RW_JOIN:            'join';
RW_INDEXOF:         'indexOf';

// =====================
// OPERADORES
// =====================
OP_SUMA:                   '+';
OP_RESTA:                  '-';
OP_MULT:                   '*';
OP_DIV:                    '/';
OP_MOD:                    '%';

OP_ASSIGN:                 '=';
OP_ADD_ASSIGN:             '+=';
OP_SUB_ASSIGN:             '-=';

OP_IGUAL:                  '==';
OP_DIFERENTE:              '!=';
OP_MENORQ:                 '<';
OP_MAYORQ:                 '>';
OP_MENORIGUAL:             '<=';
OP_MAYORIGUAL:             '>=';

OP_AND:                    '&&';
OP_OR:                     '||';
OP_NOT:                    '!';
       
OP_DECLARATION:            ':=';

// =====================
// AGRUPACIÓN Y OTROS
// =====================

PUNTO:                  '.';
COMA:                   ',';
PUNTO_Y_COMA:           ';';
DOS_PUNTOS:             ':';

PAR_IZQ:                '(';
PAR_DER:                ')';
LLAVE_IZQ:              '{';
LLAVE_DER:              '}';
CORCHETE_IZQ:           '[';
CORCHETE_DER:           ']';

// =====================
// LITERALES
// =====================

INT_LITERAL:        [0-9]+;
FLOAT_LITERAL:      [0-9]+ '.' [0-9]+;
STRING_LITERAL:     '"' ( ~["\\] | '\\' . )* '"';
RUNE_LITERAL:       '\'' . '\'';

// =====================
// IDENTIFICADORES
// =====================
ID:                 [_a-zA-Z] [_a-zA-Z0-9]*;

// =====================
// COMENTARIOS Y ESPACIOS
// =====================

WS: [ \t\r\n]+ -> skip;
COMMENT: '//' .*? ('\n' | EOF) -> skip;
MULTILINE_COMMENT: '/*' .*? '*/' -> skip;
