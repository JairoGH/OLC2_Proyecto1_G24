// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VGrammar
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
		"", "'main'", "'fn'", "'mut'", "'struct'", "'if'", "'else'", "'switch'",
		"'case'", "'default'", "'for'", "'in'", "'break'", "'continue'", "'return'",
		"'true'", "'false'", "'nil'", "'int'", "'float64'", "'string'", "'bool'",
		"'rune'", "'print'", "'println'", "'Atoi'", "'parseFloat'", "'typeof'",
		"'append'", "'len'", "'join'", "'indexOf'", "'+'", "'-'", "'*'", "'/'",
		"'%'", "'='", "'+='", "'-='", "'=='", "'!='", "'<'", "'>'", "'<='",
		"'>='", "'&&'", "'||'", "'!'", "':='", "'.'", "','", "';'", "':'", "'('",
		"')'", "'{'", "'}'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "RW_MAIN", "RW_FN", "RW_MUT", "RW_STRUCT", "RW_IF", "RW_ELSE", "RW_SWITCH",
		"RW_CASE", "RW_DEFAULT", "RW_FOR", "RW_IN", "RW_BREAK", "RW_CONTINUE",
		"RW_RETURN", "RW_TRUE", "RW_FALSE", "RW_NIL", "RW_INT", "RW_FLOAT64",
		"RW_STRING", "RW_BOOL", "RW_RUNE", "RW_PRINT", "RW_PRINTLN", "RW_ATOI",
		"RW_PARSEFLOAT", "RW_TYPEOF", "RW_APPEND", "RW_LEN", "RW_JOIN", "RW_INDEXOF",
		"OP_SUMA", "OP_RESTA", "OP_MULT", "OP_DIV", "OP_MOD", "OP_ASSIGN", "OP_ADD_ASSIGN",
		"OP_SUB_ASSIGN", "OP_IGUAL", "OP_DIFERENTE", "OP_MENORQ", "OP_MAYORQ",
		"OP_MENORIGUAL", "OP_MAYORIGUAL", "OP_AND", "OP_OR", "OP_NOT", "OP_DECLARATION",
		"PUNTO", "COMA", "PUNTO_Y_COMA", "DOS_PUNTOS", "PAR_IZQ", "PAR_DER",
		"LLAVE_IZQ", "LLAVE_DER", "CORCHETE_IZQ", "CORCHETE_DER", "INT_LITERAL",
		"FLOAT_LITERAL", "STRING_LITERAL", "RUNE_LITERAL", "ID", "WS", "COMMENT",
		"MULTILINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"programa", "stmt", "decl_stmt", "assign_stmt", "expr", "literal", "flow_stmt",
		"func_call", "if_stmt", "switch_stmt", "switch_case", "default_case",
		"for_stmt", "func_decl", "param", "struct_decl", "struct_field",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 67, 339, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 5, 0, 36, 8, 0, 10, 0, 12, 0, 39, 9, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 3, 1, 45, 8, 1, 1, 1, 1, 1, 3, 1, 49, 8, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 54, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 61, 8, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 82, 8, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 93, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 104, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4,
		145, 8, 4, 10, 4, 12, 4, 148, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 3, 5, 157, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 163, 8, 6, 3, 6, 165,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 172, 8, 7, 10, 7, 12, 7, 175,
		9, 7, 3, 7, 177, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 185, 8,
		7, 10, 7, 12, 7, 188, 9, 7, 3, 7, 190, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 5, 7, 198, 8, 7, 10, 7, 12, 7, 201, 9, 7, 3, 7, 203, 8, 7, 1,
		7, 3, 7, 206, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 212, 8, 8, 10, 8, 12,
		8, 215, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 221, 8, 8, 10, 8, 12, 8, 224,
		9, 8, 1, 8, 3, 8, 227, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 233, 8, 9, 10,
		9, 12, 9, 236, 9, 9, 1, 9, 3, 9, 239, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 5, 10, 247, 8, 10, 10, 10, 12, 10, 250, 9, 10, 1, 11, 1, 11,
		1, 11, 5, 11, 255, 8, 11, 10, 11, 12, 11, 258, 9, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 5, 12, 266, 8, 12, 10, 12, 12, 12, 269, 9, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 279, 8,
		13, 10, 13, 12, 13, 282, 9, 13, 3, 13, 284, 8, 13, 1, 13, 1, 13, 3, 13,
		288, 8, 13, 1, 13, 1, 13, 5, 13, 292, 8, 13, 10, 13, 12, 13, 295, 9, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 3, 14, 309, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 315, 8,
		15, 10, 15, 12, 15, 318, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 3, 16, 337, 8, 16, 1, 16, 0, 1, 8, 17, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 0, 2, 2, 0, 1, 1, 64, 64, 1, 0, 18, 22,
		394, 0, 37, 1, 0, 0, 0, 2, 60, 1, 0, 0, 0, 4, 81, 1, 0, 0, 0, 6, 92, 1,
		0, 0, 0, 8, 103, 1, 0, 0, 0, 10, 156, 1, 0, 0, 0, 12, 164, 1, 0, 0, 0,
		14, 205, 1, 0, 0, 0, 16, 207, 1, 0, 0, 0, 18, 228, 1, 0, 0, 0, 20, 242,
		1, 0, 0, 0, 22, 251, 1, 0, 0, 0, 24, 259, 1, 0, 0, 0, 26, 272, 1, 0, 0,
		0, 28, 308, 1, 0, 0, 0, 30, 310, 1, 0, 0, 0, 32, 336, 1, 0, 0, 0, 34, 36,
		3, 2, 1, 0, 35, 34, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0,
		37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 40, 41, 5,
		0, 0, 1, 41, 1, 1, 0, 0, 0, 42, 44, 3, 4, 2, 0, 43, 45, 5, 52, 0, 0, 44,
		43, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 61, 1, 0, 0, 0, 46, 48, 3, 6, 3,
		0, 47, 49, 5, 52, 0, 0, 48, 47, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 61,
		1, 0, 0, 0, 50, 61, 3, 12, 6, 0, 51, 53, 3, 14, 7, 0, 52, 54, 5, 52, 0,
		0, 53, 52, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 61, 1, 0, 0, 0, 55, 61,
		3, 16, 8, 0, 56, 61, 3, 18, 9, 0, 57, 61, 3, 24, 12, 0, 58, 61, 3, 26,
		13, 0, 59, 61, 3, 30, 15, 0, 60, 42, 1, 0, 0, 0, 60, 46, 1, 0, 0, 0, 60,
		50, 1, 0, 0, 0, 60, 51, 1, 0, 0, 0, 60, 55, 1, 0, 0, 0, 60, 56, 1, 0, 0,
		0, 60, 57, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 59, 1, 0, 0, 0, 61, 3, 1,
		0, 0, 0, 62, 63, 5, 3, 0, 0, 63, 64, 5, 64, 0, 0, 64, 65, 5, 49, 0, 0,
		65, 82, 3, 8, 4, 0, 66, 67, 5, 3, 0, 0, 67, 68, 5, 64, 0, 0, 68, 82, 5,
		18, 0, 0, 69, 70, 5, 3, 0, 0, 70, 71, 5, 64, 0, 0, 71, 82, 5, 20, 0, 0,
		72, 73, 5, 3, 0, 0, 73, 74, 5, 64, 0, 0, 74, 82, 5, 19, 0, 0, 75, 76, 5,
		3, 0, 0, 76, 77, 5, 64, 0, 0, 77, 82, 5, 21, 0, 0, 78, 79, 5, 3, 0, 0,
		79, 80, 5, 64, 0, 0, 80, 82, 5, 22, 0, 0, 81, 62, 1, 0, 0, 0, 81, 66, 1,
		0, 0, 0, 81, 69, 1, 0, 0, 0, 81, 72, 1, 0, 0, 0, 81, 75, 1, 0, 0, 0, 81,
		78, 1, 0, 0, 0, 82, 5, 1, 0, 0, 0, 83, 84, 5, 64, 0, 0, 84, 85, 5, 37,
		0, 0, 85, 93, 3, 8, 4, 0, 86, 87, 5, 64, 0, 0, 87, 88, 5, 38, 0, 0, 88,
		93, 3, 8, 4, 0, 89, 90, 5, 64, 0, 0, 90, 91, 5, 39, 0, 0, 91, 93, 3, 8,
		4, 0, 92, 83, 1, 0, 0, 0, 92, 86, 1, 0, 0, 0, 92, 89, 1, 0, 0, 0, 93, 7,
		1, 0, 0, 0, 94, 95, 6, 4, -1, 0, 95, 104, 3, 10, 5, 0, 96, 104, 5, 64,
		0, 0, 97, 98, 5, 48, 0, 0, 98, 104, 3, 8, 4, 2, 99, 100, 5, 54, 0, 0, 100,
		101, 3, 8, 4, 0, 101, 102, 5, 55, 0, 0, 102, 104, 1, 0, 0, 0, 103, 94,
		1, 0, 0, 0, 103, 96, 1, 0, 0, 0, 103, 97, 1, 0, 0, 0, 103, 99, 1, 0, 0,
		0, 104, 146, 1, 0, 0, 0, 105, 106, 10, 15, 0, 0, 106, 107, 5, 32, 0, 0,
		107, 145, 3, 8, 4, 16, 108, 109, 10, 14, 0, 0, 109, 110, 5, 33, 0, 0, 110,
		145, 3, 8, 4, 15, 111, 112, 10, 13, 0, 0, 112, 113, 5, 34, 0, 0, 113, 145,
		3, 8, 4, 14, 114, 115, 10, 12, 0, 0, 115, 116, 5, 35, 0, 0, 116, 145, 3,
		8, 4, 13, 117, 118, 10, 11, 0, 0, 118, 119, 5, 36, 0, 0, 119, 145, 3, 8,
		4, 12, 120, 121, 10, 10, 0, 0, 121, 122, 5, 40, 0, 0, 122, 145, 3, 8, 4,
		11, 123, 124, 10, 9, 0, 0, 124, 125, 5, 41, 0, 0, 125, 145, 3, 8, 4, 10,
		126, 127, 10, 8, 0, 0, 127, 128, 5, 42, 0, 0, 128, 145, 3, 8, 4, 9, 129,
		130, 10, 7, 0, 0, 130, 131, 5, 43, 0, 0, 131, 145, 3, 8, 4, 8, 132, 133,
		10, 6, 0, 0, 133, 134, 5, 44, 0, 0, 134, 145, 3, 8, 4, 7, 135, 136, 10,
		5, 0, 0, 136, 137, 5, 45, 0, 0, 137, 145, 3, 8, 4, 6, 138, 139, 10, 4,
		0, 0, 139, 140, 5, 46, 0, 0, 140, 145, 3, 8, 4, 5, 141, 142, 10, 3, 0,
		0, 142, 143, 5, 47, 0, 0, 143, 145, 3, 8, 4, 4, 144, 105, 1, 0, 0, 0, 144,
		108, 1, 0, 0, 0, 144, 111, 1, 0, 0, 0, 144, 114, 1, 0, 0, 0, 144, 117,
		1, 0, 0, 0, 144, 120, 1, 0, 0, 0, 144, 123, 1, 0, 0, 0, 144, 126, 1, 0,
		0, 0, 144, 129, 1, 0, 0, 0, 144, 132, 1, 0, 0, 0, 144, 135, 1, 0, 0, 0,
		144, 138, 1, 0, 0, 0, 144, 141, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146,
		144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 9, 1, 0, 0, 0, 148, 146, 1,
		0, 0, 0, 149, 157, 5, 60, 0, 0, 150, 157, 5, 61, 0, 0, 151, 157, 5, 62,
		0, 0, 152, 157, 5, 63, 0, 0, 153, 157, 5, 15, 0, 0, 154, 157, 5, 16, 0,
		0, 155, 157, 5, 17, 0, 0, 156, 149, 1, 0, 0, 0, 156, 150, 1, 0, 0, 0, 156,
		151, 1, 0, 0, 0, 156, 152, 1, 0, 0, 0, 156, 153, 1, 0, 0, 0, 156, 154,
		1, 0, 0, 0, 156, 155, 1, 0, 0, 0, 157, 11, 1, 0, 0, 0, 158, 165, 5, 12,
		0, 0, 159, 165, 5, 13, 0, 0, 160, 162, 5, 14, 0, 0, 161, 163, 3, 8, 4,
		0, 162, 161, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 165, 1, 0, 0, 0, 164,
		158, 1, 0, 0, 0, 164, 159, 1, 0, 0, 0, 164, 160, 1, 0, 0, 0, 165, 13, 1,
		0, 0, 0, 166, 167, 5, 64, 0, 0, 167, 176, 5, 54, 0, 0, 168, 173, 3, 8,
		4, 0, 169, 170, 5, 51, 0, 0, 170, 172, 3, 8, 4, 0, 171, 169, 1, 0, 0, 0,
		172, 175, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174,
		177, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 168, 1, 0, 0, 0, 176, 177,
		1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 206, 5, 55, 0, 0, 179, 180, 5, 23,
		0, 0, 180, 189, 5, 54, 0, 0, 181, 186, 3, 8, 4, 0, 182, 183, 5, 51, 0,
		0, 183, 185, 3, 8, 4, 0, 184, 182, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186,
		184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 190, 1, 0, 0, 0, 188, 186,
		1, 0, 0, 0, 189, 181, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 191, 1, 0,
		0, 0, 191, 206, 5, 55, 0, 0, 192, 193, 5, 24, 0, 0, 193, 202, 5, 54, 0,
		0, 194, 199, 3, 8, 4, 0, 195, 196, 5, 51, 0, 0, 196, 198, 3, 8, 4, 0, 197,
		195, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200,
		1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 194, 1, 0,
		0, 0, 202, 203, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 206, 5, 55, 0, 0,
		205, 166, 1, 0, 0, 0, 205, 179, 1, 0, 0, 0, 205, 192, 1, 0, 0, 0, 206,
		15, 1, 0, 0, 0, 207, 208, 5, 5, 0, 0, 208, 209, 3, 8, 4, 0, 209, 213, 5,
		56, 0, 0, 210, 212, 3, 2, 1, 0, 211, 210, 1, 0, 0, 0, 212, 215, 1, 0, 0,
		0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 216, 1, 0, 0, 0, 215,
		213, 1, 0, 0, 0, 216, 226, 5, 57, 0, 0, 217, 218, 5, 6, 0, 0, 218, 222,
		5, 56, 0, 0, 219, 221, 3, 2, 1, 0, 220, 219, 1, 0, 0, 0, 221, 224, 1, 0,
		0, 0, 222, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 225, 1, 0, 0, 0,
		224, 222, 1, 0, 0, 0, 225, 227, 5, 57, 0, 0, 226, 217, 1, 0, 0, 0, 226,
		227, 1, 0, 0, 0, 227, 17, 1, 0, 0, 0, 228, 229, 5, 7, 0, 0, 229, 230, 3,
		8, 4, 0, 230, 234, 5, 56, 0, 0, 231, 233, 3, 20, 10, 0, 232, 231, 1, 0,
		0, 0, 233, 236, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0,
		235, 238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 237, 239, 3, 22, 11, 0, 238,
		237, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 241,
		5, 57, 0, 0, 241, 19, 1, 0, 0, 0, 242, 243, 5, 8, 0, 0, 243, 244, 3, 8,
		4, 0, 244, 248, 5, 53, 0, 0, 245, 247, 3, 2, 1, 0, 246, 245, 1, 0, 0, 0,
		247, 250, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249,
		21, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 252, 5, 9, 0, 0, 252, 256, 5,
		53, 0, 0, 253, 255, 3, 2, 1, 0, 254, 253, 1, 0, 0, 0, 255, 258, 1, 0, 0,
		0, 256, 254, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 23, 1, 0, 0, 0, 258,
		256, 1, 0, 0, 0, 259, 260, 5, 10, 0, 0, 260, 261, 5, 64, 0, 0, 261, 262,
		5, 11, 0, 0, 262, 263, 3, 8, 4, 0, 263, 267, 5, 56, 0, 0, 264, 266, 3,
		2, 1, 0, 265, 264, 1, 0, 0, 0, 266, 269, 1, 0, 0, 0, 267, 265, 1, 0, 0,
		0, 267, 268, 1, 0, 0, 0, 268, 270, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 270,
		271, 5, 57, 0, 0, 271, 25, 1, 0, 0, 0, 272, 273, 5, 2, 0, 0, 273, 274,
		7, 0, 0, 0, 274, 283, 5, 54, 0, 0, 275, 280, 3, 28, 14, 0, 276, 277, 5,
		51, 0, 0, 277, 279, 3, 28, 14, 0, 278, 276, 1, 0, 0, 0, 279, 282, 1, 0,
		0, 0, 280, 278, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 284, 1, 0, 0, 0,
		282, 280, 1, 0, 0, 0, 283, 275, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284,
		285, 1, 0, 0, 0, 285, 287, 5, 55, 0, 0, 286, 288, 7, 1, 0, 0, 287, 286,
		1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 293, 5, 56,
		0, 0, 290, 292, 3, 2, 1, 0, 291, 290, 1, 0, 0, 0, 292, 295, 1, 0, 0, 0,
		293, 291, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 296, 1, 0, 0, 0, 295,
		293, 1, 0, 0, 0, 296, 297, 5, 57, 0, 0, 297, 27, 1, 0, 0, 0, 298, 299,
		5, 64, 0, 0, 299, 309, 5, 18, 0, 0, 300, 301, 5, 64, 0, 0, 301, 309, 5,
		20, 0, 0, 302, 303, 5, 64, 0, 0, 303, 309, 5, 19, 0, 0, 304, 305, 5, 64,
		0, 0, 305, 309, 5, 21, 0, 0, 306, 307, 5, 64, 0, 0, 307, 309, 5, 22, 0,
		0, 308, 298, 1, 0, 0, 0, 308, 300, 1, 0, 0, 0, 308, 302, 1, 0, 0, 0, 308,
		304, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 309, 29, 1, 0, 0, 0, 310, 311, 5,
		4, 0, 0, 311, 312, 5, 64, 0, 0, 312, 316, 5, 56, 0, 0, 313, 315, 3, 32,
		16, 0, 314, 313, 1, 0, 0, 0, 315, 318, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0,
		316, 317, 1, 0, 0, 0, 317, 319, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 319,
		320, 5, 57, 0, 0, 320, 31, 1, 0, 0, 0, 321, 322, 5, 3, 0, 0, 322, 323,
		5, 64, 0, 0, 323, 337, 5, 18, 0, 0, 324, 325, 5, 3, 0, 0, 325, 326, 5,
		64, 0, 0, 326, 337, 5, 20, 0, 0, 327, 328, 5, 3, 0, 0, 328, 329, 5, 64,
		0, 0, 329, 337, 5, 19, 0, 0, 330, 331, 5, 3, 0, 0, 331, 332, 5, 64, 0,
		0, 332, 337, 5, 21, 0, 0, 333, 334, 5, 3, 0, 0, 334, 335, 5, 64, 0, 0,
		335, 337, 5, 22, 0, 0, 336, 321, 1, 0, 0, 0, 336, 324, 1, 0, 0, 0, 336,
		327, 1, 0, 0, 0, 336, 330, 1, 0, 0, 0, 336, 333, 1, 0, 0, 0, 337, 33, 1,
		0, 0, 0, 35, 37, 44, 48, 53, 60, 81, 92, 103, 144, 146, 156, 162, 164,
		173, 176, 186, 189, 199, 202, 205, 213, 222, 226, 234, 238, 248, 256, 267,
		280, 283, 287, 293, 308, 316, 336,
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
	VGrammarEOF               = antlr.TokenEOF
	VGrammarRW_MAIN           = 1
	VGrammarRW_FN             = 2
	VGrammarRW_MUT            = 3
	VGrammarRW_STRUCT         = 4
	VGrammarRW_IF             = 5
	VGrammarRW_ELSE           = 6
	VGrammarRW_SWITCH         = 7
	VGrammarRW_CASE           = 8
	VGrammarRW_DEFAULT        = 9
	VGrammarRW_FOR            = 10
	VGrammarRW_IN             = 11
	VGrammarRW_BREAK          = 12
	VGrammarRW_CONTINUE       = 13
	VGrammarRW_RETURN         = 14
	VGrammarRW_TRUE           = 15
	VGrammarRW_FALSE          = 16
	VGrammarRW_NIL            = 17
	VGrammarRW_INT            = 18
	VGrammarRW_FLOAT64        = 19
	VGrammarRW_STRING         = 20
	VGrammarRW_BOOL           = 21
	VGrammarRW_RUNE           = 22
	VGrammarRW_PRINT          = 23
	VGrammarRW_PRINTLN        = 24
	VGrammarRW_ATOI           = 25
	VGrammarRW_PARSEFLOAT     = 26
	VGrammarRW_TYPEOF         = 27
	VGrammarRW_APPEND         = 28
	VGrammarRW_LEN            = 29
	VGrammarRW_JOIN           = 30
	VGrammarRW_INDEXOF        = 31
	VGrammarOP_SUMA           = 32
	VGrammarOP_RESTA          = 33
	VGrammarOP_MULT           = 34
	VGrammarOP_DIV            = 35
	VGrammarOP_MOD            = 36
	VGrammarOP_ASSIGN         = 37
	VGrammarOP_ADD_ASSIGN     = 38
	VGrammarOP_SUB_ASSIGN     = 39
	VGrammarOP_IGUAL          = 40
	VGrammarOP_DIFERENTE      = 41
	VGrammarOP_MENORQ         = 42
	VGrammarOP_MAYORQ         = 43
	VGrammarOP_MENORIGUAL     = 44
	VGrammarOP_MAYORIGUAL     = 45
	VGrammarOP_AND            = 46
	VGrammarOP_OR             = 47
	VGrammarOP_NOT            = 48
	VGrammarOP_DECLARATION    = 49
	VGrammarPUNTO             = 50
	VGrammarCOMA              = 51
	VGrammarPUNTO_Y_COMA      = 52
	VGrammarDOS_PUNTOS        = 53
	VGrammarPAR_IZQ           = 54
	VGrammarPAR_DER           = 55
	VGrammarLLAVE_IZQ         = 56
	VGrammarLLAVE_DER         = 57
	VGrammarCORCHETE_IZQ      = 58
	VGrammarCORCHETE_DER      = 59
	VGrammarINT_LITERAL       = 60
	VGrammarFLOAT_LITERAL     = 61
	VGrammarSTRING_LITERAL    = 62
	VGrammarRUNE_LITERAL      = 63
	VGrammarID                = 64
	VGrammarWS                = 65
	VGrammarCOMMENT           = 66
	VGrammarMULTILINE_COMMENT = 67
)

