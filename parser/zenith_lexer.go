// Code generated from Zenith.g4 by ANTLR 4.13.1. DO NOT EDIT.

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
		"", "'*'", "'/'", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "MUL", "DIV", "ADD", "SUB", "NUM", "NL", "WS",
	}
	staticData.RuleNames = []string{
		"MUL", "DIV", "ADD", "SUB", "NUM", "NL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 7, 40, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 4, 4, 25, 8, 4, 11, 4, 12, 4, 26, 1, 5, 4, 5, 30, 8, 5,
		11, 5, 12, 5, 31, 1, 6, 4, 6, 35, 8, 6, 11, 6, 12, 6, 36, 1, 6, 1, 6, 0,
		0, 7, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 1, 0, 3, 1, 0, 48, 57,
		1, 0, 10, 10, 3, 0, 9, 9, 13, 13, 32, 32, 42, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 1, 15, 1, 0, 0, 0, 3, 17, 1, 0, 0, 0, 5, 19,
		1, 0, 0, 0, 7, 21, 1, 0, 0, 0, 9, 24, 1, 0, 0, 0, 11, 29, 1, 0, 0, 0, 13,
		34, 1, 0, 0, 0, 15, 16, 5, 42, 0, 0, 16, 2, 1, 0, 0, 0, 17, 18, 5, 47,
		0, 0, 18, 4, 1, 0, 0, 0, 19, 20, 5, 43, 0, 0, 20, 6, 1, 0, 0, 0, 21, 22,
		5, 45, 0, 0, 22, 8, 1, 0, 0, 0, 23, 25, 7, 0, 0, 0, 24, 23, 1, 0, 0, 0,
		25, 26, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 10, 1,
		0, 0, 0, 28, 30, 7, 1, 0, 0, 29, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31,
		29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 12, 1, 0, 0, 0, 33, 35, 7, 2, 0,
		0, 34, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37,
		1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 39, 6, 6, 0, 0, 39, 14, 1, 0, 0, 0,
		4, 0, 26, 31, 36, 1, 6, 0, 0,
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
	l.GrammarFileName = "Zenith.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ZenithLexer tokens.
const (
	ZenithLexerMUL = 1
	ZenithLexerDIV = 2
	ZenithLexerADD = 3
	ZenithLexerSUB = 4
	ZenithLexerNUM = 5
	ZenithLexerNL  = 6
	ZenithLexerWS  = 7
)
