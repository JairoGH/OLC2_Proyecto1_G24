// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type VGrammar struct {
	*antlr.BaseParser
}

var VGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vgrammarParserInit() {
	staticData := &VGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "'mut'", "':='", "'+='", "'-='", "'='", "'int'", "",
		"'string'", "'bool'", "", "", "", "", "'nil'", "'+'", "'-'", "'*'",
		"'/'", "'%'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'&&'",
		"'||'", "'!'", "'main'", "'fn'", "'struct'", "'if'", "'else'", "'switch'",
		"'case'", "'default'", "'for'", "'while'", "'break'", "'continue'",
		"'return'", "'in'", "'inout'", "'guard'", "'mutating'", "", "'('", "')'",
		"'{'", "'}'", "'['", "']'", "'.'", "','", "';'", "':'", "'->'", "'?'",
		"'&'",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "COMENTARIO", "COMENTARIOMULT", "RW_MUT", "OP_ASSIGN", "OP_INCREMENTO",
		"OP_DECREMENTO", "OP_DECLARACION", "RW_INT", "RW_FLOAT64", "RW_STRING",
		"RW_BOOL", "INT_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL", "BOOL_LITERAL",
		"NIL_LITERAL", "OP_SUMA", "OP_RESTA", "OP_MULT", "OP_DIV", "OP_MOD",
		"OP_IGUAL", "OP_DIFERENTE", "OP_MENORQ", "OP_MAYORQ", "OP_MENORIGUAL",
		"OP_MAYORIGUAL", "OP_AND", "OP_OR", "OP_NOT", "RW_MAIN", "RW_FN", "RW_STRUCT",
		"RW_IF", "RW_ELSE", "RW_SWITCH", "RW_CASE", "RW_DEFAULT", "RW_FOR",
		"RW_WHILE", "RW_BREAK", "RW_CONTINUE", "RW_RETURN", "RW_IN", "RW_INOUT",
		"RW_GUARD", "RW_MUTATING", "ID", "PAR_IZQ", "PAR_DER", "LLAVE_IZQ",
		"LLAVE_DER", "CORCHETE_IZQ", "CORCHETE_DER", "PUNTO", "COMA", "PUNTO_Y_COMA",
		"DOS_PUNTOS", "ARROW", "INTERROGATION", "ANPERSAND",
	}
	staticData.RuleNames = []string{
		"program", "main_func", "stmt", "decl_stmt", "vector_expr", "vector_item",
		"vector_prop", "vector_func", "repeating", "var_type", "type", "vector_type",
		"matrix_type", "aux_matrix_type", "assign_stmt", "id_pattern", "literal",
		"expr", "if_stmt", "if_chain", "else_stmt", "switch_stmt", "switch_case",
		"default_case", "while_stmt", "for_stmt", "range", "guard_stmt", "transfer_stmt",
		"func_call", "arg_list", "func_arg", "func_dcl", "method_receiver",
		"param_list", "func_param", "strct_dcl", "struct_prop", "struct_vector",
		"struct_init", "struct_init_list", "struct_method_call", "struct_init_field",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 62, 587, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 5, 0, 88, 8, 0, 10, 0, 12, 0, 91, 9, 0, 1, 0, 3, 0, 94,
		8, 0, 1, 0, 3, 0, 97, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 105,
		8, 1, 10, 1, 12, 1, 108, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 124, 8, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 150,
		8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 156, 8, 3, 3, 3, 158, 8, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 5, 4, 164, 8, 4, 10, 4, 12, 4, 167, 9, 4, 3, 4, 169, 8,
		4, 1, 4, 3, 4, 172, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 4,
		5, 181, 8, 5, 11, 5, 12, 5, 182, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 3, 8, 195, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 3, 10, 216, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 3, 11, 226, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 3, 12, 234, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 3, 14, 256, 8, 14, 1, 15, 1, 15, 1, 15, 5, 15, 261,
		8, 15, 10, 15, 12, 15, 264, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3,
		16, 271, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3,
		17, 291, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5,
		17, 311, 8, 17, 10, 17, 12, 17, 314, 9, 17, 1, 18, 1, 18, 1, 18, 5, 18,
		319, 8, 18, 10, 18, 12, 18, 322, 9, 18, 1, 18, 3, 18, 325, 8, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 5, 19, 331, 8, 19, 10, 19, 12, 19, 334, 9, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 20, 5, 20, 341, 8, 20, 10, 20, 12, 20, 344,
		9, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 352, 8, 21, 10,
		21, 12, 21, 355, 9, 21, 1, 21, 3, 21, 358, 8, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 5, 22, 366, 8, 22, 10, 22, 12, 22, 369, 9, 22, 1,
		23, 1, 23, 1, 23, 5, 23, 374, 8, 23, 10, 23, 12, 23, 377, 9, 23, 1, 24,
		1, 24, 1, 24, 1, 24, 5, 24, 383, 8, 24, 10, 24, 12, 24, 386, 9, 24, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 395, 8, 25, 1, 25,
		1, 25, 5, 25, 399, 8, 25, 10, 25, 12, 25, 402, 9, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		5, 27, 417, 8, 27, 10, 27, 12, 27, 420, 9, 27, 1, 27, 1, 27, 1, 28, 1,
		28, 3, 28, 426, 8, 28, 1, 28, 1, 28, 3, 28, 430, 8, 28, 1, 29, 1, 29, 1,
		29, 3, 29, 435, 8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 5, 30, 442, 8,
		30, 10, 30, 12, 30, 445, 9, 30, 1, 30, 3, 30, 448, 8, 30, 1, 31, 1, 31,
		3, 31, 452, 8, 31, 1, 31, 3, 31, 455, 8, 31, 1, 31, 1, 31, 3, 31, 459,
		8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 465, 8, 32, 1, 32, 1, 32, 3,
		32, 469, 8, 32, 1, 32, 1, 32, 5, 32, 473, 8, 32, 10, 32, 12, 32, 476, 9,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 486,
		8, 32, 1, 32, 1, 32, 3, 32, 490, 8, 32, 1, 32, 1, 32, 5, 32, 494, 8, 32,
		10, 32, 12, 32, 497, 9, 32, 1, 32, 1, 32, 3, 32, 501, 8, 32, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 5, 34, 510, 8, 34, 10, 34, 12, 34,
		513, 9, 34, 1, 34, 3, 34, 516, 8, 34, 1, 35, 3, 35, 519, 8, 35, 1, 35,
		1, 35, 3, 35, 523, 8, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 5,
		36, 531, 8, 36, 10, 36, 12, 36, 534, 9, 36, 1, 36, 1, 36, 1, 37, 1, 37,
		1, 37, 1, 37, 3, 37, 542, 8, 37, 1, 37, 3, 37, 545, 8, 37, 1, 37, 3, 37,
		548, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1,
		39, 3, 39, 559, 8, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 5, 40, 566, 8,
		40, 10, 40, 12, 40, 569, 9, 40, 1, 40, 3, 40, 572, 8, 40, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 3, 41, 579, 8, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 42, 0, 1, 34, 43, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 0, 7, 1, 0, 6,
		7, 1, 0, 6, 8, 2, 0, 19, 19, 31, 31, 1, 0, 20, 22, 1, 0, 18, 19, 1, 0,
		25, 28, 1, 0, 23, 24, 645, 0, 89, 1, 0, 0, 0, 2, 98, 1, 0, 0, 0, 4, 123,
		1, 0, 0, 0, 6, 157, 1, 0, 0, 0, 8, 159, 1, 0, 0, 0, 10, 175, 1, 0, 0, 0,
		12, 184, 1, 0, 0, 0, 14, 188, 1, 0, 0, 0, 16, 194, 1, 0, 0, 0, 18, 206,
		1, 0, 0, 0, 20, 215, 1, 0, 0, 0, 22, 225, 1, 0, 0, 0, 24, 233, 1, 0, 0,
		0, 26, 235, 1, 0, 0, 0, 28, 255, 1, 0, 0, 0, 30, 257, 1, 0, 0, 0, 32, 270,
		1, 0, 0, 0, 34, 290, 1, 0, 0, 0, 36, 315, 1, 0, 0, 0, 38, 326, 1, 0, 0,
		0, 40, 337, 1, 0, 0, 0, 42, 347, 1, 0, 0, 0, 44, 361, 1, 0, 0, 0, 46, 370,
		1, 0, 0, 0, 48, 378, 1, 0, 0, 0, 50, 389, 1, 0, 0, 0, 52, 405, 1, 0, 0,
		0, 54, 411, 1, 0, 0, 0, 56, 429, 1, 0, 0, 0, 58, 431, 1, 0, 0, 0, 60, 438,
		1, 0, 0, 0, 62, 451, 1, 0, 0, 0, 64, 500, 1, 0, 0, 0, 66, 502, 1, 0, 0,
		0, 68, 506, 1, 0, 0, 0, 70, 518, 1, 0, 0, 0, 72, 526, 1, 0, 0, 0, 74, 547,
		1, 0, 0, 0, 76, 549, 1, 0, 0, 0, 78, 555, 1, 0, 0, 0, 80, 562, 1, 0, 0,
		0, 82, 573, 1, 0, 0, 0, 84, 582, 1, 0, 0, 0, 86, 88, 3, 4, 2, 0, 87, 86,
		1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0,
		90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 94, 3, 2, 1, 0, 93, 92, 1,
		0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 96, 1, 0, 0, 0, 95, 97, 5, 0, 0, 1, 96,
		95, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 1, 1, 0, 0, 0, 98, 99, 5, 33, 0,
		0, 99, 100, 5, 32, 0, 0, 100, 101, 5, 50, 0, 0, 101, 102, 5, 51, 0, 0,
		102, 106, 5, 52, 0, 0, 103, 105, 3, 4, 2, 0, 104, 103, 1, 0, 0, 0, 105,
		108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 109,
		1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109, 110, 5, 53, 0, 0, 110, 3, 1, 0,
		0, 0, 111, 124, 3, 6, 3, 0, 112, 124, 3, 28, 14, 0, 113, 124, 3, 56, 28,
		0, 114, 124, 3, 36, 18, 0, 115, 124, 3, 42, 21, 0, 116, 124, 3, 48, 24,
		0, 117, 124, 3, 50, 25, 0, 118, 124, 3, 54, 27, 0, 119, 124, 3, 58, 29,
		0, 120, 124, 3, 14, 7, 0, 121, 124, 3, 64, 32, 0, 122, 124, 3, 72, 36,
		0, 123, 111, 1, 0, 0, 0, 123, 112, 1, 0, 0, 0, 123, 113, 1, 0, 0, 0, 123,
		114, 1, 0, 0, 0, 123, 115, 1, 0, 0, 0, 123, 116, 1, 0, 0, 0, 123, 117,
		1, 0, 0, 0, 123, 118, 1, 0, 0, 0, 123, 119, 1, 0, 0, 0, 123, 120, 1, 0,
		0, 0, 123, 121, 1, 0, 0, 0, 123, 122, 1, 0, 0, 0, 124, 5, 1, 0, 0, 0, 125,
		126, 5, 4, 0, 0, 126, 127, 5, 49, 0, 0, 127, 128, 5, 5, 0, 0, 128, 158,
		3, 34, 17, 0, 129, 130, 5, 49, 0, 0, 130, 131, 5, 5, 0, 0, 131, 158, 3,
		34, 17, 0, 132, 133, 5, 4, 0, 0, 133, 134, 5, 49, 0, 0, 134, 135, 3, 20,
		10, 0, 135, 136, 5, 8, 0, 0, 136, 137, 3, 34, 17, 0, 137, 158, 1, 0, 0,
		0, 138, 139, 5, 4, 0, 0, 139, 140, 5, 49, 0, 0, 140, 158, 3, 20, 10, 0,
		141, 142, 5, 49, 0, 0, 142, 143, 3, 20, 10, 0, 143, 144, 5, 8, 0, 0, 144,
		145, 3, 34, 17, 0, 145, 158, 1, 0, 0, 0, 146, 147, 5, 49, 0, 0, 147, 158,
		3, 20, 10, 0, 148, 150, 5, 4, 0, 0, 149, 148, 1, 0, 0, 0, 149, 150, 1,
		0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 152, 5, 49, 0, 0, 152, 153, 5, 8, 0,
		0, 153, 155, 3, 20, 10, 0, 154, 156, 3, 34, 17, 0, 155, 154, 1, 0, 0, 0,
		155, 156, 1, 0, 0, 0, 156, 158, 1, 0, 0, 0, 157, 125, 1, 0, 0, 0, 157,
		129, 1, 0, 0, 0, 157, 132, 1, 0, 0, 0, 157, 138, 1, 0, 0, 0, 157, 141,
		1, 0, 0, 0, 157, 146, 1, 0, 0, 0, 157, 149, 1, 0, 0, 0, 158, 7, 1, 0, 0,
		0, 159, 168, 5, 52, 0, 0, 160, 165, 3, 34, 17, 0, 161, 162, 5, 57, 0, 0,
		162, 164, 3, 34, 17, 0, 163, 161, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0, 165,
		163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165,
		1, 0, 0, 0, 168, 160, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 171, 1, 0,
		0, 0, 170, 172, 5, 57, 0, 0, 171, 170, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0,
		172, 173, 1, 0, 0, 0, 173, 174, 5, 53, 0, 0, 174, 9, 1, 0, 0, 0, 175, 180,
		3, 30, 15, 0, 176, 177, 5, 54, 0, 0, 177, 178, 3, 34, 17, 0, 178, 179,
		5, 55, 0, 0, 179, 181, 1, 0, 0, 0, 180, 176, 1, 0, 0, 0, 181, 182, 1, 0,
		0, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 11, 1, 0, 0, 0,
		184, 185, 3, 10, 5, 0, 185, 186, 5, 56, 0, 0, 186, 187, 3, 30, 15, 0, 187,
		13, 1, 0, 0, 0, 188, 189, 3, 10, 5, 0, 189, 190, 5, 56, 0, 0, 190, 191,
		3, 58, 29, 0, 191, 15, 1, 0, 0, 0, 192, 195, 3, 22, 11, 0, 193, 195, 3,
		24, 12, 0, 194, 192, 1, 0, 0, 0, 194, 193, 1, 0, 0, 0, 195, 196, 1, 0,
		0, 0, 196, 197, 5, 50, 0, 0, 197, 198, 5, 49, 0, 0, 198, 199, 5, 59, 0,
		0, 199, 200, 3, 34, 17, 0, 200, 201, 5, 57, 0, 0, 201, 202, 5, 49, 0, 0,
		202, 203, 5, 59, 0, 0, 203, 204, 3, 34, 17, 0, 204, 205, 5, 51, 0, 0, 205,
		17, 1, 0, 0, 0, 206, 207, 5, 4, 0, 0, 207, 19, 1, 0, 0, 0, 208, 216, 5,
		49, 0, 0, 209, 216, 5, 9, 0, 0, 210, 216, 5, 10, 0, 0, 211, 216, 5, 11,
		0, 0, 212, 216, 5, 12, 0, 0, 213, 216, 3, 22, 11, 0, 214, 216, 3, 24, 12,
		0, 215, 208, 1, 0, 0, 0, 215, 209, 1, 0, 0, 0, 215, 210, 1, 0, 0, 0, 215,
		211, 1, 0, 0, 0, 215, 212, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 215, 214,
		1, 0, 0, 0, 216, 21, 1, 0, 0, 0, 217, 218, 5, 54, 0, 0, 218, 219, 5, 55,
		0, 0, 219, 226, 3, 20, 10, 0, 220, 221, 5, 54, 0, 0, 221, 222, 5, 55, 0,
		0, 222, 223, 5, 54, 0, 0, 223, 224, 5, 55, 0, 0, 224, 226, 3, 20, 10, 0,
		225, 217, 1, 0, 0, 0, 225, 220, 1, 0, 0, 0, 226, 23, 1, 0, 0, 0, 227, 234,
		3, 26, 13, 0, 228, 229, 5, 54, 0, 0, 229, 230, 5, 54, 0, 0, 230, 231, 5,
		49, 0, 0, 231, 232, 5, 55, 0, 0, 232, 234, 5, 55, 0, 0, 233, 227, 1, 0,
		0, 0, 233, 228, 1, 0, 0, 0, 234, 25, 1, 0, 0, 0, 235, 236, 5, 54, 0, 0,
		236, 237, 3, 24, 12, 0, 237, 238, 5, 55, 0, 0, 238, 27, 1, 0, 0, 0, 239,
		240, 3, 10, 5, 0, 240, 241, 5, 8, 0, 0, 241, 242, 3, 34, 17, 0, 242, 256,
		1, 0, 0, 0, 243, 244, 3, 30, 15, 0, 244, 245, 5, 8, 0, 0, 245, 246, 3,
		34, 17, 0, 246, 256, 1, 0, 0, 0, 247, 248, 3, 30, 15, 0, 248, 249, 7, 0,
		0, 0, 249, 250, 3, 34, 17, 0, 250, 256, 1, 0, 0, 0, 251, 252, 3, 10, 5,
		0, 252, 253, 7, 1, 0, 0, 253, 254, 3, 34, 17, 0, 254, 256, 1, 0, 0, 0,
		255, 239, 1, 0, 0, 0, 255, 243, 1, 0, 0, 0, 255, 247, 1, 0, 0, 0, 255,
		251, 1, 0, 0, 0, 256, 29, 1, 0, 0, 0, 257, 262, 5, 49, 0, 0, 258, 259,
		5, 56, 0, 0, 259, 261, 5, 49, 0, 0, 260, 258, 1, 0, 0, 0, 261, 264, 1,
		0, 0, 0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 31, 1, 0, 0,
		0, 264, 262, 1, 0, 0, 0, 265, 271, 5, 13, 0, 0, 266, 271, 5, 14, 0, 0,
		267, 271, 5, 15, 0, 0, 268, 271, 5, 16, 0, 0, 269, 271, 5, 17, 0, 0, 270,
		265, 1, 0, 0, 0, 270, 266, 1, 0, 0, 0, 270, 267, 1, 0, 0, 0, 270, 268,
		1, 0, 0, 0, 270, 269, 1, 0, 0, 0, 271, 33, 1, 0, 0, 0, 272, 273, 6, 17,
		-1, 0, 273, 274, 5, 50, 0, 0, 274, 275, 3, 34, 17, 0, 275, 276, 5, 51,
		0, 0, 276, 291, 1, 0, 0, 0, 277, 291, 3, 58, 29, 0, 278, 291, 3, 30, 15,
		0, 279, 291, 3, 10, 5, 0, 280, 291, 3, 12, 6, 0, 281, 291, 3, 14, 7, 0,
		282, 291, 3, 32, 16, 0, 283, 291, 3, 8, 4, 0, 284, 291, 3, 16, 8, 0, 285,
		291, 3, 76, 38, 0, 286, 291, 3, 78, 39, 0, 287, 291, 3, 82, 41, 0, 288,
		289, 7, 2, 0, 0, 289, 291, 3, 34, 17, 7, 290, 272, 1, 0, 0, 0, 290, 277,
		1, 0, 0, 0, 290, 278, 1, 0, 0, 0, 290, 279, 1, 0, 0, 0, 290, 280, 1, 0,
		0, 0, 290, 281, 1, 0, 0, 0, 290, 282, 1, 0, 0, 0, 290, 283, 1, 0, 0, 0,
		290, 284, 1, 0, 0, 0, 290, 285, 1, 0, 0, 0, 290, 286, 1, 0, 0, 0, 290,
		287, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 291, 312, 1, 0, 0, 0, 292, 293,
		10, 6, 0, 0, 293, 294, 7, 3, 0, 0, 294, 311, 3, 34, 17, 7, 295, 296, 10,
		5, 0, 0, 296, 297, 7, 4, 0, 0, 297, 311, 3, 34, 17, 6, 298, 299, 10, 4,
		0, 0, 299, 300, 7, 5, 0, 0, 300, 311, 3, 34, 17, 5, 301, 302, 10, 3, 0,
		0, 302, 303, 7, 6, 0, 0, 303, 311, 3, 34, 17, 4, 304, 305, 10, 2, 0, 0,
		305, 306, 5, 29, 0, 0, 306, 311, 3, 34, 17, 3, 307, 308, 10, 1, 0, 0, 308,
		309, 5, 30, 0, 0, 309, 311, 3, 34, 17, 2, 310, 292, 1, 0, 0, 0, 310, 295,
		1, 0, 0, 0, 310, 298, 1, 0, 0, 0, 310, 301, 1, 0, 0, 0, 310, 304, 1, 0,
		0, 0, 310, 307, 1, 0, 0, 0, 311, 314, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0,
		312, 313, 1, 0, 0, 0, 313, 35, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 315, 320,
		3, 38, 19, 0, 316, 317, 5, 36, 0, 0, 317, 319, 3, 38, 19, 0, 318, 316,
		1, 0, 0, 0, 319, 322, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 320, 321, 1, 0,
		0, 0, 321, 324, 1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 323, 325, 3, 40, 20,
		0, 324, 323, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 37, 1, 0, 0, 0, 326,
		327, 5, 35, 0, 0, 327, 328, 3, 34, 17, 0, 328, 332, 5, 52, 0, 0, 329, 331,
		3, 4, 2, 0, 330, 329, 1, 0, 0, 0, 331, 334, 1, 0, 0, 0, 332, 330, 1, 0,
		0, 0, 332, 333, 1, 0, 0, 0, 333, 335, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0,
		335, 336, 5, 53, 0, 0, 336, 39, 1, 0, 0, 0, 337, 338, 5, 36, 0, 0, 338,
		342, 5, 52, 0, 0, 339, 341, 3, 4, 2, 0, 340, 339, 1, 0, 0, 0, 341, 344,
		1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 345, 1, 0,
		0, 0, 344, 342, 1, 0, 0, 0, 345, 346, 5, 53, 0, 0, 346, 41, 1, 0, 0, 0,
		347, 348, 5, 37, 0, 0, 348, 349, 3, 34, 17, 0, 349, 353, 5, 52, 0, 0, 350,
		352, 3, 44, 22, 0, 351, 350, 1, 0, 0, 0, 352, 355, 1, 0, 0, 0, 353, 351,
		1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 357, 1, 0, 0, 0, 355, 353, 1, 0,
		0, 0, 356, 358, 3, 46, 23, 0, 357, 356, 1, 0, 0, 0, 357, 358, 1, 0, 0,
		0, 358, 359, 1, 0, 0, 0, 359, 360, 5, 53, 0, 0, 360, 43, 1, 0, 0, 0, 361,
		362, 5, 38, 0, 0, 362, 363, 3, 34, 17, 0, 363, 367, 5, 59, 0, 0, 364, 366,
		3, 4, 2, 0, 365, 364, 1, 0, 0, 0, 366, 369, 1, 0, 0, 0, 367, 365, 1, 0,
		0, 0, 367, 368, 1, 0, 0, 0, 368, 45, 1, 0, 0, 0, 369, 367, 1, 0, 0, 0,
		370, 371, 5, 39, 0, 0, 371, 375, 5, 59, 0, 0, 372, 374, 3, 4, 2, 0, 373,
		372, 1, 0, 0, 0, 374, 377, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 375, 376,
		1, 0, 0, 0, 376, 47, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0, 378, 379, 5, 40,
		0, 0, 379, 380, 3, 34, 17, 0, 380, 384, 5, 52, 0, 0, 381, 383, 3, 4, 2,
		0, 382, 381, 1, 0, 0, 0, 383, 386, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 384,
		385, 1, 0, 0, 0, 385, 387, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 387, 388,
		5, 53, 0, 0, 388, 49, 1, 0, 0, 0, 389, 390, 5, 40, 0, 0, 390, 391, 5, 49,
		0, 0, 391, 394, 5, 45, 0, 0, 392, 395, 3, 34, 17, 0, 393, 395, 3, 52, 26,
		0, 394, 392, 1, 0, 0, 0, 394, 393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396,
		400, 5, 52, 0, 0, 397, 399, 3, 4, 2, 0, 398, 397, 1, 0, 0, 0, 399, 402,
		1, 0, 0, 0, 400, 398, 1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 403, 1, 0,
		0, 0, 402, 400, 1, 0, 0, 0, 403, 404, 5, 53, 0, 0, 404, 51, 1, 0, 0, 0,
		405, 406, 3, 34, 17, 0, 406, 407, 5, 56, 0, 0, 407, 408, 5, 56, 0, 0, 408,
		409, 5, 56, 0, 0, 409, 410, 3, 34, 17, 0, 410, 53, 1, 0, 0, 0, 411, 412,
		5, 47, 0, 0, 412, 413, 3, 34, 17, 0, 413, 414, 5, 36, 0, 0, 414, 418, 5,
		52, 0, 0, 415, 417, 3, 4, 2, 0, 416, 415, 1, 0, 0, 0, 417, 420, 1, 0, 0,
		0, 418, 416, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 421, 1, 0, 0, 0, 420,
		418, 1, 0, 0, 0, 421, 422, 5, 53, 0, 0, 422, 55, 1, 0, 0, 0, 423, 425,
		5, 44, 0, 0, 424, 426, 3, 34, 17, 0, 425, 424, 1, 0, 0, 0, 425, 426, 1,
		0, 0, 0, 426, 430, 1, 0, 0, 0, 427, 430, 5, 42, 0, 0, 428, 430, 5, 43,
		0, 0, 429, 423, 1, 0, 0, 0, 429, 427, 1, 0, 0, 0, 429, 428, 1, 0, 0, 0,
		430, 57, 1, 0, 0, 0, 431, 432, 3, 30, 15, 0, 432, 434, 5, 50, 0, 0, 433,
		435, 3, 60, 30, 0, 434, 433, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 436,
		1, 0, 0, 0, 436, 437, 5, 51, 0, 0, 437, 59, 1, 0, 0, 0, 438, 443, 3, 62,
		31, 0, 439, 440, 5, 57, 0, 0, 440, 442, 3, 62, 31, 0, 441, 439, 1, 0, 0,
		0, 442, 445, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444,
		447, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 446, 448, 5, 57, 0, 0, 447, 446,
		1, 0, 0, 0, 447, 448, 1, 0, 0, 0, 448, 61, 1, 0, 0, 0, 449, 450, 5, 49,
		0, 0, 450, 452, 5, 59, 0, 0, 451, 449, 1, 0, 0, 0, 451, 452, 1, 0, 0, 0,
		452, 454, 1, 0, 0, 0, 453, 455, 5, 62, 0, 0, 454, 453, 1, 0, 0, 0, 454,
		455, 1, 0, 0, 0, 455, 458, 1, 0, 0, 0, 456, 459, 3, 30, 15, 0, 457, 459,
		3, 34, 17, 0, 458, 456, 1, 0, 0, 0, 458, 457, 1, 0, 0, 0, 459, 63, 1, 0,
		0, 0, 460, 461, 5, 33, 0, 0, 461, 462, 5, 49, 0, 0, 462, 464, 5, 50, 0,
		0, 463, 465, 3, 68, 34, 0, 464, 463, 1, 0, 0, 0, 464, 465, 1, 0, 0, 0,
		465, 466, 1, 0, 0, 0, 466, 468, 5, 51, 0, 0, 467, 469, 3, 20, 10, 0, 468,
		467, 1, 0, 0, 0, 468, 469, 1, 0, 0, 0, 469, 470, 1, 0, 0, 0, 470, 474,
		5, 52, 0, 0, 471, 473, 3, 4, 2, 0, 472, 471, 1, 0, 0, 0, 473, 476, 1, 0,
		0, 0, 474, 472, 1, 0, 0, 0, 474, 475, 1, 0, 0, 0, 475, 477, 1, 0, 0, 0,
		476, 474, 1, 0, 0, 0, 477, 501, 5, 53, 0, 0, 478, 479, 5, 33, 0, 0, 479,
		480, 5, 50, 0, 0, 480, 481, 3, 66, 33, 0, 481, 482, 5, 51, 0, 0, 482, 483,
		5, 49, 0, 0, 483, 485, 5, 50, 0, 0, 484, 486, 3, 68, 34, 0, 485, 484, 1,
		0, 0, 0, 485, 486, 1, 0, 0, 0, 486, 487, 1, 0, 0, 0, 487, 489, 5, 51, 0,
		0, 488, 490, 3, 20, 10, 0, 489, 488, 1, 0, 0, 0, 489, 490, 1, 0, 0, 0,
		490, 491, 1, 0, 0, 0, 491, 495, 5, 52, 0, 0, 492, 494, 3, 4, 2, 0, 493,
		492, 1, 0, 0, 0, 494, 497, 1, 0, 0, 0, 495, 493, 1, 0, 0, 0, 495, 496,
		1, 0, 0, 0, 496, 498, 1, 0, 0, 0, 497, 495, 1, 0, 0, 0, 498, 499, 5, 53,
		0, 0, 499, 501, 1, 0, 0, 0, 500, 460, 1, 0, 0, 0, 500, 478, 1, 0, 0, 0,
		501, 65, 1, 0, 0, 0, 502, 503, 5, 49, 0, 0, 503, 504, 5, 20, 0, 0, 504,
		505, 5, 49, 0, 0, 505, 67, 1, 0, 0, 0, 506, 511, 3, 70, 35, 0, 507, 508,
		5, 57, 0, 0, 508, 510, 3, 70, 35, 0, 509, 507, 1, 0, 0, 0, 510, 513, 1,
		0, 0, 0, 511, 509, 1, 0, 0, 0, 511, 512, 1, 0, 0, 0, 512, 515, 1, 0, 0,
		0, 513, 511, 1, 0, 0, 0, 514, 516, 5, 57, 0, 0, 515, 514, 1, 0, 0, 0, 515,
		516, 1, 0, 0, 0, 516, 69, 1, 0, 0, 0, 517, 519, 5, 49, 0, 0, 518, 517,
		1, 0, 0, 0, 518, 519, 1, 0, 0, 0, 519, 520, 1, 0, 0, 0, 520, 522, 5, 49,
		0, 0, 521, 523, 5, 46, 0, 0, 522, 521, 1, 0, 0, 0, 522, 523, 1, 0, 0, 0,
		523, 524, 1, 0, 0, 0, 524, 525, 3, 20, 10, 0, 525, 71, 1, 0, 0, 0, 526,
		527, 5, 34, 0, 0, 527, 528, 5, 49, 0, 0, 528, 532, 5, 52, 0, 0, 529, 531,
		3, 74, 37, 0, 530, 529, 1, 0, 0, 0, 531, 534, 1, 0, 0, 0, 532, 530, 1,
		0, 0, 0, 532, 533, 1, 0, 0, 0, 533, 535, 1, 0, 0, 0, 534, 532, 1, 0, 0,
		0, 535, 536, 5, 53, 0, 0, 536, 73, 1, 0, 0, 0, 537, 538, 3, 20, 10, 0,
		538, 541, 5, 49, 0, 0, 539, 540, 5, 8, 0, 0, 540, 542, 3, 34, 17, 0, 541,
		539, 1, 0, 0, 0, 541, 542, 1, 0, 0, 0, 542, 548, 1, 0, 0, 0, 543, 545,
		5, 48, 0, 0, 544, 543, 1, 0, 0, 0, 544, 545, 1, 0, 0, 0, 545, 546, 1, 0,
		0, 0, 546, 548, 3, 64, 32, 0, 547, 537, 1, 0, 0, 0, 547, 544, 1, 0, 0,
		0, 548, 75, 1, 0, 0, 0, 549, 550, 5, 54, 0, 0, 550, 551, 5, 49, 0, 0, 551,
		552, 5, 55, 0, 0, 552, 553, 5, 50, 0, 0, 553, 554, 5, 51, 0, 0, 554, 77,
		1, 0, 0, 0, 555, 556, 5, 49, 0, 0, 556, 558, 5, 52, 0, 0, 557, 559, 3,
		80, 40, 0, 558, 557, 1, 0, 0, 0, 558, 559, 1, 0, 0, 0, 559, 560, 1, 0,
		0, 0, 560, 561, 5, 53, 0, 0, 561, 79, 1, 0, 0, 0, 562, 567, 3, 84, 42,
		0, 563, 564, 5, 57, 0, 0, 564, 566, 3, 84, 42, 0, 565, 563, 1, 0, 0, 0,
		566, 569, 1, 0, 0, 0, 567, 565, 1, 0, 0, 0, 567, 568, 1, 0, 0, 0, 568,
		571, 1, 0, 0, 0, 569, 567, 1, 0, 0, 0, 570, 572, 5, 57, 0, 0, 571, 570,
		1, 0, 0, 0, 571, 572, 1, 0, 0, 0, 572, 81, 1, 0, 0, 0, 573, 574, 3, 30,
		15, 0, 574, 575, 5, 56, 0, 0, 575, 576, 5, 49, 0, 0, 576, 578, 5, 50, 0,
		0, 577, 579, 3, 60, 30, 0, 578, 577, 1, 0, 0, 0, 578, 579, 1, 0, 0, 0,
		579, 580, 1, 0, 0, 0, 580, 581, 5, 51, 0, 0, 581, 83, 1, 0, 0, 0, 582,
		583, 5, 49, 0, 0, 583, 584, 5, 59, 0, 0, 584, 585, 3, 34, 17, 0, 585, 85,
		1, 0, 0, 0, 61, 89, 93, 96, 106, 123, 149, 155, 157, 165, 168, 171, 182,
		194, 215, 225, 233, 255, 262, 270, 290, 310, 312, 320, 324, 332, 342, 353,
		357, 367, 375, 384, 394, 400, 418, 425, 429, 434, 443, 447, 451, 454, 458,
		464, 468, 474, 485, 489, 495, 500, 511, 515, 518, 522, 532, 541, 544, 547,
		558, 567, 571, 578,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// VGrammarInit initializes any static state used to implement VGrammar. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVGrammar(). You can call this function if you wish to initialize the static state ahead
// of time.
func VGrammarInit() {
	staticData := &VGrammarParserStaticData
	staticData.once.Do(vgrammarParserInit)
}

// NewVGrammar produces a new parser instance for the optional input antlr.TokenStream.
func NewVGrammar(input antlr.TokenStream) *VGrammar {
	VGrammarInit()
	this := new(VGrammar)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "VGrammar.g4"

	return this
}

// VGrammar tokens.
const (
	VGrammarEOF            = antlr.TokenEOF
	VGrammarWS             = 1
	VGrammarCOMENTARIO     = 2
	VGrammarCOMENTARIOMULT = 3
	VGrammarRW_MUT         = 4
	VGrammarOP_ASSIGN      = 5
	VGrammarOP_INCREMENTO  = 6
	VGrammarOP_DECREMENTO  = 7
	VGrammarOP_DECLARACION = 8
	VGrammarRW_INT         = 9
	VGrammarRW_FLOAT64     = 10
	VGrammarRW_STRING      = 11
	VGrammarRW_BOOL        = 12
	VGrammarINT_LITERAL    = 13
	VGrammarFLOAT_LITERAL  = 14
	VGrammarSTRING_LITERAL = 15
	VGrammarBOOL_LITERAL   = 16
	VGrammarNIL_LITERAL    = 17
	VGrammarOP_SUMA        = 18
	VGrammarOP_RESTA       = 19
	VGrammarOP_MULT        = 20
	VGrammarOP_DIV         = 21
	VGrammarOP_MOD         = 22
	VGrammarOP_IGUAL       = 23
	VGrammarOP_DIFERENTE   = 24
	VGrammarOP_MENORQ      = 25
	VGrammarOP_MAYORQ      = 26
	VGrammarOP_MENORIGUAL  = 27
	VGrammarOP_MAYORIGUAL  = 28
	VGrammarOP_AND         = 29
	VGrammarOP_OR          = 30
	VGrammarOP_NOT         = 31
	VGrammarRW_MAIN        = 32
	VGrammarRW_FN          = 33
	VGrammarRW_STRUCT      = 34
	VGrammarRW_IF          = 35
	VGrammarRW_ELSE        = 36
	VGrammarRW_SWITCH      = 37
	VGrammarRW_CASE        = 38
	VGrammarRW_DEFAULT     = 39
	VGrammarRW_FOR         = 40
	VGrammarRW_WHILE       = 41
	VGrammarRW_BREAK       = 42
	VGrammarRW_CONTINUE    = 43
	VGrammarRW_RETURN      = 44
	VGrammarRW_IN          = 45
	VGrammarRW_INOUT       = 46
	VGrammarRW_GUARD       = 47
	VGrammarRW_MUTATING    = 48
	VGrammarID             = 49
	VGrammarPAR_IZQ        = 50
	VGrammarPAR_DER        = 51
	VGrammarLLAVE_IZQ      = 52
	VGrammarLLAVE_DER      = 53
	VGrammarCORCHETE_IZQ   = 54
	VGrammarCORCHETE_DER   = 55
	VGrammarPUNTO          = 56
	VGrammarCOMA           = 57
	VGrammarPUNTO_Y_COMA   = 58
	VGrammarDOS_PUNTOS     = 59
	VGrammarARROW          = 60
	VGrammarINTERROGATION  = 61
	VGrammarANPERSAND      = 62
)

// VGrammar rules.
const (
	VGrammarRULE_program            = 0
	VGrammarRULE_main_func          = 1
	VGrammarRULE_stmt               = 2
	VGrammarRULE_decl_stmt          = 3
	VGrammarRULE_vector_expr        = 4
	VGrammarRULE_vector_item        = 5
	VGrammarRULE_vector_prop        = 6
	VGrammarRULE_vector_func        = 7
	VGrammarRULE_repeating          = 8
	VGrammarRULE_var_type           = 9
	VGrammarRULE_type               = 10
	VGrammarRULE_vector_type        = 11
	VGrammarRULE_matrix_type        = 12
	VGrammarRULE_aux_matrix_type    = 13
	VGrammarRULE_assign_stmt        = 14
	VGrammarRULE_id_pattern         = 15
	VGrammarRULE_literal            = 16
	VGrammarRULE_expr               = 17
	VGrammarRULE_if_stmt            = 18
	VGrammarRULE_if_chain           = 19
	VGrammarRULE_else_stmt          = 20
	VGrammarRULE_switch_stmt        = 21
	VGrammarRULE_switch_case        = 22
	VGrammarRULE_default_case       = 23
	VGrammarRULE_while_stmt         = 24
	VGrammarRULE_for_stmt           = 25
	VGrammarRULE_range              = 26
	VGrammarRULE_guard_stmt         = 27
	VGrammarRULE_transfer_stmt      = 28
	VGrammarRULE_func_call          = 29
	VGrammarRULE_arg_list           = 30
	VGrammarRULE_func_arg           = 31
	VGrammarRULE_func_dcl           = 32
	VGrammarRULE_method_receiver    = 33
	VGrammarRULE_param_list         = 34
	VGrammarRULE_func_param         = 35
	VGrammarRULE_strct_dcl          = 36
	VGrammarRULE_struct_prop        = 37
	VGrammarRULE_struct_vector      = 38
	VGrammarRULE_struct_init        = 39
	VGrammarRULE_struct_init_list   = 40
	VGrammarRULE_struct_method_call = 41
	VGrammarRULE_struct_init_field  = 42
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	Main_func() IMain_funcContext
	EOF() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ProgramContext) Main_func() IMain_funcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMain_funcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMain_funcContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(VGrammarEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VGrammarRULE_program)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(86)
				p.Stmt()
			}

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarRW_FN {
		{
			p.SetState(92)
			p.Main_func()
		}

	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(95)
			p.Match(VGrammarEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMain_funcContext is an interface to support dynamic dispatch.
type IMain_funcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMain_funcContext differentiates from other interfaces.
	IsMain_funcContext()
}

type Main_funcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMain_funcContext() *Main_funcContext {
	var p = new(Main_funcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_main_func
	return p
}

func InitEmptyMain_funcContext(p *Main_funcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_main_func
}

func (*Main_funcContext) IsMain_funcContext() {}

func NewMain_funcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Main_funcContext {
	var p = new(Main_funcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_main_func

	return p
}

func (s *Main_funcContext) GetParser() antlr.Parser { return s.parser }

func (s *Main_funcContext) CopyAll(ctx *Main_funcContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Main_funcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Main_funcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionMainContext struct {
	Main_funcContext
}

func NewFuncionMainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionMainContext {
	var p = new(FuncionMainContext)

	InitEmptyMain_funcContext(&p.Main_funcContext)
	p.parser = parser
	p.CopyAll(ctx.(*Main_funcContext))

	return p
}

func (s *FuncionMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionMainContext) RW_FN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FN, 0)
}

func (s *FuncionMainContext) RW_MAIN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MAIN, 0)
}

