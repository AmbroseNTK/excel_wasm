// Code generated from Excel.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Excel

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseExcelListener is a complete listener for a parse tree produced by ExcelParser.
type BaseExcelListener struct{}

var _ ExcelListener = &BaseExcelListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExcelListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExcelListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExcelListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExcelListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFormula is called when production formula is entered.
func (s *BaseExcelListener) EnterFormula(ctx *FormulaContext) {}

// ExitFormula is called when production formula is exited.
func (s *BaseExcelListener) ExitFormula(ctx *FormulaContext) {}

// EnterFunction_name is called when production function_name is entered.
func (s *BaseExcelListener) EnterFunction_name(ctx *Function_nameContext) {}

// ExitFunction_name is called when production function_name is exited.
func (s *BaseExcelListener) ExitFunction_name(ctx *Function_nameContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseExcelListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseExcelListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseExcelListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseExcelListener) ExitOperand(ctx *OperandContext) {}

// EnterNumeric_operator is called when production numeric_operator is entered.
func (s *BaseExcelListener) EnterNumeric_operator(ctx *Numeric_operatorContext) {}

// ExitNumeric_operator is called when production numeric_operator is exited.
func (s *BaseExcelListener) ExitNumeric_operator(ctx *Numeric_operatorContext) {}

// EnterLogic_operator is called when production logic_operator is entered.
func (s *BaseExcelListener) EnterLogic_operator(ctx *Logic_operatorContext) {}

// ExitLogic_operator is called when production logic_operator is exited.
func (s *BaseExcelListener) ExitLogic_operator(ctx *Logic_operatorContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseExcelListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseExcelListener) ExitTerm(ctx *TermContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseExcelListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseExcelListener) ExitExpr(ctx *ExprContext) {}
