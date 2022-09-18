package solved

func runningSum(nums []int) []int {
	var tempSum int
	for i, val := range nums {
		tempSum += val
		nums[i] = tempSum
	}
	return nums
}
