package core_test

import (
	"excel_wasm/core"
	"strings"
	"testing"
)

func TestCreateOneCellSheet(t *testing.T) {
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
	_, err := core.NewSheetData(input)
	if err != nil {
		t.Error(err)
	}

}

func TestCreateSheetWithInvalidInput(t *testing.T) {
	input := `invalid json`
	_, err := core.NewSheetData(input)
	if err == nil {
		t.Error("expected error")
	}
}

func TestGetExistedCell(t *testing.T) {
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
	sheetData, err := core.NewSheetData(input)
	if err != nil {
		t.Error(err)
	}
	cell, err := sheetData.Get("1", "A")
	if err != nil {
		t.Error(err)
	}
	if cell.Formula != "=1+1" {
		t.Error("cell formula is not correct")
	}
}

func TestGetNonExistedRow(t *testing.T) {
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
	sheetData, err := core.NewSheetData(input)
	if err != nil {
		t.Error(err)
	}
	_, err = sheetData.Get("2", "A")
	if err == nil {
		t.Error("expected error")
	}
}

func TestGetNonExistedColumn(t *testing.T) {
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
	sheetData, err := core.NewSheetData(input)
	if err != nil {
		t.Error(err)
	}
	_, err = sheetData.Get("1", "B")
	if err == nil {
		t.Error("expected error")
	}
}

func TestGetCellOnEmptySheet(t *testing.T) {
	input := `{
		"data": {}
	}`
	sheetData, err := core.NewSheetData(input)
	if err != nil {
		t.Error(err)
	}
	_, err = sheetData.Get("1", "A")
	if err == nil {
		t.Error("expected error")
	}
}

func TestToJSON(t *testing.T) {
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
	// remove all spaces and new lines
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\t", "")

	sheetData, err := core.NewSheetData(input)
	if err != nil {
		t.Error(err)
	}
	json, err := sheetData.ToJSON()
	if err != nil {
		t.Error(err)
	}
	if json != input {
		t.Log(json)
		t.Log(input)
		t.Error("json is not correct")
	}
}