func (s *FuncionMainContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *FuncionMainContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *FuncionMainContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *FuncionMainContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *FuncionMainContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *FuncionMainContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *FuncionMainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFuncionMain(s)
	}
}

func (s *FuncionMainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFuncionMain(s)
	}
}

func (s *FuncionMainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFuncionMain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Main_func() (localctx IMain_funcContext) {
	localctx = NewMain_funcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VGrammarRULE_main_func)
	var _la int

	localctx = NewFuncionMainContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(VGrammarRW_FN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.Match(VGrammarRW_MAIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(VGrammarPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(VGrammarPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&735770847477776) != 0 {
		{
			p.SetState(103)
			p.Stmt()
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(109)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decl_stmt() IDecl_stmtContext
	Assign_stmt() IAssign_stmtContext
	Transfer_stmt() ITransfer_stmtContext
	If_stmt() IIf_stmtContext
	Switch_stmt() ISwitch_stmtContext
	While_stmt() IWhile_stmtContext
	For_stmt() IFor_stmtContext
	Guard_stmt() IGuard_stmtContext
	Func_call() IFunc_callContext
	Vector_func() IVector_funcContext
	Func_dcl() IFunc_dclContext
	Strct_dcl() IStrct_dclContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Decl_stmt() IDecl_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecl_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecl_stmtContext)
}

func (s *StmtContext) Assign_stmt() IAssign_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *StmtContext) Transfer_stmt() ITransfer_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransfer_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransfer_stmtContext)
}

func (s *StmtContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *StmtContext) Switch_stmt() ISwitch_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_stmtContext)
}

