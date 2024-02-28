package dto

import "github.com/Easy-Job-Developer/catalog_plus/domain/records/models"

type ColumnPair struct {
	Input struct {
		SheetName string        `json:"sheetName"`
		Column    models.Column `json:"column"`
	} `json:"input"`
	Output struct {
		SheetName string        `json:"sheetName"`
		Column    models.Column `json:"column"`
	} `json:"output"`
}