// VGrammar rules.
const (
	VGrammarRULE_programa     = 0
	VGrammarRULE_stmt         = 1
	VGrammarRULE_decl_stmt    = 2
	VGrammarRULE_assign_stmt  = 3
	VGrammarRULE_expr         = 4
	VGrammarRULE_literal      = 5
	VGrammarRULE_flow_stmt    = 6
	VGrammarRULE_func_call    = 7
	VGrammarRULE_if_stmt      = 8
	VGrammarRULE_switch_stmt  = 9
	VGrammarRULE_switch_case  = 10
	VGrammarRULE_default_case = 11
	VGrammarRULE_for_stmt     = 12
	VGrammarRULE_func_decl    = 13
	VGrammarRULE_param        = 14
	VGrammarRULE_struct_decl  = 15
	VGrammarRULE_struct_field = 16
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) EOF() antlr.TerminalNode {
	return s.GetToken(VGrammarEOF, 0)
}

func (s *ProgramaContext) AllStmt() []IStmtContext {
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

func (s *ProgramaContext) Stmt(i int) IStmtContext {
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

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (s *ProgramaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitPrograma(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VGrammarRULE_programa)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&4611686018433686831) != 0 {
		{
			p.SetState(34)
			p.Stmt()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Match(VGrammarEOF)
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

func (s *StmtContext) CopyAll(ctx *StmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StmtFuncDeclContext struct {
	StmtContext
}

func NewStmtFuncDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtFuncDeclContext {
	var p = new(StmtFuncDeclContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtFuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtFuncDeclContext) Func_decl() IFunc_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_declContext)
}

func (s *StmtFuncDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmtFuncDecl(s)
	}
}

func (s *StmtFuncDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmtFuncDecl(s)
	}
}

func (s *StmtFuncDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmtFuncDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtDeclContext struct {
	StmtContext
}

func NewStmtDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtDeclContext {
	var p = new(StmtDeclContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtDeclContext) Decl_stmt() IDecl_stmtContext {
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

func (s *StmtDeclContext) PUNTO_Y_COMA() antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO_Y_COMA, 0)
}

func (s *StmtDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmtDecl(s)
	}
}

func (s *StmtDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmtDecl(s)
	}
}

func (s *StmtDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmtDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtStructDeclContext struct {
	StmtContext
}

func NewStmtStructDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtStructDeclContext {
	var p = new(StmtStructDeclContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtStructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtStructDeclContext) Struct_decl() IStruct_declContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_declContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_declContext)
}

func (s *StmtStructDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmtStructDecl(s)
	}
}