func (s *StmtContext) While_stmt() IWhile_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *StmtContext) For_stmt() IFor_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *StmtContext) Guard_stmt() IGuard_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuard_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuard_stmtContext)
}

func (s *StmtContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *StmtContext) Vector_func() IVector_funcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_funcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_funcContext)
}

func (s *StmtContext) Func_dcl() IFunc_dclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_dclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_dclContext)
}

func (s *StmtContext) Strct_dcl() IStrct_dclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrct_dclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrct_dclContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VGrammarRULE_stmt)
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Decl_stmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Assign_stmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(113)
			p.Transfer_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(114)
			p.If_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(115)
			p.Switch_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(116)
			p.While_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(117)
			p.For_stmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(118)
			p.Guard_stmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(119)
			p.Func_call()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(120)
			p.Vector_func()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(121)
			p.Func_dcl()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(122)
			p.Strct_dcl()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDecl_stmtContext is an interface to support dynamic dispatch.
type IDecl_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDecl_stmtContext differentiates from other interfaces.
	IsDecl_stmtContext()
}

type Decl_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecl_stmtContext() *Decl_stmtContext {
	var p = new(Decl_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_decl_stmt
	return p
}

func InitEmptyDecl_stmtContext(p *Decl_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_decl_stmt
}

func (*Decl_stmtContext) IsDecl_stmtContext() {}

func NewDecl_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Decl_stmtContext {
	var p = new(Decl_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_decl_stmt

	return p
}

func (s *Decl_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Decl_stmtContext) CopyAll(ctx *Decl_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Decl_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Decl_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclararVectorContext struct {
	Decl_stmtContext
}

func NewDeclararVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararVectorContext {
	var p = new(DeclararVectorContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *DeclararVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararVectorContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararVectorContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *DeclararVectorContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclararVectorContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *DeclararVectorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclararVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararVector(s)
	}
}

func (s *DeclararVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararVector(s)
	}
}

func (s *DeclararVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararVector(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclararInferenciaContext struct {
	Decl_stmtContext
}

func NewDeclararInferenciaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararInferenciaContext {
	var p = new(DeclararInferenciaContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *DeclararInferenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararInferenciaContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararInferenciaContext) OP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_ASSIGN, 0)
}

func (s *DeclararInferenciaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclararInferenciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararInferencia(s)
	}
}

func (s *DeclararInferenciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararInferencia(s)
	}
}

func (s *DeclararInferenciaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararInferencia(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclaraTipoValorContext struct {
	Decl_stmtContext
}

func NewDeclaraTipoValorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclaraTipoValorContext {
	var p = new(DeclaraTipoValorContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *DeclaraTipoValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaraTipoValorContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *DeclaraTipoValorContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclaraTipoValorContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclaraTipoValorContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *DeclaraTipoValorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclaraTipoValorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclaraTipoValor(s)
	}
}

func (s *DeclaraTipoValorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclaraTipoValor(s)
	}
}

func (s *DeclaraTipoValorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclaraTipoValor(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclararSinMutValorContext struct {
	Decl_stmtContext
}

func NewDeclararSinMutValorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararSinMutValorContext {
	var p = new(DeclararSinMutValorContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *DeclararSinMutValorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararSinMutValorContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararSinMutValorContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclararSinMutValorContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *DeclararSinMutValorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclararSinMutValorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararSinMutValor(s)
	}
}

func (s *DeclararSinMutValorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararSinMutValor(s)
	}
}

func (s *DeclararSinMutValorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararSinMutValor(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclararInferenciaMutContext struct {
	Decl_stmtContext
}

func NewDeclararInferenciaMutContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararInferenciaMutContext {
	var p = new(DeclararInferenciaMutContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *DeclararInferenciaMutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararInferenciaMutContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *DeclararInferenciaMutContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararInferenciaMutContext) OP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_ASSIGN, 0)
}

func (s *DeclararInferenciaMutContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclararInferenciaMutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararInferenciaMut(s)
	}
}

func (s *DeclararInferenciaMutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararInferenciaMut(s)
	}
}

func (s *DeclararInferenciaMutContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararInferenciaMut(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclararTipoContext struct {
	Decl_stmtContext
}

func NewDeclararTipoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararTipoContext {
	var p = new(DeclararTipoContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *DeclararTipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararTipoContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *DeclararTipoContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararTipoContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclararTipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararTipo(s)
	}
}

func (s *DeclararTipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararTipo(s)
	}
}

func (s *DeclararTipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclararSinMutTipoContext struct {
	Decl_stmtContext
}

func NewDeclararSinMutTipoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararSinMutTipoContext {
	var p = new(DeclararSinMutTipoContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *DeclararSinMutTipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararSinMutTipoContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararSinMutTipoContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclararSinMutTipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararSinMutTipo(s)
	}
}

func (s *DeclararSinMutTipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararSinMutTipo(s)
	}
}

func (s *DeclararSinMutTipoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararSinMutTipo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Decl_stmt() (localctx IDecl_stmtContext) {
	localctx = NewDecl_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VGrammarRULE_decl_stmt)
	var _la int

	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclararInferenciaMutContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Match(VGrammarOP_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.expr(0)
		}

	case 2:
		localctx = NewDeclararInferenciaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Match(VGrammarOP_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.expr(0)
		}

	case 3:
		localctx = NewDeclaraTipoValorContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.Type_()
		}
		{
			p.SetState(135)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.expr(0)
		}

	case 4:
		localctx = NewDeclararTipoContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(138)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Type_()
		}

	case 5:
		localctx = NewDeclararSinMutValorContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(141)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Type_()
		}
		{
			p.SetState(143)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.expr(0)
		}

	case 6:
		localctx = NewDeclararSinMutTipoContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(146)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Type_()
		}

	case 7:
		localctx = NewDeclararVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VGrammarRW_MUT {
			{
				p.SetState(148)
				p.Match(VGrammarRW_MUT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(151)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Type_()
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(154)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVector_exprContext is an interface to support dynamic dispatch.
type IVector_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVector_exprContext differentiates from other interfaces.
	IsVector_exprContext()
}

type Vector_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVector_exprContext() *Vector_exprContext {
	var p = new(Vector_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_vector_expr
	return p
}

func InitEmptyVector_exprContext(p *Vector_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_vector_expr
}

func (*Vector_exprContext) IsVector_exprContext() {}

func NewVector_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vector_exprContext {
	var p = new(Vector_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_vector_expr

	return p
}

func (s *Vector_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Vector_exprContext) CopyAll(ctx *Vector_exprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Vector_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vector_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListaItemsVectorContext struct {
	Vector_exprContext
}

func NewListaItemsVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListaItemsVectorContext {
	var p = new(ListaItemsVectorContext)

	InitEmptyVector_exprContext(&p.Vector_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Vector_exprContext))

	return p
}

func (s *ListaItemsVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListaItemsVectorContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *ListaItemsVectorContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *ListaItemsVectorContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ListaItemsVectorContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListaItemsVectorContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *ListaItemsVectorContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *ListaItemsVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterListaItemsVector(s)
	}
}

func (s *ListaItemsVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitListaItemsVector(s)
	}
}

func (s *ListaItemsVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitListaItemsVector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Vector_expr() (localctx IVector_exprContext) {
	localctx = NewVector_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VGrammarRULE_vector_expr)
	var _la int

	var _alt int

	localctx = NewListaItemsVectorContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&24206850145378304) != 0 {
		{
			p.SetState(160)
			p.expr(0)
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(161)
					p.Match(VGrammarCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(162)
					p.expr(0)
				}

			}
			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarCOMA {
		{
			p.SetState(170)
			p.Match(VGrammarCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(173)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVector_itemContext is an interface to support dynamic dispatch.
type IVector_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVector_itemContext differentiates from other interfaces.
	IsVector_itemContext()
}

type Vector_itemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVector_itemContext() *Vector_itemContext {
	var p = new(Vector_itemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_vector_item
	return p
}

func InitEmptyVector_itemContext(p *Vector_itemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_vector_item
}

func (*Vector_itemContext) IsVector_itemContext() {}

func NewVector_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vector_itemContext {
	var p = new(Vector_itemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_vector_item

	return p
}

func (s *Vector_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Vector_itemContext) CopyAll(ctx *Vector_itemContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Vector_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vector_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorItemContext struct {
	Vector_itemContext
}

func NewVectorItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorItemContext {
	var p = new(VectorItemContext)

	InitEmptyVector_itemContext(&p.Vector_itemContext)
	p.parser = parser
	p.CopyAll(ctx.(*Vector_itemContext))

	return p
}

func (s *VectorItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorItemContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *VectorItemContext) AllCORCHETE_IZQ() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCORCHETE_IZQ)
}

func (s *VectorItemContext) CORCHETE_IZQ(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_IZQ, i)
}

func (s *VectorItemContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *VectorItemContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VectorItemContext) AllCORCHETE_DER() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCORCHETE_DER)
}

func (s *VectorItemContext) CORCHETE_DER(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_DER, i)
}

func (s *VectorItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterVectorItem(s)
	}
}

func (s *VectorItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitVectorItem(s)
	}
}

