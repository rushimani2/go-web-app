package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

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

func ipPage(w http.ResponseWriter, r *http.Request) {
	ip, err := getPublicIP()
	if err != nil {
		http.Error(w, "Unable to get IP", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><head><title>Public IP</title>
		<style>
			body {
				font-family: 'Arial', sans-serif;
				background-color: #f0f4f8;
				color: #333;
				text-align: center;
				padding: 50px;
			}
			h2 {
				color: #4CAF50;
				font-size: 24px;
				border: 2px solid #4CAF50;
				padding: 10px 20px;
				border-radius: 10px;
				display: inline-block;
				background-color: #e8f5e9;
			}
		</style>
	</head>
	<body>
		<h2>Your public IP address is: <span style="color:#1976D2">%s</span></h2>
	</body></html>`, ip)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset= utf-8")
	fmt.Fprint(w, `<html><head><title>Hello</title>
		<style>
			body {
				font-family: 'Arial', sans-serif;
				background-color: #f0f4f8;
				color: #333;
				text-align: center;
				padding: 50px;
			}
			h1 {
				color: #0288D1;
			}
			p {
				font-size: 18px;
				color: #555;
			}
			a {
				color: #1976D2;
				text-decoration: none;
				font-weight: bold;
			}
			a:hover {
				color: #0288D1;
			}
		</style>
	</head>
	<body>
		<h1>👋 Hello, World!</h1>
		<p>Check your public IP at <a href="/ip">/ip</a></p>
	</body></html>`)
}

func main() {
	http.HandleFunc("/ip", ipPage)
	http.HandleFunc("/", homePage)
	fmt.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
