package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) * successRate) / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	successfulCarsPerHour := float64(productionRate) * (successRate / 100.0)
	minutesPerHour := 60.0
	successfulCarsPerMinute := successfulCarsPerHour / minutesPerHour
	return int(successfulCarsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	totalCost := 0

	numGroups := carsCount / 10
	totalCost += numGroups * 95000

	remainingCars := carsCount % 10
	totalCost += remainingCars * 10000

    return uint(totalCost)
}
