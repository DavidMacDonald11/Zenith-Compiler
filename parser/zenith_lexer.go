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
		"", "':='", "'+'", "'-'", "'*'", "'/'", "'%'", "'if'", "'else'", "'true'",
		"'false'", "", "", "", "'('", "')'", "'{'", "'}'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "DEFINE_EQ", "PLUS", "MINUS", "TIMES", "DIVIDE", "REM", "IF", "ELSE",
		"TRUE", "FALSE", "TYPE", "NUM", "ID", "LPAREN", "RPAREN", "LBRACE",
		"RBRACE", "SEMICOLON", "NL", "SPACE",
	}
	staticData.RuleNames = []string{
		"DEFINE_EQ", "PLUS", "MINUS", "TIMES", "DIVIDE", "REM", "IF", "ELSE",
		"TRUE", "FALSE", "TYPE", "NUM", "DEC_NUM", "DEC_SEG", "BIN_NUM", "BIN_SEG",
		"OCT_NUM", "OCT_SEG", "HEX_NUM", "HEX_SEG", "HEX_DIGIT", "ID", "LPAREN",
		"RPAREN", "LBRACE", "RBRACE", "SEMICOLON", "NL", "SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 20, 332, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 159, 8, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 3, 11, 165, 8, 11, 1, 12, 1, 12, 5, 12, 169, 8, 12,
		10, 12, 12, 12, 172, 9, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 179,
		8, 12, 10, 12, 12, 12, 182, 9, 12, 3, 12, 184, 8, 12, 1, 13, 1, 13, 5,
		13, 188, 8, 13, 10, 13, 12, 13, 191, 9, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 5, 14, 198, 8, 14, 10, 14, 12, 14, 201, 9, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 5, 14, 208, 8, 14, 10, 14, 12, 14, 211, 9, 14, 3, 14,
		213, 8, 14, 1, 15, 5, 15, 216, 8, 15, 10, 15, 12, 15, 219, 9, 15, 1, 15,
		1, 15, 5, 15, 223, 8, 15, 10, 15, 12, 15, 226, 9, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 5, 16, 233, 8, 16, 10, 16, 12, 16, 236, 9, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 243, 8, 16, 10, 16, 12, 16, 246, 9,
		16, 3, 16, 248, 8, 16, 1, 17, 5, 17, 251, 8, 17, 10, 17, 12, 17, 254, 9,
		17, 1, 17, 1, 17, 5, 17, 258, 8, 17, 10, 17, 12, 17, 261, 9, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 269, 8, 18, 10, 18, 12, 18, 272,
		9, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 280, 8, 18, 10,
		18, 12, 18, 283, 9, 18, 3, 18, 285, 8, 18, 1, 19, 1, 19, 5, 19, 289, 8,
		19, 10, 19, 12, 19, 292, 9, 19, 1, 19, 1, 19, 1, 19, 5, 19, 297, 8, 19,
		10, 19, 12, 19, 300, 9, 19, 1, 20, 1, 20, 1, 21, 1, 21, 5, 21, 306, 8,
		21, 10, 21, 12, 21, 309, 9, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 3, 27, 322, 8, 27, 1, 27, 4, 27, 325,
		8, 27, 11, 27, 12, 27, 326, 1, 28, 1, 28, 1, 28, 1, 28, 0, 0, 29, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23,
		12, 25, 0, 27, 0, 29, 0, 31, 0, 33, 0, 35, 0, 37, 0, 39, 0, 41, 0, 43,
		13, 45, 14, 47, 15, 49, 16, 51, 17, 53, 18, 55, 19, 57, 20, 1, 0, 12, 2,
		0, 48, 57, 95, 95, 1, 0, 48, 57, 2, 0, 48, 49, 95, 95, 1, 0, 48, 49, 2,
		0, 48, 55, 95, 95, 1, 0, 48, 55, 3, 0, 48, 57, 65, 70, 97, 102, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 13, 13,
		1, 0, 10, 10, 2, 0, 9, 9, 32, 32, 367, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
		1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 1, 59, 1, 0, 0, 0, 3,
		62, 1, 0, 0, 0, 5, 64, 1, 0, 0, 0, 7, 66, 1, 0, 0, 0, 9, 68, 1, 0, 0, 0,
		11, 70, 1, 0, 0, 0, 13, 72, 1, 0, 0, 0, 15, 75, 1, 0, 0, 0, 17, 80, 1,
		0, 0, 0, 19, 85, 1, 0, 0, 0, 21, 158, 1, 0, 0, 0, 23, 164, 1, 0, 0, 0,
		25, 183, 1, 0, 0, 0, 27, 185, 1, 0, 0, 0, 29, 192, 1, 0, 0, 0, 31, 217,
		1, 0, 0, 0, 33, 227, 1, 0, 0, 0, 35, 252, 1, 0, 0, 0, 37, 262, 1, 0, 0,
		0, 39, 290, 1, 0, 0, 0, 41, 301, 1, 0, 0, 0, 43, 303, 1, 0, 0, 0, 45, 310,
		1, 0, 0, 0, 47, 312, 1, 0, 0, 0, 49, 314, 1, 0, 0, 0, 51, 316, 1, 0, 0,
		0, 53, 318, 1, 0, 0, 0, 55, 324, 1, 0, 0, 0, 57, 328, 1, 0, 0, 0, 59, 60,
		5, 58, 0, 0, 60, 61, 5, 61, 0, 0, 61, 2, 1, 0, 0, 0, 62, 63, 5, 43, 0,
		0, 63, 4, 1, 0, 0, 0, 64, 65, 5, 45, 0, 0, 65, 6, 1, 0, 0, 0, 66, 67, 5,
		42, 0, 0, 67, 8, 1, 0, 0, 0, 68, 69, 5, 47, 0, 0, 69, 10, 1, 0, 0, 0, 70,
		71, 5, 37, 0, 0, 71, 12, 1, 0, 0, 0, 72, 73, 5, 105, 0, 0, 73, 74, 5, 102,
		0, 0, 74, 14, 1, 0, 0, 0, 75, 76, 5, 101, 0, 0, 76, 77, 5, 108, 0, 0, 77,
		78, 5, 115, 0, 0, 78, 79, 5, 101, 0, 0, 79, 16, 1, 0, 0, 0, 80, 81, 5,
		116, 0, 0, 81, 82, 5, 114, 0, 0, 82, 83, 5, 117, 0, 0, 83, 84, 5, 101,
		0, 0, 84, 18, 1, 0, 0, 0, 85, 86, 5, 102, 0, 0, 86, 87, 5, 97, 0, 0, 87,
		88, 5, 108, 0, 0, 88, 89, 5, 115, 0, 0, 89, 90, 5, 101, 0, 0, 90, 20, 1,
		0, 0, 0, 91, 92, 5, 117, 0, 0, 92, 93, 5, 105, 0, 0, 93, 94, 5, 110, 0,
		0, 94, 95, 5, 116, 0, 0, 95, 159, 5, 56, 0, 0, 96, 97, 5, 117, 0, 0, 97,
		98, 5, 105, 0, 0, 98, 99, 5, 110, 0, 0, 99, 100, 5, 116, 0, 0, 100, 101,
		5, 49, 0, 0, 101, 159, 5, 54, 0, 0, 102, 103, 5, 117, 0, 0, 103, 104, 5,
		105, 0, 0, 104, 105, 5, 110, 0, 0, 105, 106, 5, 116, 0, 0, 106, 107, 5,
		51, 0, 0, 107, 159, 5, 50, 0, 0, 108, 109, 5, 117, 0, 0, 109, 110, 5, 105,
		0, 0, 110, 111, 5, 110, 0, 0, 111, 112, 5, 116, 0, 0, 112, 113, 5, 54,
		0, 0, 113, 159, 5, 52, 0, 0, 114, 115, 5, 117, 0, 0, 115, 116, 5, 105,
		0, 0, 116, 117, 5, 110, 0, 0, 117, 159, 5, 116, 0, 0, 118, 119, 5, 105,
		0, 0, 119, 120, 5, 110, 0, 0, 120, 121, 5, 116, 0, 0, 121, 159, 5, 56,
		0, 0, 122, 123, 5, 105, 0, 0, 123, 124, 5, 110, 0, 0, 124, 125, 5, 116,
		0, 0, 125, 126, 5, 49, 0, 0, 126, 159, 5, 54, 0, 0, 127, 128, 5, 105, 0,
		0, 128, 129, 5, 110, 0, 0, 129, 130, 5, 116, 0, 0, 130, 131, 5, 51, 0,
		0, 131, 159, 5, 50, 0, 0, 132, 133, 5, 105, 0, 0, 133, 134, 5, 110, 0,
		0, 134, 135, 5, 116, 0, 0, 135, 136, 5, 54, 0, 0, 136, 159, 5, 52, 0, 0,
		137, 138, 5, 105, 0, 0, 138, 139, 5, 110, 0, 0, 139, 159, 5, 116, 0, 0,
		140, 141, 5, 102, 0, 0, 141, 142, 5, 108, 0, 0, 142, 143, 5, 111, 0, 0,
		143, 144, 5, 97, 0, 0, 144, 145, 5, 116, 0, 0, 145, 146, 5, 51, 0, 0, 146,
		159, 5, 50, 0, 0, 147, 148, 5, 102, 0, 0, 148, 149, 5, 108, 0, 0, 149,
		150, 5, 111, 0, 0, 150, 151, 5, 97, 0, 0, 151, 152, 5, 116, 0, 0, 152,
		153, 5, 54, 0, 0, 153, 159, 5, 52, 0, 0, 154, 155, 5, 98, 0, 0, 155, 156,
		5, 111, 0, 0, 156, 157, 5, 111, 0, 0, 157, 159, 5, 108, 0, 0, 158, 91,
		1, 0, 0, 0, 158, 96, 1, 0, 0, 0, 158, 102, 1, 0, 0, 0, 158, 108, 1, 0,
		0, 0, 158, 114, 1, 0, 0, 0, 158, 118, 1, 0, 0, 0, 158, 122, 1, 0, 0, 0,
		158, 127, 1, 0, 0, 0, 158, 132, 1, 0, 0, 0, 158, 137, 1, 0, 0, 0, 158,
		140, 1, 0, 0, 0, 158, 147, 1, 0, 0, 0, 158, 154, 1, 0, 0, 0, 159, 22, 1,
		0, 0, 0, 160, 165, 3, 25, 12, 0, 161, 165, 3, 29, 14, 0, 162, 165, 3, 33,
		16, 0, 163, 165, 3, 37, 18, 0, 164, 160, 1, 0, 0, 0, 164, 161, 1, 0, 0,
		0, 164, 162, 1, 0, 0, 0, 164, 163, 1, 0, 0, 0, 165, 24, 1, 0, 0, 0, 166,
		184, 3, 27, 13, 0, 167, 169, 7, 0, 0, 0, 168, 167, 1, 0, 0, 0, 169, 172,
		1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 173, 1, 0,
		0, 0, 172, 170, 1, 0, 0, 0, 173, 174, 5, 46, 0, 0, 174, 184, 3, 27, 13,
		0, 175, 176, 3, 27, 13, 0, 176, 180, 5, 46, 0, 0, 177, 179, 7, 0, 0, 0,
		178, 177, 1, 0, 0, 0, 179, 182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180,
		181, 1, 0, 0, 0, 181, 184, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 166,
		1, 0, 0, 0, 183, 170, 1, 0, 0, 0, 183, 175, 1, 0, 0, 0, 184, 26, 1, 0,
		0, 0, 185, 189, 7, 1, 0, 0, 186, 188, 7, 0, 0, 0, 187, 186, 1, 0, 0, 0,
		188, 191, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190,
		28, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 192, 193, 5, 48, 0, 0, 193, 194,
		5, 98, 0, 0, 194, 212, 1, 0, 0, 0, 195, 213, 3, 31, 15, 0, 196, 198, 7,
		2, 0, 0, 197, 196, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0,
		0, 199, 200, 1, 0, 0, 0, 200, 202, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202,
		203, 5, 46, 0, 0, 203, 213, 3, 31, 15, 0, 204, 205, 3, 31, 15, 0, 205,
		209, 5, 46, 0, 0, 206, 208, 7, 2, 0, 0, 207, 206, 1, 0, 0, 0, 208, 211,
		1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 213, 1, 0,
		0, 0, 211, 209, 1, 0, 0, 0, 212, 195, 1, 0, 0, 0, 212, 199, 1, 0, 0, 0,
		212, 204, 1, 0, 0, 0, 213, 30, 1, 0, 0, 0, 214, 216, 7, 2, 0, 0, 215, 214,
		1, 0, 0, 0, 216, 219, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 217, 218, 1, 0,
		0, 0, 218, 220, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 220, 224, 7, 3, 0, 0,
		221, 223, 7, 2, 0, 0, 222, 221, 1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224,
		222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 32, 1, 0, 0, 0, 226, 224, 1,
		0, 0, 0, 227, 228, 5, 48, 0, 0, 228, 229, 5, 111, 0, 0, 229, 247, 1, 0,
		0, 0, 230, 248, 3, 35, 17, 0, 231, 233, 7, 4, 0, 0, 232, 231, 1, 0, 0,
		0, 233, 236, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235,
		237, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 237, 238, 5, 46, 0, 0, 238, 248,
		3, 35, 17, 0, 239, 240, 3, 35, 17, 0, 240, 244, 5, 46, 0, 0, 241, 243,
		7, 4, 0, 0, 242, 241, 1, 0, 0, 0, 243, 246, 1, 0, 0, 0, 244, 242, 1, 0,
		0, 0, 244, 245, 1, 0, 0, 0, 245, 248, 1, 0, 0, 0, 246, 244, 1, 0, 0, 0,
		247, 230, 1, 0, 0, 0, 247, 234, 1, 0, 0, 0, 247, 239, 1, 0, 0, 0, 248,
		34, 1, 0, 0, 0, 249, 251, 7, 4, 0, 0, 250, 249, 1, 0, 0, 0, 251, 254, 1,
		0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 255, 1, 0, 0,
		0, 254, 252, 1, 0, 0, 0, 255, 259, 7, 5, 0, 0, 256, 258, 7, 4, 0, 0, 257,
		256, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259, 260,
		1, 0, 0, 0, 260, 36, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 263, 5, 48,
		0, 0, 263, 264, 5, 120, 0, 0, 264, 284, 1, 0, 0, 0, 265, 285, 3, 39, 19,
		0, 266, 269, 3, 41, 20, 0, 267, 269, 5, 95, 0, 0, 268, 266, 1, 0, 0, 0,
		268, 267, 1, 0, 0, 0, 269, 272, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 270,
		271, 1, 0, 0, 0, 271, 273, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 273, 274,
		5, 46, 0, 0, 274, 285, 3, 39, 19, 0, 275, 276, 3, 39, 19, 0, 276, 281,
		5, 46, 0, 0, 277, 280, 3, 41, 20, 0, 278, 280, 5, 95, 0, 0, 279, 277, 1,
		0, 0, 0, 279, 278, 1, 0, 0, 0, 280, 283, 1, 0, 0, 0, 281, 279, 1, 0, 0,
		0, 281, 282, 1, 0, 0, 0, 282, 285, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 284,
		265, 1, 0, 0, 0, 284, 270, 1, 0, 0, 0, 284, 275, 1, 0, 0, 0, 285, 38, 1,
		0, 0, 0, 286, 289, 3, 41, 20, 0, 287, 289, 5, 95, 0, 0, 288, 286, 1, 0,
		0, 0, 288, 287, 1, 0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0,
		290, 291, 1, 0, 0, 0, 291, 293, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293,
		298, 3, 41, 20, 0, 294, 297, 3, 41, 20, 0, 295, 297, 5, 95, 0, 0, 296,
		294, 1, 0, 0, 0, 296, 295, 1, 0, 0, 0, 297, 300, 1, 0, 0, 0, 298, 296,
		1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 40, 1, 0, 0, 0, 300, 298, 1, 0,
		0, 0, 301, 302, 7, 6, 0, 0, 302, 42, 1, 0, 0, 0, 303, 307, 7, 7, 0, 0,
		304, 306, 7, 8, 0, 0, 305, 304, 1, 0, 0, 0, 306, 309, 1, 0, 0, 0, 307,
		305, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 44, 1, 0, 0, 0, 309, 307, 1,
		0, 0, 0, 310, 311, 5, 40, 0, 0, 311, 46, 1, 0, 0, 0, 312, 313, 5, 41, 0,
		0, 313, 48, 1, 0, 0, 0, 314, 315, 5, 123, 0, 0, 315, 50, 1, 0, 0, 0, 316,
		317, 5, 125, 0, 0, 317, 52, 1, 0, 0, 0, 318, 319, 5, 59, 0, 0, 319, 54,
		1, 0, 0, 0, 320, 322, 7, 9, 0, 0, 321, 320, 1, 0, 0, 0, 321, 322, 1, 0,
		0, 0, 322, 323, 1, 0, 0, 0, 323, 325, 7, 10, 0, 0, 324, 321, 1, 0, 0, 0,
		325, 326, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327,
		56, 1, 0, 0, 0, 328, 329, 7, 11, 0, 0, 329, 330, 1, 0, 0, 0, 330, 331,
		6, 28, 0, 0, 331, 58, 1, 0, 0, 0, 29, 0, 158, 164, 170, 180, 183, 189,
		199, 209, 212, 217, 224, 234, 244, 247, 252, 259, 268, 270, 279, 281, 284,
		288, 290, 296, 298, 307, 321, 326, 1, 6, 0, 0,
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
	ZenithLexerDEFINE_EQ = 1
	ZenithLexerPLUS      = 2
	ZenithLexerMINUS     = 3
	ZenithLexerTIMES     = 4
	ZenithLexerDIVIDE    = 5
	ZenithLexerREM       = 6
	ZenithLexerIF        = 7
	ZenithLexerELSE      = 8
	ZenithLexerTRUE      = 9
	ZenithLexerFALSE     = 10
	ZenithLexerTYPE      = 11
	ZenithLexerNUM       = 12
	ZenithLexerID        = 13
	ZenithLexerLPAREN    = 14
	ZenithLexerRPAREN    = 15
	ZenithLexerLBRACE    = 16
	ZenithLexerRBRACE    = 17
	ZenithLexerSEMICOLON = 18
	ZenithLexerNL        = 19
	ZenithLexerSPACE     = 20
)
