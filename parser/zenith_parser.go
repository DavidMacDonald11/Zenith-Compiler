// Code generated from ZenithParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ZenithParser

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

type ZenithParser struct {
	*antlr.BaseParser
}

var ZenithParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func zenithparserParserInit() {
	staticData := &ZenithParserParserStaticData
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
		"fileStat", "endedStat", "lineEnd", "stat", "type", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 211, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 4, 0, 14, 8, 0, 11, 0, 12, 0, 15, 1, 0, 3, 0, 19,
		8, 0, 1, 1, 3, 1, 22, 8, 1, 1, 1, 1, 1, 1, 1, 1, 2, 4, 2, 28, 8, 2, 11,
		2, 12, 2, 29, 1, 2, 3, 2, 33, 8, 2, 1, 3, 1, 3, 3, 3, 37, 8, 3, 1, 3, 1,
		3, 3, 3, 41, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3, 46, 8, 3, 1, 3, 5, 3, 49, 8,
		3, 10, 3, 12, 3, 52, 9, 3, 1, 3, 3, 3, 55, 8, 3, 1, 3, 3, 3, 58, 8, 3,
		1, 3, 1, 3, 3, 3, 62, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 69, 8,
		4, 1, 4, 5, 4, 72, 8, 4, 10, 4, 12, 4, 75, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 83, 8, 5, 1, 5, 1, 5, 3, 5, 87, 8, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 3, 5, 95, 8, 5, 1, 5, 1, 5, 3, 5, 99, 8, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 106, 8, 5, 1, 5, 1, 5, 3, 5, 110, 8, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 117, 8, 5, 1, 5, 1, 5, 3, 5, 121, 8, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 127, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 132, 8,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 138, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3,
		5, 144, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 150, 8, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 3, 5, 156, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 162, 8, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 3, 5, 168, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 177, 8, 5, 1, 5, 1, 5, 3, 5, 181, 8, 5, 1, 5, 1, 5, 3, 5, 185,
		8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 191, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 199, 8, 5, 1, 5, 4, 5, 202, 8, 5, 11, 5, 12, 5, 203,
		5, 5, 206, 8, 5, 10, 5, 12, 5, 209, 9, 5, 1, 5, 0, 2, 8, 10, 6, 0, 2, 4,
		6, 8, 10, 0, 8, 1, 0, 42, 43, 1, 0, 33, 34, 2, 0, 18, 19, 22, 22, 2, 0,
		7, 8, 33, 33, 1, 0, 4, 6, 1, 0, 7, 8, 1, 0, 9, 10, 1, 0, 11, 16, 259, 0,
		18, 1, 0, 0, 0, 2, 21, 1, 0, 0, 0, 4, 32, 1, 0, 0, 0, 6, 61, 1, 0, 0, 0,
		8, 63, 1, 0, 0, 0, 10, 126, 1, 0, 0, 0, 12, 14, 3, 2, 1, 0, 13, 12, 1,
		0, 0, 0, 14, 15, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16,
		19, 1, 0, 0, 0, 17, 19, 5, 0, 0, 1, 18, 13, 1, 0, 0, 0, 18, 17, 1, 0, 0,
		0, 19, 1, 1, 0, 0, 0, 20, 22, 5, 43, 0, 0, 21, 20, 1, 0, 0, 0, 21, 22,
		1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 24, 3, 6, 3, 0, 24, 25, 3, 4, 2, 0,
		25, 3, 1, 0, 0, 0, 26, 28, 7, 0, 0, 0, 27, 26, 1, 0, 0, 0, 28, 29, 1, 0,
		0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 33, 1, 0, 0, 0, 31, 33,
		5, 0, 0, 1, 32, 27, 1, 0, 0, 0, 32, 31, 1, 0, 0, 0, 33, 5, 1, 0, 0, 0,
		34, 36, 5, 29, 0, 0, 35, 37, 3, 8, 4, 0, 36, 35, 1, 0, 0, 0, 36, 37, 1,
		0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40, 5, 26, 0, 0, 39, 41, 5, 43, 0, 0,
		40, 39, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 62, 3,
		10, 5, 0, 43, 45, 5, 38, 0, 0, 44, 46, 5, 43, 0, 0, 45, 44, 1, 0, 0, 0,
		45, 46, 1, 0, 0, 0, 46, 50, 1, 0, 0, 0, 47, 49, 3, 2, 1, 0, 48, 47, 1,
		0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51,
		54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 55, 3, 6, 3, 0, 54, 53, 1, 0, 0,
		0, 54, 55, 1, 0, 0, 0, 55, 57, 1, 0, 0, 0, 56, 58, 3, 4, 2, 0, 57, 56,
		1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 62, 5, 39, 0, 0,
		60, 62, 3, 10, 5, 0, 61, 34, 1, 0, 0, 0, 61, 43, 1, 0, 0, 0, 61, 60, 1,
		0, 0, 0, 62, 7, 1, 0, 0, 0, 63, 64, 6, 4, -1, 0, 64, 65, 5, 27, 0, 0, 65,
		73, 1, 0, 0, 0, 66, 68, 10, 1, 0, 0, 67, 69, 5, 35, 0, 0, 68, 67, 1, 0,
		0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 72, 7, 1, 0, 0, 71, 66,
		1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0,
		74, 9, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 77, 6, 5, -1, 0, 77, 127, 5,
		28, 0, 0, 78, 127, 5, 29, 0, 0, 79, 127, 7, 2, 0, 0, 80, 82, 5, 36, 0,
		0, 81, 83, 5, 43, 0, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84,
		1, 0, 0, 0, 84, 86, 3, 10, 5, 0, 85, 87, 5, 43, 0, 0, 86, 85, 1, 0, 0,
		0, 86, 87, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 5, 37, 0, 0, 89, 127,
		1, 0, 0, 0, 90, 91, 5, 20, 0, 0, 91, 92, 7, 1, 0, 0, 92, 94, 5, 36, 0,
		0, 93, 95, 5, 43, 0, 0, 94, 93, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 96,
		1, 0, 0, 0, 96, 98, 3, 10, 5, 0, 97, 99, 5, 43, 0, 0, 98, 97, 1, 0, 0,
		0, 98, 99, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 5, 37, 0, 0, 101,
		127, 1, 0, 0, 0, 102, 103, 5, 21, 0, 0, 103, 105, 5, 36, 0, 0, 104, 106,
		5, 43, 0, 0, 105, 104, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 1, 0,
		0, 0, 107, 109, 3, 10, 5, 0, 108, 110, 5, 43, 0, 0, 109, 108, 1, 0, 0,
		0, 109, 110, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 5, 37, 0, 0, 112,
		127, 1, 0, 0, 0, 113, 114, 5, 27, 0, 0, 114, 116, 5, 36, 0, 0, 115, 117,
		5, 43, 0, 0, 116, 115, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 1, 0,
		0, 0, 118, 120, 3, 10, 5, 0, 119, 121, 5, 43, 0, 0, 120, 119, 1, 0, 0,
		0, 120, 121, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123, 5, 37, 0, 0, 123,
		127, 1, 0, 0, 0, 124, 125, 7, 3, 0, 0, 125, 127, 3, 10, 5, 11, 126, 76,
		1, 0, 0, 0, 126, 78, 1, 0, 0, 0, 126, 79, 1, 0, 0, 0, 126, 80, 1, 0, 0,
		0, 126, 90, 1, 0, 0, 0, 126, 102, 1, 0, 0, 0, 126, 113, 1, 0, 0, 0, 126,
		124, 1, 0, 0, 0, 127, 207, 1, 0, 0, 0, 128, 129, 10, 9, 0, 0, 129, 131,
		7, 4, 0, 0, 130, 132, 5, 43, 0, 0, 131, 130, 1, 0, 0, 0, 131, 132, 1, 0,
		0, 0, 132, 133, 1, 0, 0, 0, 133, 206, 3, 10, 5, 10, 134, 135, 10, 8, 0,
		0, 135, 137, 7, 5, 0, 0, 136, 138, 5, 43, 0, 0, 137, 136, 1, 0, 0, 0, 137,
		138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 206, 3, 10, 5, 9, 140, 141,
		10, 7, 0, 0, 141, 143, 7, 6, 0, 0, 142, 144, 5, 43, 0, 0, 143, 142, 1,
		0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 206, 3, 10, 5,
		8, 146, 147, 10, 6, 0, 0, 147, 149, 5, 30, 0, 0, 148, 150, 5, 43, 0, 0,
		149, 148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151,
		206, 3, 10, 5, 7, 152, 153, 10, 5, 0, 0, 153, 155, 5, 31, 0, 0, 154, 156,
		5, 43, 0, 0, 155, 154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1, 0,
		0, 0, 157, 206, 3, 10, 5, 6, 158, 159, 10, 4, 0, 0, 159, 161, 5, 32, 0,
		0, 160, 162, 5, 43, 0, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162,
		163, 1, 0, 0, 0, 163, 206, 3, 10, 5, 5, 164, 165, 10, 3, 0, 0, 165, 167,
		7, 7, 0, 0, 166, 168, 5, 43, 0, 0, 167, 166, 1, 0, 0, 0, 167, 168, 1, 0,
		0, 0, 168, 169, 1, 0, 0, 0, 169, 206, 3, 10, 5, 4, 170, 171, 10, 2, 0,
		0, 171, 172, 5, 17, 0, 0, 172, 206, 3, 10, 5, 3, 173, 174, 10, 1, 0, 0,
		174, 176, 5, 23, 0, 0, 175, 177, 5, 43, 0, 0, 176, 175, 1, 0, 0, 0, 176,
		177, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 180, 3, 10, 5, 0, 179, 181,
		5, 43, 0, 0, 180, 179, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 182, 1, 0,
		0, 0, 182, 184, 5, 24, 0, 0, 183, 185, 5, 43, 0, 0, 184, 183, 1, 0, 0,
		0, 184, 185, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 187, 3, 10, 5, 2, 187,
		206, 1, 0, 0, 0, 188, 190, 10, 16, 0, 0, 189, 191, 5, 35, 0, 0, 190, 189,
		1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 206, 7, 1,
		0, 0, 193, 194, 10, 15, 0, 0, 194, 206, 5, 1, 0, 0, 195, 201, 10, 10, 0,
		0, 196, 198, 5, 3, 0, 0, 197, 199, 5, 43, 0, 0, 198, 197, 1, 0, 0, 0, 198,
		199, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 202, 3, 10, 5, 0, 201, 196,
		1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0,
		0, 0, 204, 206, 1, 0, 0, 0, 205, 128, 1, 0, 0, 0, 205, 134, 1, 0, 0, 0,
		205, 140, 1, 0, 0, 0, 205, 146, 1, 0, 0, 0, 205, 152, 1, 0, 0, 0, 205,
		158, 1, 0, 0, 0, 205, 164, 1, 0, 0, 0, 205, 170, 1, 0, 0, 0, 205, 173,
		1, 0, 0, 0, 205, 188, 1, 0, 0, 0, 205, 193, 1, 0, 0, 0, 205, 195, 1, 0,
		0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0,
		208, 11, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 38, 15, 18, 21, 29, 32, 36,
		40, 45, 50, 54, 57, 61, 68, 73, 82, 86, 94, 98, 105, 109, 116, 120, 126,
		131, 137, 143, 149, 155, 161, 167, 176, 180, 184, 190, 198, 203, 205, 207,
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

// ZenithParserInit initializes any static state used to implement ZenithParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewZenithParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ZenithParserInit() {
	staticData := &ZenithParserParserStaticData
	staticData.once.Do(zenithparserParserInit)
}

