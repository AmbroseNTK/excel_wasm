package core

import (
	gen "excel_wasm/gen"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Executor struct {
	sheetData   *SheetData
	syntaxTrees map[string]map[string]*gen.IFormulaContext
}

func NewExecutor(rawInput string) *Executor {
	sheetData, err := NewSheetData(rawInput)
	if err != nil {
		return nil
	}
	return &Executor{
		sheetData:   sheetData,
		syntaxTrees: make(map[string]map[string]*gen.IFormulaContext),
	}
}

func (e *Executor) ParseSheetToSyntaxTrees() {
	for row, rowData := range e.sheetData.Data {
		for col, cell := range rowData {
			if cell.Formula == "" {
				continue
			}
			if _, ok := e.syntaxTrees[row]; !ok {
				e.syntaxTrees[row] = make(map[string]*gen.IFormulaContext)
			}
			lexer := gen.NewExcelLexer(antlr.NewInputStream(cell.Formula))
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := gen.NewExcelParser(stream)
			errorListener := antlr.NewDiagnosticErrorListener(true)
			p.AddErrorListener(errorListener)
			p.BuildParseTrees = true
			tree := p.Formula()
			e.syntaxTrees[row][col] = &tree

		}
	}
}

func (e *Executor) Execute(rawInput string) (string, error) {
	// sheetData, err := NewSheetData(rawInput)
	// if err != nil {
	// 	return "", err
	// }

	return "", nil
}
