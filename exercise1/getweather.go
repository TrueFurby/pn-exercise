package main

import (
	"fmt"
	"log"
	"os"

	owm "github.com/briandowns/openweathermap"
)

var (
	apiKey   = os.Getenv("OPENWEATHER_API_KEY")
	cityName = os.Getenv("CITY_NAME")
)

func main() {
	w, err := owm.NewCurrent("C", "EN", apiKey)
	if err != nil {
		log.Fatalln(err)
	}
	if err := w.CurrentByName(cityName); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("source=%s, city=%q, description=%q, temp=%.1f, humidity=%d\n",
		"openweathermap", w.Name, w.Weather[0].Description, w.Main.Temp, w.Main.Humidity)
}
