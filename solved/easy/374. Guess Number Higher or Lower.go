package easy

//
//func guessNumber(n int) int {
//	var (
//		left  = 1
//		right = n
//		res   int
//	)
//	for left <= right {
//		mid := (left + right) / 2
//		if guess(mid) == -1 {
//			right = mid - 1
//		} else if guess(mid) == 1 {
//			left = mid + 1
//		} else {
//			res = mid
//			break
//		}
//	}
//	return res
//}

func sortedSquares(nums []int) []int {
	var (
		pointer = 0
	)
	for i := range nums {
		nums[i] = nums[i] * nums[i]
		if nums[pointer] > nums[i] {
			Swap(&nums[pointer], &nums[i])
			pointer++
		}
	}
	if nums[0] >= nums[1] {

	}
	return nums
}

func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
