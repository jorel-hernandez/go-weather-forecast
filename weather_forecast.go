// Package weather: weather forecasting manager program of weather station.
package weather

//CurrentCondition variable to store current condition.
var CurrentCondition string
//CurrentLocation variable to store current location.
var CurrentLocation string

//Forecast function to get forecast weather in a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
