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

	// EnterRefType is called when entering the refType production.
	EnterRefType(c *RefTypeContext)

	// EnterBaseType is called when entering the baseType production.
	EnterBaseType(c *BaseTypeContext)

	// EnterCastExpr is called when entering the castExpr production.
	EnterCastExpr(c *CastExprContext)

	// EnterKeyExpr is called when entering the keyExpr production.
	EnterKeyExpr(c *KeyExprContext)

	// EnterNumExpr is called when entering the numExpr production.
	EnterNumExpr(c *NumExprContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterBitXorExpr is called when entering the bitXorExpr production.
	EnterBitXorExpr(c *BitXorExprContext)

	// EnterShiftExpr is called when entering the shiftExpr production.
	EnterShiftExpr(c *ShiftExprContext)

	// EnterPrefixExpr is called when entering the prefixExpr production.
	EnterPrefixExpr(c *PrefixExprContext)

	// EnterBitOrExpr is called when entering the bitOrExpr production.
	EnterBitOrExpr(c *BitOrExprContext)

	// EnterAddExpr is called when entering the addExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterRefExpr is called when entering the refExpr production.
	EnterRefExpr(c *RefExprContext)

	// EnterCompExpr is called when entering the compExpr production.
	EnterCompExpr(c *CompExprContext)

	// EnterMulExpr is called when entering the mulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterIfExpr is called when entering the ifExpr production.
	EnterIfExpr(c *IfExprContext)

	// EnterBitAndExpr is called when entering the bitAndExpr production.
	EnterBitAndExpr(c *BitAndExprContext)

	// EnterPowExpr is called when entering the powExpr production.
	EnterPowExpr(c *PowExprContext)

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

	// ExitRefType is called when exiting the refType production.
	ExitRefType(c *RefTypeContext)

	// ExitBaseType is called when exiting the baseType production.
	ExitBaseType(c *BaseTypeContext)

	// ExitCastExpr is called when exiting the castExpr production.
	ExitCastExpr(c *CastExprContext)

	// ExitKeyExpr is called when exiting the keyExpr production.
	ExitKeyExpr(c *KeyExprContext)

	// ExitNumExpr is called when exiting the numExpr production.
	ExitNumExpr(c *NumExprContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitBitXorExpr is called when exiting the bitXorExpr production.
	ExitBitXorExpr(c *BitXorExprContext)

	// ExitShiftExpr is called when exiting the shiftExpr production.
	ExitShiftExpr(c *ShiftExprContext)

	// ExitPrefixExpr is called when exiting the prefixExpr production.
	ExitPrefixExpr(c *PrefixExprContext)

	// ExitBitOrExpr is called when exiting the bitOrExpr production.
	ExitBitOrExpr(c *BitOrExprContext)

	// ExitAddExpr is called when exiting the addExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitRefExpr is called when exiting the refExpr production.
	ExitRefExpr(c *RefExprContext)

	// ExitCompExpr is called when exiting the compExpr production.
	ExitCompExpr(c *CompExprContext)

	// ExitMulExpr is called when exiting the mulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitIfExpr is called when exiting the ifExpr production.
	ExitIfExpr(c *IfExprContext)

	// ExitBitAndExpr is called when exiting the bitAndExpr production.
	ExitBitAndExpr(c *BitAndExprContext)

	// ExitPowExpr is called when exiting the powExpr production.
	ExitPowExpr(c *PowExprContext)

	// ExitIdExpr is called when exiting the idExpr production.
	ExitIdExpr(c *IdExprContext)
}
