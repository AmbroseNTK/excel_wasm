// Code generated from Excel.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Excel

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ExcelParser struct {
	*antlr.BaseParser
}

var excelParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func excelParserInit() {
	staticData := &excelParserStaticData
	staticData.literalNames = []string{
		"", "'+'", "'-'", "'*'", "'/'", "'='", "'!='", "'<'", "'>'", "'<='",
		"'>='", "'AND'", "'OR'", "'NOT'", "'XOR'", "'('", "')'", "','", "",
		"", "", "", "", "", "'SUM'", "'AVERAGE'", "'COUNT'", "'MIN'", "'MAX'",
	}
	staticData.symbolicNames = []string{
		"", "ADD", "SUB", "MUL", "DIV", "EQ", "NEQ", "LT", "GT", "LTE", "GTE",
		"AND", "OR", "NOT", "XOR", "LPAREN", "RPAREN", "COMMA", "WS", "VAR_NAME",
		"VAR_RANGE", "NUMERIC", "STRING", "BOOL", "FN_SUM", "FN_AVERAGE", "FN_COUNT",
		"FN_MIN", "FN_MAX",
	}
	staticData.ruleNames = []string{
		"formula", "function_name", "function_call", "operand", "numeric_operator",
		"logic_operator", "term", "expr",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 73, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 28, 8, 2, 10, 2, 12, 2, 31, 9, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 41, 8, 3, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 54, 8, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 60, 8, 7, 10, 7, 12, 7, 63, 9, 7, 1, 7, 1,
		7, 1, 7, 5, 7, 68, 8, 7, 10, 7, 12, 7, 71, 9, 7, 1, 7, 0, 0, 8, 0, 2, 4,
		6, 8, 10, 12, 14, 0, 3, 1, 0, 24, 28, 1, 0, 1, 4, 1, 0, 5, 14, 74, 0, 16,
		1, 0, 0, 0, 2, 20, 1, 0, 0, 0, 4, 22, 1, 0, 0, 0, 6, 40, 1, 0, 0, 0, 8,
		42, 1, 0, 0, 0, 10, 44, 1, 0, 0, 0, 12, 53, 1, 0, 0, 0, 14, 55, 1, 0, 0,
		0, 16, 17, 5, 5, 0, 0, 17, 18, 3, 14, 7, 0, 18, 19, 5, 0, 0, 1, 19, 1,
		1, 0, 0, 0, 20, 21, 7, 0, 0, 0, 21, 3, 1, 0, 0, 0, 22, 23, 3, 2, 1, 0,
		23, 24, 5, 15, 0, 0, 24, 29, 3, 14, 7, 0, 25, 26, 5, 17, 0, 0, 26, 28,
		3, 14, 7, 0, 27, 25, 1, 0, 0, 0, 28, 31, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0,
		29, 30, 1, 0, 0, 0, 30, 32, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 32, 33, 5,
		16, 0, 0, 33, 5, 1, 0, 0, 0, 34, 41, 5, 19, 0, 0, 35, 41, 5, 20, 0, 0,
		36, 41, 3, 4, 2, 0, 37, 41, 5, 22, 0, 0, 38, 41, 5, 23, 0, 0, 39, 41, 5,
		21, 0, 0, 40, 34, 1, 0, 0, 0, 40, 35, 1, 0, 0, 0, 40, 36, 1, 0, 0, 0, 40,
		37, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 39, 1, 0, 0, 0, 41, 7, 1, 0, 0,
		0, 42, 43, 7, 1, 0, 0, 43, 9, 1, 0, 0, 0, 44, 45, 7, 2, 0, 0, 45, 11, 1,
		0, 0, 0, 46, 54, 3, 6, 3, 0, 47, 48, 5, 15, 0, 0, 48, 49, 3, 14, 7, 0,
		49, 50, 5, 16, 0, 0, 50, 54, 1, 0, 0, 0, 51, 52, 5, 2, 0, 0, 52, 54, 3,
		14, 7, 0, 53, 46, 1, 0, 0, 0, 53, 47, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 54,
		13, 1, 0, 0, 0, 55, 61, 3, 12, 6, 0, 56, 57, 3, 8, 4, 0, 57, 58, 3, 12,
		6, 0, 58, 60, 1, 0, 0, 0, 59, 56, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59,
		1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 69, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0,
		64, 65, 3, 10, 5, 0, 65, 66, 3, 12, 6, 0, 66, 68, 1, 0, 0, 0, 67, 64, 1,
		0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70,
		15, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 5, 29, 40, 53, 61, 69,
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

// ExcelParserInit initializes any static state used to implement ExcelParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewExcelParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExcelParserInit() {
	staticData := &excelParserStaticData
	staticData.once.Do(excelParserInit)
}

// NewExcelParser produces a new parser instance for the optional input antlr.TokenStream.
func NewExcelParser(input antlr.TokenStream) *ExcelParser {
	ExcelParserInit()
	this := new(ExcelParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &excelParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Excel.g4"

	return this
}

// ExcelParser tokens.
const (
	ExcelParserEOF        = antlr.TokenEOF
	ExcelParserADD        = 1
	ExcelParserSUB        = 2
	ExcelParserMUL        = 3
	ExcelParserDIV        = 4
	ExcelParserEQ         = 5
	ExcelParserNEQ        = 6
	ExcelParserLT         = 7
	ExcelParserGT         = 8
	ExcelParserLTE        = 9
	ExcelParserGTE        = 10
	ExcelParserAND        = 11
	ExcelParserOR         = 12
	ExcelParserNOT        = 13
	ExcelParserXOR        = 14
	ExcelParserLPAREN     = 15
	ExcelParserRPAREN     = 16
	ExcelParserCOMMA      = 17
	ExcelParserWS         = 18
	ExcelParserVAR_NAME   = 19
	ExcelParserVAR_RANGE  = 20
	ExcelParserNUMERIC    = 21
	ExcelParserSTRING     = 22
	ExcelParserBOOL       = 23
	ExcelParserFN_SUM     = 24
	ExcelParserFN_AVERAGE = 25
	ExcelParserFN_COUNT   = 26
	ExcelParserFN_MIN     = 27
	ExcelParserFN_MAX     = 28
)

// ExcelParser rules.
const (
	ExcelParserRULE_formula          = 0
	ExcelParserRULE_function_name    = 1
	ExcelParserRULE_function_call    = 2
	ExcelParserRULE_operand          = 3
	ExcelParserRULE_numeric_operator = 4
	ExcelParserRULE_logic_operator   = 5
	ExcelParserRULE_term             = 6
	ExcelParserRULE_expr             = 7
)

// IFormulaContext is an interface to support dynamic dispatch.
type IFormulaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	Expr() IExprContext
	EOF() antlr.TerminalNode

	// IsFormulaContext differentiates from other interfaces.
	IsFormulaContext()
}

type FormulaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormulaContext() *FormulaContext {
	var p = new(FormulaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcelParserRULE_formula
	return p
}

func (*FormulaContext) IsFormulaContext() {}

func NewFormulaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormulaContext {
	var p = new(FormulaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcelParserRULE_formula

	return p
}

func (s *FormulaContext) GetParser() antlr.Parser { return s.parser }

func (s *FormulaContext) EQ() antlr.TerminalNode {
	return s.GetToken(ExcelParserEQ, 0)
}

func (s *FormulaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FormulaContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExcelParserEOF, 0)
}

func (s *FormulaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormulaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormulaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.EnterFormula(s)
	}
}

func (s *FormulaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.ExitFormula(s)
	}
}

