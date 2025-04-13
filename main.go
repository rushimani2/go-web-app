package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Function to get the public IP address
func getPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}

// Handler to display public IP address
func ipPage(w http.ResponseWriter, r *http.Request) {
	ip, err := getPublicIP()
	if err != nil {
		http.Error(w, "Unable to retrieve IP address", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "Your public IP address is: %s", ip)
}

// Home page handler
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to the Go Web App3</h1>")
}

func main() {
	http.HandleFunc("/ip", ipPage)
	http.HandleFunc("/home", homePage)
	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