func (s *StmtStructDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmtStructDecl(s)
	}
}

func (s *StmtStructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmtStructDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtSwitchContext struct {
	StmtContext
}

func NewStmtSwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtSwitchContext {
	var p = new(StmtSwitchContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtSwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtSwitchContext) Switch_stmt() ISwitch_stmtContext {
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

func (s *StmtSwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmtSwitch(s)
	}
}

func (s *StmtSwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmtSwitch(s)
	}
}

func (s *StmtSwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmtSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtForContext struct {
	StmtContext
}

func NewStmtForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtForContext {
	var p = new(StmtForContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtForContext) For_stmt() IFor_stmtContext {
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

func (s *StmtForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmtFor(s)
	}
}

func (s *StmtForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmtFor(s)
	}
}

func (s *StmtForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmtFor(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtAssignContext struct {
	StmtContext
}

func NewStmtAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtAssignContext {
	var p = new(StmtAssignContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtAssignContext) Assign_stmt() IAssign_stmtContext {
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

func (s *StmtAssignContext) PUNTO_Y_COMA() antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO_Y_COMA, 0)
}

func (s *StmtAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmtAssign(s)
	}
}

func (s *StmtAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmtAssign(s)
	}
}

func (s *StmtAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmtAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtIfContext struct {
	StmtContext
}

func NewStmtIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtIfContext {
	var p = new(StmtIfContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtIfContext) If_stmt() IIf_stmtContext {
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

func (s *StmtIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmtIf(s)
	}
}

func (s *StmtIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmtIf(s)
	}
}

func (s *StmtIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmtIf(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtFlowContext struct {
	StmtContext
}

func NewStmtFlowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtFlowContext {
	var p = new(StmtFlowContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtFlowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtFlowContext) Flow_stmt() IFlow_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFlow_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFlow_stmtContext)
}

func (s *StmtFlowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmtFlow(s)
	}
}

func (s *StmtFlowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmtFlow(s)
	}
}

func (s *StmtFlowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmtFlow(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtFuncCallContext struct {
	StmtContext
}

func NewStmtFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtFuncCallContext {
	var p = new(StmtFuncCallContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtFuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtFuncCallContext) Func_call() IFunc_callContext {
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

func (s *StmtFuncCallContext) PUNTO_Y_COMA() antlr.TerminalNode {
	return s.GetToken(VGrammarPUNTO_Y_COMA, 0)
}

func (s *StmtFuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmtFuncCall(s)
	}
}

func (s *StmtFuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmtFuncCall(s)
	}
}

func (s *StmtFuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmtFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VGrammarRULE_stmt)
	var _la int

	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStmtDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Decl_stmt()
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VGrammarPUNTO_Y_COMA {
			{
				p.SetState(43)
				p.Match(VGrammarPUNTO_Y_COMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		localctx = NewStmtAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)
			p.Assign_stmt()
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VGrammarPUNTO_Y_COMA {
			{
				p.SetState(47)
				p.Match(VGrammarPUNTO_Y_COMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		localctx = NewStmtFlowContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Flow_stmt()
		}

	case 4:
		localctx = NewStmtFuncCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(51)
			p.Func_call()
		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == VGrammarPUNTO_Y_COMA {
			{
				p.SetState(52)
				p.Match(VGrammarPUNTO_Y_COMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 5:
		localctx = NewStmtIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(55)
			p.If_stmt()
		}

	case 6:
		localctx = NewStmtSwitchContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(56)
			p.Switch_stmt()
		}

	case 7:
		localctx = NewStmtForContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(57)
			p.For_stmt()
		}

	case 8:
		localctx = NewStmtFuncDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(58)
			p.Func_decl()
		}

	case 9:
		localctx = NewStmtStructDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(59)
			p.Struct_decl()
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

type TypedDeclBoolContext struct {
	Decl_stmtContext
}

func NewTypedDeclBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedDeclBoolContext {
	var p = new(TypedDeclBoolContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *TypedDeclBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedDeclBoolContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *TypedDeclBoolContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *TypedDeclBoolContext) RW_BOOL() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_BOOL, 0)
}

func (s *TypedDeclBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterTypedDeclBool(s)
	}
}

func (s *TypedDeclBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitTypedDeclBool(s)
	}
}

func (s *TypedDeclBoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitTypedDeclBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypedDeclFloatContext struct {
	Decl_stmtContext
}

func NewTypedDeclFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedDeclFloatContext {
	var p = new(TypedDeclFloatContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *TypedDeclFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedDeclFloatContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *TypedDeclFloatContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *TypedDeclFloatContext) RW_FLOAT64() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FLOAT64, 0)
}

func (s *TypedDeclFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterTypedDeclFloat(s)
	}
}

func (s *TypedDeclFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitTypedDeclFloat(s)
	}
}

func (s *TypedDeclFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitTypedDeclFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypedDeclStringContext struct {
	Decl_stmtContext
}

func NewTypedDeclStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedDeclStringContext {
	var p = new(TypedDeclStringContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *TypedDeclStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedDeclStringContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *TypedDeclStringContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *TypedDeclStringContext) RW_STRING() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_STRING, 0)
}

func (s *TypedDeclStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterTypedDeclString(s)
	}
}

func (s *TypedDeclStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitTypedDeclString(s)
	}
}

