//go:build js && wasm
// +build js,wasm

package main

import (
	gen "excel_wasm/gen"
	"syscall/js"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("excelExecute", js.FuncOf(execute))
	<-done
}

func execute(this js.Value, inputs []js.Value) interface{} {
	input := antlr.NewInputStream(inputs[0].String())
	lexer := gen.NewExcelLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := gen.NewExcelParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Formula()
	return tree.ToStringTree(nil, p)
}
