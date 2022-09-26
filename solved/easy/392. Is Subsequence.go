package easy

import (
	"fmt"
	"strconv"
	"strings"
)

func isSubsequence(s string, t string) bool {
	runes := make(map[uint8]int, 0)
	var (
		temp = ""
		j    = 0
	)
	for _, val := range s {
		runes[uint8(val)]++
	}
	for j <= len(t)-1 {
		if val, ok := runes[t[j]]; ok && val != 0 {
			temp += string(t[j])
			runes[t[j]]--
		}
		j++
	}
	fmt.Println(runes, temp)
	if temp == s {
		return true
	}
	return false
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		res       []int
		i         int
		j         int
		minLenght int
	)
	if len(nums1) < len(nums2) {
		minLenght = len(nums1)
	} else {
		minLenght = len(nums2)
	}
	for i < minLenght {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			j++
		} else if nums1[i] > nums2[j] {
			res = append(res, nums2[j])
			i++
		} else {
			res = append(res, nums1[i], nums2[j])
		}
	}

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var ptr1, ptr2, ptr3 int = m - 1, n - 1, m + n - 1
	for ; ptr1 >= 0 && ptr2 >= 0; ptr3-- {
		if nums2[ptr2] >= nums1[ptr1] {
			nums1[ptr3] = nums2[ptr2]
			ptr2--
		} else {
			nums1[ptr3] = nums1[ptr1]
			ptr1--
		}
	}

	if ptr2 >= 0 {
		copy(nums1[:ptr3+1], nums2[:ptr2+1])
	}
}

* Complete the 'ModifyString' function below and add imports if needed.
*
* The function is expected to return a STRING.
* The function accepts STRING str as parameter.
*/

func ModifyString(str string) string {
	var (
		temp   string
		i      int
		digits = map[uint8]int{
			'1': 1,
			'2': 2,
			'3': 3,
			'4': 4,
			'5': 5,
			'6': 6,
			'7': 7,
			'8': 8,
			'9': 9,
		}
	)
	temp = strings.Join(strings.Fields(str), " ")
	for i < len(temp) {
		if _, ok := digits[temp[i]]; ok {
			temp = strings.ReplaceAll(temp, string(temp[i]), "0")
		}
		i++
	}
	temp = strings.ReplaceAll(temp, "0", "")
	r := []rune(temp)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}


func miniMaxSum(arr []int32) {
	var (
		minSum int32
		maxSum int32
	)
	for i := 0; i < 5; i++ {
	var temp int32
	for j := 0; j < i; j++ {
	temp += arr[j]
	}
	for j := i; j < 5; j++ {
	temp += arr[j]
	}
	fmt.Println(temp)
	}
}

func timeConversion(s string) string {
	var res string
	if s[len(s) - 2:] == "PM" {
		temp, _ := strconv.Atoi(s[:2])
		if temp == 12 {
			temp += 11
		} else {
			temp += 12
		}
		res = strconv.Itoa(temp) + s[2:len(s) - 2]
	} else {
		res = s[:len(s) - 2]
	}
	return res
}
func Unique(data interface{}) {data.Tru}

func superReducedString(s string) string {
	var res string
	res = func
}