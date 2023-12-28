package weatherStation

import (
	"fmt"
	"math"
)

type WeatherInfo struct {
	Temperature float64
	Humidity    float64
	Pressure    float64
}

type Display struct {
}

func NewDisplay() *Display {
	return &Display{}
}

/*
Метод Update сделан приватным, чтобы ограничить возможность его вызова напрямую
Классу Observable он будет доступен все равно, т.к. в интерфейсе IObserver он
остается публичным
*/
func (d *Display) Update(data *WeatherInfo) {
	fmt.Printf("Current Temp %.2f\n", data.Temperature)
	fmt.Printf("Current Hum %.2f\n", data.Humidity)
	fmt.Printf("Current Pressure %.2f\n", data.Pressure)
	fmt.Println("----------------")
}

type StatsDisplay struct {
	minTemperature float64
	maxTemperature float64
	accTemperature float64
	countAcc       int
}

func NewStatsDisplay() *StatsDisplay {
	return &StatsDisplay{
		minTemperature: math.Inf(1),
		maxTemperature: math.Inf(-1),
		accTemperature: 0,
		countAcc:       0}
}

/*
Метод Update сделан приватным, чтобы ограничить возможность его вызова напрямую
Классу Observable он будет доступен все равно, т.к. в интерфейсе IObserver он
остается публичным
*/
func (s *StatsDisplay) Update(data *WeatherInfo) {
	if data.Temperature < s.minTemperature {
		s.minTemperature = data.Temperature
	}
	if data.Temperature > s.maxTemperature {
		s.maxTemperature = data.Temperature
	}
	s.accTemperature += data.Temperature
	s.countAcc++

	fmt.Printf("Max Temp %.2f\n", s.maxTemperature)
	fmt.Printf("Min Temp %.2f\n", s.minTemperature)
	fmt.Printf("Average Temp %.2f\n", s.accTemperature/float64(s.countAcc))
	fmt.Println("----------------")
}

type WeatherData struct {
	Observable[WeatherInfo]
	Temperature float64
	Humidity    float64
	Pressure    float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{Observable: *NewObservable[WeatherInfo](), Temperature: 0, Humidity: 0, Pressure: 760}
}

// Температура в градусах Цельсия
func (w *WeatherData) GetTemperature() float64 {
	return w.Temperature
}

// Относительная влажность (0...100)
func (w *WeatherData) GetHumidity() float64 {
	return w.Humidity
}

// Атмосферное давление (в мм.рт.ст)
func (w *WeatherData) GetPressure() float64 {
	return w.Pressure
}

func (w *WeatherData) MeasurementsChanged() {
	w.NotifyObservers(w.GetChangedData())
}

func (w *WeatherData) SetMeasurements(temp, humidity, pressure float64) {
	data := WeatherInfo{
		Temperature: temp,
		Humidity:    humidity,
		Pressure:    pressure,
	}
	w.NotifyObservers(&data)
}

func (w *WeatherData) GetChangedData() *WeatherInfo {
	weatherInfo := WeatherInfo{Temperature: w.GetTemperature(), Humidity: w.GetHumidity(), Pressure: w.GetPressure()}
	return &weatherInfo
}
