package callfuncs

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"weather/structs"
	"github.com/joho/godotenv"
)

var ApiKey string

// получение API-ключа
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} 
	ApiKey = os.Getenv("API_KEY")
	if ApiKey == "" {
		log.Fatal("API key is not found")
	}
}

// главная для GET-запросов функция
func MakeAPICall(url string, target any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("bad http status code")
	}
	err = json.Unmarshal(body, &target)
	return err
}

func GetCityList(stdinReader *bufio.Reader) ([]structs.City, error) {
	fmt.Printf("Enter a city name: ")
	text, _ := stdinReader.ReadString('\n')
	text = strings.TrimSpace(text)
	cities := []structs.City{}
	err := MakeAPICall(
		fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s", text, ApiKey),
		&cities)
	return cities, err
}

func GetForecast(city *structs.City, unit string, stdinReader *bufio.Reader) (structs.WeatherInfo, error) {
	weather := structs.WeatherInfo{}
	err := MakeAPICall(
	fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&units=%s&appid=%s", city.Lat, city.Lon, unit, ApiKey),
	&weather)
	return weather, err
}