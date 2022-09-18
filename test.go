package main

import (
	"fmt"
	"os"
	"sort"
)

func task1() {
	var first, second = make([]int, 4), make([]int, 4)
	for i := 0; i < len(first); i++ {
		fmt.Fscan(os.Stdin, &first[i])
	}
	for i := 0; i < len(second); i++ {
		fmt.Fscan(os.Stdin, &second[i])
	}
	var Xarray = []int{first[0], first[2], second[0], second[2]}
	var Yarray = []int{first[1], first[3], second[1], second[3]}
	sort.Slice(Xarray, func(a, b int) bool { return Xarray[a] < Xarray[b] })
	sort.Slice(Yarray, func(a, b int) bool { return Yarray[a] < Yarray[b] })
	x, y := Xarray[3]-Xarray[0], Yarray[3]-Yarray[0]
	if x < y {
		fmt.Print(y * y)
	} else {
		fmt.Print(x * x)
	}
}
