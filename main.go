// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Struct to parse the JSON response
type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: weather-cli <city-name>")
		os.Exit(1)
	}

	city := os.Args[1]
	apiKey := "26920b5b4d2aefaab73cf31efacc1800" // Replace with your API key

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Failed to get weather data: %s\n", string(body))
		os.Exit(1)
	}

	var weatherData WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherData)
	if err != nil {
		fmt.Printf("Error parsing data: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Weather for %s:\n", weatherData.Name)
	fmt.Printf("Temperature: %.1f°C\n", weatherData.Main.Temp)
	fmt.Printf("Humidity: %d%%\n", weatherData.Main.Humidity)
	if len(weatherData.Weather) > 0 {
		fmt.Printf("Condition: %s\n", weatherData.Weather[0].Description)
	}
}
