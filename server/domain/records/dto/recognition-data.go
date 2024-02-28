package dto

import (
	"github.com/Easy-Job-Developer/catalog_plus/domain/records/models"
)

type RecognitionData struct {
	DefaultHeader   []Cell            `json:"header"`
	RecognitionUnit []RecognitionUnit `json:"recognition-units"`
	Sheet           string            `json:"sheet"`
}
type RecognitionUnit struct {
	Data  [][]Cell `json:"data"`
	Count int      `json:"count"`
	Index int      `json:"index"`
}
type Cell struct {
	Content string `json:"content"`
	Row     int    `json:"row"`
	Col     int    `json:"col"`
}

func (rd *RecognitionData) InitRecognitionData(set models.Set, rows [][]string) {
	q := 0
	for key, val := range set.UnitedSimilarSets {
		if key != set.ModeKey {
			var newData [][]Cell
			for _, rowIndex := range val {
				if rowIndex >= 0 && rowIndex < len(rows) {
					var newRow []Cell
					for colIndex, content := range rows[rowIndex-1] {
						if content != "" {
							newRow = append(newRow, Cell{Content: content, Row: rowIndex, Col: colIndex})
						}
					}
					newData = append(newData, newRow)
				}
			}
			rUnit := RecognitionUnit{
				Data:  newData,
				Count: len(val),
				Index: q,
			}

			rd.RecognitionUnit = append(rd.RecognitionUnit, rUnit)
			q++
		}
	}
}
