// Code generated from ZenithLexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ZenithLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ZenithLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func zenithlexerLexerInit() {
	staticData := &ZenithLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'!!'", "'~'", "'^'", "'*'", "'/'", "'%'", "'+'", "'-'", "'<<'",
		"'>>'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'??'", "'null'",
		"'true'", "'alloc'", "'dealloc'", "'false'", "'if'", "'else'", "'='",
		"':='", "", "", "", "'&'", "'$'", "'|'", "'!'", "'?'", "'#'", "'('",
		"')'", "'{'", "'}'", "'['", "']'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "NOT_NULL", "BIT_NOT", "POW", "TIMES", "DIVIDE", "REM", "PLUS",
		"MINUS", "LSHIFT", "RSHIFT", "LT", "GT", "LTE", "GTE", "EQ", "NEQ",
		"COALESCE", "NULL", "TRUE", "ALLOC", "DEALLOC", "FALSE", "IF", "ELSE",
		"ASSIGN", "INIT_ASSIGN", "TYPE", "NUM", "ID", "AND", "DOLLAR", "PIPE",
		"EXCLAIM", "QUESTION", "HASH", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
		"LBRACK", "RBRACK", "SEMICOLON", "NL", "IGNORED",
	}
	staticData.RuleNames = []string{
		"NOT_NULL", "BIT_NOT", "POW", "TIMES", "DIVIDE", "REM", "PLUS", "MINUS",
		"LSHIFT", "RSHIFT", "LT", "GT", "LTE", "GTE", "EQ", "NEQ", "COALESCE",
		"NULL", "TRUE", "ALLOC", "DEALLOC", "FALSE", "IF", "ELSE", "ASSIGN",
		"INIT_ASSIGN", "TYPE", "NUM", "DEC_NUM", "DEC_SEG", "BIN_NUM", "BIN_SEG",
		"OCT_NUM", "OCT_SEG", "HEX_NUM", "HEX_SEG", "HEX_DIGIT", "ID", "AND",
		"DOLLAR", "PIPE", "EXCLAIM", "QUESTION", "HASH", "LPAREN", "RPAREN",
		"LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMICOLON", "NL", "IGNORED",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 470, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 260, 8, 26, 1, 27, 1, 27, 1, 27, 1,
		27, 3, 27, 266, 8, 27, 1, 28, 1, 28, 5, 28, 270, 8, 28, 10, 28, 12, 28,
		273, 9, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 280, 8, 28, 10, 28,
		12, 28, 283, 9, 28, 3, 28, 285, 8, 28, 1, 29, 1, 29, 5, 29, 289, 8, 29,
		10, 29, 12, 29, 292, 9, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 299,
		8, 30, 10, 30, 12, 30, 302, 9, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5,
		30, 309, 8, 30, 10, 30, 12, 30, 312, 9, 30, 3, 30, 314, 8, 30, 1, 31, 5,
		31, 317, 8, 31, 10, 31, 12, 31, 320, 9, 31, 1, 31, 1, 31, 5, 31, 324, 8,
		31, 10, 31, 12, 31, 327, 9, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32,
		334, 8, 32, 10, 32, 12, 32, 337, 9, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 5, 32, 344, 8, 32, 10, 32, 12, 32, 347, 9, 32, 3, 32, 349, 8, 32, 1,
		33, 5, 33, 352, 8, 33, 10, 33, 12, 33, 355, 9, 33, 1, 33, 1, 33, 5, 33,
		359, 8, 33, 10, 33, 12, 33, 362, 9, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 5, 34, 370, 8, 34, 10, 34, 12, 34, 373, 9, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 381, 8, 34, 10, 34, 12, 34, 384, 9,
		34, 3, 34, 386, 8, 34, 1, 35, 1, 35, 5, 35, 390, 8, 35, 10, 35, 12, 35,
		393, 9, 35, 1, 35, 1, 35, 1, 35, 5, 35, 398, 8, 35, 10, 35, 12, 35, 401,
		9, 35, 1, 36, 1, 36, 1, 37, 1, 37, 5, 37, 407, 8, 37, 10, 37, 12, 37, 410,
		9, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1,
		42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47,
		1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 3, 51, 439, 8, 51, 1,
		51, 4, 51, 442, 8, 51, 11, 51, 12, 51, 443, 1, 52, 1, 52, 1, 52, 1, 52,
		1, 52, 5, 52, 451, 8, 52, 10, 52, 12, 52, 454, 9, 52, 1, 52, 1, 52, 1,
		52, 1, 52, 1, 52, 1, 52, 5, 52, 462, 8, 52, 10, 52, 12, 52, 465, 9, 52,
		3, 52, 467, 8, 52, 1, 52, 1, 52, 1, 452, 0, 53, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 0, 59, 0, 61, 0, 63, 0, 65,
		0, 67, 0, 69, 0, 71, 0, 73, 0, 75, 29, 77, 30, 79, 31, 81, 32, 83, 33,
		85, 34, 87, 35, 89, 36, 91, 37, 93, 38, 95, 39, 97, 40, 99, 41, 101, 42,
		103, 43, 105, 44, 1, 0, 12, 2, 0, 48, 57, 95, 95, 1, 0, 48, 57, 2, 0, 48,
		49, 95, 95, 1, 0, 48, 49, 2, 0, 48, 55, 95, 95, 1, 0, 48, 55, 3, 0, 48,
		57, 65, 70, 97, 102, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90,
		95, 95, 97, 122, 1, 0, 13, 13, 1, 0, 10, 10, 2, 0, 9, 9, 32, 32, 509, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47,
		1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0,
		55, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0,
		0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0,
		0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0,
		0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103,
		1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 1, 107, 1, 0, 0, 0, 3, 110, 1, 0, 0, 0,
		5, 112, 1, 0, 0, 0, 7, 114, 1, 0, 0, 0, 9, 116, 1, 0, 0, 0, 11, 118, 1,
		0, 0, 0, 13, 120, 1, 0, 0, 0, 15, 122, 1, 0, 0, 0, 17, 124, 1, 0, 0, 0,
		19, 127, 1, 0, 0, 0, 21, 130, 1, 0, 0, 0, 23, 132, 1, 0, 0, 0, 25, 134,
		1, 0, 0, 0, 27, 137, 1, 0, 0, 0, 29, 140, 1, 0, 0, 0, 31, 143, 1, 0, 0,
		0, 33, 146, 1, 0, 0, 0, 35, 149, 1, 0, 0, 0, 37, 154, 1, 0, 0, 0, 39, 159,
		1, 0, 0, 0, 41, 165, 1, 0, 0, 0, 43, 173, 1, 0, 0, 0, 45, 179, 1, 0, 0,
		0, 47, 182, 1, 0, 0, 0, 49, 187, 1, 0, 0, 0, 51, 189, 1, 0, 0, 0, 53, 259,
		1, 0, 0, 0, 55, 265, 1, 0, 0, 0, 57, 284, 1, 0, 0, 0, 59, 286, 1, 0, 0,
		0, 61, 293, 1, 0, 0, 0, 63, 318, 1, 0, 0, 0, 65, 328, 1, 0, 0, 0, 67, 353,
		1, 0, 0, 0, 69, 363, 1, 0, 0, 0, 71, 391, 1, 0, 0, 0, 73, 402, 1, 0, 0,
		0, 75, 404, 1, 0, 0, 0, 77, 411, 1, 0, 0, 0, 79, 413, 1, 0, 0, 0, 81, 415,
		1, 0, 0, 0, 83, 417, 1, 0, 0, 0, 85, 419, 1, 0, 0, 0, 87, 421, 1, 0, 0,
		0, 89, 423, 1, 0, 0, 0, 91, 425, 1, 0, 0, 0, 93, 427, 1, 0, 0, 0, 95, 429,
		1, 0, 0, 0, 97, 431, 1, 0, 0, 0, 99, 433, 1, 0, 0, 0, 101, 435, 1, 0, 0,
		0, 103, 441, 1, 0, 0, 0, 105, 466, 1, 0, 0, 0, 107, 108, 5, 33, 0, 0, 108,
		109, 5, 33, 0, 0, 109, 2, 1, 0, 0, 0, 110, 111, 5, 126, 0, 0, 111, 4, 1,
		0, 0, 0, 112, 113, 5, 94, 0, 0, 113, 6, 1, 0, 0, 0, 114, 115, 5, 42, 0,
		0, 115, 8, 1, 0, 0, 0, 116, 117, 5, 47, 0, 0, 117, 10, 1, 0, 0, 0, 118,
		119, 5, 37, 0, 0, 119, 12, 1, 0, 0, 0, 120, 121, 5, 43, 0, 0, 121, 14,
		1, 0, 0, 0, 122, 123, 5, 45, 0, 0, 123, 16, 1, 0, 0, 0, 124, 125, 5, 60,
		0, 0, 125, 126, 5, 60, 0, 0, 126, 18, 1, 0, 0, 0, 127, 128, 5, 62, 0, 0,
		128, 129, 5, 62, 0, 0, 129, 20, 1, 0, 0, 0, 130, 131, 5, 60, 0, 0, 131,
		22, 1, 0, 0, 0, 132, 133, 5, 62, 0, 0, 133, 24, 1, 0, 0, 0, 134, 135, 5,
		60, 0, 0, 135, 136, 5, 61, 0, 0, 136, 26, 1, 0, 0, 0, 137, 138, 5, 62,
		0, 0, 138, 139, 5, 61, 0, 0, 139, 28, 1, 0, 0, 0, 140, 141, 5, 61, 0, 0,
		141, 142, 5, 61, 0, 0, 142, 30, 1, 0, 0, 0, 143, 144, 5, 33, 0, 0, 144,
		145, 5, 61, 0, 0, 145, 32, 1, 0, 0, 0, 146, 147, 5, 63, 0, 0, 147, 148,
		5, 63, 0, 0, 148, 34, 1, 0, 0, 0, 149, 150, 5, 110, 0, 0, 150, 151, 5,
		117, 0, 0, 151, 152, 5, 108, 0, 0, 152, 153, 5, 108, 0, 0, 153, 36, 1,
		0, 0, 0, 154, 155, 5, 116, 0, 0, 155, 156, 5, 114, 0, 0, 156, 157, 5, 117,
		0, 0, 157, 158, 5, 101, 0, 0, 158, 38, 1, 0, 0, 0, 159, 160, 5, 97, 0,
		0, 160, 161, 5, 108, 0, 0, 161, 162, 5, 108, 0, 0, 162, 163, 5, 111, 0,
		0, 163, 164, 5, 99, 0, 0, 164, 40, 1, 0, 0, 0, 165, 166, 5, 100, 0, 0,
		166, 167, 5, 101, 0, 0, 167, 168, 5, 97, 0, 0, 168, 169, 5, 108, 0, 0,
		169, 170, 5, 108, 0, 0, 170, 171, 5, 111, 0, 0, 171, 172, 5, 99, 0, 0,
		172, 42, 1, 0, 0, 0, 173, 174, 5, 102, 0, 0, 174, 175, 5, 97, 0, 0, 175,
		176, 5, 108, 0, 0, 176, 177, 5, 115, 0, 0, 177, 178, 5, 101, 0, 0, 178,
		44, 1, 0, 0, 0, 179, 180, 5, 105, 0, 0, 180, 181, 5, 102, 0, 0, 181, 46,
		1, 0, 0, 0, 182, 183, 5, 101, 0, 0, 183, 184, 5, 108, 0, 0, 184, 185, 5,
		115, 0, 0, 185, 186, 5, 101, 0, 0, 186, 48, 1, 0, 0, 0, 187, 188, 5, 61,
		0, 0, 188, 50, 1, 0, 0, 0, 189, 190, 5, 58, 0, 0, 190, 191, 5, 61, 0, 0,
		191, 52, 1, 0, 0, 0, 192, 193, 5, 85, 0, 0, 193, 194, 5, 73, 0, 0, 194,
		195, 5, 110, 0, 0, 195, 196, 5, 116, 0, 0, 196, 260, 5, 56, 0, 0, 197,
		198, 5, 85, 0, 0, 198, 199, 5, 73, 0, 0, 199, 200, 5, 110, 0, 0, 200, 201,
		5, 116, 0, 0, 201, 202, 5, 49, 0, 0, 202, 260, 5, 54, 0, 0, 203, 204, 5,
		85, 0, 0, 204, 205, 5, 73, 0, 0, 205, 206, 5, 110, 0, 0, 206, 207, 5, 116,
		0, 0, 207, 208, 5, 51, 0, 0, 208, 260, 5, 50, 0, 0, 209, 210, 5, 85, 0,
		0, 210, 211, 5, 73, 0, 0, 211, 212, 5, 110, 0, 0, 212, 213, 5, 116, 0,
		0, 213, 214, 5, 54, 0, 0, 214, 260, 5, 52, 0, 0, 215, 216, 5, 85, 0, 0,
		216, 217, 5, 73, 0, 0, 217, 218, 5, 110, 0, 0, 218, 260, 5, 116, 0, 0,
		219, 220, 5, 73, 0, 0, 220, 221, 5, 110, 0, 0, 221, 222, 5, 116, 0, 0,
		222, 260, 5, 56, 0, 0, 223, 224, 5, 73, 0, 0, 224, 225, 5, 110, 0, 0, 225,
		226, 5, 116, 0, 0, 226, 227, 5, 49, 0, 0, 227, 260, 5, 54, 0, 0, 228, 229,
		5, 73, 0, 0, 229, 230, 5, 110, 0, 0, 230, 231, 5, 116, 0, 0, 231, 232,
		5, 51, 0, 0, 232, 260, 5, 50, 0, 0, 233, 234, 5, 73, 0, 0, 234, 235, 5,
		110, 0, 0, 235, 236, 5, 116, 0, 0, 236, 237, 5, 54, 0, 0, 237, 260, 5,
		52, 0, 0, 238, 239, 5, 73, 0, 0, 239, 240, 5, 110, 0, 0, 240, 260, 5, 116,
		0, 0, 241, 242, 5, 70, 0, 0, 242, 243, 5, 108, 0, 0, 243, 244, 5, 111,
		0, 0, 244, 245, 5, 97, 0, 0, 245, 246, 5, 116, 0, 0, 246, 247, 5, 51, 0,
		0, 247, 260, 5, 50, 0, 0, 248, 249, 5, 70, 0, 0, 249, 250, 5, 108, 0, 0,
		250, 251, 5, 111, 0, 0, 251, 252, 5, 97, 0, 0, 252, 253, 5, 116, 0, 0,
		253, 254, 5, 54, 0, 0, 254, 260, 5, 52, 0, 0, 255, 256, 5, 66, 0, 0, 256,
		257, 5, 111, 0, 0, 257, 258, 5, 111, 0, 0, 258, 260, 5, 108, 0, 0, 259,
		192, 1, 0, 0, 0, 259, 197, 1, 0, 0, 0, 259, 203, 1, 0, 0, 0, 259, 209,
		1, 0, 0, 0, 259, 215, 1, 0, 0, 0, 259, 219, 1, 0, 0, 0, 259, 223, 1, 0,
		0, 0, 259, 228, 1, 0, 0, 0, 259, 233, 1, 0, 0, 0, 259, 238, 1, 0, 0, 0,
		259, 241, 1, 0, 0, 0, 259, 248, 1, 0, 0, 0, 259, 255, 1, 0, 0, 0, 260,
		54, 1, 0, 0, 0, 261, 266, 3, 57, 28, 0, 262, 266, 3, 61, 30, 0, 263, 266,
		3, 65, 32, 0, 264, 266, 3, 69, 34, 0, 265, 261, 1, 0, 0, 0, 265, 262, 1,
		0, 0, 0, 265, 263, 1, 0, 0, 0, 265, 264, 1, 0, 0, 0, 266, 56, 1, 0, 0,
		0, 267, 285, 3, 59, 29, 0, 268, 270, 7, 0, 0, 0, 269, 268, 1, 0, 0, 0,
		270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272,
		274, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 275, 5, 46, 0, 0, 275, 285,
		3, 59, 29, 0, 276, 277, 3, 59, 29, 0, 277, 281, 5, 46, 0, 0, 278, 280,
		7, 0, 0, 0, 279, 278, 1, 0, 0, 0, 280, 283, 1, 0, 0, 0, 281, 279, 1, 0,
		0, 0, 281, 282, 1, 0, 0, 0, 282, 285, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0,
		284, 267, 1, 0, 0, 0, 284, 271, 1, 0, 0, 0, 284, 276, 1, 0, 0, 0, 285,
		58, 1, 0, 0, 0, 286, 290, 7, 1, 0, 0, 287, 289, 7, 0, 0, 0, 288, 287, 1,
		0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0, 0,
		0, 291, 60, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 294, 5, 48, 0, 0, 294,
		295, 5, 98, 0, 0, 295, 313, 1, 0, 0, 0, 296, 314, 3, 63, 31, 0, 297, 299,
		7, 2, 0, 0, 298, 297, 1, 0, 0, 0, 299, 302, 1, 0, 0, 0, 300, 298, 1, 0,
		0, 0, 300, 301, 1, 0, 0, 0, 301, 303, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0,
		303, 304, 5, 46, 0, 0, 304, 314, 3, 63, 31, 0, 305, 306, 3, 63, 31, 0,
		306, 310, 5, 46, 0, 0, 307, 309, 7, 2, 0, 0, 308, 307, 1, 0, 0, 0, 309,
		312, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 314,
		1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 313, 296, 1, 0, 0, 0, 313, 300, 1, 0,
		0, 0, 313, 305, 1, 0, 0, 0, 314, 62, 1, 0, 0, 0, 315, 317, 7, 2, 0, 0,
		316, 315, 1, 0, 0, 0, 317, 320, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 318,
		319, 1, 0, 0, 0, 319, 321, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 321, 325,
		7, 3, 0, 0, 322, 324, 7, 2, 0, 0, 323, 322, 1, 0, 0, 0, 324, 327, 1, 0,
		0, 0, 325, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 64, 1, 0, 0, 0,
		327, 325, 1, 0, 0, 0, 328, 329, 5, 48, 0, 0, 329, 330, 5, 111, 0, 0, 330,
		348, 1, 0, 0, 0, 331, 349, 3, 67, 33, 0, 332, 334, 7, 4, 0, 0, 333, 332,
		1, 0, 0, 0, 334, 337, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 335, 336, 1, 0,
		0, 0, 336, 338, 1, 0, 0, 0, 337, 335, 1, 0, 0, 0, 338, 339, 5, 46, 0, 0,
		339, 349, 3, 67, 33, 0, 340, 341, 3, 67, 33, 0, 341, 345, 5, 46, 0, 0,
		342, 344, 7, 4, 0, 0, 343, 342, 1, 0, 0, 0, 344, 347, 1, 0, 0, 0, 345,
		343, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 349, 1, 0, 0, 0, 347, 345,
		1, 0, 0, 0, 348, 331, 1, 0, 0, 0, 348, 335, 1, 0, 0, 0, 348, 340, 1, 0,
		0, 0, 349, 66, 1, 0, 0, 0, 350, 352, 7, 4, 0, 0, 351, 350, 1, 0, 0, 0,
		352, 355, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354,
		356, 1, 0, 0, 0, 355, 353, 1, 0, 0, 0, 356, 360, 7, 5, 0, 0, 357, 359,
		7, 4, 0, 0, 358, 357, 1, 0, 0, 0, 359, 362, 1, 0, 0, 0, 360, 358, 1, 0,
		0, 0, 360, 361, 1, 0, 0, 0, 361, 68, 1, 0, 0, 0, 362, 360, 1, 0, 0, 0,
		363, 364, 5, 48, 0, 0, 364, 365, 5, 120, 0, 0, 365, 385, 1, 0, 0, 0, 366,
		386, 3, 71, 35, 0, 367, 370, 3, 73, 36, 0, 368, 370, 5, 95, 0, 0, 369,
		367, 1, 0, 0, 0, 369, 368, 1, 0, 0, 0, 370, 373, 1, 0, 0, 0, 371, 369,
		1, 0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 374, 1, 0, 0, 0, 373, 371, 1, 0,
		0, 0, 374, 375, 5, 46, 0, 0, 375, 386, 3, 71, 35, 0, 376, 377, 3, 71, 35,
		0, 377, 382, 5, 46, 0, 0, 378, 381, 3, 73, 36, 0, 379, 381, 5, 95, 0, 0,
		380, 378, 1, 0, 0, 0, 380, 379, 1, 0, 0, 0, 381, 384, 1, 0, 0, 0, 382,
		380, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 386, 1, 0, 0, 0, 384, 382,
		1, 0, 0, 0, 385, 366, 1, 0, 0, 0, 385, 371, 1, 0, 0, 0, 385, 376, 1, 0,
		0, 0, 386, 70, 1, 0, 0, 0, 387, 390, 3, 73, 36, 0, 388, 390, 5, 95, 0,
		0, 389, 387, 1, 0, 0, 0, 389, 388, 1, 0, 0, 0, 390, 393, 1, 0, 0, 0, 391,
		389, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 394, 1, 0, 0, 0, 393, 391,
		1, 0, 0, 0, 394, 399, 3, 73, 36, 0, 395, 398, 3, 73, 36, 0, 396, 398, 5,
		95, 0, 0, 397, 395, 1, 0, 0, 0, 397, 396, 1, 0, 0, 0, 398, 401, 1, 0, 0,
		0, 399, 397, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 72, 1, 0, 0, 0, 401,
		399, 1, 0, 0, 0, 402, 403, 7, 6, 0, 0, 403, 74, 1, 0, 0, 0, 404, 408, 7,
		7, 0, 0, 405, 407, 7, 8, 0, 0, 406, 405, 1, 0, 0, 0, 407, 410, 1, 0, 0,
		0, 408, 406, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 76, 1, 0, 0, 0, 410,
		408, 1, 0, 0, 0, 411, 412, 5, 38, 0, 0, 412, 78, 1, 0, 0, 0, 413, 414,
		5, 36, 0, 0, 414, 80, 1, 0, 0, 0, 415, 416, 5, 124, 0, 0, 416, 82, 1, 0,
		0, 0, 417, 418, 5, 33, 0, 0, 418, 84, 1, 0, 0, 0, 419, 420, 5, 63, 0, 0,
		420, 86, 1, 0, 0, 0, 421, 422, 5, 35, 0, 0, 422, 88, 1, 0, 0, 0, 423, 424,
		5, 40, 0, 0, 424, 90, 1, 0, 0, 0, 425, 426, 5, 41, 0, 0, 426, 92, 1, 0,
		0, 0, 427, 428, 5, 123, 0, 0, 428, 94, 1, 0, 0, 0, 429, 430, 5, 125, 0,
		0, 430, 96, 1, 0, 0, 0, 431, 432, 5, 91, 0, 0, 432, 98, 1, 0, 0, 0, 433,
		434, 5, 93, 0, 0, 434, 100, 1, 0, 0, 0, 435, 436, 5, 59, 0, 0, 436, 102,
		1, 0, 0, 0, 437, 439, 7, 9, 0, 0, 438, 437, 1, 0, 0, 0, 438, 439, 1, 0,
		0, 0, 439, 440, 1, 0, 0, 0, 440, 442, 7, 10, 0, 0, 441, 438, 1, 0, 0, 0,
		442, 443, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444,
		104, 1, 0, 0, 0, 445, 467, 7, 11, 0, 0, 446, 447, 5, 47, 0, 0, 447, 448,
		5, 42, 0, 0, 448, 452, 1, 0, 0, 0, 449, 451, 9, 0, 0, 0, 450, 449, 1, 0,
		0, 0, 451, 454, 1, 0, 0, 0, 452, 453, 1, 0, 0, 0, 452, 450, 1, 0, 0, 0,
		453, 455, 1, 0, 0, 0, 454, 452, 1, 0, 0, 0, 455, 456, 5, 42, 0, 0, 456,
		467, 5, 47, 0, 0, 457, 458, 5, 47, 0, 0, 458, 459, 5, 47, 0, 0, 459, 463,
		1, 0, 0, 0, 460, 462, 8, 10, 0, 0, 461, 460, 1, 0, 0, 0, 462, 465, 1, 0,
		0, 0, 463, 461, 1, 0, 0, 0, 463, 464, 1, 0, 0, 0, 464, 467, 1, 0, 0, 0,
		465, 463, 1, 0, 0, 0, 466, 445, 1, 0, 0, 0, 466, 446, 1, 0, 0, 0, 466,
		457, 1, 0, 0, 0, 467, 468, 1, 0, 0, 0, 468, 469, 6, 52, 0, 0, 469, 106,
		1, 0, 0, 0, 32, 0, 259, 265, 271, 281, 284, 290, 300, 310, 313, 318, 325,
		335, 345, 348, 353, 360, 369, 371, 380, 382, 385, 389, 391, 397, 399, 408,
		438, 443, 452, 463, 466, 1, 6, 0, 0,
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

// ZenithLexerInit initializes any static state used to implement ZenithLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewZenithLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ZenithLexerInit() {
	staticData := &ZenithLexerLexerStaticData
	staticData.once.Do(zenithlexerLexerInit)
}

// NewZenithLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewZenithLexer(input antlr.CharStream) *ZenithLexer {
	ZenithLexerInit()
	l := new(ZenithLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ZenithLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ZenithLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ZenithLexer tokens.
const (
	ZenithLexerNOT_NULL    = 1
	ZenithLexerBIT_NOT     = 2
	ZenithLexerPOW         = 3
	ZenithLexerTIMES       = 4
	ZenithLexerDIVIDE      = 5
	ZenithLexerREM         = 6
	ZenithLexerPLUS        = 7
	ZenithLexerMINUS       = 8
	ZenithLexerLSHIFT      = 9
	ZenithLexerRSHIFT      = 10
	ZenithLexerLT          = 11
	ZenithLexerGT          = 12
	ZenithLexerLTE         = 13
	ZenithLexerGTE         = 14
	ZenithLexerEQ          = 15
	ZenithLexerNEQ         = 16
	ZenithLexerCOALESCE    = 17
	ZenithLexerNULL        = 18
	ZenithLexerTRUE        = 19
	ZenithLexerALLOC       = 20
	ZenithLexerDEALLOC     = 21
	ZenithLexerFALSE       = 22
	ZenithLexerIF          = 23
	ZenithLexerELSE        = 24
	ZenithLexerASSIGN      = 25
	ZenithLexerINIT_ASSIGN = 26
	ZenithLexerTYPE        = 27
	ZenithLexerNUM         = 28
	ZenithLexerID          = 29
	ZenithLexerAND         = 30
	ZenithLexerDOLLAR      = 31
	ZenithLexerPIPE        = 32
	ZenithLexerEXCLAIM     = 33
	ZenithLexerQUESTION    = 34
	ZenithLexerHASH        = 35
	ZenithLexerLPAREN      = 36
	ZenithLexerRPAREN      = 37
	ZenithLexerLBRACE      = 38
	ZenithLexerRBRACE      = 39
	ZenithLexerLBRACK      = 40
	ZenithLexerRBRACK      = 41
	ZenithLexerSEMICOLON   = 42
	ZenithLexerNL          = 43
	ZenithLexerIGNORED     = 44
)
