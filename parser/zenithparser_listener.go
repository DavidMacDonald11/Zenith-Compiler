// Code generated from ZenithParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ZenithParser

import "github.com/antlr4-go/antlr/v4"

// ZenithParserListener is a complete listener for a parse tree produced by ZenithParser.
type ZenithParserListener interface {
	antlr.ParseTreeListener

	// EnterFileStat is called when entering the fileStat production.
	EnterFileStat(c *FileStatContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitFileStat is called when exiting the fileStat production.
	ExitFileStat(c *FileStatContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
