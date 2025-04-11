package main

import (
	"fmt"
	"io/ioutil"
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

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}

// Handler to display public IP address
func ipPage(w http.ResponseWriter, r *http.Request) {
	ip, err := getPublicIP()
	if err != nil {
		log.Printf("Error retrieving IP: %v", err)
		http.Error(w, "Unable to retrieve IP address", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Your public IP address is: %s", ip)
}

// Handler to render a basic home page
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to the Home Page</h1>")
}

func main() {
	http.HandleFunc("/", ipPage)
	http.HandleFunc("/home", homePage)

	log.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
