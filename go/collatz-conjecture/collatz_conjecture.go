package collatzconjecture

import (
	"fmt"
)

/*
- If it's even, divide it by 2.
- If it's odd, multiply it by 3 and add 1.

Then, repeat these steps with the result, continuing indefinitely.

Curious, you picked number 12 to test and began the journey:

12 ➜ 6 ➜ 3 ➜ 10 ➜ 5 ➜ 16 ➜ 8 ➜ 4 ➜ 2 ➜ 1

Counting from the second number (6), it took 9 steps to reach 1, and each time the rules repeated, the number kept changing.
At first, the sequence seemed unpredictable — jumping up, down, and all over.
Yet, the conjecture claims that no matter the starting number, we'll always end at 1.
*/
func CollatzConjecture(n int) (int, error) {
	// steps is the number of steps to reach 1
	var steps int
	if n < 1 {
		return 0, fmt.Errorf("value is 0 or minus 0")
	}
	for {

		if n == 1 {
			return steps, nil
		}

		// this is even
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (n * 3) + 1
		}
		steps++
	}

}
