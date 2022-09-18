package solved

import "sort"

func maximumWealth(accounts [][]int) int {
	var temp []int
	for i := range accounts {
		counter := 0
		for _, val := range accounts[i] {
			counter += val
		}
		temp = append(temp, counter)
	}
	sort.Slice(temp, func(i, j int) bool { return temp[i] > temp[j] })
	return temp[0]
}
