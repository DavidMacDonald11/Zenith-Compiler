// Code generated from ZenithParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ZenithParser

import "github.com/antlr4-go/antlr/v4"

// ZenithParserListener is a complete listener for a parse tree produced by ZenithParser.
type ZenithParserListener interface {
	antlr.ParseTreeListener

	// EnterFileStat is called when entering the fileStat production.
	EnterFileStat(c *FileStatContext)

	// EnterAddExpr is called when entering the addExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterMulExpr is called when entering the mulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterNumExpr is called when entering the numExpr production.
	EnterNumExpr(c *NumExprContext)

	// ExitFileStat is called when exiting the fileStat production.
	ExitFileStat(c *FileStatContext)

	// ExitAddExpr is called when exiting the addExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitMulExpr is called when exiting the mulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitNumExpr is called when exiting the numExpr production.
	ExitNumExpr(c *NumExprContext)
}
