package server

import (
	"fmt"
	"github.com/nickstern2002/junkyard/internal/box2048"
	"github.com/nickstern2002/junkyard/internal/cards"
	"io"
	"log"
	"net/http"
	"time"
)

func RegisterHandlers() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/weather", getWeather)
	http.HandleFunc("/time", getTime)
	http.HandleFunc("/blackjack", cards.GetBlackJack)
	http.HandleFunc("/2048", box2048.Get2048)
}

func getRoot(w http.ResponseWriter, r *http.Request) {

	// Ensure the correct content-type header is set
	w.Header().Set("Content-Type", "text/html")

	// Send HTML response, removing extra escape characters
	io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Simple Webpage</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        }
        header {
            background: #333;
            color: #fff;
            padding: 10px 0;
            text-align: center;
        }
        section {
            padding: 20px;
            text-align: center;
        }
        footer {
            background: #333;
            color: #fff;
            text-align: center;
            padding: 10px 0;
            position: fixed;
            bottom: 0;
            width: 100%;
        }
    </style>
</head>
<body>
    <header>
        <h1>Oh no what have you done</h1>
    </header>
    <section>
        <h2>Welcome to the Junkyard</h2>
        <p>Have Fun :|</p>
    </section>
    <footer>
        <p>&copy; 2024 My Simple Webpage</p>
    </footer>
</body>
</html>`)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /weather request \n")

	url := "https://wttr.in/Newark?format=%C+%t"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching weather data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	io.WriteString(w, string(body))

	fmt.Printf("Weather: %s\n", body)
}

func getTime(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /time request \n")
	io.WriteString(w, fmt.Sprintf("Time: %s\n", time.Now().String()))
}
