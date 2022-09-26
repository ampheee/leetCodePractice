package easy

func pivotIndex(nums []int) int {
	sum := 0
	for each := 0; each < len(nums); each++ {
		sum += nums[each]
	}
	left := 0
	for i, value := range nums {
		if left == (sum - left - value) {
			return i
		}
		left += value
	}
	return -1
}
