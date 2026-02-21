// Package weather provides information about current weather conditions and forecasts.
package weather

var (
    // CurrentCondition Stores the current weather condition.
	CurrentCondition string 
    // CurrentLocation Stores the current location.
	CurrentLocation  string 
)

// Forecast generates a weather forecast string based on the provided city and condition.
// It updates the CurrentLocation and CurrentCondition variables with the input values,
// then returns a formatted string indicating the location and current weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
