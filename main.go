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
		http.Error(w, "Unable to retrieve IP address", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Your public IP address is: %s", ip)
}

// Handler to render the home page
func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

// Handler to render the courses page
func coursePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/courses.html")
}

// Handler to render the about page
func aboutPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

// Handler to render the contact page
func contactPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/contact.html")
}

func main() {
	http.HandleFunc("/ip", ipPage)       // Route to display IP address
	http.HandleFunc("/home", homePage)   // Route for home page
	http.HandleFunc("/courses", coursePage) // Route for courses page
	http.HandleFunc("/about", aboutPage) // Route for about page
	http.HandleFunc("/contact", contactPage) // Route for contact page

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