func (p *ExcelParser) Formula() (localctx IFormulaContext) {
	this := p
	_ = this

	localctx = NewFormulaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExcelParserRULE_formula)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Match(ExcelParserEQ)
	}
	{
		p.SetState(17)
		p.Expr()
	}
	{
		p.SetState(18)
		p.Match(ExcelParserEOF)
	}

	return localctx
}

// IFunction_nameContext is an interface to support dynamic dispatch.
type IFunction_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FN_SUM() antlr.TerminalNode
	FN_AVERAGE() antlr.TerminalNode
	FN_COUNT() antlr.TerminalNode
	FN_MIN() antlr.TerminalNode
	FN_MAX() antlr.TerminalNode

	// IsFunction_nameContext differentiates from other interfaces.
	IsFunction_nameContext()
}

type Function_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_nameContext() *Function_nameContext {
	var p = new(Function_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcelParserRULE_function_name
	return p
}

func (*Function_nameContext) IsFunction_nameContext() {}

func NewFunction_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_nameContext {
	var p = new(Function_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcelParserRULE_function_name

	return p
}

func (s *Function_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_nameContext) FN_SUM() antlr.TerminalNode {
	return s.GetToken(ExcelParserFN_SUM, 0)
}

func (s *Function_nameContext) FN_AVERAGE() antlr.TerminalNode {
	return s.GetToken(ExcelParserFN_AVERAGE, 0)
}

func (s *Function_nameContext) FN_COUNT() antlr.TerminalNode {
	return s.GetToken(ExcelParserFN_COUNT, 0)
}

func (s *Function_nameContext) FN_MIN() antlr.TerminalNode {
	return s.GetToken(ExcelParserFN_MIN, 0)
}

func (s *Function_nameContext) FN_MAX() antlr.TerminalNode {
	return s.GetToken(ExcelParserFN_MAX, 0)
}

func (s *Function_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.EnterFunction_name(s)
	}
}

func (s *Function_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.ExitFunction_name(s)
	}
}

