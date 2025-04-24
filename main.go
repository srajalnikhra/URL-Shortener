package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:8]

}

func createURL(original_url string) string {
	shortURL := generateShortURL(original_url)
	id := shortURL
	urlDB[id] = URL{
		ID: id,
		OriginalURL: original_url,
		ShortURL: shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func RootPageURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	shortURL_ := createURL(data.URL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL_}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	fmt.Println("Srajal...")
	OriginalURL := "https://github.com/Prince-1501"
	generateShortURL(OriginalURL)

	http.HandleFunc("/", RootPageURL)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)

	fmt.Println(("Starting server on port 3000..."))
	err := http.ListenAndServe(":3000", nil)
	if err != nil{
		fmt.Println("error on starting server:", err)
	}
}