func (s *TypedDeclStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitTypedDeclString(s)

	default:
		return t.VisitChildren(s)
	}
}

type InferDeclContext struct {
	Decl_stmtContext
}

func NewInferDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InferDeclContext {
	var p = new(InferDeclContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *InferDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InferDeclContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *InferDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *InferDeclContext) OP_DECLARATION() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DECLARATION, 0)
}

func (s *InferDeclContext) Expr() IExprContext {
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

func (s *InferDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterInferDecl(s)
	}
}

func (s *InferDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitInferDecl(s)
	}
}

func (s *InferDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitInferDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypedDeclRuneContext struct {
	Decl_stmtContext
}

func NewTypedDeclRuneContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedDeclRuneContext {
	var p = new(TypedDeclRuneContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *TypedDeclRuneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedDeclRuneContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *TypedDeclRuneContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *TypedDeclRuneContext) RW_RUNE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_RUNE, 0)
}

func (s *TypedDeclRuneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterTypedDeclRune(s)
	}
}

func (s *TypedDeclRuneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitTypedDeclRune(s)
	}
}

func (s *TypedDeclRuneContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitTypedDeclRune(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypedDeclIntContext struct {
	Decl_stmtContext
}

func NewTypedDeclIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedDeclIntContext {
	var p = new(TypedDeclIntContext)

	InitEmptyDecl_stmtContext(&p.Decl_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Decl_stmtContext))

	return p
}

func (s *TypedDeclIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedDeclIntContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *TypedDeclIntContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *TypedDeclIntContext) RW_INT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_INT, 0)
}

func (s *TypedDeclIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterTypedDeclInt(s)
	}
}

func (s *TypedDeclIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitTypedDeclInt(s)
	}
}

func (s *TypedDeclIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitTypedDeclInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Decl_stmt() (localctx IDecl_stmtContext) {
	localctx = NewDecl_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VGrammarRULE_decl_stmt)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInferDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)
			p.Match(VGrammarOP_DECLARATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.expr(0)
		}

	case 2:
		localctx = NewTypedDeclIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Match(VGrammarRW_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewTypedDeclStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.Match(VGrammarRW_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewTypedDeclFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(72)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.Match(VGrammarRW_FLOAT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewTypedDeclBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(75)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(VGrammarRW_BOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewTypedDeclRuneContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(78)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(VGrammarRW_RUNE)
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

type AssignSimpleContext struct {
	Assign_stmtContext
}

func NewAssignSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignSimpleContext {
	var p = new(AssignSimpleContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *AssignSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignSimpleContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *AssignSimpleContext) OP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_ASSIGN, 0)
}

func (s *AssignSimpleContext) Expr() IExprContext {
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

func (s *AssignSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAssignSimple(s)
	}
}

func (s *AssignSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAssignSimple(s)
	}
}

func (s *AssignSimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAssignSimple(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignAddContext struct {
	Assign_stmtContext
}

func NewAssignAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignAddContext {
	var p = new(AssignAddContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *AssignAddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignAddContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *AssignAddContext) OP_ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_ADD_ASSIGN, 0)
}

func (s *AssignAddContext) Expr() IExprContext {
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

func (s *AssignAddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAssignAdd(s)
	}
}

func (s *AssignAddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAssignAdd(s)
	}
}

func (s *AssignAddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAssignAdd(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignSubContext struct {
	Assign_stmtContext
}

func NewAssignSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignSubContext {
	var p = new(AssignSubContext)

	InitEmptyAssign_stmtContext(&p.Assign_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_stmtContext))

	return p
}

func (s *AssignSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignSubContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *AssignSubContext) OP_SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_SUB_ASSIGN, 0)
}

func (s *AssignSubContext) Expr() IExprContext {
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

func (s *AssignSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAssignSub(s)
	}
}

func (s *AssignSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAssignSub(s)
	}
}