// NewZenithParser produces a new parser instance for the optional input antlr.TokenStream.
func NewZenithParser(input antlr.TokenStream) *ZenithParser {
	ZenithParserInit()
	this := new(ZenithParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ZenithParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ZenithParser.g4"

	return this
}

// ZenithParser tokens.
const (
	ZenithParserEOF         = antlr.TokenEOF
	ZenithParserNOT_NULL    = 1
	ZenithParserBIT_NOT     = 2
	ZenithParserPOW         = 3
	ZenithParserTIMES       = 4
	ZenithParserDIVIDE      = 5
	ZenithParserREM         = 6
	ZenithParserPLUS        = 7
	ZenithParserMINUS       = 8
	ZenithParserLSHIFT      = 9
	ZenithParserRSHIFT      = 10
	ZenithParserLT          = 11
	ZenithParserGT          = 12
	ZenithParserLTE         = 13
	ZenithParserGTE         = 14
	ZenithParserEQ          = 15
	ZenithParserNEQ         = 16
	ZenithParserCOALESCE    = 17
	ZenithParserNULL        = 18
	ZenithParserTRUE        = 19
	ZenithParserALLOC       = 20
	ZenithParserDEALLOC     = 21
	ZenithParserFALSE       = 22
	ZenithParserIF          = 23
	ZenithParserELSE        = 24
	ZenithParserASSIGN      = 25
	ZenithParserINIT_ASSIGN = 26
	ZenithParserTYPE        = 27
	ZenithParserNUM         = 28
	ZenithParserID          = 29
	ZenithParserAND         = 30
	ZenithParserDOLLAR      = 31
	ZenithParserPIPE        = 32
	ZenithParserEXCLAIM     = 33
	ZenithParserQUESTION    = 34
	ZenithParserHASH        = 35
	ZenithParserLPAREN      = 36
	ZenithParserRPAREN      = 37
	ZenithParserLBRACE      = 38
	ZenithParserRBRACE      = 39
	ZenithParserLBRACK      = 40
	ZenithParserRBRACK      = 41
	ZenithParserSEMICOLON   = 42
	ZenithParserNL          = 43
	ZenithParserIGNORED     = 44
)

// ZenithParser rules.
const (
	ZenithParserRULE_fileStat  = 0
	ZenithParserRULE_endedStat = 1
	ZenithParserRULE_lineEnd   = 2
	ZenithParserRULE_stat      = 3
	ZenithParserRULE_type      = 4
	ZenithParserRULE_expr      = 5
)

// IFileStatContext is an interface to support dynamic dispatch.
type IFileStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEndedStat() []IEndedStatContext
	EndedStat(i int) IEndedStatContext
	EOF() antlr.TerminalNode

	// IsFileStatContext differentiates from other interfaces.
	IsFileStatContext()
}

type FileStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileStatContext() *FileStatContext {
	var p = new(FileStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_fileStat
	return p
}

func InitEmptyFileStatContext(p *FileStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_fileStat
}

func (*FileStatContext) IsFileStatContext() {}

func NewFileStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileStatContext {
	var p = new(FileStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZenithParserRULE_fileStat

	return p
}

func (s *FileStatContext) GetParser() antlr.Parser { return s.parser }

func (s *FileStatContext) AllEndedStat() []IEndedStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEndedStatContext); ok {
			len++
		}
	}

	tst := make([]IEndedStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEndedStatContext); ok {
			tst[i] = t.(IEndedStatContext)
			i++
		}
	}

	return tst
}

func (s *FileStatContext) EndedStat(i int) IEndedStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndedStatContext); ok {
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

	return t.(IEndedStatContext)
}

func (s *FileStatContext) EOF() antlr.TerminalNode {
	return s.GetToken(ZenithParserEOF, 0)
}

func (s *FileStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterFileStat(s)
	}
}

func (s *FileStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitFileStat(s)
	}
}

func (s *FileStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitFileStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZenithParser) FileStat() (localctx IFileStatContext) {
	localctx = NewFileStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ZenithParserRULE_fileStat)
	var _la int

	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ZenithParserPLUS, ZenithParserMINUS, ZenithParserNULL, ZenithParserTRUE, ZenithParserALLOC, ZenithParserDEALLOC, ZenithParserFALSE, ZenithParserTYPE, ZenithParserNUM, ZenithParserID, ZenithParserEXCLAIM, ZenithParserLPAREN, ZenithParserLBRACE, ZenithParserNL:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9149227991424) != 0) {
			{
				p.SetState(12)
				p.EndedStat()
			}

			p.SetState(15)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case ZenithParserEOF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(17)
			p.Match(ZenithParserEOF)
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

// IEndedStatContext is an interface to support dynamic dispatch.
type IEndedStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Stat() IStatContext
	LineEnd() ILineEndContext
	NL() antlr.TerminalNode

	// IsEndedStatContext differentiates from other interfaces.
	IsEndedStatContext()
}

type EndedStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndedStatContext() *EndedStatContext {
	var p = new(EndedStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_endedStat
	return p
}

func InitEmptyEndedStatContext(p *EndedStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_endedStat
}

func (*EndedStatContext) IsEndedStatContext() {}

func NewEndedStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndedStatContext {
	var p = new(EndedStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZenithParserRULE_endedStat

	return p
}

func (s *EndedStatContext) GetParser() antlr.Parser { return s.parser }

func (s *EndedStatContext) Stat() IStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *EndedStatContext) LineEnd() ILineEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineEndContext)
}

func (s *EndedStatContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *EndedStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndedStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndedStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterEndedStat(s)
	}
}

func (s *EndedStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitEndedStat(s)
	}
}

