package main

import (
	"fmt"
	"io"
	"net/http"
)

// Function to get the public IP address
func getPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Use io.ReadAll instead of ioutil.ReadAll
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
	fmt.Fprintf(w, "Your public IP address is: %s", ip)
}