func (s *AssignSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAssignSub(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Assign_stmt() (localctx IAssign_stmtContext) {
	localctx = NewAssign_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VGrammarRULE_assign_stmt)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Match(VGrammarOP_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.expr(0)
		}

	case 2:
		localctx = NewAssignAddContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(VGrammarOP_ADD_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.expr(0)
		}

	case 3:
		localctx = NewAssignSubContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(89)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Match(VGrammarOP_SUB_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
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

type ExprGreaterContext struct {
	ExprContext
}

func NewExprGreaterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprGreaterContext {
	var p = new(ExprGreaterContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprGreaterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprGreaterContext) AllExpr() []IExprContext {
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

func (s *ExprGreaterContext) Expr(i int) IExprContext {
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

func (s *ExprGreaterContext) OP_MAYORQ() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MAYORQ, 0)
}

func (s *ExprGreaterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprGreater(s)
	}
}

func (s *ExprGreaterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprGreater(s)
	}
}

func (s *ExprGreaterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprGreater(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprParenContext struct {
	ExprContext
}

func NewExprParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprParenContext {
	var p = new(ExprParenContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprParenContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *ExprParenContext) Expr() IExprContext {
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

func (s *ExprParenContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *ExprParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprParen(s)
	}
}

func (s *ExprParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprParen(s)
	}
}

func (s *ExprParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprParen(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprNotEqualContext struct {
	ExprContext
}

func NewExprNotEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprNotEqualContext {
	var p = new(ExprNotEqualContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprNotEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprNotEqualContext) AllExpr() []IExprContext {
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

func (s *ExprNotEqualContext) Expr(i int) IExprContext {
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

func (s *ExprNotEqualContext) OP_DIFERENTE() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DIFERENTE, 0)
}

func (s *ExprNotEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprNotEqual(s)
	}
}

func (s *ExprNotEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprNotEqual(s)
	}
}

func (s *ExprNotEqualContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprNotEqual(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprLessEqContext struct {
	ExprContext
}

func NewExprLessEqContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprLessEqContext {
	var p = new(ExprLessEqContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprLessEqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprLessEqContext) AllExpr() []IExprContext {
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

func (s *ExprLessEqContext) Expr(i int) IExprContext {
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

func (s *ExprLessEqContext) OP_MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MENORIGUAL, 0)
}

func (s *ExprLessEqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprLessEq(s)
	}
}

func (s *ExprLessEqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprLessEq(s)
	}
}

func (s *ExprLessEqContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprLessEq(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprLessContext struct {
	ExprContext
}

func NewExprLessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprLessContext {
	var p = new(ExprLessContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprLessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprLessContext) AllExpr() []IExprContext {
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

func (s *ExprLessContext) Expr(i int) IExprContext {
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

func (s *ExprLessContext) OP_MENORQ() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MENORQ, 0)
}

func (s *ExprLessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprLess(s)
	}
}

func (s *ExprLessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprLess(s)
	}
}

func (s *ExprLessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprLess(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprNotContext struct {
	ExprContext
}

func NewExprNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprNotContext {
	var p = new(ExprNotContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprNotContext) OP_NOT() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_NOT, 0)
}

func (s *ExprNotContext) Expr() IExprContext {
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

func (s *ExprNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprNot(s)
	}
}

func (s *ExprNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprNot(s)
	}
}

func (s *ExprNotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprDivContext struct {
	ExprContext
}

func NewExprDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprDivContext {
	var p = new(ExprDivContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprDivContext) AllExpr() []IExprContext {
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

func (s *ExprDivContext) Expr(i int) IExprContext {
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

func (s *ExprDivContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DIV, 0)
}

func (s *ExprDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprDiv(s)
	}
}

func (s *ExprDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprDiv(s)
	}
}

func (s *ExprDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprAndContext struct {
	ExprContext
}

func NewExprAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAndContext {
	var p = new(ExprAndContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAndContext) AllExpr() []IExprContext {
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

func (s *ExprAndContext) Expr(i int) IExprContext {
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

func (s *ExprAndContext) OP_AND() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_AND, 0)
}

func (s *ExprAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprAnd(s)
	}
}

func (s *ExprAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprAnd(s)
	}
}

func (s *ExprAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprGreaterEqContext struct {
	ExprContext
}

func NewExprGreaterEqContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprGreaterEqContext {
	var p = new(ExprGreaterEqContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprGreaterEqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprGreaterEqContext) AllExpr() []IExprContext {
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

func (s *ExprGreaterEqContext) Expr(i int) IExprContext {
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

func (s *ExprGreaterEqContext) OP_MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MAYORIGUAL, 0)
}

func (s *ExprGreaterEqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprGreaterEq(s)
	}
}

func (s *ExprGreaterEqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprGreaterEq(s)
	}
}

func (s *ExprGreaterEqContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprGreaterEq(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprOrContext struct {
	ExprContext
}

func NewExprOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprOrContext {
	var p = new(ExprOrContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprOrContext) AllExpr() []IExprContext {
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

func (s *ExprOrContext) Expr(i int) IExprContext {
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

func (s *ExprOrContext) OP_OR() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_OR, 0)
}

func (s *ExprOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprOr(s)
	}
}

func (s *ExprOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprOr(s)
	}
}

func (s *ExprOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprSubContext struct {
	ExprContext
}

func NewExprSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprSubContext {
	var p = new(ExprSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprSubContext) AllExpr() []IExprContext {
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

func (s *ExprSubContext) Expr(i int) IExprContext {
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

func (s *ExprSubContext) OP_RESTA() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_RESTA, 0)
}

func (s *ExprSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprSub(s)
	}
}

func (s *ExprSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprSub(s)
	}
}

func (s *ExprSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprLiteralContext struct {
	ExprContext
}

func NewExprLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprLiteralContext {
	var p = new(ExprLiteralContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprLiteralContext) Literal() ILiteralContext {
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

func (s *ExprLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprLiteral(s)
	}
}

func (s *ExprLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprLiteral(s)
	}
}

func (s *ExprLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprMulContext struct {
	ExprContext
}

func NewExprMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprMulContext {
	var p = new(ExprMulContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprMulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprMulContext) AllExpr() []IExprContext {
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

func (s *ExprMulContext) Expr(i int) IExprContext {
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

func (s *ExprMulContext) OP_MULT() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MULT, 0)
}

func (s *ExprMulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprMul(s)
	}
}

func (s *ExprMulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprMul(s)
	}
}

func (s *ExprMulContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprMul(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprEqualContext struct {
	ExprContext
}

func NewExprEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprEqualContext {
	var p = new(ExprEqualContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprEqualContext) AllExpr() []IExprContext {
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

func (s *ExprEqualContext) Expr(i int) IExprContext {
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

func (s *ExprEqualContext) OP_IGUAL() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_IGUAL, 0)
}

func (s *ExprEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprEqual(s)
	}
}

func (s *ExprEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprEqual(s)
	}
}

func (s *ExprEqualContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprEqual(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprIdContext struct {
	ExprContext
}

func NewExprIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprIdContext {
	var p = new(ExprIdContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIdContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *ExprIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprId(s)
	}
}

func (s *ExprIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprId(s)
	}
}

func (s *ExprIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprId(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprAddContext struct {
	ExprContext
}

func NewExprAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAddContext {
	var p = new(ExprAddContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprAddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAddContext) AllExpr() []IExprContext {
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

func (s *ExprAddContext) Expr(i int) IExprContext {
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

func (s *ExprAddContext) OP_SUMA() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_SUMA, 0)
}

func (s *ExprAddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprAdd(s)
	}
}

func (s *ExprAddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprAdd(s)
	}
}

func (s *ExprAddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprAdd(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprModContext struct {
	ExprContext
}

func NewExprModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprModContext {
	var p = new(ExprModContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprModContext) AllExpr() []IExprContext {
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

func (s *ExprModContext) Expr(i int) IExprContext {
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

func (s *ExprModContext) OP_MOD() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MOD, 0)
}

func (s *ExprModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterExprMod(s)
	}
}

func (s *ExprModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitExprMod(s)
	}
}

func (s *ExprModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitExprMod(s)

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
	_startState := 8
	p.EnterRecursionRule(localctx, 8, VGrammarRULE_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VGrammarRW_TRUE, VGrammarRW_FALSE, VGrammarRW_NIL, VGrammarINT_LITERAL, VGrammarFLOAT_LITERAL, VGrammarSTRING_LITERAL, VGrammarRUNE_LITERAL:
		localctx = NewExprLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(95)
			p.Literal()
		}

	case VGrammarID:
		localctx = NewExprIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(96)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarOP_NOT:
		localctx = NewExprNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(97)
			p.Match(VGrammarOP_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.expr(2)
		}

	case VGrammarPAR_IZQ:
		localctx = NewExprParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(99)
			p.Match(VGrammarPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.expr(0)
		}
		{
			p.SetState(101)
			p.Match(VGrammarPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(144)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprAddContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(106)
					p.Match(VGrammarOP_SUMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(107)
					p.expr(16)
				}

			case 2:
				localctx = NewExprSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(109)
					p.Match(VGrammarOP_RESTA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(110)
					p.expr(15)
				}

			case 3:
				localctx = NewExprMulContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(111)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(112)
					p.Match(VGrammarOP_MULT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(113)
					p.expr(14)
				}

			case 4:
				localctx = NewExprDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(114)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(115)
					p.Match(VGrammarOP_DIV)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(116)
					p.expr(13)
				}

			case 5:
				localctx = NewExprModContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(117)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(118)
					p.Match(VGrammarOP_MOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(119)
					p.expr(12)
				}

			case 6:
				localctx = NewExprEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(120)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(121)
					p.Match(VGrammarOP_IGUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(122)
					p.expr(11)
				}

			case 7:
				localctx = NewExprNotEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(123)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(124)
					p.Match(VGrammarOP_DIFERENTE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(125)
					p.expr(10)
				}

			case 8:
				localctx = NewExprLessContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(126)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(127)
					p.Match(VGrammarOP_MENORQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(128)
					p.expr(9)
				}

			case 9:
				localctx = NewExprGreaterContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(129)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(130)
					p.Match(VGrammarOP_MAYORQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(131)
					p.expr(8)
				}

			case 10:
				localctx = NewExprLessEqContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(132)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(133)
					p.Match(VGrammarOP_MENORIGUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(134)
					p.expr(7)
				}

			case 11:
				localctx = NewExprGreaterEqContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(135)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(136)
					p.Match(VGrammarOP_MAYORIGUAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(137)
					p.expr(6)
				}

			case 12:
				localctx = NewExprAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(138)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(139)
					p.Match(VGrammarOP_AND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(140)
					p.expr(5)
				}

			case 13:
				localctx = NewExprOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(141)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(142)
					p.Match(VGrammarOP_OR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(143)
					p.expr(4)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
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

type LitStringContext struct {
	LiteralContext
}

func NewLitStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitStringContext {
	var p = new(LitStringContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LitStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitStringContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarSTRING_LITERAL, 0)
}

func (s *LitStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLitString(s)
	}
}

func (s *LitStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLitString(s)
	}
}

func (s *LitStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type LitRuneContext struct {
	LiteralContext
}

func NewLitRuneContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitRuneContext {
	var p = new(LitRuneContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LitRuneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitRuneContext) RUNE_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarRUNE_LITERAL, 0)
}

func (s *LitRuneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLitRune(s)
	}
}

func (s *LitRuneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLitRune(s)
	}
}

func (s *LitRuneContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLitRune(s)

	default:
		return t.VisitChildren(s)
	}
}

type LitFloatContext struct {
	LiteralContext
}

func NewLitFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitFloatContext {
	var p = new(LitFloatContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LitFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitFloatContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarFLOAT_LITERAL, 0)
}

func (s *LitFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLitFloat(s)
	}
}

func (s *LitFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLitFloat(s)
	}
}

func (s *LitFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLitFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type LitTrueContext struct {
	LiteralContext
}

func NewLitTrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitTrueContext {
	var p = new(LitTrueContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LitTrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitTrueContext) RW_TRUE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_TRUE, 0)
}

func (s *LitTrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLitTrue(s)
	}
}

func (s *LitTrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLitTrue(s)
	}
}

func (s *LitTrueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLitTrue(s)

	default:
		return t.VisitChildren(s)
	}
}

type LitIntContext struct {
	LiteralContext
}

func NewLitIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitIntContext {
	var p = new(LitIntContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LitIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitIntContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarINT_LITERAL, 0)
}

func (s *LitIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLitInt(s)
	}
}

func (s *LitIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLitInt(s)
	}
}

func (s *LitIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

type LitNilContext struct {
	LiteralContext
}

func NewLitNilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitNilContext {
	var p = new(LitNilContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LitNilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitNilContext) RW_NIL() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_NIL, 0)
}

func (s *LitNilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLitNil(s)
	}
}

func (s *LitNilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLitNil(s)
	}
}

func (s *LitNilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLitNil(s)

	default:
		return t.VisitChildren(s)
	}
}

type LitFalseContext struct {
	LiteralContext
}

func NewLitFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitFalseContext {
	var p = new(LitFalseContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *LitFalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitFalseContext) RW_FALSE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FALSE, 0)
}

func (s *LitFalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterLitFalse(s)
	}
}

