package horoshop

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func PostProductsDataHandler(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Token cookie not found", http.StatusUnauthorized)
		return
	}
	token := tokenCookie.Value

	domainCookie, err := r.Cookie("domain")
	if err != nil {
		http.Error(w, "Domain cookie not found", http.StatusUnauthorized)
		return
	}
	domain := domainCookie.Value

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	catalogRequest := struct {
		Token    string          `json:"token"`
		Products json.RawMessage `json:"products"`
	}{
		Token:    token,
		Products: bodyBytes,
	}

	jsonData, err := json.Marshal(catalogRequest)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", domain+"/api/catalog/import/", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error sending request", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(responseBody)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
