// Code generated from Excel.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ExcelLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var excellexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func excellexerLexerInit() {
	staticData := &excellexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"ADD", "SUB", "MUL", "DIV", "EQ", "NEQ", "LT", "GT", "LTE", "GTE", "AND",
		"OR", "NOT", "XOR", "LPAREN", "RPAREN", "COLON", "COMMA", "WS", "VAR_COL",
		"VAR_ROW", "VAR_NAME", "VAR_RANGE", "INT", "NUM", "DEC", "NUMERIC",
		"STRING", "BOOL", "FN_SUM", "FN_AVERAGE", "FN_COUNT", "FN_MIN", "FN_MAX",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 28, 218, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 4, 18, 117, 8,
		18, 11, 18, 12, 18, 118, 1, 18, 1, 18, 1, 19, 4, 19, 124, 8, 19, 11, 19,
		12, 19, 125, 1, 20, 1, 20, 5, 20, 130, 8, 20, 10, 20, 12, 20, 133, 9, 20,
		1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 3, 23, 143, 8,
		23, 1, 23, 1, 23, 5, 23, 147, 8, 23, 10, 23, 12, 23, 150, 9, 23, 1, 24,
		4, 24, 153, 8, 24, 11, 24, 12, 24, 154, 1, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 3, 26, 162, 8, 26, 1, 27, 1, 27, 5, 27, 166, 8, 27, 10, 27, 12, 27,
		169, 9, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 3, 28, 191, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 0, 0, 34,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 0, 35, 17, 37, 18, 39, 0, 41,
		0, 43, 19, 45, 20, 47, 0, 49, 0, 51, 0, 53, 21, 55, 22, 57, 23, 59, 24,
		61, 25, 63, 26, 65, 27, 67, 28, 1, 0, 6, 2, 0, 9, 10, 32, 32, 1, 0, 65,
		90, 1, 0, 49, 57, 1, 0, 48, 57, 1, 0, 45, 45, 1, 0, 34, 34, 222, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0,
		0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1,
		0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67,
		1, 0, 0, 0, 1, 69, 1, 0, 0, 0, 3, 71, 1, 0, 0, 0, 5, 73, 1, 0, 0, 0, 7,
		75, 1, 0, 0, 0, 9, 77, 1, 0, 0, 0, 11, 79, 1, 0, 0, 0, 13, 82, 1, 0, 0,
		0, 15, 84, 1, 0, 0, 0, 17, 86, 1, 0, 0, 0, 19, 89, 1, 0, 0, 0, 21, 92,
		1, 0, 0, 0, 23, 96, 1, 0, 0, 0, 25, 99, 1, 0, 0, 0, 27, 103, 1, 0, 0, 0,
		29, 107, 1, 0, 0, 0, 31, 109, 1, 0, 0, 0, 33, 111, 1, 0, 0, 0, 35, 113,
		1, 0, 0, 0, 37, 116, 1, 0, 0, 0, 39, 123, 1, 0, 0, 0, 41, 127, 1, 0, 0,
		0, 43, 134, 1, 0, 0, 0, 45, 137, 1, 0, 0, 0, 47, 142, 1, 0, 0, 0, 49, 152,
		1, 0, 0, 0, 51, 156, 1, 0, 0, 0, 53, 159, 1, 0, 0, 0, 55, 163, 1, 0, 0,
		0, 57, 190, 1, 0, 0, 0, 59, 192, 1, 0, 0, 0, 61, 196, 1, 0, 0, 0, 63, 204,
		1, 0, 0, 0, 65, 210, 1, 0, 0, 0, 67, 214, 1, 0, 0, 0, 69, 70, 5, 43, 0,
		0, 70, 2, 1, 0, 0, 0, 71, 72, 5, 45, 0, 0, 72, 4, 1, 0, 0, 0, 73, 74, 5,
		42, 0, 0, 74, 6, 1, 0, 0, 0, 75, 76, 5, 47, 0, 0, 76, 8, 1, 0, 0, 0, 77,
		78, 5, 61, 0, 0, 78, 10, 1, 0, 0, 0, 79, 80, 5, 33, 0, 0, 80, 81, 5, 61,
		0, 0, 81, 12, 1, 0, 0, 0, 82, 83, 5, 60, 0, 0, 83, 14, 1, 0, 0, 0, 84,
		85, 5, 62, 0, 0, 85, 16, 1, 0, 0, 0, 86, 87, 5, 60, 0, 0, 87, 88, 5, 61,
		0, 0, 88, 18, 1, 0, 0, 0, 89, 90, 5, 62, 0, 0, 90, 91, 5, 61, 0, 0, 91,
		20, 1, 0, 0, 0, 92, 93, 5, 65, 0, 0, 93, 94, 5, 78, 0, 0, 94, 95, 5, 68,
		0, 0, 95, 22, 1, 0, 0, 0, 96, 97, 5, 79, 0, 0, 97, 98, 5, 82, 0, 0, 98,
		24, 1, 0, 0, 0, 99, 100, 5, 78, 0, 0, 100, 101, 5, 79, 0, 0, 101, 102,
		5, 84, 0, 0, 102, 26, 1, 0, 0, 0, 103, 104, 5, 88, 0, 0, 104, 105, 5, 79,
		0, 0, 105, 106, 5, 82, 0, 0, 106, 28, 1, 0, 0, 0, 107, 108, 5, 40, 0, 0,
		108, 30, 1, 0, 0, 0, 109, 110, 5, 41, 0, 0, 110, 32, 1, 0, 0, 0, 111, 112,
		5, 58, 0, 0, 112, 34, 1, 0, 0, 0, 113, 114, 5, 44, 0, 0, 114, 36, 1, 0,
		0, 0, 115, 117, 7, 0, 0, 0, 116, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0,
		118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120,
		121, 6, 18, 0, 0, 121, 38, 1, 0, 0, 0, 122, 124, 7, 1, 0, 0, 123, 122,
		1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0,
		0, 0, 126, 40, 1, 0, 0, 0, 127, 131, 7, 2, 0, 0, 128, 130, 7, 3, 0, 0,
		129, 128, 1, 0, 0, 0, 130, 133, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131,
		132, 1, 0, 0, 0, 132, 42, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 134, 135, 3,
		39, 19, 0, 135, 136, 3, 41, 20, 0, 136, 44, 1, 0, 0, 0, 137, 138, 3, 43,
		21, 0, 138, 139, 3, 33, 16, 0, 139, 140, 3, 43, 21, 0, 140, 46, 1, 0, 0,
		0, 141, 143, 7, 4, 0, 0, 142, 141, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143,
		144, 1, 0, 0, 0, 144, 148, 7, 2, 0, 0, 145, 147, 7, 3, 0, 0, 146, 145,
		1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0,
		0, 0, 149, 48, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151, 153, 7, 3, 0, 0,
		152, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154,
		155, 1, 0, 0, 0, 155, 50, 1, 0, 0, 0, 156, 157, 5, 46, 0, 0, 157, 158,
		3, 49, 24, 0, 158, 52, 1, 0, 0, 0, 159, 161, 3, 47, 23, 0, 160, 162, 3,
		51, 25, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162, 54, 1, 0, 0,
		0, 163, 167, 5, 34, 0, 0, 164, 166, 8, 5, 0, 0, 165, 164, 1, 0, 0, 0, 166,
		169, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 170,
		1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170, 171, 5, 34, 0, 0, 171, 56, 1, 0,
		0, 0, 172, 173, 5, 84, 0, 0, 173, 174, 5, 82, 0, 0, 174, 175, 5, 85, 0,
		0, 175, 191, 5, 69, 0, 0, 176, 177, 5, 70, 0, 0, 177, 178, 5, 65, 0, 0,
		178, 179, 5, 76, 0, 0, 179, 180, 5, 83, 0, 0, 180, 191, 5, 69, 0, 0, 181,
		182, 5, 116, 0, 0, 182, 183, 5, 114, 0, 0, 183, 184, 5, 117, 0, 0, 184,
		191, 5, 101, 0, 0, 185, 186, 5, 102, 0, 0, 186, 187, 5, 97, 0, 0, 187,
		188, 5, 108, 0, 0, 188, 189, 5, 115, 0, 0, 189, 191, 5, 101, 0, 0, 190,
		172, 1, 0, 0, 0, 190, 176, 1, 0, 0, 0, 190, 181, 1, 0, 0, 0, 190, 185,
		1, 0, 0, 0, 191, 58, 1, 0, 0, 0, 192, 193, 5, 83, 0, 0, 193, 194, 5, 85,
		0, 0, 194, 195, 5, 77, 0, 0, 195, 60, 1, 0, 0, 0, 196, 197, 5, 65, 0, 0,
		197, 198, 5, 86, 0, 0, 198, 199, 5, 69, 0, 0, 199, 200, 5, 82, 0, 0, 200,
		201, 5, 65, 0, 0, 201, 202, 5, 71, 0, 0, 202, 203, 5, 69, 0, 0, 203, 62,
		1, 0, 0, 0, 204, 205, 5, 67, 0, 0, 205, 206, 5, 79, 0, 0, 206, 207, 5,
		85, 0, 0, 207, 208, 5, 78, 0, 0, 208, 209, 5, 84, 0, 0, 209, 64, 1, 0,
		0, 0, 210, 211, 5, 77, 0, 0, 211, 212, 5, 73, 0, 0, 212, 213, 5, 78, 0,
		0, 213, 66, 1, 0, 0, 0, 214, 215, 5, 77, 0, 0, 215, 216, 5, 65, 0, 0, 216,
		217, 5, 88, 0, 0, 217, 68, 1, 0, 0, 0, 10, 0, 118, 125, 131, 142, 148,
		154, 161, 167, 190, 1, 6, 0, 0,
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

