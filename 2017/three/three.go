package three

import (
	"fmt"
	"math"
)

const (
	up int = iota
	down
	left
	right
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

// printTable is here for debugging
// NOTE: this table is inverted meaning (0,0) is the top-left corner
// instead of bottom-right! All coordinates are based from the top-left
//
// To check this out for small numbers (~10000) enable the printTable functions
//func printTable(t [][]int) {
//	for _, y := range t {
//		fmt.Printf("[ ")
//		for _, v := range y {
//			fmt.Printf("%5d ", v)
//		}
//		fmt.Printf("]\n")
//	}
//}

// DayThreePartTwo of AoC 2017
func DayThreePartTwo(in int) error {
	sq := int(math.Ceil(math.Sqrt(float64(in))))
	var table = make([][]int, sq)
	var sum int

	for i := 0; i < sq; i++ {
		table[i] = make([]int, sq)
	}

	x := sq / 2
	y := sq / 2
	ring, step := 1, 1
	table[y][x] = 1 // center
	//	printTable(table)

	dir := right

	for {
		if sum > in {
			break
		}

		switch dir {
		case right:
			// sum += current + one-up + one-up & one-right + one-up & two-right
			sum = table[y][x] + table[y-1][x] + table[y-1][x+1] + table[y-1][x+2]
			x++
			table[y][x] = sum

			if table[y-1][x] == 0 {
				ring++
				dir = up
			}

		case left:
			// sum += current + one-down + one-left & one-down + two-left & one-down
			sum = table[y][x] + table[y+1][x] + table[y+1][x-1] + table[y+1][x-2]
			x--
			table[y][x] = sum

			if table[y+1][x] == 0 {
				dir = down
			}

		case up:
			// sum += current + one-left + one-left & one-up + one-left & two-up
			sum = table[y][x] + table[y][x-1] + table[y-1][x-1] + table[y-2][x-1]
			y--
			table[y][x] = sum

			if table[y][x-1] == 0 {
				dir = left
			}

		case down:
			// sum += current + one-right + one-right & one-down + one-right & two-down
			sum = table[y][x] + table[y][x+1] + table[y+1][x+1] + table[y+2][x+1]
			y++
			table[y][x] = sum

			if table[y][x+1] == 0 {
				dir = right
			}
		}
		step++
		//		printTable(table)
		//		fmt.Printf("sum is: %d, y: %d, x: %d, dir: %d, ring: %d, step: %d\n", sum, y, x, dir, ring, step)
	}

	fmt.Printf("DayThreePartTwo checksum is: %d, %d\n", in, sum)

	return nil
}
