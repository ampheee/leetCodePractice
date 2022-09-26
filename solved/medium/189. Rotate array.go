package medium

func rotate(nums []int, k int) {
	var temp int
	for j := 0; j < k; j++ {
		temp = nums[len(nums)-1]
		for i := len(nums) - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = temp
	}
}

func rotate2(nums []int, k int) {
	n := len(nums)
	k = k % n
	result := append(nums[n-k:], nums[:n-k]...)
	for i := 0; i < len(nums); i++ {
		nums[i] = result[i]
	}
}
