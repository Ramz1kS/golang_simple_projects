package w_io

import (
	"errors"
	"fmt"
	"weather/colors"
	"weather/structs"
)

func PrintCity(city *structs.City) {
	fmt.Printf("%s", colors.Yellow)
	fmt.Printf("City: %s, country: %s, state: %s, lat: %f, lon: %f\n",
		city.Name, city.Country, city.State, city.Lat, city.Lon)
	fmt.Printf("%s", colors.Reset)
}

func PrintCitiesList(cities []structs.City) error {
	if len(cities) == 0 {
		return errors.New("no cities in the list :(")
	}
	for index, val := range cities {
		fmt.Printf("%d. ", index+1)
		PrintCity(&val)
	}
	return nil
}

func PrintWeather(weather *structs.WeatherInfo, unit string) {
	for _, short := range weather.Short {
		fmt.Printf("%sWeather: %s - %s%s\n", colors.Blue, short.Name, short.Description, colors.Reset)
	}
	var unitTemp, unitSpeed string
	switch unit {
	default:
		unitTemp = "K"
		unitSpeed = "m/s"
	case "metric":
		unitTemp = "°C"
		unitSpeed = "m/s"
	case "imperial":
		unitTemp = "F"
		unitSpeed = "miles/h"
	}
	fmt.Printf("%sTemperature: %.2f (Feels like: %.2f %s)%s\n", colors.Green, weather.Main.Temp, weather.Main.TempFeels, unitTemp, colors.Reset)
	fmt.Printf("%sMin Temperature: %.2f %s%s\n", colors.Yellow, weather.Main.TempMin, unitTemp,  colors.Reset)
	fmt.Printf("%sMax Temperature: %.2f %s%s\n", colors.Yellow, weather.Main.TempMax, unitTemp, colors.Reset)
	fmt.Printf("%sWind Speed: %.2f %s, Direction: %.2f°%s\n", colors.Cyan, weather.Wind.Speed, unitSpeed, weather.Wind.Direction, colors.Reset)
	fmt.Printf("%sGust: %.2f %s%s\n", colors.Gray, weather.Wind.Gust, unitSpeed, colors.Reset)
	fmt.Printf("%sCloud Coverage: %.2f%%%s\n", colors.Purple, weather.Clouds.Percantage, colors.Reset)
}