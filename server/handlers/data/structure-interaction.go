package data

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/Easy-Job-Developer/catalog_plus/domain/records"
)

func GetSheets(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // 10 MB
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	switch ext {
	case ".xlam", ".xlsm", ".xlsx", ".xltm", ".xltx":
	default:
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}
	parsedSheets, err := records.GetSheetList(file)
	if err != nil {
		http.Error(w, "Error executing parse", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(parsedSheets)
	if err != nil {
		http.Error(w, "Error executing serialization", http.StatusInternalServerError)
		return
	}

	jsonString := string(jsonData)

	response := jsonString
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
