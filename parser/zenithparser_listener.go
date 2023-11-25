// Code generated from ZenithParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ZenithParser

import "github.com/antlr4-go/antlr/v4"

// ZenithParserListener is a complete listener for a parse tree produced by ZenithParser.
type ZenithParserListener interface {
	antlr.ParseTreeListener

	// EnterFileStat is called when entering the fileStat production.
	EnterFileStat(c *FileStatContext)

	// EnterEndedStat is called when entering the endedStat production.
	EnterEndedStat(c *EndedStatContext)

	// EnterLineEnd is called when entering the lineEnd production.
	EnterLineEnd(c *LineEndContext)

	// EnterDefineStat is called when entering the defineStat production.
	EnterDefineStat(c *DefineStatContext)

	// EnterMultiStat is called when entering the multiStat production.
	EnterMultiStat(c *MultiStatContext)

	// EnterExprStat is called when entering the exprStat production.
	EnterExprStat(c *ExprStatContext)

	// EnterBlankStat is called when entering the blankStat production.
	EnterBlankStat(c *BlankStatContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterPrefixExpr is called when entering the prefixExpr production.
	EnterPrefixExpr(c *PrefixExprContext)

	// EnterCastExpr is called when entering the castExpr production.
	EnterCastExpr(c *CastExprContext)

	// EnterAddExpr is called when entering the addExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterMulExpr is called when entering the mulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterIfExpr is called when entering the ifExpr production.
	EnterIfExpr(c *IfExprContext)

	// EnterKeyExpr is called when entering the keyExpr production.
	EnterKeyExpr(c *KeyExprContext)

	// EnterNumExpr is called when entering the numExpr production.
	EnterNumExpr(c *NumExprContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterIdExpr is called when entering the idExpr production.
	EnterIdExpr(c *IdExprContext)

	// ExitFileStat is called when exiting the fileStat production.
	ExitFileStat(c *FileStatContext)

	// ExitEndedStat is called when exiting the endedStat production.
	ExitEndedStat(c *EndedStatContext)

	// ExitLineEnd is called when exiting the lineEnd production.
	ExitLineEnd(c *LineEndContext)

	// ExitDefineStat is called when exiting the defineStat production.
	ExitDefineStat(c *DefineStatContext)

	// ExitMultiStat is called when exiting the multiStat production.
	ExitMultiStat(c *MultiStatContext)

	// ExitExprStat is called when exiting the exprStat production.
	ExitExprStat(c *ExprStatContext)

	// ExitBlankStat is called when exiting the blankStat production.
	ExitBlankStat(c *BlankStatContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitPrefixExpr is called when exiting the prefixExpr production.
	ExitPrefixExpr(c *PrefixExprContext)

	// ExitCastExpr is called when exiting the castExpr production.
	ExitCastExpr(c *CastExprContext)

	// ExitAddExpr is called when exiting the addExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitMulExpr is called when exiting the mulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitIfExpr is called when exiting the ifExpr production.
	ExitIfExpr(c *IfExprContext)

	// ExitKeyExpr is called when exiting the keyExpr production.
	ExitKeyExpr(c *KeyExprContext)

	// ExitNumExpr is called when exiting the numExpr production.
	ExitNumExpr(c *NumExprContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitIdExpr is called when exiting the idExpr production.
	ExitIdExpr(c *IdExprContext)
}
