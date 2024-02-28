package horoshop

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Easy-Job-Developer/catalog_plus/domain/horoshop"
)

func HoroshopAuthHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var creds horoshop.Auth
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var domain horoshop.Domain
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest("POST", domain.URL+"/api/auth/", bytes.NewBuffer(jsonData))
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
			Token string `json:"token"`
		} `json:"response"`
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare the response
	domain.Token = response.Response.Token
	jsonDataDomain, err := json.Marshal(domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonDataDomain)
}
