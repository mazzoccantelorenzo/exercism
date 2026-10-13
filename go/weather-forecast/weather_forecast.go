// Package weather for testing.
package weather

var (
	// CurrentCondition is the current condition.
	CurrentCondition string
	// CurrentLocation this is the current location.
	CurrentLocation string
)

// Forecast is the function for weather forecasting.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
