package records

import (
	"fmt"
	"mime/multipart"
	"sort"

	"github.com/Easy-Job-Developer/catalog_plus/domain/horoshop"
	"github.com/Easy-Job-Developer/catalog_plus/domain/records/dto"
	"github.com/Easy-Job-Developer/catalog_plus/domain/records/models"
	"github.com/xuri/excelize/v2"
)

func GetSheetList(file multipart.File) ([]string, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}
	defer func() {
		f.Close()
	}()

	sheets := f.GetSheetList()
	return sheets, nil
}

func ParseExcelSets(file multipart.File, selectedSheets []string) (*[]dto.RecognitionData, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}
	defer func() {
		f.Close()
	}()

	var rds []dto.RecognitionData
	for _, sh := range selectedSheets {
		rows, err := f.GetRows(sh)
		if err != nil {
			fmt.Println(err.Error())
		}

		var set models.Set

		set.InitSet(rows)
		set.FindMode()
		set.CountSimilarity()
		set.CombineSets(0.5)

		var rd dto.RecognitionData
		rd.InitRecognitionData(set, rows)
		rd.Sheet = sh

		sort.Ints(set.UnitedSimilarSets[set.ModeKey])
		row := set.UnitedSimilarSets[set.ModeKey][0] - 1
		for col, cell := range rows[row] {
			rd.DefaultHeader = append(rd.DefaultHeader, dto.Cell{
				Content: cell,
				Row:     row + 1,
				Col:     col,
			})
		}
		rds = append(rds, rd)
	}

	return &rds, nil
}

func ParseExcelData(
	file multipart.File,
	columnPairs []horoshop.ProductSheetsMap,
	selectedSheets []string,
	categorySetting interface{}) ([]horoshop.Product, []horoshop.Product, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		fmt.Println(err.Error())
		return nil, nil, err
	}
	defer func() {
		f.Close()
	}()

	var allProducts []horoshop.Product
	var failedProducts []horoshop.Product
	for _, sh := range selectedSheets {
		rows, err := f.GetRows(sh)
		if err != nil {
			return nil, nil, err
		}

		set := models.SetNew(rows, 0.5)
		productRows := set.UnitedSimilarSets[set.ModeKey]
		sort.Ints(productRows)

		var productMap *horoshop.ProductMap
		for _, cp := range columnPairs {
			if cp.Sheet == sh {
				productMap = &cp.ProductMap
				break
			}
		}

		if productMap != nil {
			p, pFailed := productMap.InitProductFromTable(rows, productRows)

			switch cs := categorySetting.(type) {
			case *dto.CategorySheet:
				var uId int
				for _, u := range cs.CategorySheetUnits {
					if u.SheetName == sh {
						uId = u.HoroshopCategoryID
						break
					}
				}
				for _, pr := range p {
					pr.Parent.ID = &uId
				}
				for _, pr := range pFailed {
					pr.Parent.ID = &uId
				}

			case *dto.CategoryFullFile:
				for _, pr := range p {
					pr.Parent.ID = &cs.HoroshopCategoryID
				}
				for _, pr := range pFailed {
					pr.Parent.ID = &cs.HoroshopCategoryID
				}

			case *dto.CategoryRow:
				var categories []dto.CategoryRowUnit
				for _, u := range cs.CategoryRowSheets {
					if u.SheetName == sh {
						categories = u.CategoryRowUnits
						break
					}
				}
				for ind, c := range categories {
					id := c.FileCategoryID.Row
					var nextId int
					if len(categories) > ind+1 {
						nextId = categories[ind+1].FileCategoryID.Row
					} else {
						nextId = len(rows) + 1
					}

					for _, pr := range p {
						if pr.Row > id && pr.Row < nextId {
							horo_id := c.HoroshopCategoryID
							pr.Parent.ID = &horo_id
						}
					}
					for _, pr := range pFailed {
						if pr.Row > id && pr.Row < nextId {
							horo_id := c.HoroshopCategoryID
							pr.Parent.ID = &horo_id
						}
					}
				}
			}

			allProducts = append(allProducts, p...)
			failedProducts = append(failedProducts, pFailed...)
		}
	}

	var toDelete []int
	for indx, p := range allProducts {
		if (p.Parent == nil) || (p.Parent.ID != nil && *p.Parent.ID == 0) && (p.Parent.Value != nil && *p.Parent.Value == "") {
			failedProducts = append(failedProducts, p)
			toDelete = append(toDelete, indx)
		}
	}

	for i := len(toDelete) - 1; i >= 0; i-- {
		indx := toDelete[i]
		allProducts = append(allProducts[:indx], allProducts[indx+1:]...)
	}

	return allProducts, failedProducts, nil

}
