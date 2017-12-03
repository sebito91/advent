package two

import "fmt"

// DayTwoPartOne of advent of code 2017
func DayTwoPartOne(in [][]int) error {
	var sum int

	for _, x := range in {
		var min = x[0]
		var max = min

		for _, y := range x {
			switch {
			case y < min:
				min = y
			case y > max:
				max = y
			}
		}
		sum += (max - min)
		fmt.Printf("DEBUG -- min: %d, max: %d, diff: %d, sum: %d\n", min, max, (max - min), sum)
	}

	fmt.Printf("checksum is: %d\n", sum)
	return nil
}
