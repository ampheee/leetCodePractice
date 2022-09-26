package easy

func removeElement(nums []int, val int) int {
	var (
		pointer = 0
	)
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[pointer] = nums[i]
			pointer++
		}
	}
	return pointer
}
