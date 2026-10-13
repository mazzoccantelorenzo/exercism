package raindrops

import "strconv"

//- is divisible by 3, add "Pling" to the result.
//- is divisible by 5, add "Plang" to the result.
//- is divisible by 7, add "Plong" to the result.
//- **is not** divisible by 3, 5, or 7, the result should be the number as a string.

func Convert(number int) string {
	var finalString string

	if (number % 3) == 0 {
		finalString = finalString + "Pling"
	}

	if (number % 5) == 0 {
		finalString = finalString + "Plang"
	}

	if (number % 7) == 0 {
		finalString = finalString + "Plong"
	}
	if finalString == "" {
		finalString = strconv.Itoa(number)
	}
	return finalString
}
