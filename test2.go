package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

func containsKey(m, k interface{}) bool {
	v := reflect.ValueOf(m).MapIndex(reflect.ValueOf(k))
	return v != reflect.Value{}
}

func task2() {
	years, max := 0, 0
	fmt.Fscan(os.Stdin, &years)
	allYearsSlice := make(map[string]int, years)
	for i := 0; i < years; i++ {
		temp := []string{}
		for k := 0; k < 3; k++ {
			var tempName string
			fmt.Fscan(os.Stdin, &tempName)
			temp = append(temp, tempName)
		}
		sort.Slice(temp, func(a, b int) bool { return temp[a] < temp[b] })
		tempString := strings.Join(temp, " ")
		allYearsSlice[tempString]++
	}
	for _, value := range allYearsSlice {
		if value > max {
			max = value
		}
	}
	fmt.Println(max)
}
