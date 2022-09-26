package main

import (
	"fmt"
	"os"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return a * -1
	} else {
		return a
	}
}

func task3() {
	var days, totalsum int
	fmt.Fscan(os.Stdin, &days)
	minus, plus := make([]int, 0), make([]int, 0)
	for i := 0; i < days; i++ {
		var day int
		fmt.Fscan(os.Stdin, &day)
		if i%2 == 0 {
			plus = append(plus, day)
			totalsum += day
		} else {
			minus = append(minus, day*-1)
			totalsum -= day
		}
	}
	sort.Slice(minus, func(i, j int) bool { return minus[i] < minus[j] })
	sort.Slice(plus, func(i, j int) bool { return plus[i] < plus[j] })
	if plus[0] < abs(minus[0]) {
		totalsum += 2 * (abs(minus[0]) - plus[0])
	}
	fmt.Print(totalsum)
}
