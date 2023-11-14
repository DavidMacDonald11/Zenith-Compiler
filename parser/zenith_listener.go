// Code generated from Zenith.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Zenith

import "github.com/antlr4-go/antlr/v4"

// ZenithListener is a complete listener for a parse tree produced by ZenithParser.
type ZenithListener interface {
	antlr.ParseTreeListener

	// EnterFileStat is called when entering the fileStat production.
	EnterFileStat(c *FileStatContext)

	// EnterMulExpr is called when entering the MulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterNum is called when entering the Num production.
	EnterNum(c *NumContext)

	// EnterAddExpr is called when entering the AddExpr production.
	EnterAddExpr(c *AddExprContext)

	// ExitFileStat is called when exiting the fileStat production.
	ExitFileStat(c *FileStatContext)

	// ExitMulExpr is called when exiting the MulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitNum is called when exiting the Num production.
	ExitNum(c *NumContext)

	// ExitAddExpr is called when exiting the AddExpr production.
	ExitAddExpr(c *AddExprContext)
}