func (s *VectorItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitVectorItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Vector_item() (localctx IVector_itemContext) {
	localctx = NewVector_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VGrammarRULE_vector_item)
	var _alt int

	localctx = NewVectorItemContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Id_pattern()
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(176)
				p.Match(VGrammarCORCHETE_IZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(177)
				p.expr(0)
			}
			{
				p.SetState(178)
				p.Match(VGrammarCORCHETE_DER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVector_propContext is an interface to support dynamic dispatch.
type IVector_propContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVector_propContext differentiates from other interfaces.
	IsVector_propContext()
}

type Vector_propContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVector_propContext() *Vector_propContext {
	var p = new(Vector_propContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_vector_prop
	return p
}

func InitEmptyVector_propContext(p *Vector_propContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_vector_prop
}

func (*Vector_propContext) IsVector_propContext() {}

func NewVector_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vector_propContext {
	var p = new(Vector_propContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_vector_prop

	return p
}

func (s *Vector_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Vector_propContext) CopyAll(ctx *Vector_propContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Vector_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vector_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PropVectorContext struct {
	Vector_propContext
}

func NewPropVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropVectorContext {
	var p = new(PropVectorContext)

	InitEmptyVector_propContext(&p.Vector_propContext)
	p.parser = parser
	p.CopyAll(ctx.(*Vector_propContext))

	return p
}

func (s *PropVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropVectorContext) Vector_item() IVector_itemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_itemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_itemContext)
}

func (s *PropVectorContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO, 0)
}

func (s *PropVectorContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *PropVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterPropVector(s)
	}
}

func (s *PropVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitPropVector(s)
	}
}

func (s *PropVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitPropVector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Vector_prop() (localctx IVector_propContext) {
	localctx = NewVector_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VGrammarRULE_vector_prop)
	localctx = NewPropVectorContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Vector_item()
	}
	{
		p.SetState(185)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Id_pattern()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVector_funcContext is an interface to support dynamic dispatch.
type IVector_funcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVector_funcContext differentiates from other interfaces.
	IsVector_funcContext()
}

type Vector_funcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVector_funcContext() *Vector_funcContext {
	var p = new(Vector_funcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_vector_func
	return p
}

func InitEmptyVector_funcContext(p *Vector_funcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_vector_func
}

func (*Vector_funcContext) IsVector_funcContext() {}

func NewVector_funcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vector_funcContext {
	var p = new(Vector_funcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_vector_func

	return p
}

func (s *Vector_funcContext) GetParser() antlr.Parser { return s.parser }

func (s *Vector_funcContext) CopyAll(ctx *Vector_funcContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Vector_funcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vector_funcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionVectorContext struct {
	Vector_funcContext
}

func NewFuncionVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionVectorContext {
	var p = new(FuncionVectorContext)

	InitEmptyVector_funcContext(&p.Vector_funcContext)
	p.parser = parser
	p.CopyAll(ctx.(*Vector_funcContext))

	return p
}

func (s *FuncionVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionVectorContext) Vector_item() IVector_itemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_itemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_itemContext)
}

func (s *FuncionVectorContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO, 0)
}

func (s *FuncionVectorContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *FuncionVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFuncionVector(s)
	}
}

func (s *FuncionVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFuncionVector(s)
	}
}

func (s *FuncionVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFuncionVector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Vector_func() (localctx IVector_funcContext) {
	localctx = NewVector_funcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VGrammarRULE_vector_func)
	localctx = NewFuncionVectorContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Vector_item()
	}
	{
		p.SetState(189)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
		p.Func_call()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepeatingContext is an interface to support dynamic dispatch.
type IRepeatingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PAR_IZQ() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllDOS_PUNTOS() []antlr.TerminalNode
	DOS_PUNTOS(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	COMA() antlr.TerminalNode
	PAR_DER() antlr.TerminalNode
	Vector_type() IVector_typeContext
	Matrix_type() IMatrix_typeContext

	// IsRepeatingContext differentiates from other interfaces.
	IsRepeatingContext()
}

type RepeatingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeatingContext() *RepeatingContext {
	var p = new(RepeatingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_repeating
	return p
}

func InitEmptyRepeatingContext(p *RepeatingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_repeating
}

func (*RepeatingContext) IsRepeatingContext() {}

func NewRepeatingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatingContext {
	var p = new(RepeatingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_repeating

	return p
}

func (s *RepeatingContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatingContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *RepeatingContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VGrammarID)
}

func (s *RepeatingContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarID, i)
}

func (s *RepeatingContext) AllDOS_PUNTOS() []antlr.TerminalNode {
	return s.GetTokens(VGrammarDOS_PUNTOS)
}

func (s *RepeatingContext) DOS_PUNTOS(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, i)
}

func (s *RepeatingContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RepeatingContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RepeatingContext) COMA() antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, 0)
}

func (s *RepeatingContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *RepeatingContext) Vector_type() IVector_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_typeContext)
}

func (s *RepeatingContext) Matrix_type() IMatrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_typeContext)
}

func (s *RepeatingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepeatingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterRepeating(s)
	}
}

func (s *RepeatingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitRepeating(s)
	}
}

func (s *RepeatingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitRepeating(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Repeating() (localctx IRepeatingContext) {
	localctx = NewRepeatingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VGrammarRULE_repeating)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(192)
			p.Vector_type()
		}

	case 2:
		{
			p.SetState(193)
			p.Matrix_type()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(196)
		p.Match(VGrammarPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.Match(VGrammarDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.expr(0)
	}
	{
		p.SetState(200)
		p.Match(VGrammarCOMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Match(VGrammarDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.expr(0)
	}
	{
		p.SetState(204)
		p.Match(VGrammarPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_MUT() antlr.TerminalNode

	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterVar_type(s)
	}
}

func (s *Var_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitVar_type(s)
	}
}

func (s *Var_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitVar_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VGrammarRULE_var_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(VGrammarRW_MUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	RW_INT() antlr.TerminalNode
	RW_FLOAT64() antlr.TerminalNode
	RW_STRING() antlr.TerminalNode
	RW_BOOL() antlr.TerminalNode
	Vector_type() IVector_typeContext
	Matrix_type() IMatrix_typeContext

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *TypeContext) RW_INT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_INT, 0)
}

func (s *TypeContext) RW_FLOAT64() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FLOAT64, 0)
}

func (s *TypeContext) RW_STRING() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_STRING, 0)
}

func (s *TypeContext) RW_BOOL() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_BOOL, 0)
}

func (s *TypeContext) Vector_type() IVector_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_typeContext)
}

func (s *TypeContext) Matrix_type() IMatrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_typeContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VGrammarRULE_type)
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(208)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(209)
			p.Match(VGrammarRW_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(210)
			p.Match(VGrammarRW_FLOAT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(211)
			p.Match(VGrammarRW_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(212)
			p.Match(VGrammarRW_BOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(213)
			p.Vector_type()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(214)
			p.Matrix_type()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVector_typeContext is an interface to support dynamic dispatch.
type IVector_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVector_typeContext differentiates from other interfaces.
	IsVector_typeContext()
}

type Vector_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVector_typeContext() *Vector_typeContext {
	var p = new(Vector_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_vector_type
	return p
}

func InitEmptyVector_typeContext(p *Vector_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_vector_type
}

func (*Vector_typeContext) IsVector_typeContext() {}

func NewVector_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vector_typeContext {
	var p = new(Vector_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_vector_type

	return p
}

func (s *Vector_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Vector_typeContext) CopyAll(ctx *Vector_typeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Vector_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vector_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VectorSimpleContext struct {
	Vector_typeContext
}

func NewVectorSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorSimpleContext {
	var p = new(VectorSimpleContext)

	InitEmptyVector_typeContext(&p.Vector_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Vector_typeContext))

	return p
}

func (s *VectorSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorSimpleContext) CORCHETE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_IZQ, 0)
}

func (s *VectorSimpleContext) CORCHETE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_DER, 0)
}

func (s *VectorSimpleContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VectorSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterVectorSimple(s)
	}
}

func (s *VectorSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitVectorSimple(s)
	}
}

func (s *VectorSimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitVectorSimple(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatrizDobleContext struct {
	Vector_typeContext
}

func NewMatrizDobleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatrizDobleContext {
	var p = new(MatrizDobleContext)

	InitEmptyVector_typeContext(&p.Vector_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Vector_typeContext))

	return p
}

func (s *MatrizDobleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatrizDobleContext) AllCORCHETE_IZQ() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCORCHETE_IZQ)
}

func (s *MatrizDobleContext) CORCHETE_IZQ(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_IZQ, i)
}

func (s *MatrizDobleContext) AllCORCHETE_DER() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCORCHETE_DER)
}

func (s *MatrizDobleContext) CORCHETE_DER(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_DER, i)
}

func (s *MatrizDobleContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *MatrizDobleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterMatrizDoble(s)
	}
}

func (s *MatrizDobleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitMatrizDoble(s)
	}
}

func (s *MatrizDobleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitMatrizDoble(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Vector_type() (localctx IVector_typeContext) {
	localctx = NewVector_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VGrammarRULE_vector_type)
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVectorSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.Match(VGrammarCORCHETE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.Match(VGrammarCORCHETE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Type_()
		}

	case 2:
		localctx = NewMatrizDobleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.Match(VGrammarCORCHETE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)
			p.Match(VGrammarCORCHETE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.Match(VGrammarCORCHETE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Match(VGrammarCORCHETE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Type_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatrix_typeContext is an interface to support dynamic dispatch.
type IMatrix_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Aux_matrix_type() IAux_matrix_typeContext
	AllCORCHETE_IZQ() []antlr.TerminalNode
	CORCHETE_IZQ(i int) antlr.TerminalNode
	ID() antlr.TerminalNode
	AllCORCHETE_DER() []antlr.TerminalNode
	CORCHETE_DER(i int) antlr.TerminalNode

	// IsMatrix_typeContext differentiates from other interfaces.
	IsMatrix_typeContext()
}

type Matrix_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatrix_typeContext() *Matrix_typeContext {
	var p = new(Matrix_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_matrix_type
	return p
}

func InitEmptyMatrix_typeContext(p *Matrix_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_matrix_type
}

func (*Matrix_typeContext) IsMatrix_typeContext() {}

func NewMatrix_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Matrix_typeContext {
	var p = new(Matrix_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_matrix_type

	return p
}

func (s *Matrix_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Matrix_typeContext) Aux_matrix_type() IAux_matrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAux_matrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAux_matrix_typeContext)
}

func (s *Matrix_typeContext) AllCORCHETE_IZQ() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCORCHETE_IZQ)
}

func (s *Matrix_typeContext) CORCHETE_IZQ(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_IZQ, i)
}

func (s *Matrix_typeContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *Matrix_typeContext) AllCORCHETE_DER() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCORCHETE_DER)
}

func (s *Matrix_typeContext) CORCHETE_DER(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_DER, i)
}

func (s *Matrix_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Matrix_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Matrix_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterMatrix_type(s)
	}
}

func (s *Matrix_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitMatrix_type(s)
	}
}

func (s *Matrix_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitMatrix_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Matrix_type() (localctx IMatrix_typeContext) {
	localctx = NewMatrix_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VGrammarRULE_matrix_type)
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.Aux_matrix_type()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.Match(VGrammarCORCHETE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Match(VGrammarCORCHETE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Match(VGrammarCORCHETE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.Match(VGrammarCORCHETE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAux_matrix_typeContext is an interface to support dynamic dispatch.
type IAux_matrix_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CORCHETE_IZQ() antlr.TerminalNode
	Matrix_type() IMatrix_typeContext
	CORCHETE_DER() antlr.TerminalNode

	// IsAux_matrix_typeContext differentiates from other interfaces.
	IsAux_matrix_typeContext()
}

type Aux_matrix_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAux_matrix_typeContext() *Aux_matrix_typeContext {
	var p = new(Aux_matrix_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_aux_matrix_type
	return p
}

func InitEmptyAux_matrix_typeContext(p *Aux_matrix_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_aux_matrix_type
}

func (*Aux_matrix_typeContext) IsAux_matrix_typeContext() {}

func NewAux_matrix_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aux_matrix_typeContext {
	var p = new(Aux_matrix_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_aux_matrix_type

	return p
}

func (s *Aux_matrix_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Aux_matrix_typeContext) CORCHETE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_IZQ, 0)
}

func (s *Aux_matrix_typeContext) Matrix_type() IMatrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_typeContext)
}

func (s *Aux_matrix_typeContext) CORCHETE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_DER, 0)
}

