package medium

func peakIndexInMountainArray(arr []int) int {
	var (
		left  = 0
		right = len(arr) - 1
	)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
