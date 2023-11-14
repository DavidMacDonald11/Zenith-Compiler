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
		"", "'+'", "'-'", "'*'", "'/'",
	}
	staticData.SymbolicNames = []string{
		"", "PLUS", "MINUS", "TIMES", "DIVIDE", "NUM", "NL", "SPACE",
	}
	staticData.RuleNames = []string{
		"PLUS", "MINUS", "TIMES", "DIVIDE", "NUM", "DEC_NUM", "DEC_SEG", "BIN_NUM",
		"BIN_SEG", "OCT_NUM", "OCT_SEG", "HEX_NUM", "HEX_SEG", "HEX_DIGIT",
		"NL", "SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 7, 196, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 3, 4, 46, 8, 4, 1, 5, 1, 5, 5, 5, 50, 8, 5, 10, 5, 12, 5, 53, 9,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 60, 8, 5, 10, 5, 12, 5, 63, 9, 5,
		3, 5, 65, 8, 5, 1, 6, 1, 6, 5, 6, 69, 8, 6, 10, 6, 12, 6, 72, 9, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 79, 8, 7, 10, 7, 12, 7, 82, 9, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 89, 8, 7, 10, 7, 12, 7, 92, 9, 7, 3, 7, 94,
		8, 7, 1, 8, 5, 8, 97, 8, 8, 10, 8, 12, 8, 100, 9, 8, 1, 8, 1, 8, 5, 8,
		104, 8, 8, 10, 8, 12, 8, 107, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9,
		114, 8, 9, 10, 9, 12, 9, 117, 9, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9,
		124, 8, 9, 10, 9, 12, 9, 127, 9, 9, 3, 9, 129, 8, 9, 1, 10, 5, 10, 132,
		8, 10, 10, 10, 12, 10, 135, 9, 10, 1, 10, 1, 10, 5, 10, 139, 8, 10, 10,
		10, 12, 10, 142, 9, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11,
		150, 8, 11, 10, 11, 12, 11, 153, 9, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 5, 11, 161, 8, 11, 10, 11, 12, 11, 164, 9, 11, 3, 11, 166, 8,
		11, 1, 12, 1, 12, 5, 12, 170, 8, 12, 10, 12, 12, 12, 173, 9, 12, 1, 12,
		1, 12, 1, 12, 5, 12, 178, 8, 12, 10, 12, 12, 12, 181, 9, 12, 1, 13, 1,
		13, 1, 14, 3, 14, 186, 8, 14, 1, 14, 4, 14, 189, 8, 14, 11, 14, 12, 14,
		190, 1, 15, 1, 15, 1, 15, 1, 15, 0, 0, 16, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 0, 13, 0, 15, 0, 17, 0, 19, 0, 21, 0, 23, 0, 25, 0, 27, 0, 29, 6, 31,
		7, 1, 0, 10, 2, 0, 48, 57, 95, 95, 1, 0, 48, 57, 2, 0, 48, 49, 95, 95,
		1, 0, 48, 49, 2, 0, 48, 55, 95, 95, 1, 0, 48, 55, 3, 0, 48, 57, 65, 70,
		97, 102, 1, 0, 13, 13, 1, 0, 10, 10, 2, 0, 9, 9, 32, 32, 218, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 1, 33, 1, 0, 0, 0, 3, 35,
		1, 0, 0, 0, 5, 37, 1, 0, 0, 0, 7, 39, 1, 0, 0, 0, 9, 45, 1, 0, 0, 0, 11,
		64, 1, 0, 0, 0, 13, 66, 1, 0, 0, 0, 15, 73, 1, 0, 0, 0, 17, 98, 1, 0, 0,
		0, 19, 108, 1, 0, 0, 0, 21, 133, 1, 0, 0, 0, 23, 143, 1, 0, 0, 0, 25, 171,
		1, 0, 0, 0, 27, 182, 1, 0, 0, 0, 29, 188, 1, 0, 0, 0, 31, 192, 1, 0, 0,
		0, 33, 34, 5, 43, 0, 0, 34, 2, 1, 0, 0, 0, 35, 36, 5, 45, 0, 0, 36, 4,
		1, 0, 0, 0, 37, 38, 5, 42, 0, 0, 38, 6, 1, 0, 0, 0, 39, 40, 5, 47, 0, 0,
		40, 8, 1, 0, 0, 0, 41, 46, 3, 11, 5, 0, 42, 46, 3, 15, 7, 0, 43, 46, 3,
		19, 9, 0, 44, 46, 3, 23, 11, 0, 45, 41, 1, 0, 0, 0, 45, 42, 1, 0, 0, 0,
		45, 43, 1, 0, 0, 0, 45, 44, 1, 0, 0, 0, 46, 10, 1, 0, 0, 0, 47, 65, 3,
		13, 6, 0, 48, 50, 7, 0, 0, 0, 49, 48, 1, 0, 0, 0, 50, 53, 1, 0, 0, 0, 51,
		49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 54, 1, 0, 0, 0, 53, 51, 1, 0, 0,
		0, 54, 55, 5, 46, 0, 0, 55, 65, 3, 13, 6, 0, 56, 57, 3, 13, 6, 0, 57, 61,
		5, 46, 0, 0, 58, 60, 7, 0, 0, 0, 59, 58, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0,
		61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 65, 1, 0, 0, 0, 63, 61, 1,
		0, 0, 0, 64, 47, 1, 0, 0, 0, 64, 51, 1, 0, 0, 0, 64, 56, 1, 0, 0, 0, 65,
		12, 1, 0, 0, 0, 66, 70, 7, 1, 0, 0, 67, 69, 7, 0, 0, 0, 68, 67, 1, 0, 0,
		0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 14,
		1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 74, 5, 48, 0, 0, 74, 75, 5, 98, 0,
		0, 75, 93, 1, 0, 0, 0, 76, 94, 3, 17, 8, 0, 77, 79, 7, 2, 0, 0, 78, 77,
		1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0,
		81, 83, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 83, 84, 5, 46, 0, 0, 84, 94, 3,
		17, 8, 0, 85, 86, 3, 17, 8, 0, 86, 90, 5, 46, 0, 0, 87, 89, 7, 2, 0, 0,
		88, 87, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1,
		0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 76, 1, 0, 0, 0, 93,
		80, 1, 0, 0, 0, 93, 85, 1, 0, 0, 0, 94, 16, 1, 0, 0, 0, 95, 97, 7, 2, 0,
		0, 96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99,
		1, 0, 0, 0, 99, 101, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 105, 7, 3, 0,
		0, 102, 104, 7, 2, 0, 0, 103, 102, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105,
		103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 18, 1, 0, 0, 0, 107, 105, 1,
		0, 0, 0, 108, 109, 5, 48, 0, 0, 109, 110, 5, 111, 0, 0, 110, 128, 1, 0,
		0, 0, 111, 129, 3, 21, 10, 0, 112, 114, 7, 4, 0, 0, 113, 112, 1, 0, 0,
		0, 114, 117, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116,
		118, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 118, 119, 5, 46, 0, 0, 119, 129,
		3, 21, 10, 0, 120, 121, 3, 21, 10, 0, 121, 125, 5, 46, 0, 0, 122, 124,
		7, 4, 0, 0, 123, 122, 1, 0, 0, 0, 124, 127, 1, 0, 0, 0, 125, 123, 1, 0,
		0, 0, 125, 126, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0,
		128, 111, 1, 0, 0, 0, 128, 115, 1, 0, 0, 0, 128, 120, 1, 0, 0, 0, 129,
		20, 1, 0, 0, 0, 130, 132, 7, 4, 0, 0, 131, 130, 1, 0, 0, 0, 132, 135, 1,
		0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 136, 1, 0, 0,
		0, 135, 133, 1, 0, 0, 0, 136, 140, 7, 5, 0, 0, 137, 139, 7, 4, 0, 0, 138,
		137, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141,
		1, 0, 0, 0, 141, 22, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 144, 5, 48,
		0, 0, 144, 145, 5, 120, 0, 0, 145, 165, 1, 0, 0, 0, 146, 166, 3, 25, 12,
		0, 147, 150, 3, 27, 13, 0, 148, 150, 5, 95, 0, 0, 149, 147, 1, 0, 0, 0,
		149, 148, 1, 0, 0, 0, 150, 153, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151,
		152, 1, 0, 0, 0, 152, 154, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 154, 155,
		5, 46, 0, 0, 155, 166, 3, 25, 12, 0, 156, 157, 3, 25, 12, 0, 157, 162,
		5, 46, 0, 0, 158, 161, 3, 27, 13, 0, 159, 161, 5, 95, 0, 0, 160, 158, 1,
		0, 0, 0, 160, 159, 1, 0, 0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0, 0,
		0, 162, 163, 1, 0, 0, 0, 163, 166, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165,
		146, 1, 0, 0, 0, 165, 151, 1, 0, 0, 0, 165, 156, 1, 0, 0, 0, 166, 24, 1,
		0, 0, 0, 167, 170, 3, 27, 13, 0, 168, 170, 5, 95, 0, 0, 169, 167, 1, 0,
		0, 0, 169, 168, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0,
		171, 172, 1, 0, 0, 0, 172, 174, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 174,
		179, 3, 27, 13, 0, 175, 178, 3, 27, 13, 0, 176, 178, 5, 95, 0, 0, 177,
		175, 1, 0, 0, 0, 177, 176, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177,
		1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 26, 1, 0, 0, 0, 181, 179, 1, 0,
		0, 0, 182, 183, 7, 6, 0, 0, 183, 28, 1, 0, 0, 0, 184, 186, 7, 7, 0, 0,
		185, 184, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187,
		189, 7, 8, 0, 0, 188, 185, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 188,
		1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 30, 1, 0, 0, 0, 192, 193, 7, 9,
		0, 0, 193, 194, 1, 0, 0, 0, 194, 195, 6, 15, 0, 0, 195, 32, 1, 0, 0, 0,
		27, 0, 45, 51, 61, 64, 70, 80, 90, 93, 98, 105, 115, 125, 128, 133, 140,
		149, 151, 160, 162, 165, 169, 171, 177, 179, 185, 190, 1, 6, 0, 0,
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
	ZenithLexerPLUS   = 1
	ZenithLexerMINUS  = 2
	ZenithLexerTIMES  = 3
	ZenithLexerDIVIDE = 4
	ZenithLexerNUM    = 5
	ZenithLexerNL     = 6
	ZenithLexerSPACE  = 7
)
