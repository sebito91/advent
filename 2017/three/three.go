package three

import (
	"fmt"
	"math"
)

// DayThreePartOne of AoC 2017
func DayThreePartOne(in int64) error {
	// lowest square <= in
	sq := int64(math.Floor(math.Sqrt(float64(in))))

	// next ring of numbers in square outside of lowest
	nr := (sq + 1) * 2

	// number of spaces left of next square endpoint
	pos := nr - (in - sq*sq)

	fmt.Printf("DayThreePartOne checksum is: %d\n", (sq+1)-pos)
	return nil
}

// DayThreePartTwo of AoC 2017
func DayThreePartTwo(in int64) error {
	fmt.Printf("DayThreePartTwo checksum is: %d\n", in)
	return nil
}