func (s *EndedStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitEndedStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZenithParser) EndedStat() (localctx IEndedStatContext) {
	localctx = NewEndedStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ZenithParserRULE_endedStat)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ZenithParserNL {
		{
			p.SetState(20)
			p.Match(ZenithParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(23)
		p.Stat()
	}
	{
		p.SetState(24)
		p.LineEnd()
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

// ILineEndContext is an interface to support dynamic dispatch.
type ILineEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	EOF() antlr.TerminalNode

	// IsLineEndContext differentiates from other interfaces.
	IsLineEndContext()
}

type LineEndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineEndContext() *LineEndContext {
	var p = new(LineEndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_lineEnd
	return p
}

func InitEmptyLineEndContext(p *LineEndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_lineEnd
}

func (*LineEndContext) IsLineEndContext() {}

func NewLineEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineEndContext {
	var p = new(LineEndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZenithParserRULE_lineEnd

	return p
}

func (s *LineEndContext) GetParser() antlr.Parser { return s.parser }

func (s *LineEndContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ZenithParserNL)
}

func (s *LineEndContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, i)
}

func (s *LineEndContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(ZenithParserSEMICOLON)
}

func (s *LineEndContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(ZenithParserSEMICOLON, i)
}

func (s *LineEndContext) EOF() antlr.TerminalNode {
	return s.GetToken(ZenithParserEOF, 0)
}

func (s *LineEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterLineEnd(s)
	}
}

func (s *LineEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitLineEnd(s)
	}
}

func (s *LineEndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitLineEnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZenithParser) LineEnd() (localctx ILineEndContext) {
	localctx = NewLineEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ZenithParserRULE_lineEnd)
	var _la int

	var _alt int

	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ZenithParserSEMICOLON, ZenithParserNL:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(26)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ZenithParserSEMICOLON || _la == ZenithParserNL) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(29)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case ZenithParserEOF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.Match(ZenithParserEOF)
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

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_stat
	return p
}

func InitEmptyStatContext(p *StatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_stat
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZenithParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) CopyAll(ctx *StatContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MultiStatContext struct {
	StatContext
}

func NewMultiStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiStatContext {
	var p = new(MultiStatContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *MultiStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiStatContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZenithParserLBRACE, 0)
}

func (s *MultiStatContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZenithParserRBRACE, 0)
}

func (s *MultiStatContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *MultiStatContext) AllEndedStat() []IEndedStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEndedStatContext); ok {
			len++
		}
	}

	tst := make([]IEndedStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEndedStatContext); ok {
			tst[i] = t.(IEndedStatContext)
			i++
		}
	}

	return tst
}

func (s *MultiStatContext) EndedStat(i int) IEndedStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndedStatContext); ok {
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

	return t.(IEndedStatContext)
}

func (s *MultiStatContext) Stat() IStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *MultiStatContext) LineEnd() ILineEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineEndContext)
}

func (s *MultiStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterMultiStat(s)
	}
}

func (s *MultiStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitMultiStat(s)
	}
}

func (s *MultiStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitMultiStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprStatContext struct {
	StatContext
}

func NewExprStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprStatContext {
	var p = new(ExprStatContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *ExprStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStatContext) Expr() IExprContext {
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

func (s *ExprStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterExprStat(s)
	}
}

func (s *ExprStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitExprStat(s)
	}
}

func (s *ExprStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitExprStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type DefineStatContext struct {
	StatContext
}

func NewDefineStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefineStatContext {
	var p = new(DefineStatContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *DefineStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineStatContext) ID() antlr.TerminalNode {
	return s.GetToken(ZenithParserID, 0)
}

func (s *DefineStatContext) INIT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(ZenithParserINIT_ASSIGN, 0)
}

func (s *DefineStatContext) Expr() IExprContext {
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

func (s *DefineStatContext) Type_() ITypeContext {
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

func (s *DefineStatContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *DefineStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterDefineStat(s)
	}
}

func (s *DefineStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitDefineStat(s)
	}
}

func (s *DefineStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitDefineStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZenithParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ZenithParserRULE_stat)
	var _la int

	var _alt int

	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDefineStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(ZenithParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserTYPE {
			{
				p.SetState(35)
				p.type_(0)
			}

		}
		{
			p.SetState(38)
			p.Match(ZenithParserINIT_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(39)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(42)
			p.expr(0)
		}

	case 2:
		localctx = NewMultiStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Match(ZenithParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(44)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(50)
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
					p.SetState(47)
					p.EndedStat()
				}

			}
			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&353134969216) != 0 {
			{
				p.SetState(53)
				p.Stat()
			}

		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la - -1)) & ^0x3f) == 0 && ((int64(1)<<(_la - -1))&26388279066625) != 0 {
			{
				p.SetState(56)
				p.LineEnd()
			}

		}
		{
			p.SetState(59)
			p.Match(ZenithParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewExprStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = ZenithParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZenithParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) CopyAll(ctx *TypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RefTypeContext struct {
	TypeContext
	Ref antlr.Token
}

func NewRefTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RefTypeContext {
	var p = new(RefTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *RefTypeContext) GetRef() antlr.Token { return s.Ref }

func (s *RefTypeContext) SetRef(v antlr.Token) { s.Ref = v }

func (s *RefTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RefTypeContext) Type_() ITypeContext {
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

func (s *RefTypeContext) EXCLAIM() antlr.TerminalNode {
	return s.GetToken(ZenithParserEXCLAIM, 0)
}

func (s *RefTypeContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(ZenithParserQUESTION, 0)
}

func (s *RefTypeContext) HASH() antlr.TerminalNode {
	return s.GetToken(ZenithParserHASH, 0)
}

func (s *RefTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterRefType(s)
	}
}

func (s *RefTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitRefType(s)
	}
}

func (s *RefTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitRefType(s)

	default:
		return t.VisitChildren(s)
	}
}

type BaseTypeContext struct {
	TypeContext
}

func NewBaseTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BaseTypeContext {
	var p = new(BaseTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *BaseTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseTypeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(ZenithParserTYPE, 0)
}

func (s *BaseTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterBaseType(s)
	}
}

func (s *BaseTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitBaseType(s)
	}
}

