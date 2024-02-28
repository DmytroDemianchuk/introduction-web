package records

import (
	"fmt"

	"github.com/Easy-Job-Developer/catalog_plus/domain/records/models"
)

func PrintSheets(data []models.Sheet) {
	for _, sheet := range data {
		fmt.Printf("Sheet: %s\n", sheet.Title)
		printTable(sheet)
		fmt.Println()
	}
}

// printTable helper function to print each sheet's data in a tabular format.
func printTable(sheet models.Sheet) {
	// Find the max width of the column content for proper alignment
	maxWidths := make([]int, len(sheet.Columns))
	for i, col := range sheet.Columns {
		maxWidth := len(col.Title)
		for _, val := range col.Values {
			if len(val) > maxWidth {
				maxWidth = len(val)
			}
		}
		maxWidths[i] = maxWidth
	}

	// Print column headers
	for i, col := range sheet.Columns {
		fmt.Printf("%-*s ", maxWidths[i], col.Title)
	}
	fmt.Println()

	// Find the maximum number of rows in any column
	maxRows := 0
	for _, col := range sheet.Columns {
		if len(col.Values) > maxRows {
			maxRows = len(col.Values)
		}
	}

	// Print the rows of the table
	for i := 0; i < maxRows; i++ {
		for j, col := range sheet.Columns {
			val := ""
			if i < len(col.Values) {
				val = col.Values[i]
			}
			fmt.Printf("%-*s ", maxWidths[j], val)
		}
		fmt.Println()
	}
}
