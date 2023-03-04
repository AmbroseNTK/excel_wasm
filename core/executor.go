package core

import gen "excel_wasm/gen"

type Executor struct {
	sheetData   *SheetData
	syntaxTrees map[string]map[string]*gen.IFormulaContext
}

func (e *Executor) Execute(rawInput string) (string, error) {
	sheetData, err := NewSheetData(rawInput)
	if err != nil {
		return "", err
	}

	return "", nil
}