// ExcelLexerInit initializes any static state used to implement ExcelLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewExcelLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExcelLexerInit() {
	staticData := &excellexerLexerStaticData
	staticData.once.Do(excellexerLexerInit)
}

// NewExcelLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewExcelLexer(input antlr.CharStream) *ExcelLexer {
	ExcelLexerInit()
	l := new(ExcelLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &excellexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Excel.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExcelLexer tokens.
const (
	ExcelLexerADD        = 1
	ExcelLexerSUB        = 2
	ExcelLexerMUL        = 3
	ExcelLexerDIV        = 4
	ExcelLexerEQ         = 5
	ExcelLexerNEQ        = 6
	ExcelLexerLT         = 7
	ExcelLexerGT         = 8
	ExcelLexerLTE        = 9
	ExcelLexerGTE        = 10
	ExcelLexerAND        = 11
	ExcelLexerOR         = 12
	ExcelLexerNOT        = 13
	ExcelLexerXOR        = 14
	ExcelLexerLPAREN     = 15
	ExcelLexerRPAREN     = 16
	ExcelLexerCOMMA      = 17
	ExcelLexerWS         = 18
	ExcelLexerVAR_NAME   = 19
	ExcelLexerVAR_RANGE  = 20
	ExcelLexerNUMERIC    = 21
	ExcelLexerSTRING     = 22
	ExcelLexerBOOL       = 23
	ExcelLexerFN_SUM     = 24
	ExcelLexerFN_AVERAGE = 25
	ExcelLexerFN_COUNT   = 26
	ExcelLexerFN_MIN     = 27
	ExcelLexerFN_MAX     = 28
)
