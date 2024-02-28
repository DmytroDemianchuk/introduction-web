package horoshop

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Easy-Job-Developer/catalog_plus/domain/horoshop"
)

func GetProductsDataHandler(w http.ResponseWriter, r *http.Request) {
	mockStr := "input"
	mockS := &mockStr
	mockInt := 0
	mockI := &mockInt
	mockFloat := 0.0
	mockF := &mockFloat
	response := horoshop.Product{
		Article: mockS,
		Title: &horoshop.LocalizedString{
			RU: mockS,
			UA: mockS,
		},
		H1Title: &horoshop.LocalizedString{
			RU: mockS,
			UA: mockS,
		},
		SEOTitle: &horoshop.LocalizedString{
			RU: mockS,
			UA: mockS,
		},
		Description: &horoshop.LocalizedString{
			RU: mockS,
			UA: mockS,
		},
		SEODescription: &horoshop.LocalizedString{
			RU: mockS,
			UA: mockS,
		},
		MarketplaceDescription: &horoshop.LocalizedString{
			RU: mockS,
			UA: mockS,
		},
		GTIN:            mockS,
		MPN:             mockS,
		Popularity:      mockI,
		GuaranteeShop:   mockS,
		GuaranteeLength: mockI,
		Parent: &horoshop.ParentID{
			ID:    mockI,
			Value: mockS,
		},
		Price:    mockF,
		PriceOld: mockF,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func GetCategories(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Token cookie is required", http.StatusUnauthorized)
		return
	}
	token := tokenCookie.Value

	domainCookie, err := r.Cookie("domain")
	if err != nil {
		http.Error(w, "Domain cookie is required", http.StatusUnauthorized)
		return
	}
	domain := domainCookie.Value

	jsonData := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}
	jsonBody, err := json.Marshal(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", domain+"/api/pages/export/", bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	var response struct {
		Status   string `json:"status"`
		Response struct {
			Pages []horoshop.Page `json:"pages"`
		} `json:"response"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filteredPages := make([]horoshop.Page, 0)
	for _, page := range response.Response.Pages {
		if page.Parent != 1 && page.Parent != 0 {
			filteredPages = append(filteredPages, page)
		}
	}

	jsonDataPages, err := json.Marshal(filteredPages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonDataPages)
}