func (s *Aux_matrix_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aux_matrix_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Aux_matrix_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAux_matrix_type(s)
	}
}

func (s *Aux_matrix_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAux_matrix_type(s)
	}
}

func (s *Aux_matrix_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAux_matrix_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Aux_matrix_type() (localctx IAux_matrix_typeContext) {
	localctx = NewAux_matrix_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VGrammarRULE_aux_matrix_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(VGrammarCORCHETE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.Matrix_type()
	}
	{
		p.SetState(237)
		p.Match(VGrammarCORCHETE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssign_stmtContext is an interface to support dynamic dispatch.
type IAssign_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssign_stmtContext differentiates from other interfaces.
	IsAssign_stmtContext()
}

type Assign_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_stmtContext() *Assign_stmtContext {
	var p = new(Assign_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_assign_stmt
	return p
}

func InitEmptyAssign_stmtContext(p *Assign_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_assign_stmt
}

func (*Assign_stmtContext) IsAssign_stmtContext() {}

func NewAssign_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_stmtContext {
	var p = new(Assign_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_assign_stmt

	return p
}

func (s *Assign_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_stmtContext) CopyAll(ctx *Assign_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Assign_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AsignacionVectorContext struct {
	Assign_stmtContext
	op antlr.Token
}

func NewAsignacionVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionVectorContext {
	var p = new(AsignacionVectorContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *AsignacionVectorContext) GetOp() antlr.Token { return s.op }

func (s *AsignacionVectorContext) SetOp(v antlr.Token) { s.op = v }

func (s *AsignacionVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionVectorContext) Vector_item() IVector_itemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_itemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_itemContext)
}

func (s *AsignacionVectorContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionVectorContext) OP_INCREMENTO() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_INCREMENTO, 0)
}

func (s *AsignacionVectorContext) OP_DECREMENTO() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECREMENTO, 0)
}

func (s *AsignacionVectorContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *AsignacionVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAsignacionVector(s)
	}
}

func (s *AsignacionVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAsignacionVector(s)
	}
}

func (s *AsignacionVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAsignacionVector(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionDirectaContext struct {
	Assign_stmtContext
}

func NewAsignacionDirectaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionDirectaContext {
	var p = new(AsignacionDirectaContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *AsignacionDirectaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionDirectaContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *AsignacionDirectaContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *AsignacionDirectaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionDirectaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAsignacionDirecta(s)
	}
}

func (s *AsignacionDirectaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAsignacionDirecta(s)
	}
}

func (s *AsignacionDirectaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAsignacionDirecta(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignVectorItemContext struct {
	Assign_stmtContext
}

func NewAssignVectorItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignVectorItemContext {
	var p = new(AssignVectorItemContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *AssignVectorItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignVectorItemContext) Vector_item() IVector_itemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_itemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_itemContext)
}

func (s *AssignVectorItemContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *AssignVectorItemContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignVectorItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAssignVectorItem(s)
	}
}

func (s *AssignVectorItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAssignVectorItem(s)
	}
}

func (s *AssignVectorItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAssignVectorItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type AsignacionAritmeticaContext struct {
	Assign_stmtContext
	op antlr.Token
}

func NewAsignacionAritmeticaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionAritmeticaContext {
	var p = new(AsignacionAritmeticaContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *AsignacionAritmeticaContext) GetOp() antlr.Token { return s.op }

func (s *AsignacionAritmeticaContext) SetOp(v antlr.Token) { s.op = v }

func (s *AsignacionAritmeticaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionAritmeticaContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *AsignacionAritmeticaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignacionAritmeticaContext) OP_INCREMENTO() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_INCREMENTO, 0)
}

func (s *AsignacionAritmeticaContext) OP_DECREMENTO() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECREMENTO, 0)
}

func (s *AsignacionAritmeticaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAsignacionAritmetica(s)
	}
}

func (s *AsignacionAritmeticaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAsignacionAritmetica(s)
	}
}

func (s *AsignacionAritmeticaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAsignacionAritmetica(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Assign_stmt() (localctx IAssign_stmtContext) {
	localctx = NewAssign_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, VGrammarRULE_assign_stmt)
	var _la int

	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignVectorItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(239)
			p.Vector_item()
		}
		{
			p.SetState(240)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.expr(0)
		}

	case 2:
		localctx = NewAsignacionDirectaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(243)
			p.Id_pattern()
		}
		{
			p.SetState(244)
			p.Match(VGrammarOP_DECLARACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.expr(0)
		}

	case 3:
		localctx = NewAsignacionAritmeticaContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(247)
			p.Id_pattern()
		}
		{
			p.SetState(248)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AsignacionAritmeticaContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VGrammarOP_INCREMENTO || _la == VGrammarOP_DECREMENTO) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AsignacionAritmeticaContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(249)
			p.expr(0)
		}

	case 4:
		localctx = NewAsignacionVectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(251)
			p.Vector_item()
		}
		{
			p.SetState(252)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AsignacionVectorContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&448) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AsignacionVectorContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(253)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IId_patternContext is an interface to support dynamic dispatch.
type IId_patternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsId_patternContext differentiates from other interfaces.
	IsId_patternContext()
}

type Id_patternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_patternContext() *Id_patternContext {
	var p = new(Id_patternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_id_pattern
	return p
}

func InitEmptyId_patternContext(p *Id_patternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_id_pattern
}

func (*Id_patternContext) IsId_patternContext() {}

func NewId_patternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_patternContext {
	var p = new(Id_patternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_id_pattern

	return p
}

func (s *Id_patternContext) GetParser() antlr.Parser { return s.parser }

func (s *Id_patternContext) CopyAll(ctx *Id_patternContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Id_patternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_patternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ID_PatronContext struct {
	Id_patternContext
}

func NewID_PatronContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ID_PatronContext {
	var p = new(ID_PatronContext)

	InitEmptyId_patternContext(&p.Id_patternContext)
	p.parser = parser
	p.CopyAll(ctx.(*Id_patternContext))

	return p
}

func (s *ID_PatronContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ID_PatronContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VGrammarID)
}

func (s *ID_PatronContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarID, i)
}

func (s *ID_PatronContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(VGrammarPUNTO)
}

func (s *ID_PatronContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO, i)
}

func (s *ID_PatronContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterID_Patron(s)
	}
}

func (s *ID_PatronContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitID_Patron(s)
	}
}

func (s *ID_PatronContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitID_Patron(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Id_pattern() (localctx IId_patternContext) {
	localctx = NewId_patternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, VGrammarRULE_id_pattern)
	var _alt int

	localctx = NewID_PatronContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(258)
				p.Match(VGrammarPUNTO)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(259)
				p.Match(VGrammarID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringLiteralContext struct {
	LiteralContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarSTRING_LITERAL, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolLiteralContext struct {
	LiteralContext
}

func NewBoolLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLiteralContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarBOOL_LITERAL, 0)
}

func (s *BoolLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitBoolLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatLiteralContext struct {
	LiteralContext
}

func NewFloatLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarFLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilLiteralContext struct {
	LiteralContext
}

func NewNilLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilLiteralContext {
	var p = new(NilLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *NilLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilLiteralContext) NIL_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarNIL_LITERAL, 0)
}

func (s *NilLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterNilLiteral(s)
	}
}

func (s *NilLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitNilLiteral(s)
	}
}

func (s *NilLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitNilLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntLiteralContext struct {
	LiteralContext
}

func NewIntLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLiteralContext {
	var p = new(IntLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarINT_LITERAL, 0)
}

func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterIntLiteral(s)
	}
}

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitIntLiteral(s)
	}
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitIntLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, VGrammarRULE_literal)
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VGrammarINT_LITERAL:
		localctx = NewIntLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
			p.Match(VGrammarINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarFLOAT_LITERAL:
		localctx = NewFloatLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.Match(VGrammarFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarSTRING_LITERAL:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(267)
			p.Match(VGrammarSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarBOOL_LITERAL:
		localctx = NewBoolLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(268)
			p.Match(VGrammarBOOL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarNIL_LITERAL:
		localctx = NewNilLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(269)
			p.Match(VGrammarNIL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LiteralExpContext struct {
	ExprContext
}

func NewLiteralExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpContext {
	var p = new(LiteralExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LiteralExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLiteralExp(s)
	}
}

func (s *LiteralExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLiteralExp(s)
	}
}

func (s *LiteralExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLiteralExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructMethodExpContext struct {
	ExprContext
}

func NewStructMethodExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructMethodExpContext {
	var p = new(StructMethodExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StructMethodExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMethodExpContext) Struct_method_call() IStruct_method_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_method_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_method_callContext)
}

func (s *StructMethodExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStructMethodExp(s)
	}
}

func (s *StructMethodExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStructMethodExp(s)
	}
}

func (s *StructMethodExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStructMethodExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructInitExpContext struct {
	ExprContext
}

func NewStructInitExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructInitExpContext {
	var p = new(StructInitExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StructInitExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInitExpContext) Struct_init() IStruct_initContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_initContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_initContext)
}

func (s *StructInitExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStructInitExp(s)
	}
}

func (s *StructInitExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStructInitExp(s)
	}
}

func (s *StructInitExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStructInitExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type RepeatingExpContext struct {
	ExprContext
}

func NewRepeatingExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RepeatingExpContext {
	var p = new(RepeatingExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RepeatingExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatingExpContext) Repeating() IRepeatingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepeatingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepeatingContext)
}

func (s *RepeatingExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterRepeatingExp(s)
	}
}

func (s *RepeatingExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitRepeatingExp(s)
	}
}

func (s *RepeatingExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitRepeatingExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructExpContext struct {
	ExprContext
}

func NewStructExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructExpContext {
	var p = new(StructExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StructExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructExpContext) Struct_vector() IStruct_vectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_vectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_vectorContext)
}

func (s *StructExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStructExp(s)
	}
}

func (s *StructExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStructExp(s)
	}
}

func (s *StructExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStructExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpBinarioContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewExpBinarioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpBinarioContext {
	var p = new(ExpBinarioContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExpBinarioContext) GetOp() antlr.Token { return s.op }

func (s *ExpBinarioContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpBinarioContext) GetLeft() IExprContext { return s.left }

func (s *ExpBinarioContext) GetRight() IExprContext { return s.right }

func (s *ExpBinarioContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExpBinarioContext) SetRight(v IExprContext) { s.right = v }

func (s *ExpBinarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpBinarioContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExpBinarioContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpBinarioContext) OP_MULT() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MULT, 0)
}

func (s *ExpBinarioContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DIV, 0)
}

func (s *ExpBinarioContext) OP_MOD() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MOD, 0)
}

func (s *ExpBinarioContext) OP_SUMA() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_SUMA, 0)
}

func (s *ExpBinarioContext) OP_RESTA() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_RESTA, 0)
}

func (s *ExpBinarioContext) OP_MENORQ() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MENORQ, 0)
}

func (s *ExpBinarioContext) OP_MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MENORIGUAL, 0)
}

func (s *ExpBinarioContext) OP_MAYORQ() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MAYORQ, 0)
}

func (s *ExpBinarioContext) OP_MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MAYORIGUAL, 0)
}

func (s *ExpBinarioContext) OP_IGUAL() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_IGUAL, 0)
}

func (s *ExpBinarioContext) OP_DIFERENTE() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DIFERENTE, 0)
}

func (s *ExpBinarioContext) OP_AND() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_AND, 0)
}

func (s *ExpBinarioContext) OP_OR() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_OR, 0)
}

func (s *ExpBinarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExpBinario(s)
	}
}

func (s *ExpBinarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExpBinario(s)
	}
}

func (s *ExpBinarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExpBinario(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorPropExpContext struct {
	ExprContext
}

func NewVectorPropExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorPropExpContext {
	var p = new(VectorPropExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VectorPropExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorPropExpContext) Vector_prop() IVector_propContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_propContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_propContext)
}

func (s *VectorPropExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterVectorPropExp(s)
	}
}

func (s *VectorPropExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitVectorPropExp(s)
	}
}

func (s *VectorPropExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitVectorPropExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorFuncExpContext struct {
	ExprContext
}

func NewVectorFuncExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorFuncExpContext {
	var p = new(VectorFuncExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VectorFuncExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorFuncExpContext) Vector_func() IVector_funcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_funcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_funcContext)
}

func (s *VectorFuncExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterVectorFuncExp(s)
	}
}

func (s *VectorFuncExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitVectorFuncExp(s)
	}
}

func (s *VectorFuncExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitVectorFuncExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParentecisExpContext struct {
	ExprContext
}

func NewParentecisExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentecisExpContext {
	var p = new(ParentecisExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParentecisExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentecisExpContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *ParentecisExpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParentecisExpContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *ParentecisExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterParentecisExp(s)
	}
}

func (s *ParentecisExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitParentecisExp(s)
	}
}

func (s *ParentecisExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitParentecisExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpUnaryContext struct {
	ExprContext
	op antlr.Token
}

func NewExpUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpUnaryContext {
	var p = new(ExpUnaryContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExpUnaryContext) GetOp() antlr.Token { return s.op }

func (s *ExpUnaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpUnaryContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpUnaryContext) OP_NOT() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_NOT, 0)
}

func (s *ExpUnaryContext) OP_RESTA() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_RESTA, 0)
}

func (s *ExpUnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExpUnary(s)
	}
}

func (s *ExpUnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExpUnary(s)
	}
}

func (s *ExpUnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExpUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExpContext struct {
	ExprContext
}

func NewIdExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExpContext {
	var p = new(IdExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExpContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *IdExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterIdExp(s)
	}
}

func (s *IdExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitIdExp(s)
	}
}

func (s *IdExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitIdExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallExpContext struct {
	ExprContext
}

func NewFunctionCallExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpContext {
	var p = new(FunctionCallExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FunctionCallExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpContext) Func_call() IFunc_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_callContext)
}

func (s *FunctionCallExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFunctionCallExp(s)
	}
}

func (s *FunctionCallExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFunctionCallExp(s)
	}
}

func (s *FunctionCallExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFunctionCallExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorItemExpContext struct {
	ExprContext
}

func NewVectorItemExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorItemExpContext {
	var p = new(VectorItemExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VectorItemExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorItemExpContext) Vector_item() IVector_itemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_itemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_itemContext)
}

func (s *VectorItemExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterVectorItemExp(s)
	}
}

func (s *VectorItemExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitVectorItemExp(s)
	}
}

func (s *VectorItemExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitVectorItemExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type VectorExpContext struct {
	ExprContext
}

func NewVectorExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VectorExpContext {
	var p = new(VectorExpContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VectorExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorExpContext) Vector_expr() IVector_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVector_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVector_exprContext)
}

func (s *VectorExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterVectorExp(s)
	}
}

func (s *VectorExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitVectorExp(s)
	}
}

