package records

import (
	"github.com/Easy-Job-Developer/catalog_plus/domain/records/models"
	"github.com/xuri/excelize/v2"
)

func WriteExcelData(sheets []models.Sheet) (*excelize.File, error) {
	f := excelize.NewFile()
	for _, sheet := range sheets {
		index, _ := f.NewSheet(sheet.Title)

		for _, col := range sheet.Columns {
			colIndex := col.Index.Column
			for rowIdx, val := range col.Values {
				cell, err := excelize.CoordinatesToCellName(colIndex, rowIdx+1) // Excel рядки починаються з 1
				if err != nil {
					return nil, err
				}
				if err := f.SetCellValue(sheet.Title, cell, val); err != nil {
					return nil, err
				}
			}
		}

		if index == 1 {
			f.SetActiveSheet(index)
		}
	}
	f.DeleteSheet("Sheet1")

	return f, nil
}