func (s *LitFalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitLitFalse(s)
	}
}

func (s *LitFalseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitLitFalse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VGrammarRULE_literal)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VGrammarINT_LITERAL:
		localctx = NewLitIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Match(VGrammarINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarFLOAT_LITERAL:
		localctx = NewLitFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.Match(VGrammarFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarSTRING_LITERAL:
		localctx = NewLitStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.Match(VGrammarSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRUNE_LITERAL:
		localctx = NewLitRuneContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(152)
			p.Match(VGrammarRUNE_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_TRUE:
		localctx = NewLitTrueContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(153)
			p.Match(VGrammarRW_TRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_FALSE:
		localctx = NewLitFalseContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(154)
			p.Match(VGrammarRW_FALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_NIL:
		localctx = NewLitNilContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(155)
			p.Match(VGrammarRW_NIL)
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

// IFlow_stmtContext is an interface to support dynamic dispatch.
type IFlow_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFlow_stmtContext differentiates from other interfaces.
	IsFlow_stmtContext()
}

type Flow_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlow_stmtContext() *Flow_stmtContext {
	var p = new(Flow_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_flow_stmt
	return p
}

func InitEmptyFlow_stmtContext(p *Flow_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_flow_stmt
}

func (*Flow_stmtContext) IsFlow_stmtContext() {}

func NewFlow_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Flow_stmtContext {
	var p = new(Flow_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_flow_stmt

	return p
}

func (s *Flow_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Flow_stmtContext) CopyAll(ctx *Flow_stmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Flow_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Flow_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StmtReturnContext struct {
	Flow_stmtContext
}

func NewStmtReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtReturnContext {
	var p = new(StmtReturnContext)

	InitEmptyFlow_stmtContext(&p.Flow_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Flow_stmtContext))

	return p
}

func (s *StmtReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtReturnContext) RW_RETURN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_RETURN, 0)
}

func (s *StmtReturnContext) Expr() IExprContext {
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

func (s *StmtReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmtReturn(s)
	}
}

func (s *StmtReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmtReturn(s)
	}
}

func (s *StmtReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmtReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtContinueContext struct {
	Flow_stmtContext
}

func NewStmtContinueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtContinueContext {
	var p = new(StmtContinueContext)

	InitEmptyFlow_stmtContext(&p.Flow_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Flow_stmtContext))

	return p
}

func (s *StmtContinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContinueContext) RW_CONTINUE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_CONTINUE, 0)
}

func (s *StmtContinueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmtContinue(s)
	}
}

func (s *StmtContinueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmtContinue(s)
	}
}

func (s *StmtContinueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmtContinue(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtBreakContext struct {
	Flow_stmtContext
}

func NewStmtBreakContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtBreakContext {
	var p = new(StmtBreakContext)

	InitEmptyFlow_stmtContext(&p.Flow_stmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*Flow_stmtContext))

	return p
}

func (s *StmtBreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtBreakContext) RW_BREAK() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_BREAK, 0)
}

func (s *StmtBreakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStmtBreak(s)
	}
}

func (s *StmtBreakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStmtBreak(s)
	}
}

