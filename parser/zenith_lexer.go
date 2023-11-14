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
		"PLUS", "MINUS", "TIMES", "DIVIDE", "NUM", "BIN_DIGIT", "OCT_DIGIT",
		"DEC_DIGIT", "HEX_DIGIT", "NL", "SPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 7, 184, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 4, 4,
		33, 8, 4, 11, 4, 12, 4, 34, 1, 4, 5, 4, 38, 8, 4, 10, 4, 12, 4, 41, 9,
		4, 1, 4, 1, 4, 4, 4, 45, 8, 4, 11, 4, 12, 4, 46, 1, 4, 4, 4, 50, 8, 4,
		11, 4, 12, 4, 51, 1, 4, 1, 4, 5, 4, 56, 8, 4, 10, 4, 12, 4, 59, 9, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 4, 4, 65, 8, 4, 11, 4, 12, 4, 66, 1, 4, 5, 4, 70,
		8, 4, 10, 4, 12, 4, 73, 9, 4, 1, 4, 1, 4, 4, 4, 77, 8, 4, 11, 4, 12, 4,
		78, 1, 4, 4, 4, 82, 8, 4, 11, 4, 12, 4, 83, 1, 4, 1, 4, 5, 4, 88, 8, 4,
		10, 4, 12, 4, 91, 9, 4, 3, 4, 93, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 4, 4, 99,
		8, 4, 11, 4, 12, 4, 100, 1, 4, 5, 4, 104, 8, 4, 10, 4, 12, 4, 107, 9, 4,
		1, 4, 1, 4, 4, 4, 111, 8, 4, 11, 4, 12, 4, 112, 1, 4, 4, 4, 116, 8, 4,
		11, 4, 12, 4, 117, 1, 4, 1, 4, 5, 4, 122, 8, 4, 10, 4, 12, 4, 125, 9, 4,
		3, 4, 127, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 4, 4, 133, 8, 4, 11, 4, 12, 4,
		134, 1, 4, 5, 4, 138, 8, 4, 10, 4, 12, 4, 141, 9, 4, 1, 4, 1, 4, 4, 4,
		145, 8, 4, 11, 4, 12, 4, 146, 1, 4, 4, 4, 150, 8, 4, 11, 4, 12, 4, 151,
		1, 4, 1, 4, 5, 4, 156, 8, 4, 10, 4, 12, 4, 159, 9, 4, 3, 4, 161, 8, 4,
		3, 4, 163, 8, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9,
		3, 9, 174, 8, 9, 1, 9, 4, 9, 177, 8, 9, 11, 9, 12, 9, 178, 1, 10, 1, 10,
		1, 10, 1, 10, 0, 0, 11, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 0, 13, 0, 15,
		0, 17, 0, 19, 6, 21, 7, 1, 0, 7, 1, 0, 48, 49, 1, 0, 48, 55, 1, 0, 48,
		57, 3, 0, 48, 57, 65, 70, 97, 102, 1, 0, 13, 13, 1, 0, 10, 10, 2, 0, 9,
		9, 32, 32, 212, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0,
		7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		1, 23, 1, 0, 0, 0, 3, 25, 1, 0, 0, 0, 5, 27, 1, 0, 0, 0, 7, 29, 1, 0, 0,
		0, 9, 162, 1, 0, 0, 0, 11, 164, 1, 0, 0, 0, 13, 166, 1, 0, 0, 0, 15, 168,
		1, 0, 0, 0, 17, 170, 1, 0, 0, 0, 19, 176, 1, 0, 0, 0, 21, 180, 1, 0, 0,
		0, 23, 24, 5, 43, 0, 0, 24, 2, 1, 0, 0, 0, 25, 26, 5, 45, 0, 0, 26, 4,
		1, 0, 0, 0, 27, 28, 5, 42, 0, 0, 28, 6, 1, 0, 0, 0, 29, 30, 5, 47, 0, 0,
		30, 8, 1, 0, 0, 0, 31, 33, 3, 15, 7, 0, 32, 31, 1, 0, 0, 0, 33, 34, 1,
		0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 163, 1, 0, 0, 0, 36,
		38, 3, 15, 7, 0, 37, 36, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0,
		0, 0, 39, 40, 1, 0, 0, 0, 40, 42, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 44,
		5, 46, 0, 0, 43, 45, 3, 15, 7, 0, 44, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0,
		0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 163, 1, 0, 0, 0, 48, 50,
		3, 15, 7, 0, 49, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0,
		51, 52, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 57, 5, 46, 0, 0, 54, 56, 3,
		15, 7, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57,
		58, 1, 0, 0, 0, 58, 163, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 61, 5, 48,
		0, 0, 61, 62, 5, 98, 0, 0, 62, 92, 1, 0, 0, 0, 63, 65, 3, 11, 5, 0, 64,
		63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0,
		0, 67, 93, 1, 0, 0, 0, 68, 70, 3, 11, 5, 0, 69, 68, 1, 0, 0, 0, 70, 73,
		1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 74, 1, 0, 0, 0,
		73, 71, 1, 0, 0, 0, 74, 76, 5, 46, 0, 0, 75, 77, 3, 11, 5, 0, 76, 75, 1,
		0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79,
		93, 1, 0, 0, 0, 80, 82, 3, 11, 5, 0, 81, 80, 1, 0, 0, 0, 82, 83, 1, 0,
		0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 89,
		5, 46, 0, 0, 86, 88, 3, 11, 5, 0, 87, 86, 1, 0, 0, 0, 88, 91, 1, 0, 0,
		0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89,
		1, 0, 0, 0, 92, 64, 1, 0, 0, 0, 92, 71, 1, 0, 0, 0, 92, 81, 1, 0, 0, 0,
		93, 163, 1, 0, 0, 0, 94, 95, 5, 48, 0, 0, 95, 96, 5, 111, 0, 0, 96, 126,
		1, 0, 0, 0, 97, 99, 3, 13, 6, 0, 98, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0,
		0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 127, 1, 0, 0, 0, 102,
		104, 3, 13, 6, 0, 103, 102, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103,
		1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 108, 1, 0, 0, 0, 107, 105, 1, 0,
		0, 0, 108, 110, 5, 46, 0, 0, 109, 111, 3, 13, 6, 0, 110, 109, 1, 0, 0,
		0, 111, 112, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113,
		127, 1, 0, 0, 0, 114, 116, 3, 13, 6, 0, 115, 114, 1, 0, 0, 0, 116, 117,
		1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 1, 0,
		0, 0, 119, 123, 5, 46, 0, 0, 120, 122, 3, 13, 6, 0, 121, 120, 1, 0, 0,
		0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124,
		127, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 98, 1, 0, 0, 0, 126, 105, 1,
		0, 0, 0, 126, 115, 1, 0, 0, 0, 127, 163, 1, 0, 0, 0, 128, 129, 5, 48, 0,
		0, 129, 130, 5, 120, 0, 0, 130, 160, 1, 0, 0, 0, 131, 133, 3, 17, 8, 0,
		132, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134,
		135, 1, 0, 0, 0, 135, 161, 1, 0, 0, 0, 136, 138, 3, 17, 8, 0, 137, 136,
		1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0,
		0, 0, 140, 142, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 144, 5, 46, 0, 0,
		143, 145, 3, 17, 8, 0, 144, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146,
		144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 161, 1, 0, 0, 0, 148, 150,
		3, 17, 8, 0, 149, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 149, 1, 0,
		0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 157, 5, 46, 0, 0,
		154, 156, 3, 17, 8, 0, 155, 154, 1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157,
		155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 161, 1, 0, 0, 0, 159, 157,
		1, 0, 0, 0, 160, 132, 1, 0, 0, 0, 160, 139, 1, 0, 0, 0, 160, 149, 1, 0,
		0, 0, 161, 163, 1, 0, 0, 0, 162, 32, 1, 0, 0, 0, 162, 39, 1, 0, 0, 0, 162,
		49, 1, 0, 0, 0, 162, 60, 1, 0, 0, 0, 162, 94, 1, 0, 0, 0, 162, 128, 1,
		0, 0, 0, 163, 10, 1, 0, 0, 0, 164, 165, 7, 0, 0, 0, 165, 12, 1, 0, 0, 0,
		166, 167, 7, 1, 0, 0, 167, 14, 1, 0, 0, 0, 168, 169, 7, 2, 0, 0, 169, 16,
		1, 0, 0, 0, 170, 171, 7, 3, 0, 0, 171, 18, 1, 0, 0, 0, 172, 174, 7, 4,
		0, 0, 173, 172, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0,
		175, 177, 7, 5, 0, 0, 176, 173, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178,
		176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 20, 1, 0, 0, 0, 180, 181, 7,
		6, 0, 0, 181, 182, 1, 0, 0, 0, 182, 183, 6, 10, 0, 0, 183, 22, 1, 0, 0,
		0, 27, 0, 34, 39, 46, 51, 57, 66, 71, 78, 83, 89, 92, 100, 105, 112, 117,
		123, 126, 134, 139, 146, 151, 157, 160, 162, 173, 178, 1, 6, 0, 0,
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
