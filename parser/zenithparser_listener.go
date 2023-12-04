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

	// EnterFullType is called when entering the fullType production.
	EnterFullType(c *FullTypeContext)

	// EnterPtrType is called when entering the ptrType production.
	EnterPtrType(c *PtrTypeContext)

	// EnterBaseType is called when entering the baseType production.
	EnterBaseType(c *BaseTypeContext)

	// EnterSliceType is called when entering the sliceType production.
	EnterSliceType(c *SliceTypeContext)

	// EnterCastExpr is called when entering the castExpr production.
	EnterCastExpr(c *CastExprContext)

	// EnterAllocExpr is called when entering the allocExpr production.
	EnterAllocExpr(c *AllocExprContext)

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

	// EnterPtrExpr is called when entering the ptrExpr production.
	EnterPtrExpr(c *PtrExprContext)

	// EnterPrefixExpr is called when entering the prefixExpr production.
	EnterPrefixExpr(c *PrefixExprContext)

	// EnterDeallocExpr is called when entering the deallocExpr production.
	EnterDeallocExpr(c *DeallocExprContext)

	// EnterBitOrExpr is called when entering the bitOrExpr production.
	EnterBitOrExpr(c *BitOrExprContext)

	// EnterAddExpr is called when entering the addExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterCompExpr is called when entering the compExpr production.
	EnterCompExpr(c *CompExprContext)

	// EnterMulExpr is called when entering the mulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterIfExpr is called when entering the ifExpr production.
	EnterIfExpr(c *IfExprContext)

	// EnterPostfixExpr is called when entering the postfixExpr production.
	EnterPostfixExpr(c *PostfixExprContext)

	// EnterInitExpr is called when entering the initExpr production.
	EnterInitExpr(c *InitExprContext)

	// EnterBitAndExpr is called when entering the bitAndExpr production.
	EnterBitAndExpr(c *BitAndExprContext)

	// EnterPowExpr is called when entering the powExpr production.
	EnterPowExpr(c *PowExprContext)

	// EnterIdExpr is called when entering the idExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterInitArgs is called when entering the initArgs production.
	EnterInitArgs(c *InitArgsContext)

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

	// ExitFullType is called when exiting the fullType production.
	ExitFullType(c *FullTypeContext)

	// ExitPtrType is called when exiting the ptrType production.
	ExitPtrType(c *PtrTypeContext)

	// ExitBaseType is called when exiting the baseType production.
	ExitBaseType(c *BaseTypeContext)

	// ExitSliceType is called when exiting the sliceType production.
	ExitSliceType(c *SliceTypeContext)

	// ExitCastExpr is called when exiting the castExpr production.
	ExitCastExpr(c *CastExprContext)

	// ExitAllocExpr is called when exiting the allocExpr production.
	ExitAllocExpr(c *AllocExprContext)

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

	// ExitPtrExpr is called when exiting the ptrExpr production.
	ExitPtrExpr(c *PtrExprContext)

	// ExitPrefixExpr is called when exiting the prefixExpr production.
	ExitPrefixExpr(c *PrefixExprContext)

	// ExitDeallocExpr is called when exiting the deallocExpr production.
	ExitDeallocExpr(c *DeallocExprContext)

	// ExitBitOrExpr is called when exiting the bitOrExpr production.
	ExitBitOrExpr(c *BitOrExprContext)

	// ExitAddExpr is called when exiting the addExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitCompExpr is called when exiting the compExpr production.
	ExitCompExpr(c *CompExprContext)

	// ExitMulExpr is called when exiting the mulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitIfExpr is called when exiting the ifExpr production.
	ExitIfExpr(c *IfExprContext)

	// ExitPostfixExpr is called when exiting the postfixExpr production.
	ExitPostfixExpr(c *PostfixExprContext)

	// ExitInitExpr is called when exiting the initExpr production.
	ExitInitExpr(c *InitExprContext)

	// ExitBitAndExpr is called when exiting the bitAndExpr production.
	ExitBitAndExpr(c *BitAndExprContext)

	// ExitPowExpr is called when exiting the powExpr production.
	ExitPowExpr(c *PowExprContext)

	// ExitIdExpr is called when exiting the idExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitInitArgs is called when exiting the initArgs production.
	ExitInitArgs(c *InitArgsContext)
}
