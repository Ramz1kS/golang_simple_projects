# Weather app. 
## Uses OpenWeatherMap API
You need to get a free OpenWeatherMap API key, create .env file and type
```
API_KEY=(your key)
```
in order to use this little app

How it works?
```
> go run main.go
Enter a city name: Moscow
1. City: Moscow, country: RU, state: Moscow, lat: 55.750446, lon: 37.617494
2. City: Moscow, country: US, state: Idaho, lat: 46.732388, lon: -117.000165
3. City: Moscow, country: US, state: Maine, lat: 45.071096, lon: -69.891586
4. City: Moscow, country: US, state: Tennessee, lat: 35.061998, lon: -89.403961
5. City: Moscow, country: US, state: Maryland, lat: 39.543701, lon: -79.005027
Please select a city you want to get a weather forecast for: 1
MEASUREMENT METHODS
1. Standard
2. Metric
3. Imperial
Please select a measurement method: 2
Weather: Clouds - overcast clouds
Temperature: 19.16 (Feels like: 18.31 °C)
Min Temperature: 17.45 °C
Max Temperature: 20.98 °C
Wind Speed: 5.45 m/s, Direction: 275.00°
Gust: 7.23 m/s
Cloud Coverage: 85.00%
```