func (p *ExcelParser) Function_name() (localctx IFunction_nameContext) {
	this := p
	_ = this

	localctx = NewFunction_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExcelParserRULE_function_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&520093696) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function_name() IFunction_nameContext
	LPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcelParserRULE_function_call
	return p
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcelParserRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) Function_name() IFunction_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_nameContext)
}

func (s *Function_callContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ExcelParserLPAREN, 0)
}

func (s *Function_callContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Function_callContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Function_callContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ExcelParserRPAREN, 0)
}

func (s *Function_callContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ExcelParserCOMMA)
}

func (s *Function_callContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ExcelParserCOMMA, i)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (p *ExcelParser) Function_call() (localctx IFunction_callContext) {
	this := p
	_ = this

	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ExcelParserRULE_function_call)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Function_name()
	}
	{
		p.SetState(23)
		p.Match(ExcelParserLPAREN)
	}
	{
		p.SetState(24)
		p.Expr()
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExcelParserCOMMA {
		{
			p.SetState(25)
			p.Match(ExcelParserCOMMA)
		}
		{
			p.SetState(26)
			p.Expr()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
		p.Match(ExcelParserRPAREN)
	}

	return localctx
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_NAME() antlr.TerminalNode
	VAR_RANGE() antlr.TerminalNode
	Function_call() IFunction_callContext
	STRING() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	NUMERIC() antlr.TerminalNode

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcelParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcelParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) VAR_NAME() antlr.TerminalNode {
	return s.GetToken(ExcelParserVAR_NAME, 0)
}

func (s *OperandContext) VAR_RANGE() antlr.TerminalNode {
	return s.GetToken(ExcelParserVAR_RANGE, 0)
}

func (s *OperandContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *OperandContext) STRING() antlr.TerminalNode {
	return s.GetToken(ExcelParserSTRING, 0)
}

func (s *OperandContext) BOOL() antlr.TerminalNode {
	return s.GetToken(ExcelParserBOOL, 0)
}

func (s *OperandContext) NUMERIC() antlr.TerminalNode {
	return s.GetToken(ExcelParserNUMERIC, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (p *ExcelParser) Operand() (localctx IOperandContext) {
	this := p
	_ = this

	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ExcelParserRULE_operand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExcelParserVAR_NAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(ExcelParserVAR_NAME)
		}

	case ExcelParserVAR_RANGE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Match(ExcelParserVAR_RANGE)
		}

	case ExcelParserFN_SUM, ExcelParserFN_AVERAGE, ExcelParserFN_COUNT, ExcelParserFN_MIN, ExcelParserFN_MAX:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.Function_call()
		}

	case ExcelParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(37)
			p.Match(ExcelParserSTRING)
		}

	case ExcelParserBOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(38)
			p.Match(ExcelParserBOOL)
		}

	case ExcelParserNUMERIC:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(39)
			p.Match(ExcelParserNUMERIC)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumeric_operatorContext is an interface to support dynamic dispatch.
type INumeric_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode

	// IsNumeric_operatorContext differentiates from other interfaces.
	IsNumeric_operatorContext()
}

type Numeric_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumeric_operatorContext() *Numeric_operatorContext {
	var p = new(Numeric_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcelParserRULE_numeric_operator
	return p
}

func (*Numeric_operatorContext) IsNumeric_operatorContext() {}

func NewNumeric_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numeric_operatorContext {
	var p = new(Numeric_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcelParserRULE_numeric_operator

	return p
}

func (s *Numeric_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Numeric_operatorContext) ADD() antlr.TerminalNode {
	return s.GetToken(ExcelParserADD, 0)
}

func (s *Numeric_operatorContext) SUB() antlr.TerminalNode {
	return s.GetToken(ExcelParserSUB, 0)
}

func (s *Numeric_operatorContext) MUL() antlr.TerminalNode {
	return s.GetToken(ExcelParserMUL, 0)
}

func (s *Numeric_operatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(ExcelParserDIV, 0)
}

func (s *Numeric_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numeric_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numeric_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.EnterNumeric_operator(s)
	}
}

func (s *Numeric_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.ExitNumeric_operator(s)
	}
}

