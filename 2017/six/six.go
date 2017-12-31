package six

import "fmt"

// rotateKeys is a helper to redistribute keys across the list
func rotateKeys() {
	var id int
	var m = keys[id]

	d := len(keys)

	for k, v := range keys {
		if v > m {
			m = v
			id = k
		}
	}

	keys[id] = 0

	for {
		if m == 0 {
			return
		}

		id = ((id + 1) % d)
		keys[id]++
		m--
	}
}

// DaySixPartOne of AoC 17
func DaySixPartOne() error {
	var count int
	combos := make(map[string]bool)

	for {
		if _, ok := combos[fmt.Sprint(keys)]; ok {
			fmt.Printf("DaySixPartOne codec found at: %d\n", count)
			return nil
		}

		combos[fmt.Sprint(keys)] = true
		count++
		rotateKeys()
	}
}

// DaySixPartTwo of AoC 17
func DaySixPartTwo() error {
	return nil
}

var keys = []int{10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6}

//var keys = []int{0, 2, 7, 0}
