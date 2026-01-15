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

// Create map for storing all URL data (Example structure: {"40c7d58f": {ID:"40c7d58f", OriginalURL:"https://google.com", ShortURL:"40c7d58f"}})
var urlDB = make(map[string]URL)

// Generate short URL
func generateShortURL(OriginalURL string) string {
	hasher := md5.New() // Creates a new MD5 hash generator (Example: &{md5 digest object})

	hasher.Write([]byte(OriginalURL))
	// Converts URL string to bytes
	// Example OriginalURL: "https://google.com"
	// []byte → [104 116 116 112 115 58 47 47 103 111 111 103 108 101 46 99 111 109]

	fmt.Println("hasher", hasher) // Prints hash generator object (not useful for hash value)

	data := hasher.Sum(nil)
	// Generates raw hash bytes
	// Example output: [64 199 213 143 176 177 228 52 159 168 101 108 63 31 115 205]
	fmt.Println("hasher data", data)

	hash := hex.EncodeToString(data)
	// Converts raw hash bytes to hex string
	// Example: "40c7d58fb0b1e4349fa8656c3f1f73cd"
	fmt.Println("EncodeToString", hash)

	return hash[:8] // Returns first 8 characters (Example: "40c7d58f")
}

// Create URL map data
func createURL(OriginalURL string) string {
	shortURL := generateShortURL(OriginalURL) // Example input: "https://google.com" → output: "40c7d58f"
	id := shortURL                            // Using shortURL as ID

	// Store in map (urlDB)
	urlDB[id] = URL{
		ID:           id,          // Example: "40c7d58f"
		OriginalURL:  OriginalURL, // Example: "https://google.com"
		ShortURL:     shortURL,    // Example: "40c7d58f"
		CreationDate: time.Now(),  // Example: 2026-01-15 12:30:22 +0530 IST
	}
	return shortURL
}

// Get URL object from map using ID
func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found") // Example: request id "aaaa1234"
	}
	return url, nil
}

// Basic handler for root path
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// Short URL handler
func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	// Decode incoming JSON body
	// Example request:
	// { "url": "https://google.com" }
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL_ := createURL(data.URL)
	// Example output stored in shortURL_: "40c7d58f"

	// Prepare JSON response
	response := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: shortURL_, // Example: "40c7d58f"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler for redirecting to original URL
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	// Example request URL: /redirect/40c7d58f
	// id = "40c7d58f"

	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
		return
	}

	// Redirect to original URL
	http.Redirect(w, r, url.OriginalURL, http.StatusNotFound)
}

func main() {
	// Register handlers
	http.HandleFunc("/", handler)                // http://localhost:8080/
	http.HandleFunc("/shorten", ShortURLHandler) // POST → create short URL

	// Start server
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error on starting server", err)
	}
}

// // Hello World api code
// package main

// import (
// 	"crypto/md5"
// 	"encoding/hex"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// type URL struct {
// 	ID           string    `json:"id"`
// 	OriginalURL  string    `json:"original_url"`
// 	ShortURL     string    `json:"short_url"`
// 	CreationDate time.Time `json:"creation_date"`
// }

// // Create map for store the data
// var urlDB = make(map[string]URL)

// // Generate short url
// func generateShortURL(OriginalURL string) string {
// 	hasher := md5.New()               // Creates a new MD5 hash generator (Example: &{[] 0xc0000120c0})
// 	hasher.Write([]byte(OriginalURL)) // Converts URL string to bytes (Example OriginalURL: "https://google.com")
// 	// []byte("https://google.com") → [104 116 116 112 ...]

// 	fmt.Println("hasher", hasher) // Prints hash object (not final hash) [104 116 116 112 ...]

// 	data := hasher.Sum(nil)          // Generates raw hash bytes (Example output: [64 199 213 143 176 177 ...])
// 	fmt.Println("hasher data", data) // Example: [64 199 213 143 176 177 228 52 ...]

// 	hash := hex.EncodeToString(data) // Convert hash bytes to hex string
// 	// Example: "40c7d58fb0b1e4349fa8656c3f1f73cd"
// 	fmt.Println("EncodeToString", hash)

// 	return hash[:8] // Returns first 8 chars (Example: "40c7d58f")
// }

// // Create URL map data
// func createURL(OriginalURL string) string {
// 	shortURL := generateShortURL(OriginalURL)
// 	id := shortURL // Use the shortURL as the ID for simplicity

// 	// Creating map data
// 	urlDB[id] = URL{
// 		ID:           id,
// 		OriginalURL:  OriginalURL,
// 		ShortURL:     shortURL,
// 		CreationDate: time.Now(),
// 	}
// 	return shortURL
// }

// // Get URL
// func getURL(id string) (URL, error) {
// 	url, ok := urlDB[id]
// 	if !ok {
// 		return URL{}, errors.New("URL not found")
// 	}
// 	return url, nil
// }

// // handler function for request
// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World")
// }

// // Short URL handler
// func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
// 	var data struct {
// 		URL string `json:"url"`
// 	}

// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}
// 	shortURL_ := createURL(data.URL)
// 	// fmt.Fprintf(w, shortURL) // first way directly return the short url without json format

// 	// Return response in json
// 	response := struct {
// 		ShortURL string `json:"short_url"`
// 	}{ShortURL: shortURL_}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Path[len("/redirect/"):] // Take the id, comes id just after the redirect word in request url
// 	url, err := getURL(id)
// 	if err != nil {
// 		http.Error(w, "Invalid request", http.StatusNotFound)
// 	}
// 	http.Redirect(w, r, url.OriginalURL, http.StatusNotFound)
// }

// func main() {
// 	// Register the handler function to handle all requests to the root URL ("/")
// 	http.HandleFunc("/", handler)
// 	http.HandleFunc("/shorten", ShortURLHandler)

// 	// Start the HTTP server on port 8080
// 	fmt.Println("Starting server on port 8080...")
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		fmt.Println("Error on starting server", err)
// 	}
// }
