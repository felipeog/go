// Package weatherforecast provides weather forecast tools.
package weatherforecast

// CurrentCondition represents a certain condition.
var CurrentCondition string

// CurrentLocation represents a certain location.
var CurrentLocation string

// Forecast returns a string value describing the weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition

	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