func (s *BaseTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitBaseType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZenithParser) Type_() (localctx ITypeContext) {
	return p.type_(0)
}

func (p *ZenithParser) type_(_p int) (localctx ITypeContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewTypeContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITypeContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, ZenithParserRULE_type, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewBaseTypeContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(64)
		p.Match(ZenithParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRefTypeContext(p, NewTypeContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_type)
			p.SetState(66)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(68)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == ZenithParserHASH {
				{
					p.SetState(67)
					p.Match(ZenithParserHASH)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(70)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*RefTypeContext).Ref = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == ZenithParserEXCLAIM || _la == ZenithParserQUESTION) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*RefTypeContext).Ref = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
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
	p.RuleIndex = ZenithParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZenithParserRULE_expr

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

type CastExprContext struct {
	ExprContext
}

func NewCastExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CastExprContext {
	var p = new(CastExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CastExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastExprContext) TYPE() antlr.TerminalNode {
	return s.GetToken(ZenithParserTYPE, 0)
}

func (s *CastExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserLPAREN, 0)
}

func (s *CastExprContext) Expr() IExprContext {
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

func (s *CastExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserRPAREN, 0)
}

func (s *CastExprContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ZenithParserNL)
}

func (s *CastExprContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, i)
}

func (s *CastExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterCastExpr(s)
	}
}

func (s *CastExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitCastExpr(s)
	}
}

func (s *CastExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitCastExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AllocExprContext struct {
	ExprContext
	Ref antlr.Token
}

func NewAllocExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllocExprContext {
	var p = new(AllocExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AllocExprContext) GetRef() antlr.Token { return s.Ref }

func (s *AllocExprContext) SetRef(v antlr.Token) { s.Ref = v }

func (s *AllocExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllocExprContext) ALLOC() antlr.TerminalNode {
	return s.GetToken(ZenithParserALLOC, 0)
}

func (s *AllocExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserLPAREN, 0)
}

func (s *AllocExprContext) Expr() IExprContext {
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

func (s *AllocExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserRPAREN, 0)
}

func (s *AllocExprContext) EXCLAIM() antlr.TerminalNode {
	return s.GetToken(ZenithParserEXCLAIM, 0)
}

func (s *AllocExprContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(ZenithParserQUESTION, 0)
}

func (s *AllocExprContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ZenithParserNL)
}

func (s *AllocExprContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, i)
}

func (s *AllocExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterAllocExpr(s)
	}
}

func (s *AllocExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitAllocExpr(s)
	}
}

func (s *AllocExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitAllocExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type KeyExprContext struct {
	ExprContext
	Key antlr.Token
}

func NewKeyExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KeyExprContext {
	var p = new(KeyExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *KeyExprContext) GetKey() antlr.Token { return s.Key }

func (s *KeyExprContext) SetKey(v antlr.Token) { s.Key = v }

func (s *KeyExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyExprContext) TRUE() antlr.TerminalNode {
	return s.GetToken(ZenithParserTRUE, 0)
}

func (s *KeyExprContext) FALSE() antlr.TerminalNode {
	return s.GetToken(ZenithParserFALSE, 0)
}

func (s *KeyExprContext) NULL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNULL, 0)
}

func (s *KeyExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterKeyExpr(s)
	}
}

func (s *KeyExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitKeyExpr(s)
	}
}

func (s *KeyExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitKeyExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumExprContext struct {
	ExprContext
}

func NewNumExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumExprContext {
	var p = new(NumExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NumExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumExprContext) NUM() antlr.TerminalNode {
	return s.GetToken(ZenithParserNUM, 0)
}

func (s *NumExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterNumExpr(s)
	}
}

func (s *NumExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitNumExpr(s)
	}
}

func (s *NumExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitNumExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	ExprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserLPAREN, 0)
}

func (s *ParenExprContext) Expr() IExprContext {
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

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserRPAREN, 0)
}

func (s *ParenExprContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ZenithParserNL)
}

func (s *ParenExprContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, i)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitXorExprContext struct {
	ExprContext
	Left  IExprContext
	Right IExprContext
}

func NewBitXorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitXorExprContext {
	var p = new(BitXorExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BitXorExprContext) GetLeft() IExprContext { return s.Left }

func (s *BitXorExprContext) GetRight() IExprContext { return s.Right }

func (s *BitXorExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *BitXorExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *BitXorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitXorExprContext) DOLLAR() antlr.TerminalNode {
	return s.GetToken(ZenithParserDOLLAR, 0)
}

func (s *BitXorExprContext) AllExpr() []IExprContext {
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

func (s *BitXorExprContext) Expr(i int) IExprContext {
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

func (s *BitXorExprContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *BitXorExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterBitXorExpr(s)
	}
}

func (s *BitXorExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitBitXorExpr(s)
	}
}

func (s *BitXorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitBitXorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CoalesceExprContext struct {
	ExprContext
	Left  IExprContext
	Right IExprContext
}

func NewCoalesceExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CoalesceExprContext {
	var p = new(CoalesceExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CoalesceExprContext) GetLeft() IExprContext { return s.Left }

func (s *CoalesceExprContext) GetRight() IExprContext { return s.Right }

func (s *CoalesceExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *CoalesceExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *CoalesceExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoalesceExprContext) COALESCE() antlr.TerminalNode {
	return s.GetToken(ZenithParserCOALESCE, 0)
}

func (s *CoalesceExprContext) AllExpr() []IExprContext {
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

func (s *CoalesceExprContext) Expr(i int) IExprContext {
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

func (s *CoalesceExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterCoalesceExpr(s)
	}
}

func (s *CoalesceExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitCoalesceExpr(s)
	}
}

func (s *CoalesceExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitCoalesceExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ShiftExprContext struct {
	ExprContext
	Left  IExprContext
	Op    antlr.Token
	Right IExprContext
}

func NewShiftExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShiftExprContext {
	var p = new(ShiftExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ShiftExprContext) GetOp() antlr.Token { return s.Op }

func (s *ShiftExprContext) SetOp(v antlr.Token) { s.Op = v }

func (s *ShiftExprContext) GetLeft() IExprContext { return s.Left }

func (s *ShiftExprContext) GetRight() IExprContext { return s.Right }

func (s *ShiftExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *ShiftExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *ShiftExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftExprContext) AllExpr() []IExprContext {
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

func (s *ShiftExprContext) Expr(i int) IExprContext {
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

func (s *ShiftExprContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(ZenithParserLSHIFT, 0)
}

func (s *ShiftExprContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(ZenithParserRSHIFT, 0)
}

func (s *ShiftExprContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *ShiftExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterShiftExpr(s)
	}
}

func (s *ShiftExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitShiftExpr(s)
	}
}

func (s *ShiftExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitShiftExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrefixExprContext struct {
	ExprContext
	Op    antlr.Token
	Right IExprContext
}

func NewPrefixExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrefixExprContext {
	var p = new(PrefixExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *PrefixExprContext) GetOp() antlr.Token { return s.Op }

func (s *PrefixExprContext) SetOp(v antlr.Token) { s.Op = v }

func (s *PrefixExprContext) GetRight() IExprContext { return s.Right }

func (s *PrefixExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *PrefixExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixExprContext) Expr() IExprContext {
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

func (s *PrefixExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ZenithParserPLUS, 0)
}

func (s *PrefixExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ZenithParserMINUS, 0)
}

func (s *PrefixExprContext) EXCLAIM() antlr.TerminalNode {
	return s.GetToken(ZenithParserEXCLAIM, 0)
}

func (s *PrefixExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterPrefixExpr(s)
	}
}

func (s *PrefixExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitPrefixExpr(s)
	}
}

func (s *PrefixExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitPrefixExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeallocExprContext struct {
	ExprContext
}

func NewDeallocExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeallocExprContext {
	var p = new(DeallocExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *DeallocExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeallocExprContext) DEALLOC() antlr.TerminalNode {
	return s.GetToken(ZenithParserDEALLOC, 0)
}

func (s *DeallocExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserLPAREN, 0)
}

func (s *DeallocExprContext) Expr() IExprContext {
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

func (s *DeallocExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserRPAREN, 0)
}

func (s *DeallocExprContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ZenithParserNL)
}

func (s *DeallocExprContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, i)
}

func (s *DeallocExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterDeallocExpr(s)
	}
}

func (s *DeallocExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitDeallocExpr(s)
	}
}

func (s *DeallocExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitDeallocExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitOrExprContext struct {
	ExprContext
	Left  IExprContext
	Right IExprContext
}

func NewBitOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitOrExprContext {
	var p = new(BitOrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BitOrExprContext) GetLeft() IExprContext { return s.Left }

func (s *BitOrExprContext) GetRight() IExprContext { return s.Right }

func (s *BitOrExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *BitOrExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *BitOrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitOrExprContext) PIPE() antlr.TerminalNode {
	return s.GetToken(ZenithParserPIPE, 0)
}

func (s *BitOrExprContext) AllExpr() []IExprContext {
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

func (s *BitOrExprContext) Expr(i int) IExprContext {
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

func (s *BitOrExprContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *BitOrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterBitOrExpr(s)
	}
}

func (s *BitOrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitBitOrExpr(s)
	}
}

func (s *BitOrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitBitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExprContext struct {
	ExprContext
	Left  IExprContext
	Op    antlr.Token
	Right IExprContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddExprContext) GetOp() antlr.Token { return s.Op }

func (s *AddExprContext) SetOp(v antlr.Token) { s.Op = v }

func (s *AddExprContext) GetLeft() IExprContext { return s.Left }

func (s *AddExprContext) GetRight() IExprContext { return s.Right }

func (s *AddExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *AddExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpr() []IExprContext {
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

func (s *AddExprContext) Expr(i int) IExprContext {
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

func (s *AddExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ZenithParserPLUS, 0)
}

func (s *AddExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ZenithParserMINUS, 0)
}

func (s *AddExprContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RefExprContext struct {
	ExprContext
	Left IExprContext
	Op   antlr.Token
}

func NewRefExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RefExprContext {
	var p = new(RefExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RefExprContext) GetOp() antlr.Token { return s.Op }

func (s *RefExprContext) SetOp(v antlr.Token) { s.Op = v }

func (s *RefExprContext) GetLeft() IExprContext { return s.Left }

func (s *RefExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *RefExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RefExprContext) Expr() IExprContext {
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

func (s *RefExprContext) EXCLAIM() antlr.TerminalNode {
	return s.GetToken(ZenithParserEXCLAIM, 0)
}

func (s *RefExprContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(ZenithParserQUESTION, 0)
}

func (s *RefExprContext) HASH() antlr.TerminalNode {
	return s.GetToken(ZenithParserHASH, 0)
}

func (s *RefExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterRefExpr(s)
	}
}

func (s *RefExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitRefExpr(s)
	}
}

func (s *RefExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitRefExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompExprContext struct {
	ExprContext
	Left  IExprContext
	Op    antlr.Token
	Right IExprContext
}

func NewCompExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompExprContext {
	var p = new(CompExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CompExprContext) GetOp() antlr.Token { return s.Op }

func (s *CompExprContext) SetOp(v antlr.Token) { s.Op = v }

func (s *CompExprContext) GetLeft() IExprContext { return s.Left }

func (s *CompExprContext) GetRight() IExprContext { return s.Right }

func (s *CompExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *CompExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *CompExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompExprContext) AllExpr() []IExprContext {
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

func (s *CompExprContext) Expr(i int) IExprContext {
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

func (s *CompExprContext) LT() antlr.TerminalNode {
	return s.GetToken(ZenithParserLT, 0)
}

func (s *CompExprContext) GT() antlr.TerminalNode {
	return s.GetToken(ZenithParserGT, 0)
}

func (s *CompExprContext) LTE() antlr.TerminalNode {
	return s.GetToken(ZenithParserLTE, 0)
}

func (s *CompExprContext) GTE() antlr.TerminalNode {
	return s.GetToken(ZenithParserGTE, 0)
}

func (s *CompExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(ZenithParserEQ, 0)
}

func (s *CompExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(ZenithParserNEQ, 0)
}

func (s *CompExprContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *CompExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterCompExpr(s)
	}
}

func (s *CompExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitCompExpr(s)
	}
}

func (s *CompExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitCompExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulExprContext struct {
	ExprContext
	Left  IExprContext
	Op    antlr.Token
	Right IExprContext
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulExprContext) GetOp() antlr.Token { return s.Op }

func (s *MulExprContext) SetOp(v antlr.Token) { s.Op = v }

func (s *MulExprContext) GetLeft() IExprContext { return s.Left }

func (s *MulExprContext) GetRight() IExprContext { return s.Right }

func (s *MulExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *MulExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllExpr() []IExprContext {
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

func (s *MulExprContext) Expr(i int) IExprContext {
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

func (s *MulExprContext) TIMES() antlr.TerminalNode {
	return s.GetToken(ZenithParserTIMES, 0)
}

func (s *MulExprContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(ZenithParserDIVIDE, 0)
}

func (s *MulExprContext) REM() antlr.TerminalNode {
	return s.GetToken(ZenithParserREM, 0)
}

func (s *MulExprContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *MulExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterMulExpr(s)
	}
}

func (s *MulExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitMulExpr(s)
	}
}

func (s *MulExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitMulExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfExprContext struct {
	ExprContext
	Left      IExprContext
	Condition IExprContext
	Right     IExprContext
}

func NewIfExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfExprContext {
	var p = new(IfExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IfExprContext) GetLeft() IExprContext { return s.Left }

func (s *IfExprContext) GetCondition() IExprContext { return s.Condition }

func (s *IfExprContext) GetRight() IExprContext { return s.Right }

func (s *IfExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *IfExprContext) SetCondition(v IExprContext) { s.Condition = v }

func (s *IfExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *IfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExprContext) IF() antlr.TerminalNode {
	return s.GetToken(ZenithParserIF, 0)
}

func (s *IfExprContext) ELSE() antlr.TerminalNode {
	return s.GetToken(ZenithParserELSE, 0)
}

func (s *IfExprContext) AllExpr() []IExprContext {
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

func (s *IfExprContext) Expr(i int) IExprContext {
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

func (s *IfExprContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ZenithParserNL)
}

func (s *IfExprContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, i)
}

func (s *IfExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterIfExpr(s)
	}
}

func (s *IfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitIfExpr(s)
	}
}

func (s *IfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitIfExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BitAndExprContext struct {
	ExprContext
	Left  IExprContext
	Right IExprContext
}

func NewBitAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitAndExprContext {
	var p = new(BitAndExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BitAndExprContext) GetLeft() IExprContext { return s.Left }

func (s *BitAndExprContext) GetRight() IExprContext { return s.Right }

func (s *BitAndExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *BitAndExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *BitAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitAndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(ZenithParserAND, 0)
}

func (s *BitAndExprContext) AllExpr() []IExprContext {
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

func (s *BitAndExprContext) Expr(i int) IExprContext {
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

func (s *BitAndExprContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *BitAndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterBitAndExpr(s)
	}
}

func (s *BitAndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitBitAndExpr(s)
	}
}

func (s *BitAndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitBitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowExprContext struct {
	ExprContext
	Left  IExprContext
	Right IExprContext
}

func NewPowExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowExprContext {
	var p = new(PowExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *PowExprContext) GetLeft() IExprContext { return s.Left }

func (s *PowExprContext) GetRight() IExprContext { return s.Right }

func (s *PowExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *PowExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *PowExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowExprContext) AllExpr() []IExprContext {
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

func (s *PowExprContext) Expr(i int) IExprContext {
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

func (s *PowExprContext) AllPOW() []antlr.TerminalNode {
	return s.GetTokens(ZenithParserPOW)
}

func (s *PowExprContext) POW(i int) antlr.TerminalNode {
	return s.GetToken(ZenithParserPOW, i)
}

func (s *PowExprContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ZenithParserNL)
}

func (s *PowExprContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, i)
}

func (s *PowExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterPowExpr(s)
	}
}

func (s *PowExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitPowExpr(s)
	}
}

func (s *PowExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitPowExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotNullExprContext struct {
	ExprContext
	Left IExprContext
}

func NewNotNullExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotNullExprContext {
	var p = new(NotNullExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotNullExprContext) GetLeft() IExprContext { return s.Left }

func (s *NotNullExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *NotNullExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotNullExprContext) NOT_NULL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNOT_NULL, 0)
}

func (s *NotNullExprContext) Expr() IExprContext {
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

func (s *NotNullExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterNotNullExpr(s)
	}
}

func (s *NotNullExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitNotNullExpr(s)
	}
}

func (s *NotNullExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitNotNullExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	ExprContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(ZenithParserID, 0)
}

func (s *IdExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterIdExpr(s)
	}
}

func (s *IdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitIdExpr(s)
	}
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZenithParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ZenithParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, ZenithParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ZenithParserNUM:
		localctx = NewNumExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(77)
			p.Match(ZenithParserNUM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZenithParserID:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(78)
			p.Match(ZenithParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZenithParserNULL, ZenithParserTRUE, ZenithParserFALSE:
		localctx = NewKeyExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(79)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*KeyExprContext).Key = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4980736) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*KeyExprContext).Key = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case ZenithParserLPAREN:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(80)
			p.Match(ZenithParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(81)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(84)
			p.expr(0)
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(85)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(88)
			p.Match(ZenithParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZenithParserALLOC:
		localctx = NewAllocExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(90)
			p.Match(ZenithParserALLOC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AllocExprContext).Ref = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ZenithParserEXCLAIM || _la == ZenithParserQUESTION) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AllocExprContext).Ref = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(92)
			p.Match(ZenithParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(93)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(96)
			p.expr(0)
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(97)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(100)
			p.Match(ZenithParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZenithParserDEALLOC:
		localctx = NewDeallocExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(102)
			p.Match(ZenithParserDEALLOC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(ZenithParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(104)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(107)
			p.expr(0)
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(108)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(111)
			p.Match(ZenithParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZenithParserTYPE:
		localctx = NewCastExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(113)
			p.Match(ZenithParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Match(ZenithParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(115)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(118)
			p.expr(0)
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(119)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(122)
			p.Match(ZenithParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZenithParserPLUS, ZenithParserMINUS, ZenithParserEXCLAIM:
		localctx = NewPrefixExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(124)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*PrefixExprContext).Op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8589934976) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*PrefixExprContext).Op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(125)

			var _x = p.expr(11)

			localctx.(*PrefixExprContext).Right = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(207)
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
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MulExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(128)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(129)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulExprContext).Op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&112) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulExprContext).Op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(131)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(130)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(133)

					var _x = p.expr(10)

					localctx.(*MulExprContext).Right = _x
				}

			case 2:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(134)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(135)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddExprContext).Op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ZenithParserPLUS || _la == ZenithParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddExprContext).Op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(137)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(136)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(139)

					var _x = p.expr(9)

					localctx.(*AddExprContext).Right = _x
				}

			case 3:
				localctx = NewShiftExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*ShiftExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(140)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(141)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ShiftExprContext).Op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ZenithParserLSHIFT || _la == ZenithParserRSHIFT) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ShiftExprContext).Op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(143)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(142)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(145)

					var _x = p.expr(8)

					localctx.(*ShiftExprContext).Right = _x
				}

			case 4:
				localctx = NewBitAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BitAndExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(147)
					p.Match(ZenithParserAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(149)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(148)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(151)

					var _x = p.expr(7)

					localctx.(*BitAndExprContext).Right = _x
				}

			case 5:
				localctx = NewBitXorExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BitXorExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(153)
					p.Match(ZenithParserDOLLAR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(155)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(154)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(157)

					var _x = p.expr(6)

					localctx.(*BitXorExprContext).Right = _x
				}

			case 6:
				localctx = NewBitOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*BitOrExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(158)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(159)
					p.Match(ZenithParserPIPE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(161)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(160)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(163)

					var _x = p.expr(5)

					localctx.(*BitOrExprContext).Right = _x
				}

			case 7:
				localctx = NewCompExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*CompExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(164)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(165)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompExprContext).Op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&129024) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompExprContext).Op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(167)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(166)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(169)

					var _x = p.expr(4)

					localctx.(*CompExprContext).Right = _x
				}

			case 8:
				localctx = NewCoalesceExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*CoalesceExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(170)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(171)
					p.Match(ZenithParserCOALESCE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(172)

					var _x = p.expr(3)

					localctx.(*CoalesceExprContext).Right = _x
				}

			case 9:
				localctx = NewIfExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*IfExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(173)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(174)
					p.Match(ZenithParserIF)
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

				if _la == ZenithParserNL {
					{
						p.SetState(175)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(178)

					var _x = p.expr(0)

					localctx.(*IfExprContext).Condition = _x
				}
				p.SetState(180)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(179)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(182)
					p.Match(ZenithParserELSE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(184)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(183)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(186)

					var _x = p.expr(2)

					localctx.(*IfExprContext).Right = _x
				}

			case 10:
				localctx = NewRefExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*RefExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(188)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				p.SetState(190)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserHASH {
					{
						p.SetState(189)
						p.Match(ZenithParserHASH)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(192)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RefExprContext).Op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ZenithParserEXCLAIM || _la == ZenithParserQUESTION) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RefExprContext).Op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			case 11:
				localctx = NewNotNullExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*NotNullExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(193)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(194)
					p.Match(ZenithParserNOT_NULL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 12:
				localctx = NewPowExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*PowExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				p.SetState(201)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(196)
							p.Match(ZenithParserPOW)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}
						p.SetState(198)
						p.GetErrorHandler().Sync(p)
						if p.HasError() {
							goto errorExit
						}
						_la = p.GetTokenStream().LA(1)

						if _la == ZenithParserNL {
							{
								p.SetState(197)
								p.Match(ZenithParserNL)
								if p.HasError() {
									// Recognition error - abort rule
									goto errorExit
								}
							}

						}
						{
							p.SetState(200)

							var _x = p.expr(0)

							localctx.(*PowExprContext).Right = _x
						}

					default:
						p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
						goto errorExit
					}

					p.SetState(203)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
					if p.HasError() {
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
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

func (p *ZenithParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *TypeContext = nil
		if localctx != nil {
			t = localctx.(*TypeContext)
		}
		return p.Type__Sempred(t, predIndex)

	case 5:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ZenithParser) Type__Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ZenithParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
