package one

import (
	"fmt"
	"strconv"
)

func parseInt(in byte) int64 {
	if s, err := strconv.ParseInt(fmt.Sprintf("%c", in), 10, 64); err == nil {
		return s
	}
	return 0
}

// DayOne solves the 'captcha' problem as follows:
//
// 1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second digit and the third digit (2) matches the fourth digit.
// 1111 produces 4 because each digit (all 1) matches the next.
// 1234 produces 0 because no digit matches the next.
// 91212129 produces 9 because the only digit that matches the next one is the last digit, 9.
// http://adventofcode.com/2017/day/1
func DayOne(in string) error {
	var sum int64
	fmt.Printf("calc sum for %s\n", in)

	// check front + back
	if in[0] == in[len(in)-1] && len(in) > 1 {
		sum += parseInt(in[0])
	}

	if len(in) == 1 {
		fmt.Printf("input is one digit, sum is 0\n")
		return nil
	}

	// check each of the others
	for i := 0; i < len(in)-1; i++ {
		if in[i] == in[i+1] {
			sum += parseInt(in[i])
		}
	}
	fmt.Printf("sum is: %d\n", sum)

	return nil
}
