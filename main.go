package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Function to fetch the public IP address
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

// IP Page – simple, clean design with larger font
func ipPage(w http.ResponseWriter, r *http.Request) {
	ip, err := getPublicIP()
	if err != nil {
		http.Error(w, "Unable to get IP", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><head><title>Your Public IP</title>
		<style>
			body {
				font-family: 'Arial', sans-serif;
				background-color: #f5f7fa;
				color: #333;
				text-align: center;
				padding: 100px 20px;
			}
			.container {
				display: inline-block;
				background-color: #ffffff;
				padding: 30px 40px;
				border: 2px solid #1976D2;
				border-radius: 12px;
				box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
			}
			h2 {
				font-size: 28px;
				color: #1976D2;
				margin-bottom: 10px;
			}
			span {
				font-size: 24px;
				font-weight: bold;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h2>Your public IP address is:</h2>
			<span>%s</span>
		</div>
	</body></html>`, ip)
}

// Homepage with animated symbol background
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Hello - Global Unconference</title>
	<link href="https://fonts.googleapis.com/css2?family=Montserrat:wght@400;700&display=swap" rel="stylesheet">
	<style>
		body {
			margin: 0;
			font-family: 'Montserrat', sans-serif;
			background: linear-gradient(to right, #2c3e50, #3498db);
			color: #fff;
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;
			height: 100vh;
			overflow: hidden;
			position: relative;
		}
		.symbol {
			position: absolute;
			font-size: 50px;
			color: rgba(255, 255, 255, 0.8);
			animation: moveSymbol 15s linear infinite;
		}
		@keyframes moveSymbol {
			0% { transform: translate(0, 0); }
			25% { transform: translate(calc(50vw + 50px), calc(50vh + 50px)); }
			50% { transform: translate(calc(-50vw - 50px), calc(50vh + 50px)); }
			75% { transform: translate(calc(50vw + 50px), calc(-50vh - 50px)); }
			100% { transform: translate(0, 0); }
		}
		.content {
			position: absolute;
			text-align: center;
			z-index: 10;
			padding: 20px;
			background-color: rgba(0, 123, 255, 0.7);
			border: 3px solid #0288D1;
			border-radius: 12px;
			font-size: 2rem;
			color: white;
			max-width: 80%;
			box-shadow: 0 4px 10px rgba(0, 0, 0, 0.5);
		}
		.content a {
			color: #fff;
			text-decoration: underline;
		}
	</style>
</head>
<body>
	<script>
		const symbols = ["∠", "∆", "√", "∑", "∞", "π", "≈", "≡", "⊥", "∩", "∪"];
		function createRandomSymbol() {
			const symbol = document.createElement("div");
			const randomSymbol = symbols[Math.floor(Math.random() * symbols.length)];
			const randomFontSize = Math.floor(Math.random() * 30) + 30 + "px";
			const randomX = Math.random() * 100;
			const randomY = Math.random() * 100;
			symbol.className = "symbol";
			symbol.textContent = randomSymbol;
			symbol.style.fontSize = randomFontSize;
			symbol.style.left = randomX + "%";
			symbol.style.top = randomY + "%";
			document.body.appendChild(symbol);
			symbol.style.animationDuration = (Math.random() * 10 + 10) + "s";
		}
		for (let i = 0; i < 50; i++) {
			createRandomSymbol();
		}
	</script>
	<div class="content">
		<h1>👋 Hello, World!!</h1>
		<p>Check your public IP at <a href="/ip">/ip</a></p>
	</div>
</body>
</html>`)
}

func main() {
	http.HandleFunc("/ip", ipPage)
	http.HandleFunc("/", homePage)
	fmt.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
