package cars

var SINGLE_CAR_COST = 10000
var TEN_CARS_COST = 95000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	//productionRate is the quantity of cars produces in an hour.
	//so productionRate = 10 means -> 10 cars every 1 hour

	//successRate is the precentage ( as integer, from 0 to 100 )
	//of cars that return without error

	return float64((float64(productionRate) * successRate) / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int((float64(productionRate)*successRate)/100) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

	// 13 cars for example should be 95000 * (10000*3)
	var groups_of_ten = int(carsCount / 10) // 13 / 10 = 1.3 so -> 1
	var single_cars = int(carsCount % 10)

	return uint(single_cars*SINGLE_CAR_COST) + uint(groups_of_ten*TEN_CARS_COST)

}