func (s *VectorExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitVectorExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *VGrammar) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, VGrammarRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParentecisExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(273)
			p.Match(VGrammarPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.expr(0)
		}
		{
			p.SetState(275)
			p.Match(VGrammarPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFunctionCallExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(277)
			p.Func_call()
		}

	case 3:
		localctx = NewIdExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(278)
			p.Id_pattern()
		}

	case 4:
		localctx = NewVectorItemExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(279)
			p.Vector_item()
		}

	case 5:
		localctx = NewVectorPropExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(280)
			p.Vector_prop()
		}

	case 6:
		localctx = NewVectorFuncExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(281)
			p.Vector_func()
		}

	case 7:
		localctx = NewLiteralExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(282)
			p.Literal()
		}

	case 8:
		localctx = NewVectorExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(283)
			p.Vector_expr()
		}

	case 9:
		localctx = NewRepeatingExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(284)
			p.Repeating()
		}

	case 10:
		localctx = NewStructExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(285)
			p.Struct_vector()
		}

	case 11:
		localctx = NewStructInitExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(286)
			p.Struct_init()
		}

	case 12:
		localctx = NewStructMethodExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(287)
			p.Struct_method_call()
		}

	case 13:
		localctx = NewExpUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(288)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpUnaryContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == VGrammarOP_RESTA || _la == VGrammarOP_NOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpUnaryContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(289)
			p.expr(7)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(310)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpBinarioContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExpBinarioContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(292)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(293)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpBinarioContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7340032) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpBinarioContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(294)

					var _x = p.expr(7)

					localctx.(*ExpBinarioContext).right = _x
				}

			case 2:
				localctx = NewExpBinarioContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExpBinarioContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(295)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(296)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpBinarioContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VGrammarOP_SUMA || _la == VGrammarOP_RESTA) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpBinarioContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(297)

					var _x = p.expr(6)

					localctx.(*ExpBinarioContext).right = _x
				}

			case 3:
				localctx = NewExpBinarioContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExpBinarioContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(298)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(299)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpBinarioContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&503316480) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpBinarioContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(300)

					var _x = p.expr(5)

					localctx.(*ExpBinarioContext).right = _x
				}

			case 4:
				localctx = NewExpBinarioContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExpBinarioContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(301)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(302)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpBinarioContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VGrammarOP_IGUAL || _la == VGrammarOP_DIFERENTE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpBinarioContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(303)

					var _x = p.expr(4)

					localctx.(*ExpBinarioContext).right = _x
				}

			case 5:
				localctx = NewExpBinarioContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExpBinarioContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(304)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(305)

					var _m = p.Match(VGrammarOP_AND)

					localctx.(*ExpBinarioContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(306)

					var _x = p.expr(3)

					localctx.(*ExpBinarioContext).right = _x
				}

			case 6:
				localctx = NewExpBinarioContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ExpBinarioContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(307)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(308)

					var _m = p.Match(VGrammarOP_OR)

					localctx.(*ExpBinarioContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(309)

					var _x = p.expr(2)

					localctx.(*ExpBinarioContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_if_stmt
	return p
}

func InitEmptyIf_stmtContext(p *If_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_if_stmt
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) CopyAll(ctx *If_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IFstmtContext struct {
	If_stmtContext
}

func NewIFstmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IFstmtContext {
	var p = new(IFstmtContext)

	InitEmptyIf_stmtContext(&p.If_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_stmtContext))

	return p
}

func (s *IFstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IFstmtContext) AllIf_chain() []IIf_chainContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIf_chainContext); ok {
			len++
		}
	}

	tst := make([]IIf_chainContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIf_chainContext); ok {
			tst[i] = t.(IIf_chainContext)
			i++
		}
	}

	return tst
}

func (s *IFstmtContext) If_chain(i int) IIf_chainContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_chainContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_chainContext)
}

func (s *IFstmtContext) AllRW_ELSE() []antlr.TerminalNode {
	return s.GetTokens(VGrammarRW_ELSE)
}

func (s *IFstmtContext) RW_ELSE(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarRW_ELSE, i)
}

func (s *IFstmtContext) Else_stmt() IElse_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_stmtContext)
}

func (s *IFstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterIFstmt(s)
	}
}

func (s *IFstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitIFstmt(s)
	}
}

func (s *IFstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitIFstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, VGrammarRULE_if_stmt)
	var _la int

	var _alt int

	localctx = NewIFstmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.If_chain()
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(316)
				p.Match(VGrammarRW_ELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(317)
				p.If_chain()
			}

		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarRW_ELSE {
		{
			p.SetState(323)
			p.Else_stmt()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_chainContext is an interface to support dynamic dispatch.
type IIf_chainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_chainContext differentiates from other interfaces.
	IsIf_chainContext()
}

type If_chainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_chainContext() *If_chainContext {
	var p = new(If_chainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_if_chain
	return p
}

func InitEmptyIf_chainContext(p *If_chainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_if_chain
}

func (*If_chainContext) IsIf_chainContext() {}

func NewIf_chainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_chainContext {
	var p = new(If_chainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_if_chain

	return p
}

func (s *If_chainContext) GetParser() antlr.Parser { return s.parser }

func (s *If_chainContext) CopyAll(ctx *If_chainContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_chainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_chainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IFcadenaContext struct {
	If_chainContext
}

func NewIFcadenaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IFcadenaContext {
	var p = new(IFcadenaContext)

	InitEmptyIf_chainContext(&p.If_chainContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_chainContext))

	return p
}

func (s *IFcadenaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IFcadenaContext) RW_IF() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_IF, 0)
}

func (s *IFcadenaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IFcadenaContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *IFcadenaContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *IFcadenaContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *IFcadenaContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *IFcadenaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterIFcadena(s)
	}
}

func (s *IFcadenaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitIFcadena(s)
	}
}

func (s *IFcadenaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitIFcadena(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) If_chain() (localctx IIf_chainContext) {
	localctx = NewIf_chainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, VGrammarRULE_if_chain)
	var _la int

	localctx = NewIFcadenaContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.Match(VGrammarRW_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(327)
		p.expr(0)
	}
	{
		p.SetState(328)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&735770847477776) != 0 {
		{
			p.SetState(329)
			p.Stmt()
		}

		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(335)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElse_stmtContext is an interface to support dynamic dispatch.
type IElse_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElse_stmtContext differentiates from other interfaces.
	IsElse_stmtContext()
}

type Else_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_stmtContext() *Else_stmtContext {
	var p = new(Else_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_else_stmt
	return p
}

func InitEmptyElse_stmtContext(p *Else_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_else_stmt
}

func (*Else_stmtContext) IsElse_stmtContext() {}

func NewElse_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_stmtContext {
	var p = new(Else_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_else_stmt

	return p
}

func (s *Else_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_stmtContext) CopyAll(ctx *Else_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Else_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElseStmtContext struct {
	Else_stmtContext
}

func NewElseStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseStmtContext {
	var p = new(ElseStmtContext)

	InitEmptyElse_stmtContext(&p.Else_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Else_stmtContext))

	return p
}

func (s *ElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStmtContext) RW_ELSE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_ELSE, 0)
}

func (s *ElseStmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *ElseStmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *ElseStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ElseStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ElseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterElseStmt(s)
	}
}

func (s *ElseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitElseStmt(s)
	}
}

func (s *ElseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitElseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Else_stmt() (localctx IElse_stmtContext) {
	localctx = NewElse_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, VGrammarRULE_else_stmt)
	var _la int

	localctx = NewElseStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(VGrammarRW_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(338)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&735770847477776) != 0 {
		{
			p.SetState(339)
			p.Stmt()
		}

		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(345)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_stmtContext is an interface to support dynamic dispatch.
type ISwitch_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_stmtContext differentiates from other interfaces.
	IsSwitch_stmtContext()
}

type Switch_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_stmtContext() *Switch_stmtContext {
	var p = new(Switch_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_switch_stmt
	return p
}

func InitEmptySwitch_stmtContext(p *Switch_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_switch_stmt
}

func (*Switch_stmtContext) IsSwitch_stmtContext() {}

func NewSwitch_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_stmtContext {
	var p = new(Switch_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_switch_stmt

	return p
}

func (s *Switch_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_stmtContext) CopyAll(ctx *Switch_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchStmtContext struct {
	Switch_stmtContext
}

func NewSwitchStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	InitEmptySwitch_stmtContext(&p.Switch_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_stmtContext))

	return p
}

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) RW_SWITCH() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_SWITCH, 0)
}

func (s *SwitchStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchStmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *SwitchStmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *SwitchStmtContext) AllSwitch_case() []ISwitch_caseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			len++
		}
	}

	tst := make([]ISwitch_caseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitch_caseContext); ok {
			tst[i] = t.(ISwitch_caseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStmtContext) Switch_case(i int) ISwitch_caseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_caseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_caseContext)
}

func (s *SwitchStmtContext) Default_case() IDefault_caseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefault_caseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefault_caseContext)
}

func (s *SwitchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitSwitchStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Switch_stmt() (localctx ISwitch_stmtContext) {
	localctx = NewSwitch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, VGrammarRULE_switch_stmt)
	var _la int

	localctx = NewSwitchStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Match(VGrammarRW_SWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)
		p.expr(0)
	}
	{
		p.SetState(349)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VGrammarRW_CASE {
		{
			p.SetState(350)
			p.Switch_case()
		}

		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarRW_DEFAULT {
		{
			p.SetState(356)
			p.Default_case()
		}

	}
	{
		p.SetState(359)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_caseContext is an interface to support dynamic dispatch.
type ISwitch_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_caseContext differentiates from other interfaces.
	IsSwitch_caseContext()
}

type Switch_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_caseContext() *Switch_caseContext {
	var p = new(Switch_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_switch_case
	return p
}

func InitEmptySwitch_caseContext(p *Switch_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_switch_case
}

func (*Switch_caseContext) IsSwitch_caseContext() {}

func NewSwitch_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_caseContext {
	var p = new(Switch_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_switch_case

	return p
}

func (s *Switch_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_caseContext) CopyAll(ctx *Switch_caseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchCaseContext struct {
	Switch_caseContext
}

func NewSwitchCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchCaseContext {
	var p = new(SwitchCaseContext)

	InitEmptySwitch_caseContext(&p.Switch_caseContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_caseContext))

	return p
}

func (s *SwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchCaseContext) RW_CASE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_CASE, 0)
}

func (s *SwitchCaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchCaseContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, 0)
}

func (s *SwitchCaseContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *SwitchCaseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *SwitchCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterSwitchCase(s)
	}
}

func (s *SwitchCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitSwitchCase(s)
	}
}

func (s *SwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitSwitchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Switch_case() (localctx ISwitch_caseContext) {
	localctx = NewSwitch_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, VGrammarRULE_switch_case)
	var _la int

	localctx = NewSwitchCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		p.Match(VGrammarRW_CASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(362)
		p.expr(0)
	}
	{
		p.SetState(363)
		p.Match(VGrammarDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&735770847477776) != 0 {
		{
			p.SetState(364)
			p.Stmt()
		}

		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefault_caseContext is an interface to support dynamic dispatch.
type IDefault_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDefault_caseContext differentiates from other interfaces.
	IsDefault_caseContext()
}

type Default_caseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_caseContext() *Default_caseContext {
	var p = new(Default_caseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_default_case
	return p
}

func InitEmptyDefault_caseContext(p *Default_caseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_default_case
}

func (*Default_caseContext) IsDefault_caseContext() {}

func NewDefault_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_caseContext {
	var p = new(Default_caseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_default_case

	return p
}

func (s *Default_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Default_caseContext) CopyAll(ctx *Default_caseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Default_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefaultCaseContext struct {
	Default_caseContext
}

func NewDefaultCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultCaseContext {
	var p = new(DefaultCaseContext)

	InitEmptyDefault_caseContext(&p.Default_caseContext)
	p.parser = parser
	p.CopyAll(ctx.(*Default_caseContext))

	return p
}

func (s *DefaultCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultCaseContext) RW_DEFAULT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_DEFAULT, 0)
}

func (s *DefaultCaseContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, 0)
}

func (s *DefaultCaseContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *DefaultCaseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *DefaultCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDefaultCase(s)
	}
}

func (s *DefaultCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDefaultCase(s)
	}
}

func (s *DefaultCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDefaultCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Default_case() (localctx IDefault_caseContext) {
	localctx = NewDefault_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, VGrammarRULE_default_case)
	var _la int

	localctx = NewDefaultCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		p.Match(VGrammarRW_DEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Match(VGrammarDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&735770847477776) != 0 {
		{
			p.SetState(372)
			p.Stmt()
		}

		p.SetState(377)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhile_stmtContext is an interface to support dynamic dispatch.
type IWhile_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWhile_stmtContext differentiates from other interfaces.
	IsWhile_stmtContext()
}

type While_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_stmtContext() *While_stmtContext {
	var p = new(While_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_while_stmt
	return p
}

func InitEmptyWhile_stmtContext(p *While_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_while_stmt
}

func (*While_stmtContext) IsWhile_stmtContext() {}

func NewWhile_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmtContext {
	var p = new(While_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_while_stmt

	return p
}

func (s *While_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmtContext) CopyAll(ctx *While_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *While_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type WhileStmtContext struct {
	While_stmtContext
}

func NewWhileStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStmtContext {
	var p = new(WhileStmtContext)

	InitEmptyWhile_stmtContext(&p.While_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*While_stmtContext))

	return p
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) RW_FOR() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FOR, 0)
}

func (s *WhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *WhileStmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *WhileStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *WhileStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) While_stmt() (localctx IWhile_stmtContext) {
	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, VGrammarRULE_while_stmt)
	var _la int

	localctx = NewWhileStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Match(VGrammarRW_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(379)
		p.expr(0)
	}
	{
		p.SetState(380)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&735770847477776) != 0 {
		{
			p.SetState(381)
			p.Stmt()
		}

		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(387)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFor_stmtContext differentiates from other interfaces.
	IsFor_stmtContext()
}

type For_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmtContext() *For_stmtContext {
	var p = new(For_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_for_stmt
	return p
}

func InitEmptyFor_stmtContext(p *For_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_for_stmt
}

func (*For_stmtContext) IsFor_stmtContext() {}

func NewFor_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmtContext {
	var p = new(For_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_for_stmt

	return p
}

func (s *For_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmtContext) CopyAll(ctx *For_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForStmtContext struct {
	For_stmtContext
}

func NewForStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStmtContext {
	var p = new(ForStmtContext)

	InitEmptyFor_stmtContext(&p.For_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_stmtContext))

	return p
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) RW_FOR() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FOR, 0)
}

func (s *ForStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *ForStmtContext) RW_IN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_IN, 0)
}

func (s *ForStmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *ForStmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *ForStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForStmtContext) Range_() IRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeContext)
}

