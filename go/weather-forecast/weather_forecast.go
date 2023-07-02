// Package weather provides information about the weather.
package weather

// CurrentCondition provides info on the current condition.
var CurrentCondition string
// CurrentLocation provides info on the current location.
var CurrentLocation string

// Forecast provides info on the current weather condition for the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
