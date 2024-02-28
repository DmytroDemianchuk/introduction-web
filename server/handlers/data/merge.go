package data

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Easy-Job-Developer/catalog_plus/domain/horoshop"
	"github.com/Easy-Job-Developer/catalog_plus/domain/records"
	"github.com/Easy-Job-Developer/catalog_plus/domain/records/dto"
)

func MergeFilesHandler(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // 10 MB
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Panicln(err.Error())
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	columnPairsJSON := r.FormValue("columnPairs")
	var columnPairs []horoshop.ProductSheetsMap

	err := json.Unmarshal([]byte(columnPairsJSON), &columnPairs)
	if err != nil {
		log.Panicln(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	inputFile, _, err := r.FormFile("inputFile")
	if err != nil {
		log.Panicln(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer inputFile.Close()

	settingInputJSON := r.FormValue("settingInput")
	type PreParse struct {
		Option string `json:"option"`
	}

	var pp PreParse
	if err := json.Unmarshal([]byte(settingInputJSON), &pp); err != nil {
		log.Println("Error parsing JSON:", err)
		http.Error(w, "Error executing parse", http.StatusInternalServerError)
		return
	}

	var setting interface{}
	switch pp.Option {
	case "category_full_file":
		setting = &dto.CategoryFullFile{}
	case "category_sheet":
		setting = &dto.CategorySheet{}
	case "category_row":
		setting = &dto.CategoryRow{}
	case "category_column":
		setting = &dto.CategoryColumn{}
	default:
		log.Println("Unknown option")
		return
	}

	selectedSheetsJSON := r.FormValue("selectedSheets")
	var selectedSheets []string
	err = json.Unmarshal([]byte(selectedSheetsJSON), &selectedSheets)
	if err != nil {
		log.Println("Error parsing JSON to specific type:", err)
		http.Error(w, "Error executing parse", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal([]byte(settingInputJSON), &setting); err != nil {
		log.Println("Error parsing JSON to specific type:", err)
		http.Error(w, "Error executing parse", http.StatusInternalServerError)
		return
	}

	products, failedProducts, err := records.ParseExcelData(inputFile, columnPairs, selectedSheets, setting)
	if err != nil {
		log.Println("Error during ParseExcelData:", err)
		http.Error(w, "Error executing parse", http.StatusInternalServerError)
		return
	}

	result := struct {
		Products        []horoshop.Product `json:"products"`
		FaildedProducts []horoshop.Product `josn:"failedProducts"`
	}{
		Products:        products,
		FaildedProducts: failedProducts,
	}

	productsJSON, _ := json.Marshal(result)
	w.Write(productsJSON)

}
