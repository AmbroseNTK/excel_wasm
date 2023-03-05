package core_test

import (
	"excel_wasm/core"
	"testing"
)

func TestCreateNewExecutor(t *testing.T) {
	input := `{
		"data": {
			"1": {
				"A": {
					"formula": "=1+1",
					"calculatedValue": "2",
					"isCalculated": true,
					"error": "",
					"dependencies": []
				}
			}
		}
	}`
	executor := core.NewExecutor(input)
	if executor == nil {
		t.Error("expected executor")
	}
}

func TestCreateNewExecutorMustFailed(t *testing.T) {
	input := `invalid json`
	executor := core.NewExecutor(input)
	if executor != nil {
		t.Error("expected nil executor")
	}
}
