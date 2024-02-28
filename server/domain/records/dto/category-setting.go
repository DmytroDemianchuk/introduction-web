package dto

type CategoryFullFile struct {
	Option             string `json:"option"`
	HoroshopCategoryID int    `json:"horoshop_category_id"`
}

// {
// 	"option": "category_full_file",
// 	"horoshop_category_id": 1062
// }

type CategorySheet struct {
	Option             string              `json:"option"`
	CategorySheetUnits []CategorySheetUnit `json:"sheets"`
}

type CategorySheetUnit struct {
	SheetName          string `json:"sheet"`
	HoroshopCategoryID int    `json:"horoshop_category_id"`
}

// {
// 	"option": "category_sheet",
// 	"sheets": [
// 	  {
// 		"sheet": "1. ВЕТ",
// 		"horoshop_category_id": 1063
// 	  },
// 	  {
// 		"sheet": "2. Свині",
// 		"horoshop_category_id": 1057
// 	  }
// 	]
// }

type CategoryRow struct {
	Option            string             `json:"option"`
	CategoryRowSheets []CategoryRowSheet `json:"categories"`
}

type CategoryRowSheet struct {
	SheetName        string            `json:"sheet"`
	CategoryRowUnits []CategoryRowUnit `json:"categories"`
}

type CategoryRowUnit struct {
	FileCategoryID     FileCategoryIndex `json:"excel_category_id"`
	HoroshopCategoryID int               `json:"horoshop_category_id"`
}

type FileCategoryIndex struct {
	Sheet string `json:"sheet"`
	Row   int    `json:"row"`
}

//   {
// 	"option": "category_row",
// 	"categories": [
// 	  {
// 		"sheet": "2. Свині",
// 		"categories": [
// 		  {
// 			"excel_category_id": {
// 			  "sheet": "2. Свині",
// 			  "row": 8
// 			},
// 			"horoshop_category_id": 1062
// 		  },
// 		  {
// 			"excel_category_id": {
// 			  "sheet": "2. Свині",
// 			  "row": 58
// 			},
// 			"horoshop_category_id": 1059
// 		  }
// 		]
// 	  },
// 	  {
// 		"sheet": "3. ВРХ",
// 		"categories": [
// 		  {
// 			"excel_category_id": {
// 			  "sheet": "3. ВРХ",
// 			  "row": 55
// 			},
// 			"horoshop_category_id": 1057
// 		  },
// 		  {
// 			"excel_category_id": {
// 			  "sheet": "3. ВРХ",
// 			  "row": 211
// 			},
// 			"horoshop_category_id": 1073
// 		  }
// 		]
// 	  }
// 	]
//   }

type CategoryColumn struct {
	Option string `json:"option"`
}

//   {
// 	"option": "category_column"
//   }