func (s *ForStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) For_stmt() (localctx IFor_stmtContext) {
	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, VGrammarRULE_for_stmt)
	var _la int

	localctx = NewForStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
		p.Match(VGrammarRW_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(390)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(391)
		p.Match(VGrammarRW_IN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(392)
			p.expr(0)
		}

	case 2:
		{
			p.SetState(393)
			p.Range_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(396)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&735770847477776) != 0 {
		{
			p.SetState(397)
			p.Stmt()
		}

		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(403)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeContext is an interface to support dynamic dispatch.
type IRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRangeContext differentiates from other interfaces.
	IsRangeContext()
}

type RangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeContext() *RangeContext {
	var p = new(RangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_range
	return p
}

func InitEmptyRangeContext(p *RangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_range
}

func (*RangeContext) IsRangeContext() {}

func NewRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeContext {
	var p = new(RangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_range

	return p
}

func (s *RangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeContext) CopyAll(ctx *RangeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RangoNumContext struct {
	RangeContext
}

func NewRangoNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RangoNumContext {
	var p = new(RangoNumContext)

	InitEmptyRangeContext(&p.RangeContext)
	p.parser = parser
	p.CopyAll(ctx.(*RangeContext))

	return p
}

func (s *RangoNumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangoNumContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RangoNumContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RangoNumContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(VGrammarPUNTO)
}

func (s *RangoNumContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO, i)
}

func (s *RangoNumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterRangoNum(s)
	}
}

func (s *RangoNumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitRangoNum(s)
	}
}

func (s *RangoNumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitRangoNum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Range_() (localctx IRangeContext) {
	localctx = NewRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, VGrammarRULE_range)
	localctx = NewRangoNumContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(405)
		p.expr(0)
	}
	{
		p.SetState(406)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(408)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(409)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuard_stmtContext is an interface to support dynamic dispatch.
type IGuard_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsGuard_stmtContext differentiates from other interfaces.
	IsGuard_stmtContext()
}

type Guard_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuard_stmtContext() *Guard_stmtContext {
	var p = new(Guard_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_guard_stmt
	return p
}

func InitEmptyGuard_stmtContext(p *Guard_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_guard_stmt
}

func (*Guard_stmtContext) IsGuard_stmtContext() {}

func NewGuard_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Guard_stmtContext {
	var p = new(Guard_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_guard_stmt

	return p
}

func (s *Guard_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Guard_stmtContext) CopyAll(ctx *Guard_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Guard_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Guard_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GuardStmtContext struct {
	Guard_stmtContext
}

func NewGuardStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GuardStmtContext {
	var p = new(GuardStmtContext)

	InitEmptyGuard_stmtContext(&p.Guard_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Guard_stmtContext))

	return p
}

func (s *GuardStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardStmtContext) RW_GUARD() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_GUARD, 0)
}

func (s *GuardStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardStmtContext) RW_ELSE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_ELSE, 0)
}

func (s *GuardStmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *GuardStmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *GuardStmtContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *GuardStmtContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *GuardStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterGuardStmt(s)
	}
}

func (s *GuardStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitGuardStmt(s)
	}
}

func (s *GuardStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitGuardStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Guard_stmt() (localctx IGuard_stmtContext) {
	localctx = NewGuard_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, VGrammarRULE_guard_stmt)
	var _la int

	localctx = NewGuardStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(411)
		p.Match(VGrammarRW_GUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(412)
		p.expr(0)
	}
	{
		p.SetState(413)
		p.Match(VGrammarRW_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(414)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&735770847477776) != 0 {
		{
			p.SetState(415)
			p.Stmt()
		}

		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(421)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITransfer_stmtContext is an interface to support dynamic dispatch.
type ITransfer_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTransfer_stmtContext differentiates from other interfaces.
	IsTransfer_stmtContext()
}

type Transfer_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransfer_stmtContext() *Transfer_stmtContext {
	var p = new(Transfer_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_transfer_stmt
	return p
}

func InitEmptyTransfer_stmtContext(p *Transfer_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_transfer_stmt
}

func (*Transfer_stmtContext) IsTransfer_stmtContext() {}

func NewTransfer_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Transfer_stmtContext {
	var p = new(Transfer_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_transfer_stmt

	return p
}

func (s *Transfer_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Transfer_stmtContext) CopyAll(ctx *Transfer_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Transfer_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Transfer_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ContinueStmtContext struct {
	Transfer_stmtContext
}

func NewContinueStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) RW_CONTINUE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_CONTINUE, 0)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStmtContext struct {
	Transfer_stmtContext
}

func NewBreakStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStmtContext {
	var p = new(BreakStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) RW_BREAK() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_BREAK, 0)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStmtContext struct {
	Transfer_stmtContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	InitEmptyTransfer_stmtContext(&p.Transfer_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Transfer_stmtContext))

	return p
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) RW_RETURN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_RETURN, 0)
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Transfer_stmt() (localctx ITransfer_stmtContext) {
	localctx = NewTransfer_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, VGrammarRULE_transfer_stmt)
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VGrammarRW_RETURN:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(423)
			p.Match(VGrammarRW_RETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(425)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(424)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case VGrammarRW_BREAK:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(427)
			p.Match(VGrammarRW_BREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_CONTINUE:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(428)
			p.Match(VGrammarRW_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_callContext is an interface to support dynamic dispatch.
type IFunc_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_callContext differentiates from other interfaces.
	IsFunc_callContext()
}

type Func_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_callContext() *Func_callContext {
	var p = new(Func_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_func_call
	return p
}

func InitEmptyFunc_callContext(p *Func_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_func_call
}

func (*Func_callContext) IsFunc_callContext() {}

func NewFunc_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_callContext {
	var p = new(Func_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_func_call

	return p
}

func (s *Func_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_callContext) CopyAll(ctx *Func_callContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LlamarFuncionContext struct {
	Func_callContext
}

func NewLlamarFuncionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LlamarFuncionContext {
	var p = new(LlamarFuncionContext)

	InitEmptyFunc_callContext(&p.Func_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_callContext))

	return p
}

func (s *LlamarFuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamarFuncionContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *LlamarFuncionContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *LlamarFuncionContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *LlamarFuncionContext) Arg_list() IArg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *LlamarFuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLlamarFuncion(s)
	}
}

func (s *LlamarFuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLlamarFuncion(s)
	}
}

func (s *LlamarFuncionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLlamarFuncion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Func_call() (localctx IFunc_callContext) {
	localctx = NewFunc_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, VGrammarRULE_func_call)
	var _la int

	localctx = NewLlamarFuncionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(431)
		p.Id_pattern()
	}
	{
		p.SetState(432)
		p.Match(VGrammarPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4635892868572766208) != 0 {
		{
			p.SetState(433)
			p.Arg_list()
		}

	}
	{
		p.SetState(436)
		p.Match(VGrammarPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArg_listContext is an interface to support dynamic dispatch.
type IArg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArg_listContext differentiates from other interfaces.
	IsArg_listContext()
}

type Arg_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArg_listContext() *Arg_listContext {
	var p = new(Arg_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_arg_list
	return p
}

func InitEmptyArg_listContext(p *Arg_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_arg_list
}

func (*Arg_listContext) IsArg_listContext() {}

func NewArg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arg_listContext {
	var p = new(Arg_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_arg_list

	return p
}

func (s *Arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Arg_listContext) CopyAll(ctx *Arg_listContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArgListContext struct {
	Arg_listContext
}

func NewArgListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgListContext {
	var p = new(ArgListContext)

	InitEmptyArg_listContext(&p.Arg_listContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arg_listContext))

	return p
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) AllFunc_arg() []IFunc_argContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_argContext); ok {
			len++
		}
	}

	tst := make([]IFunc_argContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_argContext); ok {
			tst[i] = t.(IFunc_argContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Func_arg(i int) IFunc_argContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_argContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_argContext)
}

func (s *ArgListContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *ArgListContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Arg_list() (localctx IArg_listContext) {
	localctx = NewArg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, VGrammarRULE_arg_list)
	var _la int

	var _alt int

	localctx = NewArgListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(438)
		p.Func_arg()
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(439)
				p.Match(VGrammarCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(440)
				p.Func_arg()
			}

		}
		p.SetState(445)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarCOMA {
		{
			p.SetState(446)
			p.Match(VGrammarCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_argContext is an interface to support dynamic dispatch.
type IFunc_argContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_argContext differentiates from other interfaces.
	IsFunc_argContext()
}

type Func_argContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_argContext() *Func_argContext {
	var p = new(Func_argContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_func_arg
	return p
}

func InitEmptyFunc_argContext(p *Func_argContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_func_arg
}

func (*Func_argContext) IsFunc_argContext() {}

func NewFunc_argContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_argContext {
	var p = new(Func_argContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_func_arg

	return p
}

func (s *Func_argContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_argContext) CopyAll(ctx *Func_argContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_argContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_argContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncionArgContext struct {
	Func_argContext
}

func NewFuncionArgContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionArgContext {
	var p = new(FuncionArgContext)

	InitEmptyFunc_argContext(&p.Func_argContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_argContext))

	return p
}

func (s *FuncionArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionArgContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *FuncionArgContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FuncionArgContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *FuncionArgContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, 0)
}

func (s *FuncionArgContext) ANPERSAND() antlr.TerminalNode {
	return s.GetToken(VGrammarANPERSAND, 0)
}

func (s *FuncionArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFuncionArg(s)
	}
}

func (s *FuncionArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFuncionArg(s)
	}
}

func (s *FuncionArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFuncionArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Func_arg() (localctx IFunc_argContext) {
	localctx = NewFunc_argContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, VGrammarRULE_func_arg)
	var _la int

	localctx = NewFuncionArgContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(451)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(449)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(450)
			p.Match(VGrammarDOS_PUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarANPERSAND {
		{
			p.SetState(453)
			p.Match(VGrammarANPERSAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(456)
			p.Id_pattern()
		}

	case 2:
		{
			p.SetState(457)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_dclContext is an interface to support dynamic dispatch.
type IFunc_dclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_dclContext differentiates from other interfaces.
	IsFunc_dclContext()
}

type Func_dclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_dclContext() *Func_dclContext {
	var p = new(Func_dclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_func_dcl
	return p
}

func InitEmptyFunc_dclContext(p *Func_dclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_func_dcl
}

func (*Func_dclContext) IsFunc_dclContext() {}

func NewFunc_dclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_dclContext {
	var p = new(Func_dclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_func_dcl

	return p
}

func (s *Func_dclContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_dclContext) CopyAll(ctx *Func_dclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_dclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_dclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MetodoStructContext struct {
	Func_dclContext
}

func NewMetodoStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MetodoStructContext {
	var p = new(MetodoStructContext)

	InitEmptyFunc_dclContext(&p.Func_dclContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_dclContext))

	return p
}

func (s *MetodoStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetodoStructContext) RW_FN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FN, 0)
}

func (s *MetodoStructContext) AllPAR_IZQ() []antlr.TerminalNode {
	return s.GetTokens(VGrammarPAR_IZQ)
}

func (s *MetodoStructContext) PAR_IZQ(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, i)
}

func (s *MetodoStructContext) Method_receiver() IMethod_receiverContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMethod_receiverContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMethod_receiverContext)
}

func (s *MetodoStructContext) AllPAR_DER() []antlr.TerminalNode {
	return s.GetTokens(VGrammarPAR_DER)
}

func (s *MetodoStructContext) PAR_DER(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, i)
}

func (s *MetodoStructContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *MetodoStructContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *MetodoStructContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *MetodoStructContext) Param_list() IParam_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_listContext)
}

func (s *MetodoStructContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *MetodoStructContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *MetodoStructContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *MetodoStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterMetodoStruct(s)
	}
}

func (s *MetodoStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitMetodoStruct(s)
	}
}

func (s *MetodoStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitMetodoStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncionDecleradaContext struct {
	Func_dclContext
}

func NewFuncionDecleradaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncionDecleradaContext {
	var p = new(FuncionDecleradaContext)

	InitEmptyFunc_dclContext(&p.Func_dclContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_dclContext))

	return p
}

func (s *FuncionDecleradaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionDecleradaContext) RW_FN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FN, 0)
}

func (s *FuncionDecleradaContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *FuncionDecleradaContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *FuncionDecleradaContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *FuncionDecleradaContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *FuncionDecleradaContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *FuncionDecleradaContext) Param_list() IParam_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_listContext)
}

func (s *FuncionDecleradaContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncionDecleradaContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *FuncionDecleradaContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *FuncionDecleradaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFuncionDeclerada(s)
	}
}

func (s *FuncionDecleradaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFuncionDeclerada(s)
	}
}

