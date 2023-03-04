// Code generated from Excel.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Excel

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// ExcelListener is a complete listener for a parse tree produced by ExcelParser.
type ExcelListener interface {
	antlr.ParseTreeListener

	// EnterFormula is called when entering the formula production.
	EnterFormula(c *FormulaContext)

	// EnterFunction_name is called when entering the function_name production.
	EnterFunction_name(c *Function_nameContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterNumeric_operator is called when entering the numeric_operator production.
	EnterNumeric_operator(c *Numeric_operatorContext)

	// EnterLogic_operator is called when entering the logic_operator production.
	EnterLogic_operator(c *Logic_operatorContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitFormula is called when exiting the formula production.
	ExitFormula(c *FormulaContext)

	// ExitFunction_name is called when exiting the function_name production.
	ExitFunction_name(c *Function_nameContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitNumeric_operator is called when exiting the numeric_operator production.
	ExitNumeric_operator(c *Numeric_operatorContext)

	// ExitLogic_operator is called when exiting the logic_operator production.
	ExitLogic_operator(c *Logic_operatorContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
