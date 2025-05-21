package main

import (
	"bufio"
	"log"
	"os"

	w_io "weather/io"
	callfuncs "weather/call_funcs"
)

func main() {
	stdinReader := bufio.NewReader(os.Stdin)
	availableCities, err := callfuncs.GetCityList(stdinReader)
	if err != nil {
		log.Fatal(err)
	}
	err = w_io.PrintCitiesList(availableCities)
	if err != nil {
		log.Fatal(err)
	}
	choice, err := w_io.GetChoice("Please select a city you want to get a weather forecast for: ", len(availableCities), stdinReader)
	if err != nil {
		log.Fatal(err)
	}
	unit, err := w_io.GetMeasurementMethod(stdinReader)
	if err != nil {
		log.Fatal(err)
	}
	weather, err := callfuncs.GetForecast(&availableCities[choice - 1], unit, stdinReader)
	if err != nil {
		log.Fatal(err)
	}
	w_io.PrintWeather(&weather, unit)
}