func (s *FuncionDecleradaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFuncionDeclerada(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Func_dcl() (localctx IFunc_dclContext) {
	localctx = NewFunc_dclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, VGrammarRULE_func_dcl)
	var _la int

	p.SetState(500)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFuncionDecleradaContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(460)
			p.Match(VGrammarRW_FN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.Match(VGrammarPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(464)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VGrammarID {
			{
				p.SetState(463)
				p.Param_list()
			}

		}
		{
			p.SetState(466)
			p.Match(VGrammarPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18577348462910976) != 0 {
			{
				p.SetState(467)
				p.Type_()
			}

		}
		{
			p.SetState(470)
			p.Match(VGrammarLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&735770847477776) != 0 {
			{
				p.SetState(471)
				p.Stmt()
			}

			p.SetState(476)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(477)
			p.Match(VGrammarLLAVE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewMetodoStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(478)
			p.Match(VGrammarRW_FN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(479)
			p.Match(VGrammarPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(480)
			p.Method_receiver()
		}
		{
			p.SetState(481)
			p.Match(VGrammarPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(482)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(483)
			p.Match(VGrammarPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(485)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VGrammarID {
			{
				p.SetState(484)
				p.Param_list()
			}

		}
		{
			p.SetState(487)
			p.Match(VGrammarPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(489)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18577348462910976) != 0 {
			{
				p.SetState(488)
				p.Type_()
			}

		}
		{
			p.SetState(491)
			p.Match(VGrammarLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(495)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&735770847477776) != 0 {
			{
				p.SetState(492)
				p.Stmt()
			}

			p.SetState(497)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(498)
			p.Match(VGrammarLLAVE_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMethod_receiverContext is an interface to support dynamic dispatch.
type IMethod_receiverContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMethod_receiverContext differentiates from other interfaces.
	IsMethod_receiverContext()
}

type Method_receiverContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethod_receiverContext() *Method_receiverContext {
	var p = new(Method_receiverContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_method_receiver
	return p
}

func InitEmptyMethod_receiverContext(p *Method_receiverContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_method_receiver
}

func (*Method_receiverContext) IsMethod_receiverContext() {}

func NewMethod_receiverContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Method_receiverContext {
	var p = new(Method_receiverContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_method_receiver

	return p
}

func (s *Method_receiverContext) GetParser() antlr.Parser { return s.parser }

func (s *Method_receiverContext) CopyAll(ctx *Method_receiverContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Method_receiverContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Method_receiverContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MethodReceiverContext struct {
	Method_receiverContext
}

func NewMethodReceiverContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethodReceiverContext {
	var p = new(MethodReceiverContext)

	InitEmptyMethod_receiverContext(&p.Method_receiverContext)
	p.parser = parser
	p.CopyAll(ctx.(*Method_receiverContext))

	return p
}

func (s *MethodReceiverContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodReceiverContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VGrammarID)
}

func (s *MethodReceiverContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarID, i)
}

func (s *MethodReceiverContext) OP_MULT() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MULT, 0)
}

func (s *MethodReceiverContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterMethodReceiver(s)
	}
}

func (s *MethodReceiverContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitMethodReceiver(s)
	}
}

func (s *MethodReceiverContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitMethodReceiver(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Method_receiver() (localctx IMethod_receiverContext) {
	localctx = NewMethod_receiverContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, VGrammarRULE_method_receiver)
	localctx = NewMethodReceiverContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(502)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(503)
		p.Match(VGrammarOP_MULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(504)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParam_listContext is an interface to support dynamic dispatch.
type IParam_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParam_listContext differentiates from other interfaces.
	IsParam_listContext()
}

type Param_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_listContext() *Param_listContext {
	var p = new(Param_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_param_list
	return p
}

func InitEmptyParam_listContext(p *Param_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_param_list
}

func (*Param_listContext) IsParam_listContext() {}

func NewParam_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_listContext {
	var p = new(Param_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_param_list

	return p
}

func (s *Param_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_listContext) CopyAll(ctx *Param_listContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Param_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParamListContext struct {
	Param_listContext
}

func NewParamListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamListContext {
	var p = new(ParamListContext)

	InitEmptyParam_listContext(&p.Param_listContext)
	p.parser = parser
	p.CopyAll(ctx.(*Param_listContext))

	return p
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) AllFunc_param() []IFunc_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunc_paramContext); ok {
			len++
		}
	}

	tst := make([]IFunc_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunc_paramContext); ok {
			tst[i] = t.(IFunc_paramContext)
			i++
		}
	}

	return tst
}

func (s *ParamListContext) Func_param(i int) IFunc_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_paramContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_paramContext)
}

func (s *ParamListContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *ParamListContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (s *ParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Param_list() (localctx IParam_listContext) {
	localctx = NewParam_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, VGrammarRULE_param_list)
	var _la int

	var _alt int

	localctx = NewParamListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(506)
		p.Func_param()
	}
	p.SetState(511)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(507)
				p.Match(VGrammarCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(508)
				p.Func_param()
			}

		}
		p.SetState(513)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarCOMA {
		{
			p.SetState(514)
			p.Match(VGrammarCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunc_paramContext is an interface to support dynamic dispatch.
type IFunc_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_paramContext differentiates from other interfaces.
	IsFunc_paramContext()
}

type Func_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_paramContext() *Func_paramContext {
	var p = new(Func_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_func_param
	return p
}

func InitEmptyFunc_paramContext(p *Func_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_func_param
}

func (*Func_paramContext) IsFunc_paramContext() {}

func NewFunc_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_paramContext {
	var p = new(Func_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_func_param

	return p
}

func (s *Func_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_paramContext) CopyAll(ctx *Func_paramContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FuncParamContext struct {
	Func_paramContext
}

func NewFuncParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncParamContext {
	var p = new(FuncParamContext)

	InitEmptyFunc_paramContext(&p.Func_paramContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_paramContext))

	return p
}

func (s *FuncParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncParamContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(VGrammarID)
}

func (s *FuncParamContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarID, i)
}

func (s *FuncParamContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncParamContext) RW_INOUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_INOUT, 0)
}

func (s *FuncParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFuncParam(s)
	}
}

func (s *FuncParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFuncParam(s)
	}
}

func (s *FuncParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFuncParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Func_param() (localctx IFunc_paramContext) {
	localctx = NewFunc_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, VGrammarRULE_func_param)
	var _la int

	localctx = NewFuncParamContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(518)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(517)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(520)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarRW_INOUT {
		{
			p.SetState(521)
			p.Match(VGrammarRW_INOUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(524)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStrct_dclContext is an interface to support dynamic dispatch.
type IStrct_dclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStrct_dclContext differentiates from other interfaces.
	IsStrct_dclContext()
}

type Strct_dclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrct_dclContext() *Strct_dclContext {
	var p = new(Strct_dclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_strct_dcl
	return p
}

func InitEmptyStrct_dclContext(p *Strct_dclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_strct_dcl
}

func (*Strct_dclContext) IsStrct_dclContext() {}

func NewStrct_dclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Strct_dclContext {
	var p = new(Strct_dclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_strct_dcl

	return p
}

func (s *Strct_dclContext) GetParser() antlr.Parser { return s.parser }

func (s *Strct_dclContext) CopyAll(ctx *Strct_dclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Strct_dclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Strct_dclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclararStructContext struct {
	Strct_dclContext
}

func NewDeclararStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclararStructContext {
	var p = new(DeclararStructContext)

	InitEmptyStrct_dclContext(&p.Strct_dclContext)
	p.parser = parser
	p.CopyAll(ctx.(*Strct_dclContext))

	return p
}

func (s *DeclararStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclararStructContext) RW_STRUCT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_STRUCT, 0)
}

func (s *DeclararStructContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *DeclararStructContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *DeclararStructContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *DeclararStructContext) AllStruct_prop() []IStruct_propContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStruct_propContext); ok {
			len++
		}
	}

	tst := make([]IStruct_propContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStruct_propContext); ok {
			tst[i] = t.(IStruct_propContext)
			i++
		}
	}

	return tst
}

func (s *DeclararStructContext) Struct_prop(i int) IStruct_propContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_propContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_propContext)
}

func (s *DeclararStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDeclararStruct(s)
	}
}

func (s *DeclararStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDeclararStruct(s)
	}
}

func (s *DeclararStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDeclararStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Strct_dcl() (localctx IStrct_dclContext) {
	localctx = NewStrct_dclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, VGrammarRULE_strct_dcl)
	var _la int

	localctx = NewDeclararStructContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(526)
		p.Match(VGrammarRW_STRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(527)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(528)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(532)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18858832029556224) != 0 {
		{
			p.SetState(529)
			p.Struct_prop()
		}

		p.SetState(534)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(535)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_propContext is an interface to support dynamic dispatch.
type IStruct_propContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_propContext differentiates from other interfaces.
	IsStruct_propContext()
}

type Struct_propContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_propContext() *Struct_propContext {
	var p = new(Struct_propContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_prop
	return p
}

func InitEmptyStruct_propContext(p *Struct_propContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_prop
}

func (*Struct_propContext) IsStruct_propContext() {}

func NewStruct_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_propContext {
	var p = new(Struct_propContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_struct_prop

	return p
}

func (s *Struct_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_propContext) CopyAll(ctx *Struct_propContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructAttrContext struct {
	Struct_propContext
}

func NewStructAttrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructAttrContext {
	var p = new(StructAttrContext)

	InitEmptyStruct_propContext(&p.Struct_propContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_propContext))

	return p
}

func (s *StructAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructAttrContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *StructAttrContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *StructAttrContext) OP_DECLARACION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARACION, 0)
}

func (s *StructAttrContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStructAttr(s)
	}
}

func (s *StructAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStructAttr(s)
	}
}

func (s *StructAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStructAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructFuncContext struct {
	Struct_propContext
}

func NewStructFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructFuncContext {
	var p = new(StructFuncContext)

	InitEmptyStruct_propContext(&p.Struct_propContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_propContext))

	return p
}

func (s *StructFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFuncContext) Func_dcl() IFunc_dclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_dclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_dclContext)
}

func (s *StructFuncContext) RW_MUTATING() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUTATING, 0)
}

func (s *StructFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStructFunc(s)
	}
}

func (s *StructFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStructFunc(s)
	}
}

func (s *StructFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStructFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Struct_prop() (localctx IStruct_propContext) {
	localctx = NewStruct_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, VGrammarRULE_struct_prop)
	var _la int

	p.SetState(547)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VGrammarRW_INT, VGrammarRW_FLOAT64, VGrammarRW_STRING, VGrammarRW_BOOL, VGrammarID, VGrammarCORCHETE_IZQ:
		localctx = NewStructAttrContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(537)
			p.Type_()
		}
		{
			p.SetState(538)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(541)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VGrammarOP_DECLARACION {
			{
				p.SetState(539)
				p.Match(VGrammarOP_DECLARACION)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(540)
				p.expr(0)
			}

		}

	case VGrammarRW_FN, VGrammarRW_MUTATING:
		localctx = NewStructFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(544)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VGrammarRW_MUTATING {
			{
				p.SetState(543)
				p.Match(VGrammarRW_MUTATING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(546)
			p.Func_dcl()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_vectorContext is an interface to support dynamic dispatch.
type IStruct_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_vectorContext differentiates from other interfaces.
	IsStruct_vectorContext()
}

type Struct_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_vectorContext() *Struct_vectorContext {
	var p = new(Struct_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_vector
	return p
}

func InitEmptyStruct_vectorContext(p *Struct_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_vector
}

func (*Struct_vectorContext) IsStruct_vectorContext() {}

func NewStruct_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_vectorContext {
	var p = new(Struct_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_struct_vector

	return p
}

func (s *Struct_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_vectorContext) CopyAll(ctx *Struct_vectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructVectorContext struct {
	Struct_vectorContext
}

func NewStructVectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructVectorContext {
	var p = new(StructVectorContext)

	InitEmptyStruct_vectorContext(&p.Struct_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_vectorContext))

	return p
}

func (s *StructVectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructVectorContext) CORCHETE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_IZQ, 0)
}

func (s *StructVectorContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *StructVectorContext) CORCHETE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarCORCHETE_DER, 0)
}

func (s *StructVectorContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *StructVectorContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *StructVectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStructVector(s)
	}
}

func (s *StructVectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStructVector(s)
	}
}

func (s *StructVectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStructVector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Struct_vector() (localctx IStruct_vectorContext) {
	localctx = NewStruct_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, VGrammarRULE_struct_vector)
	localctx = NewStructVectorContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(549)
		p.Match(VGrammarCORCHETE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(550)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(551)
		p.Match(VGrammarCORCHETE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(552)
		p.Match(VGrammarPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(553)
		p.Match(VGrammarPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_initContext is an interface to support dynamic dispatch.
type IStruct_initContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_initContext differentiates from other interfaces.
	IsStruct_initContext()
}

type Struct_initContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_initContext() *Struct_initContext {
	var p = new(Struct_initContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_init
	return p
}

func InitEmptyStruct_initContext(p *Struct_initContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_init
}

func (*Struct_initContext) IsStruct_initContext() {}

func NewStruct_initContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_initContext {
	var p = new(Struct_initContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_struct_init

	return p
}

func (s *Struct_initContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_initContext) CopyAll(ctx *Struct_initContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_initContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_initContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructInitContext struct {
	Struct_initContext
}

func NewStructInitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructInitContext {
	var p = new(StructInitContext)

	InitEmptyStruct_initContext(&p.Struct_initContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_initContext))

	return p
}

func (s *StructInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInitContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *StructInitContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *StructInitContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *StructInitContext) Struct_init_list() IStruct_init_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_init_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_init_listContext)
}

func (s *StructInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStructInit(s)
	}
}

func (s *StructInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStructInit(s)
	}
}

func (s *StructInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStructInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Struct_init() (localctx IStruct_initContext) {
	localctx = NewStruct_initContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, VGrammarRULE_struct_init)
	var _la int

	localctx = NewStructInitContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(555)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(556)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarID {
		{
			p.SetState(557)
			p.Struct_init_list()
		}

	}
	{
		p.SetState(560)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_init_listContext is an interface to support dynamic dispatch.
type IStruct_init_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_init_listContext differentiates from other interfaces.
	IsStruct_init_listContext()
}

type Struct_init_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_init_listContext() *Struct_init_listContext {
	var p = new(Struct_init_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_init_list
	return p
}

func InitEmptyStruct_init_listContext(p *Struct_init_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_init_list
}

func (*Struct_init_listContext) IsStruct_init_listContext() {}

func NewStruct_init_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_init_listContext {
	var p = new(Struct_init_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_struct_init_list

	return p
}

func (s *Struct_init_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_init_listContext) CopyAll(ctx *Struct_init_listContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_init_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_init_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructInitListContext struct {
	Struct_init_listContext
}

func NewStructInitListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructInitListContext {
	var p = new(StructInitListContext)

	InitEmptyStruct_init_listContext(&p.Struct_init_listContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_init_listContext))

	return p
}

func (s *StructInitListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInitListContext) AllStruct_init_field() []IStruct_init_fieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStruct_init_fieldContext); ok {
			len++
		}
	}

	tst := make([]IStruct_init_fieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStruct_init_fieldContext); ok {
			tst[i] = t.(IStruct_init_fieldContext)
			i++
		}
	}

	return tst
}

func (s *StructInitListContext) Struct_init_field(i int) IStruct_init_fieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_init_fieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_init_fieldContext)
}

func (s *StructInitListContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *StructInitListContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *StructInitListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStructInitList(s)
	}
}

func (s *StructInitListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStructInitList(s)
	}
}

func (s *StructInitListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStructInitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Struct_init_list() (localctx IStruct_init_listContext) {
	localctx = NewStruct_init_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, VGrammarRULE_struct_init_list)
	var _la int

	var _alt int

	localctx = NewStructInitListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(562)
		p.Struct_init_field()
	}
	p.SetState(567)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(563)
				p.Match(VGrammarCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(564)
				p.Struct_init_field()
			}

		}
		p.SetState(569)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(571)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarCOMA {
		{
			p.SetState(570)
			p.Match(VGrammarCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_method_callContext is an interface to support dynamic dispatch.
type IStruct_method_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_method_callContext differentiates from other interfaces.
	IsStruct_method_callContext()
}

type Struct_method_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_method_callContext() *Struct_method_callContext {
	var p = new(Struct_method_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_method_call
	return p
}

func InitEmptyStruct_method_callContext(p *Struct_method_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_method_call
}

func (*Struct_method_callContext) IsStruct_method_callContext() {}

func NewStruct_method_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_method_callContext {
	var p = new(Struct_method_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_struct_method_call

	return p
}

func (s *Struct_method_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_method_callContext) CopyAll(ctx *Struct_method_callContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_method_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_method_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructMethodCallContext struct {
	Struct_method_callContext
}

func NewStructMethodCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructMethodCallContext {
	var p = new(StructMethodCallContext)

	InitEmptyStruct_method_callContext(&p.Struct_method_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_method_callContext))

	return p
}

func (s *StructMethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMethodCallContext) Id_pattern() IId_patternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_patternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_patternContext)
}

func (s *StructMethodCallContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO, 0)
}

func (s *StructMethodCallContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *StructMethodCallContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *StructMethodCallContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *StructMethodCallContext) Arg_list() IArg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArg_listContext)
}

func (s *StructMethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStructMethodCall(s)
	}
}

func (s *StructMethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStructMethodCall(s)
	}
}

func (s *StructMethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStructMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Struct_method_call() (localctx IStruct_method_callContext) {
	localctx = NewStruct_method_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, VGrammarRULE_struct_method_call)
	var _la int

	localctx = NewStructMethodCallContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(573)
		p.Id_pattern()
	}
	{
		p.SetState(574)
		p.Match(VGrammarPUNTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(575)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(576)
		p.Match(VGrammarPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(578)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4635892868572766208) != 0 {
		{
			p.SetState(577)
			p.Arg_list()
		}

	}
	{
		p.SetState(580)
		p.Match(VGrammarPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStruct_init_fieldContext is an interface to support dynamic dispatch.
type IStruct_init_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_init_fieldContext differentiates from other interfaces.
	IsStruct_init_fieldContext()
}

type Struct_init_fieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_init_fieldContext() *Struct_init_fieldContext {
	var p = new(Struct_init_fieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_init_field
	return p
}

func InitEmptyStruct_init_fieldContext(p *Struct_init_fieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_init_field
}

func (*Struct_init_fieldContext) IsStruct_init_fieldContext() {}

func NewStruct_init_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_init_fieldContext {
	var p = new(Struct_init_fieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_struct_init_field

	return p
}

func (s *Struct_init_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_init_fieldContext) CopyAll(ctx *Struct_init_fieldContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_init_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_init_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructInitFieldContext struct {
	Struct_init_fieldContext
}

func NewStructInitFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructInitFieldContext {
	var p = new(StructInitFieldContext)

	InitEmptyStruct_init_fieldContext(&p.Struct_init_fieldContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_init_fieldContext))

	return p
}

func (s *StructInitFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructInitFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *StructInitFieldContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, 0)
}

func (s *StructInitFieldContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructInitFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStructInitField(s)
	}
}

func (s *StructInitFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStructInitField(s)
	}
}

func (s *StructInitFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStructInitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Struct_init_field() (localctx IStruct_init_fieldContext) {
	localctx = NewStruct_init_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, VGrammarRULE_struct_init_field)
	localctx = NewStructInitFieldContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(582)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(583)
		p.Match(VGrammarDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(584)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *VGrammar) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 17:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VGrammar) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
