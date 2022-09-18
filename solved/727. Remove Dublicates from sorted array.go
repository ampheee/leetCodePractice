package solved

func removeDuplicates(nums []int) int {
	var (
		counter = 1
		pointer = 0
	)
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[pointer] {
			pointer++
			counter++
			nums[pointer] = nums[i]
		}
	}
	return counter
}