func (p *ExcelParser) Numeric_operator() (localctx INumeric_operatorContext) {
	this := p
	_ = this

	localctx = NewNumeric_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ExcelParserRULE_numeric_operator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILogic_operatorContext is an interface to support dynamic dispatch.
type ILogic_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GTE() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	XOR() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsLogic_operatorContext differentiates from other interfaces.
	IsLogic_operatorContext()
}

type Logic_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogic_operatorContext() *Logic_operatorContext {
	var p = new(Logic_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcelParserRULE_logic_operator
	return p
}

func (*Logic_operatorContext) IsLogic_operatorContext() {}

func NewLogic_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logic_operatorContext {
	var p = new(Logic_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcelParserRULE_logic_operator

	return p
}

func (s *Logic_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Logic_operatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(ExcelParserEQ, 0)
}

func (s *Logic_operatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(ExcelParserNEQ, 0)
}

func (s *Logic_operatorContext) LT() antlr.TerminalNode {
	return s.GetToken(ExcelParserLT, 0)
}

func (s *Logic_operatorContext) GT() antlr.TerminalNode {
	return s.GetToken(ExcelParserGT, 0)
}

func (s *Logic_operatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(ExcelParserLTE, 0)
}

func (s *Logic_operatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(ExcelParserGTE, 0)
}

func (s *Logic_operatorContext) AND() antlr.TerminalNode {
	return s.GetToken(ExcelParserAND, 0)
}

func (s *Logic_operatorContext) OR() antlr.TerminalNode {
	return s.GetToken(ExcelParserOR, 0)
}

func (s *Logic_operatorContext) XOR() antlr.TerminalNode {
	return s.GetToken(ExcelParserXOR, 0)
}

func (s *Logic_operatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(ExcelParserNOT, 0)
}

func (s *Logic_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logic_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logic_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.EnterLogic_operator(s)
	}
}

func (s *Logic_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.ExitLogic_operator(s)
	}
}

func (p *ExcelParser) Logic_operator() (localctx ILogic_operatorContext) {
	this := p
	_ = this

	localctx = NewLogic_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ExcelParserRULE_logic_operator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32736) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	SUB() antlr.TerminalNode

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcelParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcelParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *TermContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ExcelParserLPAREN, 0)
}

func (s *TermContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TermContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ExcelParserRPAREN, 0)
}

func (s *TermContext) SUB() antlr.TerminalNode {
	return s.GetToken(ExcelParserSUB, 0)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *ExcelParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ExcelParserRULE_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(53)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExcelParserVAR_NAME, ExcelParserVAR_RANGE, ExcelParserNUMERIC, ExcelParserSTRING, ExcelParserBOOL, ExcelParserFN_SUM, ExcelParserFN_AVERAGE, ExcelParserFN_COUNT, ExcelParserFN_MIN, ExcelParserFN_MAX:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Operand()
		}

	case ExcelParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Match(ExcelParserLPAREN)
		}
		{
			p.SetState(48)
			p.Expr()
		}
		{
			p.SetState(49)
			p.Match(ExcelParserRPAREN)
		}

	case ExcelParserSUB:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(51)
			p.Match(ExcelParserSUB)
		}
		{
			p.SetState(52)
			p.Expr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext
	AllNumeric_operator() []INumeric_operatorContext
	Numeric_operator(i int) INumeric_operatorContext
	AllLogic_operator() []ILogic_operatorContext
	Logic_operator(i int) ILogic_operatorContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcelParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcelParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ExprContext) AllNumeric_operator() []INumeric_operatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumeric_operatorContext); ok {
			len++
		}
	}

	tst := make([]INumeric_operatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumeric_operatorContext); ok {
			tst[i] = t.(INumeric_operatorContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Numeric_operator(i int) INumeric_operatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumeric_operatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumeric_operatorContext)
}

func (s *ExprContext) AllLogic_operator() []ILogic_operatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogic_operatorContext); ok {
			len++
		}
	}

	tst := make([]ILogic_operatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogic_operatorContext); ok {
			tst[i] = t.(ILogic_operatorContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Logic_operator(i int) ILogic_operatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogic_operatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogic_operatorContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcelListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *ExcelParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ExcelParserRULE_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Term()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(56)
				p.Numeric_operator()
			}
			{
				p.SetState(57)
				p.Term()
			}

		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(64)
				p.Logic_operator()
			}
			{
				p.SetState(65)
				p.Term()
			}

		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}
