package core

import (
	"encoding/json"
	"errors"
)

type CellPos struct {
	Row string
	Col string
}

type SheetCell struct {
	Formula         string    `json:"formula"`
	CalculatedValue string    `json:"calculatedValue"`
	IsCalculated    bool      `json:"isCalculated"`
	Error           string    `json:"error"`
	Dependencies    []CellPos `json:"dependencies"`
}

type SheetData struct {
	Data map[string]map[string]SheetCell `json:"data"`
}

func (s *SheetData) ToJSON() (string, error) {
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func NewSheetData(rawString string) (*SheetData, error) {
	var data SheetData
	err := json.Unmarshal([]byte(rawString), &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *SheetData) Get(row, col string) (*SheetCell, error) {
	if s.Data == nil {
		return nil, errors.New("sheet data is empty")
	}
	rowData, ok := s.Data[row]
	if !ok {
		return nil, errors.New("row not found")
	}
	cell, ok := rowData[col]
	if !ok {
		return nil, errors.New("cell not found")
	}
	return &cell, nil
}
