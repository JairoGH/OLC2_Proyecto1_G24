parser grammar VGrammar;

options { tokenVocab=VLexer; }

programa : expr EOF ;

expr
    : expr op=(OP_MULT|OP_DIV) expr     # MulDiv
    | expr op=(OP_SUMA|OP_RESTA) expr  # AddSub
    | INT_LITERAL                        # Int
    | PAR_IZQ expr PAR_DER         # Parens
    ;