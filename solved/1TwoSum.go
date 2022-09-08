package solved

func twoSum(nums []int, target int) []int {
	var temp []int
	for i := 1; i < len(nums); i++ {
		if target-nums[i] == nums[i-1] {
			temp = []int{i - 1, i}
		}
	}
	return temp
}