func (s *StmtBreakContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStmtBreak(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Flow_stmt() (localctx IFlow_stmtContext) {
	localctx = NewFlow_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VGrammarRULE_flow_stmt)
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VGrammarRW_BREAK:
		localctx = NewStmtBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Match(VGrammarRW_BREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_CONTINUE:
		localctx = NewStmtContinueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.Match(VGrammarRW_CONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_RETURN:
		localctx = NewStmtReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(160)
			p.Match(VGrammarRW_RETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(161)
				p.expr(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
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

type PrintCallContext struct {
	Func_callContext
}

func NewPrintCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintCallContext {
	var p = new(PrintCallContext)

	InitEmptyFunc_callContext(&p.Func_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_callContext))

	return p
}

func (s *PrintCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintCallContext) RW_PRINT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_PRINT, 0)
}

func (s *PrintCallContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *PrintCallContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *PrintCallContext) AllExpr() []IExprContext {
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

func (s *PrintCallContext) Expr(i int) IExprContext {
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

func (s *PrintCallContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *PrintCallContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *PrintCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterPrintCall(s)
	}
}

func (s *PrintCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitPrintCall(s)
	}
}

func (s *PrintCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitPrintCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallContext struct {
	Func_callContext
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	InitEmptyFunc_callContext(&p.Func_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_callContext))

	return p
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *FunctionCallContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *FunctionCallContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *FunctionCallContext) AllExpr() []IExprContext {
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

func (s *FunctionCallContext) Expr(i int) IExprContext {
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

func (s *FunctionCallContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *FunctionCallContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintLnCallContext struct {
	Func_callContext
}

func NewPrintLnCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintLnCallContext {
	var p = new(PrintLnCallContext)

	InitEmptyFunc_callContext(&p.Func_callContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_callContext))

	return p
}

func (s *PrintLnCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintLnCallContext) RW_PRINTLN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_PRINTLN, 0)
}

func (s *PrintLnCallContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *PrintLnCallContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *PrintLnCallContext) AllExpr() []IExprContext {
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

func (s *PrintLnCallContext) Expr(i int) IExprContext {
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

func (s *PrintLnCallContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *PrintLnCallContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *PrintLnCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterPrintLnCall(s)
	}
}

func (s *PrintLnCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitPrintLnCall(s)
	}
}

func (s *PrintLnCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitPrintLnCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Func_call() (localctx IFunc_callContext) {
	localctx = NewFunc_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VGrammarRULE_func_call)
	var _la int

	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VGrammarID:
		localctx = NewFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Match(VGrammarPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&1091273880502279) != 0 {
			{
				p.SetState(168)
				p.expr(0)
			}
			p.SetState(173)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == VGrammarCOMA {
				{
					p.SetState(169)
					p.Match(VGrammarCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(170)
					p.expr(0)
				}

				p.SetState(175)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(178)
			p.Match(VGrammarPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_PRINT:
		localctx = NewPrintCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.Match(VGrammarRW_PRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Match(VGrammarPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&1091273880502279) != 0 {
			{
				p.SetState(181)
				p.expr(0)
			}
			p.SetState(186)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == VGrammarCOMA {
				{
					p.SetState(182)
					p.Match(VGrammarCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(183)
					p.expr(0)
				}

				p.SetState(188)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(191)
			p.Match(VGrammarPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarRW_PRINTLN:
		localctx = NewPrintLnCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(192)
			p.Match(VGrammarRW_PRINTLN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Match(VGrammarPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&1091273880502279) != 0 {
			{
				p.SetState(194)
				p.expr(0)
			}
			p.SetState(199)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == VGrammarCOMA {
				{
					p.SetState(195)
					p.Match(VGrammarCOMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(196)
					p.expr(0)
				}

				p.SetState(201)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(204)
			p.Match(VGrammarPAR_DER)
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

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_IF() antlr.TerminalNode
	Expr() IExprContext
	AllLLAVE_IZQ() []antlr.TerminalNode
	LLAVE_IZQ(i int) antlr.TerminalNode
	AllLLAVE_DER() []antlr.TerminalNode
	LLAVE_DER(i int) antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	RW_ELSE() antlr.TerminalNode

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

func (s *If_stmtContext) RW_IF() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_IF, 0)
}

func (s *If_stmtContext) Expr() IExprContext {
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

func (s *If_stmtContext) AllLLAVE_IZQ() []antlr.TerminalNode {
	return s.GetTokens(VGrammarLLAVE_IZQ)
}

func (s *If_stmtContext) LLAVE_IZQ(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, i)
}

func (s *If_stmtContext) AllLLAVE_DER() []antlr.TerminalNode {
	return s.GetTokens(VGrammarLLAVE_DER)
}

func (s *If_stmtContext) LLAVE_DER(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, i)
}

func (s *If_stmtContext) AllStmt() []IStmtContext {
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

func (s *If_stmtContext) Stmt(i int) IStmtContext {
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

func (s *If_stmtContext) RW_ELSE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_ELSE, 0)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterIf_stmt(s)
	}
}

func (s *If_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitIf_stmt(s)
	}
}

func (s *If_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitIf_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) If_stmt() (localctx IIf_stmtContext) {
	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VGrammarRULE_if_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(VGrammarRW_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.expr(0)
	}
	{
		p.SetState(209)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&4611686018433686831) != 0 {
		{
			p.SetState(210)
			p.Stmt()
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(216)
		p.Match(VGrammarLLAVE_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarRW_ELSE {
		{
			p.SetState(217)
			p.Match(VGrammarRW_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.Match(VGrammarLLAVE_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&4611686018433686831) != 0 {
			{
				p.SetState(219)
				p.Stmt()
			}

			p.SetState(224)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(225)
			p.Match(VGrammarLLAVE_DER)
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

// ISwitch_stmtContext is an interface to support dynamic dispatch.
type ISwitch_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_SWITCH() antlr.TerminalNode
	Expr() IExprContext
	LLAVE_IZQ() antlr.TerminalNode
	LLAVE_DER() antlr.TerminalNode
	AllSwitch_case() []ISwitch_caseContext
	Switch_case(i int) ISwitch_caseContext
	Default_case() IDefault_caseContext

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

func (s *Switch_stmtContext) RW_SWITCH() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_SWITCH, 0)
}

func (s *Switch_stmtContext) Expr() IExprContext {
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

func (s *Switch_stmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *Switch_stmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *Switch_stmtContext) AllSwitch_case() []ISwitch_caseContext {
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

func (s *Switch_stmtContext) Switch_case(i int) ISwitch_caseContext {
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

func (s *Switch_stmtContext) Default_case() IDefault_caseContext {
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

func (s *Switch_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterSwitch_stmt(s)
	}
}

func (s *Switch_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitSwitch_stmt(s)
	}
}

func (s *Switch_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitSwitch_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Switch_stmt() (localctx ISwitch_stmtContext) {
	localctx = NewSwitch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VGrammarRULE_switch_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(VGrammarRW_SWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.expr(0)
	}
	{
		p.SetState(230)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VGrammarRW_CASE {
		{
			p.SetState(231)
			p.Switch_case()
		}

		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarRW_DEFAULT {
		{
			p.SetState(237)
			p.Default_case()
		}

	}
	{
		p.SetState(240)
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

	// Getter signatures
	RW_CASE() antlr.TerminalNode
	Expr() IExprContext
	DOS_PUNTOS() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

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

func (s *Switch_caseContext) RW_CASE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_CASE, 0)
}

func (s *Switch_caseContext) Expr() IExprContext {
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

func (s *Switch_caseContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, 0)
}

func (s *Switch_caseContext) AllStmt() []IStmtContext {
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

func (s *Switch_caseContext) Stmt(i int) IStmtContext {
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

func (s *Switch_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterSwitch_case(s)
	}
}

func (s *Switch_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitSwitch_case(s)
	}
}

func (s *Switch_caseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitSwitch_case(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Switch_case() (localctx ISwitch_caseContext) {
	localctx = NewSwitch_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VGrammarRULE_switch_case)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(VGrammarRW_CASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.expr(0)
	}
	{
		p.SetState(244)
		p.Match(VGrammarDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&4611686018433686831) != 0 {
		{
			p.SetState(245)
			p.Stmt()
		}

		p.SetState(250)
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

	// Getter signatures
	RW_DEFAULT() antlr.TerminalNode
	DOS_PUNTOS() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

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

func (s *Default_caseContext) RW_DEFAULT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_DEFAULT, 0)
}

func (s *Default_caseContext) DOS_PUNTOS() antlr.TerminalNode {
	return s.GetToken(VGrammarDOS_PUNTOS, 0)
}

func (s *Default_caseContext) AllStmt() []IStmtContext {
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

func (s *Default_caseContext) Stmt(i int) IStmtContext {
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

func (s *Default_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Default_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterDefault_case(s)
	}
}

func (s *Default_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitDefault_case(s)
	}
}

func (s *Default_caseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitDefault_case(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Default_case() (localctx IDefault_caseContext) {
	localctx = NewDefault_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VGrammarRULE_default_case)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(VGrammarRW_DEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)
		p.Match(VGrammarDOS_PUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&4611686018433686831) != 0 {
		{
			p.SetState(253)
			p.Stmt()
		}

		p.SetState(258)
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

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	RW_IN() antlr.TerminalNode
	Expr() IExprContext
	LLAVE_IZQ() antlr.TerminalNode
	LLAVE_DER() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

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

func (s *For_stmtContext) RW_FOR() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FOR, 0)
}

func (s *For_stmtContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *For_stmtContext) RW_IN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_IN, 0)
}

func (s *For_stmtContext) Expr() IExprContext {
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

func (s *For_stmtContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *For_stmtContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *For_stmtContext) AllStmt() []IStmtContext {
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

func (s *For_stmtContext) Stmt(i int) IStmtContext {
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

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFor_stmt(s)
	}
}

func (s *For_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFor_stmt(s)
	}
}

func (s *For_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFor_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) For_stmt() (localctx IFor_stmtContext) {
	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VGrammarRULE_for_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(VGrammarRW_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Match(VGrammarRW_IN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.expr(0)
	}
	{
		p.SetState(263)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&4611686018433686831) != 0 {
		{
			p.SetState(264)
			p.Stmt()
		}

		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(270)
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

// IFunc_declContext is an interface to support dynamic dispatch.
type IFunc_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_FN() antlr.TerminalNode
	PAR_IZQ() antlr.TerminalNode
	PAR_DER() antlr.TerminalNode
	LLAVE_IZQ() antlr.TerminalNode
	LLAVE_DER() antlr.TerminalNode
	ID() antlr.TerminalNode
	RW_MAIN() antlr.TerminalNode
	AllParam() []IParamContext
	Param(i int) IParamContext
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext
	RW_INT() antlr.TerminalNode
	RW_STRING() antlr.TerminalNode
	RW_FLOAT64() antlr.TerminalNode
	RW_BOOL() antlr.TerminalNode
	RW_RUNE() antlr.TerminalNode
	AllCOMA() []antlr.TerminalNode
	COMA(i int) antlr.TerminalNode

	// IsFunc_declContext differentiates from other interfaces.
	IsFunc_declContext()
}

type Func_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_declContext() *Func_declContext {
	var p = new(Func_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_func_decl
	return p
}

func InitEmptyFunc_declContext(p *Func_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_func_decl
}

func (*Func_declContext) IsFunc_declContext() {}

func NewFunc_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_declContext {
	var p = new(Func_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_func_decl

	return p
}

func (s *Func_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_declContext) RW_FN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FN, 0)
}

func (s *Func_declContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *Func_declContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *Func_declContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *Func_declContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *Func_declContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *Func_declContext) RW_MAIN() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MAIN, 0)
}

func (s *Func_declContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *Func_declContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
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

	return t.(IParamContext)
}

func (s *Func_declContext) AllStmt() []IStmtContext {
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

func (s *Func_declContext) Stmt(i int) IStmtContext {
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

func (s *Func_declContext) RW_INT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_INT, 0)
}

func (s *Func_declContext) RW_STRING() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_STRING, 0)
}

func (s *Func_declContext) RW_FLOAT64() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FLOAT64, 0)
}

func (s *Func_declContext) RW_BOOL() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_BOOL, 0)
}

func (s *Func_declContext) RW_RUNE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_RUNE, 0)
}

func (s *Func_declContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(VGrammarCOMA)
}

func (s *Func_declContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(VGrammarCOMA, i)
}

func (s *Func_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterFunc_decl(s)
	}
}

func (s *Func_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitFunc_decl(s)
	}
}

func (s *Func_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitFunc_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Func_decl() (localctx IFunc_declContext) {
	localctx = NewFunc_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VGrammarRULE_func_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(VGrammarRW_FN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		_la = p.GetTokenStream().LA(1)

		if !(_la == VGrammarRW_MAIN || _la == VGrammarID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(274)
		p.Match(VGrammarPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == VGrammarID {
		{
			p.SetState(275)
			p.Param()
		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VGrammarCOMA {
			{
				p.SetState(276)
				p.Match(VGrammarCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(277)
				p.Param()
			}

			p.SetState(282)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(285)
		p.Match(VGrammarPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8126464) != 0 {
		{
			p.SetState(286)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8126464) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(289)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-2)) & ^0x3f) == 0 && ((int64(1)<<(_la-2))&4611686018433686831) != 0 {
		{
			p.SetState(290)
			p.Stmt()
		}

		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(296)
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

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	RW_INT() antlr.TerminalNode
	RW_STRING() antlr.TerminalNode
	RW_FLOAT64() antlr.TerminalNode
	RW_BOOL() antlr.TerminalNode
	RW_RUNE() antlr.TerminalNode

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *ParamContext) RW_INT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_INT, 0)
}

func (s *ParamContext) RW_STRING() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_STRING, 0)
}

func (s *ParamContext) RW_FLOAT64() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FLOAT64, 0)
}

func (s *ParamContext) RW_BOOL() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_BOOL, 0)
}

func (s *ParamContext) RW_RUNE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_RUNE, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitParam(s)
	}
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, VGrammarRULE_param)
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(298)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Match(VGrammarRW_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(300)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.Match(VGrammarRW_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(302)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.Match(VGrammarRW_FLOAT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(304)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Match(VGrammarRW_BOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(306)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.Match(VGrammarRW_RUNE)
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

// IStruct_declContext is an interface to support dynamic dispatch.
type IStruct_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_STRUCT() antlr.TerminalNode
	ID() antlr.TerminalNode
	LLAVE_IZQ() antlr.TerminalNode
	LLAVE_DER() antlr.TerminalNode
	AllStruct_field() []IStruct_fieldContext
	Struct_field(i int) IStruct_fieldContext

	// IsStruct_declContext differentiates from other interfaces.
	IsStruct_declContext()
}

type Struct_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_declContext() *Struct_declContext {
	var p = new(Struct_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_decl
	return p
}

func InitEmptyStruct_declContext(p *Struct_declContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_decl
}

func (*Struct_declContext) IsStruct_declContext() {}

func NewStruct_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_declContext {
	var p = new(Struct_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_struct_decl

	return p
}

func (s *Struct_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_declContext) RW_STRUCT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_STRUCT, 0)
}

func (s *Struct_declContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *Struct_declContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_IZQ, 0)
}

func (s *Struct_declContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarLLAVE_DER, 0)
}

func (s *Struct_declContext) AllStruct_field() []IStruct_fieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStruct_fieldContext); ok {
			len++
		}
	}

	tst := make([]IStruct_fieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStruct_fieldContext); ok {
			tst[i] = t.(IStruct_fieldContext)
			i++
		}
	}

	return tst
}

func (s *Struct_declContext) Struct_field(i int) IStruct_fieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_fieldContext); ok {
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

	return t.(IStruct_fieldContext)
}

func (s *Struct_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStruct_decl(s)
	}
}

func (s *Struct_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStruct_decl(s)
	}
}

func (s *Struct_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStruct_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Struct_decl() (localctx IStruct_declContext) {
	localctx = NewStruct_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, VGrammarRULE_struct_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(VGrammarRW_STRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(311)
		p.Match(VGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)
		p.Match(VGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VGrammarRW_MUT {
		{
			p.SetState(313)
			p.Struct_field()
		}

		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(319)
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

// IStruct_fieldContext is an interface to support dynamic dispatch.
type IStruct_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_MUT() antlr.TerminalNode
	ID() antlr.TerminalNode
	RW_INT() antlr.TerminalNode
	RW_STRING() antlr.TerminalNode
	RW_FLOAT64() antlr.TerminalNode
	RW_BOOL() antlr.TerminalNode
	RW_RUNE() antlr.TerminalNode

	// IsStruct_fieldContext differentiates from other interfaces.
	IsStruct_fieldContext()
}

type Struct_fieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_fieldContext() *Struct_fieldContext {
	var p = new(Struct_fieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_field
	return p
}

func InitEmptyStruct_fieldContext(p *Struct_fieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_struct_field
}

func (*Struct_fieldContext) IsStruct_fieldContext() {}

func NewStruct_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_fieldContext {
	var p = new(Struct_fieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_struct_field

	return p
}

func (s *Struct_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_fieldContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_MUT, 0)
}

func (s *Struct_fieldContext) ID() antlr.TerminalNode {
	return s.GetToken(VGrammarID, 0)
}

func (s *Struct_fieldContext) RW_INT() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_INT, 0)
}

func (s *Struct_fieldContext) RW_STRING() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_STRING, 0)
}

func (s *Struct_fieldContext) RW_FLOAT64() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_FLOAT64, 0)
}

func (s *Struct_fieldContext) RW_BOOL() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_BOOL, 0)
}

func (s *Struct_fieldContext) RW_RUNE() antlr.TerminalNode {
	return s.GetToken(VGrammarRW_RUNE, 0)
}

func (s *Struct_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterStruct_field(s)
	}
}

func (s *Struct_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitStruct_field(s)
	}
}

func (s *Struct_fieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitStruct_field(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Struct_field() (localctx IStruct_fieldContext) {
	localctx = NewStruct_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, VGrammarRULE_struct_field)
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.Match(VGrammarRW_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)
			p.Match(VGrammarRW_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(327)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.Match(VGrammarRW_FLOAT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(330)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(VGrammarRW_BOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(333)
			p.Match(VGrammarRW_MUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Match(VGrammarID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Match(VGrammarRW_RUNE)
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

func (p *VGrammar) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
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
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
