package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	testSlice1 := []int{1, 2, 3}
	testSlice2 := []int{67, 4, 5}
	fmt.Println(mergeTwoSlices(testSlice1, testSlice2))
}

func mergeTwoSlices(test1, test2 []int) []int {
	return append(test1, test2[0:]...)
}

func romanToIntTest(s string) int {
	var (
		intSum int
		prev   rune
	)
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for _, val := range s {
		intSum += m[val]
		if m[prev] < m[val] {
			intSum -= m[prev] * 2
		}
		prev = val
	}
	return intSum
}

func fizzBuzzTest(n int) []string {
	var answer []string
	for i := 1; i <= n; i++ {
		if i%5 == 0 && i%3 == 0 {
			answer = append(answer, "FizzBuzz")
		} else if i%5 == 0 {
			answer = append(answer, "Buzz")
		} else if i%3 == 0 {
			answer = append(answer, "Fizz")
		} else {
			answer = append(answer, strconv.Itoa(i))
		}
	}
	return answer
}

func middleNodeTest(head *ListNode) *ListNode {
	s, f := head, head
	for f != nil && f.Next != nil {
		s, f = s.Next, f.Next.Next
	}
	return s
}

func kWeakestRowsTest(mat [][]int, k int) []int {
	var sums = make([][2]int, len(mat))
	for index, str := range mat {
		for _, val := range str {
			sums[index][1] += val
		}
		sums[index][0] = index
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			if sums[i][1] < sums[j][1] || (sums[i][1] == sums[j][1] && sums[i][0] < sums[j][0]) {
				sums[i], sums[j] = sums[j], sums[i]
			}
		}
	}
	var out = make([]int, k)
	for i := 0; i < k; i++ {
		out[i] = sums[i][0]
	}
	return out
}
